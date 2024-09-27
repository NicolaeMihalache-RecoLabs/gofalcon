// Code generated by go-swagger; DO NOT EDIT.

package firewall_management

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

// QueryPlatformsReader is a Reader for the QueryPlatforms structure.
type QueryPlatformsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryPlatformsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryPlatformsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQueryPlatformsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryPlatformsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /fwmgr/queries/platforms/v1] query-platforms", response, response.Code())
	}
}

// NewQueryPlatformsOK creates a QueryPlatformsOK with default headers values
func NewQueryPlatformsOK() *QueryPlatformsOK {
	return &QueryPlatformsOK{}
}

/*
QueryPlatformsOK describes a response with status code 200, with default header values.

OK
*/
type QueryPlatformsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FwmgrMsaspecQueryResponse
}

// IsSuccess returns true when this query platforms o k response has a 2xx status code
func (o *QueryPlatformsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query platforms o k response has a 3xx status code
func (o *QueryPlatformsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query platforms o k response has a 4xx status code
func (o *QueryPlatformsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query platforms o k response has a 5xx status code
func (o *QueryPlatformsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query platforms o k response a status code equal to that given
func (o *QueryPlatformsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query platforms o k response
func (o *QueryPlatformsOK) Code() int {
	return 200
}

func (o *QueryPlatformsOK) Error() string {
	return fmt.Sprintf("[GET /fwmgr/queries/platforms/v1][%d] queryPlatformsOK  %+v", 200, o.Payload)
}

func (o *QueryPlatformsOK) String() string {
	return fmt.Sprintf("[GET /fwmgr/queries/platforms/v1][%d] queryPlatformsOK  %+v", 200, o.Payload)
}

func (o *QueryPlatformsOK) GetPayload() *models.FwmgrMsaspecQueryResponse {
	return o.Payload
}

func (o *QueryPlatformsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FwmgrMsaspecQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryPlatformsForbidden creates a QueryPlatformsForbidden with default headers values
func NewQueryPlatformsForbidden() *QueryPlatformsForbidden {
	return &QueryPlatformsForbidden{}
}

/*
QueryPlatformsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryPlatformsForbidden struct {

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

// IsSuccess returns true when this query platforms forbidden response has a 2xx status code
func (o *QueryPlatformsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query platforms forbidden response has a 3xx status code
func (o *QueryPlatformsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query platforms forbidden response has a 4xx status code
func (o *QueryPlatformsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query platforms forbidden response has a 5xx status code
func (o *QueryPlatformsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query platforms forbidden response a status code equal to that given
func (o *QueryPlatformsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query platforms forbidden response
func (o *QueryPlatformsForbidden) Code() int {
	return 403
}

func (o *QueryPlatformsForbidden) Error() string {
	return fmt.Sprintf("[GET /fwmgr/queries/platforms/v1][%d] queryPlatformsForbidden  %+v", 403, o.Payload)
}

func (o *QueryPlatformsForbidden) String() string {
	return fmt.Sprintf("[GET /fwmgr/queries/platforms/v1][%d] queryPlatformsForbidden  %+v", 403, o.Payload)
}

func (o *QueryPlatformsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryPlatformsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryPlatformsTooManyRequests creates a QueryPlatformsTooManyRequests with default headers values
func NewQueryPlatformsTooManyRequests() *QueryPlatformsTooManyRequests {
	return &QueryPlatformsTooManyRequests{}
}

/*
QueryPlatformsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryPlatformsTooManyRequests struct {

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

// IsSuccess returns true when this query platforms too many requests response has a 2xx status code
func (o *QueryPlatformsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query platforms too many requests response has a 3xx status code
func (o *QueryPlatformsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query platforms too many requests response has a 4xx status code
func (o *QueryPlatformsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query platforms too many requests response has a 5xx status code
func (o *QueryPlatformsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query platforms too many requests response a status code equal to that given
func (o *QueryPlatformsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query platforms too many requests response
func (o *QueryPlatformsTooManyRequests) Code() int {
	return 429
}

func (o *QueryPlatformsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /fwmgr/queries/platforms/v1][%d] queryPlatformsTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryPlatformsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /fwmgr/queries/platforms/v1][%d] queryPlatformsTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryPlatformsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryPlatformsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
