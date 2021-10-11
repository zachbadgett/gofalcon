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
)

// FalconxSandboxSummaryReportV1 falconx sandbox summary report v1
//
// swagger:model falconx.SandboxSummaryReportV1
type FalconxSandboxSummaryReportV1 struct {

	// classification tags
	ClassificationTags []string `json:"classification_tags"`

	// environment description
	EnvironmentDescription string `json:"environment_description,omitempty"`

	// environment id
	EnvironmentID int32 `json:"environment_id,omitempty"`

	// error message
	ErrorMessage string `json:"error_message,omitempty"`

	// error origin
	ErrorOrigin string `json:"error_origin,omitempty"`

	// error type
	ErrorType string `json:"error_type,omitempty"`

	// file type
	FileType string `json:"file_type,omitempty"`

	// incidents
	Incidents []*FalconxIncident `json:"incidents"`

	// sample flags
	SampleFlags []string `json:"sample_flags"`

	// sha256
	Sha256 string `json:"sha256,omitempty"`

	// submission type
	SubmissionType string `json:"submission_type,omitempty"`

	// submit name
	SubmitName string `json:"submit_name,omitempty"`

	// submit url
	SubmitURL string `json:"submit_url,omitempty"`

	// threat score
	ThreatScore int32 `json:"threat_score,omitempty"`

	// verdict
	Verdict string `json:"verdict,omitempty"`
}

// Validate validates this falconx sandbox summary report v1
func (m *FalconxSandboxSummaryReportV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIncidents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FalconxSandboxSummaryReportV1) validateIncidents(formats strfmt.Registry) error {
	if swag.IsZero(m.Incidents) { // not required
		return nil
	}

	for i := 0; i < len(m.Incidents); i++ {
		if swag.IsZero(m.Incidents[i]) { // not required
			continue
		}

		if m.Incidents[i] != nil {
			if err := m.Incidents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("incidents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("incidents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this falconx sandbox summary report v1 based on the context it is used
func (m *FalconxSandboxSummaryReportV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIncidents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FalconxSandboxSummaryReportV1) contextValidateIncidents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Incidents); i++ {

		if m.Incidents[i] != nil {
			if err := m.Incidents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("incidents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("incidents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FalconxSandboxSummaryReportV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FalconxSandboxSummaryReportV1) UnmarshalBinary(b []byte) error {
	var res FalconxSandboxSummaryReportV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
