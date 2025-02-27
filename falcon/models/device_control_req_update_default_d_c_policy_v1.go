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

// DeviceControlReqUpdateDefaultDCPolicyV1 device control req update default d c policy v1
//
// swagger:model device_control.ReqUpdateDefaultDCPolicyV1
type DeviceControlReqUpdateDefaultDCPolicyV1 struct {

	// custom notifications
	CustomNotifications *DeviceControlUSBCustomNotifications `json:"custom_notifications,omitempty"`
}

// Validate validates this device control req update default d c policy v1
func (m *DeviceControlReqUpdateDefaultDCPolicyV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomNotifications(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceControlReqUpdateDefaultDCPolicyV1) validateCustomNotifications(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomNotifications) { // not required
		return nil
	}

	if m.CustomNotifications != nil {
		if err := m.CustomNotifications.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("custom_notifications")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("custom_notifications")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this device control req update default d c policy v1 based on the context it is used
func (m *DeviceControlReqUpdateDefaultDCPolicyV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomNotifications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceControlReqUpdateDefaultDCPolicyV1) contextValidateCustomNotifications(ctx context.Context, formats strfmt.Registry) error {

	if m.CustomNotifications != nil {
		if err := m.CustomNotifications.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("custom_notifications")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("custom_notifications")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceControlReqUpdateDefaultDCPolicyV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceControlReqUpdateDefaultDCPolicyV1) UnmarshalBinary(b []byte) error {
	var res DeviceControlReqUpdateDefaultDCPolicyV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
