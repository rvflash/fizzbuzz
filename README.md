# FizzBuzz

[![GoDoc](https://godoc.org/github.com/rvflash/fizzbuzz?status.svg)](https://godoc.org/github.com/rvflash/fizzbuzz)
[![Build Status](https://img.shields.io/travis/rvflash/fizzbuzz.svg)](https://travis-ci.org/rvflash/fizzbuzz)
[![Code Coverage](https://img.shields.io/codecov/c/github/rvflash/fizzbuzz.svg)](http://codecov.io/github/rvflash/fizzbuzz?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/rvflash/fizzbuzz)](https://goreportcard.com/report/github.com/rvflash/fizzbuzz)


FizzBuzz REST API over HTTPS (example).

This project offers 2 endpoints:
* A simple [package](https://godoc.org/github.com/rvflash/fizzbuzz) to play with the original FizzBuzz or make your own one. Three algorithms are available.
* A [REST API](https://github.com/rvflash/fizzbuzz/tree/master/cmd/fizzbuzz) over HTTPS to make your own FizzBuzz and get its responses as JSON.  
 
 
## Installation

Fizzbuzz uses [dep](https://golang.github.io/dep/) to manage its dependencies.
Waiting for Go Modules (code name: vgo), it's the tool to do this stuff!

```bash
$ go get github.com/rvflash/fizzbuzz
$ cd $GOPATH/src/github.com/rvflash/fizzbuzz
$ dep ensure
```


## Quick start

If you just need a Go package to play with Fizzbuzz, see this [documentation](https://godoc.org/github.com/rvflash/fizzbuzz) on the interface.
This package proposes 3 different methods to apply the algorithm. The benchmark's performances have defined the method to use in the API.  

```go
import (
	"fmt"
	"github.com/rvflash/fizzbuzz"
)
// ...
fb := fizzbuzz.Default
fmt.Println(fb.Bulk(15))
// output: [1 2 fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizzbuzz]
```

Otherwise, to start the REST API, see the [cmd/fizzbuzz](https://github.com/rvflash/fizzbuzz/tree/master/cmd/fizzbuzz) for more information.
Also see, the [documentation](https://github.com/rvflash/fizzbuzz/tree/master/api) of the API handler.  


## Dependencies

### Gin

One of the best features of Go is its built-in net/http library.
However, I wanted to use Gin for this project to simplify some parts of the code and used a non-standard package to use dep.
And, it's a really fast alternative to build a micro-service like this one (see [Goa](https://github.com/goadesign/goa) for more complex projects).


## Features

* A package that exposes methods to play with FizzBuzz, with 3 different algorithms (see benchmark bellow).
* A REST API over HTTPS to use it.
* Continuous integration with Travis that run all lints and tests (benchmark included). 
* Fully tested and code that passed various [linters](https://github.com/golangci/golangci-lint) (go vet, etc.)
* A Systemd unit file to easily manage and control the service. The logs can be manipulated with Journalctl.
* The API can recover on its own panics. Systemd will restart it for others cases (server reboot, etc).
* The API exposes a health check as monitoring purpose.
* Vendored dependencies with dep.


### TODO: possible improvements

* Add a cache. As the result is deterministic, we can use a cache to speed up operations and store the API results.
* Add the cache-control header to also cache the data on client-side (max-age).
* Add the CORS headers (Access-Control-Allow-Origin, etc.) to make the API accessible by JavaScript in-browser client-side code.
* Add a system, via NGINX or other, to redirect the HTTP traffic to the HTTPS API.
* Add GZIP encoding to optimize the load time.


## Testing

As usual:

```bash
$ go test -cover -race -v $(go list ./... | grep -v /vendor/)
$ go test -bench=.
```

### Benchmark
```
goos: darwin
goarch: amd64
pkg: github.com/rvflash/fizzbuzz
BenchmarkMultiples_One-4            	20000000	        61.6 ns/op
BenchmarkMultiples_Two-4            	20000000	        59.3 ns/op
BenchmarkMultiples_Three-4          	30000000	        51.7 ns/op
BenchmarkMultiples_Bulk20-4         	 2000000	       719 ns/op
BenchmarkMultiples_BulkTwo20-4      	 2000000	       749 ns/op
BenchmarkMultiples_BulkThree20-4    	 2000000	       835 ns/op
BenchmarkMultiples_Bulk100-4        	  500000	      3422 ns/op
BenchmarkMultiples_BulkTwo100-4     	  500000	      3538 ns/op
BenchmarkMultiples_BulkThree100-4   	  300000	      3950 ns/op
PASS
ok  	github.com/rvflash/fizzbuzz	15.906s
```
