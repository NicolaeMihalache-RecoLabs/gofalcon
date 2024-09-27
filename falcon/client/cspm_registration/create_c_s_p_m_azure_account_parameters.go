// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

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

// NewCreateCSPMAzureAccountParams creates a new CreateCSPMAzureAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateCSPMAzureAccountParams() *CreateCSPMAzureAccountParams {
	return &CreateCSPMAzureAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCSPMAzureAccountParamsWithTimeout creates a new CreateCSPMAzureAccountParams object
// with the ability to set a timeout on a request.
func NewCreateCSPMAzureAccountParamsWithTimeout(timeout time.Duration) *CreateCSPMAzureAccountParams {
	return &CreateCSPMAzureAccountParams{
		timeout: timeout,
	}
}

// NewCreateCSPMAzureAccountParamsWithContext creates a new CreateCSPMAzureAccountParams object
// with the ability to set a context for a request.
func NewCreateCSPMAzureAccountParamsWithContext(ctx context.Context) *CreateCSPMAzureAccountParams {
	return &CreateCSPMAzureAccountParams{
		Context: ctx,
	}
}

// NewCreateCSPMAzureAccountParamsWithHTTPClient creates a new CreateCSPMAzureAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateCSPMAzureAccountParamsWithHTTPClient(client *http.Client) *CreateCSPMAzureAccountParams {
	return &CreateCSPMAzureAccountParams{
		HTTPClient: client,
	}
}

/*
CreateCSPMAzureAccountParams contains all the parameters to send to the API endpoint

	for the create c s p m azure account operation.

	Typically these are written to a http.Request.
*/
type CreateCSPMAzureAccountParams struct {

	// Body.
	Body *models.RegistrationAzureAccountCreateRequestExternalV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create c s p m azure account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCSPMAzureAccountParams) WithDefaults() *CreateCSPMAzureAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create c s p m azure account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCSPMAzureAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create c s p m azure account params
func (o *CreateCSPMAzureAccountParams) WithTimeout(timeout time.Duration) *CreateCSPMAzureAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create c s p m azure account params
func (o *CreateCSPMAzureAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create c s p m azure account params
func (o *CreateCSPMAzureAccountParams) WithContext(ctx context.Context) *CreateCSPMAzureAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create c s p m azure account params
func (o *CreateCSPMAzureAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create c s p m azure account params
func (o *CreateCSPMAzureAccountParams) WithHTTPClient(client *http.Client) *CreateCSPMAzureAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create c s p m azure account params
func (o *CreateCSPMAzureAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create c s p m azure account params
func (o *CreateCSPMAzureAccountParams) WithBody(body *models.RegistrationAzureAccountCreateRequestExternalV1) *CreateCSPMAzureAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create c s p m azure account params
func (o *CreateCSPMAzureAccountParams) SetBody(body *models.RegistrationAzureAccountCreateRequestExternalV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCSPMAzureAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
