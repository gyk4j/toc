package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/synchronization"
)

func NewSynchronization(params synchronization.NewSynchronizationParams) middleware.Responder {
	var res middleware.Responder

	s := Route(params.HTTPRequest)

	sync := s.NewSynchronization()
	if s != nil {
		res = synchronization.NewNewSynchronizationOK().WithPayload(sync)
	} else {
		res = synchronization.NewNewSynchronizationServiceUnavailable()
	}

	return res
}

func UpdateSynchronization(params synchronization.UpdateSynchronizationParams) middleware.Responder {
	return middleware.NotImplemented("operation UpdateSynchronization has not yet been implemented")
}
