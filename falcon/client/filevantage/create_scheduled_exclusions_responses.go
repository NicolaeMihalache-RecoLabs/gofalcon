// Code generated by go-swagger; DO NOT EDIT.

package filevantage

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

// CreateScheduledExclusionsReader is a Reader for the CreateScheduledExclusions structure.
type CreateScheduledExclusionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateScheduledExclusionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateScheduledExclusionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateScheduledExclusionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateScheduledExclusionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateScheduledExclusionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateScheduledExclusionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateScheduledExclusionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /filevantage/entities/policy-scheduled-exclusions/v1] createScheduledExclusions", response, response.Code())
	}
}

// NewCreateScheduledExclusionsOK creates a CreateScheduledExclusionsOK with default headers values
func NewCreateScheduledExclusionsOK() *CreateScheduledExclusionsOK {
	return &CreateScheduledExclusionsOK{}
}

/*
CreateScheduledExclusionsOK describes a response with status code 200, with default header values.

OK
*/
type CreateScheduledExclusionsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ScheduledexclusionsResponse
}

// IsSuccess returns true when this create scheduled exclusions o k response has a 2xx status code
func (o *CreateScheduledExclusionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create scheduled exclusions o k response has a 3xx status code
func (o *CreateScheduledExclusionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create scheduled exclusions o k response has a 4xx status code
func (o *CreateScheduledExclusionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create scheduled exclusions o k response has a 5xx status code
func (o *CreateScheduledExclusionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create scheduled exclusions o k response a status code equal to that given
func (o *CreateScheduledExclusionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create scheduled exclusions o k response
func (o *CreateScheduledExclusionsOK) Code() int {
	return 200
}

func (o *CreateScheduledExclusionsOK) Error() string {
	return fmt.Sprintf("[POST /filevantage/entities/policy-scheduled-exclusions/v1][%d] createScheduledExclusionsOK  %+v", 200, o.Payload)
}

func (o *CreateScheduledExclusionsOK) String() string {
	return fmt.Sprintf("[POST /filevantage/entities/policy-scheduled-exclusions/v1][%d] createScheduledExclusionsOK  %+v", 200, o.Payload)
}

func (o *CreateScheduledExclusionsOK) GetPayload() *models.ScheduledexclusionsResponse {
	return o.Payload
}

func (o *CreateScheduledExclusionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ScheduledexclusionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateScheduledExclusionsBadRequest creates a CreateScheduledExclusionsBadRequest with default headers values
func NewCreateScheduledExclusionsBadRequest() *CreateScheduledExclusionsBadRequest {
	return &CreateScheduledExclusionsBadRequest{}
}

/*
CreateScheduledExclusionsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateScheduledExclusionsBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this create scheduled exclusions bad request response has a 2xx status code
func (o *CreateScheduledExclusionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create scheduled exclusions bad request response has a 3xx status code
func (o *CreateScheduledExclusionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create scheduled exclusions bad request response has a 4xx status code
func (o *CreateScheduledExclusionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create scheduled exclusions bad request response has a 5xx status code
func (o *CreateScheduledExclusionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create scheduled exclusions bad request response a status code equal to that given
func (o *CreateScheduledExclusionsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create scheduled exclusions bad request response
func (o *CreateScheduledExclusionsBadRequest) Code() int {
	return 400
}

func (o *CreateScheduledExclusionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /filevantage/entities/policy-scheduled-exclusions/v1][%d] createScheduledExclusionsBadRequest  %+v", 400, o.Payload)
}

func (o *CreateScheduledExclusionsBadRequest) String() string {
	return fmt.Sprintf("[POST /filevantage/entities/policy-scheduled-exclusions/v1][%d] createScheduledExclusionsBadRequest  %+v", 400, o.Payload)
}

func (o *CreateScheduledExclusionsBadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CreateScheduledExclusionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateScheduledExclusionsForbidden creates a CreateScheduledExclusionsForbidden with default headers values
func NewCreateScheduledExclusionsForbidden() *CreateScheduledExclusionsForbidden {
	return &CreateScheduledExclusionsForbidden{}
}

/*
CreateScheduledExclusionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateScheduledExclusionsForbidden struct {

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

// IsSuccess returns true when this create scheduled exclusions forbidden response has a 2xx status code
func (o *CreateScheduledExclusionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create scheduled exclusions forbidden response has a 3xx status code
func (o *CreateScheduledExclusionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create scheduled exclusions forbidden response has a 4xx status code
func (o *CreateScheduledExclusionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create scheduled exclusions forbidden response has a 5xx status code
func (o *CreateScheduledExclusionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create scheduled exclusions forbidden response a status code equal to that given
func (o *CreateScheduledExclusionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create scheduled exclusions forbidden response
func (o *CreateScheduledExclusionsForbidden) Code() int {
	return 403
}

func (o *CreateScheduledExclusionsForbidden) Error() string {
	return fmt.Sprintf("[POST /filevantage/entities/policy-scheduled-exclusions/v1][%d] createScheduledExclusionsForbidden  %+v", 403, o.Payload)
}

func (o *CreateScheduledExclusionsForbidden) String() string {
	return fmt.Sprintf("[POST /filevantage/entities/policy-scheduled-exclusions/v1][%d] createScheduledExclusionsForbidden  %+v", 403, o.Payload)
}

func (o *CreateScheduledExclusionsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateScheduledExclusionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateScheduledExclusionsNotFound creates a CreateScheduledExclusionsNotFound with default headers values
func NewCreateScheduledExclusionsNotFound() *CreateScheduledExclusionsNotFound {
	return &CreateScheduledExclusionsNotFound{}
}

/*
CreateScheduledExclusionsNotFound describes a response with status code 404, with default header values.

The policy to add the scheduled exclusion to does not exist.
*/
type CreateScheduledExclusionsNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this create scheduled exclusions not found response has a 2xx status code
func (o *CreateScheduledExclusionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create scheduled exclusions not found response has a 3xx status code
func (o *CreateScheduledExclusionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create scheduled exclusions not found response has a 4xx status code
func (o *CreateScheduledExclusionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create scheduled exclusions not found response has a 5xx status code
func (o *CreateScheduledExclusionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create scheduled exclusions not found response a status code equal to that given
func (o *CreateScheduledExclusionsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create scheduled exclusions not found response
func (o *CreateScheduledExclusionsNotFound) Code() int {
	return 404
}

func (o *CreateScheduledExclusionsNotFound) Error() string {
	return fmt.Sprintf("[POST /filevantage/entities/policy-scheduled-exclusions/v1][%d] createScheduledExclusionsNotFound  %+v", 404, o.Payload)
}

func (o *CreateScheduledExclusionsNotFound) String() string {
	return fmt.Sprintf("[POST /filevantage/entities/policy-scheduled-exclusions/v1][%d] createScheduledExclusionsNotFound  %+v", 404, o.Payload)
}

func (o *CreateScheduledExclusionsNotFound) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CreateScheduledExclusionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateScheduledExclusionsTooManyRequests creates a CreateScheduledExclusionsTooManyRequests with default headers values
func NewCreateScheduledExclusionsTooManyRequests() *CreateScheduledExclusionsTooManyRequests {
	return &CreateScheduledExclusionsTooManyRequests{}
}

/*
CreateScheduledExclusionsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateScheduledExclusionsTooManyRequests struct {

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

// IsSuccess returns true when this create scheduled exclusions too many requests response has a 2xx status code
func (o *CreateScheduledExclusionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create scheduled exclusions too many requests response has a 3xx status code
func (o *CreateScheduledExclusionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create scheduled exclusions too many requests response has a 4xx status code
func (o *CreateScheduledExclusionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create scheduled exclusions too many requests response has a 5xx status code
func (o *CreateScheduledExclusionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create scheduled exclusions too many requests response a status code equal to that given
func (o *CreateScheduledExclusionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the create scheduled exclusions too many requests response
func (o *CreateScheduledExclusionsTooManyRequests) Code() int {
	return 429
}

func (o *CreateScheduledExclusionsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /filevantage/entities/policy-scheduled-exclusions/v1][%d] createScheduledExclusionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateScheduledExclusionsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /filevantage/entities/policy-scheduled-exclusions/v1][%d] createScheduledExclusionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateScheduledExclusionsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateScheduledExclusionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateScheduledExclusionsInternalServerError creates a CreateScheduledExclusionsInternalServerError with default headers values
func NewCreateScheduledExclusionsInternalServerError() *CreateScheduledExclusionsInternalServerError {
	return &CreateScheduledExclusionsInternalServerError{}
}

/*
CreateScheduledExclusionsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateScheduledExclusionsInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this create scheduled exclusions internal server error response has a 2xx status code
func (o *CreateScheduledExclusionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create scheduled exclusions internal server error response has a 3xx status code
func (o *CreateScheduledExclusionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create scheduled exclusions internal server error response has a 4xx status code
func (o *CreateScheduledExclusionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create scheduled exclusions internal server error response has a 5xx status code
func (o *CreateScheduledExclusionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create scheduled exclusions internal server error response a status code equal to that given
func (o *CreateScheduledExclusionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create scheduled exclusions internal server error response
func (o *CreateScheduledExclusionsInternalServerError) Code() int {
	return 500
}

func (o *CreateScheduledExclusionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /filevantage/entities/policy-scheduled-exclusions/v1][%d] createScheduledExclusionsInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateScheduledExclusionsInternalServerError) String() string {
	return fmt.Sprintf("[POST /filevantage/entities/policy-scheduled-exclusions/v1][%d] createScheduledExclusionsInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateScheduledExclusionsInternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CreateScheduledExclusionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
