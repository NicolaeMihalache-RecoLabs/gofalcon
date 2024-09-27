// Code generated by go-swagger; DO NOT EDIT.

package filevantage

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

// SignalChangesExternalReader is a Reader for the SignalChangesExternal structure.
type SignalChangesExternalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SignalChangesExternalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSignalChangesExternalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSignalChangesExternalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSignalChangesExternalForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewSignalChangesExternalConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 424:
		result := NewSignalChangesExternalFailedDependency()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewSignalChangesExternalTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSignalChangesExternalInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /filevantage/entities/workflow/v1] signalChangesExternal", response, response.Code())
	}
}

// NewSignalChangesExternalOK creates a SignalChangesExternalOK with default headers values
func NewSignalChangesExternalOK() *SignalChangesExternalOK {
	return &SignalChangesExternalOK{}
}

/*
SignalChangesExternalOK describes a response with status code 200, with default header values.

Workflow initiated
*/
type SignalChangesExternalOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.WorkflowResponse
}

// IsSuccess returns true when this signal changes external o k response has a 2xx status code
func (o *SignalChangesExternalOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this signal changes external o k response has a 3xx status code
func (o *SignalChangesExternalOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this signal changes external o k response has a 4xx status code
func (o *SignalChangesExternalOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this signal changes external o k response has a 5xx status code
func (o *SignalChangesExternalOK) IsServerError() bool {
	return false
}

// IsCode returns true when this signal changes external o k response a status code equal to that given
func (o *SignalChangesExternalOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the signal changes external o k response
func (o *SignalChangesExternalOK) Code() int {
	return 200
}

func (o *SignalChangesExternalOK) Error() string {
	return fmt.Sprintf("[POST /filevantage/entities/workflow/v1][%d] signalChangesExternalOK  %+v", 200, o.Payload)
}

func (o *SignalChangesExternalOK) String() string {
	return fmt.Sprintf("[POST /filevantage/entities/workflow/v1][%d] signalChangesExternalOK  %+v", 200, o.Payload)
}

func (o *SignalChangesExternalOK) GetPayload() *models.WorkflowResponse {
	return o.Payload
}

func (o *SignalChangesExternalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.WorkflowResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSignalChangesExternalBadRequest creates a SignalChangesExternalBadRequest with default headers values
func NewSignalChangesExternalBadRequest() *SignalChangesExternalBadRequest {
	return &SignalChangesExternalBadRequest{}
}

/*
SignalChangesExternalBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SignalChangesExternalBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this signal changes external bad request response has a 2xx status code
func (o *SignalChangesExternalBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this signal changes external bad request response has a 3xx status code
func (o *SignalChangesExternalBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this signal changes external bad request response has a 4xx status code
func (o *SignalChangesExternalBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this signal changes external bad request response has a 5xx status code
func (o *SignalChangesExternalBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this signal changes external bad request response a status code equal to that given
func (o *SignalChangesExternalBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the signal changes external bad request response
func (o *SignalChangesExternalBadRequest) Code() int {
	return 400
}

func (o *SignalChangesExternalBadRequest) Error() string {
	return fmt.Sprintf("[POST /filevantage/entities/workflow/v1][%d] signalChangesExternalBadRequest  %+v", 400, o.Payload)
}

func (o *SignalChangesExternalBadRequest) String() string {
	return fmt.Sprintf("[POST /filevantage/entities/workflow/v1][%d] signalChangesExternalBadRequest  %+v", 400, o.Payload)
}

func (o *SignalChangesExternalBadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *SignalChangesExternalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSignalChangesExternalForbidden creates a SignalChangesExternalForbidden with default headers values
func NewSignalChangesExternalForbidden() *SignalChangesExternalForbidden {
	return &SignalChangesExternalForbidden{}
}

/*
SignalChangesExternalForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SignalChangesExternalForbidden struct {

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

// IsSuccess returns true when this signal changes external forbidden response has a 2xx status code
func (o *SignalChangesExternalForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this signal changes external forbidden response has a 3xx status code
func (o *SignalChangesExternalForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this signal changes external forbidden response has a 4xx status code
func (o *SignalChangesExternalForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this signal changes external forbidden response has a 5xx status code
func (o *SignalChangesExternalForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this signal changes external forbidden response a status code equal to that given
func (o *SignalChangesExternalForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the signal changes external forbidden response
func (o *SignalChangesExternalForbidden) Code() int {
	return 403
}

func (o *SignalChangesExternalForbidden) Error() string {
	return fmt.Sprintf("[POST /filevantage/entities/workflow/v1][%d] signalChangesExternalForbidden  %+v", 403, o.Payload)
}

func (o *SignalChangesExternalForbidden) String() string {
	return fmt.Sprintf("[POST /filevantage/entities/workflow/v1][%d] signalChangesExternalForbidden  %+v", 403, o.Payload)
}

func (o *SignalChangesExternalForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *SignalChangesExternalForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSignalChangesExternalConflict creates a SignalChangesExternalConflict with default headers values
func NewSignalChangesExternalConflict() *SignalChangesExternalConflict {
	return &SignalChangesExternalConflict{}
}

/*
SignalChangesExternalConflict describes a response with status code 409, with default header values.

Conflict
*/
type SignalChangesExternalConflict struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this signal changes external conflict response has a 2xx status code
func (o *SignalChangesExternalConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this signal changes external conflict response has a 3xx status code
func (o *SignalChangesExternalConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this signal changes external conflict response has a 4xx status code
func (o *SignalChangesExternalConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this signal changes external conflict response has a 5xx status code
func (o *SignalChangesExternalConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this signal changes external conflict response a status code equal to that given
func (o *SignalChangesExternalConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the signal changes external conflict response
func (o *SignalChangesExternalConflict) Code() int {
	return 409
}

func (o *SignalChangesExternalConflict) Error() string {
	return fmt.Sprintf("[POST /filevantage/entities/workflow/v1][%d] signalChangesExternalConflict  %+v", 409, o.Payload)
}

func (o *SignalChangesExternalConflict) String() string {
	return fmt.Sprintf("[POST /filevantage/entities/workflow/v1][%d] signalChangesExternalConflict  %+v", 409, o.Payload)
}

func (o *SignalChangesExternalConflict) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *SignalChangesExternalConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSignalChangesExternalFailedDependency creates a SignalChangesExternalFailedDependency with default headers values
func NewSignalChangesExternalFailedDependency() *SignalChangesExternalFailedDependency {
	return &SignalChangesExternalFailedDependency{}
}

/*
SignalChangesExternalFailedDependency describes a response with status code 424, with default header values.

Failed Dependency
*/
type SignalChangesExternalFailedDependency struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this signal changes external failed dependency response has a 2xx status code
func (o *SignalChangesExternalFailedDependency) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this signal changes external failed dependency response has a 3xx status code
func (o *SignalChangesExternalFailedDependency) IsRedirect() bool {
	return false
}

// IsClientError returns true when this signal changes external failed dependency response has a 4xx status code
func (o *SignalChangesExternalFailedDependency) IsClientError() bool {
	return true
}

// IsServerError returns true when this signal changes external failed dependency response has a 5xx status code
func (o *SignalChangesExternalFailedDependency) IsServerError() bool {
	return false
}

// IsCode returns true when this signal changes external failed dependency response a status code equal to that given
func (o *SignalChangesExternalFailedDependency) IsCode(code int) bool {
	return code == 424
}

// Code gets the status code for the signal changes external failed dependency response
func (o *SignalChangesExternalFailedDependency) Code() int {
	return 424
}

func (o *SignalChangesExternalFailedDependency) Error() string {
	return fmt.Sprintf("[POST /filevantage/entities/workflow/v1][%d] signalChangesExternalFailedDependency  %+v", 424, o.Payload)
}

func (o *SignalChangesExternalFailedDependency) String() string {
	return fmt.Sprintf("[POST /filevantage/entities/workflow/v1][%d] signalChangesExternalFailedDependency  %+v", 424, o.Payload)
}

func (o *SignalChangesExternalFailedDependency) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *SignalChangesExternalFailedDependency) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSignalChangesExternalTooManyRequests creates a SignalChangesExternalTooManyRequests with default headers values
func NewSignalChangesExternalTooManyRequests() *SignalChangesExternalTooManyRequests {
	return &SignalChangesExternalTooManyRequests{}
}

/*
SignalChangesExternalTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type SignalChangesExternalTooManyRequests struct {

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

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this signal changes external too many requests response has a 2xx status code
func (o *SignalChangesExternalTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this signal changes external too many requests response has a 3xx status code
func (o *SignalChangesExternalTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this signal changes external too many requests response has a 4xx status code
func (o *SignalChangesExternalTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this signal changes external too many requests response has a 5xx status code
func (o *SignalChangesExternalTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this signal changes external too many requests response a status code equal to that given
func (o *SignalChangesExternalTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the signal changes external too many requests response
func (o *SignalChangesExternalTooManyRequests) Code() int {
	return 429
}

func (o *SignalChangesExternalTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /filevantage/entities/workflow/v1][%d] signalChangesExternalTooManyRequests  %+v", 429, o.Payload)
}

func (o *SignalChangesExternalTooManyRequests) String() string {
	return fmt.Sprintf("[POST /filevantage/entities/workflow/v1][%d] signalChangesExternalTooManyRequests  %+v", 429, o.Payload)
}

func (o *SignalChangesExternalTooManyRequests) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *SignalChangesExternalTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSignalChangesExternalInternalServerError creates a SignalChangesExternalInternalServerError with default headers values
func NewSignalChangesExternalInternalServerError() *SignalChangesExternalInternalServerError {
	return &SignalChangesExternalInternalServerError{}
}

/*
SignalChangesExternalInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type SignalChangesExternalInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this signal changes external internal server error response has a 2xx status code
func (o *SignalChangesExternalInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this signal changes external internal server error response has a 3xx status code
func (o *SignalChangesExternalInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this signal changes external internal server error response has a 4xx status code
func (o *SignalChangesExternalInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this signal changes external internal server error response has a 5xx status code
func (o *SignalChangesExternalInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this signal changes external internal server error response a status code equal to that given
func (o *SignalChangesExternalInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the signal changes external internal server error response
func (o *SignalChangesExternalInternalServerError) Code() int {
	return 500
}

func (o *SignalChangesExternalInternalServerError) Error() string {
	return fmt.Sprintf("[POST /filevantage/entities/workflow/v1][%d] signalChangesExternalInternalServerError  %+v", 500, o.Payload)
}

func (o *SignalChangesExternalInternalServerError) String() string {
	return fmt.Sprintf("[POST /filevantage/entities/workflow/v1][%d] signalChangesExternalInternalServerError  %+v", 500, o.Payload)
}

func (o *SignalChangesExternalInternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *SignalChangesExternalInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
