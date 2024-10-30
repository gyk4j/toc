package srvr

import (
	"strconv"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gyk4j/toc/toc-backend/models"
)

var logs = make([]*models.Log, 0)

func (s *Server) NewLog() *models.Log {
	id := int64(len(logs))
	l := models.Log{
		ID:     id,
		Status: models.LogStatusQueued,
		Time:   strfmt.DateTime(time.Now()),
		URL:    "/v1/logs/" + strconv.FormatInt(id, 10),
	}
	logs = append(logs, &l)
	return &l
}

func (s *Server) UpdateLog(log *models.Log) *models.Log {
	return log
}

func (s *Server) GetLogByID(id int64) *models.Log {
	if id >= 0 && id < int64(len(logs)) {
		return logs[id]
	} else {
		return nil
	}
}
