// Code generated by go-swagger; DO NOT EDIT.

package restoration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gyk4j/toc/toc-backend/models"
)

// GetRestorationsOKCode is the HTTP code returned for type GetRestorationsOK
const GetRestorationsOKCode int = 200

/*
GetRestorationsOK Restoration info retrieved

swagger:response getRestorationsOK
*/
type GetRestorationsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Restoration `json:"body,omitempty"`
}

// NewGetRestorationsOK creates GetRestorationsOK with default headers values
func NewGetRestorationsOK() *GetRestorationsOK {

	return &GetRestorationsOK{}
}

// WithPayload adds the payload to the get restorations o k response
func (o *GetRestorationsOK) WithPayload(payload []*models.Restoration) *GetRestorationsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get restorations o k response
func (o *GetRestorationsOK) SetPayload(payload []*models.Restoration) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRestorationsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Restoration, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetRestorationsForbiddenCode is the HTTP code returned for type GetRestorationsForbidden
const GetRestorationsForbiddenCode int = 403

/*
GetRestorationsForbidden Forbidden from querying restoration info

swagger:response getRestorationsForbidden
*/
type GetRestorationsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetRestorationsForbidden creates GetRestorationsForbidden with default headers values
func NewGetRestorationsForbidden() *GetRestorationsForbidden {

	return &GetRestorationsForbidden{}
}

// WithPayload adds the payload to the get restorations forbidden response
func (o *GetRestorationsForbidden) WithPayload(payload *models.APIResponse) *GetRestorationsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get restorations forbidden response
func (o *GetRestorationsForbidden) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRestorationsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRestorationsNotFoundCode is the HTTP code returned for type GetRestorationsNotFound
const GetRestorationsNotFoundCode int = 404

/*
GetRestorationsNotFound Restoration not found

swagger:response getRestorationsNotFound
*/
type GetRestorationsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetRestorationsNotFound creates GetRestorationsNotFound with default headers values
func NewGetRestorationsNotFound() *GetRestorationsNotFound {

	return &GetRestorationsNotFound{}
}

// WithPayload adds the payload to the get restorations not found response
func (o *GetRestorationsNotFound) WithPayload(payload *models.APIResponse) *GetRestorationsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get restorations not found response
func (o *GetRestorationsNotFound) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRestorationsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRestorationsMethodNotAllowedCode is the HTTP code returned for type GetRestorationsMethodNotAllowed
const GetRestorationsMethodNotAllowedCode int = 405

/*
GetRestorationsMethodNotAllowed Bad request

swagger:response getRestorationsMethodNotAllowed
*/
type GetRestorationsMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetRestorationsMethodNotAllowed creates GetRestorationsMethodNotAllowed with default headers values
func NewGetRestorationsMethodNotAllowed() *GetRestorationsMethodNotAllowed {

	return &GetRestorationsMethodNotAllowed{}
}

// WithPayload adds the payload to the get restorations method not allowed response
func (o *GetRestorationsMethodNotAllowed) WithPayload(payload *models.APIResponse) *GetRestorationsMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get restorations method not allowed response
func (o *GetRestorationsMethodNotAllowed) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRestorationsMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRestorationsInternalServerErrorCode is the HTTP code returned for type GetRestorationsInternalServerError
const GetRestorationsInternalServerErrorCode int = 500

/*
GetRestorationsInternalServerError TOC controller error

swagger:response getRestorationsInternalServerError
*/
type GetRestorationsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetRestorationsInternalServerError creates GetRestorationsInternalServerError with default headers values
func NewGetRestorationsInternalServerError() *GetRestorationsInternalServerError {

	return &GetRestorationsInternalServerError{}
}

// WithPayload adds the payload to the get restorations internal server error response
func (o *GetRestorationsInternalServerError) WithPayload(payload *models.APIResponse) *GetRestorationsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get restorations internal server error response
func (o *GetRestorationsInternalServerError) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRestorationsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRestorationsServiceUnavailableCode is the HTTP code returned for type GetRestorationsServiceUnavailable
const GetRestorationsServiceUnavailableCode int = 503

/*
GetRestorationsServiceUnavailable Service unavailable

swagger:response getRestorationsServiceUnavailable
*/
type GetRestorationsServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetRestorationsServiceUnavailable creates GetRestorationsServiceUnavailable with default headers values
func NewGetRestorationsServiceUnavailable() *GetRestorationsServiceUnavailable {

	return &GetRestorationsServiceUnavailable{}
}

// WithPayload adds the payload to the get restorations service unavailable response
func (o *GetRestorationsServiceUnavailable) WithPayload(payload *models.APIResponse) *GetRestorationsServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get restorations service unavailable response
func (o *GetRestorationsServiceUnavailable) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRestorationsServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
