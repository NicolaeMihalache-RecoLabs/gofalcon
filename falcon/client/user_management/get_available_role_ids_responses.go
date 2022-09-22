// Code generated by go-swagger; DO NOT EDIT.

package user_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// GetAvailableRoleIdsReader is a Reader for the GetAvailableRoleIds structure.
type GetAvailableRoleIdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAvailableRoleIdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAvailableRoleIdsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetAvailableRoleIdsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAvailableRoleIdsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAvailableRoleIdsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAvailableRoleIdsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAvailableRoleIdsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAvailableRoleIdsOK creates a GetAvailableRoleIdsOK with default headers values
func NewGetAvailableRoleIdsOK() *GetAvailableRoleIdsOK {
	return &GetAvailableRoleIdsOK{}
}

/*
GetAvailableRoleIdsOK describes a response with status code 200, with default header values.

OK
*/
type GetAvailableRoleIdsOK struct {

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

// IsSuccess returns true when this get available role ids o k response has a 2xx status code
func (o *GetAvailableRoleIdsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get available role ids o k response has a 3xx status code
func (o *GetAvailableRoleIdsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get available role ids o k response has a 4xx status code
func (o *GetAvailableRoleIdsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get available role ids o k response has a 5xx status code
func (o *GetAvailableRoleIdsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get available role ids o k response a status code equal to that given
func (o *GetAvailableRoleIdsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAvailableRoleIdsOK) Error() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-cid/v1][%d] getAvailableRoleIdsOK  %+v", 200, o.Payload)
}

func (o *GetAvailableRoleIdsOK) String() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-cid/v1][%d] getAvailableRoleIdsOK  %+v", 200, o.Payload)
}

func (o *GetAvailableRoleIdsOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *GetAvailableRoleIdsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAvailableRoleIdsForbidden creates a GetAvailableRoleIdsForbidden with default headers values
func NewGetAvailableRoleIdsForbidden() *GetAvailableRoleIdsForbidden {
	return &GetAvailableRoleIdsForbidden{}
}

/*
GetAvailableRoleIdsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAvailableRoleIdsForbidden struct {

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

// IsSuccess returns true when this get available role ids forbidden response has a 2xx status code
func (o *GetAvailableRoleIdsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get available role ids forbidden response has a 3xx status code
func (o *GetAvailableRoleIdsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get available role ids forbidden response has a 4xx status code
func (o *GetAvailableRoleIdsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get available role ids forbidden response has a 5xx status code
func (o *GetAvailableRoleIdsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get available role ids forbidden response a status code equal to that given
func (o *GetAvailableRoleIdsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAvailableRoleIdsForbidden) Error() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-cid/v1][%d] getAvailableRoleIdsForbidden  %+v", 403, o.Payload)
}

func (o *GetAvailableRoleIdsForbidden) String() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-cid/v1][%d] getAvailableRoleIdsForbidden  %+v", 403, o.Payload)
}

func (o *GetAvailableRoleIdsForbidden) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *GetAvailableRoleIdsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAvailableRoleIdsNotFound creates a GetAvailableRoleIdsNotFound with default headers values
func NewGetAvailableRoleIdsNotFound() *GetAvailableRoleIdsNotFound {
	return &GetAvailableRoleIdsNotFound{}
}

/*
GetAvailableRoleIdsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAvailableRoleIdsNotFound struct {

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

// IsSuccess returns true when this get available role ids not found response has a 2xx status code
func (o *GetAvailableRoleIdsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get available role ids not found response has a 3xx status code
func (o *GetAvailableRoleIdsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get available role ids not found response has a 4xx status code
func (o *GetAvailableRoleIdsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get available role ids not found response has a 5xx status code
func (o *GetAvailableRoleIdsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get available role ids not found response a status code equal to that given
func (o *GetAvailableRoleIdsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAvailableRoleIdsNotFound) Error() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-cid/v1][%d] getAvailableRoleIdsNotFound  %+v", 404, o.Payload)
}

func (o *GetAvailableRoleIdsNotFound) String() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-cid/v1][%d] getAvailableRoleIdsNotFound  %+v", 404, o.Payload)
}

func (o *GetAvailableRoleIdsNotFound) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *GetAvailableRoleIdsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAvailableRoleIdsTooManyRequests creates a GetAvailableRoleIdsTooManyRequests with default headers values
func NewGetAvailableRoleIdsTooManyRequests() *GetAvailableRoleIdsTooManyRequests {
	return &GetAvailableRoleIdsTooManyRequests{}
}

/*
GetAvailableRoleIdsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetAvailableRoleIdsTooManyRequests struct {

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

// IsSuccess returns true when this get available role ids too many requests response has a 2xx status code
func (o *GetAvailableRoleIdsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get available role ids too many requests response has a 3xx status code
func (o *GetAvailableRoleIdsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get available role ids too many requests response has a 4xx status code
func (o *GetAvailableRoleIdsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get available role ids too many requests response has a 5xx status code
func (o *GetAvailableRoleIdsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get available role ids too many requests response a status code equal to that given
func (o *GetAvailableRoleIdsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetAvailableRoleIdsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-cid/v1][%d] getAvailableRoleIdsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAvailableRoleIdsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-cid/v1][%d] getAvailableRoleIdsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAvailableRoleIdsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetAvailableRoleIdsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAvailableRoleIdsInternalServerError creates a GetAvailableRoleIdsInternalServerError with default headers values
func NewGetAvailableRoleIdsInternalServerError() *GetAvailableRoleIdsInternalServerError {
	return &GetAvailableRoleIdsInternalServerError{}
}

/*
GetAvailableRoleIdsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAvailableRoleIdsInternalServerError struct {

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

// IsSuccess returns true when this get available role ids internal server error response has a 2xx status code
func (o *GetAvailableRoleIdsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get available role ids internal server error response has a 3xx status code
func (o *GetAvailableRoleIdsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get available role ids internal server error response has a 4xx status code
func (o *GetAvailableRoleIdsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get available role ids internal server error response has a 5xx status code
func (o *GetAvailableRoleIdsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get available role ids internal server error response a status code equal to that given
func (o *GetAvailableRoleIdsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAvailableRoleIdsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-cid/v1][%d] getAvailableRoleIdsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAvailableRoleIdsInternalServerError) String() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-cid/v1][%d] getAvailableRoleIdsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAvailableRoleIdsInternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *GetAvailableRoleIdsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAvailableRoleIdsDefault creates a GetAvailableRoleIdsDefault with default headers values
func NewGetAvailableRoleIdsDefault(code int) *GetAvailableRoleIdsDefault {
	return &GetAvailableRoleIdsDefault{
		_statusCode: code,
	}
}

/*
GetAvailableRoleIdsDefault describes a response with status code -1, with default header values.

OK
*/
type GetAvailableRoleIdsDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// Code gets the status code for the get available role ids default response
func (o *GetAvailableRoleIdsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get available role ids default response has a 2xx status code
func (o *GetAvailableRoleIdsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get available role ids default response has a 3xx status code
func (o *GetAvailableRoleIdsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get available role ids default response has a 4xx status code
func (o *GetAvailableRoleIdsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get available role ids default response has a 5xx status code
func (o *GetAvailableRoleIdsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get available role ids default response a status code equal to that given
func (o *GetAvailableRoleIdsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetAvailableRoleIdsDefault) Error() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-cid/v1][%d] GetAvailableRoleIds default  %+v", o._statusCode, o.Payload)
}

func (o *GetAvailableRoleIdsDefault) String() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-cid/v1][%d] GetAvailableRoleIds default  %+v", o._statusCode, o.Payload)
}

func (o *GetAvailableRoleIdsDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *GetAvailableRoleIdsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
