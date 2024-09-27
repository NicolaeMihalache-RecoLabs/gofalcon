// Code generated by go-swagger; DO NOT EDIT.

package ods

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

// NewAggregateQueryScanHostMetadataParams creates a new AggregateQueryScanHostMetadataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAggregateQueryScanHostMetadataParams() *AggregateQueryScanHostMetadataParams {
	return &AggregateQueryScanHostMetadataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAggregateQueryScanHostMetadataParamsWithTimeout creates a new AggregateQueryScanHostMetadataParams object
// with the ability to set a timeout on a request.
func NewAggregateQueryScanHostMetadataParamsWithTimeout(timeout time.Duration) *AggregateQueryScanHostMetadataParams {
	return &AggregateQueryScanHostMetadataParams{
		timeout: timeout,
	}
}

// NewAggregateQueryScanHostMetadataParamsWithContext creates a new AggregateQueryScanHostMetadataParams object
// with the ability to set a context for a request.
func NewAggregateQueryScanHostMetadataParamsWithContext(ctx context.Context) *AggregateQueryScanHostMetadataParams {
	return &AggregateQueryScanHostMetadataParams{
		Context: ctx,
	}
}

// NewAggregateQueryScanHostMetadataParamsWithHTTPClient creates a new AggregateQueryScanHostMetadataParams object
// with the ability to set a custom HTTPClient for a request.
func NewAggregateQueryScanHostMetadataParamsWithHTTPClient(client *http.Client) *AggregateQueryScanHostMetadataParams {
	return &AggregateQueryScanHostMetadataParams{
		HTTPClient: client,
	}
}

/*
AggregateQueryScanHostMetadataParams contains all the parameters to send to the API endpoint

	for the aggregate query scan host metadata operation.

	Typically these are written to a http.Request.
*/
type AggregateQueryScanHostMetadataParams struct {

	// Body.
	Body []*models.MsaAggregateQueryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the aggregate query scan host metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregateQueryScanHostMetadataParams) WithDefaults() *AggregateQueryScanHostMetadataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the aggregate query scan host metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregateQueryScanHostMetadataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the aggregate query scan host metadata params
func (o *AggregateQueryScanHostMetadataParams) WithTimeout(timeout time.Duration) *AggregateQueryScanHostMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the aggregate query scan host metadata params
func (o *AggregateQueryScanHostMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the aggregate query scan host metadata params
func (o *AggregateQueryScanHostMetadataParams) WithContext(ctx context.Context) *AggregateQueryScanHostMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the aggregate query scan host metadata params
func (o *AggregateQueryScanHostMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the aggregate query scan host metadata params
func (o *AggregateQueryScanHostMetadataParams) WithHTTPClient(client *http.Client) *AggregateQueryScanHostMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the aggregate query scan host metadata params
func (o *AggregateQueryScanHostMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the aggregate query scan host metadata params
func (o *AggregateQueryScanHostMetadataParams) WithBody(body []*models.MsaAggregateQueryRequest) *AggregateQueryScanHostMetadataParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the aggregate query scan host metadata params
func (o *AggregateQueryScanHostMetadataParams) SetBody(body []*models.MsaAggregateQueryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AggregateQueryScanHostMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
