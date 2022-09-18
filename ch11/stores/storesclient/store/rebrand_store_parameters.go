// Code generated by go-swagger; DO NOT EDIT.

package store

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

	"eda-in-golang/stores/storesclient/models"
)

// NewRebrandStoreParams creates a new RebrandStoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRebrandStoreParams() *RebrandStoreParams {
	return &RebrandStoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRebrandStoreParamsWithTimeout creates a new RebrandStoreParams object
// with the ability to set a timeout on a request.
func NewRebrandStoreParamsWithTimeout(timeout time.Duration) *RebrandStoreParams {
	return &RebrandStoreParams{
		timeout: timeout,
	}
}

// NewRebrandStoreParamsWithContext creates a new RebrandStoreParams object
// with the ability to set a context for a request.
func NewRebrandStoreParamsWithContext(ctx context.Context) *RebrandStoreParams {
	return &RebrandStoreParams{
		Context: ctx,
	}
}

// NewRebrandStoreParamsWithHTTPClient creates a new RebrandStoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewRebrandStoreParamsWithHTTPClient(client *http.Client) *RebrandStoreParams {
	return &RebrandStoreParams{
		HTTPClient: client,
	}
}

/* RebrandStoreParams contains all the parameters to send to the API endpoint
   for the rebrand store operation.

   Typically these are written to a http.Request.
*/
type RebrandStoreParams struct {

	// Body.
	Body *models.RebrandStoreParamsBody

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rebrand store params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RebrandStoreParams) WithDefaults() *RebrandStoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rebrand store params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RebrandStoreParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the rebrand store params
func (o *RebrandStoreParams) WithTimeout(timeout time.Duration) *RebrandStoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rebrand store params
func (o *RebrandStoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rebrand store params
func (o *RebrandStoreParams) WithContext(ctx context.Context) *RebrandStoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rebrand store params
func (o *RebrandStoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rebrand store params
func (o *RebrandStoreParams) WithHTTPClient(client *http.Client) *RebrandStoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rebrand store params
func (o *RebrandStoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the rebrand store params
func (o *RebrandStoreParams) WithBody(body *models.RebrandStoreParamsBody) *RebrandStoreParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the rebrand store params
func (o *RebrandStoreParams) SetBody(body *models.RebrandStoreParamsBody) {
	o.Body = body
}

// WithID adds the id to the rebrand store params
func (o *RebrandStoreParams) WithID(id string) *RebrandStoreParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the rebrand store params
func (o *RebrandStoreParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RebrandStoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
