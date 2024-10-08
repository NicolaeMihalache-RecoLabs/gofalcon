// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

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

// UpdateCSPMGCPServiceAccountsExtReader is a Reader for the UpdateCSPMGCPServiceAccountsExt structure.
type UpdateCSPMGCPServiceAccountsExtReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCSPMGCPServiceAccountsExtReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCSPMGCPServiceAccountsExtOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCSPMGCPServiceAccountsExtBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateCSPMGCPServiceAccountsExtForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateCSPMGCPServiceAccountsExtTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateCSPMGCPServiceAccountsExtInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /cloud-connect-cspm-gcp/entities/service-accounts/v1] UpdateCSPMGCPServiceAccountsExt", response, response.Code())
	}
}

// NewUpdateCSPMGCPServiceAccountsExtOK creates a UpdateCSPMGCPServiceAccountsExtOK with default headers values
func NewUpdateCSPMGCPServiceAccountsExtOK() *UpdateCSPMGCPServiceAccountsExtOK {
	return &UpdateCSPMGCPServiceAccountsExtOK{}
}

/*
UpdateCSPMGCPServiceAccountsExtOK describes a response with status code 200, with default header values.

OK
*/
type UpdateCSPMGCPServiceAccountsExtOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPServiceAccountResponseExtV1
}

// IsSuccess returns true when this update c s p m g c p service accounts ext o k response has a 2xx status code
func (o *UpdateCSPMGCPServiceAccountsExtOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update c s p m g c p service accounts ext o k response has a 3xx status code
func (o *UpdateCSPMGCPServiceAccountsExtOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update c s p m g c p service accounts ext o k response has a 4xx status code
func (o *UpdateCSPMGCPServiceAccountsExtOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update c s p m g c p service accounts ext o k response has a 5xx status code
func (o *UpdateCSPMGCPServiceAccountsExtOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update c s p m g c p service accounts ext o k response a status code equal to that given
func (o *UpdateCSPMGCPServiceAccountsExtOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update c s p m g c p service accounts ext o k response
func (o *UpdateCSPMGCPServiceAccountsExtOK) Code() int {
	return 200
}

func (o *UpdateCSPMGCPServiceAccountsExtOK) Error() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-gcp/entities/service-accounts/v1][%d] updateCSPMGCPServiceAccountsExtOK  %+v", 200, o.Payload)
}

func (o *UpdateCSPMGCPServiceAccountsExtOK) String() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-gcp/entities/service-accounts/v1][%d] updateCSPMGCPServiceAccountsExtOK  %+v", 200, o.Payload)
}

func (o *UpdateCSPMGCPServiceAccountsExtOK) GetPayload() *models.RegistrationGCPServiceAccountResponseExtV1 {
	return o.Payload
}

func (o *UpdateCSPMGCPServiceAccountsExtOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPServiceAccountResponseExtV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCSPMGCPServiceAccountsExtBadRequest creates a UpdateCSPMGCPServiceAccountsExtBadRequest with default headers values
func NewUpdateCSPMGCPServiceAccountsExtBadRequest() *UpdateCSPMGCPServiceAccountsExtBadRequest {
	return &UpdateCSPMGCPServiceAccountsExtBadRequest{}
}

/*
UpdateCSPMGCPServiceAccountsExtBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateCSPMGCPServiceAccountsExtBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPServiceAccountResponseExtV1
}

// IsSuccess returns true when this update c s p m g c p service accounts ext bad request response has a 2xx status code
func (o *UpdateCSPMGCPServiceAccountsExtBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update c s p m g c p service accounts ext bad request response has a 3xx status code
func (o *UpdateCSPMGCPServiceAccountsExtBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update c s p m g c p service accounts ext bad request response has a 4xx status code
func (o *UpdateCSPMGCPServiceAccountsExtBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update c s p m g c p service accounts ext bad request response has a 5xx status code
func (o *UpdateCSPMGCPServiceAccountsExtBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update c s p m g c p service accounts ext bad request response a status code equal to that given
func (o *UpdateCSPMGCPServiceAccountsExtBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update c s p m g c p service accounts ext bad request response
func (o *UpdateCSPMGCPServiceAccountsExtBadRequest) Code() int {
	return 400
}

func (o *UpdateCSPMGCPServiceAccountsExtBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-gcp/entities/service-accounts/v1][%d] updateCSPMGCPServiceAccountsExtBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateCSPMGCPServiceAccountsExtBadRequest) String() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-gcp/entities/service-accounts/v1][%d] updateCSPMGCPServiceAccountsExtBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateCSPMGCPServiceAccountsExtBadRequest) GetPayload() *models.RegistrationGCPServiceAccountResponseExtV1 {
	return o.Payload
}

func (o *UpdateCSPMGCPServiceAccountsExtBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPServiceAccountResponseExtV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCSPMGCPServiceAccountsExtForbidden creates a UpdateCSPMGCPServiceAccountsExtForbidden with default headers values
func NewUpdateCSPMGCPServiceAccountsExtForbidden() *UpdateCSPMGCPServiceAccountsExtForbidden {
	return &UpdateCSPMGCPServiceAccountsExtForbidden{}
}

/*
UpdateCSPMGCPServiceAccountsExtForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateCSPMGCPServiceAccountsExtForbidden struct {

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

// IsSuccess returns true when this update c s p m g c p service accounts ext forbidden response has a 2xx status code
func (o *UpdateCSPMGCPServiceAccountsExtForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update c s p m g c p service accounts ext forbidden response has a 3xx status code
func (o *UpdateCSPMGCPServiceAccountsExtForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update c s p m g c p service accounts ext forbidden response has a 4xx status code
func (o *UpdateCSPMGCPServiceAccountsExtForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update c s p m g c p service accounts ext forbidden response has a 5xx status code
func (o *UpdateCSPMGCPServiceAccountsExtForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update c s p m g c p service accounts ext forbidden response a status code equal to that given
func (o *UpdateCSPMGCPServiceAccountsExtForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update c s p m g c p service accounts ext forbidden response
func (o *UpdateCSPMGCPServiceAccountsExtForbidden) Code() int {
	return 403
}

func (o *UpdateCSPMGCPServiceAccountsExtForbidden) Error() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-gcp/entities/service-accounts/v1][%d] updateCSPMGCPServiceAccountsExtForbidden  %+v", 403, o.Payload)
}

func (o *UpdateCSPMGCPServiceAccountsExtForbidden) String() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-gcp/entities/service-accounts/v1][%d] updateCSPMGCPServiceAccountsExtForbidden  %+v", 403, o.Payload)
}

func (o *UpdateCSPMGCPServiceAccountsExtForbidden) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *UpdateCSPMGCPServiceAccountsExtForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateCSPMGCPServiceAccountsExtTooManyRequests creates a UpdateCSPMGCPServiceAccountsExtTooManyRequests with default headers values
func NewUpdateCSPMGCPServiceAccountsExtTooManyRequests() *UpdateCSPMGCPServiceAccountsExtTooManyRequests {
	return &UpdateCSPMGCPServiceAccountsExtTooManyRequests{}
}

/*
UpdateCSPMGCPServiceAccountsExtTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateCSPMGCPServiceAccountsExtTooManyRequests struct {

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

// IsSuccess returns true when this update c s p m g c p service accounts ext too many requests response has a 2xx status code
func (o *UpdateCSPMGCPServiceAccountsExtTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update c s p m g c p service accounts ext too many requests response has a 3xx status code
func (o *UpdateCSPMGCPServiceAccountsExtTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update c s p m g c p service accounts ext too many requests response has a 4xx status code
func (o *UpdateCSPMGCPServiceAccountsExtTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update c s p m g c p service accounts ext too many requests response has a 5xx status code
func (o *UpdateCSPMGCPServiceAccountsExtTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update c s p m g c p service accounts ext too many requests response a status code equal to that given
func (o *UpdateCSPMGCPServiceAccountsExtTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the update c s p m g c p service accounts ext too many requests response
func (o *UpdateCSPMGCPServiceAccountsExtTooManyRequests) Code() int {
	return 429
}

func (o *UpdateCSPMGCPServiceAccountsExtTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-gcp/entities/service-accounts/v1][%d] updateCSPMGCPServiceAccountsExtTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateCSPMGCPServiceAccountsExtTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-gcp/entities/service-accounts/v1][%d] updateCSPMGCPServiceAccountsExtTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateCSPMGCPServiceAccountsExtTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateCSPMGCPServiceAccountsExtTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateCSPMGCPServiceAccountsExtInternalServerError creates a UpdateCSPMGCPServiceAccountsExtInternalServerError with default headers values
func NewUpdateCSPMGCPServiceAccountsExtInternalServerError() *UpdateCSPMGCPServiceAccountsExtInternalServerError {
	return &UpdateCSPMGCPServiceAccountsExtInternalServerError{}
}

/*
UpdateCSPMGCPServiceAccountsExtInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateCSPMGCPServiceAccountsExtInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPServiceAccountResponseExtV1
}

// IsSuccess returns true when this update c s p m g c p service accounts ext internal server error response has a 2xx status code
func (o *UpdateCSPMGCPServiceAccountsExtInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update c s p m g c p service accounts ext internal server error response has a 3xx status code
func (o *UpdateCSPMGCPServiceAccountsExtInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update c s p m g c p service accounts ext internal server error response has a 4xx status code
func (o *UpdateCSPMGCPServiceAccountsExtInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update c s p m g c p service accounts ext internal server error response has a 5xx status code
func (o *UpdateCSPMGCPServiceAccountsExtInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update c s p m g c p service accounts ext internal server error response a status code equal to that given
func (o *UpdateCSPMGCPServiceAccountsExtInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update c s p m g c p service accounts ext internal server error response
func (o *UpdateCSPMGCPServiceAccountsExtInternalServerError) Code() int {
	return 500
}

func (o *UpdateCSPMGCPServiceAccountsExtInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-gcp/entities/service-accounts/v1][%d] updateCSPMGCPServiceAccountsExtInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateCSPMGCPServiceAccountsExtInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-gcp/entities/service-accounts/v1][%d] updateCSPMGCPServiceAccountsExtInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateCSPMGCPServiceAccountsExtInternalServerError) GetPayload() *models.RegistrationGCPServiceAccountResponseExtV1 {
	return o.Payload
}

func (o *UpdateCSPMGCPServiceAccountsExtInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPServiceAccountResponseExtV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
