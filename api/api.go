// Copyright (c) 2018 Hervé Gouchet. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

// Package api exposes an handler to customize your own FizzBuzz.
package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rvflash/fizzbuzz"
)

// List of available URL parameters.
const (
	fizzTerm = "string1"
	buzzTerm = "string2"
	fizzMod  = "int1"
	buzzMod  = "int2"
	maxBound = "limit"
)

// Handler handles requests to customize your own FizzBuzz.
// It takes as parameters the following parameters:
// > string1: all multiples of int1 are replaced by this string.
// > string2: all multiples of int2 are replaced by this string.
// > int1: first int value to use as multiple.
// > int2: second int value to use as multiple.
// > limit: list the numbers from 1 to the given value.
// It returns the result, a list of strings, as a JSON.
func Handler(c *gin.Context) {
	var toInt = func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}
	// Creates a custom FizzBuzz
	fb, err := fizzbuzz.Custom(c.Query(fizzTerm), c.Query(buzzTerm), toInt(c.Query(fizzMod)), toInt(c.Query(buzzMod)))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, fb.Bulk(toInt(c.Query(maxBound))))
}

// List of environments.
const (
	// DevEnv for development environment.
	DevEnv = "debug"
	// ProdEnv for production environment.
	ProdEnv = "release"
	// TestMode for test environment.
	TestMode = "test"
)

// Mode ensures to manipulate a Gin's mode.
// By default, production is returned.
func Mode(env string) string {
	switch env {
	case DevEnv, "dev":
		return DevEnv
	case TestMode, "qa":
		return TestMode
	default:
		return ProdEnv
	}
}
