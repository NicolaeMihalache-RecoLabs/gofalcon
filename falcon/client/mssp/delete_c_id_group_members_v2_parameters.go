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

// NewDeleteCIDGroupMembersV2Params creates a new DeleteCIDGroupMembersV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCIDGroupMembersV2Params() *DeleteCIDGroupMembersV2Params {
	return &DeleteCIDGroupMembersV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCIDGroupMembersV2ParamsWithTimeout creates a new DeleteCIDGroupMembersV2Params object
// with the ability to set a timeout on a request.
func NewDeleteCIDGroupMembersV2ParamsWithTimeout(timeout time.Duration) *DeleteCIDGroupMembersV2Params {
	return &DeleteCIDGroupMembersV2Params{
		timeout: timeout,
	}
}

// NewDeleteCIDGroupMembersV2ParamsWithContext creates a new DeleteCIDGroupMembersV2Params object
// with the ability to set a context for a request.
func NewDeleteCIDGroupMembersV2ParamsWithContext(ctx context.Context) *DeleteCIDGroupMembersV2Params {
	return &DeleteCIDGroupMembersV2Params{
		Context: ctx,
	}
}

// NewDeleteCIDGroupMembersV2ParamsWithHTTPClient creates a new DeleteCIDGroupMembersV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCIDGroupMembersV2ParamsWithHTTPClient(client *http.Client) *DeleteCIDGroupMembersV2Params {
	return &DeleteCIDGroupMembersV2Params{
		HTTPClient: client,
	}
}

/*
DeleteCIDGroupMembersV2Params contains all the parameters to send to the API endpoint

	for the delete c ID group members v2 operation.

	Typically these are written to a http.Request.
*/
type DeleteCIDGroupMembersV2Params struct {

	/* Body.

	   Both 'cid_group_id' and 'cids' fields are required.
	*/
	Body *models.DomainCIDGroupMembersRequestV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete c ID group members v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCIDGroupMembersV2Params) WithDefaults() *DeleteCIDGroupMembersV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete c ID group members v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCIDGroupMembersV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete c ID group members v2 params
func (o *DeleteCIDGroupMembersV2Params) WithTimeout(timeout time.Duration) *DeleteCIDGroupMembersV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete c ID group members v2 params
func (o *DeleteCIDGroupMembersV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete c ID group members v2 params
func (o *DeleteCIDGroupMembersV2Params) WithContext(ctx context.Context) *DeleteCIDGroupMembersV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete c ID group members v2 params
func (o *DeleteCIDGroupMembersV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete c ID group members v2 params
func (o *DeleteCIDGroupMembersV2Params) WithHTTPClient(client *http.Client) *DeleteCIDGroupMembersV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete c ID group members v2 params
func (o *DeleteCIDGroupMembersV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete c ID group members v2 params
func (o *DeleteCIDGroupMembersV2Params) WithBody(body *models.DomainCIDGroupMembersRequestV1) *DeleteCIDGroupMembersV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete c ID group members v2 params
func (o *DeleteCIDGroupMembersV2Params) SetBody(body *models.DomainCIDGroupMembersRequestV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCIDGroupMembersV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
