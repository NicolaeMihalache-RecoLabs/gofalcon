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

// AggregateEscalationsReader is a Reader for the AggregateEscalations structure.
type AggregateEscalationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregateEscalationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregateEscalationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAggregateEscalationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAggregateEscalationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /falcon-complete-dashboards/aggregates/escalations/GET/v1] AggregateEscalations", response, response.Code())
	}
}

// NewAggregateEscalationsOK creates a AggregateEscalationsOK with default headers values
func NewAggregateEscalationsOK() *AggregateEscalationsOK {
	return &AggregateEscalationsOK{}
}

/*
AggregateEscalationsOK describes a response with status code 200, with default header values.

OK
*/
type AggregateEscalationsOK struct {

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

// IsSuccess returns true when this aggregate escalations o k response has a 2xx status code
func (o *AggregateEscalationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aggregate escalations o k response has a 3xx status code
func (o *AggregateEscalationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate escalations o k response has a 4xx status code
func (o *AggregateEscalationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate escalations o k response has a 5xx status code
func (o *AggregateEscalationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate escalations o k response a status code equal to that given
func (o *AggregateEscalationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the aggregate escalations o k response
func (o *AggregateEscalationsOK) Code() int {
	return 200
}

func (o *AggregateEscalationsOK) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/escalations/GET/v1][%d] aggregateEscalationsOK  %+v", 200, o.Payload)
}

func (o *AggregateEscalationsOK) String() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/escalations/GET/v1][%d] aggregateEscalationsOK  %+v", 200, o.Payload)
}

func (o *AggregateEscalationsOK) GetPayload() *models.MsaAggregatesResponse {
	return o.Payload
}

func (o *AggregateEscalationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregateEscalationsForbidden creates a AggregateEscalationsForbidden with default headers values
func NewAggregateEscalationsForbidden() *AggregateEscalationsForbidden {
	return &AggregateEscalationsForbidden{}
}

/*
AggregateEscalationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AggregateEscalationsForbidden struct {

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

// IsSuccess returns true when this aggregate escalations forbidden response has a 2xx status code
func (o *AggregateEscalationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate escalations forbidden response has a 3xx status code
func (o *AggregateEscalationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate escalations forbidden response has a 4xx status code
func (o *AggregateEscalationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate escalations forbidden response has a 5xx status code
func (o *AggregateEscalationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate escalations forbidden response a status code equal to that given
func (o *AggregateEscalationsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the aggregate escalations forbidden response
func (o *AggregateEscalationsForbidden) Code() int {
	return 403
}

func (o *AggregateEscalationsForbidden) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/escalations/GET/v1][%d] aggregateEscalationsForbidden  %+v", 403, o.Payload)
}

func (o *AggregateEscalationsForbidden) String() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/escalations/GET/v1][%d] aggregateEscalationsForbidden  %+v", 403, o.Payload)
}

func (o *AggregateEscalationsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregateEscalationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregateEscalationsTooManyRequests creates a AggregateEscalationsTooManyRequests with default headers values
func NewAggregateEscalationsTooManyRequests() *AggregateEscalationsTooManyRequests {
	return &AggregateEscalationsTooManyRequests{}
}

/*
AggregateEscalationsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AggregateEscalationsTooManyRequests struct {

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

// IsSuccess returns true when this aggregate escalations too many requests response has a 2xx status code
func (o *AggregateEscalationsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate escalations too many requests response has a 3xx status code
func (o *AggregateEscalationsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate escalations too many requests response has a 4xx status code
func (o *AggregateEscalationsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate escalations too many requests response has a 5xx status code
func (o *AggregateEscalationsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate escalations too many requests response a status code equal to that given
func (o *AggregateEscalationsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the aggregate escalations too many requests response
func (o *AggregateEscalationsTooManyRequests) Code() int {
	return 429
}

func (o *AggregateEscalationsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/escalations/GET/v1][%d] aggregateEscalationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregateEscalationsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/escalations/GET/v1][%d] aggregateEscalationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregateEscalationsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregateEscalationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
