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

// DeleteCSPMGCPAccountReader is a Reader for the DeleteCSPMGCPAccount structure.
type DeleteCSPMGCPAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCSPMGCPAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCSPMGCPAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewDeleteCSPMGCPAccountMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCSPMGCPAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteCSPMGCPAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteCSPMGCPAccountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteCSPMGCPAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /cloud-connect-cspm-gcp/entities/account/v1] DeleteCSPMGCPAccount", response, response.Code())
	}
}

// NewDeleteCSPMGCPAccountOK creates a DeleteCSPMGCPAccountOK with default headers values
func NewDeleteCSPMGCPAccountOK() *DeleteCSPMGCPAccountOK {
	return &DeleteCSPMGCPAccountOK{}
}

/*
DeleteCSPMGCPAccountOK describes a response with status code 200, with default header values.

OK
*/
type DeleteCSPMGCPAccountOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaBaseEntitiesResponse
}

// IsSuccess returns true when this delete c s p m g c p account o k response has a 2xx status code
func (o *DeleteCSPMGCPAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete c s p m g c p account o k response has a 3xx status code
func (o *DeleteCSPMGCPAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete c s p m g c p account o k response has a 4xx status code
func (o *DeleteCSPMGCPAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete c s p m g c p account o k response has a 5xx status code
func (o *DeleteCSPMGCPAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete c s p m g c p account o k response a status code equal to that given
func (o *DeleteCSPMGCPAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete c s p m g c p account o k response
func (o *DeleteCSPMGCPAccountOK) Code() int {
	return 200
}

func (o *DeleteCSPMGCPAccountOK) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-gcp/entities/account/v1][%d] deleteCSPMGCPAccountOK  %+v", 200, o.Payload)
}

func (o *DeleteCSPMGCPAccountOK) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-gcp/entities/account/v1][%d] deleteCSPMGCPAccountOK  %+v", 200, o.Payload)
}

func (o *DeleteCSPMGCPAccountOK) GetPayload() *models.MsaBaseEntitiesResponse {
	return o.Payload
}

func (o *DeleteCSPMGCPAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaBaseEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCSPMGCPAccountMultiStatus creates a DeleteCSPMGCPAccountMultiStatus with default headers values
func NewDeleteCSPMGCPAccountMultiStatus() *DeleteCSPMGCPAccountMultiStatus {
	return &DeleteCSPMGCPAccountMultiStatus{}
}

/*
DeleteCSPMGCPAccountMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type DeleteCSPMGCPAccountMultiStatus struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaBaseEntitiesResponse
}

// IsSuccess returns true when this delete c s p m g c p account multi status response has a 2xx status code
func (o *DeleteCSPMGCPAccountMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete c s p m g c p account multi status response has a 3xx status code
func (o *DeleteCSPMGCPAccountMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete c s p m g c p account multi status response has a 4xx status code
func (o *DeleteCSPMGCPAccountMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete c s p m g c p account multi status response has a 5xx status code
func (o *DeleteCSPMGCPAccountMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this delete c s p m g c p account multi status response a status code equal to that given
func (o *DeleteCSPMGCPAccountMultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the delete c s p m g c p account multi status response
func (o *DeleteCSPMGCPAccountMultiStatus) Code() int {
	return 207
}

func (o *DeleteCSPMGCPAccountMultiStatus) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-gcp/entities/account/v1][%d] deleteCSPMGCPAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *DeleteCSPMGCPAccountMultiStatus) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-gcp/entities/account/v1][%d] deleteCSPMGCPAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *DeleteCSPMGCPAccountMultiStatus) GetPayload() *models.MsaBaseEntitiesResponse {
	return o.Payload
}

func (o *DeleteCSPMGCPAccountMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaBaseEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCSPMGCPAccountBadRequest creates a DeleteCSPMGCPAccountBadRequest with default headers values
func NewDeleteCSPMGCPAccountBadRequest() *DeleteCSPMGCPAccountBadRequest {
	return &DeleteCSPMGCPAccountBadRequest{}
}

/*
DeleteCSPMGCPAccountBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteCSPMGCPAccountBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaBaseEntitiesResponse
}

// IsSuccess returns true when this delete c s p m g c p account bad request response has a 2xx status code
func (o *DeleteCSPMGCPAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete c s p m g c p account bad request response has a 3xx status code
func (o *DeleteCSPMGCPAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete c s p m g c p account bad request response has a 4xx status code
func (o *DeleteCSPMGCPAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete c s p m g c p account bad request response has a 5xx status code
func (o *DeleteCSPMGCPAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete c s p m g c p account bad request response a status code equal to that given
func (o *DeleteCSPMGCPAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete c s p m g c p account bad request response
func (o *DeleteCSPMGCPAccountBadRequest) Code() int {
	return 400
}

func (o *DeleteCSPMGCPAccountBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-gcp/entities/account/v1][%d] deleteCSPMGCPAccountBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteCSPMGCPAccountBadRequest) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-gcp/entities/account/v1][%d] deleteCSPMGCPAccountBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteCSPMGCPAccountBadRequest) GetPayload() *models.MsaBaseEntitiesResponse {
	return o.Payload
}

func (o *DeleteCSPMGCPAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaBaseEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCSPMGCPAccountForbidden creates a DeleteCSPMGCPAccountForbidden with default headers values
func NewDeleteCSPMGCPAccountForbidden() *DeleteCSPMGCPAccountForbidden {
	return &DeleteCSPMGCPAccountForbidden{}
}

/*
DeleteCSPMGCPAccountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteCSPMGCPAccountForbidden struct {

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

// IsSuccess returns true when this delete c s p m g c p account forbidden response has a 2xx status code
func (o *DeleteCSPMGCPAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete c s p m g c p account forbidden response has a 3xx status code
func (o *DeleteCSPMGCPAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete c s p m g c p account forbidden response has a 4xx status code
func (o *DeleteCSPMGCPAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete c s p m g c p account forbidden response has a 5xx status code
func (o *DeleteCSPMGCPAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete c s p m g c p account forbidden response a status code equal to that given
func (o *DeleteCSPMGCPAccountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete c s p m g c p account forbidden response
func (o *DeleteCSPMGCPAccountForbidden) Code() int {
	return 403
}

func (o *DeleteCSPMGCPAccountForbidden) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-gcp/entities/account/v1][%d] deleteCSPMGCPAccountForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCSPMGCPAccountForbidden) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-gcp/entities/account/v1][%d] deleteCSPMGCPAccountForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCSPMGCPAccountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteCSPMGCPAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteCSPMGCPAccountTooManyRequests creates a DeleteCSPMGCPAccountTooManyRequests with default headers values
func NewDeleteCSPMGCPAccountTooManyRequests() *DeleteCSPMGCPAccountTooManyRequests {
	return &DeleteCSPMGCPAccountTooManyRequests{}
}

/*
DeleteCSPMGCPAccountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteCSPMGCPAccountTooManyRequests struct {

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

// IsSuccess returns true when this delete c s p m g c p account too many requests response has a 2xx status code
func (o *DeleteCSPMGCPAccountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete c s p m g c p account too many requests response has a 3xx status code
func (o *DeleteCSPMGCPAccountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete c s p m g c p account too many requests response has a 4xx status code
func (o *DeleteCSPMGCPAccountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete c s p m g c p account too many requests response has a 5xx status code
func (o *DeleteCSPMGCPAccountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete c s p m g c p account too many requests response a status code equal to that given
func (o *DeleteCSPMGCPAccountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the delete c s p m g c p account too many requests response
func (o *DeleteCSPMGCPAccountTooManyRequests) Code() int {
	return 429
}

func (o *DeleteCSPMGCPAccountTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-gcp/entities/account/v1][%d] deleteCSPMGCPAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteCSPMGCPAccountTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-gcp/entities/account/v1][%d] deleteCSPMGCPAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteCSPMGCPAccountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteCSPMGCPAccountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteCSPMGCPAccountInternalServerError creates a DeleteCSPMGCPAccountInternalServerError with default headers values
func NewDeleteCSPMGCPAccountInternalServerError() *DeleteCSPMGCPAccountInternalServerError {
	return &DeleteCSPMGCPAccountInternalServerError{}
}

/*
DeleteCSPMGCPAccountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteCSPMGCPAccountInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaBaseEntitiesResponse
}

// IsSuccess returns true when this delete c s p m g c p account internal server error response has a 2xx status code
func (o *DeleteCSPMGCPAccountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete c s p m g c p account internal server error response has a 3xx status code
func (o *DeleteCSPMGCPAccountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete c s p m g c p account internal server error response has a 4xx status code
func (o *DeleteCSPMGCPAccountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete c s p m g c p account internal server error response has a 5xx status code
func (o *DeleteCSPMGCPAccountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete c s p m g c p account internal server error response a status code equal to that given
func (o *DeleteCSPMGCPAccountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete c s p m g c p account internal server error response
func (o *DeleteCSPMGCPAccountInternalServerError) Code() int {
	return 500
}

func (o *DeleteCSPMGCPAccountInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-gcp/entities/account/v1][%d] deleteCSPMGCPAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCSPMGCPAccountInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-gcp/entities/account/v1][%d] deleteCSPMGCPAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCSPMGCPAccountInternalServerError) GetPayload() *models.MsaBaseEntitiesResponse {
	return o.Payload
}

func (o *DeleteCSPMGCPAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaBaseEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
