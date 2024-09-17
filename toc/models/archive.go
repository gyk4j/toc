// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Archive archive
// Example: {"id":1,"status":"queued","time":"2024-12-31T23:59:59Z","url":"/archives/1"}
//
// swagger:model Archive
type Archive struct {

	// id
	ID int64 `json:"id"`

	// Archive Status
	// Enum: [queued in-progress completed failed]
	Status string `json:"status"`

	// time
	// Format: date-time
	Time strfmt.DateTime `json:"time"`

	// url
	URL string `json:"url"`
}

// Validate validates this archive
func (m *Archive) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var archiveTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["queued","in-progress","completed","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		archiveTypeStatusPropEnum = append(archiveTypeStatusPropEnum, v)
	}
}

const (

	// ArchiveStatusQueued captures enum value "queued"
	ArchiveStatusQueued string = "queued"

	// ArchiveStatusInDashProgress captures enum value "in-progress"
	ArchiveStatusInDashProgress string = "in-progress"

	// ArchiveStatusCompleted captures enum value "completed"
	ArchiveStatusCompleted string = "completed"

	// ArchiveStatusFailed captures enum value "failed"
	ArchiveStatusFailed string = "failed"
)

// prop value enum
func (m *Archive) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, archiveTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Archive) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Archive) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this archive based on context it is used
func (m *Archive) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Archive) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Archive) UnmarshalBinary(b []byte) error {
	var res Archive
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
