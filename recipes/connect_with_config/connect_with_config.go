// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

// Connect simply connects to Elasticsearch.
//
// Example
//
//
//     connect "http://127.0.0.1:9200/test-index?sniff=false&healthcheck=false"
//
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/olivere/elastic"
	"github.com/olivere/elastic/config"
)

func main() {
	flag.Parse()
	log.SetFlags(0)

	url := flag.Arg(0)
	if url == "" {
		url = "http://127.0.0.1:9200"
	}

	// Parse the URL with the config package
	cfg, err := config.Parse(url)
	if err != nil {
		log.Fatal(err)
	}

	// Create an Elasticsearch client
	client, err := elastic.NewClientFromConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}
	_ = client

	// Just a status message
	fmt.Println("Connection succeeded")
}
