# FizzBuzz

[![GoDoc](https://godoc.org/github.com/rvflash/fizzbuzz?status.svg)](https://godoc.org/github.com/rvflash/fizzbuzz)
[![Build Status](https://img.shields.io/travis/rvflash/fizzbuzz.svg)](https://travis-ci.org/rvflash/fizzbuzz)
[![Code Coverage](https://img.shields.io/codecov/c/github/rvflash/fizzbuzz.svg)](http://codecov.io/github/rvflash/fizzbuzz?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/rvflash/fizzbuzz)](https://goreportcard.com/report/github.com/rvflash/fizzbuzz)


FizzBuzz REST API over HTTPS (example).

This project offers 2 endpoints:
* A simple [package](https://godoc.org/github.com/rvflash/fizzbuzz) to play with the original FizzBuzz or make your own one.
* A [REST API](https://github.com/rvflash/fizzbuzz/blob/master/cmd/fizzbuzz/README.md) over HTTPS to make your own FizzBuzz and get its responses as JSON.  
 
 
## Installation

Fizzbuzz uses [dep](https://golang.github.io/dep/) to manage its dependencies.
Waiting for Go Modules (code name: vgo), it's the tool to do this stuff!

```bash
$ go get github.com/rvflash/fizzbuzz
$ cd $GOPATH/src/github.com/rvflash/fizzbuzz
$ dep ensure
```


## Quick start

## Dependencies

### Gin

> I'll be back!


## Features

* A package that exposes methods to play with FizzBuzz.
* A REST API over HTTPS to use it.
* Continuous integration with Travis that run all lints and tests (benchmark included). 
* Fully tested and code that passed various [linters](https://github.com/golangci/golangci-lint) (go vet, etc.)
* A Systemd unit file to easily manage and control the service. The logs can be manipulated with Journalctl.
* The API can recover on its own panics. Systemd will restart it for others cases (server reboot, etc).
* The API exposes a health check as monitoring purpose.
* Vendored dependencies with dep.


### TODO: possible improvements

* Add a cache. As the result is deterministic, we can use a cache to speed up operations and store the API results.
* Add the CORS headers (Access-Control-Allow-Origin, etc.) to make the API accessible by JavaScript in-browser client-side code.
* Add a system, via NGINX or other, to redirect the HTTP traffic to the HTTPS API.
* Add cache-control header to also cache the data on client-side (max-age).
* Add GZIP encoding to optimize the load time.


### Testing

As usual:

```bash
$ go test -cover -race -v $(go list ./... | grep -v /vendor/)
```