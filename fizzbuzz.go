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
	// ErrMode is returned if the given mode is unknown.
	ErrMode = errors.New("unknown mode")
)

// Mode represents an algorithm.
type Mode int

// List of available modes.
const (
	// One is the first algorithm (the best).
	ModeOne Mode = iota + 1
	// Two the second (almost the same performance)
	ModeTwo
	// Three is the more slowly.
	ModeThree
)

// G represents a FizzBuzz game.
type G struct {
	s1, s2 string
	m1, m2 int
	m      Mode
}

// Default returns a default fizzbuzz with "fizz" for the first multiples of three, "buzz" for five.
// Also returns "fizzbuzz" for multiples of fifteen.
var Default = &G{s1: "fizz", s2: "buzz", m1: 3, m2: 5, m: ModeOne}

// Custom returns a new instance of fizzbuzz with personalized terms and multiples.
// fizz will be print for multiples of m1, buzz for multiples of m2.
// A concatenated version of fizz and buzz will be print for multiples of m1 and m2.
// If one of the fields is invalid, an error is returned.
func Custom(fizz, buzz string, mod1, mod2 int) (*G, error) {
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
	return &G{s1: fizz, s2: buzz, m1: mod1, m2: mod2, m: ModeOne}, nil
}

// One applies the algorithm on the given number.
// It returns the corresponding string value.
// Regarding to the benchmark, it's the more faster of three methods.
func (g *G) One(i int) string {
	switch {
	case i%g.m1 == 0 && i%g.m2 == 0:
		return g.s1 + g.s2
	case i%g.m1 == 0:
		return g.s1
	case i%g.m2 == 0:
		return g.s2
	default:
		return strconv.Itoa(i)
	}
}

// Two does the same job than One, with the same performance or almost.
func (g *G) Two(i int) (s string) {
	if i%g.m1 == 0 {
		s = g.s1
	}
	if i%g.m2 == 0 {
		// Add to the existing string is enough.
		// It avoids to do an other modulo with the both value.
		s += g.s2
	}
	if s == "" {
		s = strconv.Itoa(i)
	}
	return
}

// Three does the same job than One, but more slower.
func (g *G) Three(i int) string {
	switch {
	case i%(g.m1*g.m2) == 0:
		return g.s1 + g.s2
	case i%g.m1 == 0:
		return g.s1
	case i%g.m2 == 0:
		return g.s2
	default:
		return strconv.Itoa(i)
	}
}

// Bulk returns a list of "fizzbuzz" values from 1 to the given until value.
// If until is negative or zero, a nil value is returned.
func (g *G) Bulk(until int) []string {
	if until < 1 {
		// Tooth
		return nil
	}
	var res = make([]string, until)
	for i := 1; i <= until; i++ {
		// Proof of concept: switches by algorithm.
		if g.m == ModeThree {
			res[i-1] = g.Three(i)
		} else if g.m == ModeTwo {
			res[i-1] = g.Two(i)
		} else {
			res[i-1] = g.One(i)
		}
	}
	return res
}

// Use sets the algorithm to use.
// If unknown, an error is returned.
func (g *G) Use(m Mode) error {
	switch m {
	case ModeOne, ModeTwo, ModeThree:
		// Supported modes.
		g.m = m
	default:
		return ErrMode
	}
	return nil
}
