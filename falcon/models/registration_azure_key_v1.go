// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RegistrationAzureKeyV1 registration azure key v1
//
// swagger:model registration.AzureKeyV1
type RegistrationAzureKeyV1 struct {

	// client id
	// Required: true
	ClientID *string `json:"client_id"`

	// end date
	EndDate string `json:"end_date,omitempty"`

	// public certificate
	PublicCertificate string `json:"public_certificate,omitempty"`

	// tenant id
	// Required: true
	TenantID *string `json:"tenant_id"`

	// valid
	Valid bool `json:"valid,omitempty"`
}

// Validate validates this registration azure key v1
func (m *RegistrationAzureKeyV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistrationAzureKeyV1) validateClientID(formats strfmt.Registry) error {

	if err := validate.Required("client_id", "body", m.ClientID); err != nil {
		return err
	}

	return nil
}

func (m *RegistrationAzureKeyV1) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenant_id", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this registration azure key v1 based on context it is used
func (m *RegistrationAzureKeyV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RegistrationAzureKeyV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistrationAzureKeyV1) UnmarshalBinary(b []byte) error {
	var res RegistrationAzureKeyV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
