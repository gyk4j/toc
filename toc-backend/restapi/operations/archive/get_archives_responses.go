// Code generated by go-swagger; DO NOT EDIT.

package archive

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gyk4j/toc/toc-backend/models"
)

// GetArchivesOKCode is the HTTP code returned for type GetArchivesOK
const GetArchivesOKCode int = 200

/*
GetArchivesOK Archive info retrieved

swagger:response getArchivesOK
*/
type GetArchivesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Archive `json:"body,omitempty"`
}

// NewGetArchivesOK creates GetArchivesOK with default headers values
func NewGetArchivesOK() *GetArchivesOK {

	return &GetArchivesOK{}
}

// WithPayload adds the payload to the get archives o k response
func (o *GetArchivesOK) WithPayload(payload []*models.Archive) *GetArchivesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get archives o k response
func (o *GetArchivesOK) SetPayload(payload []*models.Archive) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetArchivesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetArchivesForbiddenCode is the HTTP code returned for type GetArchivesForbidden
const GetArchivesForbiddenCode int = 403

/*
GetArchivesForbidden Forbidden from querying archive info

swagger:response getArchivesForbidden
*/
type GetArchivesForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetArchivesForbidden creates GetArchivesForbidden with default headers values
func NewGetArchivesForbidden() *GetArchivesForbidden {

	return &GetArchivesForbidden{}
}

// WithPayload adds the payload to the get archives forbidden response
func (o *GetArchivesForbidden) WithPayload(payload *models.APIResponse) *GetArchivesForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get archives forbidden response
func (o *GetArchivesForbidden) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetArchivesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetArchivesNotFoundCode is the HTTP code returned for type GetArchivesNotFound
const GetArchivesNotFoundCode int = 404

/*
GetArchivesNotFound Archive not found

swagger:response getArchivesNotFound
*/
type GetArchivesNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetArchivesNotFound creates GetArchivesNotFound with default headers values
func NewGetArchivesNotFound() *GetArchivesNotFound {

	return &GetArchivesNotFound{}
}

// WithPayload adds the payload to the get archives not found response
func (o *GetArchivesNotFound) WithPayload(payload *models.APIResponse) *GetArchivesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get archives not found response
func (o *GetArchivesNotFound) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetArchivesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetArchivesMethodNotAllowedCode is the HTTP code returned for type GetArchivesMethodNotAllowed
const GetArchivesMethodNotAllowedCode int = 405

/*
GetArchivesMethodNotAllowed Bad request

swagger:response getArchivesMethodNotAllowed
*/
type GetArchivesMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetArchivesMethodNotAllowed creates GetArchivesMethodNotAllowed with default headers values
func NewGetArchivesMethodNotAllowed() *GetArchivesMethodNotAllowed {

	return &GetArchivesMethodNotAllowed{}
}

// WithPayload adds the payload to the get archives method not allowed response
func (o *GetArchivesMethodNotAllowed) WithPayload(payload *models.APIResponse) *GetArchivesMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get archives method not allowed response
func (o *GetArchivesMethodNotAllowed) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetArchivesMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetArchivesInternalServerErrorCode is the HTTP code returned for type GetArchivesInternalServerError
const GetArchivesInternalServerErrorCode int = 500

/*
GetArchivesInternalServerError TOC controller error

swagger:response getArchivesInternalServerError
*/
type GetArchivesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetArchivesInternalServerError creates GetArchivesInternalServerError with default headers values
func NewGetArchivesInternalServerError() *GetArchivesInternalServerError {

	return &GetArchivesInternalServerError{}
}

// WithPayload adds the payload to the get archives internal server error response
func (o *GetArchivesInternalServerError) WithPayload(payload *models.APIResponse) *GetArchivesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get archives internal server error response
func (o *GetArchivesInternalServerError) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetArchivesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetArchivesServiceUnavailableCode is the HTTP code returned for type GetArchivesServiceUnavailable
const GetArchivesServiceUnavailableCode int = 503

/*
GetArchivesServiceUnavailable Service unavailable

swagger:response getArchivesServiceUnavailable
*/
type GetArchivesServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetArchivesServiceUnavailable creates GetArchivesServiceUnavailable with default headers values
func NewGetArchivesServiceUnavailable() *GetArchivesServiceUnavailable {

	return &GetArchivesServiceUnavailable{}
}

// WithPayload adds the payload to the get archives service unavailable response
func (o *GetArchivesServiceUnavailable) WithPayload(payload *models.APIResponse) *GetArchivesServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get archives service unavailable response
func (o *GetArchivesServiceUnavailable) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetArchivesServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
