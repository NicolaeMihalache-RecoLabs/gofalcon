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

// NewUpdateCSPMGCPAccountParams creates a new UpdateCSPMGCPAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCSPMGCPAccountParams() *UpdateCSPMGCPAccountParams {
	return &UpdateCSPMGCPAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCSPMGCPAccountParamsWithTimeout creates a new UpdateCSPMGCPAccountParams object
// with the ability to set a timeout on a request.
func NewUpdateCSPMGCPAccountParamsWithTimeout(timeout time.Duration) *UpdateCSPMGCPAccountParams {
	return &UpdateCSPMGCPAccountParams{
		timeout: timeout,
	}
}

// NewUpdateCSPMGCPAccountParamsWithContext creates a new UpdateCSPMGCPAccountParams object
// with the ability to set a context for a request.
func NewUpdateCSPMGCPAccountParamsWithContext(ctx context.Context) *UpdateCSPMGCPAccountParams {
	return &UpdateCSPMGCPAccountParams{
		Context: ctx,
	}
}

// NewUpdateCSPMGCPAccountParamsWithHTTPClient creates a new UpdateCSPMGCPAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCSPMGCPAccountParamsWithHTTPClient(client *http.Client) *UpdateCSPMGCPAccountParams {
	return &UpdateCSPMGCPAccountParams{
		HTTPClient: client,
	}
}

/*
UpdateCSPMGCPAccountParams contains all the parameters to send to the API endpoint

	for the update c s p m g c p account operation.

	Typically these are written to a http.Request.
*/
type UpdateCSPMGCPAccountParams struct {

	// Body.
	Body *models.RegistrationGCPAccountPatchRequestV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update c s p m g c p account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCSPMGCPAccountParams) WithDefaults() *UpdateCSPMGCPAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update c s p m g c p account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCSPMGCPAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update c s p m g c p account params
func (o *UpdateCSPMGCPAccountParams) WithTimeout(timeout time.Duration) *UpdateCSPMGCPAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update c s p m g c p account params
func (o *UpdateCSPMGCPAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update c s p m g c p account params
func (o *UpdateCSPMGCPAccountParams) WithContext(ctx context.Context) *UpdateCSPMGCPAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update c s p m g c p account params
func (o *UpdateCSPMGCPAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update c s p m g c p account params
func (o *UpdateCSPMGCPAccountParams) WithHTTPClient(client *http.Client) *UpdateCSPMGCPAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update c s p m g c p account params
func (o *UpdateCSPMGCPAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update c s p m g c p account params
func (o *UpdateCSPMGCPAccountParams) WithBody(body *models.RegistrationGCPAccountPatchRequestV1) *UpdateCSPMGCPAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update c s p m g c p account params
func (o *UpdateCSPMGCPAccountParams) SetBody(body *models.RegistrationGCPAccountPatchRequestV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCSPMGCPAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
