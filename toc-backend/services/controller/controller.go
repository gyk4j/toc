package controller

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gyk4j/toc/toc-backend/models"
	"github.com/gyk4j/toc/toc-backend/repositories"
	"github.com/gyk4j/toc/toc-backend/services"
)

//
// Backup
//

type BackupService struct {
	r repositories.IBackupRepository
}

func (s *BackupService) NewBackup() (*models.Backup, *services.AppError) {
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

	return s.r.Save(&b), nil
}

func (s *BackupService) UpdateBackup(backup *models.Backup) (*models.Backup, *services.AppError) {
	if s.r.FindById(backup.ID) != nil {
		return s.r.Save(backup), nil
	} else {
		return nil, &services.AppError{
			Error:   fmt.Errorf("not found: %s", "NewBackup"),
			Message: "not found",
			Code:    int(services.NewBackupNotFound),
		}
	}
}

func (s *BackupService) GetBackups() ([]*models.Backup, *services.AppError) {
	return s.r.FindAll(), nil
}

func (s *BackupService) GetBackupById(id int64) (*models.Backup, *services.AppError) {
	return s.r.FindById(id), nil
}

//
// Restoration
//

type RestorationService struct {
	r repositories.IRestorationRepository
}

func (s *RestorationService) NewRestoration(backup *models.Backup) (*models.Restoration, *services.AppError) {
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

		return s.r.Save(&r), nil
	} else {
		log.Println("[Error] Backup not found")
		return nil, &services.AppError{
			Error:   fmt.Errorf("not found: %s", "NewRestoration"),
			Message: "not found",
			Code:    int(services.NewRestorationNotFound),
		}
	}
}

func (s *RestorationService) UpdateRestoration(restoration *models.Restoration) (*models.Restoration, *services.AppError) {
	if s.r.FindById(restoration.ID) != nil {
		return s.r.Save(restoration), nil
	} else {
		return nil, &services.AppError{
			Error:   fmt.Errorf("not found: %s", "UpdateRestoration"),
			Message: "not found",
			Code:    int(services.UpdateRestorationNotFound),
		}
	}
}

func (s *RestorationService) GetRestorations() ([]*models.Restoration, *services.AppError) {
	return s.r.FindAll(), nil
}

func (s *RestorationService) GetRestorationById(id int64) (*models.Restoration, *services.AppError) {
	return s.r.FindById(id), nil
}

//
// Transfer
//

type TransferService struct {
	r repositories.ITransferRepository
}

func (s *TransferService) NewTransfer(backup *models.Backup) (*models.Transfer, *services.AppError) {
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

		return s.r.Save(&t), nil
	} else {
		log.Println("[Error] Backup not found")
		return nil, &services.AppError{
			Error:   fmt.Errorf("not found: %s", "NewTransfer"),
			Message: "not found",
			Code:    int(services.NewTransferNotFound),
		}
	}
}

func (s *TransferService) UpdateTransfer(transfer *models.Transfer) (*models.Transfer, *services.AppError) {
	if s.r.FindById(transfer.ID) != nil {
		return s.r.Save(transfer), nil
	} else {
		return nil, &services.AppError{
			Error:   fmt.Errorf("not found: %s", "UpdateTransfer"),
			Message: "not found",
			Code:    int(services.UpdateTransferNotFound),
		}
	}
}

func (s *TransferService) GetTransfers() ([]*models.Transfer, *services.AppError) {
	return s.r.FindAll(), nil
}

func (s *TransferService) GetTransferById(id int64) (*models.Transfer, *services.AppError) {
	return s.r.FindById(id), nil
}

//
// Synchronization
//

type SynchronizationService struct {
	r repositories.ISynchronizationRepository
}

func (s *SynchronizationService) NewSynchronization() (*models.Synchronization, *services.AppError) {
	sync := models.Synchronization{
		ID:     -1,
		Status: models.LogStatusQueued,
		Time:   strfmt.DateTime(time.Now()),
	}

	return s.r.Save(&sync), nil
}

func (s *SynchronizationService) UpdateSynchronization(synchronization *models.Synchronization) (*models.Synchronization, *services.AppError) {
	if s.r.FindById(synchronization.ID) != nil {
		return s.r.Save(synchronization), nil
	} else {
		return nil, &services.AppError{
			Error:   fmt.Errorf("not found: %s", "UpdateSynchronization"),
			Message: "not found",
			Code:    int(services.UpdateSynchronizationNotFound),
		}
	}
}

//
// Quota
//

type QuotaService struct {
	r repositories.IQuotaRepository
}

func (s *QuotaService) GetQuotas() ([]*models.Quota, *services.AppError) {
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

	return s.r.FindAll(), nil
}

//
// Log
//

type LogService struct {
	r repositories.ILogRepository
}

func (s *LogService) NewLog() (*models.Log, *services.AppError) {
	l := models.Log{
		ID:     -1,
		Status: models.LogStatusQueued,
		Time:   strfmt.DateTime(time.Now()),
		URL:    "/v1/logs/" + strconv.FormatInt(0, 10),
	}

	return s.r.Save(&l), nil
}

func (s *LogService) UpdateLog(log *models.Log) (*models.Log, *services.AppError) {
	if s.r.FindById(log.ID) != nil {
		return s.r.Save(log), nil
	} else {
		return nil, &services.AppError{
			Error:   fmt.Errorf("not found: %s", "UpdateLog"),
			Message: "not found",
			Code:    int(services.UpdateLogNotFound),
		}
	}
}

func (s *LogService) GetLogById(id int64) (*models.Log, *services.AppError) {
	return s.r.FindById(id), nil
}

//
// Archive
//

type ArchiveService struct {
	r repositories.IArchiveRepository
}

func (s *ArchiveService) NewArchive() (*models.Archive, *services.AppError) {
	a := models.Archive{
		ID:     -1,
		Status: models.ArchiveStatusQueued,
		Time:   strfmt.DateTime(time.Now()),
		URL:    "/v1/archives/" + strconv.FormatInt(0, 10),
	}

	return s.r.Save(&a), nil
}

func (s *ArchiveService) UpdateArchive(archive *models.Archive) (*models.Archive, *services.AppError) {
	if s.r.FindById(archive.ID) != nil {
		return s.r.Save(archive), nil
	} else {
		return nil, &services.AppError{
			Error:   fmt.Errorf("not found: %s", "UpdateArchive"),
			Message: "not found",
			Code:    int(services.UpdateArchiveNotFound),
		}
	}
}

func (s *ArchiveService) GetArchives() ([]*models.Archive, *services.AppError) {
	return s.r.FindAll(), nil
}

func (s *ArchiveService) GetArchiveById(id int64) (*models.Archive, *services.AppError) {
	return s.r.FindById(id), nil
}
