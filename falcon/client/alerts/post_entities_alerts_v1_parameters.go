// Code generated by go-swagger; DO NOT EDIT.

package alerts

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

// NewPostEntitiesAlertsV1Params creates a new PostEntitiesAlertsV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostEntitiesAlertsV1Params() *PostEntitiesAlertsV1Params {
	return &PostEntitiesAlertsV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostEntitiesAlertsV1ParamsWithTimeout creates a new PostEntitiesAlertsV1Params object
// with the ability to set a timeout on a request.
func NewPostEntitiesAlertsV1ParamsWithTimeout(timeout time.Duration) *PostEntitiesAlertsV1Params {
	return &PostEntitiesAlertsV1Params{
		timeout: timeout,
	}
}

// NewPostEntitiesAlertsV1ParamsWithContext creates a new PostEntitiesAlertsV1Params object
// with the ability to set a context for a request.
func NewPostEntitiesAlertsV1ParamsWithContext(ctx context.Context) *PostEntitiesAlertsV1Params {
	return &PostEntitiesAlertsV1Params{
		Context: ctx,
	}
}

// NewPostEntitiesAlertsV1ParamsWithHTTPClient creates a new PostEntitiesAlertsV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewPostEntitiesAlertsV1ParamsWithHTTPClient(client *http.Client) *PostEntitiesAlertsV1Params {
	return &PostEntitiesAlertsV1Params{
		HTTPClient: client,
	}
}

/*
PostEntitiesAlertsV1Params contains all the parameters to send to the API endpoint

	for the post entities alerts v1 operation.

	Typically these are written to a http.Request.
*/
type PostEntitiesAlertsV1Params struct {

	// Body.
	Body *models.DetectsapiPostEntitiesInvestigatablesV1Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post entities alerts v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostEntitiesAlertsV1Params) WithDefaults() *PostEntitiesAlertsV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post entities alerts v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostEntitiesAlertsV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post entities alerts v1 params
func (o *PostEntitiesAlertsV1Params) WithTimeout(timeout time.Duration) *PostEntitiesAlertsV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post entities alerts v1 params
func (o *PostEntitiesAlertsV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post entities alerts v1 params
func (o *PostEntitiesAlertsV1Params) WithContext(ctx context.Context) *PostEntitiesAlertsV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post entities alerts v1 params
func (o *PostEntitiesAlertsV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post entities alerts v1 params
func (o *PostEntitiesAlertsV1Params) WithHTTPClient(client *http.Client) *PostEntitiesAlertsV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post entities alerts v1 params
func (o *PostEntitiesAlertsV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post entities alerts v1 params
func (o *PostEntitiesAlertsV1Params) WithBody(body *models.DetectsapiPostEntitiesInvestigatablesV1Request) *PostEntitiesAlertsV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post entities alerts v1 params
func (o *PostEntitiesAlertsV1Params) SetBody(body *models.DetectsapiPostEntitiesInvestigatablesV1Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostEntitiesAlertsV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
