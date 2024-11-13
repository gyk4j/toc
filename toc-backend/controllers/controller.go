package controllers

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/archive"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/backup"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/log"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/quota"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/restoration"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/synchronization"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/transfer"
	"github.com/gyk4j/toc/toc-backend/utils"
)

//
// Backup
//

func NewBackup(params backup.NewBackupParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	b, err := s.Backup.NewBackup()
	if err == nil {
		res = backup.NewNewBackupOK().WithPayload(b)
	} else {
		apires := utils.NewAPIResponseFactory().GetAPIResponse(http.StatusInternalServerError)
		res = backup.NewNewBackupInternalServerError().WithPayload(apires)
	}

	return res
}

func UpdateBackup(params backup.UpdateBackupParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	b, err := s.Backup.UpdateBackup(params.Body)

	if err == nil {
		res = backup.NewUpdateBackupOK().WithPayload(b)
	} else {
		apires := utils.NewAPIResponseFactory().GetAPIResponse(http.StatusInternalServerError)
		res = backup.NewUpdateBackupInternalServerError().WithPayload(apires)
	}

	return res
}

func GetBackups(params backup.GetBackupsParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	b, err := s.Backup.GetBackups()
	if err == nil {
		res = backup.NewGetBackupsOK().WithPayload(b)
	} else {
		apires := utils.NewAPIResponseFactory().GetAPIResponse(http.StatusServiceUnavailable)
		res = backup.NewGetBackupsServiceUnavailable().WithPayload(apires)
	}

	return res
}

func GetBackupByID(params backup.GetBackupByIDParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	b, err := s.Backup.GetBackupById(params.BackupID)
	if err == nil {
		res = backup.NewGetBackupByIDOK().WithPayload(b)
	} else {
		apires := utils.NewAPIResponseFactory().GetAPIResponse(http.StatusNotFound)
		res = backup.NewGetBackupByIDNotFound().WithPayload(apires)
	}

	return res
}

//
// Restoration
//

func NewRestoration(params restoration.NewRestorationParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	r, err := s.Restoration.NewRestoration(params.Body)
	if err == nil {
		res = restoration.NewNewRestorationOK().WithPayload(r)
	} else {
		apires := utils.NewAPIResponseFactory().GetAPIResponse(http.StatusInternalServerError)
		res = restoration.NewNewRestorationInternalServerError().WithPayload(apires)
	}

	return res
}

func UpdateRestoration(params restoration.UpdateRestorationParams) middleware.Responder {
	var res middleware.Responder

	apires := utils.NewAPIResponseFactory().GetAPIResponse(http.StatusNotImplemented)
	res = restoration.NewUpdateRestorationMethodNotAllowed().WithPayload(apires)

	return res
}

func GetRestorations(params restoration.GetRestorationsParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	r, err := s.Restoration.GetRestorations()
	if err == nil {
		res = restoration.NewGetRestorationsOK().WithPayload(r)
	} else {
		apires := utils.NewAPIResponseFactory().GetAPIResponse(http.StatusServiceUnavailable)
		res = restoration.NewGetRestorationsServiceUnavailable().WithPayload(apires)
	}

	return res
}

func GetRestorationByID(params restoration.GetRestorationByIDParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	r, err := s.Restoration.GetRestorationById(params.RestorationID)
	if err == nil {
		res = restoration.NewGetRestorationByIDOK().WithPayload(r)
	} else {
		apires := utils.NewAPIResponseFactory().GetAPIResponse(http.StatusNotFound)
		res = restoration.NewGetRestorationByIDNotFound().WithPayload(apires)
	}

	return res
}

//
// Transfer
//

func NewTransfer(params transfer.NewTransferParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	t, err := s.Transfer.NewTransfer(params.Body)
	if err == nil {
		res = transfer.NewNewTransferOK().WithPayload(t)
	} else {
		apires := utils.NewAPIResponseFactory().GetAPIResponse(http.StatusInternalServerError)
		res = transfer.NewNewTransferInternalServerError().WithPayload(apires)
	}

	return res
}

func UpdateTransfer(params transfer.UpdateTransferParams) middleware.Responder {
	var res middleware.Responder

	apires := utils.NewAPIResponseFactory().GetAPIResponse(http.StatusNotImplemented)
	res = transfer.NewUpdateTransferMethodNotAllowed().WithPayload(apires)

	return res
}

func GetTransfers(params transfer.GetTransfersParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	t, err := s.Transfer.GetTransfers()
	if err == nil {
		res = transfer.NewGetTransfersOK().WithPayload(t)
	} else {
		apires := utils.NewAPIResponseFactory().GetAPIResponse(http.StatusServiceUnavailable)
		res = transfer.NewGetTransfersServiceUnavailable().WithPayload(apires)
	}

	return res
}

func GetTransferByID(params transfer.GetTransferByIDParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	t, err := s.Transfer.GetTransferById(params.TransferID)
	if err == nil {
		res = transfer.NewGetTransferByIDOK().WithPayload(t)
	} else {
		apires := utils.NewAPIResponseFactory().GetAPIResponse(http.StatusNotFound)
		res = transfer.NewGetTransferByIDNotFound().WithPayload(apires)
	}

	return res
}

//
// Synchronization
//

func NewSynchronization(params synchronization.NewSynchronizationParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	sync, err := s.Synchronization.NewSynchronization()
	if err == nil {
		res = synchronization.NewNewSynchronizationOK().WithPayload(sync)
	} else {
		apires := utils.NewAPIResponseFactory().GetAPIResponse(http.StatusServiceUnavailable)
		res = synchronization.NewNewSynchronizationServiceUnavailable().WithPayload(apires)
	}

	return res
}

func UpdateSynchronization(params synchronization.UpdateSynchronizationParams) middleware.Responder {
	var res middleware.Responder

	apires := utils.NewAPIResponseFactory().GetAPIResponse(http.StatusNotImplemented)
	res = synchronization.NewUpdateSynchronizationMethodNotAllowed().WithPayload(apires)

	return res
}

//
// Quota
//

func GetQuotas(params quota.GetQuotasParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	qs, err := s.Quota.GetQuotas()
	if err == nil {
		res = quota.NewGetQuotasOK().WithPayload(qs)
	} else {
		apires := utils.NewAPIResponseFactory().GetAPIResponse(http.StatusServiceUnavailable)
		res = quota.NewGetQuotasServiceUnavailable().WithPayload(apires)
	}

	return res
}

//
// Log
//

func ExportLog(params log.ExportLogParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	l, err := s.Log.NewLog()
	if err == nil {
		res = log.NewExportLogOK().WithPayload(l)
	} else {
		apires := utils.NewAPIResponseFactory().GetAPIResponse(http.StatusInternalServerError)
		res = log.NewExportLogInternalServerError().WithPayload(apires)
	}

	return res
}

func UpdateLog(params log.UpdateLogParams) middleware.Responder {
	var res middleware.Responder

	apires := utils.NewAPIResponseFactory().GetAPIResponse(http.StatusNotImplemented)
	res = log.NewUpdateLogMethodNotAllowed().WithPayload(apires)

	return res
}

//
// Archive
//

func NewArchive(params archive.NewArchiveParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	a, err := s.Archive.NewArchive()
	if err == nil {
		res = archive.NewNewArchiveOK().WithPayload(a)
	} else {
		apires := utils.NewAPIResponseFactory().GetAPIResponse(http.StatusInternalServerError)
		res = archive.NewNewArchiveInternalServerError().WithPayload(apires)
	}

	return res
}

func UpdateArchive(params archive.UpdateArchiveParams) middleware.Responder {
	var res middleware.Responder

	apires := utils.NewAPIResponseFactory().GetAPIResponse(http.StatusNotImplemented)
	res = archive.NewUpdateArchiveMethodNotAllowed().WithPayload(apires)

	return res
}

func GetArchives(params archive.GetArchivesParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	a, err := s.Archive.GetArchives()
	if err == nil {
		res = archive.NewGetArchivesOK().WithPayload(a)
	} else {
		apires := utils.NewAPIResponseFactory().GetAPIResponse(http.StatusNotFound)
		res = archive.NewGetArchivesNotFound().WithPayload(apires)
	}

	return res
}

func GetArchiveByID(params archive.GetArchiveByIDParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	a, err := s.Archive.GetArchiveById(params.ArchiveID)
	if err == nil {
		res = archive.NewGetArchiveByIDOK().WithPayload(a)
	} else {
		apires := utils.NewAPIResponseFactory().GetAPIResponse(http.StatusNotFound)
		res = archive.NewGetArchiveByIDNotFound().WithPayload(apires)
	}

	return res
}
