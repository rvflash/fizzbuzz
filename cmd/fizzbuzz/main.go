// Copyright (c) 2018 Hervé Gouchet. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rvflash/fizzbuzz/api"
)

// Default SSL certificate.
const (
	certFile = "./testdata/server.pem"
	keyFile  = "./testdata/server.key"
)

// Fields with ldflags by the deployment tool.
var buildVersion string

func main() {
	// Get configuration
	env := flag.String("env", "prod", "environment name")
	port := flag.Int("port", 4433, "service port")
	flag.Parse()

	// Set environment
	gin.SetMode(api.Mode(*env))
	r := gin.Default()
	// Middleware
	r.Use(gin.Recovery())
	// Routes
	r.GET("/", api.Handler)
	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	// Starts the REST API
	addr := ":" + strconv.Itoa(*port)
	log.Printf("Listenning on localhost%s in %s mode (server buildID: %s)", addr, *env, buildVersion)
	if err := r.RunTLS(addr, certFile, keyFile); err != nil {
		log.Fatal(err)
	}
}
