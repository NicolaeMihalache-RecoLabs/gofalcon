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

// PodCombinedReader is a Reader for the PodCombined structure.
type PodCombinedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PodCombinedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPodCombinedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPodCombinedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPodCombinedTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPodCombinedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/combined/pods/v1] PodCombined", response, response.Code())
	}
}

// NewPodCombinedOK creates a PodCombinedOK with default headers values
func NewPodCombinedOK() *PodCombinedOK {
	return &PodCombinedOK{}
}

/*
PodCombinedOK describes a response with status code 200, with default header values.

OK
*/
type PodCombinedOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ModelsPodEntityResponse
}

// IsSuccess returns true when this pod combined o k response has a 2xx status code
func (o *PodCombinedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pod combined o k response has a 3xx status code
func (o *PodCombinedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pod combined o k response has a 4xx status code
func (o *PodCombinedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pod combined o k response has a 5xx status code
func (o *PodCombinedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pod combined o k response a status code equal to that given
func (o *PodCombinedOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pod combined o k response
func (o *PodCombinedOK) Code() int {
	return 200
}

func (o *PodCombinedOK) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/pods/v1][%d] podCombinedOK  %+v", 200, o.Payload)
}

func (o *PodCombinedOK) String() string {
	return fmt.Sprintf("[GET /container-security/combined/pods/v1][%d] podCombinedOK  %+v", 200, o.Payload)
}

func (o *PodCombinedOK) GetPayload() *models.ModelsPodEntityResponse {
	return o.Payload
}

func (o *PodCombinedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ModelsPodEntityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPodCombinedForbidden creates a PodCombinedForbidden with default headers values
func NewPodCombinedForbidden() *PodCombinedForbidden {
	return &PodCombinedForbidden{}
}

/*
PodCombinedForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PodCombinedForbidden struct {

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

// IsSuccess returns true when this pod combined forbidden response has a 2xx status code
func (o *PodCombinedForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pod combined forbidden response has a 3xx status code
func (o *PodCombinedForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pod combined forbidden response has a 4xx status code
func (o *PodCombinedForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pod combined forbidden response has a 5xx status code
func (o *PodCombinedForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pod combined forbidden response a status code equal to that given
func (o *PodCombinedForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pod combined forbidden response
func (o *PodCombinedForbidden) Code() int {
	return 403
}

func (o *PodCombinedForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/pods/v1][%d] podCombinedForbidden  %+v", 403, o.Payload)
}

func (o *PodCombinedForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/combined/pods/v1][%d] podCombinedForbidden  %+v", 403, o.Payload)
}

func (o *PodCombinedForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *PodCombinedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPodCombinedTooManyRequests creates a PodCombinedTooManyRequests with default headers values
func NewPodCombinedTooManyRequests() *PodCombinedTooManyRequests {
	return &PodCombinedTooManyRequests{}
}

/*
PodCombinedTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PodCombinedTooManyRequests struct {

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

// IsSuccess returns true when this pod combined too many requests response has a 2xx status code
func (o *PodCombinedTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pod combined too many requests response has a 3xx status code
func (o *PodCombinedTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pod combined too many requests response has a 4xx status code
func (o *PodCombinedTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this pod combined too many requests response has a 5xx status code
func (o *PodCombinedTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this pod combined too many requests response a status code equal to that given
func (o *PodCombinedTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the pod combined too many requests response
func (o *PodCombinedTooManyRequests) Code() int {
	return 429
}

func (o *PodCombinedTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/pods/v1][%d] podCombinedTooManyRequests  %+v", 429, o.Payload)
}

func (o *PodCombinedTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/combined/pods/v1][%d] podCombinedTooManyRequests  %+v", 429, o.Payload)
}

func (o *PodCombinedTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *PodCombinedTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPodCombinedInternalServerError creates a PodCombinedInternalServerError with default headers values
func NewPodCombinedInternalServerError() *PodCombinedInternalServerError {
	return &PodCombinedInternalServerError{}
}

/*
PodCombinedInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PodCombinedInternalServerError struct {

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

// IsSuccess returns true when this pod combined internal server error response has a 2xx status code
func (o *PodCombinedInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pod combined internal server error response has a 3xx status code
func (o *PodCombinedInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pod combined internal server error response has a 4xx status code
func (o *PodCombinedInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pod combined internal server error response has a 5xx status code
func (o *PodCombinedInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pod combined internal server error response a status code equal to that given
func (o *PodCombinedInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pod combined internal server error response
func (o *PodCombinedInternalServerError) Code() int {
	return 500
}

func (o *PodCombinedInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/pods/v1][%d] podCombinedInternalServerError  %+v", 500, o.Payload)
}

func (o *PodCombinedInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/combined/pods/v1][%d] podCombinedInternalServerError  %+v", 500, o.Payload)
}

func (o *PodCombinedInternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *PodCombinedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
