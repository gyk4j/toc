package controllers

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gyk4j/toc/toc-backend/models"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/transfer"
)

func NewTransfer(params transfer.NewTransferParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	t := s.NewTransfer(params.Body)
	if t != nil {
		res = transfer.NewNewTransferOK().WithPayload(t)
	} else {
		apires := models.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Type:    models.APIResponseTypeError,
		}
		res = transfer.NewNewTransferInternalServerError().WithPayload(&apires)
	}

	return res
}

func UpdateTransfer(params transfer.UpdateTransferParams) middleware.Responder {
	return middleware.NotImplemented("operation UpdateTransfer has not yet been implemented")
}

func GetTransfers(params transfer.GetTransfersParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	t := s.GetTransfers()
	if t != nil {
		res = transfer.NewGetTransfersOK().WithPayload(t)
	} else {
		apires := models.APIResponse{
			Code:    http.StatusServiceUnavailable,
			Message: "Service unavailable",
			Type:    models.APIResponseTypeError,
		}
		res = transfer.NewGetTransfersServiceUnavailable().WithPayload(&apires)
	}

	return res
}

func GetTransferByID(params transfer.GetTransferByIDParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	t := s.GetTransferByID(params.TransferID)
	if t != nil {
		res = transfer.NewGetTransferByIDOK().WithPayload(t)
	} else {
		apires := models.APIResponse{
			Code:    http.StatusNotFound,
			Message: "Not found",
			Type:    models.APIResponseTypeError,
		}
		res = transfer.NewGetTransferByIDNotFound().WithPayload(&apires)
	}

	return res
}
