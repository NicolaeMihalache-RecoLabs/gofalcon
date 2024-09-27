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
	"github.com/go-openapi/swag"

	"github.com/NicolaeMihalache-RecoLab/gofalcon/falcon/models"
)

// NewBatchRefreshSessionsParams creates a new BatchRefreshSessionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBatchRefreshSessionsParams() *BatchRefreshSessionsParams {
	return &BatchRefreshSessionsParams{
		requestTimeout: cr.DefaultTimeout,
	}
}

// NewBatchRefreshSessionsParamsWithTimeout creates a new BatchRefreshSessionsParams object
// with the ability to set a timeout on a request.
func NewBatchRefreshSessionsParamsWithTimeout(timeout time.Duration) *BatchRefreshSessionsParams {
	return &BatchRefreshSessionsParams{
		requestTimeout: timeout,
	}
}

// NewBatchRefreshSessionsParamsWithContext creates a new BatchRefreshSessionsParams object
// with the ability to set a context for a request.
func NewBatchRefreshSessionsParamsWithContext(ctx context.Context) *BatchRefreshSessionsParams {
	return &BatchRefreshSessionsParams{
		Context: ctx,
	}
}

// NewBatchRefreshSessionsParamsWithHTTPClient creates a new BatchRefreshSessionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewBatchRefreshSessionsParamsWithHTTPClient(client *http.Client) *BatchRefreshSessionsParams {
	return &BatchRefreshSessionsParams{
		HTTPClient: client,
	}
}

/*
BatchRefreshSessionsParams contains all the parameters to send to the API endpoint

	for the batch refresh sessions operation.

	Typically these are written to a http.Request.
*/
type BatchRefreshSessionsParams struct {

	/* Body.

	     **`batch_id`** Batch ID to execute the command on.  Received from `/real-time-response/combined/batch-init-session/v1`.
	**`hosts_to_remove`** Hosts to remove from the batch session.  Heartbeats will no longer happen on these hosts and the sessions will expire.
	*/
	Body *models.DomainBatchRefreshSessionRequest

	/* Timeout.

	   Timeout for how long to wait for the request in seconds, default timeout is 30 seconds. Maximum is 5 minutes.

	   Default: 30
	*/
	Timeout *int64

	/* TimeoutDuration.

	   Timeout duration for how long to wait for the request in duration syntax. Example, `10s`. Valid units: `ns, us, ms, s, m, h`. Maximum is 5 minutes.

	   Default: "30s"
	*/
	TimeoutDuration *string

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithDefaults hydrates default values in the batch refresh sessions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BatchRefreshSessionsParams) WithDefaults() *BatchRefreshSessionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the batch refresh sessions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BatchRefreshSessionsParams) SetDefaults() {
	var (
		timeoutDefault = int64(30)

		timeoutDurationDefault = string("30s")
	)

	val := BatchRefreshSessionsParams{
		Timeout:         &timeoutDefault,
		TimeoutDuration: &timeoutDurationDefault,
	}

	val.requestTimeout = o.requestTimeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithRequestTimeout adds the timeout to the batch refresh sessions params
func (o *BatchRefreshSessionsParams) WithRequestTimeout(timeout time.Duration) *BatchRefreshSessionsParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the batch refresh sessions params
func (o *BatchRefreshSessionsParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the batch refresh sessions params
func (o *BatchRefreshSessionsParams) WithContext(ctx context.Context) *BatchRefreshSessionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the batch refresh sessions params
func (o *BatchRefreshSessionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the batch refresh sessions params
func (o *BatchRefreshSessionsParams) WithHTTPClient(client *http.Client) *BatchRefreshSessionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the batch refresh sessions params
func (o *BatchRefreshSessionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the batch refresh sessions params
func (o *BatchRefreshSessionsParams) WithBody(body *models.DomainBatchRefreshSessionRequest) *BatchRefreshSessionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the batch refresh sessions params
func (o *BatchRefreshSessionsParams) SetBody(body *models.DomainBatchRefreshSessionRequest) {
	o.Body = body
}

// WithTimeout adds the timeout to the batch refresh sessions params
func (o *BatchRefreshSessionsParams) WithTimeout(timeout *int64) *BatchRefreshSessionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the batch refresh sessions params
func (o *BatchRefreshSessionsParams) SetTimeout(timeout *int64) {
	o.Timeout = timeout
}

// WithTimeoutDuration adds the timeoutDuration to the batch refresh sessions params
func (o *BatchRefreshSessionsParams) WithTimeoutDuration(timeoutDuration *string) *BatchRefreshSessionsParams {
	o.SetTimeoutDuration(timeoutDuration)
	return o
}

// SetTimeoutDuration adds the timeoutDuration to the batch refresh sessions params
func (o *BatchRefreshSessionsParams) SetTimeoutDuration(timeoutDuration *string) {
	o.TimeoutDuration = timeoutDuration
}

// WriteToRequest writes these params to a swagger request
func (o *BatchRefreshSessionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.requestTimeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Timeout != nil {

		// query param timeout
		var qrTimeout int64

		if o.Timeout != nil {
			qrTimeout = *o.Timeout
		}
		qTimeout := swag.FormatInt64(qrTimeout)
		if qTimeout != "" {

			if err := r.SetQueryParam("timeout", qTimeout); err != nil {
				return err
			}
		}
	}

	if o.TimeoutDuration != nil {

		// query param timeout_duration
		var qrTimeoutDuration string

		if o.TimeoutDuration != nil {
			qrTimeoutDuration = *o.TimeoutDuration
		}
		qTimeoutDuration := qrTimeoutDuration
		if qTimeoutDuration != "" {

			if err := r.SetQueryParam("timeout_duration", qTimeoutDuration); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
