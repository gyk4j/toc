// Code generated by go-swagger; DO NOT EDIT.

package backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetBackupHandlerFunc turns a function with the right signature into a get backup handler
type GetBackupHandlerFunc func(GetBackupParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBackupHandlerFunc) Handle(params GetBackupParams) middleware.Responder {
	return fn(params)
}

// GetBackupHandler interface for that can handle valid get backup params
type GetBackupHandler interface {
	Handle(GetBackupParams) middleware.Responder
}

// NewGetBackup creates a new http.Handler for the get backup operation
func NewGetBackup(ctx *middleware.Context, handler GetBackupHandler) *GetBackup {
	return &GetBackup{Context: ctx, Handler: handler}
}

/*
	GetBackup swagger:route GET /backups backup getBackup

# Query backup info

Retrieve information and state of all backups and their snapshots
*/
type GetBackup struct {
	Context *middleware.Context
	Handler GetBackupHandler
}

func (o *GetBackup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetBackupParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
