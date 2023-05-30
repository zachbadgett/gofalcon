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

// DomainScanHostMetadata domain scan host metadata
//
// swagger:model domain.ScanHostMetadata
type DomainScanHostMetadata struct {

	// cid
	Cid string `json:"cid,omitempty"`

	// completed on
	// Format: date-time
	CompletedOn strfmt.DateTime `json:"completed_on,omitempty"`

	// filecount
	Filecount *DomainFileCount `json:"filecount,omitempty"`

	// host id
	HostID string `json:"host_id,omitempty"`

	// host scan id
	HostScanID string `json:"host_scan_id,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// last updated
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// profile id
	ProfileID string `json:"profile_id,omitempty"`

	// scan control reason
	ScanControlReason string `json:"scan_control_reason,omitempty"`

	// scan id
	ScanID string `json:"scan_id,omitempty"`

	// severity
	Severity int64 `json:"severity,omitempty"`

	// started on
	// Format: date-time
	StartedOn strfmt.DateTime `json:"started_on,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this domain scan host metadata
func (m *DomainScanHostMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilecount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedOn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainScanHostMetadata) validateCompletedOn(formats strfmt.Registry) error {
	if swag.IsZero(m.CompletedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("completed_on", "body", "date-time", m.CompletedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainScanHostMetadata) validateFilecount(formats strfmt.Registry) error {
	if swag.IsZero(m.Filecount) { // not required
		return nil
	}

	if m.Filecount != nil {
		if err := m.Filecount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filecount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filecount")
			}
			return err
		}
	}

	return nil
}

func (m *DomainScanHostMetadata) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainScanHostMetadata) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainScanHostMetadata) validateStartedOn(formats strfmt.Registry) error {
	if swag.IsZero(m.StartedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("started_on", "body", "date-time", m.StartedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this domain scan host metadata based on the context it is used
func (m *DomainScanHostMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilecount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainScanHostMetadata) contextValidateFilecount(ctx context.Context, formats strfmt.Registry) error {

	if m.Filecount != nil {
		if err := m.Filecount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filecount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filecount")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainScanHostMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainScanHostMetadata) UnmarshalBinary(b []byte) error {
	var res DomainScanHostMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
