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

// FwmgrFirewallRuleV1 fwmgr firewall rule v1
//
// swagger:model fwmgr.firewall.RuleV1
type FwmgrFirewallRuleV1 struct {

	// action
	// Required: true
	Action *string `json:"action"`

	// address family
	// Required: true
	AddressFamily *string `json:"address_family"`

	// created by
	// Required: true
	CreatedBy *string `json:"created_by"`

	// created on
	// Required: true
	CreatedOn *string `json:"created_on"`

	// customer id
	CustomerID string `json:"customer_id,omitempty"`

	// deleted
	// Required: true
	Deleted *bool `json:"deleted"`

	// description
	// Required: true
	Description *string `json:"description"`

	// direction
	// Required: true
	Direction *string `json:"direction"`

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// family
	// Required: true
	Family *string `json:"family"`

	// fields
	// Required: true
	Fields []*FwmgrFirewallFieldValue `json:"fields"`

	// icmp
	// Required: true
	Icmp *FwmgrFirewallICMP `json:"icmp"`

	// id
	// Required: true
	ID *string `json:"id"`

	// local address
	// Required: true
	LocalAddress []*FwmgrFirewallAddressRange `json:"local_address"`

	// local port
	// Required: true
	LocalPort []*FwmgrFirewallPortRange `json:"local_port"`

	// modified by
	ModifiedBy string `json:"modified_by,omitempty"`

	// modified on
	ModifiedOn string `json:"modified_on,omitempty"`

	// monitor
	// Required: true
	Monitor *FwmgrFirewallMonitoring `json:"monitor"`

	// name
	// Required: true
	Name *string `json:"name"`

	// platform ids
	// Required: true
	PlatformIds []string `json:"platform_ids"`

	// protocol
	// Required: true
	Protocol *string `json:"protocol"`

	// remote address
	// Required: true
	RemoteAddress []*FwmgrFirewallAddressRange `json:"remote_address"`

	// remote port
	// Required: true
	RemotePort []*FwmgrFirewallPortRange `json:"remote_port"`

	// rule group
	// Required: true
	RuleGroup *FwmgrFirewallRuleGroupSummaryV1 `json:"rule_group"`

	// version
	// Required: true
	Version *int64 `json:"version"`
}

// Validate validates this fwmgr firewall rule v1
func (m *FwmgrFirewallRuleV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddressFamily(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFamily(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIcmp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonitor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemotePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FwmgrFirewallRuleV1) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrFirewallRuleV1) validateAddressFamily(formats strfmt.Registry) error {

	if err := validate.Required("address_family", "body", m.AddressFamily); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrFirewallRuleV1) validateCreatedBy(formats strfmt.Registry) error {

	if err := validate.Required("created_by", "body", m.CreatedBy); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrFirewallRuleV1) validateCreatedOn(formats strfmt.Registry) error {

	if err := validate.Required("created_on", "body", m.CreatedOn); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrFirewallRuleV1) validateDeleted(formats strfmt.Registry) error {

	if err := validate.Required("deleted", "body", m.Deleted); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrFirewallRuleV1) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrFirewallRuleV1) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrFirewallRuleV1) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrFirewallRuleV1) validateFamily(formats strfmt.Registry) error {

	if err := validate.Required("family", "body", m.Family); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrFirewallRuleV1) validateFields(formats strfmt.Registry) error {

	if err := validate.Required("fields", "body", m.Fields); err != nil {
		return err
	}

	for i := 0; i < len(m.Fields); i++ {
		if swag.IsZero(m.Fields[i]) { // not required
			continue
		}

		if m.Fields[i] != nil {
			if err := m.Fields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FwmgrFirewallRuleV1) validateIcmp(formats strfmt.Registry) error {

	if err := validate.Required("icmp", "body", m.Icmp); err != nil {
		return err
	}

	if m.Icmp != nil {
		if err := m.Icmp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("icmp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("icmp")
			}
			return err
		}
	}

	return nil
}

func (m *FwmgrFirewallRuleV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrFirewallRuleV1) validateLocalAddress(formats strfmt.Registry) error {

	if err := validate.Required("local_address", "body", m.LocalAddress); err != nil {
		return err
	}

	for i := 0; i < len(m.LocalAddress); i++ {
		if swag.IsZero(m.LocalAddress[i]) { // not required
			continue
		}

		if m.LocalAddress[i] != nil {
			if err := m.LocalAddress[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("local_address" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("local_address" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FwmgrFirewallRuleV1) validateLocalPort(formats strfmt.Registry) error {

	if err := validate.Required("local_port", "body", m.LocalPort); err != nil {
		return err
	}

	for i := 0; i < len(m.LocalPort); i++ {
		if swag.IsZero(m.LocalPort[i]) { // not required
			continue
		}

		if m.LocalPort[i] != nil {
			if err := m.LocalPort[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("local_port" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("local_port" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FwmgrFirewallRuleV1) validateMonitor(formats strfmt.Registry) error {

	if err := validate.Required("monitor", "body", m.Monitor); err != nil {
		return err
	}

	if m.Monitor != nil {
		if err := m.Monitor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monitor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("monitor")
			}
			return err
		}
	}

	return nil
}

func (m *FwmgrFirewallRuleV1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrFirewallRuleV1) validatePlatformIds(formats strfmt.Registry) error {

	if err := validate.Required("platform_ids", "body", m.PlatformIds); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrFirewallRuleV1) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", m.Protocol); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrFirewallRuleV1) validateRemoteAddress(formats strfmt.Registry) error {

	if err := validate.Required("remote_address", "body", m.RemoteAddress); err != nil {
		return err
	}

	for i := 0; i < len(m.RemoteAddress); i++ {
		if swag.IsZero(m.RemoteAddress[i]) { // not required
			continue
		}

		if m.RemoteAddress[i] != nil {
			if err := m.RemoteAddress[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("remote_address" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("remote_address" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FwmgrFirewallRuleV1) validateRemotePort(formats strfmt.Registry) error {

	if err := validate.Required("remote_port", "body", m.RemotePort); err != nil {
		return err
	}

	for i := 0; i < len(m.RemotePort); i++ {
		if swag.IsZero(m.RemotePort[i]) { // not required
			continue
		}

		if m.RemotePort[i] != nil {
			if err := m.RemotePort[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("remote_port" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("remote_port" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FwmgrFirewallRuleV1) validateRuleGroup(formats strfmt.Registry) error {

	if err := validate.Required("rule_group", "body", m.RuleGroup); err != nil {
		return err
	}

	if m.RuleGroup != nil {
		if err := m.RuleGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rule_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rule_group")
			}
			return err
		}
	}

	return nil
}

func (m *FwmgrFirewallRuleV1) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this fwmgr firewall rule v1 based on the context it is used
func (m *FwmgrFirewallRuleV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIcmp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocalAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocalPort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMonitor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemoteAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemotePort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRuleGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FwmgrFirewallRuleV1) contextValidateFields(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Fields); i++ {

		if m.Fields[i] != nil {
			if err := m.Fields[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FwmgrFirewallRuleV1) contextValidateIcmp(ctx context.Context, formats strfmt.Registry) error {

	if m.Icmp != nil {
		if err := m.Icmp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("icmp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("icmp")
			}
			return err
		}
	}

	return nil
}

func (m *FwmgrFirewallRuleV1) contextValidateLocalAddress(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LocalAddress); i++ {

		if m.LocalAddress[i] != nil {
			if err := m.LocalAddress[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("local_address" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("local_address" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FwmgrFirewallRuleV1) contextValidateLocalPort(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LocalPort); i++ {

		if m.LocalPort[i] != nil {
			if err := m.LocalPort[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("local_port" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("local_port" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FwmgrFirewallRuleV1) contextValidateMonitor(ctx context.Context, formats strfmt.Registry) error {

	if m.Monitor != nil {
		if err := m.Monitor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monitor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("monitor")
			}
			return err
		}
	}

	return nil
}

func (m *FwmgrFirewallRuleV1) contextValidateRemoteAddress(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RemoteAddress); i++ {

		if m.RemoteAddress[i] != nil {
			if err := m.RemoteAddress[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("remote_address" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("remote_address" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FwmgrFirewallRuleV1) contextValidateRemotePort(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RemotePort); i++ {

		if m.RemotePort[i] != nil {
			if err := m.RemotePort[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("remote_port" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("remote_port" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FwmgrFirewallRuleV1) contextValidateRuleGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.RuleGroup != nil {
		if err := m.RuleGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rule_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rule_group")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FwmgrFirewallRuleV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FwmgrFirewallRuleV1) UnmarshalBinary(b []byte) error {
	var res FwmgrFirewallRuleV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
