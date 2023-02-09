<!--
parent:
  order: false
-->

<div align="center">
  <h1> Fieryeyes </h1>
</div>

<div align="center">
  <a href="https://github.com/savour-labs/fieryeyes/releases/latest">
    <img alt="Version" src="https://img.shields.io/github/tag/savour-labs/fieryeyes.svg" />
  </a>
  <a href="https://github.com/savour-labs/fieryeyes/blob/main/LICENSE">
    <img alt="License: Apache-2.0" src="https://img.shields.io/github/license/savour-labs/fieryeyes.svg" />
  </a>
  <a href="https://pkg.go.dev/github.com/savour-labs/fieryeyes">
    <img alt="GoDoc" src="https://godoc.org/github.com/savour-labs/fieryeyes?status.svg" />
  </a>
</div>

Fieryeyes is a project that integrates NFT data capture, NFT digital rule sorting and NFT recommendation

**Note**: Requires [Go 1.18+](https://golang.org/dl/)

## Introduce

- fe-law: The NFT configurable rule processing center is used for NFT rule data processing.
- fe-scrapy: NFT data crawler service
- fe-service: NFT service center, data cleaning storage service, providing internal RPC and openapi services
- indexer: Synchronize blockchain data and do simple processing, provide RPC service to the service center to pull data and calculate and process data according to rules.
- proxyd: proxyd.
- proto: The rpc interface definition of each service module
- savs: project script.
- specs:project specs.

![nft-data-1.png](https://github.com/savour-labs/savour-docs-chinese/blob/main/images/nft-data-1.png)

## Installation

For prerequisites and detailed build instructions please read the [Installation](https://github.com/savour-labs/fieryeyes/) instructions. Once the dependencies are installed, run:

```bash
make build
```

Or check out the latest [release](https://github.com/savour-labs/fieryeyes).

## Quick Start

```bash
make up
```

## Module development

### indexer

1.clone project

```
git clone git@github.com:savour-labs/fieryeyes.git
```

```
cd indexer
```

2.do module development


3.build and start

before you start up, you should change your evn.

```
make indexer
```
```
source .evn
```
```
./indexer
```

4.The same goes for other modules, such as `fe-law`, `fe-service` and `fe-scrapy`

## Contributing

Read through [CONTRIBUTING.md](https://github.com/savour-labs/fieryeyes/blob/main/CONTRIBUTING.md) for a general overview of our contribution process. Then check out our list of good first issues to find something fun to work on!

## Disclaimer
This code has not yet been audited, and should not be used in any production systems.
