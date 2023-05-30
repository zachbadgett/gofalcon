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

// PublicACL public ACL
//
// swagger:model public.ACL
type PublicACL struct {

	// acl permission change
	ACLPermissionChange []*PublicBasic `json:"acl_permission_change"`

	// entity
	Entity string `json:"entity,omitempty"`

	// entity id
	EntityID string `json:"entity_id,omitempty"`

	// entity name
	EntityName string `json:"entity_name,omitempty"`
}

// Validate validates this public ACL
func (m *PublicACL) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateACLPermissionChange(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicACL) validateACLPermissionChange(formats strfmt.Registry) error {
	if swag.IsZero(m.ACLPermissionChange) { // not required
		return nil
	}

	for i := 0; i < len(m.ACLPermissionChange); i++ {
		if swag.IsZero(m.ACLPermissionChange[i]) { // not required
			continue
		}

		if m.ACLPermissionChange[i] != nil {
			if err := m.ACLPermissionChange[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("acl_permission_change" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("acl_permission_change" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this public ACL based on the context it is used
func (m *PublicACL) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateACLPermissionChange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicACL) contextValidateACLPermissionChange(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ACLPermissionChange); i++ {

		if m.ACLPermissionChange[i] != nil {
			if err := m.ACLPermissionChange[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("acl_permission_change" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("acl_permission_change" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicACL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicACL) UnmarshalBinary(b []byte) error {
	var res PublicACL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
