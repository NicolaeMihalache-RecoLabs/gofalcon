// Code generated by go-swagger; DO NOT EDIT.

package message_center

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

// NewCreateCaseParams creates a new CreateCaseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateCaseParams() *CreateCaseParams {
	return &CreateCaseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCaseParamsWithTimeout creates a new CreateCaseParams object
// with the ability to set a timeout on a request.
func NewCreateCaseParamsWithTimeout(timeout time.Duration) *CreateCaseParams {
	return &CreateCaseParams{
		timeout: timeout,
	}
}

// NewCreateCaseParamsWithContext creates a new CreateCaseParams object
// with the ability to set a context for a request.
func NewCreateCaseParamsWithContext(ctx context.Context) *CreateCaseParams {
	return &CreateCaseParams{
		Context: ctx,
	}
}

// NewCreateCaseParamsWithHTTPClient creates a new CreateCaseParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateCaseParamsWithHTTPClient(client *http.Client) *CreateCaseParams {
	return &CreateCaseParams{
		HTTPClient: client,
	}
}

/*
CreateCaseParams contains all the parameters to send to the API endpoint

	for the create case operation.

	Typically these are written to a http.Request.
*/
type CreateCaseParams struct {

	// Body.
	Body *models.DomainCaseCreationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create case params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCaseParams) WithDefaults() *CreateCaseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create case params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCaseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create case params
func (o *CreateCaseParams) WithTimeout(timeout time.Duration) *CreateCaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create case params
func (o *CreateCaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create case params
func (o *CreateCaseParams) WithContext(ctx context.Context) *CreateCaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create case params
func (o *CreateCaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create case params
func (o *CreateCaseParams) WithHTTPClient(client *http.Client) *CreateCaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create case params
func (o *CreateCaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create case params
func (o *CreateCaseParams) WithBody(body *models.DomainCaseCreationRequest) *CreateCaseParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create case params
func (o *CreateCaseParams) SetBody(body *models.DomainCaseCreationRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
