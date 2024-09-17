// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Quota quota
// Example: {"hard":1073741824,"id":1,"path":"/home/toc/","soft":838860800}
//
// swagger:model Quota
type Quota struct {

	// hard
	Hard int64 `json:"hard"`

	// id
	ID int64 `json:"id"`

	// path
	Path string `json:"path"`

	// soft
	Soft int64 `json:"soft"`
}

// Validate validates this quota
func (m *Quota) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this quota based on context it is used
func (m *Quota) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Quota) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Quota) UnmarshalBinary(b []byte) error {
	var res Quota
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
