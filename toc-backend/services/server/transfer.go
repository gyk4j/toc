package server

import (
	"log"

	"github.com/gyk4j/toc/toc-backend/models"
)

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
