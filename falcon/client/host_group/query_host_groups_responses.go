// Code generated by go-swagger; DO NOT EDIT.

package host_group

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

// QueryHostGroupsReader is a Reader for the QueryHostGroups structure.
type QueryHostGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryHostGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryHostGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryHostGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryHostGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryHostGroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryHostGroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /devices/queries/host-groups/v1] queryHostGroups", response, response.Code())
	}
}

// NewQueryHostGroupsOK creates a QueryHostGroupsOK with default headers values
func NewQueryHostGroupsOK() *QueryHostGroupsOK {
	return &QueryHostGroupsOK{}
}

/*
QueryHostGroupsOK describes a response with status code 200, with default header values.

OK
*/
type QueryHostGroupsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query host groups o k response has a 2xx status code
func (o *QueryHostGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query host groups o k response has a 3xx status code
func (o *QueryHostGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query host groups o k response has a 4xx status code
func (o *QueryHostGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query host groups o k response has a 5xx status code
func (o *QueryHostGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query host groups o k response a status code equal to that given
func (o *QueryHostGroupsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query host groups o k response
func (o *QueryHostGroupsOK) Code() int {
	return 200
}

func (o *QueryHostGroupsOK) Error() string {
	return fmt.Sprintf("[GET /devices/queries/host-groups/v1][%d] queryHostGroupsOK  %+v", 200, o.Payload)
}

func (o *QueryHostGroupsOK) String() string {
	return fmt.Sprintf("[GET /devices/queries/host-groups/v1][%d] queryHostGroupsOK  %+v", 200, o.Payload)
}

func (o *QueryHostGroupsOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryHostGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryHostGroupsBadRequest creates a QueryHostGroupsBadRequest with default headers values
func NewQueryHostGroupsBadRequest() *QueryHostGroupsBadRequest {
	return &QueryHostGroupsBadRequest{}
}

/*
QueryHostGroupsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryHostGroupsBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query host groups bad request response has a 2xx status code
func (o *QueryHostGroupsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query host groups bad request response has a 3xx status code
func (o *QueryHostGroupsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query host groups bad request response has a 4xx status code
func (o *QueryHostGroupsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query host groups bad request response has a 5xx status code
func (o *QueryHostGroupsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query host groups bad request response a status code equal to that given
func (o *QueryHostGroupsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query host groups bad request response
func (o *QueryHostGroupsBadRequest) Code() int {
	return 400
}

func (o *QueryHostGroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/queries/host-groups/v1][%d] queryHostGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *QueryHostGroupsBadRequest) String() string {
	return fmt.Sprintf("[GET /devices/queries/host-groups/v1][%d] queryHostGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *QueryHostGroupsBadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryHostGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryHostGroupsForbidden creates a QueryHostGroupsForbidden with default headers values
func NewQueryHostGroupsForbidden() *QueryHostGroupsForbidden {
	return &QueryHostGroupsForbidden{}
}

/*
QueryHostGroupsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryHostGroupsForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this query host groups forbidden response has a 2xx status code
func (o *QueryHostGroupsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query host groups forbidden response has a 3xx status code
func (o *QueryHostGroupsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query host groups forbidden response has a 4xx status code
func (o *QueryHostGroupsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query host groups forbidden response has a 5xx status code
func (o *QueryHostGroupsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query host groups forbidden response a status code equal to that given
func (o *QueryHostGroupsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query host groups forbidden response
func (o *QueryHostGroupsForbidden) Code() int {
	return 403
}

func (o *QueryHostGroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/queries/host-groups/v1][%d] queryHostGroupsForbidden  %+v", 403, o.Payload)
}

func (o *QueryHostGroupsForbidden) String() string {
	return fmt.Sprintf("[GET /devices/queries/host-groups/v1][%d] queryHostGroupsForbidden  %+v", 403, o.Payload)
}

func (o *QueryHostGroupsForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryHostGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryHostGroupsTooManyRequests creates a QueryHostGroupsTooManyRequests with default headers values
func NewQueryHostGroupsTooManyRequests() *QueryHostGroupsTooManyRequests {
	return &QueryHostGroupsTooManyRequests{}
}

/*
QueryHostGroupsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryHostGroupsTooManyRequests struct {

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

// IsSuccess returns true when this query host groups too many requests response has a 2xx status code
func (o *QueryHostGroupsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query host groups too many requests response has a 3xx status code
func (o *QueryHostGroupsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query host groups too many requests response has a 4xx status code
func (o *QueryHostGroupsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query host groups too many requests response has a 5xx status code
func (o *QueryHostGroupsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query host groups too many requests response a status code equal to that given
func (o *QueryHostGroupsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query host groups too many requests response
func (o *QueryHostGroupsTooManyRequests) Code() int {
	return 429
}

func (o *QueryHostGroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /devices/queries/host-groups/v1][%d] queryHostGroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryHostGroupsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /devices/queries/host-groups/v1][%d] queryHostGroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryHostGroupsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryHostGroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryHostGroupsInternalServerError creates a QueryHostGroupsInternalServerError with default headers values
func NewQueryHostGroupsInternalServerError() *QueryHostGroupsInternalServerError {
	return &QueryHostGroupsInternalServerError{}
}

/*
QueryHostGroupsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryHostGroupsInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query host groups internal server error response has a 2xx status code
func (o *QueryHostGroupsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query host groups internal server error response has a 3xx status code
func (o *QueryHostGroupsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query host groups internal server error response has a 4xx status code
func (o *QueryHostGroupsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query host groups internal server error response has a 5xx status code
func (o *QueryHostGroupsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query host groups internal server error response a status code equal to that given
func (o *QueryHostGroupsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query host groups internal server error response
func (o *QueryHostGroupsInternalServerError) Code() int {
	return 500
}

func (o *QueryHostGroupsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/queries/host-groups/v1][%d] queryHostGroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryHostGroupsInternalServerError) String() string {
	return fmt.Sprintf("[GET /devices/queries/host-groups/v1][%d] queryHostGroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryHostGroupsInternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryHostGroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
