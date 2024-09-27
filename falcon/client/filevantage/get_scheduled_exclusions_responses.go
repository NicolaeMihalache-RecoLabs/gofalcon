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

// GetScheduledExclusionsReader is a Reader for the GetScheduledExclusions structure.
type GetScheduledExclusionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScheduledExclusionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScheduledExclusionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetScheduledExclusionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScheduledExclusionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScheduledExclusionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetScheduledExclusionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScheduledExclusionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /filevantage/entities/policy-scheduled-exclusions/v1] getScheduledExclusions", response, response.Code())
	}
}

// NewGetScheduledExclusionsOK creates a GetScheduledExclusionsOK with default headers values
func NewGetScheduledExclusionsOK() *GetScheduledExclusionsOK {
	return &GetScheduledExclusionsOK{}
}

/*
GetScheduledExclusionsOK describes a response with status code 200, with default header values.

OK
*/
type GetScheduledExclusionsOK struct {

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

// IsSuccess returns true when this get scheduled exclusions o k response has a 2xx status code
func (o *GetScheduledExclusionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get scheduled exclusions o k response has a 3xx status code
func (o *GetScheduledExclusionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scheduled exclusions o k response has a 4xx status code
func (o *GetScheduledExclusionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scheduled exclusions o k response has a 5xx status code
func (o *GetScheduledExclusionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get scheduled exclusions o k response a status code equal to that given
func (o *GetScheduledExclusionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get scheduled exclusions o k response
func (o *GetScheduledExclusionsOK) Code() int {
	return 200
}

func (o *GetScheduledExclusionsOK) Error() string {
	return fmt.Sprintf("[GET /filevantage/entities/policy-scheduled-exclusions/v1][%d] getScheduledExclusionsOK  %+v", 200, o.Payload)
}

func (o *GetScheduledExclusionsOK) String() string {
	return fmt.Sprintf("[GET /filevantage/entities/policy-scheduled-exclusions/v1][%d] getScheduledExclusionsOK  %+v", 200, o.Payload)
}

func (o *GetScheduledExclusionsOK) GetPayload() *models.ScheduledexclusionsResponse {
	return o.Payload
}

func (o *GetScheduledExclusionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetScheduledExclusionsBadRequest creates a GetScheduledExclusionsBadRequest with default headers values
func NewGetScheduledExclusionsBadRequest() *GetScheduledExclusionsBadRequest {
	return &GetScheduledExclusionsBadRequest{}
}

/*
GetScheduledExclusionsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetScheduledExclusionsBadRequest struct {

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

// IsSuccess returns true when this get scheduled exclusions bad request response has a 2xx status code
func (o *GetScheduledExclusionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scheduled exclusions bad request response has a 3xx status code
func (o *GetScheduledExclusionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scheduled exclusions bad request response has a 4xx status code
func (o *GetScheduledExclusionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scheduled exclusions bad request response has a 5xx status code
func (o *GetScheduledExclusionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get scheduled exclusions bad request response a status code equal to that given
func (o *GetScheduledExclusionsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get scheduled exclusions bad request response
func (o *GetScheduledExclusionsBadRequest) Code() int {
	return 400
}

func (o *GetScheduledExclusionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /filevantage/entities/policy-scheduled-exclusions/v1][%d] getScheduledExclusionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetScheduledExclusionsBadRequest) String() string {
	return fmt.Sprintf("[GET /filevantage/entities/policy-scheduled-exclusions/v1][%d] getScheduledExclusionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetScheduledExclusionsBadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *GetScheduledExclusionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetScheduledExclusionsForbidden creates a GetScheduledExclusionsForbidden with default headers values
func NewGetScheduledExclusionsForbidden() *GetScheduledExclusionsForbidden {
	return &GetScheduledExclusionsForbidden{}
}

/*
GetScheduledExclusionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetScheduledExclusionsForbidden struct {

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

// IsSuccess returns true when this get scheduled exclusions forbidden response has a 2xx status code
func (o *GetScheduledExclusionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scheduled exclusions forbidden response has a 3xx status code
func (o *GetScheduledExclusionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scheduled exclusions forbidden response has a 4xx status code
func (o *GetScheduledExclusionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scheduled exclusions forbidden response has a 5xx status code
func (o *GetScheduledExclusionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get scheduled exclusions forbidden response a status code equal to that given
func (o *GetScheduledExclusionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get scheduled exclusions forbidden response
func (o *GetScheduledExclusionsForbidden) Code() int {
	return 403
}

func (o *GetScheduledExclusionsForbidden) Error() string {
	return fmt.Sprintf("[GET /filevantage/entities/policy-scheduled-exclusions/v1][%d] getScheduledExclusionsForbidden  %+v", 403, o.Payload)
}

func (o *GetScheduledExclusionsForbidden) String() string {
	return fmt.Sprintf("[GET /filevantage/entities/policy-scheduled-exclusions/v1][%d] getScheduledExclusionsForbidden  %+v", 403, o.Payload)
}

func (o *GetScheduledExclusionsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetScheduledExclusionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetScheduledExclusionsNotFound creates a GetScheduledExclusionsNotFound with default headers values
func NewGetScheduledExclusionsNotFound() *GetScheduledExclusionsNotFound {
	return &GetScheduledExclusionsNotFound{}
}

/*
GetScheduledExclusionsNotFound describes a response with status code 404, with default header values.

The policy to retrieve the scheduled exclusions from does not exist.
*/
type GetScheduledExclusionsNotFound struct {

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

// IsSuccess returns true when this get scheduled exclusions not found response has a 2xx status code
func (o *GetScheduledExclusionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scheduled exclusions not found response has a 3xx status code
func (o *GetScheduledExclusionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scheduled exclusions not found response has a 4xx status code
func (o *GetScheduledExclusionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scheduled exclusions not found response has a 5xx status code
func (o *GetScheduledExclusionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get scheduled exclusions not found response a status code equal to that given
func (o *GetScheduledExclusionsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get scheduled exclusions not found response
func (o *GetScheduledExclusionsNotFound) Code() int {
	return 404
}

func (o *GetScheduledExclusionsNotFound) Error() string {
	return fmt.Sprintf("[GET /filevantage/entities/policy-scheduled-exclusions/v1][%d] getScheduledExclusionsNotFound  %+v", 404, o.Payload)
}

func (o *GetScheduledExclusionsNotFound) String() string {
	return fmt.Sprintf("[GET /filevantage/entities/policy-scheduled-exclusions/v1][%d] getScheduledExclusionsNotFound  %+v", 404, o.Payload)
}

func (o *GetScheduledExclusionsNotFound) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *GetScheduledExclusionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetScheduledExclusionsTooManyRequests creates a GetScheduledExclusionsTooManyRequests with default headers values
func NewGetScheduledExclusionsTooManyRequests() *GetScheduledExclusionsTooManyRequests {
	return &GetScheduledExclusionsTooManyRequests{}
}

/*
GetScheduledExclusionsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetScheduledExclusionsTooManyRequests struct {

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

// IsSuccess returns true when this get scheduled exclusions too many requests response has a 2xx status code
func (o *GetScheduledExclusionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scheduled exclusions too many requests response has a 3xx status code
func (o *GetScheduledExclusionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scheduled exclusions too many requests response has a 4xx status code
func (o *GetScheduledExclusionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scheduled exclusions too many requests response has a 5xx status code
func (o *GetScheduledExclusionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get scheduled exclusions too many requests response a status code equal to that given
func (o *GetScheduledExclusionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get scheduled exclusions too many requests response
func (o *GetScheduledExclusionsTooManyRequests) Code() int {
	return 429
}

func (o *GetScheduledExclusionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /filevantage/entities/policy-scheduled-exclusions/v1][%d] getScheduledExclusionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScheduledExclusionsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /filevantage/entities/policy-scheduled-exclusions/v1][%d] getScheduledExclusionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScheduledExclusionsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetScheduledExclusionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetScheduledExclusionsInternalServerError creates a GetScheduledExclusionsInternalServerError with default headers values
func NewGetScheduledExclusionsInternalServerError() *GetScheduledExclusionsInternalServerError {
	return &GetScheduledExclusionsInternalServerError{}
}

/*
GetScheduledExclusionsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetScheduledExclusionsInternalServerError struct {

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

// IsSuccess returns true when this get scheduled exclusions internal server error response has a 2xx status code
func (o *GetScheduledExclusionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scheduled exclusions internal server error response has a 3xx status code
func (o *GetScheduledExclusionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scheduled exclusions internal server error response has a 4xx status code
func (o *GetScheduledExclusionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scheduled exclusions internal server error response has a 5xx status code
func (o *GetScheduledExclusionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get scheduled exclusions internal server error response a status code equal to that given
func (o *GetScheduledExclusionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get scheduled exclusions internal server error response
func (o *GetScheduledExclusionsInternalServerError) Code() int {
	return 500
}

func (o *GetScheduledExclusionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /filevantage/entities/policy-scheduled-exclusions/v1][%d] getScheduledExclusionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScheduledExclusionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /filevantage/entities/policy-scheduled-exclusions/v1][%d] getScheduledExclusionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScheduledExclusionsInternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *GetScheduledExclusionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
