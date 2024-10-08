// Code generated by go-swagger; DO NOT EDIT.

package synchronization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// NewSynchronizationHandlerFunc turns a function with the right signature into a new synchronization handler
type NewSynchronizationHandlerFunc func(NewSynchronizationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NewSynchronizationHandlerFunc) Handle(params NewSynchronizationParams) middleware.Responder {
	return fn(params)
}

// NewSynchronizationHandler interface for that can handle valid new synchronization params
type NewSynchronizationHandler interface {
	Handle(NewSynchronizationParams) middleware.Responder
}

// NewNewSynchronization creates a new http.Handler for the new synchronization operation
func NewNewSynchronization(ctx *middleware.Context, handler NewSynchronizationHandler) *NewSynchronization {
	return &NewSynchronization{Context: ctx, Handler: handler}
}

/*
	NewSynchronization swagger:route POST /synchronizations synchronization newSynchronization

# Request backup snapshot synchronization between main and stepup controller

Initiates backup snapshot state synchronization due to network disconnection
*/
type NewSynchronization struct {
	Context *middleware.Context
	Handler NewSynchronizationHandler
}

func (o *NewSynchronization) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewNewSynchronizationParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
