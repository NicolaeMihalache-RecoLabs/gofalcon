// Code generated by go-swagger; DO NOT EDIT.

package real_time_response

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

// RTRInitSessionReader is a Reader for the RTRInitSession structure.
type RTRInitSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRInitSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRTRInitSessionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRTRInitSessionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRInitSessionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRInitSessionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRTRInitSessionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /real-time-response/entities/sessions/v1] RTR-InitSession", response, response.Code())
	}
}

// NewRTRInitSessionCreated creates a RTRInitSessionCreated with default headers values
func NewRTRInitSessionCreated() *RTRInitSessionCreated {
	return &RTRInitSessionCreated{}
}

/*
RTRInitSessionCreated describes a response with status code 201, with default header values.

Created
*/
type RTRInitSessionCreated struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainInitResponseWrapper
}

// IsSuccess returns true when this r t r init session created response has a 2xx status code
func (o *RTRInitSessionCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this r t r init session created response has a 3xx status code
func (o *RTRInitSessionCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r init session created response has a 4xx status code
func (o *RTRInitSessionCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r init session created response has a 5xx status code
func (o *RTRInitSessionCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r init session created response a status code equal to that given
func (o *RTRInitSessionCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the r t r init session created response
func (o *RTRInitSessionCreated) Code() int {
	return 201
}

func (o *RTRInitSessionCreated) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/sessions/v1][%d] rTRInitSessionCreated  %+v", 201, o.Payload)
}

func (o *RTRInitSessionCreated) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/sessions/v1][%d] rTRInitSessionCreated  %+v", 201, o.Payload)
}

func (o *RTRInitSessionCreated) GetPayload() *models.DomainInitResponseWrapper {
	return o.Payload
}

func (o *RTRInitSessionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainInitResponseWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRInitSessionBadRequest creates a RTRInitSessionBadRequest with default headers values
func NewRTRInitSessionBadRequest() *RTRInitSessionBadRequest {
	return &RTRInitSessionBadRequest{}
}

/*
RTRInitSessionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RTRInitSessionBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this r t r init session bad request response has a 2xx status code
func (o *RTRInitSessionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r init session bad request response has a 3xx status code
func (o *RTRInitSessionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r init session bad request response has a 4xx status code
func (o *RTRInitSessionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r init session bad request response has a 5xx status code
func (o *RTRInitSessionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r init session bad request response a status code equal to that given
func (o *RTRInitSessionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the r t r init session bad request response
func (o *RTRInitSessionBadRequest) Code() int {
	return 400
}

func (o *RTRInitSessionBadRequest) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/sessions/v1][%d] rTRInitSessionBadRequest  %+v", 400, o.Payload)
}

func (o *RTRInitSessionBadRequest) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/sessions/v1][%d] rTRInitSessionBadRequest  %+v", 400, o.Payload)
}

func (o *RTRInitSessionBadRequest) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRInitSessionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRInitSessionForbidden creates a RTRInitSessionForbidden with default headers values
func NewRTRInitSessionForbidden() *RTRInitSessionForbidden {
	return &RTRInitSessionForbidden{}
}

/*
RTRInitSessionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RTRInitSessionForbidden struct {

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

// IsSuccess returns true when this r t r init session forbidden response has a 2xx status code
func (o *RTRInitSessionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r init session forbidden response has a 3xx status code
func (o *RTRInitSessionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r init session forbidden response has a 4xx status code
func (o *RTRInitSessionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r init session forbidden response has a 5xx status code
func (o *RTRInitSessionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r init session forbidden response a status code equal to that given
func (o *RTRInitSessionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the r t r init session forbidden response
func (o *RTRInitSessionForbidden) Code() int {
	return 403
}

func (o *RTRInitSessionForbidden) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/sessions/v1][%d] rTRInitSessionForbidden  %+v", 403, o.Payload)
}

func (o *RTRInitSessionForbidden) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/sessions/v1][%d] rTRInitSessionForbidden  %+v", 403, o.Payload)
}

func (o *RTRInitSessionForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRInitSessionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRInitSessionTooManyRequests creates a RTRInitSessionTooManyRequests with default headers values
func NewRTRInitSessionTooManyRequests() *RTRInitSessionTooManyRequests {
	return &RTRInitSessionTooManyRequests{}
}

/*
RTRInitSessionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RTRInitSessionTooManyRequests struct {

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

// IsSuccess returns true when this r t r init session too many requests response has a 2xx status code
func (o *RTRInitSessionTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r init session too many requests response has a 3xx status code
func (o *RTRInitSessionTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r init session too many requests response has a 4xx status code
func (o *RTRInitSessionTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r init session too many requests response has a 5xx status code
func (o *RTRInitSessionTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r init session too many requests response a status code equal to that given
func (o *RTRInitSessionTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the r t r init session too many requests response
func (o *RTRInitSessionTooManyRequests) Code() int {
	return 429
}

func (o *RTRInitSessionTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/sessions/v1][%d] rTRInitSessionTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRInitSessionTooManyRequests) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/sessions/v1][%d] rTRInitSessionTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRInitSessionTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRInitSessionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRInitSessionInternalServerError creates a RTRInitSessionInternalServerError with default headers values
func NewRTRInitSessionInternalServerError() *RTRInitSessionInternalServerError {
	return &RTRInitSessionInternalServerError{}
}

/*
RTRInitSessionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type RTRInitSessionInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this r t r init session internal server error response has a 2xx status code
func (o *RTRInitSessionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r init session internal server error response has a 3xx status code
func (o *RTRInitSessionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r init session internal server error response has a 4xx status code
func (o *RTRInitSessionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r init session internal server error response has a 5xx status code
func (o *RTRInitSessionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this r t r init session internal server error response a status code equal to that given
func (o *RTRInitSessionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the r t r init session internal server error response
func (o *RTRInitSessionInternalServerError) Code() int {
	return 500
}

func (o *RTRInitSessionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/sessions/v1][%d] rTRInitSessionInternalServerError  %+v", 500, o.Payload)
}

func (o *RTRInitSessionInternalServerError) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/sessions/v1][%d] rTRInitSessionInternalServerError  %+v", 500, o.Payload)
}

func (o *RTRInitSessionInternalServerError) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRInitSessionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
