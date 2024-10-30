package server

import (
	logops "log"

	"github.com/gyk4j/toc/toc-backend/models"
)

func (s *Server) NewLog() *models.Log {
	logops.Printf("Not implemented: %s", "NewLog")
	return nil
}

func (s *Server) UpdateLog(log *models.Log) *models.Log {
	logops.Printf("Not implemented: %s", "UpdateLog")
	return nil
}

func (s *Server) GetLogByID(id int64) *models.Log {
	logops.Printf("Not implemented: %s", "GetLogByID")
	return nil
}
