package services

import (
	"github.com/gyk4j/toc/toc/server/models"
)

var restorations = make([]*models.Restoration, 0)

func NewRestoration() *models.Restoration {
	id := int64(len(restorations))

	b := GetBackupByID(id)

	// Backup must exists before it can be restored.
	// If an invalid ID is provided, then restoration is skipped.
	if b != nil {
		r := models.Restoration{
			ID:     id,
			Backup: b,
			Status: models.RestorationStatusQueued,
		}

		restorations = append(restorations, &r)
		return &r
	} else {
		return nil
	}
}

func GetRestorations() []*models.Restoration {
	return restorations
}

func GetRestorationByID(id int64) *models.Restoration {
	if id >= 0 && id < int64(len(restorations)) {
		return restorations[id]
	} else {
		return nil
	}
}
