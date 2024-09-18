package services

import (
	"strconv"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gyk4j/toc/toc/server/models"
)

var logs = make([]*models.Log, 0)

func NewLog() *models.Log {
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

func GetLogByID(id int64) *models.Log {
	if id >= 0 && id < int64(len(logs)) {
		return logs[id]
	} else {
		return nil
	}
}
