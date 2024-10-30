package inmemory

import "github.com/gyk4j/toc/toc-backend/models"

func (r *InMemory) NewArchive(archive *models.Archive) bool {
	return false
}

func (r *InMemory) UpdateArchive(archive *models.Archive) bool {
	return false
}

func (r *InMemory) GetArchives() []*models.Archive {
	return nil
}

func (r *InMemory) GetArchiveByID(id int64) *models.Archive {
	return nil
}
