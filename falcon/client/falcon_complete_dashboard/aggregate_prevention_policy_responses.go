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

// AggregatePreventionPolicyReader is a Reader for the AggregatePreventionPolicy structure.
type AggregatePreventionPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregatePreventionPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregatePreventionPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAggregatePreventionPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAggregatePreventionPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAggregatePreventionPolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAggregatePreventionPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /falcon-complete-dashboards/aggregates/prevention-policies/v1] AggregatePreventionPolicy", response, response.Code())
	}
}

// NewAggregatePreventionPolicyOK creates a AggregatePreventionPolicyOK with default headers values
func NewAggregatePreventionPolicyOK() *AggregatePreventionPolicyOK {
	return &AggregatePreventionPolicyOK{}
}

/*
AggregatePreventionPolicyOK describes a response with status code 200, with default header values.

OK
*/
type AggregatePreventionPolicyOK struct {

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

// IsSuccess returns true when this aggregate prevention policy o k response has a 2xx status code
func (o *AggregatePreventionPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aggregate prevention policy o k response has a 3xx status code
func (o *AggregatePreventionPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate prevention policy o k response has a 4xx status code
func (o *AggregatePreventionPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate prevention policy o k response has a 5xx status code
func (o *AggregatePreventionPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate prevention policy o k response a status code equal to that given
func (o *AggregatePreventionPolicyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the aggregate prevention policy o k response
func (o *AggregatePreventionPolicyOK) Code() int {
	return 200
}

func (o *AggregatePreventionPolicyOK) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/prevention-policies/v1][%d] aggregatePreventionPolicyOK  %+v", 200, o.Payload)
}

func (o *AggregatePreventionPolicyOK) String() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/prevention-policies/v1][%d] aggregatePreventionPolicyOK  %+v", 200, o.Payload)
}

func (o *AggregatePreventionPolicyOK) GetPayload() *models.MsaAggregatesResponse {
	return o.Payload
}

func (o *AggregatePreventionPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregatePreventionPolicyBadRequest creates a AggregatePreventionPolicyBadRequest with default headers values
func NewAggregatePreventionPolicyBadRequest() *AggregatePreventionPolicyBadRequest {
	return &AggregatePreventionPolicyBadRequest{}
}

/*
AggregatePreventionPolicyBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AggregatePreventionPolicyBadRequest struct {

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

// IsSuccess returns true when this aggregate prevention policy bad request response has a 2xx status code
func (o *AggregatePreventionPolicyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate prevention policy bad request response has a 3xx status code
func (o *AggregatePreventionPolicyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate prevention policy bad request response has a 4xx status code
func (o *AggregatePreventionPolicyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate prevention policy bad request response has a 5xx status code
func (o *AggregatePreventionPolicyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate prevention policy bad request response a status code equal to that given
func (o *AggregatePreventionPolicyBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the aggregate prevention policy bad request response
func (o *AggregatePreventionPolicyBadRequest) Code() int {
	return 400
}

func (o *AggregatePreventionPolicyBadRequest) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/prevention-policies/v1][%d] aggregatePreventionPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *AggregatePreventionPolicyBadRequest) String() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/prevention-policies/v1][%d] aggregatePreventionPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *AggregatePreventionPolicyBadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *AggregatePreventionPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregatePreventionPolicyForbidden creates a AggregatePreventionPolicyForbidden with default headers values
func NewAggregatePreventionPolicyForbidden() *AggregatePreventionPolicyForbidden {
	return &AggregatePreventionPolicyForbidden{}
}

/*
AggregatePreventionPolicyForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AggregatePreventionPolicyForbidden struct {

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

// IsSuccess returns true when this aggregate prevention policy forbidden response has a 2xx status code
func (o *AggregatePreventionPolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate prevention policy forbidden response has a 3xx status code
func (o *AggregatePreventionPolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate prevention policy forbidden response has a 4xx status code
func (o *AggregatePreventionPolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate prevention policy forbidden response has a 5xx status code
func (o *AggregatePreventionPolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate prevention policy forbidden response a status code equal to that given
func (o *AggregatePreventionPolicyForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the aggregate prevention policy forbidden response
func (o *AggregatePreventionPolicyForbidden) Code() int {
	return 403
}

func (o *AggregatePreventionPolicyForbidden) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/prevention-policies/v1][%d] aggregatePreventionPolicyForbidden  %+v", 403, o.Payload)
}

func (o *AggregatePreventionPolicyForbidden) String() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/prevention-policies/v1][%d] aggregatePreventionPolicyForbidden  %+v", 403, o.Payload)
}

func (o *AggregatePreventionPolicyForbidden) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *AggregatePreventionPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregatePreventionPolicyTooManyRequests creates a AggregatePreventionPolicyTooManyRequests with default headers values
func NewAggregatePreventionPolicyTooManyRequests() *AggregatePreventionPolicyTooManyRequests {
	return &AggregatePreventionPolicyTooManyRequests{}
}

/*
AggregatePreventionPolicyTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AggregatePreventionPolicyTooManyRequests struct {

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

// IsSuccess returns true when this aggregate prevention policy too many requests response has a 2xx status code
func (o *AggregatePreventionPolicyTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate prevention policy too many requests response has a 3xx status code
func (o *AggregatePreventionPolicyTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate prevention policy too many requests response has a 4xx status code
func (o *AggregatePreventionPolicyTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate prevention policy too many requests response has a 5xx status code
func (o *AggregatePreventionPolicyTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate prevention policy too many requests response a status code equal to that given
func (o *AggregatePreventionPolicyTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the aggregate prevention policy too many requests response
func (o *AggregatePreventionPolicyTooManyRequests) Code() int {
	return 429
}

func (o *AggregatePreventionPolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/prevention-policies/v1][%d] aggregatePreventionPolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregatePreventionPolicyTooManyRequests) String() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/prevention-policies/v1][%d] aggregatePreventionPolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregatePreventionPolicyTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregatePreventionPolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregatePreventionPolicyInternalServerError creates a AggregatePreventionPolicyInternalServerError with default headers values
func NewAggregatePreventionPolicyInternalServerError() *AggregatePreventionPolicyInternalServerError {
	return &AggregatePreventionPolicyInternalServerError{}
}

/*
AggregatePreventionPolicyInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type AggregatePreventionPolicyInternalServerError struct {

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

// IsSuccess returns true when this aggregate prevention policy internal server error response has a 2xx status code
func (o *AggregatePreventionPolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate prevention policy internal server error response has a 3xx status code
func (o *AggregatePreventionPolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate prevention policy internal server error response has a 4xx status code
func (o *AggregatePreventionPolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate prevention policy internal server error response has a 5xx status code
func (o *AggregatePreventionPolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this aggregate prevention policy internal server error response a status code equal to that given
func (o *AggregatePreventionPolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the aggregate prevention policy internal server error response
func (o *AggregatePreventionPolicyInternalServerError) Code() int {
	return 500
}

func (o *AggregatePreventionPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/prevention-policies/v1][%d] aggregatePreventionPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *AggregatePreventionPolicyInternalServerError) String() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/prevention-policies/v1][%d] aggregatePreventionPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *AggregatePreventionPolicyInternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *AggregatePreventionPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
