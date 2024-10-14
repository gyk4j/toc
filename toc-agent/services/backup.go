package services

import (
	"fmt"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gyk4j/toc/toc-agent/models"
)

var backups = make([]*models.Backup, 0)

func NewBackup() *models.Backup {
	id := int64(len(backups))
	b := models.Backup{
		ID:        id,
		Time:      strfmt.DateTime(time.Now()),
		Snapshots: make([]*models.Snapshot, 0),
	}

	web := models.Snapshot{
		ID:       id,
		File:     fmt.Sprintf("/toc/archives/web-%d.tar.gz", id),
		Status:   models.SnapshotStatusQueued,
		Complete: false,
	}

	file := models.Snapshot{
		ID:       id,
		File:     fmt.Sprintf("/toc/archives/file-%d.tar.gz", id),
		Status:   models.SnapshotStatusQueued,
		Complete: false,
	}

	db := models.Snapshot{
		ID:       id,
		File:     fmt.Sprintf("/toc/archives/db-%d.tar.gz", id),
		Status:   models.SnapshotStatusQueued,
		Complete: false,
	}

	mail := models.Snapshot{
		ID:       id,
		File:     fmt.Sprintf("/toc/archives/mail-%d.tar.gz", id),
		Status:   models.SnapshotStatusQueued,
		Complete: false,
	}

	b.Snapshots = append(b.Snapshots, &web)
	b.Snapshots = append(b.Snapshots, &file)
	b.Snapshots = append(b.Snapshots, &db)
	b.Snapshots = append(b.Snapshots, &mail)

	backups = append(backups, &b)
	return &b
}

func GetBackups() []*models.Backup {
	return backups
}

func GetBackupByID(id int64) *models.Backup {
	if id >= 0 && id < int64(len(backups)) {
		return backups[id]
	} else {
		return nil
	}
}
