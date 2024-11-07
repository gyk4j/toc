package repositories

import "github.com/gyk4j/toc/toc-backend/models"

//
// Backup
//

type BackupRepository struct {
	backups []*models.Backup
}

func (r *BackupRepository) Save(backup *models.Backup) *models.Backup {
	if backup.ID < 0 {
		backup.ID = int64(len(r.backups))
		r.backups = append(r.backups, backup)
	} else if backup.ID >= int64(len(r.backups)) {
		newSlice := make([]*models.Backup, backup.ID+1)
		copy(newSlice, r.backups)
		newSlice[backup.ID] = backup
		r.backups = newSlice
	} else {
		r.backups[backup.ID] = backup
	}

	return backup
}

func (r *BackupRepository) FindAll() []*models.Backup {
	return r.backups
}

func (r *BackupRepository) FindById(id int64) *models.Backup {
	if id < 0 || id >= int64(len(r.backups)) {
		return nil
	} else {
		return r.backups[id]
	}
}

//
// Restoration
//

type RestorationRepository struct {
	restorations []*models.Restoration
}

func (r *RestorationRepository) Save(restoration *models.Restoration) *models.Restoration {
	if restoration.ID < 0 {
		restoration.ID = int64(len(r.restorations))
		r.restorations = append(r.restorations, restoration)
	} else if restoration.ID >= int64(len(r.restorations)) {
		newSlice := make([]*models.Restoration, restoration.ID+1)
		copy(newSlice, r.restorations)
		newSlice[restoration.ID] = restoration
		r.restorations = newSlice
	} else {
		r.restorations[restoration.ID] = restoration
	}

	return restoration
}

func (r *RestorationRepository) FindAll() []*models.Restoration {
	return r.restorations
}

func (r *RestorationRepository) FindById(id int64) *models.Restoration {
	if id < 0 || id >= int64(len(r.restorations)) {
		return nil
	} else {
		return r.restorations[id]
	}
}

//
// Transfer
//

type TransferRepository struct {
	transfers []*models.Transfer
}

func (r *TransferRepository) Save(transfer *models.Transfer) *models.Transfer {
	if transfer.ID < 0 {
		transfer.ID = int64(len(r.transfers))
		r.transfers = append(r.transfers, transfer)
	} else if transfer.ID >= int64(len(r.transfers)) {
		newSlice := make([]*models.Transfer, transfer.ID+1)
		copy(newSlice, r.transfers)
		newSlice[transfer.ID] = transfer
		r.transfers = newSlice
	} else {
		r.transfers[transfer.ID] = transfer
	}

	return transfer
}

func (r *TransferRepository) FindAll() []*models.Transfer {
	return r.transfers
}

func (r *TransferRepository) FindById(id int64) *models.Transfer {
	if id < 0 || id >= int64(len(r.transfers)) {
		return nil
	} else {
		return r.transfers[id]
	}
}

//
// Synchronization
//

type SynchronizationRepository struct {
	synchronizations []*models.Synchronization
}

func (r *SynchronizationRepository) Save(synchronization *models.Synchronization) *models.Synchronization {
	if synchronization.ID < 0 {
		synchronization.ID = int64(len(r.synchronizations))
		r.synchronizations = append(r.synchronizations, synchronization)
	} else if synchronization.ID >= int64(len(r.synchronizations)) {
		newSlice := make([]*models.Synchronization, synchronization.ID+1)
		copy(newSlice, r.synchronizations)
		newSlice[synchronization.ID] = synchronization
		r.synchronizations = newSlice
	} else {
		r.synchronizations[synchronization.ID] = synchronization
	}

	return synchronization
}

func (r *SynchronizationRepository) FindAll() []*models.Synchronization {
	return r.synchronizations
}

func (r *SynchronizationRepository) FindById(id int64) *models.Synchronization {
	if id < 0 || id >= int64(len(r.synchronizations)) {
		return nil
	} else {
		return r.synchronizations[id]
	}
}

//
// Quota
//

type QuotaRepository struct {
	quotas []*models.Quota
}

func (r *QuotaRepository) FindAll() []*models.Quota {
	return r.quotas
}

//
// Log
//

type LogRepository struct {
	logs []*models.Log
}

func (r *LogRepository) Save(log *models.Log) *models.Log {
	if log.ID < 0 {
		log.ID = int64(len(r.logs))
		r.logs = append(r.logs, log)
	} else if log.ID >= int64(len(r.logs)) {
		newSlice := make([]*models.Log, log.ID+1)
		copy(newSlice, r.logs)
		newSlice[log.ID] = log
		r.logs = newSlice
	} else {
		r.logs[log.ID] = log
	}

	return log
}

func (r *LogRepository) FindById(id int64) *models.Log {
	if id < 0 || id >= int64(len(r.logs)) {
		return nil
	} else {
		return r.logs[id]
	}
}

//
// Archive
//

type ArchiveRepository struct {
	archives []*models.Archive
}

func (r *ArchiveRepository) Save(archive *models.Archive) *models.Archive {
	if archive.ID < 0 {
		archive.ID = int64(len(r.archives))
		r.archives = append(r.archives, archive)
	} else if archive.ID >= int64(len(r.archives)) {
		newSlice := make([]*models.Archive, archive.ID+1)
		copy(newSlice, r.archives)
		newSlice[archive.ID] = archive
		r.archives = newSlice
	} else {
		r.archives[archive.ID] = archive
	}

	return archive
}

func (r *ArchiveRepository) FindAll() []*models.Archive {
	return r.archives
}

func (r *ArchiveRepository) FindById(id int64) *models.Archive {
	if id < 0 || id >= int64(len(r.archives)) {
		return nil
	} else {
		return r.archives[id]
	}
}
