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

// DomainAPIRemediationV2 domain API remediation v2
//
// swagger:model domain.APIRemediationV2
type DomainAPIRemediationV2 struct {

	// action
	// Required: true
	Action *string `json:"action"`

	// id
	// Required: true
	ID *string `json:"id"`

	// link
	// Required: true
	Link *string `json:"link"`

	// reference
	// Required: true
	Reference *string `json:"reference"`

	// title
	// Required: true
	Title *string `json:"title"`

	// vendor url
	// Required: true
	VendorURL *string `json:"vendor_url"`
}

// Validate validates this domain API remediation v2
func (m *DomainAPIRemediationV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVendorURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAPIRemediationV2) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *DomainAPIRemediationV2) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainAPIRemediationV2) validateLink(formats strfmt.Registry) error {

	if err := validate.Required("link", "body", m.Link); err != nil {
		return err
	}

	return nil
}

func (m *DomainAPIRemediationV2) validateReference(formats strfmt.Registry) error {

	if err := validate.Required("reference", "body", m.Reference); err != nil {
		return err
	}

	return nil
}

func (m *DomainAPIRemediationV2) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *DomainAPIRemediationV2) validateVendorURL(formats strfmt.Registry) error {

	if err := validate.Required("vendor_url", "body", m.VendorURL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain API remediation v2 based on context it is used
func (m *DomainAPIRemediationV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainAPIRemediationV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainAPIRemediationV2) UnmarshalBinary(b []byte) error {
	var res DomainAPIRemediationV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
