# magnetico

**magnetico** is the first autonomous (self-hosted) BitTorrent DHT crawler suite. It allows anyone with a decent
Internet connection to access the vast amount of torrents waiting to be discovered within the BitTorrent DHT
space, *without relying on any central entity*.

**magnetico** liberates BitTorrent from the yoke of centralised trackers & web-sites and makes it *truly decentralised*.
Finally!

## Fork Features

This is a fork of the original **magnetico** project from [boramalper](https://github.com/boramalper/magnetico) who has
since moved on to other things. This fork focuses on keeping the main component (the high performance DHT scraper)
functional, with the goal of allowing users to simply run the executable to make their own pipelines. As such, the
sibling project **magneticow** (web interface) has been removed, additionally all datastores have also been removed.

This fork implements a [ZeroMQ](https://zeromq.org/) interface for interacting with the data stream, as opposed to the
original's `stdout` stream. You can find data format and more information [here](src/API.md).

> **DISCLAIMER:**
>
> **magneticod** is considered alpha software and may remain as such indefinitely, additionally it does not implement
> a rate limiting system, so use it with caution.

## Features
- Easy installation & minimal requirements:
  - [Pre-compiled static binaries](https://github.com/fabricio20/magnetico/releases) and [Docker images](https://hub.docker.com/u/boramalper) are provided.
  - Root access is *not* required to install or to use.
- Near-zero configuration:
  - **magneticod** works out of the box, no configuration required
- No reliance on any centralised entity:
  - **magneticod** trawls the BitTorrent DHT by "going" from one node to another, and fetches the
    metadata using the nodes without using trackers.
- Resilience:
  - Unlike client-server model that web applications use, P2P networks are *chaotic* and
    **magneticod** is designed to handle all the operational errors accordingly.
- High performance implementation in Go:
  - **magneticod** utilizes every bit of your resources to discover as many infohashes & metadata as
    possible.

## Why?
BitTorrent, being a distributed P2P file sharing protocol, has long suffered because of the
centralised entities that people depended on for searching torrents (websites) and for discovering
other peers (trackers). Introduction of DHT (distributed hash table) eliminated the need for
trackers, allowing peers to discover each other through other peers and to fetch metadata from the
leechers & seeders in the network. **magnetico** is the finishing move that allows users to search
for torrents in the network, hence removing the need for centralised torrent websites.

## Installation Instructions

### Binary

You can download a compiled binary on the GitHub releases page.

### Docker

This fork provides a docker image for your convenience, it's available under GitHub Packages.

## License

All the code is licensed under AGPLv3, unless stated otherwise specifically. See `COPYING` for
details.

## Donations

You can donate to the original author below. This fork does not accept donations outside of code contributions.

### Patreon
https://www.patreon.com/boramalper

### PayPal
https://paypal.me/boramalper

### Cryptocurrencies (Coinbase)
- **BTC:** `3BLWjamWug3QQzcDDGwYLwuCqJyjcfYJB8`
- **LTC:** `MRWX5SGCF7EvN15gpzT5b3KQD3Z91gH8qi`
- **BCH:** `qqn07a58hax9l8pckq9j8ys6dsh2cnu4rsyztw2kj9`
- **ETH:** `0xe5A8e80bAA6129DF7eBB1B5302F9e2Ef4C6f6E62`
- **ETC:** `0x8964EcC86eaf043Bff2CdfE875E73D8095c26a58`
