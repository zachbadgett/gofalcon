// Code generated by go-swagger; DO NOT EDIT.

package sensor_update_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new sensor update policies API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sensor update policies API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateSensorUpdatePolicies(params *CreateSensorUpdatePoliciesParams, opts ...ClientOption) (*CreateSensorUpdatePoliciesCreated, error)

	CreateSensorUpdatePoliciesV2(params *CreateSensorUpdatePoliciesV2Params, opts ...ClientOption) (*CreateSensorUpdatePoliciesV2Created, error)

	DeleteSensorUpdatePolicies(params *DeleteSensorUpdatePoliciesParams, opts ...ClientOption) (*DeleteSensorUpdatePoliciesOK, error)

	GetSensorUpdatePolicies(params *GetSensorUpdatePoliciesParams, opts ...ClientOption) (*GetSensorUpdatePoliciesOK, error)

	GetSensorUpdatePoliciesV2(params *GetSensorUpdatePoliciesV2Params, opts ...ClientOption) (*GetSensorUpdatePoliciesV2OK, error)

	PerformSensorUpdatePoliciesAction(params *PerformSensorUpdatePoliciesActionParams, opts ...ClientOption) (*PerformSensorUpdatePoliciesActionOK, error)

	QueryCombinedSensorUpdateBuilds(params *QueryCombinedSensorUpdateBuildsParams, opts ...ClientOption) (*QueryCombinedSensorUpdateBuildsOK, error)

	QueryCombinedSensorUpdateKernels(params *QueryCombinedSensorUpdateKernelsParams, opts ...ClientOption) (*QueryCombinedSensorUpdateKernelsOK, error)

	QueryCombinedSensorUpdatePolicies(params *QueryCombinedSensorUpdatePoliciesParams, opts ...ClientOption) (*QueryCombinedSensorUpdatePoliciesOK, error)

	QueryCombinedSensorUpdatePoliciesV2(params *QueryCombinedSensorUpdatePoliciesV2Params, opts ...ClientOption) (*QueryCombinedSensorUpdatePoliciesV2OK, error)

	QueryCombinedSensorUpdatePolicyMembers(params *QueryCombinedSensorUpdatePolicyMembersParams, opts ...ClientOption) (*QueryCombinedSensorUpdatePolicyMembersOK, error)

	QuerySensorUpdateKernelsDistinct(params *QuerySensorUpdateKernelsDistinctParams, opts ...ClientOption) (*QuerySensorUpdateKernelsDistinctOK, error)

	QuerySensorUpdatePolicies(params *QuerySensorUpdatePoliciesParams, opts ...ClientOption) (*QuerySensorUpdatePoliciesOK, error)

	QuerySensorUpdatePolicyMembers(params *QuerySensorUpdatePolicyMembersParams, opts ...ClientOption) (*QuerySensorUpdatePolicyMembersOK, error)

	RevealUninstallToken(params *RevealUninstallTokenParams, opts ...ClientOption) (*RevealUninstallTokenOK, error)

	SetSensorUpdatePoliciesPrecedence(params *SetSensorUpdatePoliciesPrecedenceParams, opts ...ClientOption) (*SetSensorUpdatePoliciesPrecedenceOK, error)

	UpdateSensorUpdatePolicies(params *UpdateSensorUpdatePoliciesParams, opts ...ClientOption) (*UpdateSensorUpdatePoliciesOK, error)

	UpdateSensorUpdatePoliciesV2(params *UpdateSensorUpdatePoliciesV2Params, opts ...ClientOption) (*UpdateSensorUpdatePoliciesV2OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateSensorUpdatePolicies creates sensor update policies by specifying details about the policy to create
*/
func (a *Client) CreateSensorUpdatePolicies(params *CreateSensorUpdatePoliciesParams, opts ...ClientOption) (*CreateSensorUpdatePoliciesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSensorUpdatePoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createSensorUpdatePolicies",
		Method:             "POST",
		PathPattern:        "/policy/entities/sensor-update/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSensorUpdatePoliciesReader{formats: a.formats},
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
	success, ok := result.(*CreateSensorUpdatePoliciesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createSensorUpdatePolicies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateSensorUpdatePoliciesV2 creates sensor update policies by specifying details about the policy to create with additional support for uninstall protection
*/
func (a *Client) CreateSensorUpdatePoliciesV2(params *CreateSensorUpdatePoliciesV2Params, opts ...ClientOption) (*CreateSensorUpdatePoliciesV2Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSensorUpdatePoliciesV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "createSensorUpdatePoliciesV2",
		Method:             "POST",
		PathPattern:        "/policy/entities/sensor-update/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSensorUpdatePoliciesV2Reader{formats: a.formats},
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
	success, ok := result.(*CreateSensorUpdatePoliciesV2Created)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createSensorUpdatePoliciesV2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteSensorUpdatePolicies deletes a set of sensor update policies by specifying their i ds
*/
func (a *Client) DeleteSensorUpdatePolicies(params *DeleteSensorUpdatePoliciesParams, opts ...ClientOption) (*DeleteSensorUpdatePoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSensorUpdatePoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteSensorUpdatePolicies",
		Method:             "DELETE",
		PathPattern:        "/policy/entities/sensor-update/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSensorUpdatePoliciesReader{formats: a.formats},
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
	success, ok := result.(*DeleteSensorUpdatePoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteSensorUpdatePolicies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSensorUpdatePolicies retrieves a set of sensor update policies by specifying their i ds
*/
func (a *Client) GetSensorUpdatePolicies(params *GetSensorUpdatePoliciesParams, opts ...ClientOption) (*GetSensorUpdatePoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSensorUpdatePoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSensorUpdatePolicies",
		Method:             "GET",
		PathPattern:        "/policy/entities/sensor-update/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSensorUpdatePoliciesReader{formats: a.formats},
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
	success, ok := result.(*GetSensorUpdatePoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSensorUpdatePolicies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSensorUpdatePoliciesV2 retrieves a set of sensor update policies with additional support for uninstall protection by specifying their i ds
*/
func (a *Client) GetSensorUpdatePoliciesV2(params *GetSensorUpdatePoliciesV2Params, opts ...ClientOption) (*GetSensorUpdatePoliciesV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSensorUpdatePoliciesV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSensorUpdatePoliciesV2",
		Method:             "GET",
		PathPattern:        "/policy/entities/sensor-update/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSensorUpdatePoliciesV2Reader{formats: a.formats},
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
	success, ok := result.(*GetSensorUpdatePoliciesV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSensorUpdatePoliciesV2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PerformSensorUpdatePoliciesAction performs the specified action on the sensor update policies specified in the request
*/
func (a *Client) PerformSensorUpdatePoliciesAction(params *PerformSensorUpdatePoliciesActionParams, opts ...ClientOption) (*PerformSensorUpdatePoliciesActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPerformSensorUpdatePoliciesActionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "performSensorUpdatePoliciesAction",
		Method:             "POST",
		PathPattern:        "/policy/entities/sensor-update-actions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PerformSensorUpdatePoliciesActionReader{formats: a.formats},
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
	success, ok := result.(*PerformSensorUpdatePoliciesActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PerformSensorUpdatePoliciesActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryCombinedSensorUpdateBuilds retrieves available builds for use with sensor update policies
*/
func (a *Client) QueryCombinedSensorUpdateBuilds(params *QueryCombinedSensorUpdateBuildsParams, opts ...ClientOption) (*QueryCombinedSensorUpdateBuildsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryCombinedSensorUpdateBuildsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "queryCombinedSensorUpdateBuilds",
		Method:             "GET",
		PathPattern:        "/policy/combined/sensor-update-builds/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryCombinedSensorUpdateBuildsReader{formats: a.formats},
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
	success, ok := result.(*QueryCombinedSensorUpdateBuildsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryCombinedSensorUpdateBuildsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryCombinedSensorUpdateKernels retrieves kernel compatibility info for sensor update builds
*/
func (a *Client) QueryCombinedSensorUpdateKernels(params *QueryCombinedSensorUpdateKernelsParams, opts ...ClientOption) (*QueryCombinedSensorUpdateKernelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryCombinedSensorUpdateKernelsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "queryCombinedSensorUpdateKernels",
		Method:             "GET",
		PathPattern:        "/policy/combined/sensor-update-kernels/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryCombinedSensorUpdateKernelsReader{formats: a.formats},
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
	success, ok := result.(*QueryCombinedSensorUpdateKernelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryCombinedSensorUpdateKernelsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryCombinedSensorUpdatePolicies searches for sensor update policies in your environment by providing an f q l filter and paging details returns a set of sensor update policies which match the filter criteria
*/
func (a *Client) QueryCombinedSensorUpdatePolicies(params *QueryCombinedSensorUpdatePoliciesParams, opts ...ClientOption) (*QueryCombinedSensorUpdatePoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryCombinedSensorUpdatePoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "queryCombinedSensorUpdatePolicies",
		Method:             "GET",
		PathPattern:        "/policy/combined/sensor-update/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryCombinedSensorUpdatePoliciesReader{formats: a.formats},
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
	success, ok := result.(*QueryCombinedSensorUpdatePoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryCombinedSensorUpdatePoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryCombinedSensorUpdatePoliciesV2 searches for sensor update policies with additional support for uninstall protection in your environment by providing an f q l filter and paging details returns a set of sensor update policies which match the filter criteria
*/
func (a *Client) QueryCombinedSensorUpdatePoliciesV2(params *QueryCombinedSensorUpdatePoliciesV2Params, opts ...ClientOption) (*QueryCombinedSensorUpdatePoliciesV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryCombinedSensorUpdatePoliciesV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "queryCombinedSensorUpdatePoliciesV2",
		Method:             "GET",
		PathPattern:        "/policy/combined/sensor-update/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryCombinedSensorUpdatePoliciesV2Reader{formats: a.formats},
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
	success, ok := result.(*QueryCombinedSensorUpdatePoliciesV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryCombinedSensorUpdatePoliciesV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryCombinedSensorUpdatePolicyMembers searches for members of a sensor update policy in your environment by providing an f q l filter and paging details returns a set of host details which match the filter criteria
*/
func (a *Client) QueryCombinedSensorUpdatePolicyMembers(params *QueryCombinedSensorUpdatePolicyMembersParams, opts ...ClientOption) (*QueryCombinedSensorUpdatePolicyMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryCombinedSensorUpdatePolicyMembersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "queryCombinedSensorUpdatePolicyMembers",
		Method:             "GET",
		PathPattern:        "/policy/combined/sensor-update-members/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryCombinedSensorUpdatePolicyMembersReader{formats: a.formats},
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
	success, ok := result.(*QueryCombinedSensorUpdatePolicyMembersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryCombinedSensorUpdatePolicyMembersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QuerySensorUpdateKernelsDistinct retrieves kernel compatibility info for sensor update builds
*/
func (a *Client) QuerySensorUpdateKernelsDistinct(params *QuerySensorUpdateKernelsDistinctParams, opts ...ClientOption) (*QuerySensorUpdateKernelsDistinctOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuerySensorUpdateKernelsDistinctParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "querySensorUpdateKernelsDistinct",
		Method:             "GET",
		PathPattern:        "/policy/queries/sensor-update-kernels/{distinct_field}/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QuerySensorUpdateKernelsDistinctReader{formats: a.formats},
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
	success, ok := result.(*QuerySensorUpdateKernelsDistinctOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QuerySensorUpdateKernelsDistinctDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QuerySensorUpdatePolicies searches for sensor update policies in your environment by providing an f q l filter and paging details returns a set of sensor update policy i ds which match the filter criteria
*/
func (a *Client) QuerySensorUpdatePolicies(params *QuerySensorUpdatePoliciesParams, opts ...ClientOption) (*QuerySensorUpdatePoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuerySensorUpdatePoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "querySensorUpdatePolicies",
		Method:             "GET",
		PathPattern:        "/policy/queries/sensor-update/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QuerySensorUpdatePoliciesReader{formats: a.formats},
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
	success, ok := result.(*QuerySensorUpdatePoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QuerySensorUpdatePoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QuerySensorUpdatePolicyMembers searches for members of a sensor update policy in your environment by providing an f q l filter and paging details returns a set of agent i ds which match the filter criteria
*/
func (a *Client) QuerySensorUpdatePolicyMembers(params *QuerySensorUpdatePolicyMembersParams, opts ...ClientOption) (*QuerySensorUpdatePolicyMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuerySensorUpdatePolicyMembersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "querySensorUpdatePolicyMembers",
		Method:             "GET",
		PathPattern:        "/policy/queries/sensor-update-members/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QuerySensorUpdatePolicyMembersReader{formats: a.formats},
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
	success, ok := result.(*QuerySensorUpdatePolicyMembersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QuerySensorUpdatePolicyMembersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RevealUninstallToken reveals an uninstall token for a specific device to retrieve the bulk maintenance token pass the value m a i n t e n a n c e as the value for device id
*/
func (a *Client) RevealUninstallToken(params *RevealUninstallTokenParams, opts ...ClientOption) (*RevealUninstallTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRevealUninstallTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "revealUninstallToken",
		Method:             "POST",
		PathPattern:        "/policy/combined/reveal-uninstall-token/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RevealUninstallTokenReader{formats: a.formats},
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
	success, ok := result.(*RevealUninstallTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RevealUninstallTokenDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SetSensorUpdatePoliciesPrecedence sets the precedence of sensor update policies based on the order of i ds specified in the request the first ID specified will have the highest precedence and the last ID specified will have the lowest you must specify all non default policies for a platform when updating precedence
*/
func (a *Client) SetSensorUpdatePoliciesPrecedence(params *SetSensorUpdatePoliciesPrecedenceParams, opts ...ClientOption) (*SetSensorUpdatePoliciesPrecedenceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetSensorUpdatePoliciesPrecedenceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "setSensorUpdatePoliciesPrecedence",
		Method:             "POST",
		PathPattern:        "/policy/entities/sensor-update-precedence/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetSensorUpdatePoliciesPrecedenceReader{formats: a.formats},
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
	success, ok := result.(*SetSensorUpdatePoliciesPrecedenceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SetSensorUpdatePoliciesPrecedenceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateSensorUpdatePolicies updates sensor update policies by specifying the ID of the policy and details to update
*/
func (a *Client) UpdateSensorUpdatePolicies(params *UpdateSensorUpdatePoliciesParams, opts ...ClientOption) (*UpdateSensorUpdatePoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSensorUpdatePoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateSensorUpdatePolicies",
		Method:             "PATCH",
		PathPattern:        "/policy/entities/sensor-update/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSensorUpdatePoliciesReader{formats: a.formats},
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
	success, ok := result.(*UpdateSensorUpdatePoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateSensorUpdatePoliciesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateSensorUpdatePoliciesV2 updates sensor update policies by specifying the ID of the policy and details to update with additional support for uninstall protection
*/
func (a *Client) UpdateSensorUpdatePoliciesV2(params *UpdateSensorUpdatePoliciesV2Params, opts ...ClientOption) (*UpdateSensorUpdatePoliciesV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSensorUpdatePoliciesV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateSensorUpdatePoliciesV2",
		Method:             "PATCH",
		PathPattern:        "/policy/entities/sensor-update/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSensorUpdatePoliciesV2Reader{formats: a.formats},
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
	success, ok := result.(*UpdateSensorUpdatePoliciesV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateSensorUpdatePoliciesV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
