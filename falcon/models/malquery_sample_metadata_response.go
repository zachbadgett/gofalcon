// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MalquerySampleMetadataResponse malquery sample metadata response
//
// swagger:model malquery.SampleMetadataResponse
type MalquerySampleMetadataResponse struct {

	// Errors that occurred during the request
	// Required: true
	Errors []*MalqueryQueryError `json:"errors"`

	// Meta information about the request
	// Required: true
	Meta *MalquerySamplesMetadataMetaInfo `json:"meta"`

	// List of sample metadata
	// Required: true
	Resources []*MalquerySampleMetadata `json:"resources"`
}

// Validate validates this malquery sample metadata response
func (m *MalquerySampleMetadataResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MalquerySampleMetadataResponse) validateErrors(formats strfmt.Registry) error {

	if err := validate.Required("errors", "body", m.Errors); err != nil {
		return err
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MalquerySampleMetadataResponse) validateMeta(formats strfmt.Registry) error {

	if err := validate.Required("meta", "body", m.Meta); err != nil {
		return err
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

func (m *MalquerySampleMetadataResponse) validateResources(formats strfmt.Registry) error {

	if err := validate.Required("resources", "body", m.Resources); err != nil {
		return err
	}

	for i := 0; i < len(m.Resources); i++ {
		if swag.IsZero(m.Resources[i]) { // not required
			continue
		}

		if m.Resources[i] != nil {
			if err := m.Resources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this malquery sample metadata response based on the context it is used
func (m *MalquerySampleMetadataResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MalquerySampleMetadataResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {
			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MalquerySampleMetadataResponse) contextValidateMeta(ctx context.Context, formats strfmt.Registry) error {

	if m.Meta != nil {
		if err := m.Meta.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

func (m *MalquerySampleMetadataResponse) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Resources); i++ {

		if m.Resources[i] != nil {
			if err := m.Resources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MalquerySampleMetadataResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MalquerySampleMetadataResponse) UnmarshalBinary(b []byte) error {
	var res MalquerySampleMetadataResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
