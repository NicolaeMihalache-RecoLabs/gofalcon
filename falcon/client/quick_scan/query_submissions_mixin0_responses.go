// Code generated by go-swagger; DO NOT EDIT.

package quick_scan

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

// QuerySubmissionsMixin0Reader is a Reader for the QuerySubmissionsMixin0 structure.
type QuerySubmissionsMixin0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuerySubmissionsMixin0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQuerySubmissionsMixin0OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQuerySubmissionsMixin0BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQuerySubmissionsMixin0Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQuerySubmissionsMixin0TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQuerySubmissionsMixin0InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /scanner/queries/scans/v1] QuerySubmissionsMixin0", response, response.Code())
	}
}

// NewQuerySubmissionsMixin0OK creates a QuerySubmissionsMixin0OK with default headers values
func NewQuerySubmissionsMixin0OK() *QuerySubmissionsMixin0OK {
	return &QuerySubmissionsMixin0OK{}
}

/*
QuerySubmissionsMixin0OK describes a response with status code 200, with default header values.

OK
*/
type QuerySubmissionsMixin0OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MlscannerapiQueryResponse
}

// IsSuccess returns true when this query submissions mixin0 o k response has a 2xx status code
func (o *QuerySubmissionsMixin0OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query submissions mixin0 o k response has a 3xx status code
func (o *QuerySubmissionsMixin0OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query submissions mixin0 o k response has a 4xx status code
func (o *QuerySubmissionsMixin0OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query submissions mixin0 o k response has a 5xx status code
func (o *QuerySubmissionsMixin0OK) IsServerError() bool {
	return false
}

// IsCode returns true when this query submissions mixin0 o k response a status code equal to that given
func (o *QuerySubmissionsMixin0OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query submissions mixin0 o k response
func (o *QuerySubmissionsMixin0OK) Code() int {
	return 200
}

func (o *QuerySubmissionsMixin0OK) Error() string {
	return fmt.Sprintf("[GET /scanner/queries/scans/v1][%d] querySubmissionsMixin0OK  %+v", 200, o.Payload)
}

func (o *QuerySubmissionsMixin0OK) String() string {
	return fmt.Sprintf("[GET /scanner/queries/scans/v1][%d] querySubmissionsMixin0OK  %+v", 200, o.Payload)
}

func (o *QuerySubmissionsMixin0OK) GetPayload() *models.MlscannerapiQueryResponse {
	return o.Payload
}

func (o *QuerySubmissionsMixin0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MlscannerapiQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuerySubmissionsMixin0BadRequest creates a QuerySubmissionsMixin0BadRequest with default headers values
func NewQuerySubmissionsMixin0BadRequest() *QuerySubmissionsMixin0BadRequest {
	return &QuerySubmissionsMixin0BadRequest{}
}

/*
QuerySubmissionsMixin0BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QuerySubmissionsMixin0BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MlscannerapiQueryResponse
}

// IsSuccess returns true when this query submissions mixin0 bad request response has a 2xx status code
func (o *QuerySubmissionsMixin0BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query submissions mixin0 bad request response has a 3xx status code
func (o *QuerySubmissionsMixin0BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query submissions mixin0 bad request response has a 4xx status code
func (o *QuerySubmissionsMixin0BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query submissions mixin0 bad request response has a 5xx status code
func (o *QuerySubmissionsMixin0BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query submissions mixin0 bad request response a status code equal to that given
func (o *QuerySubmissionsMixin0BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query submissions mixin0 bad request response
func (o *QuerySubmissionsMixin0BadRequest) Code() int {
	return 400
}

func (o *QuerySubmissionsMixin0BadRequest) Error() string {
	return fmt.Sprintf("[GET /scanner/queries/scans/v1][%d] querySubmissionsMixin0BadRequest  %+v", 400, o.Payload)
}

func (o *QuerySubmissionsMixin0BadRequest) String() string {
	return fmt.Sprintf("[GET /scanner/queries/scans/v1][%d] querySubmissionsMixin0BadRequest  %+v", 400, o.Payload)
}

func (o *QuerySubmissionsMixin0BadRequest) GetPayload() *models.MlscannerapiQueryResponse {
	return o.Payload
}

func (o *QuerySubmissionsMixin0BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MlscannerapiQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuerySubmissionsMixin0Forbidden creates a QuerySubmissionsMixin0Forbidden with default headers values
func NewQuerySubmissionsMixin0Forbidden() *QuerySubmissionsMixin0Forbidden {
	return &QuerySubmissionsMixin0Forbidden{}
}

/*
QuerySubmissionsMixin0Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QuerySubmissionsMixin0Forbidden struct {

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

// IsSuccess returns true when this query submissions mixin0 forbidden response has a 2xx status code
func (o *QuerySubmissionsMixin0Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query submissions mixin0 forbidden response has a 3xx status code
func (o *QuerySubmissionsMixin0Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query submissions mixin0 forbidden response has a 4xx status code
func (o *QuerySubmissionsMixin0Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query submissions mixin0 forbidden response has a 5xx status code
func (o *QuerySubmissionsMixin0Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query submissions mixin0 forbidden response a status code equal to that given
func (o *QuerySubmissionsMixin0Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query submissions mixin0 forbidden response
func (o *QuerySubmissionsMixin0Forbidden) Code() int {
	return 403
}

func (o *QuerySubmissionsMixin0Forbidden) Error() string {
	return fmt.Sprintf("[GET /scanner/queries/scans/v1][%d] querySubmissionsMixin0Forbidden  %+v", 403, o.Payload)
}

func (o *QuerySubmissionsMixin0Forbidden) String() string {
	return fmt.Sprintf("[GET /scanner/queries/scans/v1][%d] querySubmissionsMixin0Forbidden  %+v", 403, o.Payload)
}

func (o *QuerySubmissionsMixin0Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QuerySubmissionsMixin0Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQuerySubmissionsMixin0TooManyRequests creates a QuerySubmissionsMixin0TooManyRequests with default headers values
func NewQuerySubmissionsMixin0TooManyRequests() *QuerySubmissionsMixin0TooManyRequests {
	return &QuerySubmissionsMixin0TooManyRequests{}
}

/*
QuerySubmissionsMixin0TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QuerySubmissionsMixin0TooManyRequests struct {

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

// IsSuccess returns true when this query submissions mixin0 too many requests response has a 2xx status code
func (o *QuerySubmissionsMixin0TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query submissions mixin0 too many requests response has a 3xx status code
func (o *QuerySubmissionsMixin0TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query submissions mixin0 too many requests response has a 4xx status code
func (o *QuerySubmissionsMixin0TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query submissions mixin0 too many requests response has a 5xx status code
func (o *QuerySubmissionsMixin0TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query submissions mixin0 too many requests response a status code equal to that given
func (o *QuerySubmissionsMixin0TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query submissions mixin0 too many requests response
func (o *QuerySubmissionsMixin0TooManyRequests) Code() int {
	return 429
}

func (o *QuerySubmissionsMixin0TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /scanner/queries/scans/v1][%d] querySubmissionsMixin0TooManyRequests  %+v", 429, o.Payload)
}

func (o *QuerySubmissionsMixin0TooManyRequests) String() string {
	return fmt.Sprintf("[GET /scanner/queries/scans/v1][%d] querySubmissionsMixin0TooManyRequests  %+v", 429, o.Payload)
}

func (o *QuerySubmissionsMixin0TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QuerySubmissionsMixin0TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQuerySubmissionsMixin0InternalServerError creates a QuerySubmissionsMixin0InternalServerError with default headers values
func NewQuerySubmissionsMixin0InternalServerError() *QuerySubmissionsMixin0InternalServerError {
	return &QuerySubmissionsMixin0InternalServerError{}
}

/*
QuerySubmissionsMixin0InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QuerySubmissionsMixin0InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MlscannerapiQueryResponse
}

// IsSuccess returns true when this query submissions mixin0 internal server error response has a 2xx status code
func (o *QuerySubmissionsMixin0InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query submissions mixin0 internal server error response has a 3xx status code
func (o *QuerySubmissionsMixin0InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query submissions mixin0 internal server error response has a 4xx status code
func (o *QuerySubmissionsMixin0InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query submissions mixin0 internal server error response has a 5xx status code
func (o *QuerySubmissionsMixin0InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query submissions mixin0 internal server error response a status code equal to that given
func (o *QuerySubmissionsMixin0InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query submissions mixin0 internal server error response
func (o *QuerySubmissionsMixin0InternalServerError) Code() int {
	return 500
}

func (o *QuerySubmissionsMixin0InternalServerError) Error() string {
	return fmt.Sprintf("[GET /scanner/queries/scans/v1][%d] querySubmissionsMixin0InternalServerError  %+v", 500, o.Payload)
}

func (o *QuerySubmissionsMixin0InternalServerError) String() string {
	return fmt.Sprintf("[GET /scanner/queries/scans/v1][%d] querySubmissionsMixin0InternalServerError  %+v", 500, o.Payload)
}

func (o *QuerySubmissionsMixin0InternalServerError) GetPayload() *models.MlscannerapiQueryResponse {
	return o.Payload
}

func (o *QuerySubmissionsMixin0InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MlscannerapiQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
