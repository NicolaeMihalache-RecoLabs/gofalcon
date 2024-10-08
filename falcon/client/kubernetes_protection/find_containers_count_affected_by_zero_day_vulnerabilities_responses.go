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

// FindContainersCountAffectedByZeroDayVulnerabilitiesReader is a Reader for the FindContainersCountAffectedByZeroDayVulnerabilities structure.
type FindContainersCountAffectedByZeroDayVulnerabilitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindContainersCountAffectedByZeroDayVulnerabilitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewFindContainersCountAffectedByZeroDayVulnerabilitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewFindContainersCountAffectedByZeroDayVulnerabilitiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFindContainersCountAffectedByZeroDayVulnerabilitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/aggregates/containers/count-by-zero-day/v1] FindContainersCountAffectedByZeroDayVulnerabilities", response, response.Code())
	}
}

// NewFindContainersCountAffectedByZeroDayVulnerabilitiesOK creates a FindContainersCountAffectedByZeroDayVulnerabilitiesOK with default headers values
func NewFindContainersCountAffectedByZeroDayVulnerabilitiesOK() *FindContainersCountAffectedByZeroDayVulnerabilitiesOK {
	return &FindContainersCountAffectedByZeroDayVulnerabilitiesOK{}
}

/*
FindContainersCountAffectedByZeroDayVulnerabilitiesOK describes a response with status code 200, with default header values.

OK
*/
type FindContainersCountAffectedByZeroDayVulnerabilitiesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.CommonCountResponse
}

// IsSuccess returns true when this find containers count affected by zero day vulnerabilities o k response has a 2xx status code
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find containers count affected by zero day vulnerabilities o k response has a 3xx status code
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find containers count affected by zero day vulnerabilities o k response has a 4xx status code
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find containers count affected by zero day vulnerabilities o k response has a 5xx status code
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find containers count affected by zero day vulnerabilities o k response a status code equal to that given
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find containers count affected by zero day vulnerabilities o k response
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesOK) Code() int {
	return 200
}

func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesOK) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/count-by-zero-day/v1][%d] findContainersCountAffectedByZeroDayVulnerabilitiesOK  %+v", 200, o.Payload)
}

func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesOK) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/count-by-zero-day/v1][%d] findContainersCountAffectedByZeroDayVulnerabilitiesOK  %+v", 200, o.Payload)
}

func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesOK) GetPayload() *models.CommonCountResponse {
	return o.Payload
}

func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.CommonCountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindContainersCountAffectedByZeroDayVulnerabilitiesForbidden creates a FindContainersCountAffectedByZeroDayVulnerabilitiesForbidden with default headers values
func NewFindContainersCountAffectedByZeroDayVulnerabilitiesForbidden() *FindContainersCountAffectedByZeroDayVulnerabilitiesForbidden {
	return &FindContainersCountAffectedByZeroDayVulnerabilitiesForbidden{}
}

/*
FindContainersCountAffectedByZeroDayVulnerabilitiesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type FindContainersCountAffectedByZeroDayVulnerabilitiesForbidden struct {

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

// IsSuccess returns true when this find containers count affected by zero day vulnerabilities forbidden response has a 2xx status code
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this find containers count affected by zero day vulnerabilities forbidden response has a 3xx status code
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find containers count affected by zero day vulnerabilities forbidden response has a 4xx status code
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this find containers count affected by zero day vulnerabilities forbidden response has a 5xx status code
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this find containers count affected by zero day vulnerabilities forbidden response a status code equal to that given
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the find containers count affected by zero day vulnerabilities forbidden response
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesForbidden) Code() int {
	return 403
}

func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/count-by-zero-day/v1][%d] findContainersCountAffectedByZeroDayVulnerabilitiesForbidden  %+v", 403, o.Payload)
}

func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/count-by-zero-day/v1][%d] findContainersCountAffectedByZeroDayVulnerabilitiesForbidden  %+v", 403, o.Payload)
}

func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewFindContainersCountAffectedByZeroDayVulnerabilitiesTooManyRequests creates a FindContainersCountAffectedByZeroDayVulnerabilitiesTooManyRequests with default headers values
func NewFindContainersCountAffectedByZeroDayVulnerabilitiesTooManyRequests() *FindContainersCountAffectedByZeroDayVulnerabilitiesTooManyRequests {
	return &FindContainersCountAffectedByZeroDayVulnerabilitiesTooManyRequests{}
}

/*
FindContainersCountAffectedByZeroDayVulnerabilitiesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type FindContainersCountAffectedByZeroDayVulnerabilitiesTooManyRequests struct {

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

// IsSuccess returns true when this find containers count affected by zero day vulnerabilities too many requests response has a 2xx status code
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this find containers count affected by zero day vulnerabilities too many requests response has a 3xx status code
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find containers count affected by zero day vulnerabilities too many requests response has a 4xx status code
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this find containers count affected by zero day vulnerabilities too many requests response has a 5xx status code
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this find containers count affected by zero day vulnerabilities too many requests response a status code equal to that given
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the find containers count affected by zero day vulnerabilities too many requests response
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesTooManyRequests) Code() int {
	return 429
}

func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/count-by-zero-day/v1][%d] findContainersCountAffectedByZeroDayVulnerabilitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/count-by-zero-day/v1][%d] findContainersCountAffectedByZeroDayVulnerabilitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewFindContainersCountAffectedByZeroDayVulnerabilitiesInternalServerError creates a FindContainersCountAffectedByZeroDayVulnerabilitiesInternalServerError with default headers values
func NewFindContainersCountAffectedByZeroDayVulnerabilitiesInternalServerError() *FindContainersCountAffectedByZeroDayVulnerabilitiesInternalServerError {
	return &FindContainersCountAffectedByZeroDayVulnerabilitiesInternalServerError{}
}

/*
FindContainersCountAffectedByZeroDayVulnerabilitiesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type FindContainersCountAffectedByZeroDayVulnerabilitiesInternalServerError struct {

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

// IsSuccess returns true when this find containers count affected by zero day vulnerabilities internal server error response has a 2xx status code
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this find containers count affected by zero day vulnerabilities internal server error response has a 3xx status code
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find containers count affected by zero day vulnerabilities internal server error response has a 4xx status code
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this find containers count affected by zero day vulnerabilities internal server error response has a 5xx status code
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this find containers count affected by zero day vulnerabilities internal server error response a status code equal to that given
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the find containers count affected by zero day vulnerabilities internal server error response
func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesInternalServerError) Code() int {
	return 500
}

func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/count-by-zero-day/v1][%d] findContainersCountAffectedByZeroDayVulnerabilitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/count-by-zero-day/v1][%d] findContainersCountAffectedByZeroDayVulnerabilitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesInternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *FindContainersCountAffectedByZeroDayVulnerabilitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
