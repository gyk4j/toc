package controllers

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gyk4j/toc/toc-backend/models"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/archive"
)

func NewArchive(params archive.NewArchiveParams) middleware.Responder {
	var res middleware.Responder

	s := Route(params.HTTPRequest)

	a := s.NewArchive()
	if a != nil {
		res = archive.NewNewArchiveOK().WithPayload(a)
	} else {
		apires := models.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Type:    models.APIResponseTypeError,
		}
		res = archive.NewNewArchiveInternalServerError().WithPayload(&apires)
	}

	return res
}

func UpdateArchive(params archive.UpdateArchiveParams) middleware.Responder {
	return middleware.NotImplemented("operation UpdateArchive has not yet been implemented")
}

func GetArchives(params archive.GetArchivesParams) middleware.Responder {
	var res middleware.Responder

	s := Route(params.HTTPRequest)

	a := s.GetArchives()
	if a != nil {
		res = archive.NewGetArchivesOK().WithPayload(a)
	} else {
		apires := models.APIResponse{
			Code:    http.StatusNotFound,
			Message: "Not found",
			Type:    models.APIResponseTypeError,
		}
		res = archive.NewGetArchivesNotFound().WithPayload(&apires)
	}

	return res
}

func GetArchiveByID(params archive.GetArchiveByIDParams) middleware.Responder {
	var res middleware.Responder

	s := Route(params.HTTPRequest)

	a := s.GetArchiveByID(params.ArchiveID)
	if a != nil {
		res = archive.NewGetArchiveByIDOK().WithPayload(a)
	} else {
		apires := models.APIResponse{
			Code:    http.StatusNotFound,
			Message: "Not found",
			Type:    models.APIResponseTypeError,
		}
		res = archive.NewGetArchiveByIDNotFound().WithPayload(&apires)
	}

	return res
}
