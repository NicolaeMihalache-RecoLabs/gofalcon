// Code generated by go-swagger; DO NOT EDIT.

package container_vulnerabilities

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

// ReadVulnerabilitiesByImageCountReader is a Reader for the ReadVulnerabilitiesByImageCount structure.
type ReadVulnerabilitiesByImageCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadVulnerabilitiesByImageCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadVulnerabilitiesByImageCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewReadVulnerabilitiesByImageCountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewReadVulnerabilitiesByImageCountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadVulnerabilitiesByImageCountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/combined/vulnerabilities/by-image-count/v1] ReadVulnerabilitiesByImageCount", response, response.Code())
	}
}

// NewReadVulnerabilitiesByImageCountOK creates a ReadVulnerabilitiesByImageCountOK with default headers values
func NewReadVulnerabilitiesByImageCountOK() *ReadVulnerabilitiesByImageCountOK {
	return &ReadVulnerabilitiesByImageCountOK{}
}

/*
ReadVulnerabilitiesByImageCountOK describes a response with status code 200, with default header values.

OK
*/
type ReadVulnerabilitiesByImageCountOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.VulnerabilitiesAPIVulnByImageCount
}

// IsSuccess returns true when this read vulnerabilities by image count o k response has a 2xx status code
func (o *ReadVulnerabilitiesByImageCountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read vulnerabilities by image count o k response has a 3xx status code
func (o *ReadVulnerabilitiesByImageCountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read vulnerabilities by image count o k response has a 4xx status code
func (o *ReadVulnerabilitiesByImageCountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read vulnerabilities by image count o k response has a 5xx status code
func (o *ReadVulnerabilitiesByImageCountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read vulnerabilities by image count o k response a status code equal to that given
func (o *ReadVulnerabilitiesByImageCountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read vulnerabilities by image count o k response
func (o *ReadVulnerabilitiesByImageCountOK) Code() int {
	return 200
}

func (o *ReadVulnerabilitiesByImageCountOK) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/vulnerabilities/by-image-count/v1][%d] readVulnerabilitiesByImageCountOK  %+v", 200, o.Payload)
}

func (o *ReadVulnerabilitiesByImageCountOK) String() string {
	return fmt.Sprintf("[GET /container-security/combined/vulnerabilities/by-image-count/v1][%d] readVulnerabilitiesByImageCountOK  %+v", 200, o.Payload)
}

func (o *ReadVulnerabilitiesByImageCountOK) GetPayload() *models.VulnerabilitiesAPIVulnByImageCount {
	return o.Payload
}

func (o *ReadVulnerabilitiesByImageCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.VulnerabilitiesAPIVulnByImageCount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadVulnerabilitiesByImageCountForbidden creates a ReadVulnerabilitiesByImageCountForbidden with default headers values
func NewReadVulnerabilitiesByImageCountForbidden() *ReadVulnerabilitiesByImageCountForbidden {
	return &ReadVulnerabilitiesByImageCountForbidden{}
}

/*
ReadVulnerabilitiesByImageCountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ReadVulnerabilitiesByImageCountForbidden struct {

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

// IsSuccess returns true when this read vulnerabilities by image count forbidden response has a 2xx status code
func (o *ReadVulnerabilitiesByImageCountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read vulnerabilities by image count forbidden response has a 3xx status code
func (o *ReadVulnerabilitiesByImageCountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read vulnerabilities by image count forbidden response has a 4xx status code
func (o *ReadVulnerabilitiesByImageCountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this read vulnerabilities by image count forbidden response has a 5xx status code
func (o *ReadVulnerabilitiesByImageCountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this read vulnerabilities by image count forbidden response a status code equal to that given
func (o *ReadVulnerabilitiesByImageCountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the read vulnerabilities by image count forbidden response
func (o *ReadVulnerabilitiesByImageCountForbidden) Code() int {
	return 403
}

func (o *ReadVulnerabilitiesByImageCountForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/vulnerabilities/by-image-count/v1][%d] readVulnerabilitiesByImageCountForbidden  %+v", 403, o.Payload)
}

func (o *ReadVulnerabilitiesByImageCountForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/combined/vulnerabilities/by-image-count/v1][%d] readVulnerabilitiesByImageCountForbidden  %+v", 403, o.Payload)
}

func (o *ReadVulnerabilitiesByImageCountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ReadVulnerabilitiesByImageCountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReadVulnerabilitiesByImageCountTooManyRequests creates a ReadVulnerabilitiesByImageCountTooManyRequests with default headers values
func NewReadVulnerabilitiesByImageCountTooManyRequests() *ReadVulnerabilitiesByImageCountTooManyRequests {
	return &ReadVulnerabilitiesByImageCountTooManyRequests{}
}

/*
ReadVulnerabilitiesByImageCountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ReadVulnerabilitiesByImageCountTooManyRequests struct {

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

// IsSuccess returns true when this read vulnerabilities by image count too many requests response has a 2xx status code
func (o *ReadVulnerabilitiesByImageCountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read vulnerabilities by image count too many requests response has a 3xx status code
func (o *ReadVulnerabilitiesByImageCountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read vulnerabilities by image count too many requests response has a 4xx status code
func (o *ReadVulnerabilitiesByImageCountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this read vulnerabilities by image count too many requests response has a 5xx status code
func (o *ReadVulnerabilitiesByImageCountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this read vulnerabilities by image count too many requests response a status code equal to that given
func (o *ReadVulnerabilitiesByImageCountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the read vulnerabilities by image count too many requests response
func (o *ReadVulnerabilitiesByImageCountTooManyRequests) Code() int {
	return 429
}

func (o *ReadVulnerabilitiesByImageCountTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/vulnerabilities/by-image-count/v1][%d] readVulnerabilitiesByImageCountTooManyRequests  %+v", 429, o.Payload)
}

func (o *ReadVulnerabilitiesByImageCountTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/combined/vulnerabilities/by-image-count/v1][%d] readVulnerabilitiesByImageCountTooManyRequests  %+v", 429, o.Payload)
}

func (o *ReadVulnerabilitiesByImageCountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ReadVulnerabilitiesByImageCountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReadVulnerabilitiesByImageCountInternalServerError creates a ReadVulnerabilitiesByImageCountInternalServerError with default headers values
func NewReadVulnerabilitiesByImageCountInternalServerError() *ReadVulnerabilitiesByImageCountInternalServerError {
	return &ReadVulnerabilitiesByImageCountInternalServerError{}
}

/*
ReadVulnerabilitiesByImageCountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ReadVulnerabilitiesByImageCountInternalServerError struct {

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

// IsSuccess returns true when this read vulnerabilities by image count internal server error response has a 2xx status code
func (o *ReadVulnerabilitiesByImageCountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read vulnerabilities by image count internal server error response has a 3xx status code
func (o *ReadVulnerabilitiesByImageCountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read vulnerabilities by image count internal server error response has a 4xx status code
func (o *ReadVulnerabilitiesByImageCountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read vulnerabilities by image count internal server error response has a 5xx status code
func (o *ReadVulnerabilitiesByImageCountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read vulnerabilities by image count internal server error response a status code equal to that given
func (o *ReadVulnerabilitiesByImageCountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read vulnerabilities by image count internal server error response
func (o *ReadVulnerabilitiesByImageCountInternalServerError) Code() int {
	return 500
}

func (o *ReadVulnerabilitiesByImageCountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/vulnerabilities/by-image-count/v1][%d] readVulnerabilitiesByImageCountInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadVulnerabilitiesByImageCountInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/combined/vulnerabilities/by-image-count/v1][%d] readVulnerabilitiesByImageCountInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadVulnerabilitiesByImageCountInternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *ReadVulnerabilitiesByImageCountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
