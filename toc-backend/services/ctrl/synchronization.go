package ctrl

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gyk4j/toc/toc-backend/models"
)

var synchronizations = make([]*models.Synchronization, 0)

func (c *Controller) NewSynchronization() *models.Synchronization {
	id := int64(len(synchronizations))
	s := models.Synchronization{
		ID:     id,
		Status: models.LogStatusQueued,
		Time:   strfmt.DateTime(time.Now()),
	}
	synchronizations = append(synchronizations, &s)
	return &s
}

func (c *Controller) UpdateSynchronization(synchronization *models.Synchronization) *models.Synchronization {
	return synchronization
}
