package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gyk4j/toc/toc-backend/services"
	"github.com/gyk4j/toc/toc-backend/services/controller"
	"github.com/gyk4j/toc/toc-backend/services/server"
)

type Router struct {
	m map[string]*services.Server
}

func NewRouter() *Router {
	return &Router{
		m: map[string]*services.Server{
			"toc":  controller.NewController("toc"),
			"db":   server.NewServer("db"),
			"file": server.NewServer("file"),
			"mail": server.NewServer("mail"),
			"web":  server.NewServer("web"),
		},
	}
}

func (r *Router) Route(HTTPRequest *http.Request) *services.Server {
	s := r.m[HTTPRequest.Host]

	if s.Backup == nil {
		log.Printf("Unknown request host: %s", HTTPRequest.Host)
		name, err := os.Hostname()

		if err != nil {
			log.Printf("Failed to get host name")
		} else {
			s = r.m[name]

			if s.Backup != nil {
				log.Printf("Using host name: %s", name)
			} else {
				log.Printf("Unknown host name: %s", name)
			}
		}
	}

	if s.Backup == nil {
		panic(fmt.Sprintf("Unknown host name: %s", HTTPRequest.Host))
	}

	return s
}

var router *Router = NewRouter()
