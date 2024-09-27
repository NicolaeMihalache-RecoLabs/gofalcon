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

// GetFileContentForExportJobsV1Reader is a Reader for the GetFileContentForExportJobsV1 structure.
type GetFileContentForExportJobsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFileContentForExportJobsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFileContentForExportJobsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFileContentForExportJobsV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFileContentForExportJobsV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFileContentForExportJobsV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFileContentForExportJobsV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFileContentForExportJobsV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /recon/entities/export-files/v1] GetFileContentForExportJobsV1", response, response.Code())
	}
}

// NewGetFileContentForExportJobsV1OK creates a GetFileContentForExportJobsV1OK with default headers values
func NewGetFileContentForExportJobsV1OK() *GetFileContentForExportJobsV1OK {
	return &GetFileContentForExportJobsV1OK{}
}

/*
GetFileContentForExportJobsV1OK describes a response with status code 200, with default header values.

Accepted
*/
type GetFileContentForExportJobsV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload []int64
}

// IsSuccess returns true when this get file content for export jobs v1 o k response has a 2xx status code
func (o *GetFileContentForExportJobsV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get file content for export jobs v1 o k response has a 3xx status code
func (o *GetFileContentForExportJobsV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get file content for export jobs v1 o k response has a 4xx status code
func (o *GetFileContentForExportJobsV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get file content for export jobs v1 o k response has a 5xx status code
func (o *GetFileContentForExportJobsV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get file content for export jobs v1 o k response a status code equal to that given
func (o *GetFileContentForExportJobsV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get file content for export jobs v1 o k response
func (o *GetFileContentForExportJobsV1OK) Code() int {
	return 200
}

func (o *GetFileContentForExportJobsV1OK) Error() string {
	return fmt.Sprintf("[GET /recon/entities/export-files/v1][%d] getFileContentForExportJobsV1OK  %+v", 200, o.Payload)
}

func (o *GetFileContentForExportJobsV1OK) String() string {
	return fmt.Sprintf("[GET /recon/entities/export-files/v1][%d] getFileContentForExportJobsV1OK  %+v", 200, o.Payload)
}

func (o *GetFileContentForExportJobsV1OK) GetPayload() []int64 {
	return o.Payload
}

func (o *GetFileContentForExportJobsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFileContentForExportJobsV1BadRequest creates a GetFileContentForExportJobsV1BadRequest with default headers values
func NewGetFileContentForExportJobsV1BadRequest() *GetFileContentForExportJobsV1BadRequest {
	return &GetFileContentForExportJobsV1BadRequest{}
}

/*
GetFileContentForExportJobsV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetFileContentForExportJobsV1BadRequest struct {

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

// IsSuccess returns true when this get file content for export jobs v1 bad request response has a 2xx status code
func (o *GetFileContentForExportJobsV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get file content for export jobs v1 bad request response has a 3xx status code
func (o *GetFileContentForExportJobsV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get file content for export jobs v1 bad request response has a 4xx status code
func (o *GetFileContentForExportJobsV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get file content for export jobs v1 bad request response has a 5xx status code
func (o *GetFileContentForExportJobsV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get file content for export jobs v1 bad request response a status code equal to that given
func (o *GetFileContentForExportJobsV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get file content for export jobs v1 bad request response
func (o *GetFileContentForExportJobsV1BadRequest) Code() int {
	return 400
}

func (o *GetFileContentForExportJobsV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /recon/entities/export-files/v1][%d] getFileContentForExportJobsV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetFileContentForExportJobsV1BadRequest) String() string {
	return fmt.Sprintf("[GET /recon/entities/export-files/v1][%d] getFileContentForExportJobsV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetFileContentForExportJobsV1BadRequest) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *GetFileContentForExportJobsV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFileContentForExportJobsV1Unauthorized creates a GetFileContentForExportJobsV1Unauthorized with default headers values
func NewGetFileContentForExportJobsV1Unauthorized() *GetFileContentForExportJobsV1Unauthorized {
	return &GetFileContentForExportJobsV1Unauthorized{}
}

/*
GetFileContentForExportJobsV1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetFileContentForExportJobsV1Unauthorized struct {

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

// IsSuccess returns true when this get file content for export jobs v1 unauthorized response has a 2xx status code
func (o *GetFileContentForExportJobsV1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get file content for export jobs v1 unauthorized response has a 3xx status code
func (o *GetFileContentForExportJobsV1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get file content for export jobs v1 unauthorized response has a 4xx status code
func (o *GetFileContentForExportJobsV1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get file content for export jobs v1 unauthorized response has a 5xx status code
func (o *GetFileContentForExportJobsV1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get file content for export jobs v1 unauthorized response a status code equal to that given
func (o *GetFileContentForExportJobsV1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get file content for export jobs v1 unauthorized response
func (o *GetFileContentForExportJobsV1Unauthorized) Code() int {
	return 401
}

func (o *GetFileContentForExportJobsV1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /recon/entities/export-files/v1][%d] getFileContentForExportJobsV1Unauthorized  %+v", 401, o.Payload)
}

func (o *GetFileContentForExportJobsV1Unauthorized) String() string {
	return fmt.Sprintf("[GET /recon/entities/export-files/v1][%d] getFileContentForExportJobsV1Unauthorized  %+v", 401, o.Payload)
}

func (o *GetFileContentForExportJobsV1Unauthorized) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *GetFileContentForExportJobsV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFileContentForExportJobsV1Forbidden creates a GetFileContentForExportJobsV1Forbidden with default headers values
func NewGetFileContentForExportJobsV1Forbidden() *GetFileContentForExportJobsV1Forbidden {
	return &GetFileContentForExportJobsV1Forbidden{}
}

/*
GetFileContentForExportJobsV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetFileContentForExportJobsV1Forbidden struct {

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

// IsSuccess returns true when this get file content for export jobs v1 forbidden response has a 2xx status code
func (o *GetFileContentForExportJobsV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get file content for export jobs v1 forbidden response has a 3xx status code
func (o *GetFileContentForExportJobsV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get file content for export jobs v1 forbidden response has a 4xx status code
func (o *GetFileContentForExportJobsV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get file content for export jobs v1 forbidden response has a 5xx status code
func (o *GetFileContentForExportJobsV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get file content for export jobs v1 forbidden response a status code equal to that given
func (o *GetFileContentForExportJobsV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get file content for export jobs v1 forbidden response
func (o *GetFileContentForExportJobsV1Forbidden) Code() int {
	return 403
}

func (o *GetFileContentForExportJobsV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /recon/entities/export-files/v1][%d] getFileContentForExportJobsV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetFileContentForExportJobsV1Forbidden) String() string {
	return fmt.Sprintf("[GET /recon/entities/export-files/v1][%d] getFileContentForExportJobsV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetFileContentForExportJobsV1Forbidden) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *GetFileContentForExportJobsV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFileContentForExportJobsV1TooManyRequests creates a GetFileContentForExportJobsV1TooManyRequests with default headers values
func NewGetFileContentForExportJobsV1TooManyRequests() *GetFileContentForExportJobsV1TooManyRequests {
	return &GetFileContentForExportJobsV1TooManyRequests{}
}

/*
GetFileContentForExportJobsV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetFileContentForExportJobsV1TooManyRequests struct {

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

// IsSuccess returns true when this get file content for export jobs v1 too many requests response has a 2xx status code
func (o *GetFileContentForExportJobsV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get file content for export jobs v1 too many requests response has a 3xx status code
func (o *GetFileContentForExportJobsV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get file content for export jobs v1 too many requests response has a 4xx status code
func (o *GetFileContentForExportJobsV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get file content for export jobs v1 too many requests response has a 5xx status code
func (o *GetFileContentForExportJobsV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get file content for export jobs v1 too many requests response a status code equal to that given
func (o *GetFileContentForExportJobsV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get file content for export jobs v1 too many requests response
func (o *GetFileContentForExportJobsV1TooManyRequests) Code() int {
	return 429
}

func (o *GetFileContentForExportJobsV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /recon/entities/export-files/v1][%d] getFileContentForExportJobsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFileContentForExportJobsV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /recon/entities/export-files/v1][%d] getFileContentForExportJobsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFileContentForExportJobsV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetFileContentForExportJobsV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFileContentForExportJobsV1InternalServerError creates a GetFileContentForExportJobsV1InternalServerError with default headers values
func NewGetFileContentForExportJobsV1InternalServerError() *GetFileContentForExportJobsV1InternalServerError {
	return &GetFileContentForExportJobsV1InternalServerError{}
}

/*
GetFileContentForExportJobsV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetFileContentForExportJobsV1InternalServerError struct {

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

// IsSuccess returns true when this get file content for export jobs v1 internal server error response has a 2xx status code
func (o *GetFileContentForExportJobsV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get file content for export jobs v1 internal server error response has a 3xx status code
func (o *GetFileContentForExportJobsV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get file content for export jobs v1 internal server error response has a 4xx status code
func (o *GetFileContentForExportJobsV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get file content for export jobs v1 internal server error response has a 5xx status code
func (o *GetFileContentForExportJobsV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get file content for export jobs v1 internal server error response a status code equal to that given
func (o *GetFileContentForExportJobsV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get file content for export jobs v1 internal server error response
func (o *GetFileContentForExportJobsV1InternalServerError) Code() int {
	return 500
}

func (o *GetFileContentForExportJobsV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /recon/entities/export-files/v1][%d] getFileContentForExportJobsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetFileContentForExportJobsV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /recon/entities/export-files/v1][%d] getFileContentForExportJobsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetFileContentForExportJobsV1InternalServerError) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *GetFileContentForExportJobsV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
