package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gyk4j/toc/toc-backend/restapi/operations/quota"
	"github.com/gyk4j/toc/toc-backend/services"
)

func GetQuotas(params quota.GetQuotasParams) middleware.Responder {
	var res middleware.Responder

	qs := services.GetQuotas()
	if qs != nil {
		res = quota.NewGetQuotasOK().WithPayload(qs)
	} else {
		res = quota.NewGetQuotasServiceUnavailable()
	}

	return res
}
