package srvr

import (
	"math/rand"

	"github.com/gyk4j/toc/toc-backend/models"
)

func (s *Server) GetQuotas() []*models.Quota {
	qs := make([]*models.Quota, 0)

	var id int64

	id = int64(len(qs))
	q1 := models.Quota{
		ID:   id,
		Path: "/home/a/q1",
		Soft: rand.Int63(),
		Hard: rand.Int63(),
	}

	qs = append(qs, &q1)

	id = int64(len(qs))
	q2 := models.Quota{
		ID:   id,
		Path: "/home/b/q2",
		Soft: rand.Int63(),
		Hard: rand.Int63(),
	}

	qs = append(qs, &q2)

	return qs
}
