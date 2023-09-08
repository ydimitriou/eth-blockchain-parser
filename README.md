eth-blockchain-parser
==
Author: Yiannis Dimitriou

# Overview
The ``eth-blockchain-parser`` is an app for monitoring your inbound and outbound transaction on the Ethereum blockchain. 

# Features
- Subscribe with your address
- Receive all inbound and outbound transactions for a subscribed address
- Get the block number of the latest block parsed from the ethereum blockchain

# Examples

## Subscribe
To subscribe in the eth-blockchain-parser service:
- create http request body including the address you want to subscribe ``{"address": "0x1f9090aae28b8a3dceadf281b0f12828e676c326"}``
- make ``POST`` request on ``/v1/subscriber`` 

## Get latest parsed block
To get the latest parsed block from the eth-blockchain-parser:
- make ``GET`` request on ``/v1/last-block``

## Get all transactions for an address
To get all inbound and outbound transactions for an address:
- first subscribe to the eth-blockchain-parser service (description on Subscribe section)
- make ``GET`` request on ``/v1/subscriber/{address}``. e.g: ``/v1/subscriber/0x1f9090aae28b8a3dceadf281b0f12828e676c326``

# Technical Info

## Tech Stack

| Type          | Item      | Version |
|---------------|-----------|---------|
| Language      | Go        | 1.19    |

## Design overview
The design of this project follows the fundamentals of the Clean Architecture design concepts combined with the Domain Driver Design concepts.
The main layers of the project are:
- ``domain`` layer contains the components of the application (Domain Entities and Domain Services). ``domain`` layer does not have any dependencies on other layers.
- ``app`` (application) layer exposes all available use cases of the application to the outside world. On the application layer the project follows also the command query segregation pattern. The ``app`` layer depends only on ``domain`` layer.
- ``adapters`` layer is responsible for implementing domain and application interfaces by integrating with specific providers etc. For example we can implement the domain repository interface either with a Memory or a SQL provider (for our purpose we used in memory database). The ``adapters`` layer depends on the ``app`` and ``domain`` layers.
- ``ports`` layer is the entry points of the application, receiving input from the outside world. It contains the http handler and the cron that is responsible to listening for new blocks from ethereum blockchain and import the data into our application byt interacting with the ``app`` layer. The ``ports`` layer depends on the ``app`` layer.

# Developers Handbook

## Make commands

Please use `make <target>` where `<target>` is one of the following:

```
  `run`             to run the app.
  `test`            to run the tests.
```