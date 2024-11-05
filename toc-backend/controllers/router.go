package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/gyk4j/toc/toc-backend/services"
	"github.com/gyk4j/toc/toc-backend/services/controller"
	"github.com/gyk4j/toc/toc-backend/services/server"
)

var c controller.Controller = controller.Controller{}
var s server.Server = server.Server{}

var m = map[string]services.Server{
	"toc":  &c,
	"db":   &s,
	"file": &s,
	"mail": &s,
	"web":  &s,
}

func Route(HTTPRequest *http.Request) services.Server {
	s := m[HTTPRequest.Host]

	if s == nil {
		log.Printf("Unknown request host: %s", HTTPRequest.Host)
		name, err := os.Hostname()

		if err != nil {
			log.Printf("Failed to get host name")
		} else {
			s = m[name]

			if s != nil {
				log.Printf("Using host name: %s", name)
			} else {
				log.Printf("Unknown host name: %s", name)
			}
		}
	}

	if s == nil {
		log.Printf("Fallback to last resort default server")
		s = &c
	}

	return s
}
