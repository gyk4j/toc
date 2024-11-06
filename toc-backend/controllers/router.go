package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/gyk4j/toc/toc-backend/services"
)

type Router struct {
	m map[string]services.IServer
}

var ctrl services.Controller = services.Controller{}
var srvr services.Server = services.Server{}
var router = Router{
	m: map[string]services.IServer{
		"toc":  &ctrl,
		"db":   &srvr,
		"file": &srvr,
		"mail": &srvr,
		"web":  &srvr,
	},
}

func (r *Router) Route(HTTPRequest *http.Request) services.IServer {
	s := r.m[HTTPRequest.Host]

	if s == nil {
		log.Printf("Unknown request host: %s", HTTPRequest.Host)
		name, err := os.Hostname()

		if err != nil {
			log.Printf("Failed to get host name")
		} else {
			s = r.m[name]

			if s != nil {
				log.Printf("Using host name: %s", name)
			} else {
				log.Printf("Unknown host name: %s", name)
			}
		}
	}

	if s == nil {
		log.Printf("Fallback to last resort default server")
		s = &ctrl
	}

	return s
}
