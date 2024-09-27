// Code generated by go-swagger; DO NOT EDIT.

package real_time_response_admin

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

// RTRListFalconScriptsReader is a Reader for the RTRListFalconScripts structure.
type RTRListFalconScriptsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRListFalconScriptsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRTRListFalconScriptsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRTRListFalconScriptsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRListFalconScriptsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRTRListFalconScriptsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRListFalconScriptsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /real-time-response/queries/falcon-scripts/v1] RTR-ListFalconScripts", response, response.Code())
	}
}

// NewRTRListFalconScriptsOK creates a RTRListFalconScriptsOK with default headers values
func NewRTRListFalconScriptsOK() *RTRListFalconScriptsOK {
	return &RTRListFalconScriptsOK{}
}

/*
RTRListFalconScriptsOK describes a response with status code 200, with default header values.

OK
*/
type RTRListFalconScriptsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.EmpowerapiMsaIDListResponse
}

// IsSuccess returns true when this r t r list falcon scripts o k response has a 2xx status code
func (o *RTRListFalconScriptsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this r t r list falcon scripts o k response has a 3xx status code
func (o *RTRListFalconScriptsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list falcon scripts o k response has a 4xx status code
func (o *RTRListFalconScriptsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r list falcon scripts o k response has a 5xx status code
func (o *RTRListFalconScriptsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list falcon scripts o k response a status code equal to that given
func (o *RTRListFalconScriptsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the r t r list falcon scripts o k response
func (o *RTRListFalconScriptsOK) Code() int {
	return 200
}

func (o *RTRListFalconScriptsOK) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/falcon-scripts/v1][%d] rTRListFalconScriptsOK  %+v", 200, o.Payload)
}

func (o *RTRListFalconScriptsOK) String() string {
	return fmt.Sprintf("[GET /real-time-response/queries/falcon-scripts/v1][%d] rTRListFalconScriptsOK  %+v", 200, o.Payload)
}

func (o *RTRListFalconScriptsOK) GetPayload() *models.EmpowerapiMsaIDListResponse {
	return o.Payload
}

func (o *RTRListFalconScriptsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.EmpowerapiMsaIDListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListFalconScriptsBadRequest creates a RTRListFalconScriptsBadRequest with default headers values
func NewRTRListFalconScriptsBadRequest() *RTRListFalconScriptsBadRequest {
	return &RTRListFalconScriptsBadRequest{}
}

/*
RTRListFalconScriptsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RTRListFalconScriptsBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this r t r list falcon scripts bad request response has a 2xx status code
func (o *RTRListFalconScriptsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r list falcon scripts bad request response has a 3xx status code
func (o *RTRListFalconScriptsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list falcon scripts bad request response has a 4xx status code
func (o *RTRListFalconScriptsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r list falcon scripts bad request response has a 5xx status code
func (o *RTRListFalconScriptsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list falcon scripts bad request response a status code equal to that given
func (o *RTRListFalconScriptsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the r t r list falcon scripts bad request response
func (o *RTRListFalconScriptsBadRequest) Code() int {
	return 400
}

func (o *RTRListFalconScriptsBadRequest) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/falcon-scripts/v1][%d] rTRListFalconScriptsBadRequest  %+v", 400, o.Payload)
}

func (o *RTRListFalconScriptsBadRequest) String() string {
	return fmt.Sprintf("[GET /real-time-response/queries/falcon-scripts/v1][%d] rTRListFalconScriptsBadRequest  %+v", 400, o.Payload)
}

func (o *RTRListFalconScriptsBadRequest) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRListFalconScriptsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListFalconScriptsForbidden creates a RTRListFalconScriptsForbidden with default headers values
func NewRTRListFalconScriptsForbidden() *RTRListFalconScriptsForbidden {
	return &RTRListFalconScriptsForbidden{}
}

/*
RTRListFalconScriptsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RTRListFalconScriptsForbidden struct {

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

// IsSuccess returns true when this r t r list falcon scripts forbidden response has a 2xx status code
func (o *RTRListFalconScriptsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r list falcon scripts forbidden response has a 3xx status code
func (o *RTRListFalconScriptsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list falcon scripts forbidden response has a 4xx status code
func (o *RTRListFalconScriptsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r list falcon scripts forbidden response has a 5xx status code
func (o *RTRListFalconScriptsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list falcon scripts forbidden response a status code equal to that given
func (o *RTRListFalconScriptsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the r t r list falcon scripts forbidden response
func (o *RTRListFalconScriptsForbidden) Code() int {
	return 403
}

func (o *RTRListFalconScriptsForbidden) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/falcon-scripts/v1][%d] rTRListFalconScriptsForbidden  %+v", 403, o.Payload)
}

func (o *RTRListFalconScriptsForbidden) String() string {
	return fmt.Sprintf("[GET /real-time-response/queries/falcon-scripts/v1][%d] rTRListFalconScriptsForbidden  %+v", 403, o.Payload)
}

func (o *RTRListFalconScriptsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRListFalconScriptsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRListFalconScriptsNotFound creates a RTRListFalconScriptsNotFound with default headers values
func NewRTRListFalconScriptsNotFound() *RTRListFalconScriptsNotFound {
	return &RTRListFalconScriptsNotFound{}
}

/*
RTRListFalconScriptsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RTRListFalconScriptsNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this r t r list falcon scripts not found response has a 2xx status code
func (o *RTRListFalconScriptsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r list falcon scripts not found response has a 3xx status code
func (o *RTRListFalconScriptsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list falcon scripts not found response has a 4xx status code
func (o *RTRListFalconScriptsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r list falcon scripts not found response has a 5xx status code
func (o *RTRListFalconScriptsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list falcon scripts not found response a status code equal to that given
func (o *RTRListFalconScriptsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the r t r list falcon scripts not found response
func (o *RTRListFalconScriptsNotFound) Code() int {
	return 404
}

func (o *RTRListFalconScriptsNotFound) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/falcon-scripts/v1][%d] rTRListFalconScriptsNotFound  %+v", 404, o.Payload)
}

func (o *RTRListFalconScriptsNotFound) String() string {
	return fmt.Sprintf("[GET /real-time-response/queries/falcon-scripts/v1][%d] rTRListFalconScriptsNotFound  %+v", 404, o.Payload)
}

func (o *RTRListFalconScriptsNotFound) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRListFalconScriptsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListFalconScriptsTooManyRequests creates a RTRListFalconScriptsTooManyRequests with default headers values
func NewRTRListFalconScriptsTooManyRequests() *RTRListFalconScriptsTooManyRequests {
	return &RTRListFalconScriptsTooManyRequests{}
}

/*
RTRListFalconScriptsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RTRListFalconScriptsTooManyRequests struct {

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

// IsSuccess returns true when this r t r list falcon scripts too many requests response has a 2xx status code
func (o *RTRListFalconScriptsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r list falcon scripts too many requests response has a 3xx status code
func (o *RTRListFalconScriptsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list falcon scripts too many requests response has a 4xx status code
func (o *RTRListFalconScriptsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r list falcon scripts too many requests response has a 5xx status code
func (o *RTRListFalconScriptsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list falcon scripts too many requests response a status code equal to that given
func (o *RTRListFalconScriptsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the r t r list falcon scripts too many requests response
func (o *RTRListFalconScriptsTooManyRequests) Code() int {
	return 429
}

func (o *RTRListFalconScriptsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/falcon-scripts/v1][%d] rTRListFalconScriptsTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRListFalconScriptsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /real-time-response/queries/falcon-scripts/v1][%d] rTRListFalconScriptsTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRListFalconScriptsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRListFalconScriptsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
