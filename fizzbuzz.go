// Copyright (c) 2018 Hervé Gouchet. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

// Package fizzbuzz provides a client to play with FizzBuzz, customize your own version.
package fizzbuzz

import (
	"errors"
	"strconv"
	"strings"
)

// Errors.
var (
	// ErrFizz is returned if the first term is missing.
	ErrFizz = errors.New("missing first replacement term")
	// ErrBuzz is returned if the second term is missing.
	ErrBuzz = errors.New("missing second replacement term")
	// ErrFizzMod is returned if the first modulo is less than 1.
	ErrFizzMod = errors.New("invalid first modulo")
	// ErrFizzMod is returned if the second modulo is less than 1.
	ErrBuzzMod = errors.New("invalid second modulo")
)

type multiples struct {
	s1, s2 string
	m1, m2 int
}

// Default returns a default fizzbuzz with "fizz" for the first multiples of three, "buzz" for five.
// Also returns "fizzbuzz" for multiples of fifteen.
var Default = &multiples{s1: "fizz", s2: "buzz", m1: 3, m2: 5}

// Custom returns a new instance of fizzbuzz with personalized terms and multiples.
// fizz will be print for multiples of m1, buzz for multiples of m2.
// A concatenated version of fizz and buzz will be print for multiples of m1 and m2.
// If one of the fields is invalid, an error is returned.
func Custom(fizz, buzz string, mod1, mod2 int) (*multiples, error) {
	// mandatory fields
	if fizz = strings.TrimSpace(fizz); fizz == "" {
		return nil, ErrFizz
	}
	if buzz = strings.TrimSpace(buzz); buzz == "" {
		return nil, ErrBuzz
	}
	// avoids division by zero
	if mod1 < 1 {
		return nil, ErrFizzMod
	}
	if mod2 < 1 {
		return nil, ErrBuzzMod
	}
	return &multiples{s1: fizz, s2: buzz, m1: mod1, m2: mod2}, nil
}

// Do applies the the algorithm on the given number and returns the corresponding string value.
func (m *multiples) One(i int) (s string) {
	if i%m.m1 == 0 {
		s = m.s1
	}
	if i%m.m2 == 0 {
		// Add to the existing string is enough and avoids to do an other modulo for both value.
		s += m.s2
	}
	if s == "" {
		s = strconv.Itoa(i)
	}
	return
}

// Bulk returns a list of "fizzbuzz" values from 1 to the given until value.
// If until is negative or zero, a nil value is returned.
func (m *multiples) Bulk(until int) []string {
	if until < 1 {
		// Tooth
		return nil
	}
	var res = make([]string, until)
	for i := 1; i <= until; i++ {
		res[i-1] = m.One(i)
	}
	return res
}
