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

// NewCreatePolicyGroupsParams creates a new CreatePolicyGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePolicyGroupsParams() *CreatePolicyGroupsParams {
	return &CreatePolicyGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePolicyGroupsParamsWithTimeout creates a new CreatePolicyGroupsParams object
// with the ability to set a timeout on a request.
func NewCreatePolicyGroupsParamsWithTimeout(timeout time.Duration) *CreatePolicyGroupsParams {
	return &CreatePolicyGroupsParams{
		timeout: timeout,
	}
}

// NewCreatePolicyGroupsParamsWithContext creates a new CreatePolicyGroupsParams object
// with the ability to set a context for a request.
func NewCreatePolicyGroupsParamsWithContext(ctx context.Context) *CreatePolicyGroupsParams {
	return &CreatePolicyGroupsParams{
		Context: ctx,
	}
}

// NewCreatePolicyGroupsParamsWithHTTPClient creates a new CreatePolicyGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePolicyGroupsParamsWithHTTPClient(client *http.Client) *CreatePolicyGroupsParams {
	return &CreatePolicyGroupsParams{
		HTTPClient: client,
	}
}

/*
CreatePolicyGroupsParams contains all the parameters to send to the API endpoint

	for the create policy groups operation.

	Typically these are written to a http.Request.
*/
type CreatePolicyGroupsParams struct {

	// Body.
	Body *models.ModelsCreateImageGroupRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create policy groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePolicyGroupsParams) WithDefaults() *CreatePolicyGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create policy groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePolicyGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create policy groups params
func (o *CreatePolicyGroupsParams) WithTimeout(timeout time.Duration) *CreatePolicyGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create policy groups params
func (o *CreatePolicyGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create policy groups params
func (o *CreatePolicyGroupsParams) WithContext(ctx context.Context) *CreatePolicyGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create policy groups params
func (o *CreatePolicyGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create policy groups params
func (o *CreatePolicyGroupsParams) WithHTTPClient(client *http.Client) *CreatePolicyGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create policy groups params
func (o *CreatePolicyGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create policy groups params
func (o *CreatePolicyGroupsParams) WithBody(body *models.ModelsCreateImageGroupRequest) *CreatePolicyGroupsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create policy groups params
func (o *CreatePolicyGroupsParams) SetBody(body *models.ModelsCreateImageGroupRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePolicyGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
