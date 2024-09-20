package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gyk4j/toc/toc/server/restapi/operations/archive"
	"github.com/gyk4j/toc/toc/server/services"
)

func NewArchive(params archive.ArchiveDataParams) middleware.Responder {
	services.ArchiveData()
	res := archive.NewArchiveDataOK()
	return res
}

func GetArchive(params archive.GetArchiveParams) middleware.Responder {
	var res middleware.Responder

	a := services.GetArchives()
	if a != nil {
		res = archive.NewGetArchiveOK().WithPayload(a)
	} else {
		res = archive.NewGetArchiveNotFound()
	}

	return res
}

func GetArchiveByID(params archive.GetArchiveByIDParams) middleware.Responder {
	var res middleware.Responder

	a := services.GetArchiveByID(params.ArchiveID)
	if a != nil {
		res = archive.NewGetArchiveByIDOK()
	} else {
		res = archive.NewGetArchiveByIDNotFound()
	}

	return res
}
