// Code generated by go-swagger; DO NOT EDIT.

package d4c_registration

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

// GetD4CGcpAccountReader is a Reader for the GetD4CGcpAccount structure.
type GetD4CGcpAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetD4CGcpAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetD4CGcpAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewGetD4CGcpAccountMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetD4CGcpAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetD4CGcpAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetD4CGcpAccountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetD4CGcpAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cloud-connect-gcp/entities/account/v1] GetD4CGcpAccount", response, response.Code())
	}
}

// NewGetD4CGcpAccountOK creates a GetD4CGcpAccountOK with default headers values
func NewGetD4CGcpAccountOK() *GetD4CGcpAccountOK {
	return &GetD4CGcpAccountOK{}
}

/*
GetD4CGcpAccountOK describes a response with status code 200, with default header values.

OK
*/
type GetD4CGcpAccountOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountResponseV1
}

// IsSuccess returns true when this get d4 c gcp account o k response has a 2xx status code
func (o *GetD4CGcpAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get d4 c gcp account o k response has a 3xx status code
func (o *GetD4CGcpAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d4 c gcp account o k response has a 4xx status code
func (o *GetD4CGcpAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get d4 c gcp account o k response has a 5xx status code
func (o *GetD4CGcpAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get d4 c gcp account o k response a status code equal to that given
func (o *GetD4CGcpAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get d4 c gcp account o k response
func (o *GetD4CGcpAccountOK) Code() int {
	return 200
}

func (o *GetD4CGcpAccountOK) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getD4CGcpAccountOK  %+v", 200, o.Payload)
}

func (o *GetD4CGcpAccountOK) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getD4CGcpAccountOK  %+v", 200, o.Payload)
}

func (o *GetD4CGcpAccountOK) GetPayload() *models.RegistrationGCPAccountResponseV1 {
	return o.Payload
}

func (o *GetD4CGcpAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetD4CGcpAccountMultiStatus creates a GetD4CGcpAccountMultiStatus with default headers values
func NewGetD4CGcpAccountMultiStatus() *GetD4CGcpAccountMultiStatus {
	return &GetD4CGcpAccountMultiStatus{}
}

/*
GetD4CGcpAccountMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type GetD4CGcpAccountMultiStatus struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountResponseV1
}

// IsSuccess returns true when this get d4 c gcp account multi status response has a 2xx status code
func (o *GetD4CGcpAccountMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get d4 c gcp account multi status response has a 3xx status code
func (o *GetD4CGcpAccountMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d4 c gcp account multi status response has a 4xx status code
func (o *GetD4CGcpAccountMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this get d4 c gcp account multi status response has a 5xx status code
func (o *GetD4CGcpAccountMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this get d4 c gcp account multi status response a status code equal to that given
func (o *GetD4CGcpAccountMultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the get d4 c gcp account multi status response
func (o *GetD4CGcpAccountMultiStatus) Code() int {
	return 207
}

func (o *GetD4CGcpAccountMultiStatus) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getD4CGcpAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *GetD4CGcpAccountMultiStatus) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getD4CGcpAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *GetD4CGcpAccountMultiStatus) GetPayload() *models.RegistrationGCPAccountResponseV1 {
	return o.Payload
}

func (o *GetD4CGcpAccountMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetD4CGcpAccountBadRequest creates a GetD4CGcpAccountBadRequest with default headers values
func NewGetD4CGcpAccountBadRequest() *GetD4CGcpAccountBadRequest {
	return &GetD4CGcpAccountBadRequest{}
}

/*
GetD4CGcpAccountBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetD4CGcpAccountBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountResponseV1
}

// IsSuccess returns true when this get d4 c gcp account bad request response has a 2xx status code
func (o *GetD4CGcpAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get d4 c gcp account bad request response has a 3xx status code
func (o *GetD4CGcpAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d4 c gcp account bad request response has a 4xx status code
func (o *GetD4CGcpAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get d4 c gcp account bad request response has a 5xx status code
func (o *GetD4CGcpAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get d4 c gcp account bad request response a status code equal to that given
func (o *GetD4CGcpAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get d4 c gcp account bad request response
func (o *GetD4CGcpAccountBadRequest) Code() int {
	return 400
}

func (o *GetD4CGcpAccountBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getD4CGcpAccountBadRequest  %+v", 400, o.Payload)
}

func (o *GetD4CGcpAccountBadRequest) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getD4CGcpAccountBadRequest  %+v", 400, o.Payload)
}

func (o *GetD4CGcpAccountBadRequest) GetPayload() *models.RegistrationGCPAccountResponseV1 {
	return o.Payload
}

func (o *GetD4CGcpAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetD4CGcpAccountForbidden creates a GetD4CGcpAccountForbidden with default headers values
func NewGetD4CGcpAccountForbidden() *GetD4CGcpAccountForbidden {
	return &GetD4CGcpAccountForbidden{}
}

/*
GetD4CGcpAccountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetD4CGcpAccountForbidden struct {

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

// IsSuccess returns true when this get d4 c gcp account forbidden response has a 2xx status code
func (o *GetD4CGcpAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get d4 c gcp account forbidden response has a 3xx status code
func (o *GetD4CGcpAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d4 c gcp account forbidden response has a 4xx status code
func (o *GetD4CGcpAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get d4 c gcp account forbidden response has a 5xx status code
func (o *GetD4CGcpAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get d4 c gcp account forbidden response a status code equal to that given
func (o *GetD4CGcpAccountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get d4 c gcp account forbidden response
func (o *GetD4CGcpAccountForbidden) Code() int {
	return 403
}

func (o *GetD4CGcpAccountForbidden) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getD4CGcpAccountForbidden  %+v", 403, o.Payload)
}

func (o *GetD4CGcpAccountForbidden) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getD4CGcpAccountForbidden  %+v", 403, o.Payload)
}

func (o *GetD4CGcpAccountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetD4CGcpAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetD4CGcpAccountTooManyRequests creates a GetD4CGcpAccountTooManyRequests with default headers values
func NewGetD4CGcpAccountTooManyRequests() *GetD4CGcpAccountTooManyRequests {
	return &GetD4CGcpAccountTooManyRequests{}
}

/*
GetD4CGcpAccountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetD4CGcpAccountTooManyRequests struct {

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

// IsSuccess returns true when this get d4 c gcp account too many requests response has a 2xx status code
func (o *GetD4CGcpAccountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get d4 c gcp account too many requests response has a 3xx status code
func (o *GetD4CGcpAccountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d4 c gcp account too many requests response has a 4xx status code
func (o *GetD4CGcpAccountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get d4 c gcp account too many requests response has a 5xx status code
func (o *GetD4CGcpAccountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get d4 c gcp account too many requests response a status code equal to that given
func (o *GetD4CGcpAccountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get d4 c gcp account too many requests response
func (o *GetD4CGcpAccountTooManyRequests) Code() int {
	return 429
}

func (o *GetD4CGcpAccountTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getD4CGcpAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetD4CGcpAccountTooManyRequests) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getD4CGcpAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetD4CGcpAccountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetD4CGcpAccountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetD4CGcpAccountInternalServerError creates a GetD4CGcpAccountInternalServerError with default headers values
func NewGetD4CGcpAccountInternalServerError() *GetD4CGcpAccountInternalServerError {
	return &GetD4CGcpAccountInternalServerError{}
}

/*
GetD4CGcpAccountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetD4CGcpAccountInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountResponseV1
}

// IsSuccess returns true when this get d4 c gcp account internal server error response has a 2xx status code
func (o *GetD4CGcpAccountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get d4 c gcp account internal server error response has a 3xx status code
func (o *GetD4CGcpAccountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d4 c gcp account internal server error response has a 4xx status code
func (o *GetD4CGcpAccountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get d4 c gcp account internal server error response has a 5xx status code
func (o *GetD4CGcpAccountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get d4 c gcp account internal server error response a status code equal to that given
func (o *GetD4CGcpAccountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get d4 c gcp account internal server error response
func (o *GetD4CGcpAccountInternalServerError) Code() int {
	return 500
}

func (o *GetD4CGcpAccountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getD4CGcpAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *GetD4CGcpAccountInternalServerError) String() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/account/v1][%d] getD4CGcpAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *GetD4CGcpAccountInternalServerError) GetPayload() *models.RegistrationGCPAccountResponseV1 {
	return o.Payload
}

func (o *GetD4CGcpAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
