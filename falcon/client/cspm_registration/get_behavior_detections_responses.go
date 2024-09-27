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

// GetBehaviorDetectionsReader is a Reader for the GetBehaviorDetections structure.
type GetBehaviorDetectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBehaviorDetectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBehaviorDetectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBehaviorDetectionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetBehaviorDetectionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetBehaviorDetectionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetBehaviorDetectionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /detects/entities/ioa/v1] GetBehaviorDetections", response, response.Code())
	}
}

// NewGetBehaviorDetectionsOK creates a GetBehaviorDetectionsOK with default headers values
func NewGetBehaviorDetectionsOK() *GetBehaviorDetectionsOK {
	return &GetBehaviorDetectionsOK{}
}

/*
GetBehaviorDetectionsOK describes a response with status code 200, with default header values.

OK
*/
type GetBehaviorDetectionsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationExternalIOAEventResponse
}

// IsSuccess returns true when this get behavior detections o k response has a 2xx status code
func (o *GetBehaviorDetectionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get behavior detections o k response has a 3xx status code
func (o *GetBehaviorDetectionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get behavior detections o k response has a 4xx status code
func (o *GetBehaviorDetectionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get behavior detections o k response has a 5xx status code
func (o *GetBehaviorDetectionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get behavior detections o k response a status code equal to that given
func (o *GetBehaviorDetectionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get behavior detections o k response
func (o *GetBehaviorDetectionsOK) Code() int {
	return 200
}

func (o *GetBehaviorDetectionsOK) Error() string {
	return fmt.Sprintf("[GET /detects/entities/ioa/v1][%d] getBehaviorDetectionsOK  %+v", 200, o.Payload)
}

func (o *GetBehaviorDetectionsOK) String() string {
	return fmt.Sprintf("[GET /detects/entities/ioa/v1][%d] getBehaviorDetectionsOK  %+v", 200, o.Payload)
}

func (o *GetBehaviorDetectionsOK) GetPayload() *models.RegistrationExternalIOAEventResponse {
	return o.Payload
}

func (o *GetBehaviorDetectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationExternalIOAEventResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBehaviorDetectionsBadRequest creates a GetBehaviorDetectionsBadRequest with default headers values
func NewGetBehaviorDetectionsBadRequest() *GetBehaviorDetectionsBadRequest {
	return &GetBehaviorDetectionsBadRequest{}
}

/*
GetBehaviorDetectionsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetBehaviorDetectionsBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationExternalIOAEventResponse
}

// IsSuccess returns true when this get behavior detections bad request response has a 2xx status code
func (o *GetBehaviorDetectionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get behavior detections bad request response has a 3xx status code
func (o *GetBehaviorDetectionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get behavior detections bad request response has a 4xx status code
func (o *GetBehaviorDetectionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get behavior detections bad request response has a 5xx status code
func (o *GetBehaviorDetectionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get behavior detections bad request response a status code equal to that given
func (o *GetBehaviorDetectionsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get behavior detections bad request response
func (o *GetBehaviorDetectionsBadRequest) Code() int {
	return 400
}

func (o *GetBehaviorDetectionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /detects/entities/ioa/v1][%d] getBehaviorDetectionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetBehaviorDetectionsBadRequest) String() string {
	return fmt.Sprintf("[GET /detects/entities/ioa/v1][%d] getBehaviorDetectionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetBehaviorDetectionsBadRequest) GetPayload() *models.RegistrationExternalIOAEventResponse {
	return o.Payload
}

func (o *GetBehaviorDetectionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationExternalIOAEventResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBehaviorDetectionsForbidden creates a GetBehaviorDetectionsForbidden with default headers values
func NewGetBehaviorDetectionsForbidden() *GetBehaviorDetectionsForbidden {
	return &GetBehaviorDetectionsForbidden{}
}

/*
GetBehaviorDetectionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetBehaviorDetectionsForbidden struct {

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

// IsSuccess returns true when this get behavior detections forbidden response has a 2xx status code
func (o *GetBehaviorDetectionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get behavior detections forbidden response has a 3xx status code
func (o *GetBehaviorDetectionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get behavior detections forbidden response has a 4xx status code
func (o *GetBehaviorDetectionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get behavior detections forbidden response has a 5xx status code
func (o *GetBehaviorDetectionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get behavior detections forbidden response a status code equal to that given
func (o *GetBehaviorDetectionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get behavior detections forbidden response
func (o *GetBehaviorDetectionsForbidden) Code() int {
	return 403
}

func (o *GetBehaviorDetectionsForbidden) Error() string {
	return fmt.Sprintf("[GET /detects/entities/ioa/v1][%d] getBehaviorDetectionsForbidden  %+v", 403, o.Payload)
}

func (o *GetBehaviorDetectionsForbidden) String() string {
	return fmt.Sprintf("[GET /detects/entities/ioa/v1][%d] getBehaviorDetectionsForbidden  %+v", 403, o.Payload)
}

func (o *GetBehaviorDetectionsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetBehaviorDetectionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBehaviorDetectionsTooManyRequests creates a GetBehaviorDetectionsTooManyRequests with default headers values
func NewGetBehaviorDetectionsTooManyRequests() *GetBehaviorDetectionsTooManyRequests {
	return &GetBehaviorDetectionsTooManyRequests{}
}

/*
GetBehaviorDetectionsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetBehaviorDetectionsTooManyRequests struct {

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

// IsSuccess returns true when this get behavior detections too many requests response has a 2xx status code
func (o *GetBehaviorDetectionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get behavior detections too many requests response has a 3xx status code
func (o *GetBehaviorDetectionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get behavior detections too many requests response has a 4xx status code
func (o *GetBehaviorDetectionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get behavior detections too many requests response has a 5xx status code
func (o *GetBehaviorDetectionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get behavior detections too many requests response a status code equal to that given
func (o *GetBehaviorDetectionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get behavior detections too many requests response
func (o *GetBehaviorDetectionsTooManyRequests) Code() int {
	return 429
}

func (o *GetBehaviorDetectionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /detects/entities/ioa/v1][%d] getBehaviorDetectionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetBehaviorDetectionsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /detects/entities/ioa/v1][%d] getBehaviorDetectionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetBehaviorDetectionsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetBehaviorDetectionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBehaviorDetectionsInternalServerError creates a GetBehaviorDetectionsInternalServerError with default headers values
func NewGetBehaviorDetectionsInternalServerError() *GetBehaviorDetectionsInternalServerError {
	return &GetBehaviorDetectionsInternalServerError{}
}

/*
GetBehaviorDetectionsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetBehaviorDetectionsInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationExternalIOAEventResponse
}

// IsSuccess returns true when this get behavior detections internal server error response has a 2xx status code
func (o *GetBehaviorDetectionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get behavior detections internal server error response has a 3xx status code
func (o *GetBehaviorDetectionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get behavior detections internal server error response has a 4xx status code
func (o *GetBehaviorDetectionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get behavior detections internal server error response has a 5xx status code
func (o *GetBehaviorDetectionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get behavior detections internal server error response a status code equal to that given
func (o *GetBehaviorDetectionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get behavior detections internal server error response
func (o *GetBehaviorDetectionsInternalServerError) Code() int {
	return 500
}

func (o *GetBehaviorDetectionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /detects/entities/ioa/v1][%d] getBehaviorDetectionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetBehaviorDetectionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /detects/entities/ioa/v1][%d] getBehaviorDetectionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetBehaviorDetectionsInternalServerError) GetPayload() *models.RegistrationExternalIOAEventResponse {
	return o.Payload
}

func (o *GetBehaviorDetectionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationExternalIOAEventResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
