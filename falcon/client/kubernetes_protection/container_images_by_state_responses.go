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

// ContainerImagesByStateReader is a Reader for the ContainerImagesByState structure.
type ContainerImagesByStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerImagesByStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerImagesByStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewContainerImagesByStateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewContainerImagesByStateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerImagesByStateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/aggregates/containers/images-by-state/v1] ContainerImagesByState", response, response.Code())
	}
}

// NewContainerImagesByStateOK creates a ContainerImagesByStateOK with default headers values
func NewContainerImagesByStateOK() *ContainerImagesByStateOK {
	return &ContainerImagesByStateOK{}
}

/*
ContainerImagesByStateOK describes a response with status code 200, with default header values.

OK
*/
type ContainerImagesByStateOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ModelsAPIFilterResponse
}

// IsSuccess returns true when this container images by state o k response has a 2xx status code
func (o *ContainerImagesByStateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this container images by state o k response has a 3xx status code
func (o *ContainerImagesByStateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container images by state o k response has a 4xx status code
func (o *ContainerImagesByStateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this container images by state o k response has a 5xx status code
func (o *ContainerImagesByStateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this container images by state o k response a status code equal to that given
func (o *ContainerImagesByStateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the container images by state o k response
func (o *ContainerImagesByStateOK) Code() int {
	return 200
}

func (o *ContainerImagesByStateOK) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/images-by-state/v1][%d] containerImagesByStateOK  %+v", 200, o.Payload)
}

func (o *ContainerImagesByStateOK) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/images-by-state/v1][%d] containerImagesByStateOK  %+v", 200, o.Payload)
}

func (o *ContainerImagesByStateOK) GetPayload() *models.ModelsAPIFilterResponse {
	return o.Payload
}

func (o *ContainerImagesByStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ModelsAPIFilterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerImagesByStateForbidden creates a ContainerImagesByStateForbidden with default headers values
func NewContainerImagesByStateForbidden() *ContainerImagesByStateForbidden {
	return &ContainerImagesByStateForbidden{}
}

/*
ContainerImagesByStateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ContainerImagesByStateForbidden struct {

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

// IsSuccess returns true when this container images by state forbidden response has a 2xx status code
func (o *ContainerImagesByStateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container images by state forbidden response has a 3xx status code
func (o *ContainerImagesByStateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container images by state forbidden response has a 4xx status code
func (o *ContainerImagesByStateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this container images by state forbidden response has a 5xx status code
func (o *ContainerImagesByStateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this container images by state forbidden response a status code equal to that given
func (o *ContainerImagesByStateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the container images by state forbidden response
func (o *ContainerImagesByStateForbidden) Code() int {
	return 403
}

func (o *ContainerImagesByStateForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/images-by-state/v1][%d] containerImagesByStateForbidden  %+v", 403, o.Payload)
}

func (o *ContainerImagesByStateForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/images-by-state/v1][%d] containerImagesByStateForbidden  %+v", 403, o.Payload)
}

func (o *ContainerImagesByStateForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ContainerImagesByStateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewContainerImagesByStateTooManyRequests creates a ContainerImagesByStateTooManyRequests with default headers values
func NewContainerImagesByStateTooManyRequests() *ContainerImagesByStateTooManyRequests {
	return &ContainerImagesByStateTooManyRequests{}
}

/*
ContainerImagesByStateTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ContainerImagesByStateTooManyRequests struct {

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

// IsSuccess returns true when this container images by state too many requests response has a 2xx status code
func (o *ContainerImagesByStateTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container images by state too many requests response has a 3xx status code
func (o *ContainerImagesByStateTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container images by state too many requests response has a 4xx status code
func (o *ContainerImagesByStateTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this container images by state too many requests response has a 5xx status code
func (o *ContainerImagesByStateTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this container images by state too many requests response a status code equal to that given
func (o *ContainerImagesByStateTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the container images by state too many requests response
func (o *ContainerImagesByStateTooManyRequests) Code() int {
	return 429
}

func (o *ContainerImagesByStateTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/images-by-state/v1][%d] containerImagesByStateTooManyRequests  %+v", 429, o.Payload)
}

func (o *ContainerImagesByStateTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/images-by-state/v1][%d] containerImagesByStateTooManyRequests  %+v", 429, o.Payload)
}

func (o *ContainerImagesByStateTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ContainerImagesByStateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewContainerImagesByStateInternalServerError creates a ContainerImagesByStateInternalServerError with default headers values
func NewContainerImagesByStateInternalServerError() *ContainerImagesByStateInternalServerError {
	return &ContainerImagesByStateInternalServerError{}
}

/*
ContainerImagesByStateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ContainerImagesByStateInternalServerError struct {

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

// IsSuccess returns true when this container images by state internal server error response has a 2xx status code
func (o *ContainerImagesByStateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container images by state internal server error response has a 3xx status code
func (o *ContainerImagesByStateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container images by state internal server error response has a 4xx status code
func (o *ContainerImagesByStateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this container images by state internal server error response has a 5xx status code
func (o *ContainerImagesByStateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this container images by state internal server error response a status code equal to that given
func (o *ContainerImagesByStateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the container images by state internal server error response
func (o *ContainerImagesByStateInternalServerError) Code() int {
	return 500
}

func (o *ContainerImagesByStateInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/images-by-state/v1][%d] containerImagesByStateInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerImagesByStateInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/images-by-state/v1][%d] containerImagesByStateInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerImagesByStateInternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *ContainerImagesByStateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
