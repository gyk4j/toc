package services

import (
	"log"

	"github.com/gyk4j/toc/toc-backend/models"
)

type Server struct {
}

/*
 * Backup
 */

func (s *Server) NewBackup() *models.Backup {
	log.Printf("Server: %s", "NewBackup")
	return nil
}

func (s *Server) UpdateBackup(backup *models.Backup) *models.Backup {
	log.Printf("Server: %s", "UpdateBackup")
	return nil
}

func (s *Server) GetBackups() []*models.Backup {
	log.Printf("Server: %s", "GetBackups")
	return nil
}

func (s *Server) GetBackupByID(id int64) *models.Backup {
	log.Printf("Server: %s", "GetBackupByID")
	return nil
}

/*
 * Restoration
 */

func (s *Server) NewRestoration(backup *models.Backup) *models.Restoration {
	log.Printf("Not implemented: %s", "NewRestoration")
	return nil
}

func (s *Server) UpdateRestoration(restoration *models.Restoration) *models.Restoration {
	log.Printf("Not implemented: %s", "UpdateRestoration")
	return nil
}

func (s *Server) GetRestorations() []*models.Restoration {
	log.Printf("Not implemented: %s", "GetRestorations")
	return nil
}

func (s *Server) GetRestorationByID(id int64) *models.Restoration {
	log.Printf("Not implemented: %s", "GetRestorationByID")
	return nil
}

/*
 * Transfer
 */

func (s *Server) NewTransfer(backup *models.Backup) *models.Transfer {
	log.Printf("Not implemented: %s", "NewTransfer")
	return nil
}

func (s *Server) UpdateTransfer(transfer *models.Transfer) *models.Transfer {
	log.Printf("Not implemented: %s", "UpdateTransfer")
	return nil
}

func (s *Server) GetTransfers() []*models.Transfer {
	log.Printf("Not implemented: %s", "GetTransfers")
	return nil
}

func (s *Server) GetTransferByID(id int64) *models.Transfer {
	log.Printf("Not implemented: %s", "GetTransferByID")
	return nil
}

/*
 * Synchronization
 */

func (s *Server) NewSynchronization() *models.Synchronization {
	log.Printf("Not implemented: %s", "NewSynchronization")
	return nil
}

func (s *Server) UpdateSynchronization(synchronization *models.Synchronization) *models.Synchronization {
	log.Printf("Not implemented: %s", "UpdateSynchronization")
	return nil
}

/*
 * Quota
 */

func (s *Server) GetQuotas() []*models.Quota {
	log.Printf("Not implemented: %s", "GetQuotas")
	return nil
}

/*
 * Log
 */

func (s *Server) NewLog() *models.Log {
	log.Printf("Not implemented: %s", "NewLog")
	return nil
}

func (s *Server) UpdateLog(logs *models.Log) *models.Log {
	log.Printf("Not implemented: %s", "UpdateLog")
	return nil
}

func (s *Server) GetLogByID(id int64) *models.Log {
	log.Printf("Not implemented: %s", "GetLogByID")
	return nil
}

/*
 * Archive
 */

func (s *Server) NewArchive() *models.Archive {
	log.Printf("Not implemented: %s", "NewArchive")
	return nil
}

func (s *Server) UpdateArchive(archive *models.Archive) *models.Archive {
	log.Printf("Not implemented: %s", "UpdateArchive")
	return nil
}

func (s *Server) GetArchives() []*models.Archive {
	log.Printf("Not implemented: %s", "GetArchives")
	return nil
}

func (s *Server) GetArchiveByID(id int64) *models.Archive {
	log.Printf("Not implemented: %s", "GetArchiveByID")
	return nil
}
