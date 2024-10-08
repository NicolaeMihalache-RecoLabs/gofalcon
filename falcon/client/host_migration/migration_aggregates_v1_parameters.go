// Code generated by go-swagger; DO NOT EDIT.

package host_migration

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

// NewMigrationAggregatesV1Params creates a new MigrationAggregatesV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMigrationAggregatesV1Params() *MigrationAggregatesV1Params {
	return &MigrationAggregatesV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewMigrationAggregatesV1ParamsWithTimeout creates a new MigrationAggregatesV1Params object
// with the ability to set a timeout on a request.
func NewMigrationAggregatesV1ParamsWithTimeout(timeout time.Duration) *MigrationAggregatesV1Params {
	return &MigrationAggregatesV1Params{
		timeout: timeout,
	}
}

// NewMigrationAggregatesV1ParamsWithContext creates a new MigrationAggregatesV1Params object
// with the ability to set a context for a request.
func NewMigrationAggregatesV1ParamsWithContext(ctx context.Context) *MigrationAggregatesV1Params {
	return &MigrationAggregatesV1Params{
		Context: ctx,
	}
}

// NewMigrationAggregatesV1ParamsWithHTTPClient creates a new MigrationAggregatesV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewMigrationAggregatesV1ParamsWithHTTPClient(client *http.Client) *MigrationAggregatesV1Params {
	return &MigrationAggregatesV1Params{
		HTTPClient: client,
	}
}

/*
MigrationAggregatesV1Params contains all the parameters to send to the API endpoint

	for the migration aggregates v1 operation.

	Typically these are written to a http.Request.
*/
type MigrationAggregatesV1Params struct {

	// Body.
	Body []*models.MsaAggregateQueryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the migration aggregates v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MigrationAggregatesV1Params) WithDefaults() *MigrationAggregatesV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the migration aggregates v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MigrationAggregatesV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the migration aggregates v1 params
func (o *MigrationAggregatesV1Params) WithTimeout(timeout time.Duration) *MigrationAggregatesV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the migration aggregates v1 params
func (o *MigrationAggregatesV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the migration aggregates v1 params
func (o *MigrationAggregatesV1Params) WithContext(ctx context.Context) *MigrationAggregatesV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the migration aggregates v1 params
func (o *MigrationAggregatesV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the migration aggregates v1 params
func (o *MigrationAggregatesV1Params) WithHTTPClient(client *http.Client) *MigrationAggregatesV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the migration aggregates v1 params
func (o *MigrationAggregatesV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the migration aggregates v1 params
func (o *MigrationAggregatesV1Params) WithBody(body []*models.MsaAggregateQueryRequest) *MigrationAggregatesV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the migration aggregates v1 params
func (o *MigrationAggregatesV1Params) SetBody(body []*models.MsaAggregateQueryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *MigrationAggregatesV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
