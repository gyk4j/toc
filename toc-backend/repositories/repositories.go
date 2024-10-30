package repositories

import (
	"github.com/gyk4j/toc/toc-backend/models"
	"github.com/gyk4j/toc/toc-backend/repositories/inmemory"
)

type repositories struct {
	repositories []Repository
}

func (r *repositories) NewBackup(backup *models.Backup) bool {
	b := true

	for _, repo := range r.repositories {
		b = b && repo.NewBackup(backup)
	}

	return b
}

func (r *repositories) UpdateBackup(backup *models.Backup) bool {
	b := true

	for _, repo := range r.repositories {
		b = b && repo.UpdateBackup(backup)
	}

	return b
}

func (r *repositories) GetBackups() []*models.Backup {
	m := make([]*models.Backup, 0)

	for _, repo := range r.repositories {
		backups := repo.GetBackups()
		m = append(m, backups...)
	}

	return m
}

func (r *repositories) GetBackupByID(id int64) *models.Backup {
	for _, repo := range r.repositories {
		backup := repo.GetBackupByID(id)

		if backup != nil {
			return backup
		}
	}

	return nil
}

func (r *repositories) NewRestoration(restoration *models.Restoration) bool {
	b := true

	for _, repo := range r.repositories {
		b = b && repo.NewRestoration(restoration)
	}

	return b
}

func (r *repositories) UpdateRestoration(restoration *models.Restoration) bool {
	b := true

	for _, repo := range r.repositories {
		b = b && repo.UpdateRestoration(restoration)
	}

	return b
}

func (r *repositories) GetRestorations() []*models.Restoration {
	m := make([]*models.Restoration, 0)

	for _, repo := range r.repositories {
		restorations := repo.GetRestorations()
		m = append(m, restorations...)
	}

	return m
}

func (r *repositories) GetRestorationByID(id int64) *models.Restoration {
	for _, repo := range r.repositories {
		restoration := repo.GetRestorationByID(id)

		if restoration != nil {
			return restoration
		}
	}

	return nil
}

func (r *repositories) NewTransfer(transfer *models.Transfer) bool {
	b := true

	for _, repo := range r.repositories {
		b = b && repo.NewTransfer(transfer)
	}

	return b
}

func (r *repositories) UpdateTransfer(transfer *models.Transfer) bool {
	b := true

	for _, repo := range r.repositories {
		b = b && repo.UpdateTransfer(transfer)
	}

	return b
}

func (r *repositories) GetTransfers() []*models.Transfer {
	m := make([]*models.Transfer, 0)

	for _, repo := range r.repositories {
		transfers := repo.GetTransfers()
		m = append(m, transfers...)
	}

	return m
}

func (r *repositories) GetTransferByID(id int64) *models.Transfer {
	for _, repo := range r.repositories {
		transfer := repo.GetTransferByID(id)

		if transfer != nil {
			return transfer
		}
	}

	return nil
}

func (r *repositories) NewSynchronization(synchronization *models.Synchronization) bool {
	b := true

	for _, repo := range r.repositories {
		b = b && repo.NewSynchronization(synchronization)
	}

	return b
}

func (r *repositories) UpdateSynchronization(synchronization *models.Synchronization) bool {
	b := true

	for _, repo := range r.repositories {
		b = b && repo.UpdateSynchronization(synchronization)
	}

	return b
}

func (r *repositories) GetQuotas() []*models.Quota {
	m := make([]*models.Quota, 0)

	for _, repo := range r.repositories {
		quotas := repo.GetQuotas()
		m = append(m, quotas...)
	}

	return m
}

func (r *repositories) NewLog(log *models.Log) bool {
	b := true

	for _, repo := range r.repositories {
		b = b && repo.NewLog(log)
	}

	return b
}

func (r *repositories) UpdateLog(log *models.Log) bool {
	b := true

	for _, repo := range r.repositories {
		b = b && repo.UpdateLog(log)
	}

	return b
}

func (r *repositories) GetLogByID(id int64) *models.Log {
	for _, repo := range r.repositories {
		log := repo.GetLogByID(id)

		if log != nil {
			return log
		}
	}

	return nil
}

func (r *repositories) NewArchive(archive *models.Archive) bool {
	b := true

	for _, repo := range r.repositories {
		b = b && repo.NewArchive(archive)
	}

	return b
}

func (r *repositories) UpdateArchive(archive *models.Archive) bool {
	b := true

	for _, repo := range r.repositories {
		b = b && repo.UpdateArchive(archive)
	}

	return b
}

func (r *repositories) GetArchives() []*models.Archive {
	m := make([]*models.Archive, 0)

	for _, repo := range r.repositories {
		archives := repo.GetArchives()
		m = append(m, archives...)
	}

	return m
}

func (r *repositories) GetArchiveByID(id int64) *models.Archive {
	for _, repo := range r.repositories {
		archive := repo.GetArchiveByID(id)

		if archive != nil {
			return archive
		}
	}

	return nil
}

var i = inmemory.InMemory{}
var Repositories = repositories{
	repositories: []Repository{
		&i,
	},
}
