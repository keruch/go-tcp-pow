# go-tcp-pow

This repo contains the implementation of the challenge below.

## Challenge

Design and implement “Word of Wisdom” tcp server.
* TCP server should be protected from DDOS attacks with the Proof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.
* The choice of the POW algorithm should be explained.
* After Proof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.
* Docker file should be provided both for the server and for the client that solves the POW challenge

## How to run

### Server
   ```bash
   make run-server
   ```
   The command supports an optional value - `CHALLENGE_COMPLEXITY`, which specifies the number of zero bytes at the start of a hash in the Hashcash algorithm. The default value is 4.


### Client 
   ```bash
   make run-client
   ```
  The command supports an optional value - `SERVER_ADDR`, which specifies a server to connect to. The default value is `localhost:9000`. 

## PoW algorithm

The app uses the [Hashcash](https://en.wikipedia.org/wiki/Hashcash) algorithm. This implementation doesn't require the database since it requests to solve a challenge on every TCP connection.

The reasons for choosing this algorithm are:
* It is widely used (Bitcoin) and well-documented
* Simple verification on the server side
* Understandable and reliable
* Flexible complexity


