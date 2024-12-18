package server

import (
	"fmt"
	"log"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gyk4j/toc/toc-backend/models"
	"github.com/gyk4j/toc/toc-backend/repositories"
	"github.com/gyk4j/toc/toc-backend/services"
)

var servers = []string{"web", "file", "db", "mail"}

func NewServer(name string) (server *services.Server) {
	return &services.Server{
		Name:            name,
		Backup:          &BackupService{},
		Restoration:     &RestorationService{},
		Transfer:        &TransferService{},
		Synchronization: &SynchronizationService{},
		Quota:           &QuotaService{},
		Log:             &LogService{},
		Archive:         &ArchiveService{},
	}
}

//
// Backup
//

type BackupService struct {
	r repositories.IBackupRepository
}

func (s *BackupService) NewBackup() (*models.Backup, *services.AppError) {
	var err *services.AppError

	b := models.Backup{
		ID:        -1,
		Time:      strfmt.DateTime(time.Now()),
		Snapshots: &models.Snapshots{},
	}

	// Get ID issued first.
	o := s.r.Save(&b)
	i := o.ID

	for _, server := range servers {
		ss := models.Snapshot{
			ID:     int64(i),
			File:   fmt.Sprintf("/toc/archives/%d-%s.tar.gz", i, server),
			Status: models.SnapshotStatusQueued,
		}

		switch server {
		case "db":
			b.Snapshots.Db = &ss
		case "file":
			b.Snapshots.File = &ss
		case "mail":
			b.Snapshots.Mail = &ss
		case "web":
			b.Snapshots.Web = &ss
		default:
			log.Printf("Unknown server: %s", server)
		}
	}

	// Update/resave with snapshots
	o = s.r.Save(&b)

	if o == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("internal server error: %s", "NewBackup"),
			Message: "internal server error",
			Code:    int(services.NewBackupInternalServerError),
		}
	}

	return o, err
}

func (s *BackupService) UpdateBackup(backup *models.Backup) (*models.Backup, *services.AppError) {
	log.Printf("not implemented: %s", "UpdateBackup")
	return nil, &services.AppError{
		Error:   fmt.Errorf("not implemented: %s", "UpdateBackup"),
		Message: "not implemented",
		Code:    int(services.UpdateBackupForbidden),
	}
}

func (s *BackupService) GetBackups() ([]*models.Backup, *services.AppError) {
	log.Printf("not implemented: %s", "GetBackups")
	return nil, &services.AppError{
		Error:   fmt.Errorf("not implemented: %s", "GetBackups"),
		Message: "not implemented",
		Code:    int(services.GetBackupsForbidden),
	}
}

func (s *BackupService) GetBackupById(id int64) (*models.Backup, *services.AppError) {
	log.Printf("not implemented: %s", "GetBackupById")
	return nil, &services.AppError{
		Error:   fmt.Errorf("not implemented: %s", "GetBackupById"),
		Message: "not implemented",
		Code:    int(services.GetBackupByIdForbidden),
	}
}

//
// Restoration
//

type RestorationService struct {
	// r repositories.IRestorationRepository
}

func (s *RestorationService) NewRestoration(backup *models.Backup) (*models.Restoration, *services.AppError) {
	log.Printf("not implemented: %s", "NewRestoration")
	return nil, &services.AppError{
		Error:   fmt.Errorf("not implemented: %s", "NewRestoration"),
		Message: "not implemented",
		Code:    int(services.NewRestorationForbidden),
	}
}

func (s *RestorationService) UpdateRestoration(restoration *models.Restoration) (*models.Restoration, *services.AppError) {
	log.Printf("not implemented: %s", "UpdateRestoration")
	return nil, &services.AppError{
		Error:   fmt.Errorf("not implemented: %s", "UpdateRestoration"),
		Message: "not implemented",
		Code:    int(services.UpdateRestorationForbidden),
	}
}

func (s *RestorationService) GetRestorations() ([]*models.Restoration, *services.AppError) {
	log.Printf("not implemented: %s", "GetRestorations")
	return nil, &services.AppError{
		Error:   fmt.Errorf("not implemented: %s", "GetRestorations"),
		Message: "not implemented",
		Code:    int(services.GetRestorationsForbidden),
	}
}

func (s *RestorationService) GetRestorationById(id int64) (*models.Restoration, *services.AppError) {
	log.Printf("not implemented: %s", "GetRestorationById")
	return nil, &services.AppError{
		Error:   fmt.Errorf("not implemented: %s", "GetRestorationById"),
		Message: "not implemented",
		Code:    int(services.GetRestorationByIdForbidden),
	}
}

//
// Transfer
//

type TransferService struct {
	// r repositories.ITransferRepository
}

func (s *TransferService) NewTransfer(backup *models.Backup) (*models.Transfer, *services.AppError) {
	log.Printf("not implemented: %s", "NewTransfer")
	return nil, &services.AppError{
		Error:   fmt.Errorf("not implemented: %s", "NewTransfer"),
		Message: "not implemented",
		Code:    int(services.NewTransferForbidden),
	}
}

func (s *TransferService) UpdateTransfer(transfer *models.Transfer) (*models.Transfer, *services.AppError) {
	log.Printf("not implemented: %s", "UpdateTransfer")
	return nil, &services.AppError{
		Error:   fmt.Errorf("not implemented: %s", "UpdateTransfer"),
		Message: "not implemented",
		Code:    int(services.UpdateTransferForbidden),
	}
}

func (s *TransferService) GetTransfers() ([]*models.Transfer, *services.AppError) {
	log.Printf("not implemented: %s", "GetTransfers")
	return nil, &services.AppError{
		Error:   fmt.Errorf("not implemented: %s", "GetTransfers"),
		Message: "not implemented",
		Code:    int(services.GetTransfersForbidden),
	}
}

func (s *TransferService) GetTransferById(id int64) (*models.Transfer, *services.AppError) {
	log.Printf("not implemented: %s", "GetTransferById")
	return nil, &services.AppError{
		Error:   fmt.Errorf("not implemented: %s", "GetTransferById"),
		Message: "not implemented",
		Code:    int(services.GetTransferByIdForbidden),
	}
}

//
// Synchronization
//

type SynchronizationService struct {
	// r repositories.ISynchronizationRepository
}

func (s *SynchronizationService) NewSynchronization() (*models.Synchronization, *services.AppError) {
	log.Printf("not implemented: %s", "NewSynchronization")
	return nil, &services.AppError{
		Error:   fmt.Errorf("not implemented: %s", "NewSynchronization"),
		Message: "not implemented",
		Code:    int(services.NewSynchronizationForbidden),
	}
}

func (s *SynchronizationService) UpdateSynchronization(synchronization *models.Synchronization) (*models.Synchronization, *services.AppError) {
	log.Printf("not implemented: %s", "UpdateSynchronization")
	return nil, &services.AppError{
		Error:   fmt.Errorf("not implemented: %s", "UpdateSynchronization"),
		Message: "not implemented",
		Code:    int(services.UpdateSynchronizationForbidden),
	}
}

//
// Quota
//

type QuotaService struct {
	// r repositories.IQuotaRepository
}

func (s *QuotaService) GetQuotas() ([]*models.Quota, *services.AppError) {
	log.Printf("not implemented: %s", "GetQuotas")
	return nil, &services.AppError{
		Error:   fmt.Errorf("not implemented: %s", "GetQuotas"),
		Message: "not implemented",
		Code:    int(services.GetQuotasForbidden),
	}
}

//
// Log
//

type LogService struct {
	// r repositories.ILogRepository
}

func (s *LogService) NewLog() (*models.Log, *services.AppError) {
	log.Printf("not implemented: %s", "NewLog")
	return nil, &services.AppError{
		Error:   fmt.Errorf("not implemented: %s", "NewLog"),
		Message: "not implemented",
		Code:    int(services.NewLogForbidden),
	}
}

func (s *LogService) UpdateLog(logs *models.Log) (*models.Log, *services.AppError) {
	log.Printf("not implemented: %s", "UpdateLog")
	return nil, &services.AppError{
		Error:   fmt.Errorf("not implemented: %s", "UpdateLog"),
		Message: "not implemented",
		Code:    int(services.UpdateLogForbidden),
	}
}

func (s *LogService) GetLogById(id int64) (*models.Log, *services.AppError) {
	log.Printf("not implemented: %s", "GetLogById")
	return nil, &services.AppError{
		Error:   fmt.Errorf("not implemented: %s", "GetLogById"),
		Message: "not implemented",
		Code:    int(services.GetLogByIdForbidden),
	}
}

//
// Archive
//

type ArchiveService struct {
	// r repositories.IArchiveRepository
}

func (s *ArchiveService) NewArchive() (*models.Archive, *services.AppError) {
	log.Printf("not implemented: %s", "NewArchive")
	return nil, &services.AppError{
		Error:   fmt.Errorf("not implemented: %s", "NewArchive"),
		Message: "not implemented",
		Code:    int(services.NewArchiveForbidden),
	}
}

func (s *ArchiveService) UpdateArchive(archive *models.Archive) (*models.Archive, *services.AppError) {
	log.Printf("not implemented: %s", "UpdateArchive")
	return nil, &services.AppError{
		Error:   fmt.Errorf("not implemented: %s", "UpdateArchive"),
		Message: "not implemented",
		Code:    int(services.UpdateArchiveForbidden),
	}
}

func (s *ArchiveService) GetArchives() ([]*models.Archive, *services.AppError) {
	log.Printf("not implemented: %s", "GetArchives")
	return nil, &services.AppError{
		Error:   fmt.Errorf("not implemented: %s", "GetArchives"),
		Message: "not implemented",
		Code:    int(services.GetArchivesForbidden),
	}
}

func (s *ArchiveService) GetArchiveById(id int64) (*models.Archive, *services.AppError) {
	log.Printf("not implemented: %s", "GetArchiveById")
	return nil, &services.AppError{
		Error:   fmt.Errorf("not implemented: %s", "GetArchiveById"),
		Message: "not implemented",
		Code:    int(services.GetArchiveByIdForbidden),
	}
}
