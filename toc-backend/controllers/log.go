package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	logops "github.com/gyk4j/toc/toc-backend/restapi/operations/log"
)

func ExportLog(params logops.ExportLogParams) middleware.Responder {
	var res middleware.Responder

	s := dispatch(params.HTTPRequest)

	l := s.NewLog()
	if l != nil {
		res = logops.NewExportLogOK().WithPayload(l)
	} else {
		res = logops.NewExportLogInternalServerError()
	}

	return res
}

func UpdateLog(params logops.UpdateLogParams) middleware.Responder {
	return middleware.NotImplemented("operation UpdateLog has not yet been implemented")
}
