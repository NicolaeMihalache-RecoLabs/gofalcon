// Code generated by go-swagger; DO NOT EDIT.

package ioa_exclusions

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

// UpdateIOAExclusionsV1Reader is a Reader for the UpdateIOAExclusionsV1 structure.
type UpdateIOAExclusionsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateIOAExclusionsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateIOAExclusionsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateIOAExclusionsV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateIOAExclusionsV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateIOAExclusionsV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateIOAExclusionsV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /policy/entities/ioa-exclusions/v1] updateIOAExclusionsV1", response, response.Code())
	}
}

// NewUpdateIOAExclusionsV1OK creates a UpdateIOAExclusionsV1OK with default headers values
func NewUpdateIOAExclusionsV1OK() *UpdateIOAExclusionsV1OK {
	return &UpdateIOAExclusionsV1OK{}
}

/*
UpdateIOAExclusionsV1OK describes a response with status code 200, with default header values.

OK
*/
type UpdateIOAExclusionsV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.IoaExclusionsIoaExclusionsRespV1
}

// IsSuccess returns true when this update i o a exclusions v1 o k response has a 2xx status code
func (o *UpdateIOAExclusionsV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update i o a exclusions v1 o k response has a 3xx status code
func (o *UpdateIOAExclusionsV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update i o a exclusions v1 o k response has a 4xx status code
func (o *UpdateIOAExclusionsV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update i o a exclusions v1 o k response has a 5xx status code
func (o *UpdateIOAExclusionsV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this update i o a exclusions v1 o k response a status code equal to that given
func (o *UpdateIOAExclusionsV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update i o a exclusions v1 o k response
func (o *UpdateIOAExclusionsV1OK) Code() int {
	return 200
}

func (o *UpdateIOAExclusionsV1OK) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/ioa-exclusions/v1][%d] updateIOAExclusionsV1OK  %+v", 200, o.Payload)
}

func (o *UpdateIOAExclusionsV1OK) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/ioa-exclusions/v1][%d] updateIOAExclusionsV1OK  %+v", 200, o.Payload)
}

func (o *UpdateIOAExclusionsV1OK) GetPayload() *models.IoaExclusionsIoaExclusionsRespV1 {
	return o.Payload
}

func (o *UpdateIOAExclusionsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.IoaExclusionsIoaExclusionsRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIOAExclusionsV1BadRequest creates a UpdateIOAExclusionsV1BadRequest with default headers values
func NewUpdateIOAExclusionsV1BadRequest() *UpdateIOAExclusionsV1BadRequest {
	return &UpdateIOAExclusionsV1BadRequest{}
}

/*
UpdateIOAExclusionsV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateIOAExclusionsV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.IoaExclusionsIoaExclusionsRespV1
}

// IsSuccess returns true when this update i o a exclusions v1 bad request response has a 2xx status code
func (o *UpdateIOAExclusionsV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update i o a exclusions v1 bad request response has a 3xx status code
func (o *UpdateIOAExclusionsV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update i o a exclusions v1 bad request response has a 4xx status code
func (o *UpdateIOAExclusionsV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update i o a exclusions v1 bad request response has a 5xx status code
func (o *UpdateIOAExclusionsV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update i o a exclusions v1 bad request response a status code equal to that given
func (o *UpdateIOAExclusionsV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update i o a exclusions v1 bad request response
func (o *UpdateIOAExclusionsV1BadRequest) Code() int {
	return 400
}

func (o *UpdateIOAExclusionsV1BadRequest) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/ioa-exclusions/v1][%d] updateIOAExclusionsV1BadRequest  %+v", 400, o.Payload)
}

func (o *UpdateIOAExclusionsV1BadRequest) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/ioa-exclusions/v1][%d] updateIOAExclusionsV1BadRequest  %+v", 400, o.Payload)
}

func (o *UpdateIOAExclusionsV1BadRequest) GetPayload() *models.IoaExclusionsIoaExclusionsRespV1 {
	return o.Payload
}

func (o *UpdateIOAExclusionsV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.IoaExclusionsIoaExclusionsRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIOAExclusionsV1Forbidden creates a UpdateIOAExclusionsV1Forbidden with default headers values
func NewUpdateIOAExclusionsV1Forbidden() *UpdateIOAExclusionsV1Forbidden {
	return &UpdateIOAExclusionsV1Forbidden{}
}

/*
UpdateIOAExclusionsV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateIOAExclusionsV1Forbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this update i o a exclusions v1 forbidden response has a 2xx status code
func (o *UpdateIOAExclusionsV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update i o a exclusions v1 forbidden response has a 3xx status code
func (o *UpdateIOAExclusionsV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update i o a exclusions v1 forbidden response has a 4xx status code
func (o *UpdateIOAExclusionsV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update i o a exclusions v1 forbidden response has a 5xx status code
func (o *UpdateIOAExclusionsV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update i o a exclusions v1 forbidden response a status code equal to that given
func (o *UpdateIOAExclusionsV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update i o a exclusions v1 forbidden response
func (o *UpdateIOAExclusionsV1Forbidden) Code() int {
	return 403
}

func (o *UpdateIOAExclusionsV1Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/ioa-exclusions/v1][%d] updateIOAExclusionsV1Forbidden  %+v", 403, o.Payload)
}

func (o *UpdateIOAExclusionsV1Forbidden) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/ioa-exclusions/v1][%d] updateIOAExclusionsV1Forbidden  %+v", 403, o.Payload)
}

func (o *UpdateIOAExclusionsV1Forbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *UpdateIOAExclusionsV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIOAExclusionsV1TooManyRequests creates a UpdateIOAExclusionsV1TooManyRequests with default headers values
func NewUpdateIOAExclusionsV1TooManyRequests() *UpdateIOAExclusionsV1TooManyRequests {
	return &UpdateIOAExclusionsV1TooManyRequests{}
}

/*
UpdateIOAExclusionsV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateIOAExclusionsV1TooManyRequests struct {

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

// IsSuccess returns true when this update i o a exclusions v1 too many requests response has a 2xx status code
func (o *UpdateIOAExclusionsV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update i o a exclusions v1 too many requests response has a 3xx status code
func (o *UpdateIOAExclusionsV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update i o a exclusions v1 too many requests response has a 4xx status code
func (o *UpdateIOAExclusionsV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update i o a exclusions v1 too many requests response has a 5xx status code
func (o *UpdateIOAExclusionsV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update i o a exclusions v1 too many requests response a status code equal to that given
func (o *UpdateIOAExclusionsV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the update i o a exclusions v1 too many requests response
func (o *UpdateIOAExclusionsV1TooManyRequests) Code() int {
	return 429
}

func (o *UpdateIOAExclusionsV1TooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/ioa-exclusions/v1][%d] updateIOAExclusionsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateIOAExclusionsV1TooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/ioa-exclusions/v1][%d] updateIOAExclusionsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateIOAExclusionsV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateIOAExclusionsV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateIOAExclusionsV1InternalServerError creates a UpdateIOAExclusionsV1InternalServerError with default headers values
func NewUpdateIOAExclusionsV1InternalServerError() *UpdateIOAExclusionsV1InternalServerError {
	return &UpdateIOAExclusionsV1InternalServerError{}
}

/*
UpdateIOAExclusionsV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateIOAExclusionsV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.IoaExclusionsIoaExclusionsRespV1
}

// IsSuccess returns true when this update i o a exclusions v1 internal server error response has a 2xx status code
func (o *UpdateIOAExclusionsV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update i o a exclusions v1 internal server error response has a 3xx status code
func (o *UpdateIOAExclusionsV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update i o a exclusions v1 internal server error response has a 4xx status code
func (o *UpdateIOAExclusionsV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update i o a exclusions v1 internal server error response has a 5xx status code
func (o *UpdateIOAExclusionsV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update i o a exclusions v1 internal server error response a status code equal to that given
func (o *UpdateIOAExclusionsV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update i o a exclusions v1 internal server error response
func (o *UpdateIOAExclusionsV1InternalServerError) Code() int {
	return 500
}

func (o *UpdateIOAExclusionsV1InternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/ioa-exclusions/v1][%d] updateIOAExclusionsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateIOAExclusionsV1InternalServerError) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/ioa-exclusions/v1][%d] updateIOAExclusionsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateIOAExclusionsV1InternalServerError) GetPayload() *models.IoaExclusionsIoaExclusionsRespV1 {
	return o.Payload
}

func (o *UpdateIOAExclusionsV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.IoaExclusionsIoaExclusionsRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
