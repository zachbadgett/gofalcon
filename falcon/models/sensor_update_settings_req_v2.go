// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SensorUpdateSettingsReqV2 sensor update settings req v2
//
// swagger:model sensor_update.SettingsReqV2
type SensorUpdateSettingsReqV2 struct {

	// The target build to apply to the policy
	Build string `json:"build,omitempty"`

	// The schedule that disables sensor updates
	// Required: true
	Scheduler *PolicySensorUpdateScheduler `json:"scheduler"`

	// If true, early adopter builds will be visible on the sensor update policy page
	ShowEarlyAdopterBuilds bool `json:"show_early_adopter_builds,omitempty"`

	// The uninstall protection state to apply to the policy
	// Enum: [ENABLED DISABLED MAINTENANCE_MODE IGNORE UNKNOWN]
	UninstallProtection string `json:"uninstall_protection,omitempty"`

	// variants
	// Required: true
	Variants []*SensorUpdateBuildReqV1 `json:"variants"`
}

// Validate validates this sensor update settings req v2
func (m *SensorUpdateSettingsReqV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScheduler(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUninstallProtection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariants(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SensorUpdateSettingsReqV2) validateScheduler(formats strfmt.Registry) error {

	if err := validate.Required("scheduler", "body", m.Scheduler); err != nil {
		return err
	}

	if m.Scheduler != nil {
		if err := m.Scheduler.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scheduler")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scheduler")
			}
			return err
		}
	}

	return nil
}

var sensorUpdateSettingsReqV2TypeUninstallProtectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","DISABLED","MAINTENANCE_MODE","IGNORE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sensorUpdateSettingsReqV2TypeUninstallProtectionPropEnum = append(sensorUpdateSettingsReqV2TypeUninstallProtectionPropEnum, v)
	}
}

const (

	// SensorUpdateSettingsReqV2UninstallProtectionENABLED captures enum value "ENABLED"
	SensorUpdateSettingsReqV2UninstallProtectionENABLED string = "ENABLED"

	// SensorUpdateSettingsReqV2UninstallProtectionDISABLED captures enum value "DISABLED"
	SensorUpdateSettingsReqV2UninstallProtectionDISABLED string = "DISABLED"

	// SensorUpdateSettingsReqV2UninstallProtectionMAINTENANCEMODE captures enum value "MAINTENANCE_MODE"
	SensorUpdateSettingsReqV2UninstallProtectionMAINTENANCEMODE string = "MAINTENANCE_MODE"

	// SensorUpdateSettingsReqV2UninstallProtectionIGNORE captures enum value "IGNORE"
	SensorUpdateSettingsReqV2UninstallProtectionIGNORE string = "IGNORE"

	// SensorUpdateSettingsReqV2UninstallProtectionUNKNOWN captures enum value "UNKNOWN"
	SensorUpdateSettingsReqV2UninstallProtectionUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (m *SensorUpdateSettingsReqV2) validateUninstallProtectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sensorUpdateSettingsReqV2TypeUninstallProtectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SensorUpdateSettingsReqV2) validateUninstallProtection(formats strfmt.Registry) error {
	if swag.IsZero(m.UninstallProtection) { // not required
		return nil
	}

	// value enum
	if err := m.validateUninstallProtectionEnum("uninstall_protection", "body", m.UninstallProtection); err != nil {
		return err
	}

	return nil
}

func (m *SensorUpdateSettingsReqV2) validateVariants(formats strfmt.Registry) error {

	if err := validate.Required("variants", "body", m.Variants); err != nil {
		return err
	}

	for i := 0; i < len(m.Variants); i++ {
		if swag.IsZero(m.Variants[i]) { // not required
			continue
		}

		if m.Variants[i] != nil {
			if err := m.Variants[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variants" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variants" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this sensor update settings req v2 based on the context it is used
func (m *SensorUpdateSettingsReqV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScheduler(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVariants(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SensorUpdateSettingsReqV2) contextValidateScheduler(ctx context.Context, formats strfmt.Registry) error {

	if m.Scheduler != nil {
		if err := m.Scheduler.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scheduler")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scheduler")
			}
			return err
		}
	}

	return nil
}

func (m *SensorUpdateSettingsReqV2) contextValidateVariants(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Variants); i++ {

		if m.Variants[i] != nil {
			if err := m.Variants[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variants" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variants" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SensorUpdateSettingsReqV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SensorUpdateSettingsReqV2) UnmarshalBinary(b []byte) error {
	var res SensorUpdateSettingsReqV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
