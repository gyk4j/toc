package services

import "github.com/gyk4j/toc/toc-backend/models"

type IServer interface {
	NewBackup() *models.Backup
	UpdateBackup(backup *models.Backup) *models.Backup
	GetBackups() []*models.Backup
	GetBackupByID(id int64) *models.Backup

	NewRestoration(backup *models.Backup) *models.Restoration
	UpdateRestoration(backup *models.Restoration) *models.Restoration
	GetRestorations() []*models.Restoration
	GetRestorationByID(id int64) *models.Restoration

	NewTransfer(backup *models.Backup) *models.Transfer
	UpdateTransfer(backup *models.Transfer) *models.Transfer
	GetTransfers() []*models.Transfer
	GetTransferByID(id int64) *models.Transfer

	NewSynchronization() *models.Synchronization
	UpdateSynchronization(synchronization *models.Synchronization) *models.Synchronization

	GetQuotas() []*models.Quota

	NewLog() *models.Log
	UpdateLog(log *models.Log) *models.Log
	GetLogByID(id int64) *models.Log

	NewArchive() *models.Archive
	UpdateArchive(archive *models.Archive) *models.Archive
	GetArchives() []*models.Archive
	GetArchiveByID(id int64) *models.Archive
}
