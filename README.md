# RockVerify

RockVerify is a command line interface that leverages the Ethereum blockchain to add an extra layer of verification 
when downloading files from URL.

To interact easily with the blockchain, we use [Rockside](https://www.rockside.io).

Example:

* register a downloadable item on the blockchain: `rockverify register https://example.com/releases/1.5.4/binary` 
* send this command to a third party: `rockverify https://example.com/releases/1.5.4/binary`
* automatically the item is downloaded with the extra verification step performed, making the file available locally on success 
  
The process when we get an item from an URL is simple: 

1. Validates URL format and normalizes it
2. Lookup on the blockchain the entry for this URL using the URL shasum. If no entry abort.
3. Get the data from the entry on the blockchain: URL (will be the same), shasum of file
4. Download file from the URL and calculates its shasum
5. Compare shasum register in the blockchain with shasum of downloaded file.
6. If no shasum mismatch, the file will be available locally.

## File and fingerprints, problem?

When we advertise on a website a file to be downloaded and the corresponding file fingerprint to be verified after download,
we try to mitigate an in transit corruption of the file (it being malicious or not). 

Unfortunately, the fingerprint is not always a guarantee and can even add false confidence. Indeed, a motivated attacker that finds 
a way to replace maliciously the file on your servers will most likely not have extra difficulty replacing also the file fingerprint, 
therefore rendering your fingerprint verification silently useless.

Here RockVerify allows to simply register your item on the blockchain anonymously, making it easily downloadable later on while removing
the possibility of the attacker scenario above happening.

## Register on the blockchain

To register an item on the blockchain export your Rockside API key and register the URL of the item:

```console
export ROCKSIDE_API_KEY
rockverify register https://...
``` 

## Install

To install the CLI locally run `go get github.com/simcap/rockverify` or grab the latest binary here.
