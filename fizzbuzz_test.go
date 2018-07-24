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

func ExampleG_One() {
	fb := fizzbuzz.Default
	fmt.Println(fb.One(3))
	// output: fizz
}

// 20000000 - 60.1 ns/op
func BenchmarkG_One(b *testing.B) {
	// runs the One function b.N times with 15 as value (worst case: both multiples).
	for n := 0; n < b.N; n++ {
		_ = fizzbuzz.Default.One(15)
	}
}

// 20000000 - 61.9 ns/op
func BenchmarkG_Two(b *testing.B) {
	// runs the One function b.N times with 15 as value (worst case: both multiples).
	for n := 0; n < b.N; n++ {
		_ = fizzbuzz.Default.Two(15)
	}
}

// 30000000 - 51.5 ns/op
func BenchmarkG_Three(b *testing.B) {
	// runs the One function b.N times with 15 as value (worst case: both multiples).
	for n := 0; n < b.N; n++ {
		_ = fizzbuzz.Default.Three(15)
	}
}

// 2000000 - 730 ns/op
func BenchmarkG_Bulk20(b *testing.B) {
	fb := fizzbuzz.Default
	_ = fb.Use(fizzbuzz.ModeOne)
	// runs the Bulk function b.N times with 20 as value.
	for n := 0; n < b.N; n++ {
		_ = fb.Bulk(20)
	}
}

// 2000000 - 773 ns/op
func BenchmarkG_BulkTwo20(b *testing.B) {
	fb := fizzbuzz.Default
	_ = fb.Use(fizzbuzz.ModeTwo)
	// runs the Bulk function b.N times with 20 as value.
	for n := 0; n < b.N; n++ {
		_ = fb.Bulk(20)
	}
}

// 2000000 - 881 ns/op
func BenchmarkG_BulkThree20(b *testing.B) {
	fb := fizzbuzz.Default
	_ = fb.Use(fizzbuzz.ModeThree)
	// runs the Bulk function b.N times with 20 as value.
	for n := 0; n < b.N; n++ {
		_ = fb.Bulk(20)
	}
}

// 300000 - 3480 ns/op
func BenchmarkG_Bulk100(b *testing.B) {
	fb := fizzbuzz.Default
	_ = fb.Use(fizzbuzz.ModeOne)
	// runs the Bulk function b.N times with 20 as value.
	for n := 0; n < b.N; n++ {
		_ = fb.Bulk(100)
	}
}

// 500000 - 4104 ns/op
func BenchmarkG_BulkTwo100(b *testing.B) {
	fb := fizzbuzz.Default
	_ = fb.Use(fizzbuzz.ModeTwo)
	// runs the Bulk function b.N times with 20 as value.
	for n := 0; n < b.N; n++ {
		_ = fb.Bulk(100)
	}
}

// 300000 - 4186 ns/op
func BenchmarkG_BulkThree100(b *testing.B) {
	fb := fizzbuzz.Default
	_ = fb.Use(fizzbuzz.ModeThree)
	// runs the Bulk function b.N times with 20 as value.
	for n := 0; n < b.N; n++ {
		_ = fb.Bulk(100)
	}
}

func ExampleG_Bulk() {
	fb := fizzbuzz.Default
	fmt.Println(fb.Bulk(15))
	// output: [1 2 fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizzbuzz]
}

// Expected result of the first fifteen values with the default fizzbuzz behavior.
var exp = []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"}

func newFizzBuzz(m fizzbuzz.Mode) (*fizzbuzz.G, error) {
	fb, err := fizzbuzz.Custom("fizz", "buzz", 3, 5)
	if err != nil {
		return nil, err
	}
	if err := fb.Use(m); err != nil {
		return nil, err
	}
	return fb, nil
}

func TestG_Bulk(t *testing.T) {
	// By methods
	one, err := newFizzBuzz(fizzbuzz.ModeOne)
	if err != nil {
		t.Fatal(err)
	}
	two, err := newFizzBuzz(fizzbuzz.ModeTwo)
	if err != nil {
		t.Fatal(err)
	}
	three, err := newFizzBuzz(fizzbuzz.ModeThree)
	if err != nil {
		t.Fatal(err)
	}
	// Test all the algorithms.
	var dt = []struct {
		in  *fizzbuzz.G
		out []string
	}{
		{in: one, out: exp},
		{in: two, out: exp},
		{in: three, out: exp},
	}
	for i, tt := range dt {
		if out := tt.in.Bulk(15); !reflect.DeepEqual(out, tt.out) {
			t.Errorf("%d. mismatch content:\ngot=%q\nexp=%q", i, out, tt.out)
		}
	}
}

func TestG_Bulk2(t *testing.T) {
	// Expected nil with value inferior to one.
	d := fizzbuzz.Default.Bulk(-1)
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

func TestG_Use(t *testing.T) {
	var dt = []struct {
		in  fizzbuzz.Mode
		err error
	}{
		{err: fizzbuzz.ErrMode},
		{in: fizzbuzz.ModeOne},
		{in: fizzbuzz.ModeTwo},
		{in: fizzbuzz.ModeThree},
		{in: 4, err: fizzbuzz.ErrMode},
	}
	// Checks it!
	for i, tt := range dt {
		if err := fizzbuzz.Default.Use(tt.in); err != tt.err {
			t.Errorf("%d. mismatch result: got=%q, exp=%q", i, err, tt.err)
		}
	}
}
