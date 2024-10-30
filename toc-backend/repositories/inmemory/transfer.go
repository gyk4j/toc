package inmemory

import "github.com/gyk4j/toc/toc-backend/models"

func (r *InMemory) NewTransfer(transfer *models.Transfer) bool {
	return false
}

func (r *InMemory) UpdateTransfer(transfer *models.Transfer) bool {
	return false
}

func (r *InMemory) GetTransfers() []*models.Transfer {
	return nil
}

func (r *InMemory) GetTransferByID(id int64) *models.Transfer {
	return nil
}
