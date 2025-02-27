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

// DomainAPIDetectionDocument domain API detection document
//
// swagger:model domain.APIDetectionDocument
type DomainAPIDetectionDocument struct {

	// adversary ids
	AdversaryIds []int64 `json:"adversary_ids"`

	// assigned to name
	AssignedToName string `json:"assigned_to_name,omitempty"`

	// assigned to uid
	AssignedToUID string `json:"assigned_to_uid,omitempty"`

	// behaviors
	Behaviors []*DetectsBehavior `json:"behaviors"`

	// behaviors processed
	// Required: true
	BehaviorsProcessed []string `json:"behaviors_processed"`

	// cid
	// Required: true
	Cid *string `json:"cid"`

	// created timestamp
	// Required: true
	// Format: date-time
	CreatedTimestamp *strfmt.DateTime `json:"created_timestamp"`

	// detection id
	// Required: true
	DetectionID *string `json:"detection_id"`

	// device
	// Required: true
	Device *DetectsDeviceDetailIndexed `json:"device"`

	// email sent
	// Required: true
	EmailSent *bool `json:"email_sent"`

	// first behavior
	// Required: true
	// Format: date-time
	FirstBehavior *strfmt.DateTime `json:"first_behavior"`

	// hostinfo
	// Required: true
	Hostinfo *DetectsHostInfo `json:"hostinfo"`

	// last behavior
	// Required: true
	// Format: date-time
	LastBehavior *strfmt.DateTime `json:"last_behavior"`

	// max confidence
	// Required: true
	MaxConfidence *int32 `json:"max_confidence"`

	// max severity
	// Required: true
	MaxSeverity *int32 `json:"max_severity"`

	// max severity displayname
	// Required: true
	MaxSeverityDisplayname *string `json:"max_severity_displayname"`

	// overwatch notes
	OverwatchNotes string `json:"overwatch_notes,omitempty"`

	// quarantined files
	QuarantinedFiles []*DetectsQuarantinedFile `json:"quarantined_files"`

	// seconds to resolved
	// Required: true
	SecondsToResolved *int64 `json:"seconds_to_resolved"`

	// seconds to triaged
	// Required: true
	SecondsToTriaged *int64 `json:"seconds_to_triaged"`

	// show in ui
	// Required: true
	ShowInUI *bool `json:"show_in_ui"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this domain API detection document
func (m *DomainAPIDetectionDocument) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBehaviors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBehaviorsProcessed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDetectionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailSent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstBehavior(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostinfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastBehavior(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxConfidence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxSeverity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxSeverityDisplayname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuarantinedFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondsToResolved(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondsToTriaged(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShowInUI(formats); err != nil {
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

func (m *DomainAPIDetectionDocument) validateBehaviors(formats strfmt.Registry) error {
	if swag.IsZero(m.Behaviors) { // not required
		return nil
	}

	for i := 0; i < len(m.Behaviors); i++ {
		if swag.IsZero(m.Behaviors[i]) { // not required
			continue
		}

		if m.Behaviors[i] != nil {
			if err := m.Behaviors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("behaviors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("behaviors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainAPIDetectionDocument) validateBehaviorsProcessed(formats strfmt.Registry) error {

	if err := validate.Required("behaviors_processed", "body", m.BehaviorsProcessed); err != nil {
		return err
	}

	return nil
}

func (m *DomainAPIDetectionDocument) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	return nil
}

func (m *DomainAPIDetectionDocument) validateCreatedTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("created_timestamp", "body", m.CreatedTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("created_timestamp", "body", "date-time", m.CreatedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainAPIDetectionDocument) validateDetectionID(formats strfmt.Registry) error {

	if err := validate.Required("detection_id", "body", m.DetectionID); err != nil {
		return err
	}

	return nil
}

func (m *DomainAPIDetectionDocument) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *DomainAPIDetectionDocument) validateEmailSent(formats strfmt.Registry) error {

	if err := validate.Required("email_sent", "body", m.EmailSent); err != nil {
		return err
	}

	return nil
}

func (m *DomainAPIDetectionDocument) validateFirstBehavior(formats strfmt.Registry) error {

	if err := validate.Required("first_behavior", "body", m.FirstBehavior); err != nil {
		return err
	}

	if err := validate.FormatOf("first_behavior", "body", "date-time", m.FirstBehavior.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainAPIDetectionDocument) validateHostinfo(formats strfmt.Registry) error {

	if err := validate.Required("hostinfo", "body", m.Hostinfo); err != nil {
		return err
	}

	if m.Hostinfo != nil {
		if err := m.Hostinfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostinfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hostinfo")
			}
			return err
		}
	}

	return nil
}

func (m *DomainAPIDetectionDocument) validateLastBehavior(formats strfmt.Registry) error {

	if err := validate.Required("last_behavior", "body", m.LastBehavior); err != nil {
		return err
	}

	if err := validate.FormatOf("last_behavior", "body", "date-time", m.LastBehavior.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainAPIDetectionDocument) validateMaxConfidence(formats strfmt.Registry) error {

	if err := validate.Required("max_confidence", "body", m.MaxConfidence); err != nil {
		return err
	}

	return nil
}

func (m *DomainAPIDetectionDocument) validateMaxSeverity(formats strfmt.Registry) error {

	if err := validate.Required("max_severity", "body", m.MaxSeverity); err != nil {
		return err
	}

	return nil
}

func (m *DomainAPIDetectionDocument) validateMaxSeverityDisplayname(formats strfmt.Registry) error {

	if err := validate.Required("max_severity_displayname", "body", m.MaxSeverityDisplayname); err != nil {
		return err
	}

	return nil
}

func (m *DomainAPIDetectionDocument) validateQuarantinedFiles(formats strfmt.Registry) error {
	if swag.IsZero(m.QuarantinedFiles) { // not required
		return nil
	}

	for i := 0; i < len(m.QuarantinedFiles); i++ {
		if swag.IsZero(m.QuarantinedFiles[i]) { // not required
			continue
		}

		if m.QuarantinedFiles[i] != nil {
			if err := m.QuarantinedFiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("quarantined_files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("quarantined_files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainAPIDetectionDocument) validateSecondsToResolved(formats strfmt.Registry) error {

	if err := validate.Required("seconds_to_resolved", "body", m.SecondsToResolved); err != nil {
		return err
	}

	return nil
}

func (m *DomainAPIDetectionDocument) validateSecondsToTriaged(formats strfmt.Registry) error {

	if err := validate.Required("seconds_to_triaged", "body", m.SecondsToTriaged); err != nil {
		return err
	}

	return nil
}

func (m *DomainAPIDetectionDocument) validateShowInUI(formats strfmt.Registry) error {

	if err := validate.Required("show_in_ui", "body", m.ShowInUI); err != nil {
		return err
	}

	return nil
}

func (m *DomainAPIDetectionDocument) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this domain API detection document based on the context it is used
func (m *DomainAPIDetectionDocument) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBehaviors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHostinfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuarantinedFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAPIDetectionDocument) contextValidateBehaviors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Behaviors); i++ {

		if m.Behaviors[i] != nil {
			if err := m.Behaviors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("behaviors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("behaviors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainAPIDetectionDocument) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.Device != nil {
		if err := m.Device.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *DomainAPIDetectionDocument) contextValidateHostinfo(ctx context.Context, formats strfmt.Registry) error {

	if m.Hostinfo != nil {
		if err := m.Hostinfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostinfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hostinfo")
			}
			return err
		}
	}

	return nil
}

func (m *DomainAPIDetectionDocument) contextValidateQuarantinedFiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.QuarantinedFiles); i++ {

		if m.QuarantinedFiles[i] != nil {
			if err := m.QuarantinedFiles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("quarantined_files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("quarantined_files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainAPIDetectionDocument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainAPIDetectionDocument) UnmarshalBinary(b []byte) error {
	var res DomainAPIDetectionDocument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
