package server

import (
	"log"

	"github.com/gyk4j/toc/toc-backend/models"
)

//
// Backup
//

type BackupService struct {
	// r repositories.IBackupRepository
}

func (s *BackupService) NewBackup() *models.Backup {
	log.Printf("Server: %s", "NewBackup")
	return nil
}

func (s *BackupService) UpdateBackup(backup *models.Backup) *models.Backup {
	log.Printf("Server: %s", "UpdateBackup")
	return nil
}

func (s *BackupService) GetBackups() []*models.Backup {
	log.Printf("Server: %s", "GetBackups")
	return nil
}

func (s *BackupService) GetBackupById(id int64) *models.Backup {
	log.Printf("Server: %s", "GetBackupByID")
	return nil
}

//
// Restoration
//

type RestorationService struct {
	// r repositories.IRestorationRepository
}

func (s *RestorationService) NewRestoration(backup *models.Backup) *models.Restoration {
	log.Printf("Not implemented: %s", "NewRestoration")
	return nil
}

func (s *RestorationService) UpdateRestoration(restoration *models.Restoration) *models.Restoration {
	log.Printf("Not implemented: %s", "UpdateRestoration")
	return nil
}

func (s *RestorationService) GetRestorations() []*models.Restoration {
	log.Printf("Not implemented: %s", "GetRestorations")
	return nil
}

func (s *RestorationService) GetRestorationById(id int64) *models.Restoration {
	log.Printf("Not implemented: %s", "GetRestorationByID")
	return nil
}

//
// Transfer
//

type TransferService struct {
	// r repositories.ITransferRepository
}

func (s *TransferService) NewTransfer(backup *models.Backup) *models.Transfer {
	log.Printf("Not implemented: %s", "NewTransfer")
	return nil
}

func (s *TransferService) UpdateTransfer(transfer *models.Transfer) *models.Transfer {
	log.Printf("Not implemented: %s", "UpdateTransfer")
	return nil
}

func (s *TransferService) GetTransfers() []*models.Transfer {
	log.Printf("Not implemented: %s", "GetTransfers")
	return nil
}

func (s *TransferService) GetTransferById(id int64) *models.Transfer {
	log.Printf("Not implemented: %s", "GetTransferByID")
	return nil
}

//
// Synchronization
//

type SynchronizationService struct {
	// r repositories.ISynchronizationRepository
}

func (s *SynchronizationService) NewSynchronization() *models.Synchronization {
	log.Printf("Not implemented: %s", "NewSynchronization")
	return nil
}

func (s *SynchronizationService) UpdateSynchronization(synchronization *models.Synchronization) *models.Synchronization {
	log.Printf("Not implemented: %s", "UpdateSynchronization")
	return nil
}

//
// Quota
//

type QuotaService struct {
	// r repositories.IQuotaRepository
}

func (s *QuotaService) GetQuotas() []*models.Quota {
	log.Printf("Not implemented: %s", "GetQuotas")
	return nil
}

//
// Log
//

type LogService struct {
	// r repositories.ILogRepository
}

func (s *LogService) NewLog() *models.Log {
	log.Printf("Not implemented: %s", "NewLog")
	return nil
}

func (s *LogService) UpdateLog(logs *models.Log) *models.Log {
	log.Printf("Not implemented: %s", "UpdateLog")
	return nil
}

func (s *LogService) GetLogById(id int64) *models.Log {
	log.Printf("Not implemented: %s", "GetLogByID")
	return nil
}

//
// Archive
//

type ArchiveService struct {
	// r repositories.IArchiveRepository
}

func (s *ArchiveService) NewArchive() *models.Archive {
	log.Printf("Not implemented: %s", "NewArchive")
	return nil
}

func (s *ArchiveService) UpdateArchive(archive *models.Archive) *models.Archive {
	log.Printf("Not implemented: %s", "UpdateArchive")
	return nil
}

func (s *ArchiveService) GetArchives() []*models.Archive {
	log.Printf("Not implemented: %s", "GetArchives")
	return nil
}

func (s *ArchiveService) GetArchiveById(id int64) *models.Archive {
	log.Printf("Not implemented: %s", "GetArchiveByID")
	return nil
}
