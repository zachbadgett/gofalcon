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

// DomainIncident domain incident
//
// swagger:model domain.Incident
type DomainIncident struct {

	// assigned to
	AssignedTo string `json:"assigned_to,omitempty"`

	// assigned to name
	AssignedToName string `json:"assigned_to_name,omitempty"`

	// cid
	// Required: true
	Cid *string `json:"cid"`

	// created
	// Required: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created"`

	// description
	Description string `json:"description,omitempty"`

	// email state
	EmailState string `json:"email_state,omitempty"`

	// end
	// Required: true
	// Format: date-time
	End *strfmt.DateTime `json:"end"`

	// events histogram
	EventsHistogram []*DomainEventHistogram `json:"events_histogram"`

	// fine score
	// Required: true
	FineScore *int32 `json:"fine_score"`

	// host ids
	// Required: true
	HostIds []string `json:"host_ids"`

	// hosts
	Hosts []*DetectsDeviceDetailIndexed `json:"hosts"`

	// incident id
	// Required: true
	IncidentID *string `json:"incident_id"`

	// incident type
	IncidentType int64 `json:"incident_type,omitempty"`

	// lm host ids
	LmHostIds []string `json:"lm_host_ids"`

	// lm hosts capped
	LmHostsCapped bool `json:"lm_hosts_capped,omitempty"`

	// lm types
	LmTypes int64 `json:"lm_types,omitempty"`

	// lmra host ids
	LmraHostIds []string `json:"lmra_host_ids"`

	// lmra hosts capped
	LmraHostsCapped bool `json:"lmra_hosts_capped,omitempty"`

	// modified timestamp
	// Format: date-time
	ModifiedTimestamp strfmt.DateTime `json:"modified_timestamp,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// objectives
	Objectives []string `json:"objectives"`

	// start
	// Required: true
	// Format: date-time
	Start *strfmt.DateTime `json:"start"`

	// state
	// Required: true
	State *string `json:"state"`

	// status
	Status int32 `json:"status,omitempty"`

	// tactics
	Tactics []string `json:"tactics"`

	// tags
	Tags []string `json:"tags"`

	// techniques
	Techniques []string `json:"techniques"`

	// users
	Users []string `json:"users"`

	// visibility
	Visibility int32 `json:"visibility,omitempty"`
}

// Validate validates this domain incident
func (m *DomainIncident) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventsHistogram(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFineScore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncidentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainIncident) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	return nil
}

func (m *DomainIncident) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", m.Created); err != nil {
		return err
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainIncident) validateEnd(formats strfmt.Registry) error {

	if err := validate.Required("end", "body", m.End); err != nil {
		return err
	}

	if err := validate.FormatOf("end", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainIncident) validateEventsHistogram(formats strfmt.Registry) error {
	if swag.IsZero(m.EventsHistogram) { // not required
		return nil
	}

	for i := 0; i < len(m.EventsHistogram); i++ {
		if swag.IsZero(m.EventsHistogram[i]) { // not required
			continue
		}

		if m.EventsHistogram[i] != nil {
			if err := m.EventsHistogram[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("events_histogram" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("events_histogram" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainIncident) validateFineScore(formats strfmt.Registry) error {

	if err := validate.Required("fine_score", "body", m.FineScore); err != nil {
		return err
	}

	return nil
}

func (m *DomainIncident) validateHostIds(formats strfmt.Registry) error {

	if err := validate.Required("host_ids", "body", m.HostIds); err != nil {
		return err
	}

	return nil
}

func (m *DomainIncident) validateHosts(formats strfmt.Registry) error {
	if swag.IsZero(m.Hosts) { // not required
		return nil
	}

	for i := 0; i < len(m.Hosts); i++ {
		if swag.IsZero(m.Hosts[i]) { // not required
			continue
		}

		if m.Hosts[i] != nil {
			if err := m.Hosts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hosts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainIncident) validateIncidentID(formats strfmt.Registry) error {

	if err := validate.Required("incident_id", "body", m.IncidentID); err != nil {
		return err
	}

	return nil
}

func (m *DomainIncident) validateModifiedTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.ModifiedTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_timestamp", "body", "date-time", m.ModifiedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainIncident) validateStart(formats strfmt.Registry) error {

	if err := validate.Required("start", "body", m.Start); err != nil {
		return err
	}

	if err := validate.FormatOf("start", "body", "date-time", m.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainIncident) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this domain incident based on the context it is used
func (m *DomainIncident) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEventsHistogram(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHosts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainIncident) contextValidateEventsHistogram(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EventsHistogram); i++ {

		if m.EventsHistogram[i] != nil {
			if err := m.EventsHistogram[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("events_histogram" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("events_histogram" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainIncident) contextValidateHosts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Hosts); i++ {

		if m.Hosts[i] != nil {
			if err := m.Hosts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hosts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainIncident) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainIncident) UnmarshalBinary(b []byte) error {
	var res DomainIncident
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
