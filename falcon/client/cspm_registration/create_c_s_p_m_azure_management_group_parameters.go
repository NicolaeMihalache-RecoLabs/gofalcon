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

// NewCreateCSPMAzureManagementGroupParams creates a new CreateCSPMAzureManagementGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateCSPMAzureManagementGroupParams() *CreateCSPMAzureManagementGroupParams {
	return &CreateCSPMAzureManagementGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCSPMAzureManagementGroupParamsWithTimeout creates a new CreateCSPMAzureManagementGroupParams object
// with the ability to set a timeout on a request.
func NewCreateCSPMAzureManagementGroupParamsWithTimeout(timeout time.Duration) *CreateCSPMAzureManagementGroupParams {
	return &CreateCSPMAzureManagementGroupParams{
		timeout: timeout,
	}
}

// NewCreateCSPMAzureManagementGroupParamsWithContext creates a new CreateCSPMAzureManagementGroupParams object
// with the ability to set a context for a request.
func NewCreateCSPMAzureManagementGroupParamsWithContext(ctx context.Context) *CreateCSPMAzureManagementGroupParams {
	return &CreateCSPMAzureManagementGroupParams{
		Context: ctx,
	}
}

// NewCreateCSPMAzureManagementGroupParamsWithHTTPClient creates a new CreateCSPMAzureManagementGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateCSPMAzureManagementGroupParamsWithHTTPClient(client *http.Client) *CreateCSPMAzureManagementGroupParams {
	return &CreateCSPMAzureManagementGroupParams{
		HTTPClient: client,
	}
}

/*
CreateCSPMAzureManagementGroupParams contains all the parameters to send to the API endpoint

	for the create c s p m azure management group operation.

	Typically these are written to a http.Request.
*/
type CreateCSPMAzureManagementGroupParams struct {

	// Body.
	Body *models.RegistrationAzureManagementGroupCreateRequestExternalV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create c s p m azure management group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCSPMAzureManagementGroupParams) WithDefaults() *CreateCSPMAzureManagementGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create c s p m azure management group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCSPMAzureManagementGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create c s p m azure management group params
func (o *CreateCSPMAzureManagementGroupParams) WithTimeout(timeout time.Duration) *CreateCSPMAzureManagementGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create c s p m azure management group params
func (o *CreateCSPMAzureManagementGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create c s p m azure management group params
func (o *CreateCSPMAzureManagementGroupParams) WithContext(ctx context.Context) *CreateCSPMAzureManagementGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create c s p m azure management group params
func (o *CreateCSPMAzureManagementGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create c s p m azure management group params
func (o *CreateCSPMAzureManagementGroupParams) WithHTTPClient(client *http.Client) *CreateCSPMAzureManagementGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create c s p m azure management group params
func (o *CreateCSPMAzureManagementGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create c s p m azure management group params
func (o *CreateCSPMAzureManagementGroupParams) WithBody(body *models.RegistrationAzureManagementGroupCreateRequestExternalV1) *CreateCSPMAzureManagementGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create c s p m azure management group params
func (o *CreateCSPMAzureManagementGroupParams) SetBody(body *models.RegistrationAzureManagementGroupCreateRequestExternalV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCSPMAzureManagementGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
