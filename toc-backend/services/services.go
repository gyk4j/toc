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
	NewBackup() (*models.Backup, *AppError)
	UpdateBackup(backup *models.Backup) (*models.Backup, *AppError)
	GetBackups() ([]*models.Backup, *AppError)
	GetBackupById(id int64) (*models.Backup, *AppError)
}

type IRestorationService interface {
	NewRestoration(backup *models.Backup) (*models.Restoration, *AppError)
	UpdateRestoration(backup *models.Restoration) (*models.Restoration, *AppError)
	GetRestorations() ([]*models.Restoration, *AppError)
	GetRestorationById(id int64) (*models.Restoration, *AppError)
}

type ITransferService interface {
	NewTransfer(backup *models.Backup) (*models.Transfer, *AppError)
	UpdateTransfer(backup *models.Transfer) (*models.Transfer, *AppError)
	GetTransfers() ([]*models.Transfer, *AppError)
	GetTransferById(id int64) (*models.Transfer, *AppError)
}

type ISynchronizationService interface {
	NewSynchronization() (*models.Synchronization, *AppError)
	UpdateSynchronization(synchronization *models.Synchronization) (*models.Synchronization, *AppError)
}

type IQuotaService interface {
	GetQuotas() ([]*models.Quota, *AppError)
}

type ILogService interface {
	NewLog() (*models.Log, *AppError)
	UpdateLog(log *models.Log) (*models.Log, *AppError)
	GetLogById(id int64) (*models.Log, *AppError)
}

type IArchiveService interface {
	NewArchive() (*models.Archive, *AppError)
	UpdateArchive(archive *models.Archive) (*models.Archive, *AppError)
	GetArchives() ([]*models.Archive, *AppError)
	GetArchiveById(id int64) (*models.Archive, *AppError)
}
