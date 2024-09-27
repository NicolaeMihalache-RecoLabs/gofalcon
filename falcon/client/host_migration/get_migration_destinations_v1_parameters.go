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

// NewGetMigrationDestinationsV1Params creates a new GetMigrationDestinationsV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMigrationDestinationsV1Params() *GetMigrationDestinationsV1Params {
	return &GetMigrationDestinationsV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMigrationDestinationsV1ParamsWithTimeout creates a new GetMigrationDestinationsV1Params object
// with the ability to set a timeout on a request.
func NewGetMigrationDestinationsV1ParamsWithTimeout(timeout time.Duration) *GetMigrationDestinationsV1Params {
	return &GetMigrationDestinationsV1Params{
		timeout: timeout,
	}
}

// NewGetMigrationDestinationsV1ParamsWithContext creates a new GetMigrationDestinationsV1Params object
// with the ability to set a context for a request.
func NewGetMigrationDestinationsV1ParamsWithContext(ctx context.Context) *GetMigrationDestinationsV1Params {
	return &GetMigrationDestinationsV1Params{
		Context: ctx,
	}
}

// NewGetMigrationDestinationsV1ParamsWithHTTPClient creates a new GetMigrationDestinationsV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetMigrationDestinationsV1ParamsWithHTTPClient(client *http.Client) *GetMigrationDestinationsV1Params {
	return &GetMigrationDestinationsV1Params{
		HTTPClient: client,
	}
}

/*
GetMigrationDestinationsV1Params contains all the parameters to send to the API endpoint

	for the get migration destinations v1 operation.

	Typically these are written to a http.Request.
*/
type GetMigrationDestinationsV1Params struct {

	// Body.
	Body *models.APIGetMigrationDestinationsRequestBodyV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get migration destinations v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMigrationDestinationsV1Params) WithDefaults() *GetMigrationDestinationsV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get migration destinations v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMigrationDestinationsV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get migration destinations v1 params
func (o *GetMigrationDestinationsV1Params) WithTimeout(timeout time.Duration) *GetMigrationDestinationsV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get migration destinations v1 params
func (o *GetMigrationDestinationsV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get migration destinations v1 params
func (o *GetMigrationDestinationsV1Params) WithContext(ctx context.Context) *GetMigrationDestinationsV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get migration destinations v1 params
func (o *GetMigrationDestinationsV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get migration destinations v1 params
func (o *GetMigrationDestinationsV1Params) WithHTTPClient(client *http.Client) *GetMigrationDestinationsV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get migration destinations v1 params
func (o *GetMigrationDestinationsV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get migration destinations v1 params
func (o *GetMigrationDestinationsV1Params) WithBody(body *models.APIGetMigrationDestinationsRequestBodyV1) *GetMigrationDestinationsV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get migration destinations v1 params
func (o *GetMigrationDestinationsV1Params) SetBody(body *models.APIGetMigrationDestinationsRequestBodyV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetMigrationDestinationsV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
