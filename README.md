# textile

[![Made by Textile](https://img.shields.io/badge/made%20by-Textile-informational.svg?style=popout-square)](https://textile.io)
[![Chat on Slack](https://img.shields.io/badge/slack-slack.textile.io-informational.svg?style=popout-square)](https://slack.textile.io)
[![GitHub license](https://img.shields.io/github/license/textileio/filecoin-data-transfer-service.svg?style=popout-square)](./LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/textileio/filecoin-data-transfer-service?style=flat-square)](https://goreportcard.com/report/github.com/textileio/filecoin-data-transfer-service?style=flat-square)
[![GitHub action](https://github.com/textileio/filecoin-data-transfer-service/workflows/Tests/badge.svg?style=popout-square)](https://github.com/textileio/filecoin-data-transfer-service/actions)
[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=popout-square)](https://github.com/RichardLitt/standard-readme)

> Command-line application to create and manage data migration to Filecoin

The Filecoin Data Transfer Service (FTS) creates a queue of storage tasks based on local data and leverages the [Powergate](https://github.com/textileio/powergate) API to move data to deals on the Filecoin network. 

Textile connects and extends [Libp2p](https://libp2p.io/), [IPFS](https://ipfs.io/), and [Filecoin](https://filecoin.io/). Three interoperable technologies makeup Textile:

Join us on our [public Slack channel](https://slack.textile.io/) for news, discussions, and status updates. [Check out our blog](https://blog.textile.io/) for the latest posts and announcements.

*Warning* This project is still **pre-release** and is not ready for production usage.

## Table of Contents

-   [Prerequisites](#prerequisites)
-   [Installation](#installation)
-   [Localnet mode](#localnet-mode)
-   [Production setup](#production-setup)
-   [Tests](#tests)
-   [Benchmark](#benchmark)
-   [Contributing](#contributing)
-   [Changelog](#changelog)
-   [License](#license)

## Prerequisites

To build from source, you need to have Go 1.14 or newer installed.

## Installation

To build and install the FTS command-line tool, run:
```bash
make install
```

## Usage

```
▶ fts run --help
Run a storage task pipeline to migrate large collections of data to Filecoin

Usage:
  fts run [flags]

Flags:
  -d, --dry-run               run through steps without pushing data to Powergate API.
      --folder string         path to folder containing directories or files to process. (default "f")
  -h, --help                  help for run
  -i, --include-all           include hidden files & folders from the top-level folder.
      --ipfsrevproxy string   the reverse proxy multiaddr of IPFS node. (default "127.0.0.1:6002")
  -m, --mainnet               sets staging limits based on localnet or mainnet.
  -o, --output string         the location to store intermediate and final results. (default "results.json")
  -r, --resume                resume tasks stored in the local results output file.
  -e, --retry-errors          retry tasks with error status.

Global Flags:
  -t, --token string            the user auth token
```

### Examples

**dry-run**

```
fts run --folder ~/archives/tasks --dry-run
```

**production run**

```
fts run --folder ~/archives/tasks --mainnet
```

**resume**

```
fts run --resume
```

### Output

Default output is stored in `results.json`. Each task is an entry in a results array that is continuously updated while the job is running. 

```json
 {
  "name": "high-res-pdfs",
  "path": "/archives/tasks/n-q/high-res-pdfs",
  "bytes": 9073,
  "cid": "QmPWjS6YizvHTiP5oVAycvSPJBAhvSHnKsDHdr6ex54VzP",
  "jobID": "441c9c6e-f5c9-4933-913f-5a739ccb0b38",
  "stage": "Complete",
  "deals": [
   {
    "id": "441c9c6e-f5c9-4933-913f-5a739ccb0b38",
    "api_id": "f92ce374-dd51-42d9-9eb2-10a0ef24db66",
    "cid": "QmPWjS6YizvHTiP5oVAycvSPJBAhvSHnKsDHdr6ex54VzP",
    "status": 5,
    "deal_info": [
     {
      "proposal_cid": "bafyreihgf3ezv3ameqthzurd2sfbt5to36iroazztlsbosiqjyphkjbjj4",
      "state_id": 7,
      "state_name": "StorageDealActive",
      "miner": "f01000",
      "piece_cid": "baga6ea4seaqch4hlqplqpcnjfl2rbzo7hgrbvahpwugkgzyfq24sgbkhntnggay",
      "size": 16256,
      "price_per_epoch": 7629,
      "start_epoch": 6096,
      "duration": 521601,
      "deal_id": 4,
      "activation_epoch": 579
     }
    ],
    "created_at": 1608080158
   }
  ]
 }
```

## Design

![design overview](https://github.com/textileio/filecoin-data-transfer-service/blob/main/assets/fts.png?raw=true)


## Contributing

This project is a work in progress. As such, there's a few things you can do right now to help out:

-   **Ask questions**! We'll try to help. Be sure to drop a note (on the above issue) if there is anything you'd like to work on and we'll update the issue to let others know. Also [get in touch](https://slack.textile.io) on Slack.
-   **Open issues**, [file issues](https://github.com/textileio/filecoin-data-transfer-service/issues), submit pull requests!
-   **Perform code reviews**. More eyes will help a) speed the project along b) ensure quality and c) reduce possible future bugs.
-   **Take a look at the code**. Contributions here that would be most helpful are **top-level comments** about how it should look based on your understanding. Again, the more eyes the better.
-   **Add tests**. There can never be enough tests.

## Changelog

[Changelog is published to Releases.](https://github.com/textileio/filecoin-data-transfer-service/releases)

## License

[MIT](LICENSE)

