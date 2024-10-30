package srvr

import (
	"log"

	"github.com/gyk4j/toc/toc-backend/models"
)

var transfers = make([]*models.Transfer, 0)

func (s *Server) NewTransfer(backup *models.Backup) *models.Transfer {
	id := int64(len(transfers))

	b := backup // Trust client data?
	//b := GetBackupByID((*backup).ID) // More robust method

	// Backup must exists before it can be transferred.
	// If an invalid ID is provided, then transfer is skipped.
	if b != nil {
		log.Printf("Backup found: %d = %s\n", (*b).ID, (*b).Time)
		r := models.Transfer{
			ID:     id,
			Backup: b,
			Status: models.TransferStatusQueued,
		}

		transfers = append(transfers, &r)
		return &r
	} else {
		log.Println("[Error] Backup not found")
		return nil
	}
}

func (s *Server) UpdateTransfer(transfer *models.Transfer) *models.Transfer {
	return transfer
}

func (s *Server) GetTransfers() []*models.Transfer {
	return transfers
}

func (s *Server) GetTransferByID(id int64) *models.Transfer {
	if id >= 0 && id < int64(len(transfers)) {
		return transfers[id]
	} else {
		return nil
	}
}
