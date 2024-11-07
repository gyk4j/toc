package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/gyk4j/toc/toc-backend/services"
	"github.com/gyk4j/toc/toc-backend/services/controller"
	"github.com/gyk4j/toc/toc-backend/services/server"
)

type Router struct {
	m map[string]services.Server
}

var ctrl = services.Server{
	Backup:          &controller.BackupService{},
	Restoration:     &controller.RestorationService{},
	Transfer:        &controller.TransferService{},
	Synchronization: &controller.SynchronizationService{},
	Quota:           &controller.QuotaService{},
	Log:             &controller.LogService{},
	Archive:         &controller.ArchiveService{},
}

var srvr = services.Server{
	Backup:          &server.BackupService{},
	Restoration:     &server.RestorationService{},
	Transfer:        &server.TransferService{},
	Synchronization: &server.SynchronizationService{},
	Quota:           &server.QuotaService{},
	Log:             &server.LogService{},
	Archive:         &server.ArchiveService{},
}

var router = Router{
	m: map[string]services.Server{
		"toc":  ctrl,
		"db":   srvr,
		"file": srvr,
		"mail": srvr,
		"web":  srvr,
	},
}

func (r *Router) Route(HTTPRequest *http.Request) services.Server {
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
		log.Printf("Fallback to last resort default server")
		s = ctrl
	}

	return s
}
