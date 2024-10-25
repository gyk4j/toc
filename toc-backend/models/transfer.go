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

// Transfer transfer
// Example: {"id":1,"status":"failed"}
//
// swagger:model Transfer
type Transfer struct {

	// backup
	Backup *Backup `json:"backup,omitempty"`

	// id
	ID int64 `json:"id"`

	// Transfer Status
	// Enum: [queued in-progress completed failed]
	Status string `json:"status"`
}

// Validate validates this transfer
func (m *Transfer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Transfer) validateBackup(formats strfmt.Registry) error {
	if swag.IsZero(m.Backup) { // not required
		return nil
	}

	if m.Backup != nil {
		if err := m.Backup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backup")
			}
			return err
		}
	}

	return nil
}

var transferTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["queued","in-progress","completed","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transferTypeStatusPropEnum = append(transferTypeStatusPropEnum, v)
	}
}

const (

	// TransferStatusQueued captures enum value "queued"
	TransferStatusQueued string = "queued"

	// TransferStatusInDashProgress captures enum value "in-progress"
	TransferStatusInDashProgress string = "in-progress"

	// TransferStatusCompleted captures enum value "completed"
	TransferStatusCompleted string = "completed"

	// TransferStatusFailed captures enum value "failed"
	TransferStatusFailed string = "failed"
)

// prop value enum
func (m *Transfer) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transferTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Transfer) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this transfer based on the context it is used
func (m *Transfer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Transfer) contextValidateBackup(ctx context.Context, formats strfmt.Registry) error {

	if m.Backup != nil {

		if swag.IsZero(m.Backup) { // not required
			return nil
		}

		if err := m.Backup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backup")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Transfer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Transfer) UnmarshalBinary(b []byte) error {
	var res Transfer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}