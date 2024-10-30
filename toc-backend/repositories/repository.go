package repositories

import (
	"github.com/gyk4j/toc/toc-backend/models"
)

type Repository interface {
	NewBackup(backup *models.Backup) bool
	UpdateBackup(backup *models.Backup) bool
	GetBackups() []*models.Backup
	GetBackupByID(id int64) *models.Backup

	NewRestoration(restoration *models.Restoration) bool
	UpdateRestoration(restoration *models.Restoration) bool
	GetRestorations() []*models.Restoration
	GetRestorationByID(id int64) *models.Restoration

	NewTransfer(transfer *models.Transfer) bool
	UpdateTransfer(transfer *models.Transfer) bool
	GetTransfers() []*models.Transfer
	GetTransferByID(id int64) *models.Transfer

	NewSynchronization(synchronization *models.Synchronization) bool
	UpdateSynchronization(synchronization *models.Synchronization) bool

	GetQuotas() []*models.Quota

	NewLog(log *models.Log) bool
	UpdateLog(log *models.Log) bool
	GetLogByID(id int64) *models.Log

	NewArchive(archive *models.Archive) bool
	UpdateArchive(archive *models.Archive) bool
	GetArchives() []*models.Archive
	GetArchiveByID(id int64) *models.Archive
}
