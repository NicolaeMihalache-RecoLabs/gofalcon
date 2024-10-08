// Code generated by go-swagger; DO NOT EDIT.

package workflows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/NicolaeMihalache-RecoLab/gofalcon/falcon/models"
)

// WorkflowExecuteInternalReader is a Reader for the WorkflowExecuteInternal structure.
type WorkflowExecuteInternalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowExecuteInternalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowExecuteInternalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewWorkflowExecuteInternalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewWorkflowExecuteInternalForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewWorkflowExecuteInternalNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewWorkflowExecuteInternalTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewWorkflowExecuteInternalInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /workflows/entities/execute/internal/v1] WorkflowExecuteInternal", response, response.Code())
	}
}

// NewWorkflowExecuteInternalOK creates a WorkflowExecuteInternalOK with default headers values
func NewWorkflowExecuteInternalOK() *WorkflowExecuteInternalOK {
	return &WorkflowExecuteInternalOK{}
}

/*
WorkflowExecuteInternalOK describes a response with status code 200, with default header values.

OK
*/
type WorkflowExecuteInternalOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIResourceIDsResponse
}

// IsSuccess returns true when this workflow execute internal o k response has a 2xx status code
func (o *WorkflowExecuteInternalOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow execute internal o k response has a 3xx status code
func (o *WorkflowExecuteInternalOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow execute internal o k response has a 4xx status code
func (o *WorkflowExecuteInternalOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow execute internal o k response has a 5xx status code
func (o *WorkflowExecuteInternalOK) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow execute internal o k response a status code equal to that given
func (o *WorkflowExecuteInternalOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the workflow execute internal o k response
func (o *WorkflowExecuteInternalOK) Code() int {
	return 200
}

func (o *WorkflowExecuteInternalOK) Error() string {
	return fmt.Sprintf("[POST /workflows/entities/execute/internal/v1][%d] workflowExecuteInternalOK  %+v", 200, o.Payload)
}

func (o *WorkflowExecuteInternalOK) String() string {
	return fmt.Sprintf("[POST /workflows/entities/execute/internal/v1][%d] workflowExecuteInternalOK  %+v", 200, o.Payload)
}

func (o *WorkflowExecuteInternalOK) GetPayload() *models.APIResourceIDsResponse {
	return o.Payload
}

func (o *WorkflowExecuteInternalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.APIResourceIDsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWorkflowExecuteInternalBadRequest creates a WorkflowExecuteInternalBadRequest with default headers values
func NewWorkflowExecuteInternalBadRequest() *WorkflowExecuteInternalBadRequest {
	return &WorkflowExecuteInternalBadRequest{}
}

/*
WorkflowExecuteInternalBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type WorkflowExecuteInternalBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIResourceIDsResponse
}

// IsSuccess returns true when this workflow execute internal bad request response has a 2xx status code
func (o *WorkflowExecuteInternalBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this workflow execute internal bad request response has a 3xx status code
func (o *WorkflowExecuteInternalBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow execute internal bad request response has a 4xx status code
func (o *WorkflowExecuteInternalBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this workflow execute internal bad request response has a 5xx status code
func (o *WorkflowExecuteInternalBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow execute internal bad request response a status code equal to that given
func (o *WorkflowExecuteInternalBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the workflow execute internal bad request response
func (o *WorkflowExecuteInternalBadRequest) Code() int {
	return 400
}

func (o *WorkflowExecuteInternalBadRequest) Error() string {
	return fmt.Sprintf("[POST /workflows/entities/execute/internal/v1][%d] workflowExecuteInternalBadRequest  %+v", 400, o.Payload)
}

func (o *WorkflowExecuteInternalBadRequest) String() string {
	return fmt.Sprintf("[POST /workflows/entities/execute/internal/v1][%d] workflowExecuteInternalBadRequest  %+v", 400, o.Payload)
}

func (o *WorkflowExecuteInternalBadRequest) GetPayload() *models.APIResourceIDsResponse {
	return o.Payload
}

func (o *WorkflowExecuteInternalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.APIResourceIDsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWorkflowExecuteInternalForbidden creates a WorkflowExecuteInternalForbidden with default headers values
func NewWorkflowExecuteInternalForbidden() *WorkflowExecuteInternalForbidden {
	return &WorkflowExecuteInternalForbidden{}
}

/*
WorkflowExecuteInternalForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type WorkflowExecuteInternalForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this workflow execute internal forbidden response has a 2xx status code
func (o *WorkflowExecuteInternalForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this workflow execute internal forbidden response has a 3xx status code
func (o *WorkflowExecuteInternalForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow execute internal forbidden response has a 4xx status code
func (o *WorkflowExecuteInternalForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this workflow execute internal forbidden response has a 5xx status code
func (o *WorkflowExecuteInternalForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow execute internal forbidden response a status code equal to that given
func (o *WorkflowExecuteInternalForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the workflow execute internal forbidden response
func (o *WorkflowExecuteInternalForbidden) Code() int {
	return 403
}

func (o *WorkflowExecuteInternalForbidden) Error() string {
	return fmt.Sprintf("[POST /workflows/entities/execute/internal/v1][%d] workflowExecuteInternalForbidden  %+v", 403, o.Payload)
}

func (o *WorkflowExecuteInternalForbidden) String() string {
	return fmt.Sprintf("[POST /workflows/entities/execute/internal/v1][%d] workflowExecuteInternalForbidden  %+v", 403, o.Payload)
}

func (o *WorkflowExecuteInternalForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *WorkflowExecuteInternalForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWorkflowExecuteInternalNotFound creates a WorkflowExecuteInternalNotFound with default headers values
func NewWorkflowExecuteInternalNotFound() *WorkflowExecuteInternalNotFound {
	return &WorkflowExecuteInternalNotFound{}
}

/*
WorkflowExecuteInternalNotFound describes a response with status code 404, with default header values.

Not Found
*/
type WorkflowExecuteInternalNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIResourceIDsResponse
}

// IsSuccess returns true when this workflow execute internal not found response has a 2xx status code
func (o *WorkflowExecuteInternalNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this workflow execute internal not found response has a 3xx status code
func (o *WorkflowExecuteInternalNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow execute internal not found response has a 4xx status code
func (o *WorkflowExecuteInternalNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this workflow execute internal not found response has a 5xx status code
func (o *WorkflowExecuteInternalNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow execute internal not found response a status code equal to that given
func (o *WorkflowExecuteInternalNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the workflow execute internal not found response
func (o *WorkflowExecuteInternalNotFound) Code() int {
	return 404
}

func (o *WorkflowExecuteInternalNotFound) Error() string {
	return fmt.Sprintf("[POST /workflows/entities/execute/internal/v1][%d] workflowExecuteInternalNotFound  %+v", 404, o.Payload)
}

func (o *WorkflowExecuteInternalNotFound) String() string {
	return fmt.Sprintf("[POST /workflows/entities/execute/internal/v1][%d] workflowExecuteInternalNotFound  %+v", 404, o.Payload)
}

func (o *WorkflowExecuteInternalNotFound) GetPayload() *models.APIResourceIDsResponse {
	return o.Payload
}

func (o *WorkflowExecuteInternalNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.APIResourceIDsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWorkflowExecuteInternalTooManyRequests creates a WorkflowExecuteInternalTooManyRequests with default headers values
func NewWorkflowExecuteInternalTooManyRequests() *WorkflowExecuteInternalTooManyRequests {
	return &WorkflowExecuteInternalTooManyRequests{}
}

/*
WorkflowExecuteInternalTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type WorkflowExecuteInternalTooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this workflow execute internal too many requests response has a 2xx status code
func (o *WorkflowExecuteInternalTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this workflow execute internal too many requests response has a 3xx status code
func (o *WorkflowExecuteInternalTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow execute internal too many requests response has a 4xx status code
func (o *WorkflowExecuteInternalTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this workflow execute internal too many requests response has a 5xx status code
func (o *WorkflowExecuteInternalTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow execute internal too many requests response a status code equal to that given
func (o *WorkflowExecuteInternalTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the workflow execute internal too many requests response
func (o *WorkflowExecuteInternalTooManyRequests) Code() int {
	return 429
}

func (o *WorkflowExecuteInternalTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /workflows/entities/execute/internal/v1][%d] workflowExecuteInternalTooManyRequests  %+v", 429, o.Payload)
}

func (o *WorkflowExecuteInternalTooManyRequests) String() string {
	return fmt.Sprintf("[POST /workflows/entities/execute/internal/v1][%d] workflowExecuteInternalTooManyRequests  %+v", 429, o.Payload)
}

func (o *WorkflowExecuteInternalTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *WorkflowExecuteInternalTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWorkflowExecuteInternalInternalServerError creates a WorkflowExecuteInternalInternalServerError with default headers values
func NewWorkflowExecuteInternalInternalServerError() *WorkflowExecuteInternalInternalServerError {
	return &WorkflowExecuteInternalInternalServerError{}
}

/*
WorkflowExecuteInternalInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type WorkflowExecuteInternalInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIResourceIDsResponse
}

// IsSuccess returns true when this workflow execute internal internal server error response has a 2xx status code
func (o *WorkflowExecuteInternalInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this workflow execute internal internal server error response has a 3xx status code
func (o *WorkflowExecuteInternalInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow execute internal internal server error response has a 4xx status code
func (o *WorkflowExecuteInternalInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow execute internal internal server error response has a 5xx status code
func (o *WorkflowExecuteInternalInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this workflow execute internal internal server error response a status code equal to that given
func (o *WorkflowExecuteInternalInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the workflow execute internal internal server error response
func (o *WorkflowExecuteInternalInternalServerError) Code() int {
	return 500
}

func (o *WorkflowExecuteInternalInternalServerError) Error() string {
	return fmt.Sprintf("[POST /workflows/entities/execute/internal/v1][%d] workflowExecuteInternalInternalServerError  %+v", 500, o.Payload)
}

func (o *WorkflowExecuteInternalInternalServerError) String() string {
	return fmt.Sprintf("[POST /workflows/entities/execute/internal/v1][%d] workflowExecuteInternalInternalServerError  %+v", 500, o.Payload)
}

func (o *WorkflowExecuteInternalInternalServerError) GetPayload() *models.APIResourceIDsResponse {
	return o.Payload
}

func (o *WorkflowExecuteInternalInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.APIResourceIDsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
