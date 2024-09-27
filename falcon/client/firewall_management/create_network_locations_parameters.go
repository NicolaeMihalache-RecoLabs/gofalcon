// Code generated by go-swagger; DO NOT EDIT.

package firewall_management

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

// NewCreateNetworkLocationsParams creates a new CreateNetworkLocationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNetworkLocationsParams() *CreateNetworkLocationsParams {
	return &CreateNetworkLocationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNetworkLocationsParamsWithTimeout creates a new CreateNetworkLocationsParams object
// with the ability to set a timeout on a request.
func NewCreateNetworkLocationsParamsWithTimeout(timeout time.Duration) *CreateNetworkLocationsParams {
	return &CreateNetworkLocationsParams{
		timeout: timeout,
	}
}

// NewCreateNetworkLocationsParamsWithContext creates a new CreateNetworkLocationsParams object
// with the ability to set a context for a request.
func NewCreateNetworkLocationsParamsWithContext(ctx context.Context) *CreateNetworkLocationsParams {
	return &CreateNetworkLocationsParams{
		Context: ctx,
	}
}

// NewCreateNetworkLocationsParamsWithHTTPClient creates a new CreateNetworkLocationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNetworkLocationsParamsWithHTTPClient(client *http.Client) *CreateNetworkLocationsParams {
	return &CreateNetworkLocationsParams{
		HTTPClient: client,
	}
}

/*
CreateNetworkLocationsParams contains all the parameters to send to the API endpoint

	for the create network locations operation.

	Typically these are written to a http.Request.
*/
type CreateNetworkLocationsParams struct {

	/* AddFwRules.

	   A boolean to determine whether the cloned location needs to be added to the same firewall rules that original location is added to.
	*/
	AddFwRules *bool

	// Body.
	Body *models.FwmgrAPINetworkLocationCreateRequestV1

	/* CloneID.

	   A network location ID from which to copy location. If this is provided then the body of the request is ignored.
	*/
	CloneID *string

	/* Comment.

	   Audit log comment for this action
	*/
	Comment *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create network locations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkLocationsParams) WithDefaults() *CreateNetworkLocationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create network locations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkLocationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create network locations params
func (o *CreateNetworkLocationsParams) WithTimeout(timeout time.Duration) *CreateNetworkLocationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create network locations params
func (o *CreateNetworkLocationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create network locations params
func (o *CreateNetworkLocationsParams) WithContext(ctx context.Context) *CreateNetworkLocationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create network locations params
func (o *CreateNetworkLocationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create network locations params
func (o *CreateNetworkLocationsParams) WithHTTPClient(client *http.Client) *CreateNetworkLocationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create network locations params
func (o *CreateNetworkLocationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddFwRules adds the addFwRules to the create network locations params
func (o *CreateNetworkLocationsParams) WithAddFwRules(addFwRules *bool) *CreateNetworkLocationsParams {
	o.SetAddFwRules(addFwRules)
	return o
}

// SetAddFwRules adds the addFwRules to the create network locations params
func (o *CreateNetworkLocationsParams) SetAddFwRules(addFwRules *bool) {
	o.AddFwRules = addFwRules
}

// WithBody adds the body to the create network locations params
func (o *CreateNetworkLocationsParams) WithBody(body *models.FwmgrAPINetworkLocationCreateRequestV1) *CreateNetworkLocationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create network locations params
func (o *CreateNetworkLocationsParams) SetBody(body *models.FwmgrAPINetworkLocationCreateRequestV1) {
	o.Body = body
}

// WithCloneID adds the cloneID to the create network locations params
func (o *CreateNetworkLocationsParams) WithCloneID(cloneID *string) *CreateNetworkLocationsParams {
	o.SetCloneID(cloneID)
	return o
}

// SetCloneID adds the cloneId to the create network locations params
func (o *CreateNetworkLocationsParams) SetCloneID(cloneID *string) {
	o.CloneID = cloneID
}

// WithComment adds the comment to the create network locations params
func (o *CreateNetworkLocationsParams) WithComment(comment *string) *CreateNetworkLocationsParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the create network locations params
func (o *CreateNetworkLocationsParams) SetComment(comment *string) {
	o.Comment = comment
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNetworkLocationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AddFwRules != nil {

		// query param add_fw_rules
		var qrAddFwRules bool

		if o.AddFwRules != nil {
			qrAddFwRules = *o.AddFwRules
		}
		qAddFwRules := swag.FormatBool(qrAddFwRules)
		if qAddFwRules != "" {

			if err := r.SetQueryParam("add_fw_rules", qAddFwRules); err != nil {
				return err
			}
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.CloneID != nil {

		// query param clone_id
		var qrCloneID string

		if o.CloneID != nil {
			qrCloneID = *o.CloneID
		}
		qCloneID := qrCloneID
		if qCloneID != "" {

			if err := r.SetQueryParam("clone_id", qCloneID); err != nil {
				return err
			}
		}
	}

	if o.Comment != nil {

		// query param comment
		var qrComment string

		if o.Comment != nil {
			qrComment = *o.Comment
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
