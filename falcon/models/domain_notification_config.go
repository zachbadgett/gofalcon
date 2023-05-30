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

// DomainNotificationConfig domain notification config
//
// swagger:model domain.NotificationConfig
type DomainNotificationConfig struct {

	// cid
	// Required: true
	Cid *string `json:"cid"`

	// config id
	// Required: true
	ConfigID *string `json:"config_id"`

	// plugin id
	// Required: true
	PluginID *string `json:"plugin_id"`

	// recipients
	// Required: true
	Recipients []string `json:"recipients"`

	// severity
	// Required: true
	Severity *string `json:"severity"`
}

// Validate validates this domain notification config
func (m *DomainNotificationConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePluginID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainNotificationConfig) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	return nil
}

func (m *DomainNotificationConfig) validateConfigID(formats strfmt.Registry) error {

	if err := validate.Required("config_id", "body", m.ConfigID); err != nil {
		return err
	}

	return nil
}

func (m *DomainNotificationConfig) validatePluginID(formats strfmt.Registry) error {

	if err := validate.Required("plugin_id", "body", m.PluginID); err != nil {
		return err
	}

	return nil
}

func (m *DomainNotificationConfig) validateRecipients(formats strfmt.Registry) error {

	if err := validate.Required("recipients", "body", m.Recipients); err != nil {
		return err
	}

	return nil
}

func (m *DomainNotificationConfig) validateSeverity(formats strfmt.Registry) error {

	if err := validate.Required("severity", "body", m.Severity); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain notification config based on context it is used
func (m *DomainNotificationConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainNotificationConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainNotificationConfig) UnmarshalBinary(b []byte) error {
	var res DomainNotificationConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
