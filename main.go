//go:generate abigen --sol rockverify.sol --pkg main --out abi.go

package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rocksideio/rockside-sdk-go"
)

var (
	rocksideAPIKey  = os.Getenv("ROCKSIDE_API_KEY")
	contractAddress = common.HexToAddress("0xa2c13b62d34613191578f901dde757c1b86f6484")
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	switch flag.Arg(0) {
	case "":
		flag.PrintDefaults()
	case "register":
		url := normalizeURL(flag.Arg(1))
		if err := registerURL(url); err != nil {
			log.Fatal(err)
		}
	default:
		url := normalizeURL(flag.Arg(1))
		result, err := lookupURL(url)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("content shasum '%x' for url %s", result, url)
	}
}

func registerURL(url *url.URL) error {
	contentShasum, err := shasumContentAt(url)
	if err != nil {
		return err
	}

	contractABI, err := abi.JSON(strings.NewReader(RockVerifyABI))
	if err != nil {
		return err
	}

	urlShasum := sha256.Sum256([]byte(url.String()))
	call, err := contractABI.Pack("register", urlShasum, contentShasum)
	if err != nil {
		return err
	}
	log.Printf("call %x", call)

	client, err := rockside.New(rocksideAPIKey)
	if err != nil {
		return err
	}

	identities, _, err := client.Identities.List(rockside.Ropsten)
	if err != nil {
		return err
	}

	log.Println(identities)
	log.Printf("registering %s successful, transaction=%s", url.String(), "")

	return nil
}

func lookupURL(u *url.URL) ([32]byte, error) {
	urlShasum := sha256.Sum256([]byte(u.String()))
	rockverify, err := NewRockVerifyCaller(contractAddress, rpcClient())
	if err != nil {
		return [32]byte{}, err
	}
	return rockverify.Lookup(nil, urlShasum)
}

func rpcClient() *ethclient.Client {
	if rocksideAPIKey == "" {
		log.Fatal("missing ROCKSIDE_API_KEY env variable to build RPC client")
	}
	client, err := ethclient.Dial(fmt.Sprintf("https://api.rockside.io/ethereum/ropsten/jsonrpc?apikey=%s", rocksideAPIKey))
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func shasumContentAt(url *url.URL) ([32]byte, error) {
	var shasum [32]byte

	resp, err := http.Get(url.String())
	if err != nil {
		return shasum, fmt.Errorf("cannot GET %s: %s", url, err)
	}
	defer resp.Body.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, resp.Body); err != nil {
		return shasum, err
	}

	copy(shasum[:], hasher.Sum(nil))
	return shasum, nil
}

func normalizeURL(u string) *url.URL {
	parsed, err := url.Parse(u)
	if err != nil {
		log.Fatal(err)
	}
	return parsed
}
