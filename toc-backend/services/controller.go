package services

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gyk4j/toc/toc-backend/models"
	"github.com/gyk4j/toc/toc-backend/repositories"
)

type Controller struct {
	backups          []*models.Backup
	restorations     []*models.Restoration
	transfers        []*models.Transfer
	synchronizations []*models.Synchronization
	logs             []*models.Log
	archives         []*models.Archive
}

/*
 * Backup
 */

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
		c.backups = append(c.backups, &b)
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

/*
 * Restoration
 */

func (c *Controller) NewRestoration(backup *models.Backup) *models.Restoration {
	id := int64(len(c.restorations))

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

		c.restorations = append(c.restorations, &r)
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
	return c.restorations
}

func (c *Controller) GetRestorationByID(id int64) *models.Restoration {
	if id >= 0 && id < int64(len(c.restorations)) {
		return c.restorations[id]
	} else {
		return nil
	}
}

/*
 * Transfer
 */

func (c *Controller) NewTransfer(backup *models.Backup) *models.Transfer {
	id := int64(len(c.transfers))

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

		c.transfers = append(c.transfers, &r)
		return &r
	} else {
		log.Println("[Error] Backup not found")
		return nil
	}
}

func (c *Controller) UpdateTransfer(transfer *models.Transfer) *models.Transfer {
	return transfer
}

func (c *Controller) GetTransfers() []*models.Transfer {
	return c.transfers
}

func (c *Controller) GetTransferByID(id int64) *models.Transfer {
	if id >= 0 && id < int64(len(c.transfers)) {
		return c.transfers[id]
	} else {
		return nil
	}
}

/*
 * Synchronization
 */

func (c *Controller) NewSynchronization() *models.Synchronization {
	id := int64(len(c.synchronizations))
	s := models.Synchronization{
		ID:     id,
		Status: models.LogStatusQueued,
		Time:   strfmt.DateTime(time.Now()),
	}
	c.synchronizations = append(c.synchronizations, &s)
	return &s
}

func (c *Controller) UpdateSynchronization(synchronization *models.Synchronization) *models.Synchronization {
	return synchronization
}

/*
 * Quota
 */

func (c *Controller) GetQuotas() []*models.Quota {
	qs := make([]*models.Quota, 0)

	var id int64

	id = int64(len(qs))
	q1 := models.Quota{
		ID:   id,
		Path: "/home/a/q1",
		Soft: rand.Int63(),
		Hard: rand.Int63(),
	}

	qs = append(qs, &q1)

	id = int64(len(qs))
	q2 := models.Quota{
		ID:   id,
		Path: "/home/b/q2",
		Soft: rand.Int63(),
		Hard: rand.Int63(),
	}

	qs = append(qs, &q2)

	return qs
}

/*
 * Log
 */

func (c *Controller) NewLog() *models.Log {
	id := int64(len(c.logs))
	l := models.Log{
		ID:     id,
		Status: models.LogStatusQueued,
		Time:   strfmt.DateTime(time.Now()),
		URL:    "/v1/logs/" + strconv.FormatInt(id, 10),
	}
	c.logs = append(c.logs, &l)
	return &l
}

func (c *Controller) UpdateLog(log *models.Log) *models.Log {
	return log
}

func (c *Controller) GetLogByID(id int64) *models.Log {
	if id >= 0 && id < int64(len(c.logs)) {
		return c.logs[id]
	} else {
		return nil
	}
}

/*
 * Archive
 */

func (c *Controller) NewArchive() *models.Archive {
	id := int64(len(c.archives))
	a := models.Archive{
		ID:     id,
		Status: models.ArchiveStatusQueued,
		Time:   strfmt.DateTime(time.Now()),
		URL:    "/v1/archives/" + strconv.FormatInt(id, 10),
	}
	c.archives = append(c.archives, &a)
	return &a
}

func (c *Controller) UpdateArchive(archive *models.Archive) *models.Archive {
	return archive
}

func (c *Controller) GetArchives() []*models.Archive {
	return c.archives
}

func (c *Controller) GetArchiveByID(id int64) *models.Archive {
	if id >= 0 && id < int64(len(c.archives)) {
		return c.archives[id]
	} else {
		return nil
	}
}
