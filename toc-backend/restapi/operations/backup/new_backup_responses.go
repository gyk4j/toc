// Code generated by go-swagger; DO NOT EDIT.

package backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gyk4j/toc/toc-backend/models"
)

// NewBackupOKCode is the HTTP code returned for type NewBackupOK
const NewBackupOKCode int = 200

/*
NewBackupOK Backup started

swagger:response newBackupOK
*/
type NewBackupOK struct {

	/*
	  In: Body
	*/
	Payload *models.Backup `json:"body,omitempty"`
}

// NewNewBackupOK creates NewBackupOK with default headers values
func NewNewBackupOK() *NewBackupOK {

	return &NewBackupOK{}
}

// WithPayload adds the payload to the new backup o k response
func (o *NewBackupOK) WithPayload(payload *models.Backup) *NewBackupOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the new backup o k response
func (o *NewBackupOK) SetPayload(payload *models.Backup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NewBackupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NewBackupForbiddenCode is the HTTP code returned for type NewBackupForbidden
const NewBackupForbiddenCode int = 403

/*
NewBackupForbidden Forbidden from creating duplicate/repeat backup

swagger:response newBackupForbidden
*/
type NewBackupForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewNewBackupForbidden creates NewBackupForbidden with default headers values
func NewNewBackupForbidden() *NewBackupForbidden {

	return &NewBackupForbidden{}
}

// WithPayload adds the payload to the new backup forbidden response
func (o *NewBackupForbidden) WithPayload(payload *models.APIResponse) *NewBackupForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the new backup forbidden response
func (o *NewBackupForbidden) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NewBackupForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NewBackupNotFoundCode is the HTTP code returned for type NewBackupNotFound
const NewBackupNotFoundCode int = 404

/*
NewBackupNotFound Main backup server not found

swagger:response newBackupNotFound
*/
type NewBackupNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewNewBackupNotFound creates NewBackupNotFound with default headers values
func NewNewBackupNotFound() *NewBackupNotFound {

	return &NewBackupNotFound{}
}

// WithPayload adds the payload to the new backup not found response
func (o *NewBackupNotFound) WithPayload(payload *models.APIResponse) *NewBackupNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the new backup not found response
func (o *NewBackupNotFound) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NewBackupNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NewBackupMethodNotAllowedCode is the HTTP code returned for type NewBackupMethodNotAllowed
const NewBackupMethodNotAllowedCode int = 405

/*
NewBackupMethodNotAllowed Bad request

swagger:response newBackupMethodNotAllowed
*/
type NewBackupMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewNewBackupMethodNotAllowed creates NewBackupMethodNotAllowed with default headers values
func NewNewBackupMethodNotAllowed() *NewBackupMethodNotAllowed {

	return &NewBackupMethodNotAllowed{}
}

// WithPayload adds the payload to the new backup method not allowed response
func (o *NewBackupMethodNotAllowed) WithPayload(payload *models.APIResponse) *NewBackupMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the new backup method not allowed response
func (o *NewBackupMethodNotAllowed) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NewBackupMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NewBackupInternalServerErrorCode is the HTTP code returned for type NewBackupInternalServerError
const NewBackupInternalServerErrorCode int = 500

/*
NewBackupInternalServerError Main TOC controller error

swagger:response newBackupInternalServerError
*/
type NewBackupInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewNewBackupInternalServerError creates NewBackupInternalServerError with default headers values
func NewNewBackupInternalServerError() *NewBackupInternalServerError {

	return &NewBackupInternalServerError{}
}

// WithPayload adds the payload to the new backup internal server error response
func (o *NewBackupInternalServerError) WithPayload(payload *models.APIResponse) *NewBackupInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the new backup internal server error response
func (o *NewBackupInternalServerError) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NewBackupInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NewBackupBadGatewayCode is the HTTP code returned for type NewBackupBadGateway
const NewBackupBadGatewayCode int = 502

/*
NewBackupBadGateway Bad gateway. Main backup server error

swagger:response newBackupBadGateway
*/
type NewBackupBadGateway struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewNewBackupBadGateway creates NewBackupBadGateway with default headers values
func NewNewBackupBadGateway() *NewBackupBadGateway {

	return &NewBackupBadGateway{}
}

// WithPayload adds the payload to the new backup bad gateway response
func (o *NewBackupBadGateway) WithPayload(payload *models.APIResponse) *NewBackupBadGateway {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the new backup bad gateway response
func (o *NewBackupBadGateway) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NewBackupBadGateway) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(502)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NewBackupServiceUnavailableCode is the HTTP code returned for type NewBackupServiceUnavailable
const NewBackupServiceUnavailableCode int = 503

/*
NewBackupServiceUnavailable Service unavailable on main TOC controller

swagger:response newBackupServiceUnavailable
*/
type NewBackupServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewNewBackupServiceUnavailable creates NewBackupServiceUnavailable with default headers values
func NewNewBackupServiceUnavailable() *NewBackupServiceUnavailable {

	return &NewBackupServiceUnavailable{}
}

// WithPayload adds the payload to the new backup service unavailable response
func (o *NewBackupServiceUnavailable) WithPayload(payload *models.APIResponse) *NewBackupServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the new backup service unavailable response
func (o *NewBackupServiceUnavailable) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NewBackupServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NewBackupGatewayTimeoutCode is the HTTP code returned for type NewBackupGatewayTimeout
const NewBackupGatewayTimeoutCode int = 504

/*
NewBackupGatewayTimeout Gateway timeout. Main backup server did not reply.

swagger:response newBackupGatewayTimeout
*/
type NewBackupGatewayTimeout struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewNewBackupGatewayTimeout creates NewBackupGatewayTimeout with default headers values
func NewNewBackupGatewayTimeout() *NewBackupGatewayTimeout {

	return &NewBackupGatewayTimeout{}
}

// WithPayload adds the payload to the new backup gateway timeout response
func (o *NewBackupGatewayTimeout) WithPayload(payload *models.APIResponse) *NewBackupGatewayTimeout {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the new backup gateway timeout response
func (o *NewBackupGatewayTimeout) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NewBackupGatewayTimeout) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(504)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
