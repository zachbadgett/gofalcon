// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_protection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetHelmValuesYamlParams creates a new GetHelmValuesYamlParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHelmValuesYamlParams() *GetHelmValuesYamlParams {
	return &GetHelmValuesYamlParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHelmValuesYamlParamsWithTimeout creates a new GetHelmValuesYamlParams object
// with the ability to set a timeout on a request.
func NewGetHelmValuesYamlParamsWithTimeout(timeout time.Duration) *GetHelmValuesYamlParams {
	return &GetHelmValuesYamlParams{
		timeout: timeout,
	}
}

// NewGetHelmValuesYamlParamsWithContext creates a new GetHelmValuesYamlParams object
// with the ability to set a context for a request.
func NewGetHelmValuesYamlParamsWithContext(ctx context.Context) *GetHelmValuesYamlParams {
	return &GetHelmValuesYamlParams{
		Context: ctx,
	}
}

// NewGetHelmValuesYamlParamsWithHTTPClient creates a new GetHelmValuesYamlParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetHelmValuesYamlParamsWithHTTPClient(client *http.Client) *GetHelmValuesYamlParams {
	return &GetHelmValuesYamlParams{
		HTTPClient: client,
	}
}

/*
GetHelmValuesYamlParams contains all the parameters to send to the API endpoint

	for the get helm values yaml operation.

	Typically these are written to a http.Request.
*/
type GetHelmValuesYamlParams struct {

	/* ClusterName.

	   Cluster name. For EKS it will be cluster ARN.
	*/
	ClusterName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get helm values yaml params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHelmValuesYamlParams) WithDefaults() *GetHelmValuesYamlParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get helm values yaml params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHelmValuesYamlParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get helm values yaml params
func (o *GetHelmValuesYamlParams) WithTimeout(timeout time.Duration) *GetHelmValuesYamlParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get helm values yaml params
func (o *GetHelmValuesYamlParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get helm values yaml params
func (o *GetHelmValuesYamlParams) WithContext(ctx context.Context) *GetHelmValuesYamlParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get helm values yaml params
func (o *GetHelmValuesYamlParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get helm values yaml params
func (o *GetHelmValuesYamlParams) WithHTTPClient(client *http.Client) *GetHelmValuesYamlParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get helm values yaml params
func (o *GetHelmValuesYamlParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterName adds the clusterName to the get helm values yaml params
func (o *GetHelmValuesYamlParams) WithClusterName(clusterName string) *GetHelmValuesYamlParams {
	o.SetClusterName(clusterName)
	return o
}

// SetClusterName adds the clusterName to the get helm values yaml params
func (o *GetHelmValuesYamlParams) SetClusterName(clusterName string) {
	o.ClusterName = clusterName
}

// WriteToRequest writes these params to a swagger request
func (o *GetHelmValuesYamlParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param cluster_name
	qrClusterName := o.ClusterName
	qClusterName := qrClusterName
	if qClusterName != "" {

		if err := r.SetQueryParam("cluster_name", qClusterName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
