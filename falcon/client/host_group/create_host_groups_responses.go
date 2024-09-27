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

// CreateHostGroupsReader is a Reader for the CreateHostGroups structure.
type CreateHostGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateHostGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateHostGroupsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateHostGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateHostGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateHostGroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateHostGroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /devices/entities/host-groups/v1] createHostGroups", response, response.Code())
	}
}

// NewCreateHostGroupsCreated creates a CreateHostGroupsCreated with default headers values
func NewCreateHostGroupsCreated() *CreateHostGroupsCreated {
	return &CreateHostGroupsCreated{}
}

/*
CreateHostGroupsCreated describes a response with status code 201, with default header values.

Created
*/
type CreateHostGroupsCreated struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.HostGroupsRespV1
}

// IsSuccess returns true when this create host groups created response has a 2xx status code
func (o *CreateHostGroupsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create host groups created response has a 3xx status code
func (o *CreateHostGroupsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create host groups created response has a 4xx status code
func (o *CreateHostGroupsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create host groups created response has a 5xx status code
func (o *CreateHostGroupsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create host groups created response a status code equal to that given
func (o *CreateHostGroupsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create host groups created response
func (o *CreateHostGroupsCreated) Code() int {
	return 201
}

func (o *CreateHostGroupsCreated) Error() string {
	return fmt.Sprintf("[POST /devices/entities/host-groups/v1][%d] createHostGroupsCreated  %+v", 201, o.Payload)
}

func (o *CreateHostGroupsCreated) String() string {
	return fmt.Sprintf("[POST /devices/entities/host-groups/v1][%d] createHostGroupsCreated  %+v", 201, o.Payload)
}

func (o *CreateHostGroupsCreated) GetPayload() *models.HostGroupsRespV1 {
	return o.Payload
}

func (o *CreateHostGroupsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.HostGroupsRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHostGroupsBadRequest creates a CreateHostGroupsBadRequest with default headers values
func NewCreateHostGroupsBadRequest() *CreateHostGroupsBadRequest {
	return &CreateHostGroupsBadRequest{}
}

/*
CreateHostGroupsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateHostGroupsBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.HostGroupsRespV1
}

// IsSuccess returns true when this create host groups bad request response has a 2xx status code
func (o *CreateHostGroupsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create host groups bad request response has a 3xx status code
func (o *CreateHostGroupsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create host groups bad request response has a 4xx status code
func (o *CreateHostGroupsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create host groups bad request response has a 5xx status code
func (o *CreateHostGroupsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create host groups bad request response a status code equal to that given
func (o *CreateHostGroupsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create host groups bad request response
func (o *CreateHostGroupsBadRequest) Code() int {
	return 400
}

func (o *CreateHostGroupsBadRequest) Error() string {
	return fmt.Sprintf("[POST /devices/entities/host-groups/v1][%d] createHostGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *CreateHostGroupsBadRequest) String() string {
	return fmt.Sprintf("[POST /devices/entities/host-groups/v1][%d] createHostGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *CreateHostGroupsBadRequest) GetPayload() *models.HostGroupsRespV1 {
	return o.Payload
}

func (o *CreateHostGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.HostGroupsRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHostGroupsForbidden creates a CreateHostGroupsForbidden with default headers values
func NewCreateHostGroupsForbidden() *CreateHostGroupsForbidden {
	return &CreateHostGroupsForbidden{}
}

/*
CreateHostGroupsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateHostGroupsForbidden struct {

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

// IsSuccess returns true when this create host groups forbidden response has a 2xx status code
func (o *CreateHostGroupsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create host groups forbidden response has a 3xx status code
func (o *CreateHostGroupsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create host groups forbidden response has a 4xx status code
func (o *CreateHostGroupsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create host groups forbidden response has a 5xx status code
func (o *CreateHostGroupsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create host groups forbidden response a status code equal to that given
func (o *CreateHostGroupsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create host groups forbidden response
func (o *CreateHostGroupsForbidden) Code() int {
	return 403
}

func (o *CreateHostGroupsForbidden) Error() string {
	return fmt.Sprintf("[POST /devices/entities/host-groups/v1][%d] createHostGroupsForbidden  %+v", 403, o.Payload)
}

func (o *CreateHostGroupsForbidden) String() string {
	return fmt.Sprintf("[POST /devices/entities/host-groups/v1][%d] createHostGroupsForbidden  %+v", 403, o.Payload)
}

func (o *CreateHostGroupsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateHostGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateHostGroupsTooManyRequests creates a CreateHostGroupsTooManyRequests with default headers values
func NewCreateHostGroupsTooManyRequests() *CreateHostGroupsTooManyRequests {
	return &CreateHostGroupsTooManyRequests{}
}

/*
CreateHostGroupsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateHostGroupsTooManyRequests struct {

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

// IsSuccess returns true when this create host groups too many requests response has a 2xx status code
func (o *CreateHostGroupsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create host groups too many requests response has a 3xx status code
func (o *CreateHostGroupsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create host groups too many requests response has a 4xx status code
func (o *CreateHostGroupsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create host groups too many requests response has a 5xx status code
func (o *CreateHostGroupsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create host groups too many requests response a status code equal to that given
func (o *CreateHostGroupsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the create host groups too many requests response
func (o *CreateHostGroupsTooManyRequests) Code() int {
	return 429
}

func (o *CreateHostGroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /devices/entities/host-groups/v1][%d] createHostGroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateHostGroupsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /devices/entities/host-groups/v1][%d] createHostGroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateHostGroupsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateHostGroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateHostGroupsInternalServerError creates a CreateHostGroupsInternalServerError with default headers values
func NewCreateHostGroupsInternalServerError() *CreateHostGroupsInternalServerError {
	return &CreateHostGroupsInternalServerError{}
}

/*
CreateHostGroupsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateHostGroupsInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.HostGroupsRespV1
}

// IsSuccess returns true when this create host groups internal server error response has a 2xx status code
func (o *CreateHostGroupsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create host groups internal server error response has a 3xx status code
func (o *CreateHostGroupsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create host groups internal server error response has a 4xx status code
func (o *CreateHostGroupsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create host groups internal server error response has a 5xx status code
func (o *CreateHostGroupsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create host groups internal server error response a status code equal to that given
func (o *CreateHostGroupsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create host groups internal server error response
func (o *CreateHostGroupsInternalServerError) Code() int {
	return 500
}

func (o *CreateHostGroupsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /devices/entities/host-groups/v1][%d] createHostGroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateHostGroupsInternalServerError) String() string {
	return fmt.Sprintf("[POST /devices/entities/host-groups/v1][%d] createHostGroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateHostGroupsInternalServerError) GetPayload() *models.HostGroupsRespV1 {
	return o.Payload
}

func (o *CreateHostGroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.HostGroupsRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
