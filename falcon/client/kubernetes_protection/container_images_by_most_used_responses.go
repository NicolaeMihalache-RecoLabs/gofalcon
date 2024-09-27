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

// ContainerImagesByMostUsedReader is a Reader for the ContainerImagesByMostUsed structure.
type ContainerImagesByMostUsedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerImagesByMostUsedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerImagesByMostUsedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewContainerImagesByMostUsedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewContainerImagesByMostUsedTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerImagesByMostUsedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/aggregates/images/most-used/v1] ContainerImagesByMostUsed", response, response.Code())
	}
}

// NewContainerImagesByMostUsedOK creates a ContainerImagesByMostUsedOK with default headers values
func NewContainerImagesByMostUsedOK() *ContainerImagesByMostUsedOK {
	return &ContainerImagesByMostUsedOK{}
}

/*
ContainerImagesByMostUsedOK describes a response with status code 200, with default header values.

OK
*/
type ContainerImagesByMostUsedOK struct {

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

// IsSuccess returns true when this container images by most used o k response has a 2xx status code
func (o *ContainerImagesByMostUsedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this container images by most used o k response has a 3xx status code
func (o *ContainerImagesByMostUsedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container images by most used o k response has a 4xx status code
func (o *ContainerImagesByMostUsedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this container images by most used o k response has a 5xx status code
func (o *ContainerImagesByMostUsedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this container images by most used o k response a status code equal to that given
func (o *ContainerImagesByMostUsedOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the container images by most used o k response
func (o *ContainerImagesByMostUsedOK) Code() int {
	return 200
}

func (o *ContainerImagesByMostUsedOK) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/images/most-used/v1][%d] containerImagesByMostUsedOK  %+v", 200, o.Payload)
}

func (o *ContainerImagesByMostUsedOK) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/images/most-used/v1][%d] containerImagesByMostUsedOK  %+v", 200, o.Payload)
}

func (o *ContainerImagesByMostUsedOK) GetPayload() *models.ModelsAggregateValuesByFieldResponse {
	return o.Payload
}

func (o *ContainerImagesByMostUsedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewContainerImagesByMostUsedForbidden creates a ContainerImagesByMostUsedForbidden with default headers values
func NewContainerImagesByMostUsedForbidden() *ContainerImagesByMostUsedForbidden {
	return &ContainerImagesByMostUsedForbidden{}
}

/*
ContainerImagesByMostUsedForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ContainerImagesByMostUsedForbidden struct {

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

// IsSuccess returns true when this container images by most used forbidden response has a 2xx status code
func (o *ContainerImagesByMostUsedForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container images by most used forbidden response has a 3xx status code
func (o *ContainerImagesByMostUsedForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container images by most used forbidden response has a 4xx status code
func (o *ContainerImagesByMostUsedForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this container images by most used forbidden response has a 5xx status code
func (o *ContainerImagesByMostUsedForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this container images by most used forbidden response a status code equal to that given
func (o *ContainerImagesByMostUsedForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the container images by most used forbidden response
func (o *ContainerImagesByMostUsedForbidden) Code() int {
	return 403
}

func (o *ContainerImagesByMostUsedForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/images/most-used/v1][%d] containerImagesByMostUsedForbidden  %+v", 403, o.Payload)
}

func (o *ContainerImagesByMostUsedForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/images/most-used/v1][%d] containerImagesByMostUsedForbidden  %+v", 403, o.Payload)
}

func (o *ContainerImagesByMostUsedForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ContainerImagesByMostUsedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewContainerImagesByMostUsedTooManyRequests creates a ContainerImagesByMostUsedTooManyRequests with default headers values
func NewContainerImagesByMostUsedTooManyRequests() *ContainerImagesByMostUsedTooManyRequests {
	return &ContainerImagesByMostUsedTooManyRequests{}
}

/*
ContainerImagesByMostUsedTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ContainerImagesByMostUsedTooManyRequests struct {

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

// IsSuccess returns true when this container images by most used too many requests response has a 2xx status code
func (o *ContainerImagesByMostUsedTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container images by most used too many requests response has a 3xx status code
func (o *ContainerImagesByMostUsedTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container images by most used too many requests response has a 4xx status code
func (o *ContainerImagesByMostUsedTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this container images by most used too many requests response has a 5xx status code
func (o *ContainerImagesByMostUsedTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this container images by most used too many requests response a status code equal to that given
func (o *ContainerImagesByMostUsedTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the container images by most used too many requests response
func (o *ContainerImagesByMostUsedTooManyRequests) Code() int {
	return 429
}

func (o *ContainerImagesByMostUsedTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/images/most-used/v1][%d] containerImagesByMostUsedTooManyRequests  %+v", 429, o.Payload)
}

func (o *ContainerImagesByMostUsedTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/images/most-used/v1][%d] containerImagesByMostUsedTooManyRequests  %+v", 429, o.Payload)
}

func (o *ContainerImagesByMostUsedTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ContainerImagesByMostUsedTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewContainerImagesByMostUsedInternalServerError creates a ContainerImagesByMostUsedInternalServerError with default headers values
func NewContainerImagesByMostUsedInternalServerError() *ContainerImagesByMostUsedInternalServerError {
	return &ContainerImagesByMostUsedInternalServerError{}
}

/*
ContainerImagesByMostUsedInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ContainerImagesByMostUsedInternalServerError struct {

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

// IsSuccess returns true when this container images by most used internal server error response has a 2xx status code
func (o *ContainerImagesByMostUsedInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container images by most used internal server error response has a 3xx status code
func (o *ContainerImagesByMostUsedInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container images by most used internal server error response has a 4xx status code
func (o *ContainerImagesByMostUsedInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this container images by most used internal server error response has a 5xx status code
func (o *ContainerImagesByMostUsedInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this container images by most used internal server error response a status code equal to that given
func (o *ContainerImagesByMostUsedInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the container images by most used internal server error response
func (o *ContainerImagesByMostUsedInternalServerError) Code() int {
	return 500
}

func (o *ContainerImagesByMostUsedInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/images/most-used/v1][%d] containerImagesByMostUsedInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerImagesByMostUsedInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/images/most-used/v1][%d] containerImagesByMostUsedInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerImagesByMostUsedInternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *ContainerImagesByMostUsedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
