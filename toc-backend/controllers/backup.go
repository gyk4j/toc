package controllers

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gyk4j/toc/toc-backend/models"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/backup"
)

func NewBackup(params backup.NewBackupParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	b := s.NewBackup()
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

func UpdateBackup(params backup.UpdateBackupParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	b := s.UpdateBackup(params.Body)

	if b != nil {
		res = backup.NewUpdateBackupOK().WithPayload(b)
	} else {
		apires := models.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Type:    models.APIResponseTypeError,
		}
		res = backup.NewUpdateBackupInternalServerError().WithPayload(&apires)
	}

	return res
}

func GetBackups(params backup.GetBackupsParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	b := s.GetBackups()
	if b != nil {
		res = backup.NewGetBackupsOK().WithPayload(b)
	} else {
		apires := models.APIResponse{
			Code:    http.StatusServiceUnavailable,
			Message: "Service unavailable",
			Type:    models.APIResponseTypeError,
		}
		res = backup.NewGetBackupsServiceUnavailable().WithPayload(&apires)
	}

	return res
}

func GetBackupByID(params backup.GetBackupByIDParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	b := s.GetBackupByID(params.BackupID)
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
