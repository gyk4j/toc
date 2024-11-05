package inmemory

import "github.com/gyk4j/toc/toc-backend/models"

var backups = make([]*models.Backup, 0)

func (r *InMemory) NewBackup(backup *models.Backup) bool {
	backups = append(backups, backup)
	return false
}

func (r *InMemory) UpdateBackup(backup *models.Backup) bool {
	return false
}

func (r *InMemory) GetBackups() []*models.Backup {
	return nil
}

func (r *InMemory) GetBackupByID(id int64) *models.Backup {
	return nil
}
