// Code generated by go-swagger; DO NOT EDIT.

package filevantage

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

// NewCreatePoliciesParams creates a new CreatePoliciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePoliciesParams() *CreatePoliciesParams {
	return &CreatePoliciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePoliciesParamsWithTimeout creates a new CreatePoliciesParams object
// with the ability to set a timeout on a request.
func NewCreatePoliciesParamsWithTimeout(timeout time.Duration) *CreatePoliciesParams {
	return &CreatePoliciesParams{
		timeout: timeout,
	}
}

// NewCreatePoliciesParamsWithContext creates a new CreatePoliciesParams object
// with the ability to set a context for a request.
func NewCreatePoliciesParamsWithContext(ctx context.Context) *CreatePoliciesParams {
	return &CreatePoliciesParams{
		Context: ctx,
	}
}

// NewCreatePoliciesParamsWithHTTPClient creates a new CreatePoliciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePoliciesParamsWithHTTPClient(client *http.Client) *CreatePoliciesParams {
	return &CreatePoliciesParams{
		HTTPClient: client,
	}
}

/*
CreatePoliciesParams contains all the parameters to send to the API endpoint

	for the create policies operation.

	Typically these are written to a http.Request.
*/
type CreatePoliciesParams struct {

	/* Body.

	    Create a new policy.

	* `name` must be between 1 and 100 characters.

	* `description` can be between 0 and 500 characters.

	* `platform` must be one of `Windows`, `Linux`, or `Mac`

	Rule and host group assignment and policy precedence setting is performed via their respective patch end-points.
	*/
	Body *models.PoliciesCreateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePoliciesParams) WithDefaults() *CreatePoliciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePoliciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create policies params
func (o *CreatePoliciesParams) WithTimeout(timeout time.Duration) *CreatePoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create policies params
func (o *CreatePoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create policies params
func (o *CreatePoliciesParams) WithContext(ctx context.Context) *CreatePoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create policies params
func (o *CreatePoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create policies params
func (o *CreatePoliciesParams) WithHTTPClient(client *http.Client) *CreatePoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create policies params
func (o *CreatePoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create policies params
func (o *CreatePoliciesParams) WithBody(body *models.PoliciesCreateRequest) *CreatePoliciesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create policies params
func (o *CreatePoliciesParams) SetBody(body *models.PoliciesCreateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
