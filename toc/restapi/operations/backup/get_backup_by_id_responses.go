// Code generated by go-swagger; DO NOT EDIT.

package backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gyk4j/toc/toc/models"
)

// GetBackupByIDOKCode is the HTTP code returned for type GetBackupByIDOK
const GetBackupByIDOKCode int = 200

/*
GetBackupByIDOK successful operation

swagger:response getBackupByIdOK
*/
type GetBackupByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Backup `json:"body,omitempty"`
}

// NewGetBackupByIDOK creates GetBackupByIDOK with default headers values
func NewGetBackupByIDOK() *GetBackupByIDOK {

	return &GetBackupByIDOK{}
}

// WithPayload adds the payload to the get backup by Id o k response
func (o *GetBackupByIDOK) WithPayload(payload *models.Backup) *GetBackupByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get backup by Id o k response
func (o *GetBackupByIDOK) SetPayload(payload *models.Backup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBackupByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBackupByIDBadRequestCode is the HTTP code returned for type GetBackupByIDBadRequest
const GetBackupByIDBadRequestCode int = 400

/*
GetBackupByIDBadRequest Invalid ID supplied

swagger:response getBackupByIdBadRequest
*/
type GetBackupByIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetBackupByIDBadRequest creates GetBackupByIDBadRequest with default headers values
func NewGetBackupByIDBadRequest() *GetBackupByIDBadRequest {

	return &GetBackupByIDBadRequest{}
}

// WithPayload adds the payload to the get backup by Id bad request response
func (o *GetBackupByIDBadRequest) WithPayload(payload *models.APIResponse) *GetBackupByIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get backup by Id bad request response
func (o *GetBackupByIDBadRequest) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBackupByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBackupByIDNotFoundCode is the HTTP code returned for type GetBackupByIDNotFound
const GetBackupByIDNotFoundCode int = 404

/*
GetBackupByIDNotFound Backup not found

swagger:response getBackupByIdNotFound
*/
type GetBackupByIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetBackupByIDNotFound creates GetBackupByIDNotFound with default headers values
func NewGetBackupByIDNotFound() *GetBackupByIDNotFound {

	return &GetBackupByIDNotFound{}
}

// WithPayload adds the payload to the get backup by Id not found response
func (o *GetBackupByIDNotFound) WithPayload(payload *models.APIResponse) *GetBackupByIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get backup by Id not found response
func (o *GetBackupByIDNotFound) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBackupByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
