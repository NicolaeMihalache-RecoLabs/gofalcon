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

// NodeEnrichmentReader is a Reader for the NodeEnrichment structure.
type NodeEnrichmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodeEnrichmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodeEnrichmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewNodeEnrichmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewNodeEnrichmentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNodeEnrichmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/aggregates/enrichment/nodes/entities/v1] NodeEnrichment", response, response.Code())
	}
}

// NewNodeEnrichmentOK creates a NodeEnrichmentOK with default headers values
func NewNodeEnrichmentOK() *NodeEnrichmentOK {
	return &NodeEnrichmentOK{}
}

/*
NodeEnrichmentOK describes a response with status code 200, with default header values.

OK
*/
type NodeEnrichmentOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.K8sassetsNodeEnrichmentResponse
}

// IsSuccess returns true when this node enrichment o k response has a 2xx status code
func (o *NodeEnrichmentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this node enrichment o k response has a 3xx status code
func (o *NodeEnrichmentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node enrichment o k response has a 4xx status code
func (o *NodeEnrichmentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this node enrichment o k response has a 5xx status code
func (o *NodeEnrichmentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this node enrichment o k response a status code equal to that given
func (o *NodeEnrichmentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the node enrichment o k response
func (o *NodeEnrichmentOK) Code() int {
	return 200
}

func (o *NodeEnrichmentOK) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/enrichment/nodes/entities/v1][%d] nodeEnrichmentOK  %+v", 200, o.Payload)
}

func (o *NodeEnrichmentOK) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/enrichment/nodes/entities/v1][%d] nodeEnrichmentOK  %+v", 200, o.Payload)
}

func (o *NodeEnrichmentOK) GetPayload() *models.K8sassetsNodeEnrichmentResponse {
	return o.Payload
}

func (o *NodeEnrichmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.K8sassetsNodeEnrichmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeEnrichmentForbidden creates a NodeEnrichmentForbidden with default headers values
func NewNodeEnrichmentForbidden() *NodeEnrichmentForbidden {
	return &NodeEnrichmentForbidden{}
}

/*
NodeEnrichmentForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type NodeEnrichmentForbidden struct {

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

// IsSuccess returns true when this node enrichment forbidden response has a 2xx status code
func (o *NodeEnrichmentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this node enrichment forbidden response has a 3xx status code
func (o *NodeEnrichmentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node enrichment forbidden response has a 4xx status code
func (o *NodeEnrichmentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this node enrichment forbidden response has a 5xx status code
func (o *NodeEnrichmentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this node enrichment forbidden response a status code equal to that given
func (o *NodeEnrichmentForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the node enrichment forbidden response
func (o *NodeEnrichmentForbidden) Code() int {
	return 403
}

func (o *NodeEnrichmentForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/enrichment/nodes/entities/v1][%d] nodeEnrichmentForbidden  %+v", 403, o.Payload)
}

func (o *NodeEnrichmentForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/enrichment/nodes/entities/v1][%d] nodeEnrichmentForbidden  %+v", 403, o.Payload)
}

func (o *NodeEnrichmentForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *NodeEnrichmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewNodeEnrichmentTooManyRequests creates a NodeEnrichmentTooManyRequests with default headers values
func NewNodeEnrichmentTooManyRequests() *NodeEnrichmentTooManyRequests {
	return &NodeEnrichmentTooManyRequests{}
}

/*
NodeEnrichmentTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type NodeEnrichmentTooManyRequests struct {

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

// IsSuccess returns true when this node enrichment too many requests response has a 2xx status code
func (o *NodeEnrichmentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this node enrichment too many requests response has a 3xx status code
func (o *NodeEnrichmentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node enrichment too many requests response has a 4xx status code
func (o *NodeEnrichmentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this node enrichment too many requests response has a 5xx status code
func (o *NodeEnrichmentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this node enrichment too many requests response a status code equal to that given
func (o *NodeEnrichmentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the node enrichment too many requests response
func (o *NodeEnrichmentTooManyRequests) Code() int {
	return 429
}

func (o *NodeEnrichmentTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/enrichment/nodes/entities/v1][%d] nodeEnrichmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *NodeEnrichmentTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/enrichment/nodes/entities/v1][%d] nodeEnrichmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *NodeEnrichmentTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *NodeEnrichmentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewNodeEnrichmentInternalServerError creates a NodeEnrichmentInternalServerError with default headers values
func NewNodeEnrichmentInternalServerError() *NodeEnrichmentInternalServerError {
	return &NodeEnrichmentInternalServerError{}
}

/*
NodeEnrichmentInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type NodeEnrichmentInternalServerError struct {

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

// IsSuccess returns true when this node enrichment internal server error response has a 2xx status code
func (o *NodeEnrichmentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this node enrichment internal server error response has a 3xx status code
func (o *NodeEnrichmentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node enrichment internal server error response has a 4xx status code
func (o *NodeEnrichmentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this node enrichment internal server error response has a 5xx status code
func (o *NodeEnrichmentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this node enrichment internal server error response a status code equal to that given
func (o *NodeEnrichmentInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the node enrichment internal server error response
func (o *NodeEnrichmentInternalServerError) Code() int {
	return 500
}

func (o *NodeEnrichmentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/enrichment/nodes/entities/v1][%d] nodeEnrichmentInternalServerError  %+v", 500, o.Payload)
}

func (o *NodeEnrichmentInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/enrichment/nodes/entities/v1][%d] nodeEnrichmentInternalServerError  %+v", 500, o.Payload)
}

func (o *NodeEnrichmentInternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *NodeEnrichmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
