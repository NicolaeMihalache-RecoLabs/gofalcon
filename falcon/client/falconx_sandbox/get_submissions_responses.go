// Code generated by go-swagger; DO NOT EDIT.

package falconx_sandbox

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

// GetSubmissionsReader is a Reader for the GetSubmissions structure.
type GetSubmissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubmissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSubmissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSubmissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSubmissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSubmissionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSubmissionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /falconx/entities/submissions/v1] GetSubmissions", response, response.Code())
	}
}

// NewGetSubmissionsOK creates a GetSubmissionsOK with default headers values
func NewGetSubmissionsOK() *GetSubmissionsOK {
	return &GetSubmissionsOK{}
}

/*
GetSubmissionsOK describes a response with status code 200, with default header values.

OK
*/
type GetSubmissionsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FalconxSubmissionV1Response
}

// IsSuccess returns true when this get submissions o k response has a 2xx status code
func (o *GetSubmissionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get submissions o k response has a 3xx status code
func (o *GetSubmissionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get submissions o k response has a 4xx status code
func (o *GetSubmissionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get submissions o k response has a 5xx status code
func (o *GetSubmissionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get submissions o k response a status code equal to that given
func (o *GetSubmissionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get submissions o k response
func (o *GetSubmissionsOK) Code() int {
	return 200
}

func (o *GetSubmissionsOK) Error() string {
	return fmt.Sprintf("[GET /falconx/entities/submissions/v1][%d] getSubmissionsOK  %+v", 200, o.Payload)
}

func (o *GetSubmissionsOK) String() string {
	return fmt.Sprintf("[GET /falconx/entities/submissions/v1][%d] getSubmissionsOK  %+v", 200, o.Payload)
}

func (o *GetSubmissionsOK) GetPayload() *models.FalconxSubmissionV1Response {
	return o.Payload
}

func (o *GetSubmissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FalconxSubmissionV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubmissionsBadRequest creates a GetSubmissionsBadRequest with default headers values
func NewGetSubmissionsBadRequest() *GetSubmissionsBadRequest {
	return &GetSubmissionsBadRequest{}
}

/*
GetSubmissionsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetSubmissionsBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FalconxSubmissionV1Response
}

// IsSuccess returns true when this get submissions bad request response has a 2xx status code
func (o *GetSubmissionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get submissions bad request response has a 3xx status code
func (o *GetSubmissionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get submissions bad request response has a 4xx status code
func (o *GetSubmissionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get submissions bad request response has a 5xx status code
func (o *GetSubmissionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get submissions bad request response a status code equal to that given
func (o *GetSubmissionsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get submissions bad request response
func (o *GetSubmissionsBadRequest) Code() int {
	return 400
}

func (o *GetSubmissionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /falconx/entities/submissions/v1][%d] getSubmissionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetSubmissionsBadRequest) String() string {
	return fmt.Sprintf("[GET /falconx/entities/submissions/v1][%d] getSubmissionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetSubmissionsBadRequest) GetPayload() *models.FalconxSubmissionV1Response {
	return o.Payload
}

func (o *GetSubmissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FalconxSubmissionV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubmissionsForbidden creates a GetSubmissionsForbidden with default headers values
func NewGetSubmissionsForbidden() *GetSubmissionsForbidden {
	return &GetSubmissionsForbidden{}
}

/*
GetSubmissionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSubmissionsForbidden struct {

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

// IsSuccess returns true when this get submissions forbidden response has a 2xx status code
func (o *GetSubmissionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get submissions forbidden response has a 3xx status code
func (o *GetSubmissionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get submissions forbidden response has a 4xx status code
func (o *GetSubmissionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get submissions forbidden response has a 5xx status code
func (o *GetSubmissionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get submissions forbidden response a status code equal to that given
func (o *GetSubmissionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get submissions forbidden response
func (o *GetSubmissionsForbidden) Code() int {
	return 403
}

func (o *GetSubmissionsForbidden) Error() string {
	return fmt.Sprintf("[GET /falconx/entities/submissions/v1][%d] getSubmissionsForbidden  %+v", 403, o.Payload)
}

func (o *GetSubmissionsForbidden) String() string {
	return fmt.Sprintf("[GET /falconx/entities/submissions/v1][%d] getSubmissionsForbidden  %+v", 403, o.Payload)
}

func (o *GetSubmissionsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetSubmissionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSubmissionsTooManyRequests creates a GetSubmissionsTooManyRequests with default headers values
func NewGetSubmissionsTooManyRequests() *GetSubmissionsTooManyRequests {
	return &GetSubmissionsTooManyRequests{}
}

/*
GetSubmissionsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetSubmissionsTooManyRequests struct {

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

// IsSuccess returns true when this get submissions too many requests response has a 2xx status code
func (o *GetSubmissionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get submissions too many requests response has a 3xx status code
func (o *GetSubmissionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get submissions too many requests response has a 4xx status code
func (o *GetSubmissionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get submissions too many requests response has a 5xx status code
func (o *GetSubmissionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get submissions too many requests response a status code equal to that given
func (o *GetSubmissionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get submissions too many requests response
func (o *GetSubmissionsTooManyRequests) Code() int {
	return 429
}

func (o *GetSubmissionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /falconx/entities/submissions/v1][%d] getSubmissionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSubmissionsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /falconx/entities/submissions/v1][%d] getSubmissionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSubmissionsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetSubmissionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSubmissionsInternalServerError creates a GetSubmissionsInternalServerError with default headers values
func NewGetSubmissionsInternalServerError() *GetSubmissionsInternalServerError {
	return &GetSubmissionsInternalServerError{}
}

/*
GetSubmissionsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetSubmissionsInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FalconxSubmissionV1Response
}

// IsSuccess returns true when this get submissions internal server error response has a 2xx status code
func (o *GetSubmissionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get submissions internal server error response has a 3xx status code
func (o *GetSubmissionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get submissions internal server error response has a 4xx status code
func (o *GetSubmissionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get submissions internal server error response has a 5xx status code
func (o *GetSubmissionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get submissions internal server error response a status code equal to that given
func (o *GetSubmissionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get submissions internal server error response
func (o *GetSubmissionsInternalServerError) Code() int {
	return 500
}

func (o *GetSubmissionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /falconx/entities/submissions/v1][%d] getSubmissionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSubmissionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /falconx/entities/submissions/v1][%d] getSubmissionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSubmissionsInternalServerError) GetPayload() *models.FalconxSubmissionV1Response {
	return o.Payload
}

func (o *GetSubmissionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FalconxSubmissionV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
