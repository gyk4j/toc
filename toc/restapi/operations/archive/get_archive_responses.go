// Code generated by go-swagger; DO NOT EDIT.

package archive

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gyk4j/toc/toc/models"
)

// GetArchiveOKCode is the HTTP code returned for type GetArchiveOK
const GetArchiveOKCode int = 200

/*
GetArchiveOK Archive info retrieved

swagger:response getArchiveOK
*/
type GetArchiveOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Archive `json:"body,omitempty"`
}

// NewGetArchiveOK creates GetArchiveOK with default headers values
func NewGetArchiveOK() *GetArchiveOK {

	return &GetArchiveOK{}
}

// WithPayload adds the payload to the get archive o k response
func (o *GetArchiveOK) WithPayload(payload []*models.Archive) *GetArchiveOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get archive o k response
func (o *GetArchiveOK) SetPayload(payload []*models.Archive) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetArchiveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Archive, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetArchiveForbiddenCode is the HTTP code returned for type GetArchiveForbidden
const GetArchiveForbiddenCode int = 403

/*
GetArchiveForbidden Forbidden from querying archive info

swagger:response getArchiveForbidden
*/
type GetArchiveForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetArchiveForbidden creates GetArchiveForbidden with default headers values
func NewGetArchiveForbidden() *GetArchiveForbidden {

	return &GetArchiveForbidden{}
}

// WithPayload adds the payload to the get archive forbidden response
func (o *GetArchiveForbidden) WithPayload(payload *models.APIResponse) *GetArchiveForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get archive forbidden response
func (o *GetArchiveForbidden) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetArchiveForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetArchiveNotFoundCode is the HTTP code returned for type GetArchiveNotFound
const GetArchiveNotFoundCode int = 404

/*
GetArchiveNotFound Archive not found

swagger:response getArchiveNotFound
*/
type GetArchiveNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetArchiveNotFound creates GetArchiveNotFound with default headers values
func NewGetArchiveNotFound() *GetArchiveNotFound {

	return &GetArchiveNotFound{}
}

// WithPayload adds the payload to the get archive not found response
func (o *GetArchiveNotFound) WithPayload(payload *models.APIResponse) *GetArchiveNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get archive not found response
func (o *GetArchiveNotFound) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetArchiveNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetArchiveMethodNotAllowedCode is the HTTP code returned for type GetArchiveMethodNotAllowed
const GetArchiveMethodNotAllowedCode int = 405

/*
GetArchiveMethodNotAllowed Bad request

swagger:response getArchiveMethodNotAllowed
*/
type GetArchiveMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetArchiveMethodNotAllowed creates GetArchiveMethodNotAllowed with default headers values
func NewGetArchiveMethodNotAllowed() *GetArchiveMethodNotAllowed {

	return &GetArchiveMethodNotAllowed{}
}

// WithPayload adds the payload to the get archive method not allowed response
func (o *GetArchiveMethodNotAllowed) WithPayload(payload *models.APIResponse) *GetArchiveMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get archive method not allowed response
func (o *GetArchiveMethodNotAllowed) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetArchiveMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetArchiveInternalServerErrorCode is the HTTP code returned for type GetArchiveInternalServerError
const GetArchiveInternalServerErrorCode int = 500

/*
GetArchiveInternalServerError TOC controller error

swagger:response getArchiveInternalServerError
*/
type GetArchiveInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetArchiveInternalServerError creates GetArchiveInternalServerError with default headers values
func NewGetArchiveInternalServerError() *GetArchiveInternalServerError {

	return &GetArchiveInternalServerError{}
}

// WithPayload adds the payload to the get archive internal server error response
func (o *GetArchiveInternalServerError) WithPayload(payload *models.APIResponse) *GetArchiveInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get archive internal server error response
func (o *GetArchiveInternalServerError) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetArchiveInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetArchiveServiceUnavailableCode is the HTTP code returned for type GetArchiveServiceUnavailable
const GetArchiveServiceUnavailableCode int = 503

/*
GetArchiveServiceUnavailable Service unavailable

swagger:response getArchiveServiceUnavailable
*/
type GetArchiveServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetArchiveServiceUnavailable creates GetArchiveServiceUnavailable with default headers values
func NewGetArchiveServiceUnavailable() *GetArchiveServiceUnavailable {

	return &GetArchiveServiceUnavailable{}
}

// WithPayload adds the payload to the get archive service unavailable response
func (o *GetArchiveServiceUnavailable) WithPayload(payload *models.APIResponse) *GetArchiveServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get archive service unavailable response
func (o *GetArchiveServiceUnavailable) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetArchiveServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}