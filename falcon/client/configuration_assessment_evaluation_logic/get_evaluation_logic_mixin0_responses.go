// Code generated by go-swagger; DO NOT EDIT.

package configuration_assessment_evaluation_logic

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

// GetEvaluationLogicMixin0Reader is a Reader for the GetEvaluationLogicMixin0 structure.
type GetEvaluationLogicMixin0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEvaluationLogicMixin0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEvaluationLogicMixin0OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEvaluationLogicMixin0BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEvaluationLogicMixin0Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetEvaluationLogicMixin0TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEvaluationLogicMixin0InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /configuration-assessment/entities/evaluation-logic/v1] getEvaluationLogicMixin0", response, response.Code())
	}
}

// NewGetEvaluationLogicMixin0OK creates a GetEvaluationLogicMixin0OK with default headers values
func NewGetEvaluationLogicMixin0OK() *GetEvaluationLogicMixin0OK {
	return &GetEvaluationLogicMixin0OK{}
}

/*
GetEvaluationLogicMixin0OK describes a response with status code 200, with default header values.

OK
*/
type GetEvaluationLogicMixin0OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIEvaluationLogicEntitiesResponseV1
}

// IsSuccess returns true when this get evaluation logic mixin0 o k response has a 2xx status code
func (o *GetEvaluationLogicMixin0OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get evaluation logic mixin0 o k response has a 3xx status code
func (o *GetEvaluationLogicMixin0OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get evaluation logic mixin0 o k response has a 4xx status code
func (o *GetEvaluationLogicMixin0OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get evaluation logic mixin0 o k response has a 5xx status code
func (o *GetEvaluationLogicMixin0OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get evaluation logic mixin0 o k response a status code equal to that given
func (o *GetEvaluationLogicMixin0OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get evaluation logic mixin0 o k response
func (o *GetEvaluationLogicMixin0OK) Code() int {
	return 200
}

func (o *GetEvaluationLogicMixin0OK) Error() string {
	return fmt.Sprintf("[GET /configuration-assessment/entities/evaluation-logic/v1][%d] getEvaluationLogicMixin0OK  %+v", 200, o.Payload)
}

func (o *GetEvaluationLogicMixin0OK) String() string {
	return fmt.Sprintf("[GET /configuration-assessment/entities/evaluation-logic/v1][%d] getEvaluationLogicMixin0OK  %+v", 200, o.Payload)
}

func (o *GetEvaluationLogicMixin0OK) GetPayload() *models.DomainAPIEvaluationLogicEntitiesResponseV1 {
	return o.Payload
}

func (o *GetEvaluationLogicMixin0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIEvaluationLogicEntitiesResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEvaluationLogicMixin0BadRequest creates a GetEvaluationLogicMixin0BadRequest with default headers values
func NewGetEvaluationLogicMixin0BadRequest() *GetEvaluationLogicMixin0BadRequest {
	return &GetEvaluationLogicMixin0BadRequest{}
}

/*
GetEvaluationLogicMixin0BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetEvaluationLogicMixin0BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64
}

// IsSuccess returns true when this get evaluation logic mixin0 bad request response has a 2xx status code
func (o *GetEvaluationLogicMixin0BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get evaluation logic mixin0 bad request response has a 3xx status code
func (o *GetEvaluationLogicMixin0BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get evaluation logic mixin0 bad request response has a 4xx status code
func (o *GetEvaluationLogicMixin0BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get evaluation logic mixin0 bad request response has a 5xx status code
func (o *GetEvaluationLogicMixin0BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get evaluation logic mixin0 bad request response a status code equal to that given
func (o *GetEvaluationLogicMixin0BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get evaluation logic mixin0 bad request response
func (o *GetEvaluationLogicMixin0BadRequest) Code() int {
	return 400
}

func (o *GetEvaluationLogicMixin0BadRequest) Error() string {
	return fmt.Sprintf("[GET /configuration-assessment/entities/evaluation-logic/v1][%d] getEvaluationLogicMixin0BadRequest ", 400)
}

func (o *GetEvaluationLogicMixin0BadRequest) String() string {
	return fmt.Sprintf("[GET /configuration-assessment/entities/evaluation-logic/v1][%d] getEvaluationLogicMixin0BadRequest ", 400)
}

func (o *GetEvaluationLogicMixin0BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	return nil
}

// NewGetEvaluationLogicMixin0Forbidden creates a GetEvaluationLogicMixin0Forbidden with default headers values
func NewGetEvaluationLogicMixin0Forbidden() *GetEvaluationLogicMixin0Forbidden {
	return &GetEvaluationLogicMixin0Forbidden{}
}

/*
GetEvaluationLogicMixin0Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetEvaluationLogicMixin0Forbidden struct {

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

// IsSuccess returns true when this get evaluation logic mixin0 forbidden response has a 2xx status code
func (o *GetEvaluationLogicMixin0Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get evaluation logic mixin0 forbidden response has a 3xx status code
func (o *GetEvaluationLogicMixin0Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get evaluation logic mixin0 forbidden response has a 4xx status code
func (o *GetEvaluationLogicMixin0Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get evaluation logic mixin0 forbidden response has a 5xx status code
func (o *GetEvaluationLogicMixin0Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get evaluation logic mixin0 forbidden response a status code equal to that given
func (o *GetEvaluationLogicMixin0Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get evaluation logic mixin0 forbidden response
func (o *GetEvaluationLogicMixin0Forbidden) Code() int {
	return 403
}

func (o *GetEvaluationLogicMixin0Forbidden) Error() string {
	return fmt.Sprintf("[GET /configuration-assessment/entities/evaluation-logic/v1][%d] getEvaluationLogicMixin0Forbidden  %+v", 403, o.Payload)
}

func (o *GetEvaluationLogicMixin0Forbidden) String() string {
	return fmt.Sprintf("[GET /configuration-assessment/entities/evaluation-logic/v1][%d] getEvaluationLogicMixin0Forbidden  %+v", 403, o.Payload)
}

func (o *GetEvaluationLogicMixin0Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetEvaluationLogicMixin0Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetEvaluationLogicMixin0TooManyRequests creates a GetEvaluationLogicMixin0TooManyRequests with default headers values
func NewGetEvaluationLogicMixin0TooManyRequests() *GetEvaluationLogicMixin0TooManyRequests {
	return &GetEvaluationLogicMixin0TooManyRequests{}
}

/*
GetEvaluationLogicMixin0TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetEvaluationLogicMixin0TooManyRequests struct {

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

// IsSuccess returns true when this get evaluation logic mixin0 too many requests response has a 2xx status code
func (o *GetEvaluationLogicMixin0TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get evaluation logic mixin0 too many requests response has a 3xx status code
func (o *GetEvaluationLogicMixin0TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get evaluation logic mixin0 too many requests response has a 4xx status code
func (o *GetEvaluationLogicMixin0TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get evaluation logic mixin0 too many requests response has a 5xx status code
func (o *GetEvaluationLogicMixin0TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get evaluation logic mixin0 too many requests response a status code equal to that given
func (o *GetEvaluationLogicMixin0TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get evaluation logic mixin0 too many requests response
func (o *GetEvaluationLogicMixin0TooManyRequests) Code() int {
	return 429
}

func (o *GetEvaluationLogicMixin0TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /configuration-assessment/entities/evaluation-logic/v1][%d] getEvaluationLogicMixin0TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetEvaluationLogicMixin0TooManyRequests) String() string {
	return fmt.Sprintf("[GET /configuration-assessment/entities/evaluation-logic/v1][%d] getEvaluationLogicMixin0TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetEvaluationLogicMixin0TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetEvaluationLogicMixin0TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetEvaluationLogicMixin0InternalServerError creates a GetEvaluationLogicMixin0InternalServerError with default headers values
func NewGetEvaluationLogicMixin0InternalServerError() *GetEvaluationLogicMixin0InternalServerError {
	return &GetEvaluationLogicMixin0InternalServerError{}
}

/*
GetEvaluationLogicMixin0InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetEvaluationLogicMixin0InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64
}

// IsSuccess returns true when this get evaluation logic mixin0 internal server error response has a 2xx status code
func (o *GetEvaluationLogicMixin0InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get evaluation logic mixin0 internal server error response has a 3xx status code
func (o *GetEvaluationLogicMixin0InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get evaluation logic mixin0 internal server error response has a 4xx status code
func (o *GetEvaluationLogicMixin0InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get evaluation logic mixin0 internal server error response has a 5xx status code
func (o *GetEvaluationLogicMixin0InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get evaluation logic mixin0 internal server error response a status code equal to that given
func (o *GetEvaluationLogicMixin0InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get evaluation logic mixin0 internal server error response
func (o *GetEvaluationLogicMixin0InternalServerError) Code() int {
	return 500
}

func (o *GetEvaluationLogicMixin0InternalServerError) Error() string {
	return fmt.Sprintf("[GET /configuration-assessment/entities/evaluation-logic/v1][%d] getEvaluationLogicMixin0InternalServerError ", 500)
}

func (o *GetEvaluationLogicMixin0InternalServerError) String() string {
	return fmt.Sprintf("[GET /configuration-assessment/entities/evaluation-logic/v1][%d] getEvaluationLogicMixin0InternalServerError ", 500)
}

func (o *GetEvaluationLogicMixin0InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	return nil
}
