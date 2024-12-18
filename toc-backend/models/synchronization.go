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

// Synchronization synchronization
// Example: {"id":1,"status":"running","time":"2024-12-31T23:59:59Z"}
//
// swagger:model Synchronization
type Synchronization struct {

	// id
	ID int64 `json:"id"`

	// Synchronization Status
	// Enum: [queued running completed failed]
	Status string `json:"status"`

	// time
	// Format: date-time
	Time strfmt.DateTime `json:"time"`
}

// Validate validates this synchronization
func (m *Synchronization) Validate(formats strfmt.Registry) error {
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

var synchronizationTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["queued","running","completed","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		synchronizationTypeStatusPropEnum = append(synchronizationTypeStatusPropEnum, v)
	}
}

const (

	// SynchronizationStatusQueued captures enum value "queued"
	SynchronizationStatusQueued string = "queued"

	// SynchronizationStatusRunning captures enum value "running"
	SynchronizationStatusRunning string = "running"

	// SynchronizationStatusCompleted captures enum value "completed"
	SynchronizationStatusCompleted string = "completed"

	// SynchronizationStatusFailed captures enum value "failed"
	SynchronizationStatusFailed string = "failed"
)

// prop value enum
func (m *Synchronization) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, synchronizationTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Synchronization) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Synchronization) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this synchronization based on context it is used
func (m *Synchronization) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Synchronization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Synchronization) UnmarshalBinary(b []byte) error {
	var res Synchronization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
