package controllers

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gyk4j/toc/toc-backend/models"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/restoration"
)

func NewRestoration(params restoration.NewRestorationParams) middleware.Responder {
	var res middleware.Responder

	s := dispatch(params.HTTPRequest)

	r := s.NewRestoration(params.Body)
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

func UpdateRestoration(params restoration.UpdateRestorationParams) middleware.Responder {
	return middleware.NotImplemented("operation UpdateRestoration has not yet been implemented")
}

func GetRestorations(params restoration.GetRestorationsParams) middleware.Responder {
	var res middleware.Responder

	s := dispatch(params.HTTPRequest)

	r := s.GetRestorations()
	if r != nil {
		res = restoration.NewGetRestorationsOK().WithPayload(r)
	} else {
		apires := models.APIResponse{
			Code:    http.StatusServiceUnavailable,
			Message: "Service unavailable",
			Type:    models.APIResponseTypeError,
		}
		res = restoration.NewGetRestorationsServiceUnavailable().WithPayload(&apires)
	}

	return res
}

func GetRestorationByID(params restoration.GetRestorationByIDParams) middleware.Responder {
	var res middleware.Responder

	s := dispatch(params.HTTPRequest)

	r := s.GetRestorationByID(params.RestorationID)
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
