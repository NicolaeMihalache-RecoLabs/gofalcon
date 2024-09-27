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

// DeleteExportJobsV1Reader is a Reader for the DeleteExportJobsV1 structure.
type DeleteExportJobsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteExportJobsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteExportJobsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteExportJobsV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteExportJobsV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteExportJobsV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteExportJobsV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteExportJobsV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /recon/entities/exports/v1] DeleteExportJobsV1", response, response.Code())
	}
}

// NewDeleteExportJobsV1OK creates a DeleteExportJobsV1OK with default headers values
func NewDeleteExportJobsV1OK() *DeleteExportJobsV1OK {
	return &DeleteExportJobsV1OK{}
}

/*
DeleteExportJobsV1OK describes a response with status code 200, with default header values.

OK
*/
type DeleteExportJobsV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainExportJobIDResponseV1
}

// IsSuccess returns true when this delete export jobs v1 o k response has a 2xx status code
func (o *DeleteExportJobsV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete export jobs v1 o k response has a 3xx status code
func (o *DeleteExportJobsV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete export jobs v1 o k response has a 4xx status code
func (o *DeleteExportJobsV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete export jobs v1 o k response has a 5xx status code
func (o *DeleteExportJobsV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete export jobs v1 o k response a status code equal to that given
func (o *DeleteExportJobsV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete export jobs v1 o k response
func (o *DeleteExportJobsV1OK) Code() int {
	return 200
}

func (o *DeleteExportJobsV1OK) Error() string {
	return fmt.Sprintf("[DELETE /recon/entities/exports/v1][%d] deleteExportJobsV1OK  %+v", 200, o.Payload)
}

func (o *DeleteExportJobsV1OK) String() string {
	return fmt.Sprintf("[DELETE /recon/entities/exports/v1][%d] deleteExportJobsV1OK  %+v", 200, o.Payload)
}

func (o *DeleteExportJobsV1OK) GetPayload() *models.DomainExportJobIDResponseV1 {
	return o.Payload
}

func (o *DeleteExportJobsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainExportJobIDResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExportJobsV1BadRequest creates a DeleteExportJobsV1BadRequest with default headers values
func NewDeleteExportJobsV1BadRequest() *DeleteExportJobsV1BadRequest {
	return &DeleteExportJobsV1BadRequest{}
}

/*
DeleteExportJobsV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteExportJobsV1BadRequest struct {

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

// IsSuccess returns true when this delete export jobs v1 bad request response has a 2xx status code
func (o *DeleteExportJobsV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete export jobs v1 bad request response has a 3xx status code
func (o *DeleteExportJobsV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete export jobs v1 bad request response has a 4xx status code
func (o *DeleteExportJobsV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete export jobs v1 bad request response has a 5xx status code
func (o *DeleteExportJobsV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete export jobs v1 bad request response a status code equal to that given
func (o *DeleteExportJobsV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete export jobs v1 bad request response
func (o *DeleteExportJobsV1BadRequest) Code() int {
	return 400
}

func (o *DeleteExportJobsV1BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /recon/entities/exports/v1][%d] deleteExportJobsV1BadRequest  %+v", 400, o.Payload)
}

func (o *DeleteExportJobsV1BadRequest) String() string {
	return fmt.Sprintf("[DELETE /recon/entities/exports/v1][%d] deleteExportJobsV1BadRequest  %+v", 400, o.Payload)
}

func (o *DeleteExportJobsV1BadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *DeleteExportJobsV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteExportJobsV1Unauthorized creates a DeleteExportJobsV1Unauthorized with default headers values
func NewDeleteExportJobsV1Unauthorized() *DeleteExportJobsV1Unauthorized {
	return &DeleteExportJobsV1Unauthorized{}
}

/*
DeleteExportJobsV1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteExportJobsV1Unauthorized struct {

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

// IsSuccess returns true when this delete export jobs v1 unauthorized response has a 2xx status code
func (o *DeleteExportJobsV1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete export jobs v1 unauthorized response has a 3xx status code
func (o *DeleteExportJobsV1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete export jobs v1 unauthorized response has a 4xx status code
func (o *DeleteExportJobsV1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete export jobs v1 unauthorized response has a 5xx status code
func (o *DeleteExportJobsV1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete export jobs v1 unauthorized response a status code equal to that given
func (o *DeleteExportJobsV1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete export jobs v1 unauthorized response
func (o *DeleteExportJobsV1Unauthorized) Code() int {
	return 401
}

func (o *DeleteExportJobsV1Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /recon/entities/exports/v1][%d] deleteExportJobsV1Unauthorized  %+v", 401, o.Payload)
}

func (o *DeleteExportJobsV1Unauthorized) String() string {
	return fmt.Sprintf("[DELETE /recon/entities/exports/v1][%d] deleteExportJobsV1Unauthorized  %+v", 401, o.Payload)
}

func (o *DeleteExportJobsV1Unauthorized) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *DeleteExportJobsV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteExportJobsV1Forbidden creates a DeleteExportJobsV1Forbidden with default headers values
func NewDeleteExportJobsV1Forbidden() *DeleteExportJobsV1Forbidden {
	return &DeleteExportJobsV1Forbidden{}
}

/*
DeleteExportJobsV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteExportJobsV1Forbidden struct {

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

// IsSuccess returns true when this delete export jobs v1 forbidden response has a 2xx status code
func (o *DeleteExportJobsV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete export jobs v1 forbidden response has a 3xx status code
func (o *DeleteExportJobsV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete export jobs v1 forbidden response has a 4xx status code
func (o *DeleteExportJobsV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete export jobs v1 forbidden response has a 5xx status code
func (o *DeleteExportJobsV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete export jobs v1 forbidden response a status code equal to that given
func (o *DeleteExportJobsV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete export jobs v1 forbidden response
func (o *DeleteExportJobsV1Forbidden) Code() int {
	return 403
}

func (o *DeleteExportJobsV1Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /recon/entities/exports/v1][%d] deleteExportJobsV1Forbidden  %+v", 403, o.Payload)
}

func (o *DeleteExportJobsV1Forbidden) String() string {
	return fmt.Sprintf("[DELETE /recon/entities/exports/v1][%d] deleteExportJobsV1Forbidden  %+v", 403, o.Payload)
}

func (o *DeleteExportJobsV1Forbidden) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *DeleteExportJobsV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteExportJobsV1TooManyRequests creates a DeleteExportJobsV1TooManyRequests with default headers values
func NewDeleteExportJobsV1TooManyRequests() *DeleteExportJobsV1TooManyRequests {
	return &DeleteExportJobsV1TooManyRequests{}
}

/*
DeleteExportJobsV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteExportJobsV1TooManyRequests struct {

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

// IsSuccess returns true when this delete export jobs v1 too many requests response has a 2xx status code
func (o *DeleteExportJobsV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete export jobs v1 too many requests response has a 3xx status code
func (o *DeleteExportJobsV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete export jobs v1 too many requests response has a 4xx status code
func (o *DeleteExportJobsV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete export jobs v1 too many requests response has a 5xx status code
func (o *DeleteExportJobsV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete export jobs v1 too many requests response a status code equal to that given
func (o *DeleteExportJobsV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the delete export jobs v1 too many requests response
func (o *DeleteExportJobsV1TooManyRequests) Code() int {
	return 429
}

func (o *DeleteExportJobsV1TooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /recon/entities/exports/v1][%d] deleteExportJobsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteExportJobsV1TooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /recon/entities/exports/v1][%d] deleteExportJobsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteExportJobsV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteExportJobsV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteExportJobsV1InternalServerError creates a DeleteExportJobsV1InternalServerError with default headers values
func NewDeleteExportJobsV1InternalServerError() *DeleteExportJobsV1InternalServerError {
	return &DeleteExportJobsV1InternalServerError{}
}

/*
DeleteExportJobsV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteExportJobsV1InternalServerError struct {

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

// IsSuccess returns true when this delete export jobs v1 internal server error response has a 2xx status code
func (o *DeleteExportJobsV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete export jobs v1 internal server error response has a 3xx status code
func (o *DeleteExportJobsV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete export jobs v1 internal server error response has a 4xx status code
func (o *DeleteExportJobsV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete export jobs v1 internal server error response has a 5xx status code
func (o *DeleteExportJobsV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete export jobs v1 internal server error response a status code equal to that given
func (o *DeleteExportJobsV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete export jobs v1 internal server error response
func (o *DeleteExportJobsV1InternalServerError) Code() int {
	return 500
}

func (o *DeleteExportJobsV1InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /recon/entities/exports/v1][%d] deleteExportJobsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteExportJobsV1InternalServerError) String() string {
	return fmt.Sprintf("[DELETE /recon/entities/exports/v1][%d] deleteExportJobsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteExportJobsV1InternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *DeleteExportJobsV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
