// Code generated by go-swagger; DO NOT EDIT.

package compliance_assessments

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

// ExtAggregateFailedRulesByClustersReader is a Reader for the ExtAggregateFailedRulesByClusters structure.
type ExtAggregateFailedRulesByClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtAggregateFailedRulesByClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtAggregateFailedRulesByClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExtAggregateFailedRulesByClustersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewExtAggregateFailedRulesByClustersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExtAggregateFailedRulesByClustersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewExtAggregateFailedRulesByClustersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewExtAggregateFailedRulesByClustersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-compliance/aggregates/failed-rules-by-clusters/v2] extAggregateFailedRulesByClusters", response, response.Code())
	}
}

// NewExtAggregateFailedRulesByClustersOK creates a ExtAggregateFailedRulesByClustersOK with default headers values
func NewExtAggregateFailedRulesByClustersOK() *ExtAggregateFailedRulesByClustersOK {
	return &ExtAggregateFailedRulesByClustersOK{}
}

/*
ExtAggregateFailedRulesByClustersOK describes a response with status code 200, with default header values.

OK
*/
type ExtAggregateFailedRulesByClustersOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAggregateFailedRulesByClustersResponse
}

// IsSuccess returns true when this ext aggregate failed rules by clusters o k response has a 2xx status code
func (o *ExtAggregateFailedRulesByClustersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ext aggregate failed rules by clusters o k response has a 3xx status code
func (o *ExtAggregateFailedRulesByClustersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ext aggregate failed rules by clusters o k response has a 4xx status code
func (o *ExtAggregateFailedRulesByClustersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ext aggregate failed rules by clusters o k response has a 5xx status code
func (o *ExtAggregateFailedRulesByClustersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ext aggregate failed rules by clusters o k response a status code equal to that given
func (o *ExtAggregateFailedRulesByClustersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ext aggregate failed rules by clusters o k response
func (o *ExtAggregateFailedRulesByClustersOK) Code() int {
	return 200
}

func (o *ExtAggregateFailedRulesByClustersOK) Error() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-rules-by-clusters/v2][%d] extAggregateFailedRulesByClustersOK  %+v", 200, o.Payload)
}

func (o *ExtAggregateFailedRulesByClustersOK) String() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-rules-by-clusters/v2][%d] extAggregateFailedRulesByClustersOK  %+v", 200, o.Payload)
}

func (o *ExtAggregateFailedRulesByClustersOK) GetPayload() *models.DomainAggregateFailedRulesByClustersResponse {
	return o.Payload
}

func (o *ExtAggregateFailedRulesByClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAggregateFailedRulesByClustersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtAggregateFailedRulesByClustersBadRequest creates a ExtAggregateFailedRulesByClustersBadRequest with default headers values
func NewExtAggregateFailedRulesByClustersBadRequest() *ExtAggregateFailedRulesByClustersBadRequest {
	return &ExtAggregateFailedRulesByClustersBadRequest{}
}

/*
ExtAggregateFailedRulesByClustersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExtAggregateFailedRulesByClustersBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAggregateFailedRulesByClustersResponse
}

// IsSuccess returns true when this ext aggregate failed rules by clusters bad request response has a 2xx status code
func (o *ExtAggregateFailedRulesByClustersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ext aggregate failed rules by clusters bad request response has a 3xx status code
func (o *ExtAggregateFailedRulesByClustersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ext aggregate failed rules by clusters bad request response has a 4xx status code
func (o *ExtAggregateFailedRulesByClustersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ext aggregate failed rules by clusters bad request response has a 5xx status code
func (o *ExtAggregateFailedRulesByClustersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ext aggregate failed rules by clusters bad request response a status code equal to that given
func (o *ExtAggregateFailedRulesByClustersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the ext aggregate failed rules by clusters bad request response
func (o *ExtAggregateFailedRulesByClustersBadRequest) Code() int {
	return 400
}

func (o *ExtAggregateFailedRulesByClustersBadRequest) Error() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-rules-by-clusters/v2][%d] extAggregateFailedRulesByClustersBadRequest  %+v", 400, o.Payload)
}

func (o *ExtAggregateFailedRulesByClustersBadRequest) String() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-rules-by-clusters/v2][%d] extAggregateFailedRulesByClustersBadRequest  %+v", 400, o.Payload)
}

func (o *ExtAggregateFailedRulesByClustersBadRequest) GetPayload() *models.DomainAggregateFailedRulesByClustersResponse {
	return o.Payload
}

func (o *ExtAggregateFailedRulesByClustersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAggregateFailedRulesByClustersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtAggregateFailedRulesByClustersUnauthorized creates a ExtAggregateFailedRulesByClustersUnauthorized with default headers values
func NewExtAggregateFailedRulesByClustersUnauthorized() *ExtAggregateFailedRulesByClustersUnauthorized {
	return &ExtAggregateFailedRulesByClustersUnauthorized{}
}

/*
ExtAggregateFailedRulesByClustersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ExtAggregateFailedRulesByClustersUnauthorized struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAggregateFailedRulesByClustersResponse
}

// IsSuccess returns true when this ext aggregate failed rules by clusters unauthorized response has a 2xx status code
func (o *ExtAggregateFailedRulesByClustersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ext aggregate failed rules by clusters unauthorized response has a 3xx status code
func (o *ExtAggregateFailedRulesByClustersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ext aggregate failed rules by clusters unauthorized response has a 4xx status code
func (o *ExtAggregateFailedRulesByClustersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ext aggregate failed rules by clusters unauthorized response has a 5xx status code
func (o *ExtAggregateFailedRulesByClustersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ext aggregate failed rules by clusters unauthorized response a status code equal to that given
func (o *ExtAggregateFailedRulesByClustersUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the ext aggregate failed rules by clusters unauthorized response
func (o *ExtAggregateFailedRulesByClustersUnauthorized) Code() int {
	return 401
}

func (o *ExtAggregateFailedRulesByClustersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-rules-by-clusters/v2][%d] extAggregateFailedRulesByClustersUnauthorized  %+v", 401, o.Payload)
}

func (o *ExtAggregateFailedRulesByClustersUnauthorized) String() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-rules-by-clusters/v2][%d] extAggregateFailedRulesByClustersUnauthorized  %+v", 401, o.Payload)
}

func (o *ExtAggregateFailedRulesByClustersUnauthorized) GetPayload() *models.DomainAggregateFailedRulesByClustersResponse {
	return o.Payload
}

func (o *ExtAggregateFailedRulesByClustersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAggregateFailedRulesByClustersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtAggregateFailedRulesByClustersForbidden creates a ExtAggregateFailedRulesByClustersForbidden with default headers values
func NewExtAggregateFailedRulesByClustersForbidden() *ExtAggregateFailedRulesByClustersForbidden {
	return &ExtAggregateFailedRulesByClustersForbidden{}
}

/*
ExtAggregateFailedRulesByClustersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ExtAggregateFailedRulesByClustersForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAggregateFailedRulesByClustersResponse
}

// IsSuccess returns true when this ext aggregate failed rules by clusters forbidden response has a 2xx status code
func (o *ExtAggregateFailedRulesByClustersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ext aggregate failed rules by clusters forbidden response has a 3xx status code
func (o *ExtAggregateFailedRulesByClustersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ext aggregate failed rules by clusters forbidden response has a 4xx status code
func (o *ExtAggregateFailedRulesByClustersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ext aggregate failed rules by clusters forbidden response has a 5xx status code
func (o *ExtAggregateFailedRulesByClustersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ext aggregate failed rules by clusters forbidden response a status code equal to that given
func (o *ExtAggregateFailedRulesByClustersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the ext aggregate failed rules by clusters forbidden response
func (o *ExtAggregateFailedRulesByClustersForbidden) Code() int {
	return 403
}

func (o *ExtAggregateFailedRulesByClustersForbidden) Error() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-rules-by-clusters/v2][%d] extAggregateFailedRulesByClustersForbidden  %+v", 403, o.Payload)
}

func (o *ExtAggregateFailedRulesByClustersForbidden) String() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-rules-by-clusters/v2][%d] extAggregateFailedRulesByClustersForbidden  %+v", 403, o.Payload)
}

func (o *ExtAggregateFailedRulesByClustersForbidden) GetPayload() *models.DomainAggregateFailedRulesByClustersResponse {
	return o.Payload
}

func (o *ExtAggregateFailedRulesByClustersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAggregateFailedRulesByClustersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtAggregateFailedRulesByClustersTooManyRequests creates a ExtAggregateFailedRulesByClustersTooManyRequests with default headers values
func NewExtAggregateFailedRulesByClustersTooManyRequests() *ExtAggregateFailedRulesByClustersTooManyRequests {
	return &ExtAggregateFailedRulesByClustersTooManyRequests{}
}

/*
ExtAggregateFailedRulesByClustersTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ExtAggregateFailedRulesByClustersTooManyRequests struct {

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

// IsSuccess returns true when this ext aggregate failed rules by clusters too many requests response has a 2xx status code
func (o *ExtAggregateFailedRulesByClustersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ext aggregate failed rules by clusters too many requests response has a 3xx status code
func (o *ExtAggregateFailedRulesByClustersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ext aggregate failed rules by clusters too many requests response has a 4xx status code
func (o *ExtAggregateFailedRulesByClustersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this ext aggregate failed rules by clusters too many requests response has a 5xx status code
func (o *ExtAggregateFailedRulesByClustersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this ext aggregate failed rules by clusters too many requests response a status code equal to that given
func (o *ExtAggregateFailedRulesByClustersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the ext aggregate failed rules by clusters too many requests response
func (o *ExtAggregateFailedRulesByClustersTooManyRequests) Code() int {
	return 429
}

func (o *ExtAggregateFailedRulesByClustersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-rules-by-clusters/v2][%d] extAggregateFailedRulesByClustersTooManyRequests  %+v", 429, o.Payload)
}

func (o *ExtAggregateFailedRulesByClustersTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-rules-by-clusters/v2][%d] extAggregateFailedRulesByClustersTooManyRequests  %+v", 429, o.Payload)
}

func (o *ExtAggregateFailedRulesByClustersTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ExtAggregateFailedRulesByClustersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewExtAggregateFailedRulesByClustersInternalServerError creates a ExtAggregateFailedRulesByClustersInternalServerError with default headers values
func NewExtAggregateFailedRulesByClustersInternalServerError() *ExtAggregateFailedRulesByClustersInternalServerError {
	return &ExtAggregateFailedRulesByClustersInternalServerError{}
}

/*
ExtAggregateFailedRulesByClustersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ExtAggregateFailedRulesByClustersInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAggregateFailedRulesByClustersResponse
}

// IsSuccess returns true when this ext aggregate failed rules by clusters internal server error response has a 2xx status code
func (o *ExtAggregateFailedRulesByClustersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ext aggregate failed rules by clusters internal server error response has a 3xx status code
func (o *ExtAggregateFailedRulesByClustersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ext aggregate failed rules by clusters internal server error response has a 4xx status code
func (o *ExtAggregateFailedRulesByClustersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ext aggregate failed rules by clusters internal server error response has a 5xx status code
func (o *ExtAggregateFailedRulesByClustersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ext aggregate failed rules by clusters internal server error response a status code equal to that given
func (o *ExtAggregateFailedRulesByClustersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the ext aggregate failed rules by clusters internal server error response
func (o *ExtAggregateFailedRulesByClustersInternalServerError) Code() int {
	return 500
}

func (o *ExtAggregateFailedRulesByClustersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-rules-by-clusters/v2][%d] extAggregateFailedRulesByClustersInternalServerError  %+v", 500, o.Payload)
}

func (o *ExtAggregateFailedRulesByClustersInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-rules-by-clusters/v2][%d] extAggregateFailedRulesByClustersInternalServerError  %+v", 500, o.Payload)
}

func (o *ExtAggregateFailedRulesByClustersInternalServerError) GetPayload() *models.DomainAggregateFailedRulesByClustersResponse {
	return o.Payload
}

func (o *ExtAggregateFailedRulesByClustersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAggregateFailedRulesByClustersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
