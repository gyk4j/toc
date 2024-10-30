package inmemory

import "github.com/gyk4j/toc/toc-backend/models"

func (r *InMemory) NewRestoration(restoration *models.Restoration) bool {
	return false
}

func (r *InMemory) UpdateRestoration(restoration *models.Restoration) bool {
	return false
}

func (r *InMemory) GetRestorations() []*models.Restoration {
	return nil
}

func (r *InMemory) GetRestorationByID(id int64) *models.Restoration {
	return nil
}
