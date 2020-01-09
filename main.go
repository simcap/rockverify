//go:generate abigen --sol rockverify.sol --pkg main --out abi.go

package main

import (
	"crypto/sha256"
	"errors"
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"syscall"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rocksideio/rockside-sdk-go"
)

var (
	rocksideAPIKey    = os.Getenv("ROCKSIDE_API_KEY")
	rocksideAPIURL    = "https://api.rockside.io"
	contractAddress   = common.HexToAddress("0xa2c13b62d34613191578f901dde757c1b86f6484")
	testnetFlag       = flag.Bool("testnet", true, "Use testnet (Ropsten) instead of mainnet")
	basicAuthUsername = flag.String("basic-auth-username", "", "Username for HTTP basic authentication, password then asked by prompt")
	network           = rockside.Mainnet
)

func main() {
	flag.Parse()

	if *testnetFlag {
		network = rockside.Testnet
	}

	switch flag.Arg(0) {
	case "":
		flag.PrintDefaults()
	case "register":
		url, err := normalizeURL(flag.Arg(1))
		exitOn(err)

		exitOn(registerURL(url))
	default:
		url, err := normalizeURL(flag.Arg(0))
		exitOn(err)

		exitOn(downloadContent(url))
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

	client, err := rockside.NewClient(rocksideAPIURL, rocksideAPIKey)
	if err != nil {
		return err
	}
	client.SetNetwork(rockside.Network(network))

	transaction := rockside.Transaction{
		From: "0x4b706a10eb18EEd7f5d5faf756984f7cAE85e713",
		To:   "0xa2c13b62d34613191578f901dde757c1b86f6484",
		Data: fmt.Sprintf("0x%x", call),
	}

	printInfo("performing blockchain transaction to register fingerprints")
	if _, _, err := client.Transaction.Send(transaction); err != nil {
		printError("cannot perform transaction: %s", err)
		return err
	}
	printInfo("URL and content fingerprints have been registered successfully to the blockchain. Thanks Rockside!")

	return nil
}

func downloadContent(u *url.URL) error {
	urlShasum := sha256.Sum256([]byte(u.String()))
	rockverify, err := NewRockVerifyCaller(contractAddress, rpcClient())
	if err != nil {
		return err
	}

	printInfo("reading blockchain entry for '%s'", u)
	contentShasum, err := rockverify.Lookup(nil, urlShasum)
	if err != nil {
		printError("cannot read blockchain entry: %s", err)
		return err
	}

	if contentShasum == [32]byte{} {
		printWarn("nothing registered for '%s'", u)
		return nil
	}

	file, err := ioutil.TempFile("", fmt.Sprintf("rockverify-*"))
	if err != nil {
		return err
	}
	defer file.Close()

	actualShasum, err := shasumContentAt(u, file)
	if err != nil {
		return err
	}
	printInfo("content downloaded to local file %s", file.Name())

	if actualShasum != contentShasum {
		printError("mismatch: actual content fingerprint != registered blockchain fingerprint")
		if err := os.Remove(file.Name()); err != nil {
			printError("cannot remove local file at %s", file.Name())
		}
		printInfo("removed downloaded file")
	} else {
		printInfo("fingerprint of downloaded content matches registered fingerprint on blockchain")
		printInfo("verification is successful. Thanks Rockside!")
	}

	return nil
}

func rpcClient() *ethclient.Client {
	if rocksideAPIKey == "" {
		exitOn(errors.New("missing ROCKSIDE_API_KEY env variable to build RPC client"))
	}
	client, err := ethclient.Dial(fmt.Sprintf("https://api.rockside.io/ethereum/ropsten/jsonrpc?apikey=%s", rocksideAPIKey))
	if err != nil {
		exitOn(err)
	}

	return client
}

func shasumContentAt(url *url.URL, writers ...io.Writer) ([32]byte, error) {
	var shasum [32]byte

	writer := ioutil.Discard
	if len(writers) > 0 {
		writer = writers[0]
	}

	resp, err := httpGet(url.String())
	if err != nil {
		return shasum, fmt.Errorf("cannot GET %s: %s", url, err)
	}
	defer resp.Body.Close()

	hasher := sha256.New()
	if _, err := io.Copy(io.MultiWriter(writer, hasher), resp.Body); err != nil {
		return shasum, err
	}

	copy(shasum[:], hasher.Sum(nil))
	return shasum, nil
}

func httpGet(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	if *basicAuthUsername != "" {
		fmt.Fprintf(os.Stderr, "Enter Password: ")
		password, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return nil, fmt.Errorf("cannot read password: %s", err)
		}
		fmt.Fprintln(os.Stderr)
		req.SetBasicAuth(*basicAuthUsername, string(password))
	}

	printInfo("computing fingerprint of content found at '%s'", url)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot request %s: %s", url, err)
	}
	if code := resp.StatusCode; code < http.StatusOK || code > 299 {
		return resp, fmt.Errorf("cannot request %s, got %d", url, code)
	}

	return resp, nil
}

func normalizeURL(u string) (*url.URL, error) {
	parsed, err := url.Parse(u)
	if err != nil {
		return parsed, err
	}
	printInfo("normalizing given URL")
	parsed.Path = strings.TrimSuffix(parsed.Path, "/")
	return parsed, nil
}

func exitOn(err error) {
	if err != nil {
		printError(err.Error())
		os.Exit(1)
	}
}

func printError(s string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, "\033[31m[-]\033[m %s\n", fmt.Sprintf(s, a...))
}

func printWarn(s string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, "\033[33m[?]\033[m %s\n", fmt.Sprintf(s, a...))
}

func printInfo(s string, a ...interface{}) {
	fmt.Fprintf(os.Stdout, "\033[32m[+]\033[m %s\n", fmt.Sprintf(s, a...))
}
