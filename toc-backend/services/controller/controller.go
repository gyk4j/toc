package controller

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gyk4j/toc/toc-backend/models"
	"github.com/gyk4j/toc/toc-backend/repositories"
)

//
// Backup
//

type BackupService struct {
	r repositories.IBackupRepository
}

func (s *BackupService) NewBackup() *models.Backup {
	b := models.Backup{
		ID:        -1,
		Time:      strfmt.DateTime(time.Now()),
		Snapshots: make([]*models.Snapshot, 0),
	}

	servers := []string{"web", "file", "db", "mail"}

	for i, s := range servers {
		ss := models.Snapshot{
			ID:       int64(i),
			File:     fmt.Sprintf("/toc/archives/%s-%d.tar.gz", s, i),
			Status:   models.SnapshotStatusQueued,
			Complete: false,
		}
		b.Snapshots = append(b.Snapshots, &ss)
	}

	return s.r.Save(&b)
}

func (s *BackupService) UpdateBackup(backup *models.Backup) *models.Backup {
	if s.r.FindById(backup.ID) != nil {
		return s.r.Save(backup)
	} else {
		return nil
	}
}

func (s *BackupService) GetBackups() []*models.Backup {
	return s.r.FindAll()
}

func (s *BackupService) GetBackupById(id int64) *models.Backup {
	return s.r.FindById(id)
}

//
// Restoration
//

type RestorationService struct {
	r repositories.IRestorationRepository
}

func (s *RestorationService) NewRestoration(backup *models.Backup) *models.Restoration {
	b := backup // Trust client data?
	//b := GetBackupByID((*backup).ID) // More robust method

	// Backup must exists before it can be restored.
	// If an invalid ID is provided, then restoration is skipped.
	if b != nil {
		log.Printf("Backup found: %d = %s\n", (*b).ID, (*b).Time)
		r := models.Restoration{
			ID:     -1,
			Backup: b,
			Status: models.RestorationStatusQueued,
		}

		return s.r.Save(&r)
	} else {
		log.Println("[Error] Backup not found")
		return nil
	}
}

func (s *RestorationService) UpdateRestoration(restoration *models.Restoration) *models.Restoration {
	if s.r.FindById(restoration.ID) != nil {
		return s.r.Save(restoration)
	} else {
		return nil
	}
}

func (s *RestorationService) GetRestorations() []*models.Restoration {
	return s.r.FindAll()
}

func (s *RestorationService) GetRestorationById(id int64) *models.Restoration {
	return s.r.FindById(id)
}

//
// Transfer
//

type TransferService struct {
	r repositories.ITransferRepository
}

func (s *TransferService) NewTransfer(backup *models.Backup) *models.Transfer {
	b := backup // Trust client data?
	//b := GetBackupByID((*backup).ID) // More robust method

	// Backup must exists before it can be restored.
	// If an invalid ID is provided, then restoration is skipped.
	if b != nil {
		log.Printf("Backup found: %d = %s\n", (*b).ID, (*b).Time)
		t := models.Transfer{
			ID:     -1,
			Backup: b,
			Status: models.RestorationStatusQueued,
		}

		return s.r.Save(&t)
	} else {
		log.Println("[Error] Backup not found")
		return nil
	}
}

func (s *TransferService) UpdateTransfer(transfer *models.Transfer) *models.Transfer {
	if s.r.FindById(transfer.ID) != nil {
		return s.r.Save(transfer)
	} else {
		return nil
	}
}

func (s *TransferService) GetTransfers() []*models.Transfer {
	return s.r.FindAll()
}

func (s *TransferService) GetTransferById(id int64) *models.Transfer {
	return s.r.FindById(id)
}

//
// Synchronization
//

type SynchronizationService struct {
	r repositories.ISynchronizationRepository
}

func (s *SynchronizationService) NewSynchronization() *models.Synchronization {
	sync := models.Synchronization{
		ID:     -1,
		Status: models.LogStatusQueued,
		Time:   strfmt.DateTime(time.Now()),
	}

	return s.r.Save(&sync)
}

func (s *SynchronizationService) UpdateSynchronization(synchronization *models.Synchronization) *models.Synchronization {
	if s.r.FindById(synchronization.ID) != nil {
		return s.r.Save(synchronization)
	} else {
		return nil
	}
}

//
// Quota
//

type QuotaService struct {
	r repositories.IQuotaRepository
}

func (s *QuotaService) GetQuotas() []*models.Quota {
	/*
		q1 := models.Quota{
			ID:   -1,
			Path: "/home/a/q1",
			Soft: rand.Int63(),
			Hard: rand.Int63(),
		}

		q2 := models.Quota{
			ID:   -1,
			Path: "/home/b/q2",
			Soft: rand.Int63(),
			Hard: rand.Int63(),
		}
	*/

	return s.r.FindAll()
}

//
// Log
//

type LogService struct {
	r repositories.ILogRepository
}

func (s *LogService) NewLog() *models.Log {
	l := models.Log{
		ID:     -1,
		Status: models.LogStatusQueued,
		Time:   strfmt.DateTime(time.Now()),
		URL:    "/v1/logs/" + strconv.FormatInt(0, 10),
	}

	return s.r.Save(&l)
}

func (s *LogService) UpdateLog(log *models.Log) *models.Log {
	if s.r.FindById(log.ID) != nil {
		return s.r.Save(log)
	} else {
		return nil
	}
}

func (s *LogService) GetLogById(id int64) *models.Log {
	return s.r.FindById(id)
}

//
// Archive
//

type ArchiveService struct {
	r repositories.IArchiveRepository
}

func (s *ArchiveService) NewArchive() *models.Archive {
	a := models.Archive{
		ID:     -1,
		Status: models.ArchiveStatusQueued,
		Time:   strfmt.DateTime(time.Now()),
		URL:    "/v1/archives/" + strconv.FormatInt(0, 10),
	}

	return s.r.Save(&a)
}

func (s *ArchiveService) UpdateArchive(archive *models.Archive) *models.Archive {
	if s.r.FindById(archive.ID) != nil {
		return s.r.Save(archive)
	} else {
		return nil
	}
}

func (s *ArchiveService) GetArchives() []*models.Archive {
	return s.r.FindAll()
}

func (s *ArchiveService) GetArchiveById(id int64) *models.Archive {
	return s.r.FindById(id)
}
