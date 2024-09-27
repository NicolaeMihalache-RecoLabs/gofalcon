// Code generated by go-swagger; DO NOT EDIT.

package mssp

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

// NewDeleteUserGroupMembersParams creates a new DeleteUserGroupMembersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUserGroupMembersParams() *DeleteUserGroupMembersParams {
	return &DeleteUserGroupMembersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserGroupMembersParamsWithTimeout creates a new DeleteUserGroupMembersParams object
// with the ability to set a timeout on a request.
func NewDeleteUserGroupMembersParamsWithTimeout(timeout time.Duration) *DeleteUserGroupMembersParams {
	return &DeleteUserGroupMembersParams{
		timeout: timeout,
	}
}

// NewDeleteUserGroupMembersParamsWithContext creates a new DeleteUserGroupMembersParams object
// with the ability to set a context for a request.
func NewDeleteUserGroupMembersParamsWithContext(ctx context.Context) *DeleteUserGroupMembersParams {
	return &DeleteUserGroupMembersParams{
		Context: ctx,
	}
}

// NewDeleteUserGroupMembersParamsWithHTTPClient creates a new DeleteUserGroupMembersParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUserGroupMembersParamsWithHTTPClient(client *http.Client) *DeleteUserGroupMembersParams {
	return &DeleteUserGroupMembersParams{
		HTTPClient: client,
	}
}

/*
DeleteUserGroupMembersParams contains all the parameters to send to the API endpoint

	for the delete user group members operation.

	Typically these are written to a http.Request.
*/
type DeleteUserGroupMembersParams struct {

	/* Body.

	   Both 'user_group_id' and 'user_uuids' fields are required.
	*/
	Body *models.DomainUserGroupMembersRequestV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete user group members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserGroupMembersParams) WithDefaults() *DeleteUserGroupMembersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete user group members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserGroupMembersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete user group members params
func (o *DeleteUserGroupMembersParams) WithTimeout(timeout time.Duration) *DeleteUserGroupMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user group members params
func (o *DeleteUserGroupMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user group members params
func (o *DeleteUserGroupMembersParams) WithContext(ctx context.Context) *DeleteUserGroupMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user group members params
func (o *DeleteUserGroupMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user group members params
func (o *DeleteUserGroupMembersParams) WithHTTPClient(client *http.Client) *DeleteUserGroupMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user group members params
func (o *DeleteUserGroupMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete user group members params
func (o *DeleteUserGroupMembersParams) WithBody(body *models.DomainUserGroupMembersRequestV1) *DeleteUserGroupMembersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete user group members params
func (o *DeleteUserGroupMembersParams) SetBody(body *models.DomainUserGroupMembersRequestV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserGroupMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
