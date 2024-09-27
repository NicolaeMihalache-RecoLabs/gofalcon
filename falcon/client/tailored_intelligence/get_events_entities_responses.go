// Code generated by go-swagger; DO NOT EDIT.

package tailored_intelligence

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

// GetEventsEntitiesReader is a Reader for the GetEventsEntities structure.
type GetEventsEntitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventsEntitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEventsEntitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEventsEntitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEventsEntitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetEventsEntitiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEventsEntitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /ti/events/entities/events/GET/v2] GetEventsEntities", response, response.Code())
	}
}

// NewGetEventsEntitiesOK creates a GetEventsEntitiesOK with default headers values
func NewGetEventsEntitiesOK() *GetEventsEntitiesOK {
	return &GetEventsEntitiesOK{}
}

/*
GetEventsEntitiesOK describes a response with status code 200, with default header values.

OK
*/
type GetEventsEntitiesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainEventEntitiesResponse
}

// IsSuccess returns true when this get events entities o k response has a 2xx status code
func (o *GetEventsEntitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get events entities o k response has a 3xx status code
func (o *GetEventsEntitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get events entities o k response has a 4xx status code
func (o *GetEventsEntitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get events entities o k response has a 5xx status code
func (o *GetEventsEntitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get events entities o k response a status code equal to that given
func (o *GetEventsEntitiesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get events entities o k response
func (o *GetEventsEntitiesOK) Code() int {
	return 200
}

func (o *GetEventsEntitiesOK) Error() string {
	return fmt.Sprintf("[POST /ti/events/entities/events/GET/v2][%d] getEventsEntitiesOK  %+v", 200, o.Payload)
}

func (o *GetEventsEntitiesOK) String() string {
	return fmt.Sprintf("[POST /ti/events/entities/events/GET/v2][%d] getEventsEntitiesOK  %+v", 200, o.Payload)
}

func (o *GetEventsEntitiesOK) GetPayload() *models.DomainEventEntitiesResponse {
	return o.Payload
}

func (o *GetEventsEntitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainEventEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventsEntitiesBadRequest creates a GetEventsEntitiesBadRequest with default headers values
func NewGetEventsEntitiesBadRequest() *GetEventsEntitiesBadRequest {
	return &GetEventsEntitiesBadRequest{}
}

/*
GetEventsEntitiesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetEventsEntitiesBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainEventEntitiesResponse
}

// IsSuccess returns true when this get events entities bad request response has a 2xx status code
func (o *GetEventsEntitiesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get events entities bad request response has a 3xx status code
func (o *GetEventsEntitiesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get events entities bad request response has a 4xx status code
func (o *GetEventsEntitiesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get events entities bad request response has a 5xx status code
func (o *GetEventsEntitiesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get events entities bad request response a status code equal to that given
func (o *GetEventsEntitiesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get events entities bad request response
func (o *GetEventsEntitiesBadRequest) Code() int {
	return 400
}

func (o *GetEventsEntitiesBadRequest) Error() string {
	return fmt.Sprintf("[POST /ti/events/entities/events/GET/v2][%d] getEventsEntitiesBadRequest  %+v", 400, o.Payload)
}

func (o *GetEventsEntitiesBadRequest) String() string {
	return fmt.Sprintf("[POST /ti/events/entities/events/GET/v2][%d] getEventsEntitiesBadRequest  %+v", 400, o.Payload)
}

func (o *GetEventsEntitiesBadRequest) GetPayload() *models.DomainEventEntitiesResponse {
	return o.Payload
}

func (o *GetEventsEntitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainEventEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventsEntitiesForbidden creates a GetEventsEntitiesForbidden with default headers values
func NewGetEventsEntitiesForbidden() *GetEventsEntitiesForbidden {
	return &GetEventsEntitiesForbidden{}
}

/*
GetEventsEntitiesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetEventsEntitiesForbidden struct {

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

// IsSuccess returns true when this get events entities forbidden response has a 2xx status code
func (o *GetEventsEntitiesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get events entities forbidden response has a 3xx status code
func (o *GetEventsEntitiesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get events entities forbidden response has a 4xx status code
func (o *GetEventsEntitiesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get events entities forbidden response has a 5xx status code
func (o *GetEventsEntitiesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get events entities forbidden response a status code equal to that given
func (o *GetEventsEntitiesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get events entities forbidden response
func (o *GetEventsEntitiesForbidden) Code() int {
	return 403
}

func (o *GetEventsEntitiesForbidden) Error() string {
	return fmt.Sprintf("[POST /ti/events/entities/events/GET/v2][%d] getEventsEntitiesForbidden ", 403)
}

func (o *GetEventsEntitiesForbidden) String() string {
	return fmt.Sprintf("[POST /ti/events/entities/events/GET/v2][%d] getEventsEntitiesForbidden ", 403)
}

func (o *GetEventsEntitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetEventsEntitiesTooManyRequests creates a GetEventsEntitiesTooManyRequests with default headers values
func NewGetEventsEntitiesTooManyRequests() *GetEventsEntitiesTooManyRequests {
	return &GetEventsEntitiesTooManyRequests{}
}

/*
GetEventsEntitiesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetEventsEntitiesTooManyRequests struct {

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

// IsSuccess returns true when this get events entities too many requests response has a 2xx status code
func (o *GetEventsEntitiesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get events entities too many requests response has a 3xx status code
func (o *GetEventsEntitiesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get events entities too many requests response has a 4xx status code
func (o *GetEventsEntitiesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get events entities too many requests response has a 5xx status code
func (o *GetEventsEntitiesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get events entities too many requests response a status code equal to that given
func (o *GetEventsEntitiesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get events entities too many requests response
func (o *GetEventsEntitiesTooManyRequests) Code() int {
	return 429
}

func (o *GetEventsEntitiesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /ti/events/entities/events/GET/v2][%d] getEventsEntitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetEventsEntitiesTooManyRequests) String() string {
	return fmt.Sprintf("[POST /ti/events/entities/events/GET/v2][%d] getEventsEntitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetEventsEntitiesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetEventsEntitiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetEventsEntitiesInternalServerError creates a GetEventsEntitiesInternalServerError with default headers values
func NewGetEventsEntitiesInternalServerError() *GetEventsEntitiesInternalServerError {
	return &GetEventsEntitiesInternalServerError{}
}

/*
GetEventsEntitiesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetEventsEntitiesInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainEventEntitiesResponse
}

// IsSuccess returns true when this get events entities internal server error response has a 2xx status code
func (o *GetEventsEntitiesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get events entities internal server error response has a 3xx status code
func (o *GetEventsEntitiesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get events entities internal server error response has a 4xx status code
func (o *GetEventsEntitiesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get events entities internal server error response has a 5xx status code
func (o *GetEventsEntitiesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get events entities internal server error response a status code equal to that given
func (o *GetEventsEntitiesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get events entities internal server error response
func (o *GetEventsEntitiesInternalServerError) Code() int {
	return 500
}

func (o *GetEventsEntitiesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /ti/events/entities/events/GET/v2][%d] getEventsEntitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetEventsEntitiesInternalServerError) String() string {
	return fmt.Sprintf("[POST /ti/events/entities/events/GET/v2][%d] getEventsEntitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetEventsEntitiesInternalServerError) GetPayload() *models.DomainEventEntitiesResponse {
	return o.Payload
}

func (o *GetEventsEntitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainEventEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
