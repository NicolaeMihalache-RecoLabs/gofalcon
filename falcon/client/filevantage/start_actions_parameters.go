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

// NewStartActionsParams creates a new StartActionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStartActionsParams() *StartActionsParams {
	return &StartActionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStartActionsParamsWithTimeout creates a new StartActionsParams object
// with the ability to set a timeout on a request.
func NewStartActionsParamsWithTimeout(timeout time.Duration) *StartActionsParams {
	return &StartActionsParams{
		timeout: timeout,
	}
}

// NewStartActionsParamsWithContext creates a new StartActionsParams object
// with the ability to set a context for a request.
func NewStartActionsParamsWithContext(ctx context.Context) *StartActionsParams {
	return &StartActionsParams{
		Context: ctx,
	}
}

// NewStartActionsParamsWithHTTPClient creates a new StartActionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewStartActionsParamsWithHTTPClient(client *http.Client) *StartActionsParams {
	return &StartActionsParams{
		HTTPClient: client,
	}
}

/*
StartActionsParams contains all the parameters to send to the API endpoint

	for the start actions operation.

	Typically these are written to a http.Request.
*/
type StartActionsParams struct {

	/* Body.

	    Create a new action.

	* `operation` must be one of the `suppress`, `unsuppress`, or `purge`

	* `change_ids` represent the ids of the changes the operation will perform; limited to 100 ids per action

	* `comment` optional comment to describe the reason for the action
	*/
	Body *models.ActionsCreateActionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the start actions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartActionsParams) WithDefaults() *StartActionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the start actions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartActionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the start actions params
func (o *StartActionsParams) WithTimeout(timeout time.Duration) *StartActionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start actions params
func (o *StartActionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start actions params
func (o *StartActionsParams) WithContext(ctx context.Context) *StartActionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start actions params
func (o *StartActionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start actions params
func (o *StartActionsParams) WithHTTPClient(client *http.Client) *StartActionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start actions params
func (o *StartActionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the start actions params
func (o *StartActionsParams) WithBody(body *models.ActionsCreateActionRequest) *StartActionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the start actions params
func (o *StartActionsParams) SetBody(body *models.ActionsCreateActionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *StartActionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
