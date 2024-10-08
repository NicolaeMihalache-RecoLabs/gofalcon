// Code generated by go-swagger; DO NOT EDIT.

package installation_tokens

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

// NewTokensCreateParams creates a new TokensCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTokensCreateParams() *TokensCreateParams {
	return &TokensCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTokensCreateParamsWithTimeout creates a new TokensCreateParams object
// with the ability to set a timeout on a request.
func NewTokensCreateParamsWithTimeout(timeout time.Duration) *TokensCreateParams {
	return &TokensCreateParams{
		timeout: timeout,
	}
}

// NewTokensCreateParamsWithContext creates a new TokensCreateParams object
// with the ability to set a context for a request.
func NewTokensCreateParamsWithContext(ctx context.Context) *TokensCreateParams {
	return &TokensCreateParams{
		Context: ctx,
	}
}

// NewTokensCreateParamsWithHTTPClient creates a new TokensCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTokensCreateParamsWithHTTPClient(client *http.Client) *TokensCreateParams {
	return &TokensCreateParams{
		HTTPClient: client,
	}
}

/*
TokensCreateParams contains all the parameters to send to the API endpoint

	for the tokens create operation.

	Typically these are written to a http.Request.
*/
type TokensCreateParams struct {

	// Body.
	Body *models.APITokenCreateRequestV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tokens create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TokensCreateParams) WithDefaults() *TokensCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tokens create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TokensCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tokens create params
func (o *TokensCreateParams) WithTimeout(timeout time.Duration) *TokensCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tokens create params
func (o *TokensCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tokens create params
func (o *TokensCreateParams) WithContext(ctx context.Context) *TokensCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tokens create params
func (o *TokensCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tokens create params
func (o *TokensCreateParams) WithHTTPClient(client *http.Client) *TokensCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tokens create params
func (o *TokensCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the tokens create params
func (o *TokensCreateParams) WithBody(body *models.APITokenCreateRequestV1) *TokensCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the tokens create params
func (o *TokensCreateParams) SetBody(body *models.APITokenCreateRequestV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TokensCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
