// Code generated by go-swagger; DO NOT EDIT.

package overwatch_dashboard

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

// AggregatesOWEventsGlobalCountsReader is a Reader for the AggregatesOWEventsGlobalCounts structure.
type AggregatesOWEventsGlobalCountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregatesOWEventsGlobalCountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregatesOWEventsGlobalCountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAggregatesOWEventsGlobalCountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAggregatesOWEventsGlobalCountsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /overwatch-dashboards/aggregates/ow-events-global-counts/v1] AggregatesOWEventsGlobalCounts", response, response.Code())
	}
}

// NewAggregatesOWEventsGlobalCountsOK creates a AggregatesOWEventsGlobalCountsOK with default headers values
func NewAggregatesOWEventsGlobalCountsOK() *AggregatesOWEventsGlobalCountsOK {
	return &AggregatesOWEventsGlobalCountsOK{}
}

/*
AggregatesOWEventsGlobalCountsOK describes a response with status code 200, with default header values.

OK
*/
type AggregatesOWEventsGlobalCountsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaFacetsResponse
}

// IsSuccess returns true when this aggregates o w events global counts o k response has a 2xx status code
func (o *AggregatesOWEventsGlobalCountsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aggregates o w events global counts o k response has a 3xx status code
func (o *AggregatesOWEventsGlobalCountsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregates o w events global counts o k response has a 4xx status code
func (o *AggregatesOWEventsGlobalCountsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregates o w events global counts o k response has a 5xx status code
func (o *AggregatesOWEventsGlobalCountsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregates o w events global counts o k response a status code equal to that given
func (o *AggregatesOWEventsGlobalCountsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the aggregates o w events global counts o k response
func (o *AggregatesOWEventsGlobalCountsOK) Code() int {
	return 200
}

func (o *AggregatesOWEventsGlobalCountsOK) Error() string {
	return fmt.Sprintf("[GET /overwatch-dashboards/aggregates/ow-events-global-counts/v1][%d] aggregatesOWEventsGlobalCountsOK  %+v", 200, o.Payload)
}

func (o *AggregatesOWEventsGlobalCountsOK) String() string {
	return fmt.Sprintf("[GET /overwatch-dashboards/aggregates/ow-events-global-counts/v1][%d] aggregatesOWEventsGlobalCountsOK  %+v", 200, o.Payload)
}

func (o *AggregatesOWEventsGlobalCountsOK) GetPayload() *models.MsaFacetsResponse {
	return o.Payload
}

func (o *AggregatesOWEventsGlobalCountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaFacetsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregatesOWEventsGlobalCountsForbidden creates a AggregatesOWEventsGlobalCountsForbidden with default headers values
func NewAggregatesOWEventsGlobalCountsForbidden() *AggregatesOWEventsGlobalCountsForbidden {
	return &AggregatesOWEventsGlobalCountsForbidden{}
}

/*
AggregatesOWEventsGlobalCountsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AggregatesOWEventsGlobalCountsForbidden struct {

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

// IsSuccess returns true when this aggregates o w events global counts forbidden response has a 2xx status code
func (o *AggregatesOWEventsGlobalCountsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregates o w events global counts forbidden response has a 3xx status code
func (o *AggregatesOWEventsGlobalCountsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregates o w events global counts forbidden response has a 4xx status code
func (o *AggregatesOWEventsGlobalCountsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregates o w events global counts forbidden response has a 5xx status code
func (o *AggregatesOWEventsGlobalCountsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregates o w events global counts forbidden response a status code equal to that given
func (o *AggregatesOWEventsGlobalCountsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the aggregates o w events global counts forbidden response
func (o *AggregatesOWEventsGlobalCountsForbidden) Code() int {
	return 403
}

func (o *AggregatesOWEventsGlobalCountsForbidden) Error() string {
	return fmt.Sprintf("[GET /overwatch-dashboards/aggregates/ow-events-global-counts/v1][%d] aggregatesOWEventsGlobalCountsForbidden  %+v", 403, o.Payload)
}

func (o *AggregatesOWEventsGlobalCountsForbidden) String() string {
	return fmt.Sprintf("[GET /overwatch-dashboards/aggregates/ow-events-global-counts/v1][%d] aggregatesOWEventsGlobalCountsForbidden  %+v", 403, o.Payload)
}

func (o *AggregatesOWEventsGlobalCountsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregatesOWEventsGlobalCountsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregatesOWEventsGlobalCountsTooManyRequests creates a AggregatesOWEventsGlobalCountsTooManyRequests with default headers values
func NewAggregatesOWEventsGlobalCountsTooManyRequests() *AggregatesOWEventsGlobalCountsTooManyRequests {
	return &AggregatesOWEventsGlobalCountsTooManyRequests{}
}

/*
AggregatesOWEventsGlobalCountsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AggregatesOWEventsGlobalCountsTooManyRequests struct {

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

// IsSuccess returns true when this aggregates o w events global counts too many requests response has a 2xx status code
func (o *AggregatesOWEventsGlobalCountsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregates o w events global counts too many requests response has a 3xx status code
func (o *AggregatesOWEventsGlobalCountsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregates o w events global counts too many requests response has a 4xx status code
func (o *AggregatesOWEventsGlobalCountsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregates o w events global counts too many requests response has a 5xx status code
func (o *AggregatesOWEventsGlobalCountsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregates o w events global counts too many requests response a status code equal to that given
func (o *AggregatesOWEventsGlobalCountsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the aggregates o w events global counts too many requests response
func (o *AggregatesOWEventsGlobalCountsTooManyRequests) Code() int {
	return 429
}

func (o *AggregatesOWEventsGlobalCountsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /overwatch-dashboards/aggregates/ow-events-global-counts/v1][%d] aggregatesOWEventsGlobalCountsTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregatesOWEventsGlobalCountsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /overwatch-dashboards/aggregates/ow-events-global-counts/v1][%d] aggregatesOWEventsGlobalCountsTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregatesOWEventsGlobalCountsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregatesOWEventsGlobalCountsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
