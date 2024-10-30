package srvr

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gyk4j/toc/toc-backend/models"
)

var synchronizations = make([]*models.Synchronization, 0)

func (s *Server) NewSynchronization() *models.Synchronization {
	id := int64(len(synchronizations))
	sync := models.Synchronization{
		ID:     id,
		Status: models.LogStatusQueued,
		Time:   strfmt.DateTime(time.Now()),
	}
	synchronizations = append(synchronizations, &sync)
	return &sync
}

func (s *Server) UpdateSynchronization(synchronization *models.Synchronization) *models.Synchronization {
	return synchronization
}
