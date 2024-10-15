package controllers

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gyk4j/toc/toc-backend/models"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/archive"
	"github.com/gyk4j/toc/toc-backend/services"
)

func NewArchive(params archive.ArchiveDataParams) middleware.Responder {
	var res middleware.Responder

	a := services.ArchiveData()
	if a != nil {
		res = archive.NewArchiveDataOK().WithPayload(a)
	} else {
		apires := models.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Type:    models.APIResponseTypeError,
		}
		res = archive.NewArchiveDataInternalServerError().WithPayload(&apires)
	}

	return res
}

func GetArchive(params archive.GetArchiveParams) middleware.Responder {
	var res middleware.Responder

	a := services.GetArchives()
	if a != nil {
		res = archive.NewGetArchiveOK().WithPayload(a)
	} else {
		apires := models.APIResponse{
			Code:    http.StatusNotFound,
			Message: "Not found",
			Type:    models.APIResponseTypeError,
		}
		res = archive.NewGetArchiveNotFound().WithPayload(&apires)
	}

	return res
}

func GetArchiveByID(params archive.GetArchiveByIDParams) middleware.Responder {
	var res middleware.Responder

	a := services.GetArchiveByID(params.ArchiveID)
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
