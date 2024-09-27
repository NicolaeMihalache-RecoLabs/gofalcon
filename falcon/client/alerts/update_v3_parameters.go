// Code generated by go-swagger; DO NOT EDIT.

package alerts

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

// NewUpdateV3Params creates a new UpdateV3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateV3Params() *UpdateV3Params {
	return &UpdateV3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateV3ParamsWithTimeout creates a new UpdateV3Params object
// with the ability to set a timeout on a request.
func NewUpdateV3ParamsWithTimeout(timeout time.Duration) *UpdateV3Params {
	return &UpdateV3Params{
		timeout: timeout,
	}
}

// NewUpdateV3ParamsWithContext creates a new UpdateV3Params object
// with the ability to set a context for a request.
func NewUpdateV3ParamsWithContext(ctx context.Context) *UpdateV3Params {
	return &UpdateV3Params{
		Context: ctx,
	}
}

// NewUpdateV3ParamsWithHTTPClient creates a new UpdateV3Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateV3ParamsWithHTTPClient(client *http.Client) *UpdateV3Params {
	return &UpdateV3Params{
		HTTPClient: client,
	}
}

/*
UpdateV3Params contains all the parameters to send to the API endpoint

	for the update v3 operation.

	Typically these are written to a http.Request.
*/
type UpdateV3Params struct {

	/* Body.

	   request body takes a list of action parameter request that is applied against all "composite_ids" provided
	*/
	Body *models.DetectsapiPatchEntitiesAlertsV3Request

	/* IncludeHidden.

	   allows previously hidden alerts to be retrieved

	   Default: true
	*/
	IncludeHidden *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update v3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateV3Params) WithDefaults() *UpdateV3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update v3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateV3Params) SetDefaults() {
	var (
		includeHiddenDefault = bool(true)
	)

	val := UpdateV3Params{
		IncludeHidden: &includeHiddenDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update v3 params
func (o *UpdateV3Params) WithTimeout(timeout time.Duration) *UpdateV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update v3 params
func (o *UpdateV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update v3 params
func (o *UpdateV3Params) WithContext(ctx context.Context) *UpdateV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update v3 params
func (o *UpdateV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update v3 params
func (o *UpdateV3Params) WithHTTPClient(client *http.Client) *UpdateV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update v3 params
func (o *UpdateV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update v3 params
func (o *UpdateV3Params) WithBody(body *models.DetectsapiPatchEntitiesAlertsV3Request) *UpdateV3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update v3 params
func (o *UpdateV3Params) SetBody(body *models.DetectsapiPatchEntitiesAlertsV3Request) {
	o.Body = body
}

// WithIncludeHidden adds the includeHidden to the update v3 params
func (o *UpdateV3Params) WithIncludeHidden(includeHidden *bool) *UpdateV3Params {
	o.SetIncludeHidden(includeHidden)
	return o
}

// SetIncludeHidden adds the includeHidden to the update v3 params
func (o *UpdateV3Params) SetIncludeHidden(includeHidden *bool) {
	o.IncludeHidden = includeHidden
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.IncludeHidden != nil {

		// query param include_hidden
		var qrIncludeHidden bool

		if o.IncludeHidden != nil {
			qrIncludeHidden = *o.IncludeHidden
		}
		qIncludeHidden := swag.FormatBool(qrIncludeHidden)
		if qIncludeHidden != "" {

			if err := r.SetQueryParam("include_hidden", qIncludeHidden); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
