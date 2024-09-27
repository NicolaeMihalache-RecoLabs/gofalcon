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

// QueryChangesReader is a Reader for the QueryChanges structure.
type QueryChangesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryChangesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryChangesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryChangesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryChangesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryChangesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryChangesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /filevantage/queries/changes/v2] queryChanges", response, response.Code())
	}
}

// NewQueryChangesOK creates a QueryChangesOK with default headers values
func NewQueryChangesOK() *QueryChangesOK {
	return &QueryChangesOK{}
}

/*
QueryChangesOK describes a response with status code 200, with default header values.

OK
*/
type QueryChangesOK struct {

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

// IsSuccess returns true when this query changes o k response has a 2xx status code
func (o *QueryChangesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query changes o k response has a 3xx status code
func (o *QueryChangesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query changes o k response has a 4xx status code
func (o *QueryChangesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query changes o k response has a 5xx status code
func (o *QueryChangesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query changes o k response a status code equal to that given
func (o *QueryChangesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query changes o k response
func (o *QueryChangesOK) Code() int {
	return 200
}

func (o *QueryChangesOK) Error() string {
	return fmt.Sprintf("[GET /filevantage/queries/changes/v2][%d] queryChangesOK  %+v", 200, o.Payload)
}

func (o *QueryChangesOK) String() string {
	return fmt.Sprintf("[GET /filevantage/queries/changes/v2][%d] queryChangesOK  %+v", 200, o.Payload)
}

func (o *QueryChangesOK) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *QueryChangesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryChangesBadRequest creates a QueryChangesBadRequest with default headers values
func NewQueryChangesBadRequest() *QueryChangesBadRequest {
	return &QueryChangesBadRequest{}
}

/*
QueryChangesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryChangesBadRequest struct {

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

// IsSuccess returns true when this query changes bad request response has a 2xx status code
func (o *QueryChangesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query changes bad request response has a 3xx status code
func (o *QueryChangesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query changes bad request response has a 4xx status code
func (o *QueryChangesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query changes bad request response has a 5xx status code
func (o *QueryChangesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query changes bad request response a status code equal to that given
func (o *QueryChangesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query changes bad request response
func (o *QueryChangesBadRequest) Code() int {
	return 400
}

func (o *QueryChangesBadRequest) Error() string {
	return fmt.Sprintf("[GET /filevantage/queries/changes/v2][%d] queryChangesBadRequest  %+v", 400, o.Payload)
}

func (o *QueryChangesBadRequest) String() string {
	return fmt.Sprintf("[GET /filevantage/queries/changes/v2][%d] queryChangesBadRequest  %+v", 400, o.Payload)
}

func (o *QueryChangesBadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *QueryChangesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryChangesForbidden creates a QueryChangesForbidden with default headers values
func NewQueryChangesForbidden() *QueryChangesForbidden {
	return &QueryChangesForbidden{}
}

/*
QueryChangesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryChangesForbidden struct {

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

// IsSuccess returns true when this query changes forbidden response has a 2xx status code
func (o *QueryChangesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query changes forbidden response has a 3xx status code
func (o *QueryChangesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query changes forbidden response has a 4xx status code
func (o *QueryChangesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query changes forbidden response has a 5xx status code
func (o *QueryChangesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query changes forbidden response a status code equal to that given
func (o *QueryChangesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query changes forbidden response
func (o *QueryChangesForbidden) Code() int {
	return 403
}

func (o *QueryChangesForbidden) Error() string {
	return fmt.Sprintf("[GET /filevantage/queries/changes/v2][%d] queryChangesForbidden  %+v", 403, o.Payload)
}

func (o *QueryChangesForbidden) String() string {
	return fmt.Sprintf("[GET /filevantage/queries/changes/v2][%d] queryChangesForbidden  %+v", 403, o.Payload)
}

func (o *QueryChangesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryChangesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryChangesTooManyRequests creates a QueryChangesTooManyRequests with default headers values
func NewQueryChangesTooManyRequests() *QueryChangesTooManyRequests {
	return &QueryChangesTooManyRequests{}
}

/*
QueryChangesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryChangesTooManyRequests struct {

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

// IsSuccess returns true when this query changes too many requests response has a 2xx status code
func (o *QueryChangesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query changes too many requests response has a 3xx status code
func (o *QueryChangesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query changes too many requests response has a 4xx status code
func (o *QueryChangesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query changes too many requests response has a 5xx status code
func (o *QueryChangesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query changes too many requests response a status code equal to that given
func (o *QueryChangesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query changes too many requests response
func (o *QueryChangesTooManyRequests) Code() int {
	return 429
}

func (o *QueryChangesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /filevantage/queries/changes/v2][%d] queryChangesTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryChangesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /filevantage/queries/changes/v2][%d] queryChangesTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryChangesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryChangesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryChangesInternalServerError creates a QueryChangesInternalServerError with default headers values
func NewQueryChangesInternalServerError() *QueryChangesInternalServerError {
	return &QueryChangesInternalServerError{}
}

/*
QueryChangesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryChangesInternalServerError struct {

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

// IsSuccess returns true when this query changes internal server error response has a 2xx status code
func (o *QueryChangesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query changes internal server error response has a 3xx status code
func (o *QueryChangesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query changes internal server error response has a 4xx status code
func (o *QueryChangesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query changes internal server error response has a 5xx status code
func (o *QueryChangesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query changes internal server error response a status code equal to that given
func (o *QueryChangesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query changes internal server error response
func (o *QueryChangesInternalServerError) Code() int {
	return 500
}

func (o *QueryChangesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /filevantage/queries/changes/v2][%d] queryChangesInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryChangesInternalServerError) String() string {
	return fmt.Sprintf("[GET /filevantage/queries/changes/v2][%d] queryChangesInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryChangesInternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *QueryChangesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
