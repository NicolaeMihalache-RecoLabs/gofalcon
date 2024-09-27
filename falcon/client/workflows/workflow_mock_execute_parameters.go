// Code generated by go-swagger; DO NOT EDIT.

package workflows

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

// NewWorkflowMockExecuteParams creates a new WorkflowMockExecuteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkflowMockExecuteParams() *WorkflowMockExecuteParams {
	return &WorkflowMockExecuteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowMockExecuteParamsWithTimeout creates a new WorkflowMockExecuteParams object
// with the ability to set a timeout on a request.
func NewWorkflowMockExecuteParamsWithTimeout(timeout time.Duration) *WorkflowMockExecuteParams {
	return &WorkflowMockExecuteParams{
		timeout: timeout,
	}
}

// NewWorkflowMockExecuteParamsWithContext creates a new WorkflowMockExecuteParams object
// with the ability to set a context for a request.
func NewWorkflowMockExecuteParamsWithContext(ctx context.Context) *WorkflowMockExecuteParams {
	return &WorkflowMockExecuteParams{
		Context: ctx,
	}
}

// NewWorkflowMockExecuteParamsWithHTTPClient creates a new WorkflowMockExecuteParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkflowMockExecuteParamsWithHTTPClient(client *http.Client) *WorkflowMockExecuteParams {
	return &WorkflowMockExecuteParams{
		HTTPClient: client,
	}
}

/*
WorkflowMockExecuteParams contains all the parameters to send to the API endpoint

	for the workflow mock execute operation.

	Typically these are written to a http.Request.
*/
type WorkflowMockExecuteParams struct {

	// Body.
	Body *models.ModelsMockExecutionCreateRequestV1

	/* DefinitionID.

	   Definition ID to execute, either a name or an ID, or the definition itself in the request body, can be specified.
	*/
	DefinitionID *string

	/* Depth.

	   Used to record the execution depth to help limit execution loops when a workflow triggers another. The maximum depth is 4.
	*/
	Depth *int64

	/* ExecutionCid.

	   CID(s) to execute on. This can be a child if this is a flight control enabled definition. If unset the definition CID is used.
	*/
	ExecutionCid []string

	/* Key.

	   Key used to help deduplicate executions, if unset a new UUID is used
	*/
	Key *string

	/* Name.

	   Workflow name to execute, either a name or an ID, or the definition itself in the request body, can be specified.
	*/
	Name *string

	/* SourceEventURL.

	   Used to record a URL to the source that led to triggering this workflow
	*/
	SourceEventURL *string

	/* ValidateOnly.

	   When enabled, prevents execution after validating mocks against definition
	*/
	ValidateOnly *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the workflow mock execute params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowMockExecuteParams) WithDefaults() *WorkflowMockExecuteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workflow mock execute params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowMockExecuteParams) SetDefaults() {
	var (
		validateOnlyDefault = bool(false)
	)

	val := WorkflowMockExecuteParams{
		ValidateOnly: &validateOnlyDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the workflow mock execute params
func (o *WorkflowMockExecuteParams) WithTimeout(timeout time.Duration) *WorkflowMockExecuteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow mock execute params
func (o *WorkflowMockExecuteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow mock execute params
func (o *WorkflowMockExecuteParams) WithContext(ctx context.Context) *WorkflowMockExecuteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow mock execute params
func (o *WorkflowMockExecuteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow mock execute params
func (o *WorkflowMockExecuteParams) WithHTTPClient(client *http.Client) *WorkflowMockExecuteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow mock execute params
func (o *WorkflowMockExecuteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the workflow mock execute params
func (o *WorkflowMockExecuteParams) WithBody(body *models.ModelsMockExecutionCreateRequestV1) *WorkflowMockExecuteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the workflow mock execute params
func (o *WorkflowMockExecuteParams) SetBody(body *models.ModelsMockExecutionCreateRequestV1) {
	o.Body = body
}

// WithDefinitionID adds the definitionID to the workflow mock execute params
func (o *WorkflowMockExecuteParams) WithDefinitionID(definitionID *string) *WorkflowMockExecuteParams {
	o.SetDefinitionID(definitionID)
	return o
}

// SetDefinitionID adds the definitionId to the workflow mock execute params
func (o *WorkflowMockExecuteParams) SetDefinitionID(definitionID *string) {
	o.DefinitionID = definitionID
}

// WithDepth adds the depth to the workflow mock execute params
func (o *WorkflowMockExecuteParams) WithDepth(depth *int64) *WorkflowMockExecuteParams {
	o.SetDepth(depth)
	return o
}

// SetDepth adds the depth to the workflow mock execute params
func (o *WorkflowMockExecuteParams) SetDepth(depth *int64) {
	o.Depth = depth
}

// WithExecutionCid adds the executionCid to the workflow mock execute params
func (o *WorkflowMockExecuteParams) WithExecutionCid(executionCid []string) *WorkflowMockExecuteParams {
	o.SetExecutionCid(executionCid)
	return o
}

// SetExecutionCid adds the executionCid to the workflow mock execute params
func (o *WorkflowMockExecuteParams) SetExecutionCid(executionCid []string) {
	o.ExecutionCid = executionCid
}

// WithKey adds the key to the workflow mock execute params
func (o *WorkflowMockExecuteParams) WithKey(key *string) *WorkflowMockExecuteParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the workflow mock execute params
func (o *WorkflowMockExecuteParams) SetKey(key *string) {
	o.Key = key
}

// WithName adds the name to the workflow mock execute params
func (o *WorkflowMockExecuteParams) WithName(name *string) *WorkflowMockExecuteParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the workflow mock execute params
func (o *WorkflowMockExecuteParams) SetName(name *string) {
	o.Name = name
}

// WithSourceEventURL adds the sourceEventURL to the workflow mock execute params
func (o *WorkflowMockExecuteParams) WithSourceEventURL(sourceEventURL *string) *WorkflowMockExecuteParams {
	o.SetSourceEventURL(sourceEventURL)
	return o
}

// SetSourceEventURL adds the sourceEventUrl to the workflow mock execute params
func (o *WorkflowMockExecuteParams) SetSourceEventURL(sourceEventURL *string) {
	o.SourceEventURL = sourceEventURL
}

// WithValidateOnly adds the validateOnly to the workflow mock execute params
func (o *WorkflowMockExecuteParams) WithValidateOnly(validateOnly *bool) *WorkflowMockExecuteParams {
	o.SetValidateOnly(validateOnly)
	return o
}

// SetValidateOnly adds the validateOnly to the workflow mock execute params
func (o *WorkflowMockExecuteParams) SetValidateOnly(validateOnly *bool) {
	o.ValidateOnly = validateOnly
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowMockExecuteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.DefinitionID != nil {

		// query param definition_id
		var qrDefinitionID string

		if o.DefinitionID != nil {
			qrDefinitionID = *o.DefinitionID
		}
		qDefinitionID := qrDefinitionID
		if qDefinitionID != "" {

			if err := r.SetQueryParam("definition_id", qDefinitionID); err != nil {
				return err
			}
		}
	}

	if o.Depth != nil {

		// query param depth
		var qrDepth int64

		if o.Depth != nil {
			qrDepth = *o.Depth
		}
		qDepth := swag.FormatInt64(qrDepth)
		if qDepth != "" {

			if err := r.SetQueryParam("depth", qDepth); err != nil {
				return err
			}
		}
	}

	if o.ExecutionCid != nil {

		// binding items for execution_cid
		joinedExecutionCid := o.bindParamExecutionCid(reg)

		// query array param execution_cid
		if err := r.SetQueryParam("execution_cid", joinedExecutionCid...); err != nil {
			return err
		}
	}

	if o.Key != nil {

		// query param key
		var qrKey string

		if o.Key != nil {
			qrKey = *o.Key
		}
		qKey := qrKey
		if qKey != "" {

			if err := r.SetQueryParam("key", qKey); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.SourceEventURL != nil {

		// query param source_event_url
		var qrSourceEventURL string

		if o.SourceEventURL != nil {
			qrSourceEventURL = *o.SourceEventURL
		}
		qSourceEventURL := qrSourceEventURL
		if qSourceEventURL != "" {

			if err := r.SetQueryParam("source_event_url", qSourceEventURL); err != nil {
				return err
			}
		}
	}

	if o.ValidateOnly != nil {

		// query param validate_only
		var qrValidateOnly bool

		if o.ValidateOnly != nil {
			qrValidateOnly = *o.ValidateOnly
		}
		qValidateOnly := swag.FormatBool(qrValidateOnly)
		if qValidateOnly != "" {

			if err := r.SetQueryParam("validate_only", qValidateOnly); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamWorkflowMockExecute binds the parameter execution_cid
func (o *WorkflowMockExecuteParams) bindParamExecutionCid(formats strfmt.Registry) []string {
	executionCidIR := o.ExecutionCid

	var executionCidIC []string
	for _, executionCidIIR := range executionCidIR { // explode []string

		executionCidIIV := executionCidIIR // string as string
		executionCidIC = append(executionCidIC, executionCidIIV)
	}

	// items.CollectionFormat: "csv"
	executionCidIS := swag.JoinByFormat(executionCidIC, "csv")

	return executionCidIS
}
