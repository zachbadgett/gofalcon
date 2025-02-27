// Code generated by go-swagger; DO NOT EDIT.

package zero_trust_assessment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new zero trust assessment API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for zero trust assessment API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAssessmentV1(params *GetAssessmentV1Params, opts ...ClientOption) (*GetAssessmentV1OK, error)

	GetComplianceV1(params *GetComplianceV1Params, opts ...ClientOption) (*GetComplianceV1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetAssessmentV1 gets zero trust assessment data for one or more hosts by providing agent i ds a ID and a customer ID c ID
*/
func (a *Client) GetAssessmentV1(params *GetAssessmentV1Params, opts ...ClientOption) (*GetAssessmentV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAssessmentV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAssessmentV1",
		Method:             "GET",
		PathPattern:        "/zero-trust-assessment/entities/assessments/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAssessmentV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAssessmentV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAssessmentV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetComplianceV1 gets the zero trust assessment compliance report for one customer ID c ID
*/
func (a *Client) GetComplianceV1(params *GetComplianceV1Params, opts ...ClientOption) (*GetComplianceV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComplianceV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getComplianceV1",
		Method:             "GET",
		PathPattern:        "/zero-trust-assessment/entities/audit/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetComplianceV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetComplianceV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetComplianceV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
