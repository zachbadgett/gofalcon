// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeviceMappedDevicePolicies device mapped device policies
//
// swagger:model device.MappedDevicePolicies
type DeviceMappedDevicePolicies struct {

	// airlock
	Airlock *DeviceDevicePolicy `json:"airlock,omitempty"`

	// automox
	Automox *DeviceDevicePolicy `json:"automox,omitempty"`

	// device control
	DeviceControl *DeviceDevicePolicy `json:"device_control,omitempty"`

	// firewall
	Firewall *DeviceDevicePolicy `json:"firewall,omitempty"`

	// global config
	GlobalConfig *DeviceDevicePolicy `json:"global_config,omitempty"`

	// identity protection
	IdentityProtection *DeviceDevicePolicy `json:"identity-protection,omitempty"`

	// mobile
	Mobile *DeviceDevicePolicy `json:"mobile,omitempty"`

	// netskope
	Netskope *DeviceDevicePolicy `json:"netskope,omitempty"`

	// prevention
	Prevention *DeviceDevicePolicy `json:"prevention,omitempty"`

	// remote response
	RemoteResponse *DeviceDevicePolicy `json:"remote_response,omitempty"`

	// sensor update
	SensorUpdate *DeviceDevicePolicy `json:"sensor_update,omitempty"`
}

// Validate validates this device mapped device policies
func (m *DeviceMappedDevicePolicies) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAirlock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutomox(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceControl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirewall(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGlobalConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentityProtection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMobile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetskope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrevention(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteResponse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSensorUpdate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceMappedDevicePolicies) validateAirlock(formats strfmt.Registry) error {
	if swag.IsZero(m.Airlock) { // not required
		return nil
	}

	if m.Airlock != nil {
		if err := m.Airlock.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("airlock")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("airlock")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceMappedDevicePolicies) validateAutomox(formats strfmt.Registry) error {
	if swag.IsZero(m.Automox) { // not required
		return nil
	}

	if m.Automox != nil {
		if err := m.Automox.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("automox")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("automox")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceMappedDevicePolicies) validateDeviceControl(formats strfmt.Registry) error {
	if swag.IsZero(m.DeviceControl) { // not required
		return nil
	}

	if m.DeviceControl != nil {
		if err := m.DeviceControl.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_control")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device_control")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceMappedDevicePolicies) validateFirewall(formats strfmt.Registry) error {
	if swag.IsZero(m.Firewall) { // not required
		return nil
	}

	if m.Firewall != nil {
		if err := m.Firewall.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firewall")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("firewall")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceMappedDevicePolicies) validateGlobalConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.GlobalConfig) { // not required
		return nil
	}

	if m.GlobalConfig != nil {
		if err := m.GlobalConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("global_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("global_config")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceMappedDevicePolicies) validateIdentityProtection(formats strfmt.Registry) error {
	if swag.IsZero(m.IdentityProtection) { // not required
		return nil
	}

	if m.IdentityProtection != nil {
		if err := m.IdentityProtection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity-protection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identity-protection")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceMappedDevicePolicies) validateMobile(formats strfmt.Registry) error {
	if swag.IsZero(m.Mobile) { // not required
		return nil
	}

	if m.Mobile != nil {
		if err := m.Mobile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mobile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mobile")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceMappedDevicePolicies) validateNetskope(formats strfmt.Registry) error {
	if swag.IsZero(m.Netskope) { // not required
		return nil
	}

	if m.Netskope != nil {
		if err := m.Netskope.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netskope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("netskope")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceMappedDevicePolicies) validatePrevention(formats strfmt.Registry) error {
	if swag.IsZero(m.Prevention) { // not required
		return nil
	}

	if m.Prevention != nil {
		if err := m.Prevention.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prevention")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("prevention")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceMappedDevicePolicies) validateRemoteResponse(formats strfmt.Registry) error {
	if swag.IsZero(m.RemoteResponse) { // not required
		return nil
	}

	if m.RemoteResponse != nil {
		if err := m.RemoteResponse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remote_response")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("remote_response")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceMappedDevicePolicies) validateSensorUpdate(formats strfmt.Registry) error {
	if swag.IsZero(m.SensorUpdate) { // not required
		return nil
	}

	if m.SensorUpdate != nil {
		if err := m.SensorUpdate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sensor_update")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sensor_update")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this device mapped device policies based on the context it is used
func (m *DeviceMappedDevicePolicies) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAirlock(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAutomox(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceControl(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirewall(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGlobalConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentityProtection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMobile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetskope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrevention(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemoteResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSensorUpdate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceMappedDevicePolicies) contextValidateAirlock(ctx context.Context, formats strfmt.Registry) error {

	if m.Airlock != nil {
		if err := m.Airlock.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("airlock")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("airlock")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceMappedDevicePolicies) contextValidateAutomox(ctx context.Context, formats strfmt.Registry) error {

	if m.Automox != nil {
		if err := m.Automox.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("automox")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("automox")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceMappedDevicePolicies) contextValidateDeviceControl(ctx context.Context, formats strfmt.Registry) error {

	if m.DeviceControl != nil {
		if err := m.DeviceControl.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_control")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device_control")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceMappedDevicePolicies) contextValidateFirewall(ctx context.Context, formats strfmt.Registry) error {

	if m.Firewall != nil {
		if err := m.Firewall.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firewall")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("firewall")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceMappedDevicePolicies) contextValidateGlobalConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.GlobalConfig != nil {
		if err := m.GlobalConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("global_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("global_config")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceMappedDevicePolicies) contextValidateIdentityProtection(ctx context.Context, formats strfmt.Registry) error {

	if m.IdentityProtection != nil {
		if err := m.IdentityProtection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity-protection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identity-protection")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceMappedDevicePolicies) contextValidateMobile(ctx context.Context, formats strfmt.Registry) error {

	if m.Mobile != nil {
		if err := m.Mobile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mobile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mobile")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceMappedDevicePolicies) contextValidateNetskope(ctx context.Context, formats strfmt.Registry) error {

	if m.Netskope != nil {
		if err := m.Netskope.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netskope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("netskope")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceMappedDevicePolicies) contextValidatePrevention(ctx context.Context, formats strfmt.Registry) error {

	if m.Prevention != nil {
		if err := m.Prevention.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prevention")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("prevention")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceMappedDevicePolicies) contextValidateRemoteResponse(ctx context.Context, formats strfmt.Registry) error {

	if m.RemoteResponse != nil {
		if err := m.RemoteResponse.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remote_response")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("remote_response")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceMappedDevicePolicies) contextValidateSensorUpdate(ctx context.Context, formats strfmt.Registry) error {

	if m.SensorUpdate != nil {
		if err := m.SensorUpdate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sensor_update")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sensor_update")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceMappedDevicePolicies) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceMappedDevicePolicies) UnmarshalBinary(b []byte) error {
	var res DeviceMappedDevicePolicies
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
