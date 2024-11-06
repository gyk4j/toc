package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/quota"
)

func GetQuotas(params quota.GetQuotasParams) middleware.Responder {
	var res middleware.Responder

	s := router.Route(params.HTTPRequest)

	qs := s.GetQuotas()
	if qs != nil {
		res = quota.NewGetQuotasOK().WithPayload(qs)
	} else {
		res = quota.NewGetQuotasServiceUnavailable()
	}

	return res
}
