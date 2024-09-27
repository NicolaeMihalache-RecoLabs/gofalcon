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

// NewAggregateDetectionsParams creates a new AggregateDetectionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAggregateDetectionsParams() *AggregateDetectionsParams {
	return &AggregateDetectionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAggregateDetectionsParamsWithTimeout creates a new AggregateDetectionsParams object
// with the ability to set a timeout on a request.
func NewAggregateDetectionsParamsWithTimeout(timeout time.Duration) *AggregateDetectionsParams {
	return &AggregateDetectionsParams{
		timeout: timeout,
	}
}

// NewAggregateDetectionsParamsWithContext creates a new AggregateDetectionsParams object
// with the ability to set a context for a request.
func NewAggregateDetectionsParamsWithContext(ctx context.Context) *AggregateDetectionsParams {
	return &AggregateDetectionsParams{
		Context: ctx,
	}
}

// NewAggregateDetectionsParamsWithHTTPClient creates a new AggregateDetectionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAggregateDetectionsParamsWithHTTPClient(client *http.Client) *AggregateDetectionsParams {
	return &AggregateDetectionsParams{
		HTTPClient: client,
	}
}

/*
AggregateDetectionsParams contains all the parameters to send to the API endpoint

	for the aggregate detections operation.

	Typically these are written to a http.Request.
*/
type AggregateDetectionsParams struct {

	// Body.
	Body []*models.MsaAggregateQueryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the aggregate detections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregateDetectionsParams) WithDefaults() *AggregateDetectionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the aggregate detections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregateDetectionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the aggregate detections params
func (o *AggregateDetectionsParams) WithTimeout(timeout time.Duration) *AggregateDetectionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the aggregate detections params
func (o *AggregateDetectionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the aggregate detections params
func (o *AggregateDetectionsParams) WithContext(ctx context.Context) *AggregateDetectionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the aggregate detections params
func (o *AggregateDetectionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the aggregate detections params
func (o *AggregateDetectionsParams) WithHTTPClient(client *http.Client) *AggregateDetectionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the aggregate detections params
func (o *AggregateDetectionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the aggregate detections params
func (o *AggregateDetectionsParams) WithBody(body []*models.MsaAggregateQueryRequest) *AggregateDetectionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the aggregate detections params
func (o *AggregateDetectionsParams) SetBody(body []*models.MsaAggregateQueryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AggregateDetectionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
