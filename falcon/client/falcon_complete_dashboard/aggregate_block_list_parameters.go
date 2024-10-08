// Code generated by go-swagger; DO NOT EDIT.

package falcon_complete_dashboard

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

	"github.com/NicolaeMihalache-RecoLab/gofalcon/falcon/models"
)

// NewAggregateBlockListParams creates a new AggregateBlockListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAggregateBlockListParams() *AggregateBlockListParams {
	return &AggregateBlockListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAggregateBlockListParamsWithTimeout creates a new AggregateBlockListParams object
// with the ability to set a timeout on a request.
func NewAggregateBlockListParamsWithTimeout(timeout time.Duration) *AggregateBlockListParams {
	return &AggregateBlockListParams{
		timeout: timeout,
	}
}

// NewAggregateBlockListParamsWithContext creates a new AggregateBlockListParams object
// with the ability to set a context for a request.
func NewAggregateBlockListParamsWithContext(ctx context.Context) *AggregateBlockListParams {
	return &AggregateBlockListParams{
		Context: ctx,
	}
}

// NewAggregateBlockListParamsWithHTTPClient creates a new AggregateBlockListParams object
// with the ability to set a custom HTTPClient for a request.
func NewAggregateBlockListParamsWithHTTPClient(client *http.Client) *AggregateBlockListParams {
	return &AggregateBlockListParams{
		HTTPClient: client,
	}
}

/*
AggregateBlockListParams contains all the parameters to send to the API endpoint

	for the aggregate block list operation.

	Typically these are written to a http.Request.
*/
type AggregateBlockListParams struct {

	// Body.
	Body []*models.MsaAggregateQueryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the aggregate block list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregateBlockListParams) WithDefaults() *AggregateBlockListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the aggregate block list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregateBlockListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the aggregate block list params
func (o *AggregateBlockListParams) WithTimeout(timeout time.Duration) *AggregateBlockListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the aggregate block list params
func (o *AggregateBlockListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the aggregate block list params
func (o *AggregateBlockListParams) WithContext(ctx context.Context) *AggregateBlockListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the aggregate block list params
func (o *AggregateBlockListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the aggregate block list params
func (o *AggregateBlockListParams) WithHTTPClient(client *http.Client) *AggregateBlockListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the aggregate block list params
func (o *AggregateBlockListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the aggregate block list params
func (o *AggregateBlockListParams) WithBody(body []*models.MsaAggregateQueryRequest) *AggregateBlockListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the aggregate block list params
func (o *AggregateBlockListParams) SetBody(body []*models.MsaAggregateQueryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AggregateBlockListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
