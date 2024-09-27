// Code generated by go-swagger; DO NOT EDIT.

package mssp

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

// AddUserGroupMembersReader is a Reader for the AddUserGroupMembers structure.
type AddUserGroupMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddUserGroupMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddUserGroupMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewAddUserGroupMembersMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddUserGroupMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddUserGroupMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAddUserGroupMembersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /mssp/entities/user-group-members/v1] addUserGroupMembers", response, response.Code())
	}
}

// NewAddUserGroupMembersOK creates a AddUserGroupMembersOK with default headers values
func NewAddUserGroupMembersOK() *AddUserGroupMembersOK {
	return &AddUserGroupMembersOK{}
}

/*
AddUserGroupMembersOK describes a response with status code 200, with default header values.

OK
*/
type AddUserGroupMembersOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainUserGroupMembersResponseV1
}

// IsSuccess returns true when this add user group members o k response has a 2xx status code
func (o *AddUserGroupMembersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add user group members o k response has a 3xx status code
func (o *AddUserGroupMembersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add user group members o k response has a 4xx status code
func (o *AddUserGroupMembersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add user group members o k response has a 5xx status code
func (o *AddUserGroupMembersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add user group members o k response a status code equal to that given
func (o *AddUserGroupMembersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add user group members o k response
func (o *AddUserGroupMembersOK) Code() int {
	return 200
}

func (o *AddUserGroupMembersOK) Error() string {
	return fmt.Sprintf("[POST /mssp/entities/user-group-members/v1][%d] addUserGroupMembersOK  %+v", 200, o.Payload)
}

func (o *AddUserGroupMembersOK) String() string {
	return fmt.Sprintf("[POST /mssp/entities/user-group-members/v1][%d] addUserGroupMembersOK  %+v", 200, o.Payload)
}

func (o *AddUserGroupMembersOK) GetPayload() *models.DomainUserGroupMembersResponseV1 {
	return o.Payload
}

func (o *AddUserGroupMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainUserGroupMembersResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserGroupMembersMultiStatus creates a AddUserGroupMembersMultiStatus with default headers values
func NewAddUserGroupMembersMultiStatus() *AddUserGroupMembersMultiStatus {
	return &AddUserGroupMembersMultiStatus{}
}

/*
AddUserGroupMembersMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type AddUserGroupMembersMultiStatus struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainUserGroupMembersResponseV1
}

// IsSuccess returns true when this add user group members multi status response has a 2xx status code
func (o *AddUserGroupMembersMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add user group members multi status response has a 3xx status code
func (o *AddUserGroupMembersMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add user group members multi status response has a 4xx status code
func (o *AddUserGroupMembersMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this add user group members multi status response has a 5xx status code
func (o *AddUserGroupMembersMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this add user group members multi status response a status code equal to that given
func (o *AddUserGroupMembersMultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the add user group members multi status response
func (o *AddUserGroupMembersMultiStatus) Code() int {
	return 207
}

func (o *AddUserGroupMembersMultiStatus) Error() string {
	return fmt.Sprintf("[POST /mssp/entities/user-group-members/v1][%d] addUserGroupMembersMultiStatus  %+v", 207, o.Payload)
}

func (o *AddUserGroupMembersMultiStatus) String() string {
	return fmt.Sprintf("[POST /mssp/entities/user-group-members/v1][%d] addUserGroupMembersMultiStatus  %+v", 207, o.Payload)
}

func (o *AddUserGroupMembersMultiStatus) GetPayload() *models.DomainUserGroupMembersResponseV1 {
	return o.Payload
}

func (o *AddUserGroupMembersMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainUserGroupMembersResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserGroupMembersBadRequest creates a AddUserGroupMembersBadRequest with default headers values
func NewAddUserGroupMembersBadRequest() *AddUserGroupMembersBadRequest {
	return &AddUserGroupMembersBadRequest{}
}

/*
AddUserGroupMembersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AddUserGroupMembersBadRequest struct {

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

// IsSuccess returns true when this add user group members bad request response has a 2xx status code
func (o *AddUserGroupMembersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add user group members bad request response has a 3xx status code
func (o *AddUserGroupMembersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add user group members bad request response has a 4xx status code
func (o *AddUserGroupMembersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add user group members bad request response has a 5xx status code
func (o *AddUserGroupMembersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add user group members bad request response a status code equal to that given
func (o *AddUserGroupMembersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the add user group members bad request response
func (o *AddUserGroupMembersBadRequest) Code() int {
	return 400
}

func (o *AddUserGroupMembersBadRequest) Error() string {
	return fmt.Sprintf("[POST /mssp/entities/user-group-members/v1][%d] addUserGroupMembersBadRequest  %+v", 400, o.Payload)
}

func (o *AddUserGroupMembersBadRequest) String() string {
	return fmt.Sprintf("[POST /mssp/entities/user-group-members/v1][%d] addUserGroupMembersBadRequest  %+v", 400, o.Payload)
}

func (o *AddUserGroupMembersBadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *AddUserGroupMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddUserGroupMembersForbidden creates a AddUserGroupMembersForbidden with default headers values
func NewAddUserGroupMembersForbidden() *AddUserGroupMembersForbidden {
	return &AddUserGroupMembersForbidden{}
}

/*
AddUserGroupMembersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AddUserGroupMembersForbidden struct {

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

// IsSuccess returns true when this add user group members forbidden response has a 2xx status code
func (o *AddUserGroupMembersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add user group members forbidden response has a 3xx status code
func (o *AddUserGroupMembersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add user group members forbidden response has a 4xx status code
func (o *AddUserGroupMembersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this add user group members forbidden response has a 5xx status code
func (o *AddUserGroupMembersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this add user group members forbidden response a status code equal to that given
func (o *AddUserGroupMembersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the add user group members forbidden response
func (o *AddUserGroupMembersForbidden) Code() int {
	return 403
}

func (o *AddUserGroupMembersForbidden) Error() string {
	return fmt.Sprintf("[POST /mssp/entities/user-group-members/v1][%d] addUserGroupMembersForbidden  %+v", 403, o.Payload)
}

func (o *AddUserGroupMembersForbidden) String() string {
	return fmt.Sprintf("[POST /mssp/entities/user-group-members/v1][%d] addUserGroupMembersForbidden  %+v", 403, o.Payload)
}

func (o *AddUserGroupMembersForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *AddUserGroupMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddUserGroupMembersTooManyRequests creates a AddUserGroupMembersTooManyRequests with default headers values
func NewAddUserGroupMembersTooManyRequests() *AddUserGroupMembersTooManyRequests {
	return &AddUserGroupMembersTooManyRequests{}
}

/*
AddUserGroupMembersTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AddUserGroupMembersTooManyRequests struct {

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

// IsSuccess returns true when this add user group members too many requests response has a 2xx status code
func (o *AddUserGroupMembersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add user group members too many requests response has a 3xx status code
func (o *AddUserGroupMembersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add user group members too many requests response has a 4xx status code
func (o *AddUserGroupMembersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this add user group members too many requests response has a 5xx status code
func (o *AddUserGroupMembersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this add user group members too many requests response a status code equal to that given
func (o *AddUserGroupMembersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the add user group members too many requests response
func (o *AddUserGroupMembersTooManyRequests) Code() int {
	return 429
}

func (o *AddUserGroupMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /mssp/entities/user-group-members/v1][%d] addUserGroupMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *AddUserGroupMembersTooManyRequests) String() string {
	return fmt.Sprintf("[POST /mssp/entities/user-group-members/v1][%d] addUserGroupMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *AddUserGroupMembersTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AddUserGroupMembersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
