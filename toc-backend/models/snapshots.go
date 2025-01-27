// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Snapshots snapshots
//
// swagger:model Snapshots
type Snapshots struct {

	// db
	Db *Snapshot `json:"db,omitempty"`

	// file
	File *Snapshot `json:"file,omitempty"`

	// mail
	Mail *Snapshot `json:"mail,omitempty"`

	// web
	Web *Snapshot `json:"web,omitempty"`
}

// Validate validates this snapshots
func (m *Snapshots) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDb(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeb(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Snapshots) validateDb(formats strfmt.Registry) error {
	if swag.IsZero(m.Db) { // not required
		return nil
	}

	if m.Db != nil {
		if err := m.Db.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("db")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("db")
			}
			return err
		}
	}

	return nil
}

func (m *Snapshots) validateFile(formats strfmt.Registry) error {
	if swag.IsZero(m.File) { // not required
		return nil
	}

	if m.File != nil {
		if err := m.File.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("file")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("file")
			}
			return err
		}
	}

	return nil
}

func (m *Snapshots) validateMail(formats strfmt.Registry) error {
	if swag.IsZero(m.Mail) { // not required
		return nil
	}

	if m.Mail != nil {
		if err := m.Mail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mail")
			}
			return err
		}
	}

	return nil
}

func (m *Snapshots) validateWeb(formats strfmt.Registry) error {
	if swag.IsZero(m.Web) { // not required
		return nil
	}

	if m.Web != nil {
		if err := m.Web.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("web")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("web")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapshots based on the context it is used
func (m *Snapshots) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDb(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWeb(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Snapshots) contextValidateDb(ctx context.Context, formats strfmt.Registry) error {

	if m.Db != nil {

		if swag.IsZero(m.Db) { // not required
			return nil
		}

		if err := m.Db.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("db")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("db")
			}
			return err
		}
	}

	return nil
}

func (m *Snapshots) contextValidateFile(ctx context.Context, formats strfmt.Registry) error {

	if m.File != nil {

		if swag.IsZero(m.File) { // not required
			return nil
		}

		if err := m.File.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("file")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("file")
			}
			return err
		}
	}

	return nil
}

func (m *Snapshots) contextValidateMail(ctx context.Context, formats strfmt.Registry) error {

	if m.Mail != nil {

		if swag.IsZero(m.Mail) { // not required
			return nil
		}

		if err := m.Mail.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mail")
			}
			return err
		}
	}

	return nil
}

func (m *Snapshots) contextValidateWeb(ctx context.Context, formats strfmt.Registry) error {

	if m.Web != nil {

		if swag.IsZero(m.Web) { // not required
			return nil
		}

		if err := m.Web.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("web")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("web")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Snapshots) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Snapshots) UnmarshalBinary(b []byte) error {
	var res Snapshots
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
