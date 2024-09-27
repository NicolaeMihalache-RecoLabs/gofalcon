// Code generated by go-swagger; DO NOT EDIT.

package message_center

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

// CaseAddAttachmentReader is a Reader for the CaseAddAttachment structure.
type CaseAddAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CaseAddAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCaseAddAttachmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCaseAddAttachmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCaseAddAttachmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCaseAddAttachmentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCaseAddAttachmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /message-center/entities/case-attachment/v1] CaseAddAttachment", response, response.Code())
	}
}

// NewCaseAddAttachmentOK creates a CaseAddAttachmentOK with default headers values
func NewCaseAddAttachmentOK() *CaseAddAttachmentOK {
	return &CaseAddAttachmentOK{}
}

/*
CaseAddAttachmentOK describes a response with status code 200, with default header values.

OK
*/
type CaseAddAttachmentOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIMessageCenterAttachmentUploadResponse
}

// IsSuccess returns true when this case add attachment o k response has a 2xx status code
func (o *CaseAddAttachmentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this case add attachment o k response has a 3xx status code
func (o *CaseAddAttachmentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this case add attachment o k response has a 4xx status code
func (o *CaseAddAttachmentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this case add attachment o k response has a 5xx status code
func (o *CaseAddAttachmentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this case add attachment o k response a status code equal to that given
func (o *CaseAddAttachmentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the case add attachment o k response
func (o *CaseAddAttachmentOK) Code() int {
	return 200
}

func (o *CaseAddAttachmentOK) Error() string {
	return fmt.Sprintf("[POST /message-center/entities/case-attachment/v1][%d] caseAddAttachmentOK  %+v", 200, o.Payload)
}

func (o *CaseAddAttachmentOK) String() string {
	return fmt.Sprintf("[POST /message-center/entities/case-attachment/v1][%d] caseAddAttachmentOK  %+v", 200, o.Payload)
}

func (o *CaseAddAttachmentOK) GetPayload() *models.APIMessageCenterAttachmentUploadResponse {
	return o.Payload
}

func (o *CaseAddAttachmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIMessageCenterAttachmentUploadResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCaseAddAttachmentBadRequest creates a CaseAddAttachmentBadRequest with default headers values
func NewCaseAddAttachmentBadRequest() *CaseAddAttachmentBadRequest {
	return &CaseAddAttachmentBadRequest{}
}

/*
CaseAddAttachmentBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CaseAddAttachmentBadRequest struct {

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

// IsSuccess returns true when this case add attachment bad request response has a 2xx status code
func (o *CaseAddAttachmentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this case add attachment bad request response has a 3xx status code
func (o *CaseAddAttachmentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this case add attachment bad request response has a 4xx status code
func (o *CaseAddAttachmentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this case add attachment bad request response has a 5xx status code
func (o *CaseAddAttachmentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this case add attachment bad request response a status code equal to that given
func (o *CaseAddAttachmentBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the case add attachment bad request response
func (o *CaseAddAttachmentBadRequest) Code() int {
	return 400
}

func (o *CaseAddAttachmentBadRequest) Error() string {
	return fmt.Sprintf("[POST /message-center/entities/case-attachment/v1][%d] caseAddAttachmentBadRequest  %+v", 400, o.Payload)
}

func (o *CaseAddAttachmentBadRequest) String() string {
	return fmt.Sprintf("[POST /message-center/entities/case-attachment/v1][%d] caseAddAttachmentBadRequest  %+v", 400, o.Payload)
}

func (o *CaseAddAttachmentBadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CaseAddAttachmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCaseAddAttachmentForbidden creates a CaseAddAttachmentForbidden with default headers values
func NewCaseAddAttachmentForbidden() *CaseAddAttachmentForbidden {
	return &CaseAddAttachmentForbidden{}
}

/*
CaseAddAttachmentForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CaseAddAttachmentForbidden struct {

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

// IsSuccess returns true when this case add attachment forbidden response has a 2xx status code
func (o *CaseAddAttachmentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this case add attachment forbidden response has a 3xx status code
func (o *CaseAddAttachmentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this case add attachment forbidden response has a 4xx status code
func (o *CaseAddAttachmentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this case add attachment forbidden response has a 5xx status code
func (o *CaseAddAttachmentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this case add attachment forbidden response a status code equal to that given
func (o *CaseAddAttachmentForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the case add attachment forbidden response
func (o *CaseAddAttachmentForbidden) Code() int {
	return 403
}

func (o *CaseAddAttachmentForbidden) Error() string {
	return fmt.Sprintf("[POST /message-center/entities/case-attachment/v1][%d] caseAddAttachmentForbidden  %+v", 403, o.Payload)
}

func (o *CaseAddAttachmentForbidden) String() string {
	return fmt.Sprintf("[POST /message-center/entities/case-attachment/v1][%d] caseAddAttachmentForbidden  %+v", 403, o.Payload)
}

func (o *CaseAddAttachmentForbidden) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CaseAddAttachmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCaseAddAttachmentTooManyRequests creates a CaseAddAttachmentTooManyRequests with default headers values
func NewCaseAddAttachmentTooManyRequests() *CaseAddAttachmentTooManyRequests {
	return &CaseAddAttachmentTooManyRequests{}
}

/*
CaseAddAttachmentTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CaseAddAttachmentTooManyRequests struct {

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

// IsSuccess returns true when this case add attachment too many requests response has a 2xx status code
func (o *CaseAddAttachmentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this case add attachment too many requests response has a 3xx status code
func (o *CaseAddAttachmentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this case add attachment too many requests response has a 4xx status code
func (o *CaseAddAttachmentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this case add attachment too many requests response has a 5xx status code
func (o *CaseAddAttachmentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this case add attachment too many requests response a status code equal to that given
func (o *CaseAddAttachmentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the case add attachment too many requests response
func (o *CaseAddAttachmentTooManyRequests) Code() int {
	return 429
}

func (o *CaseAddAttachmentTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /message-center/entities/case-attachment/v1][%d] caseAddAttachmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *CaseAddAttachmentTooManyRequests) String() string {
	return fmt.Sprintf("[POST /message-center/entities/case-attachment/v1][%d] caseAddAttachmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *CaseAddAttachmentTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CaseAddAttachmentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCaseAddAttachmentInternalServerError creates a CaseAddAttachmentInternalServerError with default headers values
func NewCaseAddAttachmentInternalServerError() *CaseAddAttachmentInternalServerError {
	return &CaseAddAttachmentInternalServerError{}
}

/*
CaseAddAttachmentInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CaseAddAttachmentInternalServerError struct {

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

// IsSuccess returns true when this case add attachment internal server error response has a 2xx status code
func (o *CaseAddAttachmentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this case add attachment internal server error response has a 3xx status code
func (o *CaseAddAttachmentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this case add attachment internal server error response has a 4xx status code
func (o *CaseAddAttachmentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this case add attachment internal server error response has a 5xx status code
func (o *CaseAddAttachmentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this case add attachment internal server error response a status code equal to that given
func (o *CaseAddAttachmentInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the case add attachment internal server error response
func (o *CaseAddAttachmentInternalServerError) Code() int {
	return 500
}

func (o *CaseAddAttachmentInternalServerError) Error() string {
	return fmt.Sprintf("[POST /message-center/entities/case-attachment/v1][%d] caseAddAttachmentInternalServerError  %+v", 500, o.Payload)
}

func (o *CaseAddAttachmentInternalServerError) String() string {
	return fmt.Sprintf("[POST /message-center/entities/case-attachment/v1][%d] caseAddAttachmentInternalServerError  %+v", 500, o.Payload)
}

func (o *CaseAddAttachmentInternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CaseAddAttachmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
