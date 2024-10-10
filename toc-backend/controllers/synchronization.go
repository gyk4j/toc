package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gyk4j/toc/toc/server/restapi/operations/synchronization"
	"github.com/gyk4j/toc/toc/server/services"
)

func NewSynchronization(params synchronization.NewSynchronizationParams) middleware.Responder {
	var res middleware.Responder

	s := services.NewSynchronization()
	if s != nil {
		res = synchronization.NewNewSynchronizationOK().WithPayload(s)
	} else {
		res = synchronization.NewNewSynchronizationServiceUnavailable()
	}

	return res
}
