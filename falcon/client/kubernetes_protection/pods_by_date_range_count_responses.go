// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_protection

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

// PodsByDateRangeCountReader is a Reader for the PodsByDateRangeCount structure.
type PodsByDateRangeCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PodsByDateRangeCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPodsByDateRangeCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPodsByDateRangeCountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPodsByDateRangeCountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPodsByDateRangeCountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/aggregates/pods/count-by-date/v1] PodsByDateRangeCount", response, response.Code())
	}
}

// NewPodsByDateRangeCountOK creates a PodsByDateRangeCountOK with default headers values
func NewPodsByDateRangeCountOK() *PodsByDateRangeCountOK {
	return &PodsByDateRangeCountOK{}
}

/*
PodsByDateRangeCountOK describes a response with status code 200, with default header values.

OK
*/
type PodsByDateRangeCountOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ModelsAggregateValuesByFieldResponse
}

// IsSuccess returns true when this pods by date range count o k response has a 2xx status code
func (o *PodsByDateRangeCountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pods by date range count o k response has a 3xx status code
func (o *PodsByDateRangeCountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pods by date range count o k response has a 4xx status code
func (o *PodsByDateRangeCountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pods by date range count o k response has a 5xx status code
func (o *PodsByDateRangeCountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pods by date range count o k response a status code equal to that given
func (o *PodsByDateRangeCountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pods by date range count o k response
func (o *PodsByDateRangeCountOK) Code() int {
	return 200
}

func (o *PodsByDateRangeCountOK) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/pods/count-by-date/v1][%d] podsByDateRangeCountOK  %+v", 200, o.Payload)
}

func (o *PodsByDateRangeCountOK) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/pods/count-by-date/v1][%d] podsByDateRangeCountOK  %+v", 200, o.Payload)
}

func (o *PodsByDateRangeCountOK) GetPayload() *models.ModelsAggregateValuesByFieldResponse {
	return o.Payload
}

func (o *PodsByDateRangeCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ModelsAggregateValuesByFieldResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPodsByDateRangeCountForbidden creates a PodsByDateRangeCountForbidden with default headers values
func NewPodsByDateRangeCountForbidden() *PodsByDateRangeCountForbidden {
	return &PodsByDateRangeCountForbidden{}
}

/*
PodsByDateRangeCountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PodsByDateRangeCountForbidden struct {

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

// IsSuccess returns true when this pods by date range count forbidden response has a 2xx status code
func (o *PodsByDateRangeCountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pods by date range count forbidden response has a 3xx status code
func (o *PodsByDateRangeCountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pods by date range count forbidden response has a 4xx status code
func (o *PodsByDateRangeCountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pods by date range count forbidden response has a 5xx status code
func (o *PodsByDateRangeCountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pods by date range count forbidden response a status code equal to that given
func (o *PodsByDateRangeCountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pods by date range count forbidden response
func (o *PodsByDateRangeCountForbidden) Code() int {
	return 403
}

func (o *PodsByDateRangeCountForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/pods/count-by-date/v1][%d] podsByDateRangeCountForbidden  %+v", 403, o.Payload)
}

func (o *PodsByDateRangeCountForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/pods/count-by-date/v1][%d] podsByDateRangeCountForbidden  %+v", 403, o.Payload)
}

func (o *PodsByDateRangeCountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *PodsByDateRangeCountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPodsByDateRangeCountTooManyRequests creates a PodsByDateRangeCountTooManyRequests with default headers values
func NewPodsByDateRangeCountTooManyRequests() *PodsByDateRangeCountTooManyRequests {
	return &PodsByDateRangeCountTooManyRequests{}
}

/*
PodsByDateRangeCountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PodsByDateRangeCountTooManyRequests struct {

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

// IsSuccess returns true when this pods by date range count too many requests response has a 2xx status code
func (o *PodsByDateRangeCountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pods by date range count too many requests response has a 3xx status code
func (o *PodsByDateRangeCountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pods by date range count too many requests response has a 4xx status code
func (o *PodsByDateRangeCountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this pods by date range count too many requests response has a 5xx status code
func (o *PodsByDateRangeCountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this pods by date range count too many requests response a status code equal to that given
func (o *PodsByDateRangeCountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the pods by date range count too many requests response
func (o *PodsByDateRangeCountTooManyRequests) Code() int {
	return 429
}

func (o *PodsByDateRangeCountTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/pods/count-by-date/v1][%d] podsByDateRangeCountTooManyRequests  %+v", 429, o.Payload)
}

func (o *PodsByDateRangeCountTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/pods/count-by-date/v1][%d] podsByDateRangeCountTooManyRequests  %+v", 429, o.Payload)
}

func (o *PodsByDateRangeCountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *PodsByDateRangeCountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPodsByDateRangeCountInternalServerError creates a PodsByDateRangeCountInternalServerError with default headers values
func NewPodsByDateRangeCountInternalServerError() *PodsByDateRangeCountInternalServerError {
	return &PodsByDateRangeCountInternalServerError{}
}

/*
PodsByDateRangeCountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PodsByDateRangeCountInternalServerError struct {

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

// IsSuccess returns true when this pods by date range count internal server error response has a 2xx status code
func (o *PodsByDateRangeCountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pods by date range count internal server error response has a 3xx status code
func (o *PodsByDateRangeCountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pods by date range count internal server error response has a 4xx status code
func (o *PodsByDateRangeCountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pods by date range count internal server error response has a 5xx status code
func (o *PodsByDateRangeCountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pods by date range count internal server error response a status code equal to that given
func (o *PodsByDateRangeCountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pods by date range count internal server error response
func (o *PodsByDateRangeCountInternalServerError) Code() int {
	return 500
}

func (o *PodsByDateRangeCountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/pods/count-by-date/v1][%d] podsByDateRangeCountInternalServerError  %+v", 500, o.Payload)
}

func (o *PodsByDateRangeCountInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/pods/count-by-date/v1][%d] podsByDateRangeCountInternalServerError  %+v", 500, o.Payload)
}

func (o *PodsByDateRangeCountInternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *PodsByDateRangeCountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
