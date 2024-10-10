// Code generated by go-swagger; DO NOT EDIT.

package archive

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetArchiveHandlerFunc turns a function with the right signature into a get archive handler
type GetArchiveHandlerFunc func(GetArchiveParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetArchiveHandlerFunc) Handle(params GetArchiveParams) middleware.Responder {
	return fn(params)
}

// GetArchiveHandler interface for that can handle valid get archive params
type GetArchiveHandler interface {
	Handle(GetArchiveParams) middleware.Responder
}

// NewGetArchive creates a new http.Handler for the get archive operation
func NewGetArchive(ctx *middleware.Context, handler GetArchiveHandler) *GetArchive {
	return &GetArchive{Context: ctx, Handler: handler}
}

/*
	GetArchive swagger:route GET /archives archive getArchive

# Query archive info

Retrieve progress, information and state of all archives
*/
type GetArchive struct {
	Context *middleware.Context
	Handler GetArchiveHandler
}

func (o *GetArchive) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetArchiveParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}