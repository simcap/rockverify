# RockVerify

RockVerify is a command line interface that leverages the Ethereum blockchain to add an extra layer of verification 
when downloading files from URL.

To interact easily with the blockchain, we use [Rockside](https://www.rockside.io).

## Example

You want a third party to download securely content from a URL you have. 

First, register on the blockchain the fingerprint of the URL with:

```sh
rockverify register https://example.com/releases/1.5.4/binary
```

(Note the URL does not become public as only its fingerprints is registered on the blockchain)

Then send to your third party the following command to run:

```sh
rockverify https://example.com/releases/1.5.4/binary 
```

Your third party will get the content automatically & securily on its local machine. He will see the following output:

```console
rockverify https://example.com/releases/1.5.4/binary
[+] normalizing given URL
[+] reading blockchain entry for 'https://example.com/releases/1.5.4/binary'
[+] content downloaded to local file /tmp/rockverify-131833903
[+] fingerprint of downloaded content matches registered fingerprint on blockchain
[+] verification is successful. Thanks Rockside!
```

## What are we trying to solve?

When we advertise on a website a file to be downloaded and the corresponding file fingerprint to be verified after download,
we try to mitigate an in transit corruption of the file (it being malicious or not). 

Unfortunately, the fingerprint is not always a guarantee and can even add false confidence. Indeed, a motivated attacker that finds 
a way to replace maliciously the file on your servers will most likely not have extra difficulty replacing also the file fingerprint, 
therefore rendering your fingerprint verification silently useless.

Here RockVerify allows to simply register your item on the blockchain anonymously, making it easily downloadable later on while removing
the possibility of the attacker scenario above happening.

## Register an entry on the blockchain

In order to perfom the command `rockverify register ...`, you will need to get a [Rockside](https://www.rockside.io) API key 
since we need to write on the blockchain.

Then do: 
 
```sh
export ROCKSIDE_API_KEY
rockverify register https://...
``` 

## Install

To install the CLI locally run `go get github.com/simcap/rockverify` or grab the latest binary here.
