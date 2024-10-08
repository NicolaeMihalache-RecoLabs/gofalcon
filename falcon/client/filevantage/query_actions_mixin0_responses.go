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

// QueryActionsMixin0Reader is a Reader for the QueryActionsMixin0 structure.
type QueryActionsMixin0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryActionsMixin0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryActionsMixin0OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryActionsMixin0BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryActionsMixin0Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryActionsMixin0TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryActionsMixin0InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /filevantage/queries/actions/v1] queryActionsMixin0", response, response.Code())
	}
}

// NewQueryActionsMixin0OK creates a QueryActionsMixin0OK with default headers values
func NewQueryActionsMixin0OK() *QueryActionsMixin0OK {
	return &QueryActionsMixin0OK{}
}

/*
QueryActionsMixin0OK describes a response with status code 200, with default header values.

OK
*/
type QueryActionsMixin0OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecQueryResponse
}

// IsSuccess returns true when this query actions mixin0 o k response has a 2xx status code
func (o *QueryActionsMixin0OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query actions mixin0 o k response has a 3xx status code
func (o *QueryActionsMixin0OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query actions mixin0 o k response has a 4xx status code
func (o *QueryActionsMixin0OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query actions mixin0 o k response has a 5xx status code
func (o *QueryActionsMixin0OK) IsServerError() bool {
	return false
}

// IsCode returns true when this query actions mixin0 o k response a status code equal to that given
func (o *QueryActionsMixin0OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query actions mixin0 o k response
func (o *QueryActionsMixin0OK) Code() int {
	return 200
}

func (o *QueryActionsMixin0OK) Error() string {
	return fmt.Sprintf("[GET /filevantage/queries/actions/v1][%d] queryActionsMixin0OK  %+v", 200, o.Payload)
}

func (o *QueryActionsMixin0OK) String() string {
	return fmt.Sprintf("[GET /filevantage/queries/actions/v1][%d] queryActionsMixin0OK  %+v", 200, o.Payload)
}

func (o *QueryActionsMixin0OK) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *QueryActionsMixin0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryActionsMixin0BadRequest creates a QueryActionsMixin0BadRequest with default headers values
func NewQueryActionsMixin0BadRequest() *QueryActionsMixin0BadRequest {
	return &QueryActionsMixin0BadRequest{}
}

/*
QueryActionsMixin0BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryActionsMixin0BadRequest struct {

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

// IsSuccess returns true when this query actions mixin0 bad request response has a 2xx status code
func (o *QueryActionsMixin0BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query actions mixin0 bad request response has a 3xx status code
func (o *QueryActionsMixin0BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query actions mixin0 bad request response has a 4xx status code
func (o *QueryActionsMixin0BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query actions mixin0 bad request response has a 5xx status code
func (o *QueryActionsMixin0BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query actions mixin0 bad request response a status code equal to that given
func (o *QueryActionsMixin0BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query actions mixin0 bad request response
func (o *QueryActionsMixin0BadRequest) Code() int {
	return 400
}

func (o *QueryActionsMixin0BadRequest) Error() string {
	return fmt.Sprintf("[GET /filevantage/queries/actions/v1][%d] queryActionsMixin0BadRequest  %+v", 400, o.Payload)
}

func (o *QueryActionsMixin0BadRequest) String() string {
	return fmt.Sprintf("[GET /filevantage/queries/actions/v1][%d] queryActionsMixin0BadRequest  %+v", 400, o.Payload)
}

func (o *QueryActionsMixin0BadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *QueryActionsMixin0BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryActionsMixin0Forbidden creates a QueryActionsMixin0Forbidden with default headers values
func NewQueryActionsMixin0Forbidden() *QueryActionsMixin0Forbidden {
	return &QueryActionsMixin0Forbidden{}
}

/*
QueryActionsMixin0Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryActionsMixin0Forbidden struct {

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

// IsSuccess returns true when this query actions mixin0 forbidden response has a 2xx status code
func (o *QueryActionsMixin0Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query actions mixin0 forbidden response has a 3xx status code
func (o *QueryActionsMixin0Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query actions mixin0 forbidden response has a 4xx status code
func (o *QueryActionsMixin0Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query actions mixin0 forbidden response has a 5xx status code
func (o *QueryActionsMixin0Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query actions mixin0 forbidden response a status code equal to that given
func (o *QueryActionsMixin0Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query actions mixin0 forbidden response
func (o *QueryActionsMixin0Forbidden) Code() int {
	return 403
}

func (o *QueryActionsMixin0Forbidden) Error() string {
	return fmt.Sprintf("[GET /filevantage/queries/actions/v1][%d] queryActionsMixin0Forbidden  %+v", 403, o.Payload)
}

func (o *QueryActionsMixin0Forbidden) String() string {
	return fmt.Sprintf("[GET /filevantage/queries/actions/v1][%d] queryActionsMixin0Forbidden  %+v", 403, o.Payload)
}

func (o *QueryActionsMixin0Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryActionsMixin0Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryActionsMixin0TooManyRequests creates a QueryActionsMixin0TooManyRequests with default headers values
func NewQueryActionsMixin0TooManyRequests() *QueryActionsMixin0TooManyRequests {
	return &QueryActionsMixin0TooManyRequests{}
}

/*
QueryActionsMixin0TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryActionsMixin0TooManyRequests struct {

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

// IsSuccess returns true when this query actions mixin0 too many requests response has a 2xx status code
func (o *QueryActionsMixin0TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query actions mixin0 too many requests response has a 3xx status code
func (o *QueryActionsMixin0TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query actions mixin0 too many requests response has a 4xx status code
func (o *QueryActionsMixin0TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query actions mixin0 too many requests response has a 5xx status code
func (o *QueryActionsMixin0TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query actions mixin0 too many requests response a status code equal to that given
func (o *QueryActionsMixin0TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query actions mixin0 too many requests response
func (o *QueryActionsMixin0TooManyRequests) Code() int {
	return 429
}

func (o *QueryActionsMixin0TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /filevantage/queries/actions/v1][%d] queryActionsMixin0TooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryActionsMixin0TooManyRequests) String() string {
	return fmt.Sprintf("[GET /filevantage/queries/actions/v1][%d] queryActionsMixin0TooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryActionsMixin0TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryActionsMixin0TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryActionsMixin0InternalServerError creates a QueryActionsMixin0InternalServerError with default headers values
func NewQueryActionsMixin0InternalServerError() *QueryActionsMixin0InternalServerError {
	return &QueryActionsMixin0InternalServerError{}
}

/*
QueryActionsMixin0InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryActionsMixin0InternalServerError struct {

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

// IsSuccess returns true when this query actions mixin0 internal server error response has a 2xx status code
func (o *QueryActionsMixin0InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query actions mixin0 internal server error response has a 3xx status code
func (o *QueryActionsMixin0InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query actions mixin0 internal server error response has a 4xx status code
func (o *QueryActionsMixin0InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query actions mixin0 internal server error response has a 5xx status code
func (o *QueryActionsMixin0InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query actions mixin0 internal server error response a status code equal to that given
func (o *QueryActionsMixin0InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query actions mixin0 internal server error response
func (o *QueryActionsMixin0InternalServerError) Code() int {
	return 500
}

func (o *QueryActionsMixin0InternalServerError) Error() string {
	return fmt.Sprintf("[GET /filevantage/queries/actions/v1][%d] queryActionsMixin0InternalServerError  %+v", 500, o.Payload)
}

func (o *QueryActionsMixin0InternalServerError) String() string {
	return fmt.Sprintf("[GET /filevantage/queries/actions/v1][%d] queryActionsMixin0InternalServerError  %+v", 500, o.Payload)
}

func (o *QueryActionsMixin0InternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *QueryActionsMixin0InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
