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

// NewUpdateScheduledExclusionsParams creates a new UpdateScheduledExclusionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateScheduledExclusionsParams() *UpdateScheduledExclusionsParams {
	return &UpdateScheduledExclusionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateScheduledExclusionsParamsWithTimeout creates a new UpdateScheduledExclusionsParams object
// with the ability to set a timeout on a request.
func NewUpdateScheduledExclusionsParamsWithTimeout(timeout time.Duration) *UpdateScheduledExclusionsParams {
	return &UpdateScheduledExclusionsParams{
		timeout: timeout,
	}
}

// NewUpdateScheduledExclusionsParamsWithContext creates a new UpdateScheduledExclusionsParams object
// with the ability to set a context for a request.
func NewUpdateScheduledExclusionsParamsWithContext(ctx context.Context) *UpdateScheduledExclusionsParams {
	return &UpdateScheduledExclusionsParams{
		Context: ctx,
	}
}

// NewUpdateScheduledExclusionsParamsWithHTTPClient creates a new UpdateScheduledExclusionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateScheduledExclusionsParamsWithHTTPClient(client *http.Client) *UpdateScheduledExclusionsParams {
	return &UpdateScheduledExclusionsParams{
		HTTPClient: client,
	}
}

/*
UpdateScheduledExclusionsParams contains all the parameters to send to the API endpoint

	for the update scheduled exclusions operation.

	Typically these are written to a http.Request.
*/
type UpdateScheduledExclusionsParams struct {

	/* Body.

	    Update an existing scheduled exclusion for the specified policy.



	* `policy_id` to add the scheduled exclusion to.

	* `name` must be between 1 and 100 characters.

	* `description` can be between 0 and 500 characters.

	* `users` can be between 0 and 500 characters representing a comma separated list of user to exclude their changes.

	   *  admin* excludes changes made by all usernames that begin with admin. Falon GLOB syntax is supported.

	* `processes` can be between 0 and 500 characters representing a comma separated list of processes to exclude their changes.

	   * **\RunMe.exe or *[*]/RunMe.sh excludes changes made by RunMe.exe or RunMe.sh in any location.

	* `schedule_start` must be provided to indicate the start of the schedule. This date/time must be an rfc3339 formatted string  https://datatracker.ietf.org/doc/html/rfc3339.

	* `schedule_end` optionally provided to indicate the end of the schedule. This date/time must be an rfc3339 formatted string  https://datatracker.ietf.org/doc/html/rfc3339.

	* `timezone`  must be provided to indicate the TimeZone Name set for the provided `scheduled_start` and `scheduled_end` values. See https://en.wikipedia.org/wiki/List_of_tz_database_time_zones.

	* `repeated` optionally provided to indicate that the exclusion is applied repeatedly within the `scheduled_start` and `scheduled_end` time.

	   * `start_time` must be the hour(00-23) and minute(00-59) of the day formatted as `HH:MM`. Required if `all_day` is not set to `true`

	   * `end_time` must be the hour(00-23) and minute(00-59) of the day formatted as `HH:MM`. Required if `all_day` is not set to `true`

	   * `all_day` must be `true` or `false` to indicate the exclusion is applied all day.

	   * `frequency` must be one of `daily`, `weekly` or `monthly`.

	   * `occurrence` must be one of the following when `frequency` is set to `monthly`:

	     * `1st`, `2nd`, `3rd`, `4th` or `Last` represents the week.

	     * `Days` represents specific calendar days.

	   * `weekly_days` must be one or more of `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday` or `Sunday` when `frequency` is set to `weekly` or `frequency` is set to `monthly` and `occurrence` is NOT set to `Days`.

	   * `monthly_days` must be set to one or more calendar days, between 1 and 31  when `frequency` is set to `monthly` and `occurrence` is set to `Days`.
	*/
	Body *models.ScheduledexclusionsUpdateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update scheduled exclusions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateScheduledExclusionsParams) WithDefaults() *UpdateScheduledExclusionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update scheduled exclusions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateScheduledExclusionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update scheduled exclusions params
func (o *UpdateScheduledExclusionsParams) WithTimeout(timeout time.Duration) *UpdateScheduledExclusionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update scheduled exclusions params
func (o *UpdateScheduledExclusionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update scheduled exclusions params
func (o *UpdateScheduledExclusionsParams) WithContext(ctx context.Context) *UpdateScheduledExclusionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update scheduled exclusions params
func (o *UpdateScheduledExclusionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update scheduled exclusions params
func (o *UpdateScheduledExclusionsParams) WithHTTPClient(client *http.Client) *UpdateScheduledExclusionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update scheduled exclusions params
func (o *UpdateScheduledExclusionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update scheduled exclusions params
func (o *UpdateScheduledExclusionsParams) WithBody(body *models.ScheduledexclusionsUpdateRequest) *UpdateScheduledExclusionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update scheduled exclusions params
func (o *UpdateScheduledExclusionsParams) SetBody(body *models.ScheduledexclusionsUpdateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateScheduledExclusionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
