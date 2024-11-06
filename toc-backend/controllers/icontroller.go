package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/archive"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/backup"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/log"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/quota"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/restoration"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/synchronization"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/transfer"
)

type IController interface {
	// Backup
	NewBackup(params backup.NewBackupParams) middleware.Responder
	UpdateBackup(params backup.UpdateBackupParams) middleware.Responder
	GetBackups(params backup.GetBackupsParams) middleware.Responder
	GetBackupByID(params backup.GetBackupByIDParams) middleware.Responder

	// Restoration
	NewRestoration(params restoration.NewRestorationParams) middleware.Responder
	UpdateRestoration(params restoration.UpdateRestorationParams) middleware.Responder
	GetRestorations(params restoration.GetRestorationsParams) middleware.Responder
	GetRestorationByID(params restoration.GetRestorationByIDParams) middleware.Responder

	// Transfer
	NewTransfer(params transfer.NewTransferParams) middleware.Responder
	UpdateTransfer(params transfer.UpdateTransferParams) middleware.Responder
	GetTransfers(params transfer.GetTransfersParams) middleware.Responder
	GetTransferByID(params transfer.GetTransferByIDParams) middleware.Responder

	// Synchronization
	NewSynchronization(params synchronization.NewSynchronizationParams) middleware.Responder
	UpdateSynchronization(params synchronization.UpdateSynchronizationParams) middleware.Responder

	// Quota
	GetQuotas(params quota.GetQuotasParams) middleware.Responder

	// Log
	ExportLog(params log.ExportLogParams) middleware.Responder
	UpdateLog(params log.UpdateLogParams) middleware.Responder

	// Archive
	NewArchive(params archive.NewArchiveParams) middleware.Responder
	UpdateArchive(params archive.UpdateArchiveParams) middleware.Responder
	GetArchives(params archive.GetArchivesParams) middleware.Responder
	GetArchiveByID(params archive.GetArchiveByIDParams) middleware.Responder
}
