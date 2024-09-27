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

// NewDeletedRolesParams creates a new DeletedRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletedRolesParams() *DeletedRolesParams {
	return &DeletedRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletedRolesParamsWithTimeout creates a new DeletedRolesParams object
// with the ability to set a timeout on a request.
func NewDeletedRolesParamsWithTimeout(timeout time.Duration) *DeletedRolesParams {
	return &DeletedRolesParams{
		timeout: timeout,
	}
}

// NewDeletedRolesParamsWithContext creates a new DeletedRolesParams object
// with the ability to set a context for a request.
func NewDeletedRolesParamsWithContext(ctx context.Context) *DeletedRolesParams {
	return &DeletedRolesParams{
		Context: ctx,
	}
}

// NewDeletedRolesParamsWithHTTPClient creates a new DeletedRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletedRolesParamsWithHTTPClient(client *http.Client) *DeletedRolesParams {
	return &DeletedRolesParams{
		HTTPClient: client,
	}
}

/*
DeletedRolesParams contains all the parameters to send to the API endpoint

	for the deleted roles operation.

	Typically these are written to a http.Request.
*/
type DeletedRolesParams struct {

	/* Body.

	   'user_group_id' and 'cid_group_id' fields are required. 'role_ids' field is optional. Remaining fields are ignored.
	*/
	Body *models.DomainMSSPRoleRequestV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the deleted roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletedRolesParams) WithDefaults() *DeletedRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the deleted roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletedRolesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the deleted roles params
func (o *DeletedRolesParams) WithTimeout(timeout time.Duration) *DeletedRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deleted roles params
func (o *DeletedRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deleted roles params
func (o *DeletedRolesParams) WithContext(ctx context.Context) *DeletedRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deleted roles params
func (o *DeletedRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deleted roles params
func (o *DeletedRolesParams) WithHTTPClient(client *http.Client) *DeletedRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deleted roles params
func (o *DeletedRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the deleted roles params
func (o *DeletedRolesParams) WithBody(body *models.DomainMSSPRoleRequestV1) *DeletedRolesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the deleted roles params
func (o *DeletedRolesParams) SetBody(body *models.DomainMSSPRoleRequestV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeletedRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
