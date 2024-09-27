// Code generated by go-swagger; DO NOT EDIT.

package hosts

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

// NewPerformActionV2Params creates a new PerformActionV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformActionV2Params() *PerformActionV2Params {
	return &PerformActionV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformActionV2ParamsWithTimeout creates a new PerformActionV2Params object
// with the ability to set a timeout on a request.
func NewPerformActionV2ParamsWithTimeout(timeout time.Duration) *PerformActionV2Params {
	return &PerformActionV2Params{
		timeout: timeout,
	}
}

// NewPerformActionV2ParamsWithContext creates a new PerformActionV2Params object
// with the ability to set a context for a request.
func NewPerformActionV2ParamsWithContext(ctx context.Context) *PerformActionV2Params {
	return &PerformActionV2Params{
		Context: ctx,
	}
}

// NewPerformActionV2ParamsWithHTTPClient creates a new PerformActionV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewPerformActionV2ParamsWithHTTPClient(client *http.Client) *PerformActionV2Params {
	return &PerformActionV2Params{
		HTTPClient: client,
	}
}

/*
PerformActionV2Params contains all the parameters to send to the API endpoint

	for the perform action v2 operation.

	Typically these are written to a http.Request.
*/
type PerformActionV2Params struct {

	/* ActionName.

	     Specify one of these actions:

	- `contain` - This action contains the host, which stops any network communications to locations other than the CrowdStrike cloud and IPs specified in your [containment policy](https://falcon.crowdstrike.com/support/documentation/11/getting-started-guide#containmentpolicy)
	- `lift_containment`: This action lifts containment on the host, which returns its network communications to normal
	- `hide_host`: This action will delete a host. After the host is deleted, no new detections for that host will be reported via UI or APIs
	- `unhide_host`: This action will restore a host. Detection reporting will resume after the host is restored
	*/
	ActionName string

	/* Body.

	     The host agent ID (AID) of the host you want to contain. Get an agent ID from a detection, the Falcon console, or the Streaming API.

	Provide the ID in JSON format with the key `ids` and the value in square brackets, such as:

	`"ids": ["123456789"]`
	*/
	Body *models.MsaEntityActionRequestV2

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the perform action v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformActionV2Params) WithDefaults() *PerformActionV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the perform action v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformActionV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the perform action v2 params
func (o *PerformActionV2Params) WithTimeout(timeout time.Duration) *PerformActionV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the perform action v2 params
func (o *PerformActionV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the perform action v2 params
func (o *PerformActionV2Params) WithContext(ctx context.Context) *PerformActionV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the perform action v2 params
func (o *PerformActionV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the perform action v2 params
func (o *PerformActionV2Params) WithHTTPClient(client *http.Client) *PerformActionV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the perform action v2 params
func (o *PerformActionV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionName adds the actionName to the perform action v2 params
func (o *PerformActionV2Params) WithActionName(actionName string) *PerformActionV2Params {
	o.SetActionName(actionName)
	return o
}

// SetActionName adds the actionName to the perform action v2 params
func (o *PerformActionV2Params) SetActionName(actionName string) {
	o.ActionName = actionName
}

// WithBody adds the body to the perform action v2 params
func (o *PerformActionV2Params) WithBody(body *models.MsaEntityActionRequestV2) *PerformActionV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the perform action v2 params
func (o *PerformActionV2Params) SetBody(body *models.MsaEntityActionRequestV2) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PerformActionV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
