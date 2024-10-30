package inmemory

import "github.com/gyk4j/toc/toc-backend/models"

func (r *InMemory) NewLog(log *models.Log) bool {
	return false
}

func (r *InMemory) UpdateLog(log *models.Log) bool {
	return false
}

func (r *InMemory) GetLogs() []*models.Log {
	return nil
}

func (r *InMemory) GetLogByID(id int64) *models.Log {
	return nil
}
