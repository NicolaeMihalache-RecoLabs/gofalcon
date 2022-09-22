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

// GetRolesReader is a Reader for the GetRoles structure.
type GetRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRolesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRolesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetRolesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRolesOK creates a GetRolesOK with default headers values
func NewGetRolesOK() *GetRolesOK {
	return &GetRolesOK{}
}

/*
GetRolesOK describes a response with status code 200, with default header values.

OK
*/
type GetRolesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIUserRoleResponse
}

// IsSuccess returns true when this get roles o k response has a 2xx status code
func (o *GetRolesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get roles o k response has a 3xx status code
func (o *GetRolesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles o k response has a 4xx status code
func (o *GetRolesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get roles o k response has a 5xx status code
func (o *GetRolesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles o k response a status code equal to that given
func (o *GetRolesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRolesOK) Error() string {
	return fmt.Sprintf("[GET /user-roles/entities/user-roles/v1][%d] getRolesOK  %+v", 200, o.Payload)
}

func (o *GetRolesOK) String() string {
	return fmt.Sprintf("[GET /user-roles/entities/user-roles/v1][%d] getRolesOK  %+v", 200, o.Payload)
}

func (o *GetRolesOK) GetPayload() *models.APIUserRoleResponse {
	return o.Payload
}

func (o *GetRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIUserRoleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesBadRequest creates a GetRolesBadRequest with default headers values
func NewGetRolesBadRequest() *GetRolesBadRequest {
	return &GetRolesBadRequest{}
}

/*
GetRolesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetRolesBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaEntitiesResponse
}

// IsSuccess returns true when this get roles bad request response has a 2xx status code
func (o *GetRolesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles bad request response has a 3xx status code
func (o *GetRolesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles bad request response has a 4xx status code
func (o *GetRolesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get roles bad request response has a 5xx status code
func (o *GetRolesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles bad request response a status code equal to that given
func (o *GetRolesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRolesBadRequest) Error() string {
	return fmt.Sprintf("[GET /user-roles/entities/user-roles/v1][%d] getRolesBadRequest  %+v", 400, o.Payload)
}

func (o *GetRolesBadRequest) String() string {
	return fmt.Sprintf("[GET /user-roles/entities/user-roles/v1][%d] getRolesBadRequest  %+v", 400, o.Payload)
}

func (o *GetRolesBadRequest) GetPayload() *models.MsaEntitiesResponse {
	return o.Payload
}

func (o *GetRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesForbidden creates a GetRolesForbidden with default headers values
func NewGetRolesForbidden() *GetRolesForbidden {
	return &GetRolesForbidden{}
}

/*
GetRolesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRolesForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaEntitiesResponse
}

// IsSuccess returns true when this get roles forbidden response has a 2xx status code
func (o *GetRolesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles forbidden response has a 3xx status code
func (o *GetRolesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles forbidden response has a 4xx status code
func (o *GetRolesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get roles forbidden response has a 5xx status code
func (o *GetRolesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles forbidden response a status code equal to that given
func (o *GetRolesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /user-roles/entities/user-roles/v1][%d] getRolesForbidden  %+v", 403, o.Payload)
}

func (o *GetRolesForbidden) String() string {
	return fmt.Sprintf("[GET /user-roles/entities/user-roles/v1][%d] getRolesForbidden  %+v", 403, o.Payload)
}

func (o *GetRolesForbidden) GetPayload() *models.MsaEntitiesResponse {
	return o.Payload
}

func (o *GetRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesNotFound creates a GetRolesNotFound with default headers values
func NewGetRolesNotFound() *GetRolesNotFound {
	return &GetRolesNotFound{}
}

/*
GetRolesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetRolesNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaEntitiesResponse
}

// IsSuccess returns true when this get roles not found response has a 2xx status code
func (o *GetRolesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles not found response has a 3xx status code
func (o *GetRolesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles not found response has a 4xx status code
func (o *GetRolesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get roles not found response has a 5xx status code
func (o *GetRolesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles not found response a status code equal to that given
func (o *GetRolesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRolesNotFound) Error() string {
	return fmt.Sprintf("[GET /user-roles/entities/user-roles/v1][%d] getRolesNotFound  %+v", 404, o.Payload)
}

func (o *GetRolesNotFound) String() string {
	return fmt.Sprintf("[GET /user-roles/entities/user-roles/v1][%d] getRolesNotFound  %+v", 404, o.Payload)
}

func (o *GetRolesNotFound) GetPayload() *models.MsaEntitiesResponse {
	return o.Payload
}

func (o *GetRolesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesTooManyRequests creates a GetRolesTooManyRequests with default headers values
func NewGetRolesTooManyRequests() *GetRolesTooManyRequests {
	return &GetRolesTooManyRequests{}
}

/*
GetRolesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetRolesTooManyRequests struct {

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

// IsSuccess returns true when this get roles too many requests response has a 2xx status code
func (o *GetRolesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles too many requests response has a 3xx status code
func (o *GetRolesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles too many requests response has a 4xx status code
func (o *GetRolesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get roles too many requests response has a 5xx status code
func (o *GetRolesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles too many requests response a status code equal to that given
func (o *GetRolesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRolesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /user-roles/entities/user-roles/v1][%d] getRolesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRolesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /user-roles/entities/user-roles/v1][%d] getRolesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRolesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetRolesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRolesInternalServerError creates a GetRolesInternalServerError with default headers values
func NewGetRolesInternalServerError() *GetRolesInternalServerError {
	return &GetRolesInternalServerError{}
}

/*
GetRolesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetRolesInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaEntitiesResponse
}

// IsSuccess returns true when this get roles internal server error response has a 2xx status code
func (o *GetRolesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles internal server error response has a 3xx status code
func (o *GetRolesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles internal server error response has a 4xx status code
func (o *GetRolesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get roles internal server error response has a 5xx status code
func (o *GetRolesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get roles internal server error response a status code equal to that given
func (o *GetRolesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRolesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /user-roles/entities/user-roles/v1][%d] getRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRolesInternalServerError) String() string {
	return fmt.Sprintf("[GET /user-roles/entities/user-roles/v1][%d] getRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRolesInternalServerError) GetPayload() *models.MsaEntitiesResponse {
	return o.Payload
}

func (o *GetRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesDefault creates a GetRolesDefault with default headers values
func NewGetRolesDefault(code int) *GetRolesDefault {
	return &GetRolesDefault{
		_statusCode: code,
	}
}

/*
GetRolesDefault describes a response with status code -1, with default header values.

OK
*/
type GetRolesDefault struct {
	_statusCode int

	Payload *models.APIUserRoleResponse
}

// Code gets the status code for the get roles default response
func (o *GetRolesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get roles default response has a 2xx status code
func (o *GetRolesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get roles default response has a 3xx status code
func (o *GetRolesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get roles default response has a 4xx status code
func (o *GetRolesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get roles default response has a 5xx status code
func (o *GetRolesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get roles default response a status code equal to that given
func (o *GetRolesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetRolesDefault) Error() string {
	return fmt.Sprintf("[GET /user-roles/entities/user-roles/v1][%d] GetRoles default  %+v", o._statusCode, o.Payload)
}

func (o *GetRolesDefault) String() string {
	return fmt.Sprintf("[GET /user-roles/entities/user-roles/v1][%d] GetRoles default  %+v", o._statusCode, o.Payload)
}

func (o *GetRolesDefault) GetPayload() *models.APIUserRoleResponse {
	return o.Payload
}

func (o *GetRolesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIUserRoleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
