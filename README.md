btcrpcclient
============

[![Build Status](https://travis-ci.org/mably/btcrpcclient.png?branch=master)]
(https://travis-ci.org/mably/btcrpcclient)
[![GoDoc](https://godoc.org/github.com/mably/btcrpcclient?status.png)]
(http://godoc.org/github.com/mably/btcrpcclient)

btcrpcclient implements a Websocket-enabled Bitcoin JSON-RPC client package
written in [Go](http://golang.org/).  It provides a robust and easy to use
client for interfacing with a Bitcoin RPC server that uses a ppcd/bitcoin core
compatible Bitcoin JSON-RPC API.

## Status

This package is currently under active development.  It is already stable and
the infrastructure is complete.  However, there are still several RPCs left to
implement and the API is not stable yet.

## Documentation

* [API Reference](http://godoc.org/github.com/mably/btcrpcclient)
* [ppcd Websockets Example](https://github.com/mably/btcrpcclient/blob/master/examples/btcdwebsockets)  
  Connects to a ppcd RPC server using TLS-secured websockets, registers for
  block connected and block disconnected notifications, and gets the current
  block count
* [ppcwallet Websockets Example](https://github.com/mably/btcrpcclient/blob/master/examples/ppcwalletwebsockets)  
  Connects to a ppcwallet RPC server using TLS-secured websockets, registers for
  notifications about changes to account balances, and gets a list of unspent
  transaction outputs (utxos) the wallet can sign
* [Bitcoin Core HTTP POST Example](https://github.com/mably/btcrpcclient/blob/master/examples/bitcoincorehttp)  
  Connects to a bitcoin core RPC server using HTTP POST mode with TLS disabled
  and gets the current block count

## Major Features

* Supports Websockets (ppcd/ppcwallet) and HTTP POST mode (bitcoin core)
* Provides callback and registration functions for ppcd/ppcwallet notifications
* Supports ppcd extensions
* Translates to and from higher-level and easier to use Go types
* Offers a synchronous (blocking) and asynchronous API
* When running in Websockets mode (the default):
  * Automatic reconnect handling (can be disabled)
  * Outstanding commands are automatically reissued
  * Registered notifications are automatically reregistered
  * Back-off support on reconnect attempts

## Installation

```bash
$ go get github.com/mably/btcrpcclient
```

## License

Package btcrpcclient is licensed under the [copyfree](http://copyfree.org) ISC
License.
