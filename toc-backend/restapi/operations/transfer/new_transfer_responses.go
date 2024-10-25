// Code generated by go-swagger; DO NOT EDIT.

package transfer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gyk4j/toc/toc-backend/models"
)

// NewTransferOKCode is the HTTP code returned for type NewTransferOK
const NewTransferOKCode int = 200

/*
NewTransferOK Transfer started

swagger:response newTransferOK
*/
type NewTransferOK struct {

	/*
	  In: Body
	*/
	Payload *models.Transfer `json:"body,omitempty"`
}

// NewNewTransferOK creates NewTransferOK with default headers values
func NewNewTransferOK() *NewTransferOK {

	return &NewTransferOK{}
}

// WithPayload adds the payload to the new transfer o k response
func (o *NewTransferOK) WithPayload(payload *models.Transfer) *NewTransferOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the new transfer o k response
func (o *NewTransferOK) SetPayload(payload *models.Transfer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NewTransferOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NewTransferForbiddenCode is the HTTP code returned for type NewTransferForbidden
const NewTransferForbiddenCode int = 403

/*
NewTransferForbidden Forbidden from creating duplicate/repeat transfer

swagger:response newTransferForbidden
*/
type NewTransferForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewNewTransferForbidden creates NewTransferForbidden with default headers values
func NewNewTransferForbidden() *NewTransferForbidden {

	return &NewTransferForbidden{}
}

// WithPayload adds the payload to the new transfer forbidden response
func (o *NewTransferForbidden) WithPayload(payload *models.APIResponse) *NewTransferForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the new transfer forbidden response
func (o *NewTransferForbidden) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NewTransferForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NewTransferNotFoundCode is the HTTP code returned for type NewTransferNotFound
const NewTransferNotFoundCode int = 404

/*
NewTransferNotFound Stepup backup server not found

swagger:response newTransferNotFound
*/
type NewTransferNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewNewTransferNotFound creates NewTransferNotFound with default headers values
func NewNewTransferNotFound() *NewTransferNotFound {

	return &NewTransferNotFound{}
}

// WithPayload adds the payload to the new transfer not found response
func (o *NewTransferNotFound) WithPayload(payload *models.APIResponse) *NewTransferNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the new transfer not found response
func (o *NewTransferNotFound) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NewTransferNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NewTransferMethodNotAllowedCode is the HTTP code returned for type NewTransferMethodNotAllowed
const NewTransferMethodNotAllowedCode int = 405

/*
NewTransferMethodNotAllowed Bad request

swagger:response newTransferMethodNotAllowed
*/
type NewTransferMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewNewTransferMethodNotAllowed creates NewTransferMethodNotAllowed with default headers values
func NewNewTransferMethodNotAllowed() *NewTransferMethodNotAllowed {

	return &NewTransferMethodNotAllowed{}
}

// WithPayload adds the payload to the new transfer method not allowed response
func (o *NewTransferMethodNotAllowed) WithPayload(payload *models.APIResponse) *NewTransferMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the new transfer method not allowed response
func (o *NewTransferMethodNotAllowed) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NewTransferMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NewTransferInternalServerErrorCode is the HTTP code returned for type NewTransferInternalServerError
const NewTransferInternalServerErrorCode int = 500

/*
NewTransferInternalServerError Stepup TOC controller error

swagger:response newTransferInternalServerError
*/
type NewTransferInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewNewTransferInternalServerError creates NewTransferInternalServerError with default headers values
func NewNewTransferInternalServerError() *NewTransferInternalServerError {

	return &NewTransferInternalServerError{}
}

// WithPayload adds the payload to the new transfer internal server error response
func (o *NewTransferInternalServerError) WithPayload(payload *models.APIResponse) *NewTransferInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the new transfer internal server error response
func (o *NewTransferInternalServerError) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NewTransferInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NewTransferBadGatewayCode is the HTTP code returned for type NewTransferBadGateway
const NewTransferBadGatewayCode int = 502

/*
NewTransferBadGateway Bad gateway. Stepup backup server error

swagger:response newTransferBadGateway
*/
type NewTransferBadGateway struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewNewTransferBadGateway creates NewTransferBadGateway with default headers values
func NewNewTransferBadGateway() *NewTransferBadGateway {

	return &NewTransferBadGateway{}
}

// WithPayload adds the payload to the new transfer bad gateway response
func (o *NewTransferBadGateway) WithPayload(payload *models.APIResponse) *NewTransferBadGateway {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the new transfer bad gateway response
func (o *NewTransferBadGateway) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NewTransferBadGateway) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(502)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NewTransferServiceUnavailableCode is the HTTP code returned for type NewTransferServiceUnavailable
const NewTransferServiceUnavailableCode int = 503

/*
NewTransferServiceUnavailable Service unavailable on stepup TOC controller

swagger:response newTransferServiceUnavailable
*/
type NewTransferServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewNewTransferServiceUnavailable creates NewTransferServiceUnavailable with default headers values
func NewNewTransferServiceUnavailable() *NewTransferServiceUnavailable {

	return &NewTransferServiceUnavailable{}
}

// WithPayload adds the payload to the new transfer service unavailable response
func (o *NewTransferServiceUnavailable) WithPayload(payload *models.APIResponse) *NewTransferServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the new transfer service unavailable response
func (o *NewTransferServiceUnavailable) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NewTransferServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NewTransferGatewayTimeoutCode is the HTTP code returned for type NewTransferGatewayTimeout
const NewTransferGatewayTimeoutCode int = 504

/*
NewTransferGatewayTimeout Gateway timeout. Stepup backup server did not reply.

swagger:response newTransferGatewayTimeout
*/
type NewTransferGatewayTimeout struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewNewTransferGatewayTimeout creates NewTransferGatewayTimeout with default headers values
func NewNewTransferGatewayTimeout() *NewTransferGatewayTimeout {

	return &NewTransferGatewayTimeout{}
}

// WithPayload adds the payload to the new transfer gateway timeout response
func (o *NewTransferGatewayTimeout) WithPayload(payload *models.APIResponse) *NewTransferGatewayTimeout {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the new transfer gateway timeout response
func (o *NewTransferGatewayTimeout) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NewTransferGatewayTimeout) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(504)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
