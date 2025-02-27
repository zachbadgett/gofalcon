// Code generated by go-swagger; DO NOT EDIT.

package ml_exclusions

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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewUpdateMLExclusionsV1Params creates a new UpdateMLExclusionsV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateMLExclusionsV1Params() *UpdateMLExclusionsV1Params {
	return &UpdateMLExclusionsV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMLExclusionsV1ParamsWithTimeout creates a new UpdateMLExclusionsV1Params object
// with the ability to set a timeout on a request.
func NewUpdateMLExclusionsV1ParamsWithTimeout(timeout time.Duration) *UpdateMLExclusionsV1Params {
	return &UpdateMLExclusionsV1Params{
		timeout: timeout,
	}
}

// NewUpdateMLExclusionsV1ParamsWithContext creates a new UpdateMLExclusionsV1Params object
// with the ability to set a context for a request.
func NewUpdateMLExclusionsV1ParamsWithContext(ctx context.Context) *UpdateMLExclusionsV1Params {
	return &UpdateMLExclusionsV1Params{
		Context: ctx,
	}
}

// NewUpdateMLExclusionsV1ParamsWithHTTPClient creates a new UpdateMLExclusionsV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateMLExclusionsV1ParamsWithHTTPClient(client *http.Client) *UpdateMLExclusionsV1Params {
	return &UpdateMLExclusionsV1Params{
		HTTPClient: client,
	}
}

/*
UpdateMLExclusionsV1Params contains all the parameters to send to the API endpoint

	for the update m l exclusions v1 operation.

	Typically these are written to a http.Request.
*/
type UpdateMLExclusionsV1Params struct {

	// Body.
	Body *models.RequestsSvExclusionUpdateReqV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update m l exclusions v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMLExclusionsV1Params) WithDefaults() *UpdateMLExclusionsV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update m l exclusions v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMLExclusionsV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update m l exclusions v1 params
func (o *UpdateMLExclusionsV1Params) WithTimeout(timeout time.Duration) *UpdateMLExclusionsV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update m l exclusions v1 params
func (o *UpdateMLExclusionsV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update m l exclusions v1 params
func (o *UpdateMLExclusionsV1Params) WithContext(ctx context.Context) *UpdateMLExclusionsV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update m l exclusions v1 params
func (o *UpdateMLExclusionsV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update m l exclusions v1 params
func (o *UpdateMLExclusionsV1Params) WithHTTPClient(client *http.Client) *UpdateMLExclusionsV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update m l exclusions v1 params
func (o *UpdateMLExclusionsV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update m l exclusions v1 params
func (o *UpdateMLExclusionsV1Params) WithBody(body *models.RequestsSvExclusionUpdateReqV1) *UpdateMLExclusionsV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update m l exclusions v1 params
func (o *UpdateMLExclusionsV1Params) SetBody(body *models.RequestsSvExclusionUpdateReqV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMLExclusionsV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
