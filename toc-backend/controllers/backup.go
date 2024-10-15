package controllers

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gyk4j/toc/toc-backend/models"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/backup"
	"github.com/gyk4j/toc/toc-backend/services"
)

func NewBackup(params backup.NewBackupParams) middleware.Responder {
	var res middleware.Responder

	b := services.NewBackup()
	if b != nil {
		res = backup.NewNewBackupOK().WithPayload(b)
	} else {
		apires := models.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Type:    models.APIResponseTypeError,
		}
		res = backup.NewNewBackupInternalServerError().WithPayload(&apires)
	}

	return res
}

func GetBackups(params backup.GetBackupParams) middleware.Responder {
	var res middleware.Responder

	b := services.GetBackups()
	if b != nil {
		res = backup.NewGetBackupOK().WithPayload(b)
	} else {
		apires := models.APIResponse{
			Code:    http.StatusServiceUnavailable,
			Message: "Service unavailable",
			Type:    models.APIResponseTypeError,
		}
		res = backup.NewGetBackupServiceUnavailable().WithPayload(&apires)
	}

	return res
}

func GetBackupByID(params backup.GetBackupByIDParams) middleware.Responder {
	var res middleware.Responder

	b := services.GetBackupByID(params.BackupID)
	if b != nil {
		res = backup.NewGetBackupByIDOK().WithPayload(b)
	} else {
		apires := models.APIResponse{
			Code:    http.StatusNotFound,
			Message: "Not found",
			Type:    models.APIResponseTypeError,
		}
		res = backup.NewGetBackupByIDNotFound().WithPayload(&apires)
	}

	return res
}
