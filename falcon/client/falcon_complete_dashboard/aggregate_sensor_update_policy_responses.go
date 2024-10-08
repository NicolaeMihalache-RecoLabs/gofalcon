// Code generated by go-swagger; DO NOT EDIT.

package falcon_complete_dashboard

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

// AggregateSensorUpdatePolicyReader is a Reader for the AggregateSensorUpdatePolicy structure.
type AggregateSensorUpdatePolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregateSensorUpdatePolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregateSensorUpdatePolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAggregateSensorUpdatePolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAggregateSensorUpdatePolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAggregateSensorUpdatePolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAggregateSensorUpdatePolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /falcon-complete-dashboards/aggregates/sensor-update-policies/v1] AggregateSensorUpdatePolicy", response, response.Code())
	}
}

// NewAggregateSensorUpdatePolicyOK creates a AggregateSensorUpdatePolicyOK with default headers values
func NewAggregateSensorUpdatePolicyOK() *AggregateSensorUpdatePolicyOK {
	return &AggregateSensorUpdatePolicyOK{}
}

/*
AggregateSensorUpdatePolicyOK describes a response with status code 200, with default header values.

OK
*/
type AggregateSensorUpdatePolicyOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaAggregatesResponse
}

// IsSuccess returns true when this aggregate sensor update policy o k response has a 2xx status code
func (o *AggregateSensorUpdatePolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aggregate sensor update policy o k response has a 3xx status code
func (o *AggregateSensorUpdatePolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate sensor update policy o k response has a 4xx status code
func (o *AggregateSensorUpdatePolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate sensor update policy o k response has a 5xx status code
func (o *AggregateSensorUpdatePolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate sensor update policy o k response a status code equal to that given
func (o *AggregateSensorUpdatePolicyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the aggregate sensor update policy o k response
func (o *AggregateSensorUpdatePolicyOK) Code() int {
	return 200
}

func (o *AggregateSensorUpdatePolicyOK) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/sensor-update-policies/v1][%d] aggregateSensorUpdatePolicyOK  %+v", 200, o.Payload)
}

func (o *AggregateSensorUpdatePolicyOK) String() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/sensor-update-policies/v1][%d] aggregateSensorUpdatePolicyOK  %+v", 200, o.Payload)
}

func (o *AggregateSensorUpdatePolicyOK) GetPayload() *models.MsaAggregatesResponse {
	return o.Payload
}

func (o *AggregateSensorUpdatePolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateSensorUpdatePolicyBadRequest creates a AggregateSensorUpdatePolicyBadRequest with default headers values
func NewAggregateSensorUpdatePolicyBadRequest() *AggregateSensorUpdatePolicyBadRequest {
	return &AggregateSensorUpdatePolicyBadRequest{}
}

/*
AggregateSensorUpdatePolicyBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AggregateSensorUpdatePolicyBadRequest struct {

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

// IsSuccess returns true when this aggregate sensor update policy bad request response has a 2xx status code
func (o *AggregateSensorUpdatePolicyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate sensor update policy bad request response has a 3xx status code
func (o *AggregateSensorUpdatePolicyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate sensor update policy bad request response has a 4xx status code
func (o *AggregateSensorUpdatePolicyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate sensor update policy bad request response has a 5xx status code
func (o *AggregateSensorUpdatePolicyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate sensor update policy bad request response a status code equal to that given
func (o *AggregateSensorUpdatePolicyBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the aggregate sensor update policy bad request response
func (o *AggregateSensorUpdatePolicyBadRequest) Code() int {
	return 400
}

func (o *AggregateSensorUpdatePolicyBadRequest) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/sensor-update-policies/v1][%d] aggregateSensorUpdatePolicyBadRequest  %+v", 400, o.Payload)
}

func (o *AggregateSensorUpdatePolicyBadRequest) String() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/sensor-update-policies/v1][%d] aggregateSensorUpdatePolicyBadRequest  %+v", 400, o.Payload)
}

func (o *AggregateSensorUpdatePolicyBadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *AggregateSensorUpdatePolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregateSensorUpdatePolicyForbidden creates a AggregateSensorUpdatePolicyForbidden with default headers values
func NewAggregateSensorUpdatePolicyForbidden() *AggregateSensorUpdatePolicyForbidden {
	return &AggregateSensorUpdatePolicyForbidden{}
}

/*
AggregateSensorUpdatePolicyForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AggregateSensorUpdatePolicyForbidden struct {

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

// IsSuccess returns true when this aggregate sensor update policy forbidden response has a 2xx status code
func (o *AggregateSensorUpdatePolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate sensor update policy forbidden response has a 3xx status code
func (o *AggregateSensorUpdatePolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate sensor update policy forbidden response has a 4xx status code
func (o *AggregateSensorUpdatePolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate sensor update policy forbidden response has a 5xx status code
func (o *AggregateSensorUpdatePolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate sensor update policy forbidden response a status code equal to that given
func (o *AggregateSensorUpdatePolicyForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the aggregate sensor update policy forbidden response
func (o *AggregateSensorUpdatePolicyForbidden) Code() int {
	return 403
}

func (o *AggregateSensorUpdatePolicyForbidden) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/sensor-update-policies/v1][%d] aggregateSensorUpdatePolicyForbidden  %+v", 403, o.Payload)
}

func (o *AggregateSensorUpdatePolicyForbidden) String() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/sensor-update-policies/v1][%d] aggregateSensorUpdatePolicyForbidden  %+v", 403, o.Payload)
}

func (o *AggregateSensorUpdatePolicyForbidden) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *AggregateSensorUpdatePolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregateSensorUpdatePolicyTooManyRequests creates a AggregateSensorUpdatePolicyTooManyRequests with default headers values
func NewAggregateSensorUpdatePolicyTooManyRequests() *AggregateSensorUpdatePolicyTooManyRequests {
	return &AggregateSensorUpdatePolicyTooManyRequests{}
}

/*
AggregateSensorUpdatePolicyTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AggregateSensorUpdatePolicyTooManyRequests struct {

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

// IsSuccess returns true when this aggregate sensor update policy too many requests response has a 2xx status code
func (o *AggregateSensorUpdatePolicyTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate sensor update policy too many requests response has a 3xx status code
func (o *AggregateSensorUpdatePolicyTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate sensor update policy too many requests response has a 4xx status code
func (o *AggregateSensorUpdatePolicyTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate sensor update policy too many requests response has a 5xx status code
func (o *AggregateSensorUpdatePolicyTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate sensor update policy too many requests response a status code equal to that given
func (o *AggregateSensorUpdatePolicyTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the aggregate sensor update policy too many requests response
func (o *AggregateSensorUpdatePolicyTooManyRequests) Code() int {
	return 429
}

func (o *AggregateSensorUpdatePolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/sensor-update-policies/v1][%d] aggregateSensorUpdatePolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregateSensorUpdatePolicyTooManyRequests) String() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/sensor-update-policies/v1][%d] aggregateSensorUpdatePolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregateSensorUpdatePolicyTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregateSensorUpdatePolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregateSensorUpdatePolicyInternalServerError creates a AggregateSensorUpdatePolicyInternalServerError with default headers values
func NewAggregateSensorUpdatePolicyInternalServerError() *AggregateSensorUpdatePolicyInternalServerError {
	return &AggregateSensorUpdatePolicyInternalServerError{}
}

/*
AggregateSensorUpdatePolicyInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type AggregateSensorUpdatePolicyInternalServerError struct {

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

// IsSuccess returns true when this aggregate sensor update policy internal server error response has a 2xx status code
func (o *AggregateSensorUpdatePolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate sensor update policy internal server error response has a 3xx status code
func (o *AggregateSensorUpdatePolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate sensor update policy internal server error response has a 4xx status code
func (o *AggregateSensorUpdatePolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate sensor update policy internal server error response has a 5xx status code
func (o *AggregateSensorUpdatePolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this aggregate sensor update policy internal server error response a status code equal to that given
func (o *AggregateSensorUpdatePolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the aggregate sensor update policy internal server error response
func (o *AggregateSensorUpdatePolicyInternalServerError) Code() int {
	return 500
}

func (o *AggregateSensorUpdatePolicyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/sensor-update-policies/v1][%d] aggregateSensorUpdatePolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *AggregateSensorUpdatePolicyInternalServerError) String() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/sensor-update-policies/v1][%d] aggregateSensorUpdatePolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *AggregateSensorUpdatePolicyInternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *AggregateSensorUpdatePolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
