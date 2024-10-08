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

// GroupContainersByManagedReader is a Reader for the GroupContainersByManaged structure.
type GroupContainersByManagedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupContainersByManagedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupContainersByManagedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGroupContainersByManagedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGroupContainersByManagedTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGroupContainersByManagedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/aggregates/containers/group-by-managed/v1] GroupContainersByManaged", response, response.Code())
	}
}

// NewGroupContainersByManagedOK creates a GroupContainersByManagedOK with default headers values
func NewGroupContainersByManagedOK() *GroupContainersByManagedOK {
	return &GroupContainersByManagedOK{}
}

/*
GroupContainersByManagedOK describes a response with status code 200, with default header values.

OK
*/
type GroupContainersByManagedOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ModelsContainerCoverageResponseEntity
}

// IsSuccess returns true when this group containers by managed o k response has a 2xx status code
func (o *GroupContainersByManagedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this group containers by managed o k response has a 3xx status code
func (o *GroupContainersByManagedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this group containers by managed o k response has a 4xx status code
func (o *GroupContainersByManagedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this group containers by managed o k response has a 5xx status code
func (o *GroupContainersByManagedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this group containers by managed o k response a status code equal to that given
func (o *GroupContainersByManagedOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the group containers by managed o k response
func (o *GroupContainersByManagedOK) Code() int {
	return 200
}

func (o *GroupContainersByManagedOK) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/group-by-managed/v1][%d] groupContainersByManagedOK  %+v", 200, o.Payload)
}

func (o *GroupContainersByManagedOK) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/group-by-managed/v1][%d] groupContainersByManagedOK  %+v", 200, o.Payload)
}

func (o *GroupContainersByManagedOK) GetPayload() *models.ModelsContainerCoverageResponseEntity {
	return o.Payload
}

func (o *GroupContainersByManagedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ModelsContainerCoverageResponseEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupContainersByManagedForbidden creates a GroupContainersByManagedForbidden with default headers values
func NewGroupContainersByManagedForbidden() *GroupContainersByManagedForbidden {
	return &GroupContainersByManagedForbidden{}
}

/*
GroupContainersByManagedForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GroupContainersByManagedForbidden struct {

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

// IsSuccess returns true when this group containers by managed forbidden response has a 2xx status code
func (o *GroupContainersByManagedForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this group containers by managed forbidden response has a 3xx status code
func (o *GroupContainersByManagedForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this group containers by managed forbidden response has a 4xx status code
func (o *GroupContainersByManagedForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this group containers by managed forbidden response has a 5xx status code
func (o *GroupContainersByManagedForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this group containers by managed forbidden response a status code equal to that given
func (o *GroupContainersByManagedForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the group containers by managed forbidden response
func (o *GroupContainersByManagedForbidden) Code() int {
	return 403
}

func (o *GroupContainersByManagedForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/group-by-managed/v1][%d] groupContainersByManagedForbidden  %+v", 403, o.Payload)
}

func (o *GroupContainersByManagedForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/group-by-managed/v1][%d] groupContainersByManagedForbidden  %+v", 403, o.Payload)
}

func (o *GroupContainersByManagedForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GroupContainersByManagedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGroupContainersByManagedTooManyRequests creates a GroupContainersByManagedTooManyRequests with default headers values
func NewGroupContainersByManagedTooManyRequests() *GroupContainersByManagedTooManyRequests {
	return &GroupContainersByManagedTooManyRequests{}
}

/*
GroupContainersByManagedTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GroupContainersByManagedTooManyRequests struct {

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

// IsSuccess returns true when this group containers by managed too many requests response has a 2xx status code
func (o *GroupContainersByManagedTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this group containers by managed too many requests response has a 3xx status code
func (o *GroupContainersByManagedTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this group containers by managed too many requests response has a 4xx status code
func (o *GroupContainersByManagedTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this group containers by managed too many requests response has a 5xx status code
func (o *GroupContainersByManagedTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this group containers by managed too many requests response a status code equal to that given
func (o *GroupContainersByManagedTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the group containers by managed too many requests response
func (o *GroupContainersByManagedTooManyRequests) Code() int {
	return 429
}

func (o *GroupContainersByManagedTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/group-by-managed/v1][%d] groupContainersByManagedTooManyRequests  %+v", 429, o.Payload)
}

func (o *GroupContainersByManagedTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/group-by-managed/v1][%d] groupContainersByManagedTooManyRequests  %+v", 429, o.Payload)
}

func (o *GroupContainersByManagedTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GroupContainersByManagedTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGroupContainersByManagedInternalServerError creates a GroupContainersByManagedInternalServerError with default headers values
func NewGroupContainersByManagedInternalServerError() *GroupContainersByManagedInternalServerError {
	return &GroupContainersByManagedInternalServerError{}
}

/*
GroupContainersByManagedInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GroupContainersByManagedInternalServerError struct {

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

// IsSuccess returns true when this group containers by managed internal server error response has a 2xx status code
func (o *GroupContainersByManagedInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this group containers by managed internal server error response has a 3xx status code
func (o *GroupContainersByManagedInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this group containers by managed internal server error response has a 4xx status code
func (o *GroupContainersByManagedInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this group containers by managed internal server error response has a 5xx status code
func (o *GroupContainersByManagedInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this group containers by managed internal server error response a status code equal to that given
func (o *GroupContainersByManagedInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the group containers by managed internal server error response
func (o *GroupContainersByManagedInternalServerError) Code() int {
	return 500
}

func (o *GroupContainersByManagedInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/group-by-managed/v1][%d] groupContainersByManagedInternalServerError  %+v", 500, o.Payload)
}

func (o *GroupContainersByManagedInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/group-by-managed/v1][%d] groupContainersByManagedInternalServerError  %+v", 500, o.Payload)
}

func (o *GroupContainersByManagedInternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *GroupContainersByManagedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
