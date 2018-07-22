// Copyright (c) 2018 Hervé Gouchet. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rvflash/fizzbuzz"
	"github.com/rvflash/fizzbuzz/api"
)

func TestHandler(t *testing.T) {
	// Returns the expected JSON string of the given error.
	var toErr = func(err error) string {
		return fmt.Sprintf(`{"error":%q}`, err)
	}
	// Test cases
	var dt = []struct {
		// Request
		query url.Values
		// Response
		code int
		body string
	}{
		{
			code: http.StatusBadRequest,
			body: toErr(fizzbuzz.ErrFizz),
		},
		{
			query: url.Values{"string1": {"fizz"}},
			code:  http.StatusBadRequest,
			body:  toErr(fizzbuzz.ErrBuzz),
		},
		{
			query: url.Values{
				"string1": {"fizz"},
				"string2": {"buzz"},
			},
			code: http.StatusBadRequest,
			body: toErr(fizzbuzz.ErrFizzMod),
		},
		{
			query: url.Values{
				"string1": {"fizz"},
				"string2": {"buzz"},
				"int1":    {"3"},
			},
			code: http.StatusBadRequest,
			body: toErr(fizzbuzz.ErrBuzzMod),
		},
		{
			query: url.Values{
				"string1": {"fizz"},
				"string2": {"buzz"},
				"int1":    {"3"},
				"int2":    {"5"},
			},
			code: http.StatusOK,
			body: "null",
		},
		{
			query: url.Values{
				"string1": {"fizz"},
				"string2": {"buzz"},
				"int1":    {"2"},
				"int2":    {"3"},
				"limit":   {"6"},
			},
			code: http.StatusOK,
			body: `["1","fizz","buzz","fizz","5","fizzbuzz"]`,
		},
	}

	// Creates the test environment.
	gin.SetMode(gin.TestMode)
	h := gin.New()
	h.GET("/", api.Handler)

	for i, tt := range dt {
		// Builds the request.
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		// Adds given parameters.
		if len(tt.query) != 0 {
			req.URL.RawQuery = tt.query.Encode()
		}
		// Listens the response.
		w := httptest.NewRecorder()
		h.ServeHTTP(w, req)
		if w.Code != tt.code {
			t.Errorf("%d. mismatch status code: got=%d exp=%d\n", i, w.Code, tt.code)
		}
		if w.Body.String() != tt.body {
			t.Errorf("%d. mismatch content:\ngot=%q\nexp=%q\n", i, w.Body, tt.body)
		}
	}
}

func TestMode(t *testing.T) {
	var dt = []struct {
		in, out string
	}{
		{in: "", out: api.ProdEnv},
		{in: "prod", out: api.ProdEnv},
		{in: "release", out: api.ProdEnv},
		{in: "tooth", out: api.ProdEnv},
		{in: "qa", out: api.TestMode},
		{in: "test", out: api.TestMode},
		{in: "debug", out: api.DevEnv},
		{in: "dev", out: api.DevEnv},
	}
	for i, tt := range dt {
		if out := api.Mode(tt.in); out != tt.out {
			t.Errorf("%d. mismatch content: got=%q, exp=%q\n", i, out, tt.out)
		}
	}
}
