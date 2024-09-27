// Code generated by go-swagger; DO NOT EDIT.

package incidents

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
	"github.com/go-openapi/swag"

	"github.com/NicolaeMihalache-RecoLab/gofalcon/falcon/models"
)

// NewPerformIncidentActionParams creates a new PerformIncidentActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformIncidentActionParams() *PerformIncidentActionParams {
	return &PerformIncidentActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformIncidentActionParamsWithTimeout creates a new PerformIncidentActionParams object
// with the ability to set a timeout on a request.
func NewPerformIncidentActionParamsWithTimeout(timeout time.Duration) *PerformIncidentActionParams {
	return &PerformIncidentActionParams{
		timeout: timeout,
	}
}

// NewPerformIncidentActionParamsWithContext creates a new PerformIncidentActionParams object
// with the ability to set a context for a request.
func NewPerformIncidentActionParamsWithContext(ctx context.Context) *PerformIncidentActionParams {
	return &PerformIncidentActionParams{
		Context: ctx,
	}
}

// NewPerformIncidentActionParamsWithHTTPClient creates a new PerformIncidentActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformIncidentActionParamsWithHTTPClient(client *http.Client) *PerformIncidentActionParams {
	return &PerformIncidentActionParams{
		HTTPClient: client,
	}
}

/*
PerformIncidentActionParams contains all the parameters to send to the API endpoint

	for the perform incident action operation.

	Typically these are written to a http.Request.
*/
type PerformIncidentActionParams struct {

	/* Body.

	   Incident Update request body containing minimum 1 and maximum 5000 Incident ID(s) and action param(s) to be performed action against.
	*/
	Body *models.DomainEntityActionRequest

	/* OverwriteDetects.

	   If true and update-detects is true, the assigned-to-uuid or status for ALL detections associated with the incident(s) will be overwritten. If false, only detects that have default values for assigned-to-uuid and/or status will be updated. Defaults to false. Ignored if 'update-detects' is missing or false.
	*/
	OverwriteDetects *bool

	/* UpdateDetects.

	   If true, update assigned-to-uuid and or status of detections associated with the incident(s). Defaults to false
	*/
	UpdateDetects *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the perform incident action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformIncidentActionParams) WithDefaults() *PerformIncidentActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the perform incident action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformIncidentActionParams) SetDefaults() {
	var (
		overwriteDetectsDefault = bool(false)

		updateDetectsDefault = bool(false)
	)

	val := PerformIncidentActionParams{
		OverwriteDetects: &overwriteDetectsDefault,
		UpdateDetects:    &updateDetectsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the perform incident action params
func (o *PerformIncidentActionParams) WithTimeout(timeout time.Duration) *PerformIncidentActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the perform incident action params
func (o *PerformIncidentActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the perform incident action params
func (o *PerformIncidentActionParams) WithContext(ctx context.Context) *PerformIncidentActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the perform incident action params
func (o *PerformIncidentActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the perform incident action params
func (o *PerformIncidentActionParams) WithHTTPClient(client *http.Client) *PerformIncidentActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the perform incident action params
func (o *PerformIncidentActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the perform incident action params
func (o *PerformIncidentActionParams) WithBody(body *models.DomainEntityActionRequest) *PerformIncidentActionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the perform incident action params
func (o *PerformIncidentActionParams) SetBody(body *models.DomainEntityActionRequest) {
	o.Body = body
}

// WithOverwriteDetects adds the overwriteDetects to the perform incident action params
func (o *PerformIncidentActionParams) WithOverwriteDetects(overwriteDetects *bool) *PerformIncidentActionParams {
	o.SetOverwriteDetects(overwriteDetects)
	return o
}

// SetOverwriteDetects adds the overwriteDetects to the perform incident action params
func (o *PerformIncidentActionParams) SetOverwriteDetects(overwriteDetects *bool) {
	o.OverwriteDetects = overwriteDetects
}

// WithUpdateDetects adds the updateDetects to the perform incident action params
func (o *PerformIncidentActionParams) WithUpdateDetects(updateDetects *bool) *PerformIncidentActionParams {
	o.SetUpdateDetects(updateDetects)
	return o
}

// SetUpdateDetects adds the updateDetects to the perform incident action params
func (o *PerformIncidentActionParams) SetUpdateDetects(updateDetects *bool) {
	o.UpdateDetects = updateDetects
}

// WriteToRequest writes these params to a swagger request
func (o *PerformIncidentActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.OverwriteDetects != nil {

		// query param overwrite_detects
		var qrOverwriteDetects bool

		if o.OverwriteDetects != nil {
			qrOverwriteDetects = *o.OverwriteDetects
		}
		qOverwriteDetects := swag.FormatBool(qrOverwriteDetects)
		if qOverwriteDetects != "" {

			if err := r.SetQueryParam("overwrite_detects", qOverwriteDetects); err != nil {
				return err
			}
		}
	}

	if o.UpdateDetects != nil {

		// query param update_detects
		var qrUpdateDetects bool

		if o.UpdateDetects != nil {
			qrUpdateDetects = *o.UpdateDetects
		}
		qUpdateDetects := swag.FormatBool(qrUpdateDetects)
		if qUpdateDetects != "" {

			if err := r.SetQueryParam("update_detects", qUpdateDetects); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
