package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gyk4j/toc/toc/server/restapi/operations/quota"
	"github.com/gyk4j/toc/toc/server/services"
)

func GetQuotas(params quota.GetQuotaParams) middleware.Responder {
	var res middleware.Responder

	qs := services.GetQuotas()
	if qs != nil {
		res = quota.NewGetQuotaOK().WithPayload(qs)
	} else {
		res = quota.NewGetQuotaServiceUnavailable()
	}

	return res
}
