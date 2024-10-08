// Code generated by go-swagger; DO NOT EDIT.

package real_time_response

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

// NewRTRListSessionsParams creates a new RTRListSessionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRTRListSessionsParams() *RTRListSessionsParams {
	return &RTRListSessionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRTRListSessionsParamsWithTimeout creates a new RTRListSessionsParams object
// with the ability to set a timeout on a request.
func NewRTRListSessionsParamsWithTimeout(timeout time.Duration) *RTRListSessionsParams {
	return &RTRListSessionsParams{
		timeout: timeout,
	}
}

// NewRTRListSessionsParamsWithContext creates a new RTRListSessionsParams object
// with the ability to set a context for a request.
func NewRTRListSessionsParamsWithContext(ctx context.Context) *RTRListSessionsParams {
	return &RTRListSessionsParams{
		Context: ctx,
	}
}

// NewRTRListSessionsParamsWithHTTPClient creates a new RTRListSessionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRTRListSessionsParamsWithHTTPClient(client *http.Client) *RTRListSessionsParams {
	return &RTRListSessionsParams{
		HTTPClient: client,
	}
}

/*
RTRListSessionsParams contains all the parameters to send to the API endpoint

	for the r t r list sessions operation.

	Typically these are written to a http.Request.
*/
type RTRListSessionsParams struct {

	/* Body.

	 **`ids`** List of RTR sessions to retrieve.  RTR will only return the sessions that were created by the calling user
	 */
	Body *models.MsaIdsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the r t r list sessions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RTRListSessionsParams) WithDefaults() *RTRListSessionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the r t r list sessions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RTRListSessionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the r t r list sessions params
func (o *RTRListSessionsParams) WithTimeout(timeout time.Duration) *RTRListSessionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the r t r list sessions params
func (o *RTRListSessionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the r t r list sessions params
func (o *RTRListSessionsParams) WithContext(ctx context.Context) *RTRListSessionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the r t r list sessions params
func (o *RTRListSessionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the r t r list sessions params
func (o *RTRListSessionsParams) WithHTTPClient(client *http.Client) *RTRListSessionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the r t r list sessions params
func (o *RTRListSessionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the r t r list sessions params
func (o *RTRListSessionsParams) WithBody(body *models.MsaIdsRequest) *RTRListSessionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the r t r list sessions params
func (o *RTRListSessionsParams) SetBody(body *models.MsaIdsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RTRListSessionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
