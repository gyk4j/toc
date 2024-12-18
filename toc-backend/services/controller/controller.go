package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gyk4j/toc/toc-backend/models"
	"github.com/gyk4j/toc/toc-backend/repositories"
	"github.com/gyk4j/toc/toc-backend/services"
	"github.com/gyk4j/toc/toc-backend/utils"
)

var servers = []string{"web", "file", "db", "mail"}

func NewController(name string) (server *services.Server) {
	return &services.Server{
		Name:            name,
		Backup:          &BackupService{},
		Restoration:     &RestorationService{},
		Transfer:        &TransferService{},
		Synchronization: &SynchronizationService{},
		Quota:           &QuotaService{},
		Log:             &LogService{},
		Archive:         &ArchiveService{},
	}
}

//
// Backup
//

type BackupService struct {
	r repositories.IBackupRepository
}

func (s *BackupService) NewBackup() (*models.Backup, *services.AppError) {
	var err *services.AppError

	b := models.Backup{
		ID:        -1,
		Time:      strfmt.DateTime(time.Now()),
		Snapshots: &models.Snapshots{},
	}

	o := s.r.Save(&b)

	if o == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("internal server error: %s", "NewBackup"),
			Message: "internal server error",
			Code:    int(services.NewBackupInternalServerError),
		}
		return nil, err
	}

	fails := 0
	for _, server := range servers {
		url := fmt.Sprintf("http://%s/v1/Backup", server)
		bodyReader := bytes.NewReader([]byte(""))
		res, errHttp := http.Post(url, "text/json", bodyReader)

		if errHttp != nil || (res.StatusCode != http.StatusOK && res.StatusCode != http.StatusAccepted) {
			err = &services.AppError{
				Error:   errHttp,
				Message: res.Status,
				Code:    res.StatusCode,
			}
			fails++
		}

		log.Printf("%s => %s\n", res.Status, url)

		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(body)
		}
	}

	// if fails == len(servers) {
	// 	o.Status = models.BackupStatusFailed
	// 	o = s.r.Save(o)
	// }

	return o, err
}

func (s *BackupService) UpdateBackup(backup *models.Backup) (*models.Backup, *services.AppError) {
	var err *services.AppError

	b := s.r.FindById(backup.ID)
	if b == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("not found: %s", "UpdateBackup"),
			Message: "not found",
			Code:    int(services.UpdateBackupNotFound),
		}
		return nil, err
	}

	// Update snapshots' status
	utils.UpdateBackupSnapshots(b.Snapshots, backup.Snapshots)

	// Update restoration status
	log.Printf("Restoration: %d, Status: %s\n", b.ID, utils.GetBackupStatus(b).String())

	// Update backup status

	o := s.r.Save(b)

	if o == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("internal server error: %s", "UpdateBackup"),
			Message: "internal server error",
			Code:    int(services.UpdateBackupInternalServerError),
		}
		return nil, err
	}

	return o, err
}

func (s *BackupService) GetBackups() ([]*models.Backup, *services.AppError) {
	var err *services.AppError

	o := s.r.FindAll()

	if o == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("internal server error: %s", "GetBackups"),
			Message: "internal server error",
			Code:    int(services.GetBackupsInternalServerError),
		}
		return nil, err
	}

	return o, err
}

func (s *BackupService) GetBackupById(id int64) (*models.Backup, *services.AppError) {
	var err *services.AppError

	o := s.r.FindById(id)

	if o == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("not found: %s", "GetBackupById"),
			Message: "not found",
			Code:    int(services.GetBackupByIdNotFound),
		}
		return nil, err
	}

	return o, err
}

//
// Restoration
//

type RestorationService struct {
	r repositories.IRestorationRepository
}

func (s *RestorationService) NewRestoration(backup *models.Backup) (*models.Restoration, *services.AppError) {
	var err *services.AppError

	b := backup // Trust client data?
	// b := GetBackupByID(backup.ID) // More robust method

	// Backup must exists before it can be restored.
	// If an invalid ID is provided, then restoration is skipped.
	if b == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("backup not found: %s", "NewRestoration"),
			Message: "backup not found",
			Code:    int(services.NewRestorationInternalServerError),
		}
		return nil, err
	}

	log.Printf("Backup found: %d = %s\n", b.ID, b.Time)
	r := models.Restoration{
		ID:     -1,
		Backup: b,
		Status: models.RestorationStatusQueued,
	}

	o := s.r.Save(&r)

	if o == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("internal server error: %s", "NewRestoration"),
			Message: "internal server error",
			Code:    int(services.NewRestorationInternalServerError),
		}
		return nil, err
	}

	fails := 0
	for _, server := range servers {
		url := fmt.Sprintf("http://%s/v1/Restoration", server)
		body, errJson := json.Marshal(b)
		if errJson != nil {
			log.Println("Error: JSON encoding failed")
			continue
		}
		bodyReader := bytes.NewReader([]byte(body))
		res, errHttp := http.Post(url, "text/json", bodyReader)

		if errHttp != nil || (res.StatusCode != http.StatusOK && res.StatusCode != http.StatusAccepted) {
			err = &services.AppError{
				Error:   errHttp,
				Message: res.Status,
				Code:    res.StatusCode,
			}
			fails++
		}

		log.Printf("%s => %s\n", res.Status, url)

		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(body)
		}
	}

	if fails == len(servers) {
		o.Status = models.RestorationStatusFailed
		o = s.r.Save(o)
	}

	return o, err
}

func (s *RestorationService) UpdateRestoration(restoration *models.Restoration) (*models.Restoration, *services.AppError) {
	var err *services.AppError

	r := s.r.FindById(restoration.ID)
	if r == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("not found: %s", "UpdateRestoration"),
			Message: "not found",
			Code:    int(services.UpdateRestorationNotFound),
		}
		return nil, err
	}

	// Update snapshots' status
	utils.UpdateBackupSnapshots(r.Backup.Snapshots, restoration.Backup.Snapshots)

	// Update restoration status
	log.Printf("Restoration: %d, Status: %s\n", r.ID, utils.GetBackupStatus(r.Backup).String())

	o := s.r.Save(r)

	if o == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("internal server error: %s", "UpdateRestoration"),
			Message: "internal server error",
			Code:    int(services.UpdateRestorationInternalServerError),
		}
		return nil, err
	}

	return o, err
}

func (s *RestorationService) GetRestorations() ([]*models.Restoration, *services.AppError) {
	var err *services.AppError

	o := s.r.FindAll()

	if o == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("internal server error: %s", "GetRestorations"),
			Message: "internal server error",
			Code:    int(services.UpdateRestorationInternalServerError),
		}
		return nil, err
	}

	return o, err
}

func (s *RestorationService) GetRestorationById(id int64) (*models.Restoration, *services.AppError) {
	var err *services.AppError

	o := s.r.FindById(id)

	if o == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("not found: %s", "GetRestorationById"),
			Message: "not found",
			Code:    int(services.UpdateRestorationNotFound),
		}
		return nil, err
	}

	return o, err
}

//
// Transfer
//

type TransferService struct {
	r repositories.ITransferRepository
}

func (s *TransferService) NewTransfer(backup *models.Backup) (*models.Transfer, *services.AppError) {
	var err *services.AppError

	b := backup // Trust client data?
	//b := GetBackupByID((*backup).ID) // More robust method

	// Backup must exists before it can be restored.
	// If an invalid ID is provided, then restoration is skipped.
	if b == nil {
		log.Println("[Error] Backup not found")
		err = &services.AppError{
			Error:   fmt.Errorf("not found: %s", "NewTransfer"),
			Message: "not found",
			Code:    int(services.NewTransferNotFound),
		}
		return nil, err
	}

	log.Printf("Backup found: %d = %s\n", b.ID, b.Time)
	t := models.Transfer{
		ID:     -1,
		Backup: b,
		Status: models.TransferStatusQueued,
	}

	o := s.r.Save(&t)

	if o == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("internal server error: %s", "NewTransfer"),
			Message: "internal server error",
			Code:    int(services.NewTransferInternalServerError),
		}
		return nil, err
	}

	fails := 0
	for _, server := range servers {
		url := fmt.Sprintf("http://%s/v1/Transfer", server)
		body, errJson := json.Marshal(b)
		if errJson != nil {
			log.Println("Error: JSON encoding failed")
			continue
		}
		bodyReader := bytes.NewReader([]byte(body))
		res, errHttp := http.Post(url, "text/json", bodyReader)

		if errHttp != nil || (res.StatusCode != http.StatusOK && res.StatusCode != http.StatusAccepted) {
			err = &services.AppError{
				Error:   errHttp,
				Message: res.Status,
				Code:    res.StatusCode,
			}
			fails++
		}

		log.Printf("%s => %s\n", res.Status, url)

		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(body)
		}
	}

	if fails == len(servers) {
		o.Status = models.TransferStatusFailed
		o = s.r.Save(o)
	}

	return o, err
}

func (s *TransferService) UpdateTransfer(transfer *models.Transfer) (*models.Transfer, *services.AppError) {
	var err *services.AppError

	t := s.r.FindById(transfer.ID)

	if t == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("not found: %s", "UpdateTransfer"),
			Message: "not found",
			Code:    int(services.UpdateTransferNotFound),
		}
		return nil, err
	}

	// Update snapshots' status
	utils.UpdateBackupSnapshots(t.Backup.Snapshots, transfer.Backup.Snapshots)

	// Update transfer status
	log.Printf("Transfer: %d, Status: %s\n", t.ID, utils.GetBackupStatus(t.Backup).String())

	o := s.r.Save(t)

	if o == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("internal server error: %s", "UpdateTransfer"),
			Message: "internal server error",
			Code:    int(services.UpdateTransferInternalServerError),
		}
		return nil, err
	}

	return o, err
}

func (s *TransferService) GetTransfers() ([]*models.Transfer, *services.AppError) {
	var err *services.AppError

	o := s.r.FindAll()

	if o == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("internal server error: %s", "GetTransfers"),
			Message: "internal server error",
			Code:    int(services.GetTransfersInternalServerError),
		}
	}

	return o, err
}

func (s *TransferService) GetTransferById(id int64) (*models.Transfer, *services.AppError) {
	var err *services.AppError

	o := s.r.FindById(id)

	if o == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("not found: %s", "GetTransferById"),
			Message: "not found",
			Code:    int(services.GetTransferByIdNotFound),
		}
	}

	return o, err
}

//
// Synchronization
//

type SynchronizationService struct {
	r repositories.ISynchronizationRepository
}

func (s *SynchronizationService) NewSynchronization() (*models.Synchronization, *services.AppError) {
	var err *services.AppError

	sync := models.Synchronization{
		ID:     -1,
		Status: models.LogStatusQueued,
		Time:   strfmt.DateTime(time.Now()),
	}

	o := s.r.Save(&sync)

	if o == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("internal server error: %s", "NewSynchronization"),
			Message: "internal server error",
			Code:    int(services.NewSynchronizationInternalServerError),
		}
	}

	return o, err
}

func (s *SynchronizationService) UpdateSynchronization(synchronization *models.Synchronization) (*models.Synchronization, *services.AppError) {
	var err *services.AppError

	if s.r.FindById(synchronization.ID) == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("not found: %s", "UpdateSynchronization"),
			Message: "not found",
			Code:    int(services.UpdateSynchronizationNotFound),
		}
		return nil, err
	}

	o := s.r.Save(synchronization)

	if o == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("internal server error: %s", "UpdateSynchronization"),
			Message: "internal server error",
			Code:    int(services.UpdateSynchronizationInternalServerError),
		}
	}

	return o, err
}

//
// Quota
//

type QuotaService struct {
	r repositories.IQuotaRepository
}

func (s *QuotaService) GetQuotas() ([]*models.Quota, *services.AppError) {
	var err *services.AppError

	o := s.r.FindAll()

	if o == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("internal server error: %s", "GetQuotas"),
			Message: "internal server error",
			Code:    int(services.GetQuotasInternalServerError),
		}
	}

	return o, err
}

//
// Log
//

type LogService struct {
	r repositories.ILogRepository
}

func (s *LogService) NewLog() (*models.Log, *services.AppError) {
	var err *services.AppError

	l := models.Log{
		ID:     -1,
		Status: models.LogStatusQueued,
		Time:   strfmt.DateTime(time.Now()),
		URL:    "/v1/logs/" + strconv.FormatInt(0, 10),
	}

	o := s.r.Save(&l)

	if o == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("internal server error: %s", "NewLog"),
			Message: "internal server error",
			Code:    int(services.NewLogInternalServerError),
		}
	}

	return o, err
}

func (s *LogService) UpdateLog(log *models.Log) (*models.Log, *services.AppError) {
	var err *services.AppError

	if s.r.FindById(log.ID) == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("not found: %s", "UpdateLog"),
			Message: "not found",
			Code:    int(services.UpdateLogNotFound),
		}

		return nil, err
	}

	o := s.r.Save(log)

	if o == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("internal server error: %s", "UpdateLog"),
			Message: "internal server error",
			Code:    int(services.UpdateLogInternalServerError),
		}
	}

	return o, err
}

func (s *LogService) GetLogById(id int64) (*models.Log, *services.AppError) {
	var err *services.AppError

	o := s.r.FindById(id)

	if o == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("not found: %s", "GetLogById"),
			Message: "not found",
			Code:    int(services.GetLogByIdNotFound),
		}
	}

	return o, err
}

//
// Archive
//

type ArchiveService struct {
	r repositories.IArchiveRepository
}

func (s *ArchiveService) NewArchive() (*models.Archive, *services.AppError) {
	var err *services.AppError

	a := models.Archive{
		ID:     -1,
		Status: models.ArchiveStatusQueued,
		Time:   strfmt.DateTime(time.Now()),
		URL:    "/v1/archives/" + strconv.FormatInt(0, 10),
	}

	o := s.r.Save(&a)

	if o == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("internal server error: %s", "NewArchive"),
			Message: "internal server error",
			Code:    int(services.NewArchiveInternalServerError),
		}
	}

	return o, err
}

func (s *ArchiveService) UpdateArchive(archive *models.Archive) (*models.Archive, *services.AppError) {
	var err *services.AppError

	if s.r.FindById(archive.ID) == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("not found: %s", "UpdateArchive"),
			Message: "not found",
			Code:    int(services.UpdateArchiveNotFound),
		}

		return nil, err
	}

	o := s.r.Save(archive)

	if o == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("internal server error: %s", "UpdateArchive"),
			Message: "internal server error",
			Code:    int(services.UpdateArchiveInternalServerError),
		}
	}

	return o, err
}

func (s *ArchiveService) GetArchives() ([]*models.Archive, *services.AppError) {
	var err *services.AppError

	o := s.r.FindAll()

	if o == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("internal server error: %s", "GetArchives"),
			Message: "internal server error",
			Code:    int(services.GetArchivesInternalServerError),
		}
	}

	return o, err
}

func (s *ArchiveService) GetArchiveById(id int64) (*models.Archive, *services.AppError) {
	var err *services.AppError

	o := s.r.FindById(id)

	if o == nil {
		err = &services.AppError{
			Error:   fmt.Errorf("not found: %s", "GetArchiveById"),
			Message: "not found",
			Code:    int(services.GetArchiveByIdNotFound),
		}
	}

	return o, err
}
