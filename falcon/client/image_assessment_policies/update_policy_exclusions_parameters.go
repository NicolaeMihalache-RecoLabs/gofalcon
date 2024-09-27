// Code generated by go-swagger; DO NOT EDIT.

package image_assessment_policies

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

// NewUpdatePolicyExclusionsParams creates a new UpdatePolicyExclusionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdatePolicyExclusionsParams() *UpdatePolicyExclusionsParams {
	return &UpdatePolicyExclusionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePolicyExclusionsParamsWithTimeout creates a new UpdatePolicyExclusionsParams object
// with the ability to set a timeout on a request.
func NewUpdatePolicyExclusionsParamsWithTimeout(timeout time.Duration) *UpdatePolicyExclusionsParams {
	return &UpdatePolicyExclusionsParams{
		timeout: timeout,
	}
}

// NewUpdatePolicyExclusionsParamsWithContext creates a new UpdatePolicyExclusionsParams object
// with the ability to set a context for a request.
func NewUpdatePolicyExclusionsParamsWithContext(ctx context.Context) *UpdatePolicyExclusionsParams {
	return &UpdatePolicyExclusionsParams{
		Context: ctx,
	}
}

// NewUpdatePolicyExclusionsParamsWithHTTPClient creates a new UpdatePolicyExclusionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdatePolicyExclusionsParamsWithHTTPClient(client *http.Client) *UpdatePolicyExclusionsParams {
	return &UpdatePolicyExclusionsParams{
		HTTPClient: client,
	}
}

/*
UpdatePolicyExclusionsParams contains all the parameters to send to the API endpoint

	for the update policy exclusions operation.

	Typically these are written to a http.Request.
*/
type UpdatePolicyExclusionsParams struct {

	// Body.
	Body *models.ModelsUpdateExclusionsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update policy exclusions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePolicyExclusionsParams) WithDefaults() *UpdatePolicyExclusionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update policy exclusions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePolicyExclusionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update policy exclusions params
func (o *UpdatePolicyExclusionsParams) WithTimeout(timeout time.Duration) *UpdatePolicyExclusionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update policy exclusions params
func (o *UpdatePolicyExclusionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update policy exclusions params
func (o *UpdatePolicyExclusionsParams) WithContext(ctx context.Context) *UpdatePolicyExclusionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update policy exclusions params
func (o *UpdatePolicyExclusionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update policy exclusions params
func (o *UpdatePolicyExclusionsParams) WithHTTPClient(client *http.Client) *UpdatePolicyExclusionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update policy exclusions params
func (o *UpdatePolicyExclusionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update policy exclusions params
func (o *UpdatePolicyExclusionsParams) WithBody(body *models.ModelsUpdateExclusionsRequest) *UpdatePolicyExclusionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update policy exclusions params
func (o *UpdatePolicyExclusionsParams) SetBody(body *models.ModelsUpdateExclusionsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePolicyExclusionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
