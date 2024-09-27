// Code generated by go-swagger; DO NOT EDIT.

package recon

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

// DeleteActionV1Reader is a Reader for the DeleteActionV1 structure.
type DeleteActionV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteActionV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteActionV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteActionV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteActionV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteActionV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteActionV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteActionV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /recon/entities/actions/v1] DeleteActionV1", response, response.Code())
	}
}

// NewDeleteActionV1OK creates a DeleteActionV1OK with default headers values
func NewDeleteActionV1OK() *DeleteActionV1OK {
	return &DeleteActionV1OK{}
}

/*
DeleteActionV1OK describes a response with status code 200, with default header values.

OK
*/
type DeleteActionV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainQueryResponse
}

// IsSuccess returns true when this delete action v1 o k response has a 2xx status code
func (o *DeleteActionV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete action v1 o k response has a 3xx status code
func (o *DeleteActionV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete action v1 o k response has a 4xx status code
func (o *DeleteActionV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete action v1 o k response has a 5xx status code
func (o *DeleteActionV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete action v1 o k response a status code equal to that given
func (o *DeleteActionV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete action v1 o k response
func (o *DeleteActionV1OK) Code() int {
	return 200
}

func (o *DeleteActionV1OK) Error() string {
	return fmt.Sprintf("[DELETE /recon/entities/actions/v1][%d] deleteActionV1OK  %+v", 200, o.Payload)
}

func (o *DeleteActionV1OK) String() string {
	return fmt.Sprintf("[DELETE /recon/entities/actions/v1][%d] deleteActionV1OK  %+v", 200, o.Payload)
}

func (o *DeleteActionV1OK) GetPayload() *models.DomainQueryResponse {
	return o.Payload
}

func (o *DeleteActionV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteActionV1BadRequest creates a DeleteActionV1BadRequest with default headers values
func NewDeleteActionV1BadRequest() *DeleteActionV1BadRequest {
	return &DeleteActionV1BadRequest{}
}

/*
DeleteActionV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteActionV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this delete action v1 bad request response has a 2xx status code
func (o *DeleteActionV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete action v1 bad request response has a 3xx status code
func (o *DeleteActionV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete action v1 bad request response has a 4xx status code
func (o *DeleteActionV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete action v1 bad request response has a 5xx status code
func (o *DeleteActionV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete action v1 bad request response a status code equal to that given
func (o *DeleteActionV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete action v1 bad request response
func (o *DeleteActionV1BadRequest) Code() int {
	return 400
}

func (o *DeleteActionV1BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /recon/entities/actions/v1][%d] deleteActionV1BadRequest  %+v", 400, o.Payload)
}

func (o *DeleteActionV1BadRequest) String() string {
	return fmt.Sprintf("[DELETE /recon/entities/actions/v1][%d] deleteActionV1BadRequest  %+v", 400, o.Payload)
}

func (o *DeleteActionV1BadRequest) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *DeleteActionV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteActionV1Unauthorized creates a DeleteActionV1Unauthorized with default headers values
func NewDeleteActionV1Unauthorized() *DeleteActionV1Unauthorized {
	return &DeleteActionV1Unauthorized{}
}

/*
DeleteActionV1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteActionV1Unauthorized struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this delete action v1 unauthorized response has a 2xx status code
func (o *DeleteActionV1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete action v1 unauthorized response has a 3xx status code
func (o *DeleteActionV1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete action v1 unauthorized response has a 4xx status code
func (o *DeleteActionV1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete action v1 unauthorized response has a 5xx status code
func (o *DeleteActionV1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete action v1 unauthorized response a status code equal to that given
func (o *DeleteActionV1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete action v1 unauthorized response
func (o *DeleteActionV1Unauthorized) Code() int {
	return 401
}

func (o *DeleteActionV1Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /recon/entities/actions/v1][%d] deleteActionV1Unauthorized  %+v", 401, o.Payload)
}

func (o *DeleteActionV1Unauthorized) String() string {
	return fmt.Sprintf("[DELETE /recon/entities/actions/v1][%d] deleteActionV1Unauthorized  %+v", 401, o.Payload)
}

func (o *DeleteActionV1Unauthorized) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *DeleteActionV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteActionV1Forbidden creates a DeleteActionV1Forbidden with default headers values
func NewDeleteActionV1Forbidden() *DeleteActionV1Forbidden {
	return &DeleteActionV1Forbidden{}
}

/*
DeleteActionV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteActionV1Forbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this delete action v1 forbidden response has a 2xx status code
func (o *DeleteActionV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete action v1 forbidden response has a 3xx status code
func (o *DeleteActionV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete action v1 forbidden response has a 4xx status code
func (o *DeleteActionV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete action v1 forbidden response has a 5xx status code
func (o *DeleteActionV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete action v1 forbidden response a status code equal to that given
func (o *DeleteActionV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete action v1 forbidden response
func (o *DeleteActionV1Forbidden) Code() int {
	return 403
}

func (o *DeleteActionV1Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /recon/entities/actions/v1][%d] deleteActionV1Forbidden  %+v", 403, o.Payload)
}

func (o *DeleteActionV1Forbidden) String() string {
	return fmt.Sprintf("[DELETE /recon/entities/actions/v1][%d] deleteActionV1Forbidden  %+v", 403, o.Payload)
}

func (o *DeleteActionV1Forbidden) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *DeleteActionV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteActionV1TooManyRequests creates a DeleteActionV1TooManyRequests with default headers values
func NewDeleteActionV1TooManyRequests() *DeleteActionV1TooManyRequests {
	return &DeleteActionV1TooManyRequests{}
}

/*
DeleteActionV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteActionV1TooManyRequests struct {

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

// IsSuccess returns true when this delete action v1 too many requests response has a 2xx status code
func (o *DeleteActionV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete action v1 too many requests response has a 3xx status code
func (o *DeleteActionV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete action v1 too many requests response has a 4xx status code
func (o *DeleteActionV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete action v1 too many requests response has a 5xx status code
func (o *DeleteActionV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete action v1 too many requests response a status code equal to that given
func (o *DeleteActionV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the delete action v1 too many requests response
func (o *DeleteActionV1TooManyRequests) Code() int {
	return 429
}

func (o *DeleteActionV1TooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /recon/entities/actions/v1][%d] deleteActionV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteActionV1TooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /recon/entities/actions/v1][%d] deleteActionV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteActionV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteActionV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteActionV1InternalServerError creates a DeleteActionV1InternalServerError with default headers values
func NewDeleteActionV1InternalServerError() *DeleteActionV1InternalServerError {
	return &DeleteActionV1InternalServerError{}
}

/*
DeleteActionV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteActionV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this delete action v1 internal server error response has a 2xx status code
func (o *DeleteActionV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete action v1 internal server error response has a 3xx status code
func (o *DeleteActionV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete action v1 internal server error response has a 4xx status code
func (o *DeleteActionV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete action v1 internal server error response has a 5xx status code
func (o *DeleteActionV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete action v1 internal server error response a status code equal to that given
func (o *DeleteActionV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete action v1 internal server error response
func (o *DeleteActionV1InternalServerError) Code() int {
	return 500
}

func (o *DeleteActionV1InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /recon/entities/actions/v1][%d] deleteActionV1InternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteActionV1InternalServerError) String() string {
	return fmt.Sprintf("[DELETE /recon/entities/actions/v1][%d] deleteActionV1InternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteActionV1InternalServerError) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *DeleteActionV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
