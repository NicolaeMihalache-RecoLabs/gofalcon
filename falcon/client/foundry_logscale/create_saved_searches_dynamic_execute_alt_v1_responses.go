// Code generated by go-swagger; DO NOT EDIT.

package foundry_logscale

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

// CreateSavedSearchesDynamicExecuteAltV1Reader is a Reader for the CreateSavedSearchesDynamicExecuteAltV1 structure.
type CreateSavedSearchesDynamicExecuteAltV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSavedSearchesDynamicExecuteAltV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSavedSearchesDynamicExecuteAltV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateSavedSearchesDynamicExecuteAltV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateSavedSearchesDynamicExecuteAltV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateSavedSearchesDynamicExecuteAltV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateSavedSearchesDynamicExecuteAltV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateSavedSearchesDynamicExecuteAltV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /loggingapi/entities/saved-searches-dynamic-execute/v1] CreateSavedSearchesDynamicExecuteAltV1", response, response.Code())
	}
}

// NewCreateSavedSearchesDynamicExecuteAltV1OK creates a CreateSavedSearchesDynamicExecuteAltV1OK with default headers values
func NewCreateSavedSearchesDynamicExecuteAltV1OK() *CreateSavedSearchesDynamicExecuteAltV1OK {
	return &CreateSavedSearchesDynamicExecuteAltV1OK{}
}

/*
CreateSavedSearchesDynamicExecuteAltV1OK describes a response with status code 200, with default header values.

OK
*/
type CreateSavedSearchesDynamicExecuteAltV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ApidomainQueryResponseWrapperV1
}

// IsSuccess returns true when this create saved searches dynamic execute alt v1 o k response has a 2xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create saved searches dynamic execute alt v1 o k response has a 3xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create saved searches dynamic execute alt v1 o k response has a 4xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create saved searches dynamic execute alt v1 o k response has a 5xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this create saved searches dynamic execute alt v1 o k response a status code equal to that given
func (o *CreateSavedSearchesDynamicExecuteAltV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create saved searches dynamic execute alt v1 o k response
func (o *CreateSavedSearchesDynamicExecuteAltV1OK) Code() int {
	return 200
}

func (o *CreateSavedSearchesDynamicExecuteAltV1OK) Error() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches-dynamic-execute/v1][%d] createSavedSearchesDynamicExecuteAltV1OK  %+v", 200, o.Payload)
}

func (o *CreateSavedSearchesDynamicExecuteAltV1OK) String() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches-dynamic-execute/v1][%d] createSavedSearchesDynamicExecuteAltV1OK  %+v", 200, o.Payload)
}

func (o *CreateSavedSearchesDynamicExecuteAltV1OK) GetPayload() *models.ApidomainQueryResponseWrapperV1 {
	return o.Payload
}

func (o *CreateSavedSearchesDynamicExecuteAltV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ApidomainQueryResponseWrapperV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSavedSearchesDynamicExecuteAltV1BadRequest creates a CreateSavedSearchesDynamicExecuteAltV1BadRequest with default headers values
func NewCreateSavedSearchesDynamicExecuteAltV1BadRequest() *CreateSavedSearchesDynamicExecuteAltV1BadRequest {
	return &CreateSavedSearchesDynamicExecuteAltV1BadRequest{}
}

/*
CreateSavedSearchesDynamicExecuteAltV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateSavedSearchesDynamicExecuteAltV1BadRequest struct {

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

// IsSuccess returns true when this create saved searches dynamic execute alt v1 bad request response has a 2xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create saved searches dynamic execute alt v1 bad request response has a 3xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create saved searches dynamic execute alt v1 bad request response has a 4xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create saved searches dynamic execute alt v1 bad request response has a 5xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create saved searches dynamic execute alt v1 bad request response a status code equal to that given
func (o *CreateSavedSearchesDynamicExecuteAltV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create saved searches dynamic execute alt v1 bad request response
func (o *CreateSavedSearchesDynamicExecuteAltV1BadRequest) Code() int {
	return 400
}

func (o *CreateSavedSearchesDynamicExecuteAltV1BadRequest) Error() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches-dynamic-execute/v1][%d] createSavedSearchesDynamicExecuteAltV1BadRequest  %+v", 400, o.Payload)
}

func (o *CreateSavedSearchesDynamicExecuteAltV1BadRequest) String() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches-dynamic-execute/v1][%d] createSavedSearchesDynamicExecuteAltV1BadRequest  %+v", 400, o.Payload)
}

func (o *CreateSavedSearchesDynamicExecuteAltV1BadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CreateSavedSearchesDynamicExecuteAltV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateSavedSearchesDynamicExecuteAltV1Forbidden creates a CreateSavedSearchesDynamicExecuteAltV1Forbidden with default headers values
func NewCreateSavedSearchesDynamicExecuteAltV1Forbidden() *CreateSavedSearchesDynamicExecuteAltV1Forbidden {
	return &CreateSavedSearchesDynamicExecuteAltV1Forbidden{}
}

/*
CreateSavedSearchesDynamicExecuteAltV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateSavedSearchesDynamicExecuteAltV1Forbidden struct {

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

// IsSuccess returns true when this create saved searches dynamic execute alt v1 forbidden response has a 2xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create saved searches dynamic execute alt v1 forbidden response has a 3xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create saved searches dynamic execute alt v1 forbidden response has a 4xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create saved searches dynamic execute alt v1 forbidden response has a 5xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create saved searches dynamic execute alt v1 forbidden response a status code equal to that given
func (o *CreateSavedSearchesDynamicExecuteAltV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create saved searches dynamic execute alt v1 forbidden response
func (o *CreateSavedSearchesDynamicExecuteAltV1Forbidden) Code() int {
	return 403
}

func (o *CreateSavedSearchesDynamicExecuteAltV1Forbidden) Error() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches-dynamic-execute/v1][%d] createSavedSearchesDynamicExecuteAltV1Forbidden  %+v", 403, o.Payload)
}

func (o *CreateSavedSearchesDynamicExecuteAltV1Forbidden) String() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches-dynamic-execute/v1][%d] createSavedSearchesDynamicExecuteAltV1Forbidden  %+v", 403, o.Payload)
}

func (o *CreateSavedSearchesDynamicExecuteAltV1Forbidden) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CreateSavedSearchesDynamicExecuteAltV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateSavedSearchesDynamicExecuteAltV1NotFound creates a CreateSavedSearchesDynamicExecuteAltV1NotFound with default headers values
func NewCreateSavedSearchesDynamicExecuteAltV1NotFound() *CreateSavedSearchesDynamicExecuteAltV1NotFound {
	return &CreateSavedSearchesDynamicExecuteAltV1NotFound{}
}

/*
CreateSavedSearchesDynamicExecuteAltV1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateSavedSearchesDynamicExecuteAltV1NotFound struct {

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

// IsSuccess returns true when this create saved searches dynamic execute alt v1 not found response has a 2xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create saved searches dynamic execute alt v1 not found response has a 3xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create saved searches dynamic execute alt v1 not found response has a 4xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create saved searches dynamic execute alt v1 not found response has a 5xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create saved searches dynamic execute alt v1 not found response a status code equal to that given
func (o *CreateSavedSearchesDynamicExecuteAltV1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create saved searches dynamic execute alt v1 not found response
func (o *CreateSavedSearchesDynamicExecuteAltV1NotFound) Code() int {
	return 404
}

func (o *CreateSavedSearchesDynamicExecuteAltV1NotFound) Error() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches-dynamic-execute/v1][%d] createSavedSearchesDynamicExecuteAltV1NotFound  %+v", 404, o.Payload)
}

func (o *CreateSavedSearchesDynamicExecuteAltV1NotFound) String() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches-dynamic-execute/v1][%d] createSavedSearchesDynamicExecuteAltV1NotFound  %+v", 404, o.Payload)
}

func (o *CreateSavedSearchesDynamicExecuteAltV1NotFound) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CreateSavedSearchesDynamicExecuteAltV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateSavedSearchesDynamicExecuteAltV1TooManyRequests creates a CreateSavedSearchesDynamicExecuteAltV1TooManyRequests with default headers values
func NewCreateSavedSearchesDynamicExecuteAltV1TooManyRequests() *CreateSavedSearchesDynamicExecuteAltV1TooManyRequests {
	return &CreateSavedSearchesDynamicExecuteAltV1TooManyRequests{}
}

/*
CreateSavedSearchesDynamicExecuteAltV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateSavedSearchesDynamicExecuteAltV1TooManyRequests struct {

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

// IsSuccess returns true when this create saved searches dynamic execute alt v1 too many requests response has a 2xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create saved searches dynamic execute alt v1 too many requests response has a 3xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create saved searches dynamic execute alt v1 too many requests response has a 4xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create saved searches dynamic execute alt v1 too many requests response has a 5xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create saved searches dynamic execute alt v1 too many requests response a status code equal to that given
func (o *CreateSavedSearchesDynamicExecuteAltV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the create saved searches dynamic execute alt v1 too many requests response
func (o *CreateSavedSearchesDynamicExecuteAltV1TooManyRequests) Code() int {
	return 429
}

func (o *CreateSavedSearchesDynamicExecuteAltV1TooManyRequests) Error() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches-dynamic-execute/v1][%d] createSavedSearchesDynamicExecuteAltV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateSavedSearchesDynamicExecuteAltV1TooManyRequests) String() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches-dynamic-execute/v1][%d] createSavedSearchesDynamicExecuteAltV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateSavedSearchesDynamicExecuteAltV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateSavedSearchesDynamicExecuteAltV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateSavedSearchesDynamicExecuteAltV1InternalServerError creates a CreateSavedSearchesDynamicExecuteAltV1InternalServerError with default headers values
func NewCreateSavedSearchesDynamicExecuteAltV1InternalServerError() *CreateSavedSearchesDynamicExecuteAltV1InternalServerError {
	return &CreateSavedSearchesDynamicExecuteAltV1InternalServerError{}
}

/*
CreateSavedSearchesDynamicExecuteAltV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateSavedSearchesDynamicExecuteAltV1InternalServerError struct {

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

// IsSuccess returns true when this create saved searches dynamic execute alt v1 internal server error response has a 2xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create saved searches dynamic execute alt v1 internal server error response has a 3xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create saved searches dynamic execute alt v1 internal server error response has a 4xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create saved searches dynamic execute alt v1 internal server error response has a 5xx status code
func (o *CreateSavedSearchesDynamicExecuteAltV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create saved searches dynamic execute alt v1 internal server error response a status code equal to that given
func (o *CreateSavedSearchesDynamicExecuteAltV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create saved searches dynamic execute alt v1 internal server error response
func (o *CreateSavedSearchesDynamicExecuteAltV1InternalServerError) Code() int {
	return 500
}

func (o *CreateSavedSearchesDynamicExecuteAltV1InternalServerError) Error() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches-dynamic-execute/v1][%d] createSavedSearchesDynamicExecuteAltV1InternalServerError  %+v", 500, o.Payload)
}

func (o *CreateSavedSearchesDynamicExecuteAltV1InternalServerError) String() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches-dynamic-execute/v1][%d] createSavedSearchesDynamicExecuteAltV1InternalServerError  %+v", 500, o.Payload)
}

func (o *CreateSavedSearchesDynamicExecuteAltV1InternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CreateSavedSearchesDynamicExecuteAltV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
