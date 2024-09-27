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

// NewUpdateRuleGroupsParams creates a new UpdateRuleGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRuleGroupsParams() *UpdateRuleGroupsParams {
	return &UpdateRuleGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRuleGroupsParamsWithTimeout creates a new UpdateRuleGroupsParams object
// with the ability to set a timeout on a request.
func NewUpdateRuleGroupsParamsWithTimeout(timeout time.Duration) *UpdateRuleGroupsParams {
	return &UpdateRuleGroupsParams{
		timeout: timeout,
	}
}

// NewUpdateRuleGroupsParamsWithContext creates a new UpdateRuleGroupsParams object
// with the ability to set a context for a request.
func NewUpdateRuleGroupsParamsWithContext(ctx context.Context) *UpdateRuleGroupsParams {
	return &UpdateRuleGroupsParams{
		Context: ctx,
	}
}

// NewUpdateRuleGroupsParamsWithHTTPClient creates a new UpdateRuleGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRuleGroupsParamsWithHTTPClient(client *http.Client) *UpdateRuleGroupsParams {
	return &UpdateRuleGroupsParams{
		HTTPClient: client,
	}
}

/*
UpdateRuleGroupsParams contains all the parameters to send to the API endpoint

	for the update rule groups operation.

	Typically these are written to a http.Request.
*/
type UpdateRuleGroupsParams struct {

	/* Body.

	    Enables updates to the following fields for an existing rule group.

	* `id` of the rule group to update.

	* `name` must be between 1 and 100 characters.

	* `description` can be between 0 and 500 characters.

	* `type` may not be modified after the rule group is created.

	Note: rules are added/removed from rule groups using their dedicated end-points.
	*/
	Body *models.RulegroupsUpdateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update rule groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRuleGroupsParams) WithDefaults() *UpdateRuleGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update rule groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRuleGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update rule groups params
func (o *UpdateRuleGroupsParams) WithTimeout(timeout time.Duration) *UpdateRuleGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update rule groups params
func (o *UpdateRuleGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update rule groups params
func (o *UpdateRuleGroupsParams) WithContext(ctx context.Context) *UpdateRuleGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update rule groups params
func (o *UpdateRuleGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update rule groups params
func (o *UpdateRuleGroupsParams) WithHTTPClient(client *http.Client) *UpdateRuleGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update rule groups params
func (o *UpdateRuleGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update rule groups params
func (o *UpdateRuleGroupsParams) WithBody(body *models.RulegroupsUpdateRequest) *UpdateRuleGroupsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update rule groups params
func (o *UpdateRuleGroupsParams) SetBody(body *models.RulegroupsUpdateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRuleGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
