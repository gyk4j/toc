package repositories

import (
	"log"

	"github.com/gyk4j/toc/toc-backend/models"
)

//
// Backup
//

type BackupRepository struct {
	backups []*models.Backup
}

func (r *BackupRepository) Save(backup *models.Backup) *models.Backup {
	if backup.ID < 0 {
		backup.ID = int64(len(r.backups))
		log.Printf("INFO: Creating new backup (%d)\n", backup.ID)
		r.backups = append(r.backups, backup)
	} else if backup.ID == int64(len(r.backups)) {
		log.Printf("INFO: Adding new backup (%d)\n", backup.ID)
		r.backups = append(r.backups, backup)
	} else if backup.ID > int64(len(r.backups)) {
		log.Printf("WARN: Saving backup (%d) beyond length %d\n", backup.ID, len(r.backups))
		newSlice := make([]*models.Backup, backup.ID+1)
		copy(newSlice, r.backups)
		newSlice[backup.ID] = backup
		r.backups = newSlice
	} else {
		log.Printf("INFO: Updating backup (%d)\n", backup.ID)
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
		log.Printf("INFO: Creating new restoration (%d)\n", restoration.ID)
		r.restorations = append(r.restorations, restoration)
	} else if restoration.ID == int64(len(r.restorations)) {
		log.Printf("INFO: Adding new restoration (%d)\n", restoration.ID)
		r.restorations = append(r.restorations, restoration)
	} else if restoration.ID > int64(len(r.restorations)) {
		log.Printf("WARN: Saving restoration (%d) beyond length %d\n", restoration.ID, len(r.restorations))
		newSlice := make([]*models.Restoration, restoration.ID+1)
		copy(newSlice, r.restorations)
		newSlice[restoration.ID] = restoration
		r.restorations = newSlice
	} else {
		log.Printf("INFO: Updating restoration (%d)\n", restoration.ID)
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
		log.Printf("INFO: Creating new transfer (%d)\n", transfer.ID)
		r.transfers = append(r.transfers, transfer)
	} else if transfer.ID == int64(len(r.transfers)) {
		log.Printf("INFO: Adding new transfer (%d)\n", transfer.ID)
		r.transfers = append(r.transfers, transfer)
	} else if transfer.ID > int64(len(r.transfers)) {
		log.Printf("WARN: Saving transfer (%d) beyond length %d\n", transfer.ID, len(r.transfers))
		newSlice := make([]*models.Transfer, transfer.ID+1)
		copy(newSlice, r.transfers)
		newSlice[transfer.ID] = transfer
		r.transfers = newSlice
	} else {
		log.Printf("INFO: Updating transfer (%d)\n", transfer.ID)
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
		log.Printf("INFO: Creating new synchronization (%d)\n", synchronization.ID)
		r.synchronizations = append(r.synchronizations, synchronization)
	} else if synchronization.ID == int64(len(r.synchronizations)) {
		log.Printf("INFO: Adding new synchronization (%d)\n", synchronization.ID)
		r.synchronizations = append(r.synchronizations, synchronization)
	} else if synchronization.ID > int64(len(r.synchronizations)) {
		log.Printf("WARN: Saving synchronization (%d) beyond length %d\n", synchronization.ID, len(r.synchronizations))
		newSlice := make([]*models.Synchronization, synchronization.ID+1)
		copy(newSlice, r.synchronizations)
		newSlice[synchronization.ID] = synchronization
		r.synchronizations = newSlice
	} else {
		log.Printf("INFO: Updating synchronization (%d)\n", synchronization.ID)
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

func (r *LogRepository) Save(logs *models.Log) *models.Log {
	if logs.ID < 0 {
		logs.ID = int64(len(r.logs))
		log.Printf("INFO: Creating new logs (%d)\n", logs.ID)
		r.logs = append(r.logs, logs)
	} else if logs.ID == int64(len(r.logs)) {
		log.Printf("INFO: Adding new logs (%d)\n", logs.ID)
		r.logs = append(r.logs, logs)
	} else if logs.ID > int64(len(r.logs)) {
		log.Printf("WARN: Saving log (%d) beyond length %d\n", logs.ID, len(r.logs))
		newSlice := make([]*models.Log, logs.ID+1)
		copy(newSlice, r.logs)
		newSlice[logs.ID] = logs
		r.logs = newSlice
	} else {
		log.Printf("INFO: Updating log (%d)\n", logs.ID)
		r.logs[logs.ID] = logs
	}

	return logs
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
		log.Printf("INFO: Creating new archive (%d)\n", archive.ID)
		r.archives = append(r.archives, archive)
	} else if archive.ID == int64(len(r.archives)) {
		log.Printf("INFO: Adding new archive (%d)\n", archive.ID)
		r.archives = append(r.archives, archive)
	} else if archive.ID > int64(len(r.archives)) {
		log.Printf("WARN: Saving archive (%d) beyond length %d\n", archive.ID, len(r.archives))
		newSlice := make([]*models.Archive, archive.ID+1)
		copy(newSlice, r.archives)
		newSlice[archive.ID] = archive
		r.archives = newSlice
	} else {
		log.Printf("INFO: Updating archive (%d)\n", archive.ID)
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
