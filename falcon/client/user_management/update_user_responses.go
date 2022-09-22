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

// UpdateUserReader is a Reader for the UpdateUser structure.
type UpdateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateUserTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateUserOK creates a UpdateUserOK with default headers values
func NewUpdateUserOK() *UpdateUserOK {
	return &UpdateUserOK{}
}

/*
UpdateUserOK describes a response with status code 200, with default header values.

OK
*/
type UpdateUserOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIUserMetadataResponse
}

// IsSuccess returns true when this update user o k response has a 2xx status code
func (o *UpdateUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update user o k response has a 3xx status code
func (o *UpdateUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user o k response has a 4xx status code
func (o *UpdateUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update user o k response has a 5xx status code
func (o *UpdateUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update user o k response a status code equal to that given
func (o *UpdateUserOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateUserOK) Error() string {
	return fmt.Sprintf("[PATCH /users/entities/users/v1][%d] updateUserOK  %+v", 200, o.Payload)
}

func (o *UpdateUserOK) String() string {
	return fmt.Sprintf("[PATCH /users/entities/users/v1][%d] updateUserOK  %+v", 200, o.Payload)
}

func (o *UpdateUserOK) GetPayload() *models.APIUserMetadataResponse {
	return o.Payload
}

func (o *UpdateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIUserMetadataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserBadRequest creates a UpdateUserBadRequest with default headers values
func NewUpdateUserBadRequest() *UpdateUserBadRequest {
	return &UpdateUserBadRequest{}
}

/*
UpdateUserBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateUserBadRequest struct {

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

// IsSuccess returns true when this update user bad request response has a 2xx status code
func (o *UpdateUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user bad request response has a 3xx status code
func (o *UpdateUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user bad request response has a 4xx status code
func (o *UpdateUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user bad request response has a 5xx status code
func (o *UpdateUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update user bad request response a status code equal to that given
func (o *UpdateUserBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateUserBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /users/entities/users/v1][%d] updateUserBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateUserBadRequest) String() string {
	return fmt.Sprintf("[PATCH /users/entities/users/v1][%d] updateUserBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateUserBadRequest) GetPayload() *models.MsaEntitiesResponse {
	return o.Payload
}

func (o *UpdateUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateUserForbidden creates a UpdateUserForbidden with default headers values
func NewUpdateUserForbidden() *UpdateUserForbidden {
	return &UpdateUserForbidden{}
}

/*
UpdateUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateUserForbidden struct {

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

// IsSuccess returns true when this update user forbidden response has a 2xx status code
func (o *UpdateUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user forbidden response has a 3xx status code
func (o *UpdateUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user forbidden response has a 4xx status code
func (o *UpdateUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user forbidden response has a 5xx status code
func (o *UpdateUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update user forbidden response a status code equal to that given
func (o *UpdateUserForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateUserForbidden) Error() string {
	return fmt.Sprintf("[PATCH /users/entities/users/v1][%d] updateUserForbidden  %+v", 403, o.Payload)
}

func (o *UpdateUserForbidden) String() string {
	return fmt.Sprintf("[PATCH /users/entities/users/v1][%d] updateUserForbidden  %+v", 403, o.Payload)
}

func (o *UpdateUserForbidden) GetPayload() *models.MsaEntitiesResponse {
	return o.Payload
}

func (o *UpdateUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateUserNotFound creates a UpdateUserNotFound with default headers values
func NewUpdateUserNotFound() *UpdateUserNotFound {
	return &UpdateUserNotFound{}
}

/*
UpdateUserNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateUserNotFound struct {

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

// IsSuccess returns true when this update user not found response has a 2xx status code
func (o *UpdateUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user not found response has a 3xx status code
func (o *UpdateUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user not found response has a 4xx status code
func (o *UpdateUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user not found response has a 5xx status code
func (o *UpdateUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update user not found response a status code equal to that given
func (o *UpdateUserNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateUserNotFound) Error() string {
	return fmt.Sprintf("[PATCH /users/entities/users/v1][%d] updateUserNotFound  %+v", 404, o.Payload)
}

func (o *UpdateUserNotFound) String() string {
	return fmt.Sprintf("[PATCH /users/entities/users/v1][%d] updateUserNotFound  %+v", 404, o.Payload)
}

func (o *UpdateUserNotFound) GetPayload() *models.MsaEntitiesResponse {
	return o.Payload
}

func (o *UpdateUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateUserTooManyRequests creates a UpdateUserTooManyRequests with default headers values
func NewUpdateUserTooManyRequests() *UpdateUserTooManyRequests {
	return &UpdateUserTooManyRequests{}
}

/*
UpdateUserTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateUserTooManyRequests struct {

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

// IsSuccess returns true when this update user too many requests response has a 2xx status code
func (o *UpdateUserTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user too many requests response has a 3xx status code
func (o *UpdateUserTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user too many requests response has a 4xx status code
func (o *UpdateUserTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user too many requests response has a 5xx status code
func (o *UpdateUserTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update user too many requests response a status code equal to that given
func (o *UpdateUserTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *UpdateUserTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /users/entities/users/v1][%d] updateUserTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateUserTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /users/entities/users/v1][%d] updateUserTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateUserTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateUserTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateUserDefault creates a UpdateUserDefault with default headers values
func NewUpdateUserDefault(code int) *UpdateUserDefault {
	return &UpdateUserDefault{
		_statusCode: code,
	}
}

/*
UpdateUserDefault describes a response with status code -1, with default header values.

OK
*/
type UpdateUserDefault struct {
	_statusCode int

	Payload *models.APIUserMetadataResponse
}

// Code gets the status code for the update user default response
func (o *UpdateUserDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update user default response has a 2xx status code
func (o *UpdateUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update user default response has a 3xx status code
func (o *UpdateUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update user default response has a 4xx status code
func (o *UpdateUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update user default response has a 5xx status code
func (o *UpdateUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update user default response a status code equal to that given
func (o *UpdateUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdateUserDefault) Error() string {
	return fmt.Sprintf("[PATCH /users/entities/users/v1][%d] UpdateUser default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateUserDefault) String() string {
	return fmt.Sprintf("[PATCH /users/entities/users/v1][%d] UpdateUser default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateUserDefault) GetPayload() *models.APIUserMetadataResponse {
	return o.Payload
}

func (o *UpdateUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIUserMetadataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
