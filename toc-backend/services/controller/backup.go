package controller

import (
	"fmt"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gyk4j/toc/toc-backend/models"
	"github.com/gyk4j/toc/toc-backend/repositories"
)

func (c *Controller) NewBackup() *models.Backup {
	id := time.Now()

	b := models.Backup{
		ID:        id.Unix(),
		Time:      strfmt.DateTime(id),
		Snapshots: make([]*models.Snapshot, 0),
	}

	servers := []string{"web", "file", "db", "mail"}

	for _, s := range servers {
		ss := models.Snapshot{
			ID:       id.Unix(),
			File:     fmt.Sprintf("/toc/archives/%s-%d.tar.gz", s, id.Unix()),
			Status:   models.SnapshotStatusQueued,
			Complete: false,
		}
		b.Snapshots = append(b.Snapshots, &ss)
	}

	if repositories.Repositories.NewBackup(&b) {
		return &b
	} else {
		return nil
	}
}

func (c *Controller) UpdateBackup(backup *models.Backup) *models.Backup {
	if repositories.Repositories.UpdateBackup(backup) {
		return backup
	} else {
		return nil
	}
}

func (c *Controller) GetBackups() []*models.Backup {
	return repositories.Repositories.GetBackups()
}

func (c *Controller) GetBackupByID(id int64) *models.Backup {
	return repositories.Repositories.GetBackupByID(id)
}
