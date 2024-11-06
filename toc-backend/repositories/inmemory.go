package repositories

import "github.com/gyk4j/toc/toc-backend/models"

type InMemory struct{}

//
// Backup
//

func (r *InMemory) NewBackup(backup *models.Backup) bool {
	return false
}

func (r *InMemory) UpdateBackup(backup *models.Backup) bool {
	return false
}

func (r *InMemory) GetBackups() []*models.Backup {
	return nil
}

func (r *InMemory) GetBackupByID(id int64) *models.Backup {
	return nil
}

//
// Restoration
//

func (r *InMemory) NewRestoration(restoration *models.Restoration) bool {
	return false
}

func (r *InMemory) UpdateRestoration(restoration *models.Restoration) bool {
	return false
}

func (r *InMemory) GetRestorations() []*models.Restoration {
	return nil
}

func (r *InMemory) GetRestorationByID(id int64) *models.Restoration {
	return nil
}

//
// Transfer
//

func (r *InMemory) NewTransfer(transfer *models.Transfer) bool {
	return false
}

func (r *InMemory) UpdateTransfer(transfer *models.Transfer) bool {
	return false
}

func (r *InMemory) GetTransfers() []*models.Transfer {
	return nil
}

func (r *InMemory) GetTransferByID(id int64) *models.Transfer {
	return nil
}

//
// Synchronization
//

func (r *InMemory) NewSynchronization(synchronization *models.Synchronization) bool {
	return false
}

func (r *InMemory) UpdateSynchronization(synchronization *models.Synchronization) bool {
	return false
}

func (r *InMemory) GetSynchronizations() []*models.Synchronization {
	return nil
}

func (r *InMemory) GetSynchronizationByID(id int64) *models.Synchronization {
	return nil
}

//
// Quota
//

func (r *InMemory) GetQuotas() []*models.Quota {
	return nil
}

//
// Log
//

func (r *InMemory) NewLog(log *models.Log) bool {
	return false
}

func (r *InMemory) UpdateLog(log *models.Log) bool {
	return false
}

func (r *InMemory) GetLogs() []*models.Log {
	return nil
}

func (r *InMemory) GetLogByID(id int64) *models.Log {
	return nil
}

//
// Archive
//

func (r *InMemory) NewArchive(archive *models.Archive) bool {
	return false
}

func (r *InMemory) UpdateArchive(archive *models.Archive) bool {
	return false
}

func (r *InMemory) GetArchives() []*models.Archive {
	return nil
}

func (r *InMemory) GetArchiveByID(id int64) *models.Archive {
	return nil
}
