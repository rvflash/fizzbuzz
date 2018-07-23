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

func ExampleMultiples_One() {
	fb := fizzbuzz.Default
	fmt.Println(fb.One(3))
	// output: fizz
}

// 20000000 - 60.1 ns/op
func BenchmarkMultiples_One(b *testing.B) {
	// runs the One function b.N times with 15 as value (worst case: both multiples).
	for n := 0; n < b.N; n++ {
		_ = fizzbuzz.Default.One(15)
	}
}

// 20000000 - 61.9 ns/op
func BenchmarkMultiples_Two(b *testing.B) {
	// runs the One function b.N times with 15 as value (worst case: both multiples).
	for n := 0; n < b.N; n++ {
		_ = fizzbuzz.Default.Two(15)
	}
}

// 30000000 - 51.5 ns/op
func BenchmarkMultiples_Three(b *testing.B) {
	// runs the One function b.N times with 15 as value (worst case: both multiples).
	for n := 0; n < b.N; n++ {
		_ = fizzbuzz.Default.Three(15)
	}
}

func ExampleMultiples_Bulk() {
	fb := fizzbuzz.Default
	fmt.Println(fb.Bulk(15))
	// output: [1 2 fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizzbuzz]
}

// 2000000 - 730 ns/op
func BenchmarkMultiples_Bulk20(b *testing.B) {
	// runs the Bulk function b.N times with 20 as value.
	for n := 0; n < b.N; n++ {
		_ = fizzbuzz.Default.Bulk(20)
	}
}

// 2000000 - 773 ns/op
func BenchmarkMultiples_BulkTwo20(b *testing.B) {
	// runs the Bulk function b.N times with 20 as value.
	for n := 0; n < b.N; n++ {
		_ = fizzbuzz.Default.BulkTwo(20)
	}
}

// 2000000 - 881 ns/op
func BenchmarkMultiples_BulkThree20(b *testing.B) {
	// runs the Bulk function b.N times with 20 as value.
	for n := 0; n < b.N; n++ {
		_ = fizzbuzz.Default.BulkThree(20)
	}
}

// 300000 - 3480 ns/op
func BenchmarkMultiples_Bulk100(b *testing.B) {
	// runs the Bulk function b.N times with 20 as value.
	for n := 0; n < b.N; n++ {
		_ = fizzbuzz.Default.Bulk(100)
	}
}

// 500000 - 4104 ns/op
func BenchmarkMultiples_BulkTwo100(b *testing.B) {
	// runs the Bulk function b.N times with 20 as value.
	for n := 0; n < b.N; n++ {
		_ = fizzbuzz.Default.BulkTwo(100)
	}
}

// 300000 - 4186 ns/op
func BenchmarkMultiples_BulkThree100(b *testing.B) {
	// runs the Bulk function b.N times with 20 as value.
	for n := 0; n < b.N; n++ {
		_ = fizzbuzz.Default.BulkThree(100)
	}
}

// Expected result of the first fifteen values with the default fizzbuzz behavior.
var exp = []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"}

func TestMultiples_Bulk(t *testing.T) {
	// The first fifteen values with the default fizzbuzz behavior.
	d := fizzbuzz.Default.Bulk(15)
	// Checks it!
	if !reflect.DeepEqual(d, exp) {
		t.Fatalf("unexpected result\ngot=%q\nexp=%q\n", d, exp)
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

func TestMultiples_BulkTwo(t *testing.T) {
	// The first fifteen values with the default fizzbuzz behavior.
	d := fizzbuzz.Default.BulkTwo(15)
	// Checks it!
	if !reflect.DeepEqual(d, exp) {
		t.Fatalf("unexpected result\ngot=%q\nexp=%q\n", d, exp)
	}
}

func TestMultiples_BulkTwo2(t *testing.T) {
	// Expected nil with value inferior to one.
	d := fizzbuzz.Default.BulkTwo(-1)
	// Checks it!
	if len(d) != 0 {
		t.Fatalf("expected nothing: got=%q\n", d)
	}
}

func TestMultiples_BulkThree(t *testing.T) {
	// The first fifteen values with the default fizzbuzz behavior.
	d := fizzbuzz.Default.BulkThree(15)
	// Checks it!
	if !reflect.DeepEqual(d, exp) {
		t.Fatalf("unexpected result\ngot=%q\nexp=%q\n", d, exp)
	}
}

func TestMultiples_BulkThree2(t *testing.T) {
	// Expected nil with value inferior to one.
	d := fizzbuzz.Default.BulkThree(-1)
	// Checks it!
	if len(d) != 0 {
		t.Fatalf("expected nothing: got=%q\n", d)
	}
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
