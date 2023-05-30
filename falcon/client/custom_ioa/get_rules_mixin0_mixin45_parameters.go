// Code generated by go-swagger; DO NOT EDIT.

package custom_ioa

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
	"github.com/go-openapi/swag"
)

// NewGetRulesMixin0Mixin45Params creates a new GetRulesMixin0Mixin45Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRulesMixin0Mixin45Params() *GetRulesMixin0Mixin45Params {
	return &GetRulesMixin0Mixin45Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRulesMixin0Mixin45ParamsWithTimeout creates a new GetRulesMixin0Mixin45Params object
// with the ability to set a timeout on a request.
func NewGetRulesMixin0Mixin45ParamsWithTimeout(timeout time.Duration) *GetRulesMixin0Mixin45Params {
	return &GetRulesMixin0Mixin45Params{
		timeout: timeout,
	}
}

// NewGetRulesMixin0Mixin45ParamsWithContext creates a new GetRulesMixin0Mixin45Params object
// with the ability to set a context for a request.
func NewGetRulesMixin0Mixin45ParamsWithContext(ctx context.Context) *GetRulesMixin0Mixin45Params {
	return &GetRulesMixin0Mixin45Params{
		Context: ctx,
	}
}

// NewGetRulesMixin0Mixin45ParamsWithHTTPClient creates a new GetRulesMixin0Mixin45Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetRulesMixin0Mixin45ParamsWithHTTPClient(client *http.Client) *GetRulesMixin0Mixin45Params {
	return &GetRulesMixin0Mixin45Params{
		HTTPClient: client,
	}
}

/*
GetRulesMixin0Mixin45Params contains all the parameters to send to the API endpoint

	for the get rules mixin0 mixin45 operation.

	Typically these are written to a http.Request.
*/
type GetRulesMixin0Mixin45Params struct {

	/* Ids.

	   The IDs of the entities
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get rules mixin0 mixin45 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRulesMixin0Mixin45Params) WithDefaults() *GetRulesMixin0Mixin45Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get rules mixin0 mixin45 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRulesMixin0Mixin45Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get rules mixin0 mixin45 params
func (o *GetRulesMixin0Mixin45Params) WithTimeout(timeout time.Duration) *GetRulesMixin0Mixin45Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get rules mixin0 mixin45 params
func (o *GetRulesMixin0Mixin45Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get rules mixin0 mixin45 params
func (o *GetRulesMixin0Mixin45Params) WithContext(ctx context.Context) *GetRulesMixin0Mixin45Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get rules mixin0 mixin45 params
func (o *GetRulesMixin0Mixin45Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get rules mixin0 mixin45 params
func (o *GetRulesMixin0Mixin45Params) WithHTTPClient(client *http.Client) *GetRulesMixin0Mixin45Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get rules mixin0 mixin45 params
func (o *GetRulesMixin0Mixin45Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get rules mixin0 mixin45 params
func (o *GetRulesMixin0Mixin45Params) WithIds(ids []string) *GetRulesMixin0Mixin45Params {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get rules mixin0 mixin45 params
func (o *GetRulesMixin0Mixin45Params) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetRulesMixin0Mixin45Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Ids != nil {

		// binding items for ids
		joinedIds := o.bindParamIds(reg)

		// query array param ids
		if err := r.SetQueryParam("ids", joinedIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetRulesMixin0Mixin45 binds the parameter ids
func (o *GetRulesMixin0Mixin45Params) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []string

		idsIIV := idsIIR // string as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: "multi"
	idsIS := swag.JoinByFormat(idsIC, "multi")

	return idsIS
}
