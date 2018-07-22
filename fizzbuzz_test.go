// Copyright (c) 2018 Hervé Gouchet. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package fizzbuzz_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/rvflash/fizzbuzz"
)

// 500000 ~3580 ns/op
func BenchmarkMultiples_Bulk100(b *testing.B) {
	fb := fizzbuzz.Default
	// runs the Bulk function b.N times
	for n := 0; n < b.N; n++ {
		_ = fb.Bulk(100)
	}
}

func ExampleMultiples_One() {
	fb := fizzbuzz.Default
	fmt.Println(fb.One(3))
	// output: fizz
}

func ExampleMultiples_Bulk() {
	fb := fizzbuzz.Default
	fmt.Println(fb.Bulk(15))
	// output: [1 2 fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizzbuzz]
}

func TestCustom(t *testing.T) {
	// Creates a new fizzbuzz with "€" for multiples of three and "$" for multiples of 8.
	fb, err := fizzbuzz.Custom("€", "$", 3, 8)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	// Gets the first twenty four numbers.
	d := fb.Bulk(24)
	// Expected result
	w := []string{
		"1", "2", "€", "4", "5", "€", "7", "$", "€", "10", "11", "€", "13", "14", "€", "$", "17", "€", "19", "20",
		"€", "22", "23", "€$",
	}
	// Checks it!
	if !reflect.DeepEqual(d, w) {
		t.Fatalf("unexpected result\ngot=%q\nexp=%q\n", d, w)
	}
}

func TestCustom2(t *testing.T) {
	var dt = []struct {
		s1, s2 string
		m1, m2 int
		err    error
	}{
		{err: fizzbuzz.ErrFizz},
		{s1: " ", err: fizzbuzz.ErrFizz},
		{s1: "fizz", err: fizzbuzz.ErrBuzz},
		{s1: "fizz", s2: " ", err: fizzbuzz.ErrBuzz},
		{s1: "fizz", s2: "buzz", m1: -3, err: fizzbuzz.ErrFizzMod},
		{s1: "fizz", s2: "buzz", m1: 3, err: fizzbuzz.ErrBuzzMod},
		{s1: "fizz", s2: "buzz", m1: 3, m2: -5, err: fizzbuzz.ErrBuzzMod},
		{s1: "fizz", s2: "buzz", m1: 3, m2: 5},
	}
	// Only checks error returned.
	for i, tt := range dt {
		if _, err := fizzbuzz.Custom(tt.s1, tt.s2, tt.m1, tt.m2); err != tt.err {
			t.Errorf("%d. mismatch content: got=%q, exp=%q", i, err, tt.err)
		}
	}
}

func TestMultiples_Bulk(t *testing.T) {
	// The first fifteen values with the default fizzbuzz behavior.
	d := fizzbuzz.Default.Bulk(15)
	// Expected result
	w := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"}
	// Checks it!
	if !reflect.DeepEqual(d, w) {
		t.Fatalf("unexpected result\ngot=%q\nexp=%q\n", d, w)
	}
}

func TestMultiples_Bulk2(t *testing.T) {
	// Expected nil with value inferior to one.
	d := fizzbuzz.Default.Bulk(-1)
	// Checks it!
	if len(d) != 0 {
		t.Fatalf("expected nothing: got=%q\n", d)
	}
}
