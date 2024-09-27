// Code generated by go-swagger; DO NOT EDIT.

package sample_uploads

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

// ArchiveGetV1Reader is a Reader for the ArchiveGetV1 structure.
type ArchiveGetV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArchiveGetV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewArchiveGetV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewArchiveGetV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewArchiveGetV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewArchiveGetV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewArchiveGetV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /archives/entities/archives/v1] ArchiveGetV1", response, response.Code())
	}
}

// NewArchiveGetV1OK creates a ArchiveGetV1OK with default headers values
func NewArchiveGetV1OK() *ArchiveGetV1OK {
	return &ArchiveGetV1OK{}
}

/*
ArchiveGetV1OK describes a response with status code 200, with default header values.

OK
*/
type ArchiveGetV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ClientArchiveCreateResponseV1
}

// IsSuccess returns true when this archive get v1 o k response has a 2xx status code
func (o *ArchiveGetV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this archive get v1 o k response has a 3xx status code
func (o *ArchiveGetV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive get v1 o k response has a 4xx status code
func (o *ArchiveGetV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this archive get v1 o k response has a 5xx status code
func (o *ArchiveGetV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this archive get v1 o k response a status code equal to that given
func (o *ArchiveGetV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the archive get v1 o k response
func (o *ArchiveGetV1OK) Code() int {
	return 200
}

func (o *ArchiveGetV1OK) Error() string {
	return fmt.Sprintf("[GET /archives/entities/archives/v1][%d] archiveGetV1OK  %+v", 200, o.Payload)
}

func (o *ArchiveGetV1OK) String() string {
	return fmt.Sprintf("[GET /archives/entities/archives/v1][%d] archiveGetV1OK  %+v", 200, o.Payload)
}

func (o *ArchiveGetV1OK) GetPayload() *models.ClientArchiveCreateResponseV1 {
	return o.Payload
}

func (o *ArchiveGetV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ClientArchiveCreateResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArchiveGetV1BadRequest creates a ArchiveGetV1BadRequest with default headers values
func NewArchiveGetV1BadRequest() *ArchiveGetV1BadRequest {
	return &ArchiveGetV1BadRequest{}
}

/*
ArchiveGetV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ArchiveGetV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ClientArchiveCreateResponseV1
}

// IsSuccess returns true when this archive get v1 bad request response has a 2xx status code
func (o *ArchiveGetV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this archive get v1 bad request response has a 3xx status code
func (o *ArchiveGetV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive get v1 bad request response has a 4xx status code
func (o *ArchiveGetV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this archive get v1 bad request response has a 5xx status code
func (o *ArchiveGetV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this archive get v1 bad request response a status code equal to that given
func (o *ArchiveGetV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the archive get v1 bad request response
func (o *ArchiveGetV1BadRequest) Code() int {
	return 400
}

func (o *ArchiveGetV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /archives/entities/archives/v1][%d] archiveGetV1BadRequest  %+v", 400, o.Payload)
}

func (o *ArchiveGetV1BadRequest) String() string {
	return fmt.Sprintf("[GET /archives/entities/archives/v1][%d] archiveGetV1BadRequest  %+v", 400, o.Payload)
}

func (o *ArchiveGetV1BadRequest) GetPayload() *models.ClientArchiveCreateResponseV1 {
	return o.Payload
}

func (o *ArchiveGetV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ClientArchiveCreateResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArchiveGetV1Forbidden creates a ArchiveGetV1Forbidden with default headers values
func NewArchiveGetV1Forbidden() *ArchiveGetV1Forbidden {
	return &ArchiveGetV1Forbidden{}
}

/*
ArchiveGetV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ArchiveGetV1Forbidden struct {

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

// IsSuccess returns true when this archive get v1 forbidden response has a 2xx status code
func (o *ArchiveGetV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this archive get v1 forbidden response has a 3xx status code
func (o *ArchiveGetV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive get v1 forbidden response has a 4xx status code
func (o *ArchiveGetV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this archive get v1 forbidden response has a 5xx status code
func (o *ArchiveGetV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this archive get v1 forbidden response a status code equal to that given
func (o *ArchiveGetV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the archive get v1 forbidden response
func (o *ArchiveGetV1Forbidden) Code() int {
	return 403
}

func (o *ArchiveGetV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /archives/entities/archives/v1][%d] archiveGetV1Forbidden  %+v", 403, o.Payload)
}

func (o *ArchiveGetV1Forbidden) String() string {
	return fmt.Sprintf("[GET /archives/entities/archives/v1][%d] archiveGetV1Forbidden  %+v", 403, o.Payload)
}

func (o *ArchiveGetV1Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ArchiveGetV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewArchiveGetV1TooManyRequests creates a ArchiveGetV1TooManyRequests with default headers values
func NewArchiveGetV1TooManyRequests() *ArchiveGetV1TooManyRequests {
	return &ArchiveGetV1TooManyRequests{}
}

/*
ArchiveGetV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ArchiveGetV1TooManyRequests struct {

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

// IsSuccess returns true when this archive get v1 too many requests response has a 2xx status code
func (o *ArchiveGetV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this archive get v1 too many requests response has a 3xx status code
func (o *ArchiveGetV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive get v1 too many requests response has a 4xx status code
func (o *ArchiveGetV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this archive get v1 too many requests response has a 5xx status code
func (o *ArchiveGetV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this archive get v1 too many requests response a status code equal to that given
func (o *ArchiveGetV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the archive get v1 too many requests response
func (o *ArchiveGetV1TooManyRequests) Code() int {
	return 429
}

func (o *ArchiveGetV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /archives/entities/archives/v1][%d] archiveGetV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *ArchiveGetV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /archives/entities/archives/v1][%d] archiveGetV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *ArchiveGetV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ArchiveGetV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewArchiveGetV1InternalServerError creates a ArchiveGetV1InternalServerError with default headers values
func NewArchiveGetV1InternalServerError() *ArchiveGetV1InternalServerError {
	return &ArchiveGetV1InternalServerError{}
}

/*
ArchiveGetV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ArchiveGetV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ClientArchiveCreateResponseV1
}

// IsSuccess returns true when this archive get v1 internal server error response has a 2xx status code
func (o *ArchiveGetV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this archive get v1 internal server error response has a 3xx status code
func (o *ArchiveGetV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive get v1 internal server error response has a 4xx status code
func (o *ArchiveGetV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this archive get v1 internal server error response has a 5xx status code
func (o *ArchiveGetV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this archive get v1 internal server error response a status code equal to that given
func (o *ArchiveGetV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the archive get v1 internal server error response
func (o *ArchiveGetV1InternalServerError) Code() int {
	return 500
}

func (o *ArchiveGetV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /archives/entities/archives/v1][%d] archiveGetV1InternalServerError  %+v", 500, o.Payload)
}

func (o *ArchiveGetV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /archives/entities/archives/v1][%d] archiveGetV1InternalServerError  %+v", 500, o.Payload)
}

func (o *ArchiveGetV1InternalServerError) GetPayload() *models.ClientArchiveCreateResponseV1 {
	return o.Payload
}

func (o *ArchiveGetV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ClientArchiveCreateResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
