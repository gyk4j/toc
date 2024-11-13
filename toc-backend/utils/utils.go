package utils

import (
	"github.com/gyk4j/toc/toc-backend/models"
)

func UpdateBackupSnapshots(old *models.Snapshots, new *models.Snapshots) int {
	updated := 0
	if new.Db != nil {
		old.Db = new.Db
		updated++
	}

	if new.File != nil {
		old.File = new.File
		updated++
	}

	if new.Mail != nil {
		old.Mail = new.Mail
		updated++
	}

	if new.Web != nil {
		old.Web = new.Web
		updated++
	}

	return updated
}

type BackupStatus int

const (
	BackupStatusUnknown       BackupStatus = iota
	BackupStatusQueued                     // Waiting to create backup
	BackupStatusCreating                   // Creating backup
	BackupStatusUncreated                  // Failed to create backup
	BackupStatusCreated                    // Backup created successfully
	BackupStatusTransferring               // Transferring backup between main and stepup
	BackupStatusUntransferred              // Failed to transfer backup (e.g. network disconnection, storage limit)
	BackupStatusTransferred                // Backup transferred successfully
	BackupStatusRestoring                  // Restoring backup
	BackupStatusUnrestored                 // Failed to restore backup (e.g. corrupt backup, missing backup files, server not ready/running)
	BackupStatusRestored                   // Backup restored successfully
)

func (s BackupStatus) String() string {
	return [...]string{
		"unknown",
		"queued",
		"creating", "uncreated", "created",
		"transferring", "untransferred", "transferred",
		"restoring", "unrestored", "restored",
	}[s]
}

func (s BackupStatus) EnumIndex() int {
	return int(s)
}

var backupStatusMap = map[string]BackupStatus{
	"":              BackupStatusUnknown,
	"queued":        BackupStatusQueued,
	"creating":      BackupStatusCreating,
	"uncreated":     BackupStatusUncreated,
	"created":       BackupStatusCreated,
	"transferring":  BackupStatusTransferring,
	"untransferred": BackupStatusUntransferred,
	"transferred":   BackupStatusTransferred,
	"restoring":     BackupStatusRestoring,
	"unrestored":    BackupStatusUnrestored,
	"restored":      BackupStatusRestored,
}

func GetBackupStatus(b *models.Backup) BackupStatus {
	status := BackupStatusRestored

	var last BackupStatus

	last = GetBackupSnapshotStatus(b.Snapshots.Db)
	if last < status {
		status = last
	}

	last = GetBackupSnapshotStatus(b.Snapshots.File)
	if last < status {
		status = last
	}

	last = GetBackupSnapshotStatus(b.Snapshots.Mail)
	if last < status {
		status = last
	}

	last = GetBackupSnapshotStatus(b.Snapshots.Web)
	if last < status {
		status = last
	}

	return status
}

func GetBackupSnapshotStatus(s *models.Snapshot) BackupStatus {
	return backupStatusMap[s.Status]
}
