# RockVerify

RockVerify is a command line interface that leverages the Ethereum blockchain to add an extra layer of verification 
when downloading files from URL.

To interact easily with the blockchain, we use [Rockside](https://www.rockside.io).

## Example

You want a third party to download securely content from a URL you have. 

_First_, register on the blockchain the fingerprint of the URL with:

```sh
# note the URL does not become public as only its fingerprints is registered on the blockchain
rockverify register https://example.com/releases/1.5.4/binary
```

_Secondly_, send to your third party the following command to run:

```sh
rockverify https://example.com/releases/1.5.4/binary 
```

Your third party will get the content automatically & securely on its local machine, according to the following output:

```console
rockverify https://example.com/releases/1.5.4/binary
[+] normalizing given URL
[+] reading blockchain entry for 'https://example.com/releases/1.5.4/binary'
[+] content downloaded to local file /tmp/rockverify-131833903
[+] fingerprint of downloaded content matches registered fingerprint on blockchain
[+] verification is successful. Thanks Rockside!
```

(For HTTPS authentication based accessible content see below.)

## What are we trying to solve?

When we advertise on a website a file to be downloaded and the corresponding file fingerprint to be verified after download,
we try to mitigate an in transit corruption of the file or a man in the middle attack. 

Unfortunately, the fingerprint is not always a guarantee and can even add false confidence. 

Indeed, a motivated attacker that finds a way to replace maliciously the file on your servers will have no difficulty replacing also the file fingerprint, 
therefore rendering your fingerprint verification silently useless.

Here RockVerify allows to simply register the fingerprints of the URL and content on the blockchain, making it easily downloadable later on while removing
the possibility of the attacker scenario above happening.

## Register an entry on the blockchain

In order to perform the command `rockverify register ...`, you will need to get a [Rockside API key](https://www.rockside.io) 
since we need to write on the blockchain.

Once you have one, do: 
 
```sh
export ROCKSIDE_API_KEY=....
rockverify register https://...
``` 

## Install

To install the CLI locally, grab the latest linux/windows [binaries here](https://github.com/simcap/rockverify/releases). Or if you have GO just run `go get github.com/simcap/rockverify` 

## HTTPS Basic Authentication

If your content is protected with HTTPS basic authentication, use the `basic-auth-username` flag (password being asked via prompt).

For example:

```console
rockverify --basic-auth-username john register https://example.com/releases/1.5.4/binary
[+] normalizing given URL
[+] computing fingerprint of content found at 'https://example.com/releases/1.5.4/binary'
Enter Password: 

```

Then on the download side of things:

```console
rockverify --basic-auth-username john https://example.com/releases/1.5.4/binary
[+] normalizing given URL
[+] reading blockchain entry for 'https://example.com/releases/1.5.4/binary'
Enter Password:

```