package controllers

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gyk4j/toc/toc-backend/models"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/archive"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/backup"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/log"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/quota"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/restoration"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/synchronization"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/transfer"
)

type Controller struct{}

//
// Backup
//

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

//
// Restoration
//

func NewRestoration(params restoration.NewRestorationParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

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

	s := router.Route(params.HTTPRequest)

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

	s := router.Route(params.HTTPRequest)

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

//
// Transfer
//

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

//
// Synchronization
//

func NewSynchronization(params synchronization.NewSynchronizationParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	sync := s.NewSynchronization()
	if s != nil {
		res = synchronization.NewNewSynchronizationOK().WithPayload(sync)
	} else {
		res = synchronization.NewNewSynchronizationServiceUnavailable()
	}

	return res
}

func UpdateSynchronization(params synchronization.UpdateSynchronizationParams) middleware.Responder {
	return middleware.NotImplemented("operation UpdateSynchronization has not yet been implemented")
}

//
// Quota
//

func GetQuotas(params quota.GetQuotasParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	qs := s.GetQuotas()
	if qs != nil {
		res = quota.NewGetQuotasOK().WithPayload(qs)
	} else {
		res = quota.NewGetQuotasServiceUnavailable()
	}

	return res
}

//
// Log
//

func ExportLog(params log.ExportLogParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	l := s.NewLog()
	if l != nil {
		res = log.NewExportLogOK().WithPayload(l)
	} else {
		res = log.NewExportLogInternalServerError()
	}

	return res
}

func UpdateLog(params log.UpdateLogParams) middleware.Responder {
	return middleware.NotImplemented("operation UpdateLog has not yet been implemented")
}

//
// Archive
//

func NewArchive(params archive.NewArchiveParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

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

	s := router.Route(params.HTTPRequest)

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

	s := router.Route(params.HTTPRequest)

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
