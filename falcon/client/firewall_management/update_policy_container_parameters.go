// Code generated by go-swagger; DO NOT EDIT.

package firewall_management

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

// NewUpdatePolicyContainerParams creates a new UpdatePolicyContainerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdatePolicyContainerParams() *UpdatePolicyContainerParams {
	return &UpdatePolicyContainerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePolicyContainerParamsWithTimeout creates a new UpdatePolicyContainerParams object
// with the ability to set a timeout on a request.
func NewUpdatePolicyContainerParamsWithTimeout(timeout time.Duration) *UpdatePolicyContainerParams {
	return &UpdatePolicyContainerParams{
		timeout: timeout,
	}
}

// NewUpdatePolicyContainerParamsWithContext creates a new UpdatePolicyContainerParams object
// with the ability to set a context for a request.
func NewUpdatePolicyContainerParamsWithContext(ctx context.Context) *UpdatePolicyContainerParams {
	return &UpdatePolicyContainerParams{
		Context: ctx,
	}
}

// NewUpdatePolicyContainerParamsWithHTTPClient creates a new UpdatePolicyContainerParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdatePolicyContainerParamsWithHTTPClient(client *http.Client) *UpdatePolicyContainerParams {
	return &UpdatePolicyContainerParams{
		HTTPClient: client,
	}
}

/*
UpdatePolicyContainerParams contains all the parameters to send to the API endpoint

	for the update policy container operation.

	Typically these are written to a http.Request.
*/
type UpdatePolicyContainerParams struct {

	// Body.
	Body *models.FwmgrAPIPolicyContainerUpsertRequestV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update policy container params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePolicyContainerParams) WithDefaults() *UpdatePolicyContainerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update policy container params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePolicyContainerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update policy container params
func (o *UpdatePolicyContainerParams) WithTimeout(timeout time.Duration) *UpdatePolicyContainerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update policy container params
func (o *UpdatePolicyContainerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update policy container params
func (o *UpdatePolicyContainerParams) WithContext(ctx context.Context) *UpdatePolicyContainerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update policy container params
func (o *UpdatePolicyContainerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update policy container params
func (o *UpdatePolicyContainerParams) WithHTTPClient(client *http.Client) *UpdatePolicyContainerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update policy container params
func (o *UpdatePolicyContainerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update policy container params
func (o *UpdatePolicyContainerParams) WithBody(body *models.FwmgrAPIPolicyContainerUpsertRequestV1) *UpdatePolicyContainerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update policy container params
func (o *UpdatePolicyContainerParams) SetBody(body *models.FwmgrAPIPolicyContainerUpsertRequestV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePolicyContainerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
