// Code generated by go-swagger; DO NOT EDIT.

package intel

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

// GetIntelRuleFileReader is a Reader for the GetIntelRuleFile structure.
type GetIntelRuleFileReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetIntelRuleFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntelRuleFileOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntelRuleFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntelRuleFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntelRuleFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntelRuleFileTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntelRuleFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /intel/entities/rules-files/v1] GetIntelRuleFile", response, response.Code())
	}
}

// NewGetIntelRuleFileOK creates a GetIntelRuleFileOK with default headers values
func NewGetIntelRuleFileOK(writer io.Writer) *GetIntelRuleFileOK {
	return &GetIntelRuleFileOK{

		Payload: writer,
	}
}

/*
GetIntelRuleFileOK describes a response with status code 200, with default header values.

OK
*/
type GetIntelRuleFileOK struct {

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

// IsSuccess returns true when this get intel rule file o k response has a 2xx status code
func (o *GetIntelRuleFileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get intel rule file o k response has a 3xx status code
func (o *GetIntelRuleFileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get intel rule file o k response has a 4xx status code
func (o *GetIntelRuleFileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get intel rule file o k response has a 5xx status code
func (o *GetIntelRuleFileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get intel rule file o k response a status code equal to that given
func (o *GetIntelRuleFileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get intel rule file o k response
func (o *GetIntelRuleFileOK) Code() int {
	return 200
}

func (o *GetIntelRuleFileOK) Error() string {
	return fmt.Sprintf("[GET /intel/entities/rules-files/v1][%d] getIntelRuleFileOK  %+v", 200, o.Payload)
}

func (o *GetIntelRuleFileOK) String() string {
	return fmt.Sprintf("[GET /intel/entities/rules-files/v1][%d] getIntelRuleFileOK  %+v", 200, o.Payload)
}

func (o *GetIntelRuleFileOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *GetIntelRuleFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIntelRuleFileBadRequest creates a GetIntelRuleFileBadRequest with default headers values
func NewGetIntelRuleFileBadRequest() *GetIntelRuleFileBadRequest {
	return &GetIntelRuleFileBadRequest{}
}

/*
GetIntelRuleFileBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetIntelRuleFileBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this get intel rule file bad request response has a 2xx status code
func (o *GetIntelRuleFileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get intel rule file bad request response has a 3xx status code
func (o *GetIntelRuleFileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get intel rule file bad request response has a 4xx status code
func (o *GetIntelRuleFileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get intel rule file bad request response has a 5xx status code
func (o *GetIntelRuleFileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get intel rule file bad request response a status code equal to that given
func (o *GetIntelRuleFileBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get intel rule file bad request response
func (o *GetIntelRuleFileBadRequest) Code() int {
	return 400
}

func (o *GetIntelRuleFileBadRequest) Error() string {
	return fmt.Sprintf("[GET /intel/entities/rules-files/v1][%d] getIntelRuleFileBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntelRuleFileBadRequest) String() string {
	return fmt.Sprintf("[GET /intel/entities/rules-files/v1][%d] getIntelRuleFileBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntelRuleFileBadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetIntelRuleFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntelRuleFileForbidden creates a GetIntelRuleFileForbidden with default headers values
func NewGetIntelRuleFileForbidden() *GetIntelRuleFileForbidden {
	return &GetIntelRuleFileForbidden{}
}

/*
GetIntelRuleFileForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetIntelRuleFileForbidden struct {

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

// IsSuccess returns true when this get intel rule file forbidden response has a 2xx status code
func (o *GetIntelRuleFileForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get intel rule file forbidden response has a 3xx status code
func (o *GetIntelRuleFileForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get intel rule file forbidden response has a 4xx status code
func (o *GetIntelRuleFileForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get intel rule file forbidden response has a 5xx status code
func (o *GetIntelRuleFileForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get intel rule file forbidden response a status code equal to that given
func (o *GetIntelRuleFileForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get intel rule file forbidden response
func (o *GetIntelRuleFileForbidden) Code() int {
	return 403
}

func (o *GetIntelRuleFileForbidden) Error() string {
	return fmt.Sprintf("[GET /intel/entities/rules-files/v1][%d] getIntelRuleFileForbidden  %+v", 403, o.Payload)
}

func (o *GetIntelRuleFileForbidden) String() string {
	return fmt.Sprintf("[GET /intel/entities/rules-files/v1][%d] getIntelRuleFileForbidden  %+v", 403, o.Payload)
}

func (o *GetIntelRuleFileForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetIntelRuleFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIntelRuleFileNotFound creates a GetIntelRuleFileNotFound with default headers values
func NewGetIntelRuleFileNotFound() *GetIntelRuleFileNotFound {
	return &GetIntelRuleFileNotFound{}
}

/*
GetIntelRuleFileNotFound describes a response with status code 404, with default header values.

Bad Request
*/
type GetIntelRuleFileNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this get intel rule file not found response has a 2xx status code
func (o *GetIntelRuleFileNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get intel rule file not found response has a 3xx status code
func (o *GetIntelRuleFileNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get intel rule file not found response has a 4xx status code
func (o *GetIntelRuleFileNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get intel rule file not found response has a 5xx status code
func (o *GetIntelRuleFileNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get intel rule file not found response a status code equal to that given
func (o *GetIntelRuleFileNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get intel rule file not found response
func (o *GetIntelRuleFileNotFound) Code() int {
	return 404
}

func (o *GetIntelRuleFileNotFound) Error() string {
	return fmt.Sprintf("[GET /intel/entities/rules-files/v1][%d] getIntelRuleFileNotFound  %+v", 404, o.Payload)
}

func (o *GetIntelRuleFileNotFound) String() string {
	return fmt.Sprintf("[GET /intel/entities/rules-files/v1][%d] getIntelRuleFileNotFound  %+v", 404, o.Payload)
}

func (o *GetIntelRuleFileNotFound) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetIntelRuleFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntelRuleFileTooManyRequests creates a GetIntelRuleFileTooManyRequests with default headers values
func NewGetIntelRuleFileTooManyRequests() *GetIntelRuleFileTooManyRequests {
	return &GetIntelRuleFileTooManyRequests{}
}

/*
GetIntelRuleFileTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetIntelRuleFileTooManyRequests struct {

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

// IsSuccess returns true when this get intel rule file too many requests response has a 2xx status code
func (o *GetIntelRuleFileTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get intel rule file too many requests response has a 3xx status code
func (o *GetIntelRuleFileTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get intel rule file too many requests response has a 4xx status code
func (o *GetIntelRuleFileTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get intel rule file too many requests response has a 5xx status code
func (o *GetIntelRuleFileTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get intel rule file too many requests response a status code equal to that given
func (o *GetIntelRuleFileTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get intel rule file too many requests response
func (o *GetIntelRuleFileTooManyRequests) Code() int {
	return 429
}

func (o *GetIntelRuleFileTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /intel/entities/rules-files/v1][%d] getIntelRuleFileTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntelRuleFileTooManyRequests) String() string {
	return fmt.Sprintf("[GET /intel/entities/rules-files/v1][%d] getIntelRuleFileTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntelRuleFileTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetIntelRuleFileTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIntelRuleFileInternalServerError creates a GetIntelRuleFileInternalServerError with default headers values
func NewGetIntelRuleFileInternalServerError() *GetIntelRuleFileInternalServerError {
	return &GetIntelRuleFileInternalServerError{}
}

/*
GetIntelRuleFileInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetIntelRuleFileInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this get intel rule file internal server error response has a 2xx status code
func (o *GetIntelRuleFileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get intel rule file internal server error response has a 3xx status code
func (o *GetIntelRuleFileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get intel rule file internal server error response has a 4xx status code
func (o *GetIntelRuleFileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get intel rule file internal server error response has a 5xx status code
func (o *GetIntelRuleFileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get intel rule file internal server error response a status code equal to that given
func (o *GetIntelRuleFileInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get intel rule file internal server error response
func (o *GetIntelRuleFileInternalServerError) Code() int {
	return 500
}

func (o *GetIntelRuleFileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /intel/entities/rules-files/v1][%d] getIntelRuleFileInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntelRuleFileInternalServerError) String() string {
	return fmt.Sprintf("[GET /intel/entities/rules-files/v1][%d] getIntelRuleFileInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntelRuleFileInternalServerError) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetIntelRuleFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
