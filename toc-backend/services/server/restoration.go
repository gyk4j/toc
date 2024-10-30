package server

import (
	"log"

	"github.com/gyk4j/toc/toc-backend/models"
)

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
