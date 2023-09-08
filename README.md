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