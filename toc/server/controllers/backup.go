package controllers

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gyk4j/toc/toc/server/models"
	"github.com/gyk4j/toc/toc/server/restapi/operations/backup"
	"github.com/gyk4j/toc/toc/server/services"
)

func NewBackup(params backup.NewBackupParams) middleware.Responder {
	var res middleware.Responder

	b := services.NewBackup()
	if b != nil {
		// Swagger definition dictates an API response as payload.
		// To re-generate if required.
		apires := models.APIResponse{
			Code:    http.StatusOK,
			Message: "OK",
			Type:    models.APIResponseTypeInfo,
		}
		res = backup.NewNewBackupOK().WithPayload(&apires)
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
		res = backup.NewGetBackupServiceUnavailable()
	}

	return res
}

func GetBackupByID(params backup.GetBackupByIDParams) middleware.Responder {
	var res middleware.Responder

	b := services.GetBackupByID(params.BackupID)
	if b != nil {
		res = backup.NewGetBackupByIDOK().WithPayload(b)
	} else {
		res = backup.NewGetBackupByIDNotFound()
	}

	return res
}
