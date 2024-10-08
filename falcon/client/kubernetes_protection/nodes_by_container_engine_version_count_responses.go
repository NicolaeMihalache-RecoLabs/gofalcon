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

// NodesByContainerEngineVersionCountReader is a Reader for the NodesByContainerEngineVersionCount structure.
type NodesByContainerEngineVersionCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodesByContainerEngineVersionCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodesByContainerEngineVersionCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewNodesByContainerEngineVersionCountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewNodesByContainerEngineVersionCountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNodesByContainerEngineVersionCountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/aggregates/nodes/count-by-container-engine-version/v1] NodesByContainerEngineVersionCount", response, response.Code())
	}
}

// NewNodesByContainerEngineVersionCountOK creates a NodesByContainerEngineVersionCountOK with default headers values
func NewNodesByContainerEngineVersionCountOK() *NodesByContainerEngineVersionCountOK {
	return &NodesByContainerEngineVersionCountOK{}
}

/*
NodesByContainerEngineVersionCountOK describes a response with status code 200, with default header values.

OK
*/
type NodesByContainerEngineVersionCountOK struct {

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

// IsSuccess returns true when this nodes by container engine version count o k response has a 2xx status code
func (o *NodesByContainerEngineVersionCountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nodes by container engine version count o k response has a 3xx status code
func (o *NodesByContainerEngineVersionCountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nodes by container engine version count o k response has a 4xx status code
func (o *NodesByContainerEngineVersionCountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nodes by container engine version count o k response has a 5xx status code
func (o *NodesByContainerEngineVersionCountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nodes by container engine version count o k response a status code equal to that given
func (o *NodesByContainerEngineVersionCountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nodes by container engine version count o k response
func (o *NodesByContainerEngineVersionCountOK) Code() int {
	return 200
}

func (o *NodesByContainerEngineVersionCountOK) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/nodes/count-by-container-engine-version/v1][%d] nodesByContainerEngineVersionCountOK  %+v", 200, o.Payload)
}

func (o *NodesByContainerEngineVersionCountOK) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/nodes/count-by-container-engine-version/v1][%d] nodesByContainerEngineVersionCountOK  %+v", 200, o.Payload)
}

func (o *NodesByContainerEngineVersionCountOK) GetPayload() *models.ModelsAggregateValuesByFieldResponse {
	return o.Payload
}

func (o *NodesByContainerEngineVersionCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewNodesByContainerEngineVersionCountForbidden creates a NodesByContainerEngineVersionCountForbidden with default headers values
func NewNodesByContainerEngineVersionCountForbidden() *NodesByContainerEngineVersionCountForbidden {
	return &NodesByContainerEngineVersionCountForbidden{}
}

/*
NodesByContainerEngineVersionCountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type NodesByContainerEngineVersionCountForbidden struct {

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

// IsSuccess returns true when this nodes by container engine version count forbidden response has a 2xx status code
func (o *NodesByContainerEngineVersionCountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this nodes by container engine version count forbidden response has a 3xx status code
func (o *NodesByContainerEngineVersionCountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nodes by container engine version count forbidden response has a 4xx status code
func (o *NodesByContainerEngineVersionCountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this nodes by container engine version count forbidden response has a 5xx status code
func (o *NodesByContainerEngineVersionCountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this nodes by container engine version count forbidden response a status code equal to that given
func (o *NodesByContainerEngineVersionCountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the nodes by container engine version count forbidden response
func (o *NodesByContainerEngineVersionCountForbidden) Code() int {
	return 403
}

func (o *NodesByContainerEngineVersionCountForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/nodes/count-by-container-engine-version/v1][%d] nodesByContainerEngineVersionCountForbidden  %+v", 403, o.Payload)
}

func (o *NodesByContainerEngineVersionCountForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/nodes/count-by-container-engine-version/v1][%d] nodesByContainerEngineVersionCountForbidden  %+v", 403, o.Payload)
}

func (o *NodesByContainerEngineVersionCountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *NodesByContainerEngineVersionCountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewNodesByContainerEngineVersionCountTooManyRequests creates a NodesByContainerEngineVersionCountTooManyRequests with default headers values
func NewNodesByContainerEngineVersionCountTooManyRequests() *NodesByContainerEngineVersionCountTooManyRequests {
	return &NodesByContainerEngineVersionCountTooManyRequests{}
}

/*
NodesByContainerEngineVersionCountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type NodesByContainerEngineVersionCountTooManyRequests struct {

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

// IsSuccess returns true when this nodes by container engine version count too many requests response has a 2xx status code
func (o *NodesByContainerEngineVersionCountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this nodes by container engine version count too many requests response has a 3xx status code
func (o *NodesByContainerEngineVersionCountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nodes by container engine version count too many requests response has a 4xx status code
func (o *NodesByContainerEngineVersionCountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this nodes by container engine version count too many requests response has a 5xx status code
func (o *NodesByContainerEngineVersionCountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this nodes by container engine version count too many requests response a status code equal to that given
func (o *NodesByContainerEngineVersionCountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the nodes by container engine version count too many requests response
func (o *NodesByContainerEngineVersionCountTooManyRequests) Code() int {
	return 429
}

func (o *NodesByContainerEngineVersionCountTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/nodes/count-by-container-engine-version/v1][%d] nodesByContainerEngineVersionCountTooManyRequests  %+v", 429, o.Payload)
}

func (o *NodesByContainerEngineVersionCountTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/nodes/count-by-container-engine-version/v1][%d] nodesByContainerEngineVersionCountTooManyRequests  %+v", 429, o.Payload)
}

func (o *NodesByContainerEngineVersionCountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *NodesByContainerEngineVersionCountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewNodesByContainerEngineVersionCountInternalServerError creates a NodesByContainerEngineVersionCountInternalServerError with default headers values
func NewNodesByContainerEngineVersionCountInternalServerError() *NodesByContainerEngineVersionCountInternalServerError {
	return &NodesByContainerEngineVersionCountInternalServerError{}
}

/*
NodesByContainerEngineVersionCountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type NodesByContainerEngineVersionCountInternalServerError struct {

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

// IsSuccess returns true when this nodes by container engine version count internal server error response has a 2xx status code
func (o *NodesByContainerEngineVersionCountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this nodes by container engine version count internal server error response has a 3xx status code
func (o *NodesByContainerEngineVersionCountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nodes by container engine version count internal server error response has a 4xx status code
func (o *NodesByContainerEngineVersionCountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this nodes by container engine version count internal server error response has a 5xx status code
func (o *NodesByContainerEngineVersionCountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this nodes by container engine version count internal server error response a status code equal to that given
func (o *NodesByContainerEngineVersionCountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the nodes by container engine version count internal server error response
func (o *NodesByContainerEngineVersionCountInternalServerError) Code() int {
	return 500
}

func (o *NodesByContainerEngineVersionCountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/nodes/count-by-container-engine-version/v1][%d] nodesByContainerEngineVersionCountInternalServerError  %+v", 500, o.Payload)
}

func (o *NodesByContainerEngineVersionCountInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/nodes/count-by-container-engine-version/v1][%d] nodesByContainerEngineVersionCountInternalServerError  %+v", 500, o.Payload)
}

func (o *NodesByContainerEngineVersionCountInternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *NodesByContainerEngineVersionCountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
