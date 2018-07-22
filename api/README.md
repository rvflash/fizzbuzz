# API

[![GoDoc](https://godoc.org/github.com/rvflash/fizzbuzz/api?status.svg)](https://godoc.org/github.com/rvflash/fizzbuzz/api)
[![Build Status](https://img.shields.io/travis/rvflash/fizzbuzz/api.svg)](https://travis-ci.org/rvflash/fizzbuzz/api)
[![Code Coverage](https://img.shields.io/codecov/c/github/rvflash/fizzbuzz/api.svg)](http://codecov.io/github/rvflash/fizzbuzz/api?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/rvflash/fizzbuzz/api)](https://goreportcard.com/report/github.com/rvflash/fizzbuzz/api)

The FizzBuzz API provides a Gin handler to catch requests to customize your own FizzBuzz.
It returns the result as a JSON string.

This handler takes the following GET parameters:

| parameters | required? | description                                                      |
|------------|:---------:|------------------------------------------------------------------|
| string1    | mandatory | all multiples of int1 are replaced by this string.               |
| string2    | mandatory | all multiples of int2 are replaced by this string.               |
| int1       | mandatory | first int value to use as multiple.                              |
| int2       | mandatory | second int value to use as multiple.                             |
| limit      |  optional | list the numbers to return from 1 to the given value, default 0. |