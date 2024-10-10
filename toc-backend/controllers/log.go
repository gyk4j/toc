package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	logops "github.com/gyk4j/toc/toc/server/restapi/operations/log"
	"github.com/gyk4j/toc/toc/server/services"
)

func ExportLog(params logops.ExportLogParams) middleware.Responder {
	var res middleware.Responder

	l := services.NewLog()
	if l != nil {
		res = logops.NewExportLogOK().WithPayload(l)
	} else {
		res = logops.NewExportLogInternalServerError()
	}

	return res
}
