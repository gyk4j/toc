package inmemory

import "github.com/gyk4j/toc/toc-backend/models"

func (r *InMemory) NewSynchronization(synchronization *models.Synchronization) bool {
	return false
}

func (r *InMemory) UpdateSynchronization(synchronization *models.Synchronization) bool {
	return false
}

func (r *InMemory) GetSynchronizations() []*models.Synchronization {
	return nil
}

func (r *InMemory) GetSynchronizationByID(id int64) *models.Synchronization {
	return nil
}
