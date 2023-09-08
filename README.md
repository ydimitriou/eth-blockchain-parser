eth-blockchain-parser
==
Author: Yiannis Dimitriou

# Overview
The ``eth-blockchain-parser`` is an app for monitoring your inbound and outbound transactions on the Ethereum blockchain. 

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
- ``ports`` layer is the entry points of the application, receiving input from the outside world. It contains the http handler and the cron that is responsible to listening for new blocks (using a polling mechanism every 5 sec) from ethereum blockchain and import the data into our application by interacting with the ``app`` layer. The ``ports`` layer depends on the ``app`` layer.

NOTE: The only external package that have been used is only for testing reasons.

## Cron logic explanation
Given that for the purpose of this project we don't use a web socket in order to getting informed from ethereum blockchain every time a new block is created we implemented a polling mechanism.
When the application starts we create a go routine ``go ports.Worker.Run()``. The responsibility of this goroutine is:
1. on the init of the routine we make a call to the ethereum blockchain in order to get the first block for our application and persist it in memory.
2. then we start a ticker in order to do the following steps every 5 seconds.
3. ask ethereum blockchain for the latest block number.
4. check in memory database the latest parsed block number for the application.
5. given that all blocks are a sequence of integers we check if the lastParsedBlock+1 (which is the next block that we expect to parse) is less than or equal to ethereum's last block number. The reason is to guarantee that we will not miss any blocks between pollings. If the number for the expected next block to parse is greater than the number of ethereums last block number we do nothing since that means we have already parse the latest ethereum block after time interval pass we go to step 3. Otherwise we continue to step 6.
6. we request the next expected block to parse from the ethereum blockchain including the transactions
7. iterate over block's transactions and check From and To addresses if they refer to any of the subscribers of the service. Update transactions of the affected subscribers. 
9. after checking all transactions update the latest parsed block in memory
10. after time interval pass go to step 3.

## Future improvements
- enhance unti tests
- add integration tests
- use env variables for setting up the `host` for the ethereum blockchain, `recuring time` for the polling in cron

# Developers Handbook

## Make commands

Please use `make <target>` where `<target>` is one of the following:

```
  `run`             to run the app.
  `test`            to run the tests.
  `lint`            to perform linting.
```