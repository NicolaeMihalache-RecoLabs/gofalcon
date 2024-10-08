// Code generated by go-swagger; DO NOT EDIT.

package mobile_enrollment

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

// NewRequestDeviceEnrollmentV4Params creates a new RequestDeviceEnrollmentV4Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRequestDeviceEnrollmentV4Params() *RequestDeviceEnrollmentV4Params {
	return &RequestDeviceEnrollmentV4Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewRequestDeviceEnrollmentV4ParamsWithTimeout creates a new RequestDeviceEnrollmentV4Params object
// with the ability to set a timeout on a request.
func NewRequestDeviceEnrollmentV4ParamsWithTimeout(timeout time.Duration) *RequestDeviceEnrollmentV4Params {
	return &RequestDeviceEnrollmentV4Params{
		timeout: timeout,
	}
}

// NewRequestDeviceEnrollmentV4ParamsWithContext creates a new RequestDeviceEnrollmentV4Params object
// with the ability to set a context for a request.
func NewRequestDeviceEnrollmentV4ParamsWithContext(ctx context.Context) *RequestDeviceEnrollmentV4Params {
	return &RequestDeviceEnrollmentV4Params{
		Context: ctx,
	}
}

// NewRequestDeviceEnrollmentV4ParamsWithHTTPClient creates a new RequestDeviceEnrollmentV4Params object
// with the ability to set a custom HTTPClient for a request.
func NewRequestDeviceEnrollmentV4ParamsWithHTTPClient(client *http.Client) *RequestDeviceEnrollmentV4Params {
	return &RequestDeviceEnrollmentV4Params{
		HTTPClient: client,
	}
}

/*
RequestDeviceEnrollmentV4Params contains all the parameters to send to the API endpoint

	for the request device enrollment v4 operation.

	Typically these are written to a http.Request.
*/
type RequestDeviceEnrollmentV4Params struct {

	/* ActionName.

	   Action to perform
	*/
	ActionName string

	// Body.
	Body *models.APIPostEnrollmentDetailsV4

	/* Filter.

	   FQL filter
	*/
	Filter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the request device enrollment v4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RequestDeviceEnrollmentV4Params) WithDefaults() *RequestDeviceEnrollmentV4Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the request device enrollment v4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RequestDeviceEnrollmentV4Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the request device enrollment v4 params
func (o *RequestDeviceEnrollmentV4Params) WithTimeout(timeout time.Duration) *RequestDeviceEnrollmentV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the request device enrollment v4 params
func (o *RequestDeviceEnrollmentV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the request device enrollment v4 params
func (o *RequestDeviceEnrollmentV4Params) WithContext(ctx context.Context) *RequestDeviceEnrollmentV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the request device enrollment v4 params
func (o *RequestDeviceEnrollmentV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the request device enrollment v4 params
func (o *RequestDeviceEnrollmentV4Params) WithHTTPClient(client *http.Client) *RequestDeviceEnrollmentV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the request device enrollment v4 params
func (o *RequestDeviceEnrollmentV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionName adds the actionName to the request device enrollment v4 params
func (o *RequestDeviceEnrollmentV4Params) WithActionName(actionName string) *RequestDeviceEnrollmentV4Params {
	o.SetActionName(actionName)
	return o
}

// SetActionName adds the actionName to the request device enrollment v4 params
func (o *RequestDeviceEnrollmentV4Params) SetActionName(actionName string) {
	o.ActionName = actionName
}

// WithBody adds the body to the request device enrollment v4 params
func (o *RequestDeviceEnrollmentV4Params) WithBody(body *models.APIPostEnrollmentDetailsV4) *RequestDeviceEnrollmentV4Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the request device enrollment v4 params
func (o *RequestDeviceEnrollmentV4Params) SetBody(body *models.APIPostEnrollmentDetailsV4) {
	o.Body = body
}

// WithFilter adds the filter to the request device enrollment v4 params
func (o *RequestDeviceEnrollmentV4Params) WithFilter(filter *string) *RequestDeviceEnrollmentV4Params {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the request device enrollment v4 params
func (o *RequestDeviceEnrollmentV4Params) SetFilter(filter *string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *RequestDeviceEnrollmentV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param action_name
	qrActionName := o.ActionName
	qActionName := qrActionName

	if err := r.SetQueryParam("action_name", qActionName); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
