package server

import (
	"log"

	"github.com/gyk4j/toc/toc-backend/models"
)

func (s *Server) NewBackup() *models.Backup {
	log.Printf("Not implemented: %s", "NewBackup")
	return nil
}

func (s *Server) UpdateBackup(backup *models.Backup) *models.Backup {
	log.Printf("Not implemented: %s", "UpdateBackup")
	return nil
}

func (s *Server) GetBackups() []*models.Backup {
	log.Printf("Not implemented: %s", "GetBackups")
	return nil
}

func (s *Server) GetBackupByID(id int64) *models.Backup {
	log.Printf("Not implemented: %s", "GetBackupByID")
	return nil
}
