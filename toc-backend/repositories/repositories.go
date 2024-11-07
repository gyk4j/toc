package repositories

import (
	"github.com/gyk4j/toc/toc-backend/models"
)

type IBackupRepository interface {
	Save(backup *models.Backup) *models.Backup
	FindAll() []*models.Backup
	FindById(id int64) *models.Backup
}

type IRestorationRepository interface {
	Save(restoration *models.Restoration) *models.Restoration
	FindAll() []*models.Restoration
	FindById(id int64) *models.Restoration
}

type ITransferRepository interface {
	Save(transfer *models.Transfer) *models.Transfer
	FindAll() []*models.Transfer
	FindById(id int64) *models.Transfer
}

type ISynchronizationRepository interface {
	Save(synchronization *models.Synchronization) *models.Synchronization
	FindAll() []*models.Synchronization
	FindById(id int64) *models.Synchronization
}

type IQuotaRepository interface {
	FindAll() []*models.Quota
}

type ILogRepository interface {
	Save(log *models.Log) *models.Log
	FindById(id int64) *models.Log
}

type IArchiveRepository interface {
	Save(archive *models.Archive) *models.Archive
	FindAll() []*models.Archive
	FindById(id int64) *models.Archive
}
