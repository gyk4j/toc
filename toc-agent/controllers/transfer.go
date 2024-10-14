package controllers

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gyk4j/toc/toc-agent/models"
	"github.com/gyk4j/toc/toc-agent/restapi/operations/transfer"
	"github.com/gyk4j/toc/toc-agent/services"
)

func NewTransfer(params transfer.NewTransferParams) middleware.Responder {
	var res middleware.Responder

	t := services.NewTransfer(params.Body)
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

func GetTransfers(params transfer.GetTransferParams) middleware.Responder {
	var res middleware.Responder

	t := services.GetTransfers()
	if t != nil {
		res = transfer.NewGetTransferOK().WithPayload(t)
	} else {
		apires := models.APIResponse{
			Code:    http.StatusServiceUnavailable,
			Message: "Service unavailable",
			Type:    models.APIResponseTypeError,
		}
		res = transfer.NewGetTransferServiceUnavailable().WithPayload(&apires)
	}

	return res
}

func GetTransferByID(params transfer.GetTransferByIDParams) middleware.Responder {
	var res middleware.Responder

	t := services.GetTransferByID(params.TransferID)
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
