package services

import (
	"log"

	"github.com/gyk4j/toc/toc-agent/models"
)

var transfers = make([]*models.Transfer, 0)

func NewTransfer(backup *models.Backup) *models.Transfer {
	id := int64(len(transfers))

	b := backup // Trust client data?
	//b := GetBackupByID((*backup).ID) // More robust method

	// Backup must exists before it can be restored.
	// If an invalid ID is provided, then restoration is skipped.
	if b != nil {
		log.Printf("Backup found: %d = %s\n", (*b).ID, (*b).Time)
		t := models.Transfer{
			ID:     id,
			Backup: b,
			Status: models.TransferStatusQueued,
		}

		transfers = append(transfers, &t)
		return &t
	} else {
		log.Println("[Error] Backup not found")
		return nil
	}
}

func GetTransfers() []*models.Transfer {
	return transfers
}

func GetTransferByID(id int64) *models.Transfer {
	if id >= 0 && id < int64(len(restorations)) {
		return transfers[id]
	} else {
		return nil
	}
}
