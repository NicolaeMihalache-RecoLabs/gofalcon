// Code generated by go-swagger; DO NOT EDIT.

package discover

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

// GetApplicationsReader is a Reader for the GetApplications structure.
type GetApplicationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApplicationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetApplicationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetApplicationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetApplicationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /discover/entities/applications/v1] get-applications", response, response.Code())
	}
}

// NewGetApplicationsOK creates a GetApplicationsOK with default headers values
func NewGetApplicationsOK() *GetApplicationsOK {
	return &GetApplicationsOK{}
}

/*
GetApplicationsOK describes a response with status code 200, with default header values.

OK
*/
type GetApplicationsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainDiscoverAPIApplicationEntitiesResponse
}

// IsSuccess returns true when this get applications o k response has a 2xx status code
func (o *GetApplicationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get applications o k response has a 3xx status code
func (o *GetApplicationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get applications o k response has a 4xx status code
func (o *GetApplicationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get applications o k response has a 5xx status code
func (o *GetApplicationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get applications o k response a status code equal to that given
func (o *GetApplicationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get applications o k response
func (o *GetApplicationsOK) Code() int {
	return 200
}

func (o *GetApplicationsOK) Error() string {
	return fmt.Sprintf("[GET /discover/entities/applications/v1][%d] getApplicationsOK  %+v", 200, o.Payload)
}

func (o *GetApplicationsOK) String() string {
	return fmt.Sprintf("[GET /discover/entities/applications/v1][%d] getApplicationsOK  %+v", 200, o.Payload)
}

func (o *GetApplicationsOK) GetPayload() *models.DomainDiscoverAPIApplicationEntitiesResponse {
	return o.Payload
}

func (o *GetApplicationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainDiscoverAPIApplicationEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationsBadRequest creates a GetApplicationsBadRequest with default headers values
func NewGetApplicationsBadRequest() *GetApplicationsBadRequest {
	return &GetApplicationsBadRequest{}
}

/*
GetApplicationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetApplicationsBadRequest struct {

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

// IsSuccess returns true when this get applications bad request response has a 2xx status code
func (o *GetApplicationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get applications bad request response has a 3xx status code
func (o *GetApplicationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get applications bad request response has a 4xx status code
func (o *GetApplicationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get applications bad request response has a 5xx status code
func (o *GetApplicationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get applications bad request response a status code equal to that given
func (o *GetApplicationsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get applications bad request response
func (o *GetApplicationsBadRequest) Code() int {
	return 400
}

func (o *GetApplicationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /discover/entities/applications/v1][%d] getApplicationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetApplicationsBadRequest) String() string {
	return fmt.Sprintf("[GET /discover/entities/applications/v1][%d] getApplicationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetApplicationsBadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *GetApplicationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetApplicationsForbidden creates a GetApplicationsForbidden with default headers values
func NewGetApplicationsForbidden() *GetApplicationsForbidden {
	return &GetApplicationsForbidden{}
}

/*
GetApplicationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetApplicationsForbidden struct {

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

// IsSuccess returns true when this get applications forbidden response has a 2xx status code
func (o *GetApplicationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get applications forbidden response has a 3xx status code
func (o *GetApplicationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get applications forbidden response has a 4xx status code
func (o *GetApplicationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get applications forbidden response has a 5xx status code
func (o *GetApplicationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get applications forbidden response a status code equal to that given
func (o *GetApplicationsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get applications forbidden response
func (o *GetApplicationsForbidden) Code() int {
	return 403
}

func (o *GetApplicationsForbidden) Error() string {
	return fmt.Sprintf("[GET /discover/entities/applications/v1][%d] getApplicationsForbidden  %+v", 403, o.Payload)
}

func (o *GetApplicationsForbidden) String() string {
	return fmt.Sprintf("[GET /discover/entities/applications/v1][%d] getApplicationsForbidden  %+v", 403, o.Payload)
}

func (o *GetApplicationsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetApplicationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetApplicationsTooManyRequests creates a GetApplicationsTooManyRequests with default headers values
func NewGetApplicationsTooManyRequests() *GetApplicationsTooManyRequests {
	return &GetApplicationsTooManyRequests{}
}

/*
GetApplicationsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetApplicationsTooManyRequests struct {

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

// IsSuccess returns true when this get applications too many requests response has a 2xx status code
func (o *GetApplicationsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get applications too many requests response has a 3xx status code
func (o *GetApplicationsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get applications too many requests response has a 4xx status code
func (o *GetApplicationsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get applications too many requests response has a 5xx status code
func (o *GetApplicationsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get applications too many requests response a status code equal to that given
func (o *GetApplicationsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get applications too many requests response
func (o *GetApplicationsTooManyRequests) Code() int {
	return 429
}

func (o *GetApplicationsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /discover/entities/applications/v1][%d] getApplicationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetApplicationsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /discover/entities/applications/v1][%d] getApplicationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetApplicationsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetApplicationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetApplicationsInternalServerError creates a GetApplicationsInternalServerError with default headers values
func NewGetApplicationsInternalServerError() *GetApplicationsInternalServerError {
	return &GetApplicationsInternalServerError{}
}

/*
GetApplicationsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetApplicationsInternalServerError struct {

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

// IsSuccess returns true when this get applications internal server error response has a 2xx status code
func (o *GetApplicationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get applications internal server error response has a 3xx status code
func (o *GetApplicationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get applications internal server error response has a 4xx status code
func (o *GetApplicationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get applications internal server error response has a 5xx status code
func (o *GetApplicationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get applications internal server error response a status code equal to that given
func (o *GetApplicationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get applications internal server error response
func (o *GetApplicationsInternalServerError) Code() int {
	return 500
}

func (o *GetApplicationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /discover/entities/applications/v1][%d] getApplicationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetApplicationsInternalServerError) String() string {
	return fmt.Sprintf("[GET /discover/entities/applications/v1][%d] getApplicationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetApplicationsInternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *GetApplicationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
