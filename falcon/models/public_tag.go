// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PublicTag public tag
//
// swagger:model public.Tag
type PublicTag struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this public tag
func (m *PublicTag) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this public tag based on context it is used
func (m *PublicTag) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PublicTag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicTag) UnmarshalBinary(b []byte) error {
	var res PublicTag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
