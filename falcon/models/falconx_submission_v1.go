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

// FalconxSubmissionV1 falconx submission v1
//
// swagger:model falconx.SubmissionV1
type FalconxSubmissionV1 struct {

	// cid
	Cid string `json:"cid,omitempty"`

	// created timestamp
	CreatedTimestamp string `json:"created_timestamp,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// index timestamp
	IndexTimestamp string `json:"index_timestamp,omitempty"`

	// origin
	Origin string `json:"origin,omitempty"`

	// sandbox
	Sandbox []*FalconxSandboxParametersV1 `json:"sandbox"`

	// send email notification
	SendEmailNotification bool `json:"send_email_notification,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`

	// user name
	UserName string `json:"user_name,omitempty"`

	// user tags
	UserTags []string `json:"user_tags"`

	// user uuid
	UserUUID string `json:"user_uuid,omitempty"`
}

// Validate validates this falconx submission v1
func (m *FalconxSubmissionV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSandbox(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FalconxSubmissionV1) validateSandbox(formats strfmt.Registry) error {
	if swag.IsZero(m.Sandbox) { // not required
		return nil
	}

	for i := 0; i < len(m.Sandbox); i++ {
		if swag.IsZero(m.Sandbox[i]) { // not required
			continue
		}

		if m.Sandbox[i] != nil {
			if err := m.Sandbox[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sandbox" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sandbox" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this falconx submission v1 based on the context it is used
func (m *FalconxSubmissionV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSandbox(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FalconxSubmissionV1) contextValidateSandbox(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Sandbox); i++ {

		if m.Sandbox[i] != nil {
			if err := m.Sandbox[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sandbox" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sandbox" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FalconxSubmissionV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FalconxSubmissionV1) UnmarshalBinary(b []byte) error {
	var res FalconxSubmissionV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
