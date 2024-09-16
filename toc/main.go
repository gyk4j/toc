/*
 * TOC API
 *
 * An API to create backup and restoration
 *
 * API version: 1.0.0
 * Contact: 147011991+gyk4j@users.noreply.github.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"log"
	"net/http"

	sw "github.com/gyk4j/toc/toc/go"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
