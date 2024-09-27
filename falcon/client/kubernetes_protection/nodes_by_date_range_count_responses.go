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

// NodesByDateRangeCountReader is a Reader for the NodesByDateRangeCount structure.
type NodesByDateRangeCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodesByDateRangeCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodesByDateRangeCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewNodesByDateRangeCountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewNodesByDateRangeCountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNodesByDateRangeCountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/aggregates/nodes/count-by-date/v1] NodesByDateRangeCount", response, response.Code())
	}
}

// NewNodesByDateRangeCountOK creates a NodesByDateRangeCountOK with default headers values
func NewNodesByDateRangeCountOK() *NodesByDateRangeCountOK {
	return &NodesByDateRangeCountOK{}
}

/*
NodesByDateRangeCountOK describes a response with status code 200, with default header values.

OK
*/
type NodesByDateRangeCountOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ModelsAggregateValuesByFieldResponse
}

// IsSuccess returns true when this nodes by date range count o k response has a 2xx status code
func (o *NodesByDateRangeCountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nodes by date range count o k response has a 3xx status code
func (o *NodesByDateRangeCountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nodes by date range count o k response has a 4xx status code
func (o *NodesByDateRangeCountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nodes by date range count o k response has a 5xx status code
func (o *NodesByDateRangeCountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nodes by date range count o k response a status code equal to that given
func (o *NodesByDateRangeCountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nodes by date range count o k response
func (o *NodesByDateRangeCountOK) Code() int {
	return 200
}

func (o *NodesByDateRangeCountOK) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/nodes/count-by-date/v1][%d] nodesByDateRangeCountOK  %+v", 200, o.Payload)
}

func (o *NodesByDateRangeCountOK) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/nodes/count-by-date/v1][%d] nodesByDateRangeCountOK  %+v", 200, o.Payload)
}

func (o *NodesByDateRangeCountOK) GetPayload() *models.ModelsAggregateValuesByFieldResponse {
	return o.Payload
}

func (o *NodesByDateRangeCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ModelsAggregateValuesByFieldResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodesByDateRangeCountForbidden creates a NodesByDateRangeCountForbidden with default headers values
func NewNodesByDateRangeCountForbidden() *NodesByDateRangeCountForbidden {
	return &NodesByDateRangeCountForbidden{}
}

/*
NodesByDateRangeCountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type NodesByDateRangeCountForbidden struct {

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

// IsSuccess returns true when this nodes by date range count forbidden response has a 2xx status code
func (o *NodesByDateRangeCountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this nodes by date range count forbidden response has a 3xx status code
func (o *NodesByDateRangeCountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nodes by date range count forbidden response has a 4xx status code
func (o *NodesByDateRangeCountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this nodes by date range count forbidden response has a 5xx status code
func (o *NodesByDateRangeCountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this nodes by date range count forbidden response a status code equal to that given
func (o *NodesByDateRangeCountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the nodes by date range count forbidden response
func (o *NodesByDateRangeCountForbidden) Code() int {
	return 403
}

func (o *NodesByDateRangeCountForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/nodes/count-by-date/v1][%d] nodesByDateRangeCountForbidden  %+v", 403, o.Payload)
}

func (o *NodesByDateRangeCountForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/nodes/count-by-date/v1][%d] nodesByDateRangeCountForbidden  %+v", 403, o.Payload)
}

func (o *NodesByDateRangeCountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *NodesByDateRangeCountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewNodesByDateRangeCountTooManyRequests creates a NodesByDateRangeCountTooManyRequests with default headers values
func NewNodesByDateRangeCountTooManyRequests() *NodesByDateRangeCountTooManyRequests {
	return &NodesByDateRangeCountTooManyRequests{}
}

/*
NodesByDateRangeCountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type NodesByDateRangeCountTooManyRequests struct {

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

// IsSuccess returns true when this nodes by date range count too many requests response has a 2xx status code
func (o *NodesByDateRangeCountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this nodes by date range count too many requests response has a 3xx status code
func (o *NodesByDateRangeCountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nodes by date range count too many requests response has a 4xx status code
func (o *NodesByDateRangeCountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this nodes by date range count too many requests response has a 5xx status code
func (o *NodesByDateRangeCountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this nodes by date range count too many requests response a status code equal to that given
func (o *NodesByDateRangeCountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the nodes by date range count too many requests response
func (o *NodesByDateRangeCountTooManyRequests) Code() int {
	return 429
}

func (o *NodesByDateRangeCountTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/nodes/count-by-date/v1][%d] nodesByDateRangeCountTooManyRequests  %+v", 429, o.Payload)
}

func (o *NodesByDateRangeCountTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/nodes/count-by-date/v1][%d] nodesByDateRangeCountTooManyRequests  %+v", 429, o.Payload)
}

func (o *NodesByDateRangeCountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *NodesByDateRangeCountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewNodesByDateRangeCountInternalServerError creates a NodesByDateRangeCountInternalServerError with default headers values
func NewNodesByDateRangeCountInternalServerError() *NodesByDateRangeCountInternalServerError {
	return &NodesByDateRangeCountInternalServerError{}
}

/*
NodesByDateRangeCountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type NodesByDateRangeCountInternalServerError struct {

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

// IsSuccess returns true when this nodes by date range count internal server error response has a 2xx status code
func (o *NodesByDateRangeCountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this nodes by date range count internal server error response has a 3xx status code
func (o *NodesByDateRangeCountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nodes by date range count internal server error response has a 4xx status code
func (o *NodesByDateRangeCountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this nodes by date range count internal server error response has a 5xx status code
func (o *NodesByDateRangeCountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this nodes by date range count internal server error response a status code equal to that given
func (o *NodesByDateRangeCountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the nodes by date range count internal server error response
func (o *NodesByDateRangeCountInternalServerError) Code() int {
	return 500
}

func (o *NodesByDateRangeCountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/nodes/count-by-date/v1][%d] nodesByDateRangeCountInternalServerError  %+v", 500, o.Payload)
}

func (o *NodesByDateRangeCountInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/nodes/count-by-date/v1][%d] nodesByDateRangeCountInternalServerError  %+v", 500, o.Payload)
}

func (o *NodesByDateRangeCountInternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *NodesByDateRangeCountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
