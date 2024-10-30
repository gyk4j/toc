package ctrl

import (
	"strconv"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gyk4j/toc/toc-backend/models"
)

var archives = make([]*models.Archive, 0)

func (c *Controller) NewArchive() *models.Archive {
	id := int64(len(archives))
	a := models.Archive{
		ID:     id,
		Status: models.ArchiveStatusQueued,
		Time:   strfmt.DateTime(time.Now()),
		URL:    "/v1/archives/" + strconv.FormatInt(id, 10),
	}
	archives = append(archives, &a)
	return &a
}

func (c *Controller) UpdateArchive(archive *models.Archive) *models.Archive {
	return archive
}

func (c *Controller) GetArchives() []*models.Archive {
	return archives
}

func (c *Controller) GetArchiveByID(id int64) *models.Archive {
	if id >= 0 && id < int64(len(archives)) {
		return archives[id]
	} else {
		return nil
	}
}
