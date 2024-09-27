// Code generated by go-swagger; DO NOT EDIT.

package unidentified_containers

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

// CountReader is a Reader for the Count structure.
type CountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/aggregates/unidentified-containers/count/v1] Count", response, response.Code())
	}
}

// NewCountOK creates a CountOK with default headers values
func NewCountOK() *CountOK {
	return &CountOK{}
}

/*
CountOK describes a response with status code 200, with default header values.

OK
*/
type CountOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.UnidentifiedcontainersUnidentifiedContainersCountValue
}

// IsSuccess returns true when this count o k response has a 2xx status code
func (o *CountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this count o k response has a 3xx status code
func (o *CountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this count o k response has a 4xx status code
func (o *CountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this count o k response has a 5xx status code
func (o *CountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this count o k response a status code equal to that given
func (o *CountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the count o k response
func (o *CountOK) Code() int {
	return 200
}

func (o *CountOK) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/unidentified-containers/count/v1][%d] countOK  %+v", 200, o.Payload)
}

func (o *CountOK) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/unidentified-containers/count/v1][%d] countOK  %+v", 200, o.Payload)
}

func (o *CountOK) GetPayload() *models.UnidentifiedcontainersUnidentifiedContainersCountValue {
	return o.Payload
}

func (o *CountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.UnidentifiedcontainersUnidentifiedContainersCountValue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCountForbidden creates a CountForbidden with default headers values
func NewCountForbidden() *CountForbidden {
	return &CountForbidden{}
}

/*
CountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CountForbidden struct {

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

// IsSuccess returns true when this count forbidden response has a 2xx status code
func (o *CountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this count forbidden response has a 3xx status code
func (o *CountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this count forbidden response has a 4xx status code
func (o *CountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this count forbidden response has a 5xx status code
func (o *CountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this count forbidden response a status code equal to that given
func (o *CountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the count forbidden response
func (o *CountForbidden) Code() int {
	return 403
}

func (o *CountForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/unidentified-containers/count/v1][%d] countForbidden  %+v", 403, o.Payload)
}

func (o *CountForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/unidentified-containers/count/v1][%d] countForbidden  %+v", 403, o.Payload)
}

func (o *CountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCountTooManyRequests creates a CountTooManyRequests with default headers values
func NewCountTooManyRequests() *CountTooManyRequests {
	return &CountTooManyRequests{}
}

/*
CountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CountTooManyRequests struct {

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

// IsSuccess returns true when this count too many requests response has a 2xx status code
func (o *CountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this count too many requests response has a 3xx status code
func (o *CountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this count too many requests response has a 4xx status code
func (o *CountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this count too many requests response has a 5xx status code
func (o *CountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this count too many requests response a status code equal to that given
func (o *CountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the count too many requests response
func (o *CountTooManyRequests) Code() int {
	return 429
}

func (o *CountTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/unidentified-containers/count/v1][%d] countTooManyRequests  %+v", 429, o.Payload)
}

func (o *CountTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/unidentified-containers/count/v1][%d] countTooManyRequests  %+v", 429, o.Payload)
}

func (o *CountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCountInternalServerError creates a CountInternalServerError with default headers values
func NewCountInternalServerError() *CountInternalServerError {
	return &CountInternalServerError{}
}

/*
CountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CountInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.CoreEntitiesResponse
}

// IsSuccess returns true when this count internal server error response has a 2xx status code
func (o *CountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this count internal server error response has a 3xx status code
func (o *CountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this count internal server error response has a 4xx status code
func (o *CountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this count internal server error response has a 5xx status code
func (o *CountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this count internal server error response a status code equal to that given
func (o *CountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the count internal server error response
func (o *CountInternalServerError) Code() int {
	return 500
}

func (o *CountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/unidentified-containers/count/v1][%d] countInternalServerError  %+v", 500, o.Payload)
}

func (o *CountInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/unidentified-containers/count/v1][%d] countInternalServerError  %+v", 500, o.Payload)
}

func (o *CountInternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *CountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.CoreEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
