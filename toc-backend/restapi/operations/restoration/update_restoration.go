// Code generated by go-swagger; DO NOT EDIT.

package restoration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateRestorationHandlerFunc turns a function with the right signature into a update restoration handler
type UpdateRestorationHandlerFunc func(UpdateRestorationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateRestorationHandlerFunc) Handle(params UpdateRestorationParams) middleware.Responder {
	return fn(params)
}

// UpdateRestorationHandler interface for that can handle valid update restoration params
type UpdateRestorationHandler interface {
	Handle(UpdateRestorationParams) middleware.Responder
}

// NewUpdateRestoration creates a new http.Handler for the update restoration operation
func NewUpdateRestoration(ctx *middleware.Context, handler UpdateRestorationHandler) *UpdateRestoration {
	return &UpdateRestoration{Context: ctx, Handler: handler}
}

/*
	UpdateRestoration swagger:route PUT /restorations restoration updateRestoration

Update restoration status/info

Update restoration status/info
*/
type UpdateRestoration struct {
	Context *middleware.Context
	Handler UpdateRestorationHandler
}

func (o *UpdateRestoration) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateRestorationParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}