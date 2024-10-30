package server

import (
	"log"

	"github.com/gyk4j/toc/toc-backend/models"
)

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
