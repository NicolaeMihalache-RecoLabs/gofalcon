// Code generated by go-swagger; DO NOT EDIT.

package sensor_download

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

// DownloadSensorInstallerByIDReader is a Reader for the DownloadSensorInstallerByID structure.
type DownloadSensorInstallerByIDReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *DownloadSensorInstallerByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDownloadSensorInstallerByIDOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDownloadSensorInstallerByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDownloadSensorInstallerByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDownloadSensorInstallerByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDownloadSensorInstallerByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /sensors/entities/download-installer/v1] DownloadSensorInstallerById", response, response.Code())
	}
}

// NewDownloadSensorInstallerByIDOK creates a DownloadSensorInstallerByIDOK with default headers values
func NewDownloadSensorInstallerByIDOK(writer io.Writer) *DownloadSensorInstallerByIDOK {
	return &DownloadSensorInstallerByIDOK{

		Payload: writer,
	}
}

/*
DownloadSensorInstallerByIDOK describes a response with status code 200, with default header values.

OK
*/
type DownloadSensorInstallerByIDOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload io.Writer
}

// IsSuccess returns true when this download sensor installer by Id o k response has a 2xx status code
func (o *DownloadSensorInstallerByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this download sensor installer by Id o k response has a 3xx status code
func (o *DownloadSensorInstallerByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download sensor installer by Id o k response has a 4xx status code
func (o *DownloadSensorInstallerByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this download sensor installer by Id o k response has a 5xx status code
func (o *DownloadSensorInstallerByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this download sensor installer by Id o k response a status code equal to that given
func (o *DownloadSensorInstallerByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the download sensor installer by Id o k response
func (o *DownloadSensorInstallerByIDOK) Code() int {
	return 200
}

func (o *DownloadSensorInstallerByIDOK) Error() string {
	return fmt.Sprintf("[GET /sensors/entities/download-installer/v1][%d] downloadSensorInstallerByIdOK  %+v", 200, o.Payload)
}

func (o *DownloadSensorInstallerByIDOK) String() string {
	return fmt.Sprintf("[GET /sensors/entities/download-installer/v1][%d] downloadSensorInstallerByIdOK  %+v", 200, o.Payload)
}

func (o *DownloadSensorInstallerByIDOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *DownloadSensorInstallerByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadSensorInstallerByIDBadRequest creates a DownloadSensorInstallerByIDBadRequest with default headers values
func NewDownloadSensorInstallerByIDBadRequest() *DownloadSensorInstallerByIDBadRequest {
	return &DownloadSensorInstallerByIDBadRequest{}
}

/*
DownloadSensorInstallerByIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DownloadSensorInstallerByIDBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecQueryResponse
}

// IsSuccess returns true when this download sensor installer by Id bad request response has a 2xx status code
func (o *DownloadSensorInstallerByIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this download sensor installer by Id bad request response has a 3xx status code
func (o *DownloadSensorInstallerByIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download sensor installer by Id bad request response has a 4xx status code
func (o *DownloadSensorInstallerByIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this download sensor installer by Id bad request response has a 5xx status code
func (o *DownloadSensorInstallerByIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this download sensor installer by Id bad request response a status code equal to that given
func (o *DownloadSensorInstallerByIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the download sensor installer by Id bad request response
func (o *DownloadSensorInstallerByIDBadRequest) Code() int {
	return 400
}

func (o *DownloadSensorInstallerByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /sensors/entities/download-installer/v1][%d] downloadSensorInstallerByIdBadRequest  %+v", 400, o.Payload)
}

func (o *DownloadSensorInstallerByIDBadRequest) String() string {
	return fmt.Sprintf("[GET /sensors/entities/download-installer/v1][%d] downloadSensorInstallerByIdBadRequest  %+v", 400, o.Payload)
}

func (o *DownloadSensorInstallerByIDBadRequest) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *DownloadSensorInstallerByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadSensorInstallerByIDForbidden creates a DownloadSensorInstallerByIDForbidden with default headers values
func NewDownloadSensorInstallerByIDForbidden() *DownloadSensorInstallerByIDForbidden {
	return &DownloadSensorInstallerByIDForbidden{}
}

/*
DownloadSensorInstallerByIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DownloadSensorInstallerByIDForbidden struct {

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

// IsSuccess returns true when this download sensor installer by Id forbidden response has a 2xx status code
func (o *DownloadSensorInstallerByIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this download sensor installer by Id forbidden response has a 3xx status code
func (o *DownloadSensorInstallerByIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download sensor installer by Id forbidden response has a 4xx status code
func (o *DownloadSensorInstallerByIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this download sensor installer by Id forbidden response has a 5xx status code
func (o *DownloadSensorInstallerByIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this download sensor installer by Id forbidden response a status code equal to that given
func (o *DownloadSensorInstallerByIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the download sensor installer by Id forbidden response
func (o *DownloadSensorInstallerByIDForbidden) Code() int {
	return 403
}

func (o *DownloadSensorInstallerByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /sensors/entities/download-installer/v1][%d] downloadSensorInstallerByIdForbidden  %+v", 403, o.Payload)
}

func (o *DownloadSensorInstallerByIDForbidden) String() string {
	return fmt.Sprintf("[GET /sensors/entities/download-installer/v1][%d] downloadSensorInstallerByIdForbidden  %+v", 403, o.Payload)
}

func (o *DownloadSensorInstallerByIDForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DownloadSensorInstallerByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDownloadSensorInstallerByIDNotFound creates a DownloadSensorInstallerByIDNotFound with default headers values
func NewDownloadSensorInstallerByIDNotFound() *DownloadSensorInstallerByIDNotFound {
	return &DownloadSensorInstallerByIDNotFound{}
}

/*
DownloadSensorInstallerByIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DownloadSensorInstallerByIDNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecQueryResponse
}

// IsSuccess returns true when this download sensor installer by Id not found response has a 2xx status code
func (o *DownloadSensorInstallerByIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this download sensor installer by Id not found response has a 3xx status code
func (o *DownloadSensorInstallerByIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download sensor installer by Id not found response has a 4xx status code
func (o *DownloadSensorInstallerByIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this download sensor installer by Id not found response has a 5xx status code
func (o *DownloadSensorInstallerByIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this download sensor installer by Id not found response a status code equal to that given
func (o *DownloadSensorInstallerByIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the download sensor installer by Id not found response
func (o *DownloadSensorInstallerByIDNotFound) Code() int {
	return 404
}

func (o *DownloadSensorInstallerByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /sensors/entities/download-installer/v1][%d] downloadSensorInstallerByIdNotFound  %+v", 404, o.Payload)
}

func (o *DownloadSensorInstallerByIDNotFound) String() string {
	return fmt.Sprintf("[GET /sensors/entities/download-installer/v1][%d] downloadSensorInstallerByIdNotFound  %+v", 404, o.Payload)
}

func (o *DownloadSensorInstallerByIDNotFound) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *DownloadSensorInstallerByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadSensorInstallerByIDTooManyRequests creates a DownloadSensorInstallerByIDTooManyRequests with default headers values
func NewDownloadSensorInstallerByIDTooManyRequests() *DownloadSensorInstallerByIDTooManyRequests {
	return &DownloadSensorInstallerByIDTooManyRequests{}
}

/*
DownloadSensorInstallerByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DownloadSensorInstallerByIDTooManyRequests struct {

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

// IsSuccess returns true when this download sensor installer by Id too many requests response has a 2xx status code
func (o *DownloadSensorInstallerByIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this download sensor installer by Id too many requests response has a 3xx status code
func (o *DownloadSensorInstallerByIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download sensor installer by Id too many requests response has a 4xx status code
func (o *DownloadSensorInstallerByIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this download sensor installer by Id too many requests response has a 5xx status code
func (o *DownloadSensorInstallerByIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this download sensor installer by Id too many requests response a status code equal to that given
func (o *DownloadSensorInstallerByIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the download sensor installer by Id too many requests response
func (o *DownloadSensorInstallerByIDTooManyRequests) Code() int {
	return 429
}

func (o *DownloadSensorInstallerByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /sensors/entities/download-installer/v1][%d] downloadSensorInstallerByIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DownloadSensorInstallerByIDTooManyRequests) String() string {
	return fmt.Sprintf("[GET /sensors/entities/download-installer/v1][%d] downloadSensorInstallerByIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DownloadSensorInstallerByIDTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DownloadSensorInstallerByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
