package controllers

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gyk4j/toc/toc/server/models"
	"github.com/gyk4j/toc/toc/server/restapi/operations/restoration"
	"github.com/gyk4j/toc/toc/server/services"
)

func NewRestoration(params restoration.NewRestorationParams) middleware.Responder {
	var res middleware.Responder

	r := services.NewRestoration()
	if r != nil {
		res = restoration.NewNewRestorationOK().WithPayload(r)
	} else {
		apires := models.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Type:    models.APIResponseTypeError,
		}
		res = restoration.NewNewRestorationInternalServerError().WithPayload(&apires)
	}

	return res
}

func GetRestorations(params restoration.GetRestorationParams) middleware.Responder {
	var res middleware.Responder

	r := services.GetRestorations()
	if r != nil {
		res = restoration.NewGetRestorationOK().WithPayload(r)
	} else {
		apires := models.APIResponse{
			Code:    http.StatusServiceUnavailable,
			Message: "Service unavailable",
			Type:    models.APIResponseTypeError,
		}
		res = restoration.NewGetRestorationServiceUnavailable().WithPayload(&apires)
	}

	return res
}

func GetRestorationByID(params restoration.GetRestorationByIDParams) middleware.Responder {
	var res middleware.Responder

	r := services.GetRestorationByID(params.RestorationID)
	if r != nil {
		res = restoration.NewGetRestorationByIDOK().WithPayload(r)
	} else {
		apires := models.APIResponse{
			Code:    http.StatusNotFound,
			Message: "Not found",
			Type:    models.APIResponseTypeError,
		}
		res = restoration.NewGetRestorationByIDNotFound().WithPayload(&apires)
	}

	return res
}
