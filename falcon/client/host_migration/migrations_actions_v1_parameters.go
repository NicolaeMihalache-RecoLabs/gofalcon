// Code generated by go-swagger; DO NOT EDIT.

package host_migration

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

// NewMigrationsActionsV1Params creates a new MigrationsActionsV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMigrationsActionsV1Params() *MigrationsActionsV1Params {
	return &MigrationsActionsV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewMigrationsActionsV1ParamsWithTimeout creates a new MigrationsActionsV1Params object
// with the ability to set a timeout on a request.
func NewMigrationsActionsV1ParamsWithTimeout(timeout time.Duration) *MigrationsActionsV1Params {
	return &MigrationsActionsV1Params{
		timeout: timeout,
	}
}

// NewMigrationsActionsV1ParamsWithContext creates a new MigrationsActionsV1Params object
// with the ability to set a context for a request.
func NewMigrationsActionsV1ParamsWithContext(ctx context.Context) *MigrationsActionsV1Params {
	return &MigrationsActionsV1Params{
		Context: ctx,
	}
}

// NewMigrationsActionsV1ParamsWithHTTPClient creates a new MigrationsActionsV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewMigrationsActionsV1ParamsWithHTTPClient(client *http.Client) *MigrationsActionsV1Params {
	return &MigrationsActionsV1Params{
		HTTPClient: client,
	}
}

/*
MigrationsActionsV1Params contains all the parameters to send to the API endpoint

	for the migrations actions v1 operation.

	Typically these are written to a http.Request.
*/
type MigrationsActionsV1Params struct {

	/* ActionName.

	   The action to perform
	*/
	ActionName string

	// Body.
	Body *models.MsaEntityActionRequestV3

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the migrations actions v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MigrationsActionsV1Params) WithDefaults() *MigrationsActionsV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the migrations actions v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MigrationsActionsV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the migrations actions v1 params
func (o *MigrationsActionsV1Params) WithTimeout(timeout time.Duration) *MigrationsActionsV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the migrations actions v1 params
func (o *MigrationsActionsV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the migrations actions v1 params
func (o *MigrationsActionsV1Params) WithContext(ctx context.Context) *MigrationsActionsV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the migrations actions v1 params
func (o *MigrationsActionsV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the migrations actions v1 params
func (o *MigrationsActionsV1Params) WithHTTPClient(client *http.Client) *MigrationsActionsV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the migrations actions v1 params
func (o *MigrationsActionsV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionName adds the actionName to the migrations actions v1 params
func (o *MigrationsActionsV1Params) WithActionName(actionName string) *MigrationsActionsV1Params {
	o.SetActionName(actionName)
	return o
}

// SetActionName adds the actionName to the migrations actions v1 params
func (o *MigrationsActionsV1Params) SetActionName(actionName string) {
	o.ActionName = actionName
}

// WithBody adds the body to the migrations actions v1 params
func (o *MigrationsActionsV1Params) WithBody(body *models.MsaEntityActionRequestV3) *MigrationsActionsV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the migrations actions v1 params
func (o *MigrationsActionsV1Params) SetBody(body *models.MsaEntityActionRequestV3) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *MigrationsActionsV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param action_name
	qrActionName := o.ActionName
	qActionName := qrActionName
	if qActionName != "" {

		if err := r.SetQueryParam("action_name", qActionName); err != nil {
			return err
		}
	}
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
