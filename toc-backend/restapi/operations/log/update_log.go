// Code generated by go-swagger; DO NOT EDIT.

package log

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateLogHandlerFunc turns a function with the right signature into a update log handler
type UpdateLogHandlerFunc func(UpdateLogParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateLogHandlerFunc) Handle(params UpdateLogParams) middleware.Responder {
	return fn(params)
}

// UpdateLogHandler interface for that can handle valid update log params
type UpdateLogHandler interface {
	Handle(UpdateLogParams) middleware.Responder
}

// NewUpdateLog creates a new http.Handler for the update log operation
func NewUpdateLog(ctx *middleware.Context, handler UpdateLogHandler) *UpdateLog {
	return &UpdateLog{Context: ctx, Handler: handler}
}

/*
	UpdateLog swagger:route PUT /logs log updateLog

Update log status/info

Update log status/info
*/
type UpdateLog struct {
	Context *middleware.Context
	Handler UpdateLogHandler
}

func (o *UpdateLog) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateLogParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
