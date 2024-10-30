package ctrl

import (
	"log"

	"github.com/gyk4j/toc/toc-backend/models"
)

var restorations = make([]*models.Restoration, 0)

func (c *Controller) NewRestoration(backup *models.Backup) *models.Restoration {
	id := int64(len(restorations))

	b := backup // Trust client data?
	//b := GetBackupByID((*backup).ID) // More robust method

	// Backup must exists before it can be restored.
	// If an invalid ID is provided, then restoration is skipped.
	if b != nil {
		log.Printf("Backup found: %d = %s\n", (*b).ID, (*b).Time)
		r := models.Restoration{
			ID:     id,
			Backup: b,
			Status: models.RestorationStatusQueued,
		}

		restorations = append(restorations, &r)
		return &r
	} else {
		log.Println("[Error] Backup not found")
		return nil
	}
}

func (c *Controller) UpdateRestoration(restoration *models.Restoration) *models.Restoration {
	return restoration
}

func (c *Controller) GetRestorations() []*models.Restoration {
	return restorations
}

func (c *Controller) GetRestorationByID(id int64) *models.Restoration {
	if id >= 0 && id < int64(len(restorations)) {
		return restorations[id]
	} else {
		return nil
	}
}
