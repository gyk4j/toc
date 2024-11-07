package services

import "github.com/gyk4j/toc/toc-backend/models"

type Server struct {
	Backup          IBackupService
	Restoration     IRestorationService
	Transfer        ITransferService
	Synchronization ISynchronizationService
	Quota           IQuotaService
	Log             ILogService
	Archive         IArchiveService
}

type IBackupService interface {
	NewBackup() *models.Backup
	UpdateBackup(backup *models.Backup) *models.Backup
	GetBackups() []*models.Backup
	GetBackupById(id int64) *models.Backup
}

type IRestorationService interface {
	NewRestoration(backup *models.Backup) *models.Restoration
	UpdateRestoration(backup *models.Restoration) *models.Restoration
	GetRestorations() []*models.Restoration
	GetRestorationById(id int64) *models.Restoration
}

type ITransferService interface {
	NewTransfer(backup *models.Backup) *models.Transfer
	UpdateTransfer(backup *models.Transfer) *models.Transfer
	GetTransfers() []*models.Transfer
	GetTransferById(id int64) *models.Transfer
}

type ISynchronizationService interface {
	NewSynchronization() *models.Synchronization
	UpdateSynchronization(synchronization *models.Synchronization) *models.Synchronization
}

type IQuotaService interface {
	GetQuotas() []*models.Quota
}

type ILogService interface {
	NewLog() *models.Log
	UpdateLog(log *models.Log) *models.Log
	GetLogById(id int64) *models.Log
}

type IArchiveService interface {
	NewArchive() *models.Archive
	UpdateArchive(archive *models.Archive) *models.Archive
	GetArchives() []*models.Archive
	GetArchiveById(id int64) *models.Archive
}
