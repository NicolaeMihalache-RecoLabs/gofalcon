// Code generated by go-swagger; DO NOT EDIT.

package runtime_detections

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

// GetRuntimeDetectionsCombinedV2Reader is a Reader for the GetRuntimeDetectionsCombinedV2 structure.
type GetRuntimeDetectionsCombinedV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRuntimeDetectionsCombinedV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRuntimeDetectionsCombinedV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRuntimeDetectionsCombinedV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRuntimeDetectionsCombinedV2TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRuntimeDetectionsCombinedV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/combined/runtime-detections/v2] GetRuntimeDetectionsCombinedV2", response, response.Code())
	}
}

// NewGetRuntimeDetectionsCombinedV2OK creates a GetRuntimeDetectionsCombinedV2OK with default headers values
func NewGetRuntimeDetectionsCombinedV2OK() *GetRuntimeDetectionsCombinedV2OK {
	return &GetRuntimeDetectionsCombinedV2OK{}
}

/*
GetRuntimeDetectionsCombinedV2OK describes a response with status code 200, with default header values.

OK
*/
type GetRuntimeDetectionsCombinedV2OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RuntimedetectionsDetectionsEntityResponse
}

// IsSuccess returns true when this get runtime detections combined v2 o k response has a 2xx status code
func (o *GetRuntimeDetectionsCombinedV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get runtime detections combined v2 o k response has a 3xx status code
func (o *GetRuntimeDetectionsCombinedV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runtime detections combined v2 o k response has a 4xx status code
func (o *GetRuntimeDetectionsCombinedV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get runtime detections combined v2 o k response has a 5xx status code
func (o *GetRuntimeDetectionsCombinedV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get runtime detections combined v2 o k response a status code equal to that given
func (o *GetRuntimeDetectionsCombinedV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get runtime detections combined v2 o k response
func (o *GetRuntimeDetectionsCombinedV2OK) Code() int {
	return 200
}

func (o *GetRuntimeDetectionsCombinedV2OK) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/runtime-detections/v2][%d] getRuntimeDetectionsCombinedV2OK  %+v", 200, o.Payload)
}

func (o *GetRuntimeDetectionsCombinedV2OK) String() string {
	return fmt.Sprintf("[GET /container-security/combined/runtime-detections/v2][%d] getRuntimeDetectionsCombinedV2OK  %+v", 200, o.Payload)
}

func (o *GetRuntimeDetectionsCombinedV2OK) GetPayload() *models.RuntimedetectionsDetectionsEntityResponse {
	return o.Payload
}

func (o *GetRuntimeDetectionsCombinedV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RuntimedetectionsDetectionsEntityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuntimeDetectionsCombinedV2Forbidden creates a GetRuntimeDetectionsCombinedV2Forbidden with default headers values
func NewGetRuntimeDetectionsCombinedV2Forbidden() *GetRuntimeDetectionsCombinedV2Forbidden {
	return &GetRuntimeDetectionsCombinedV2Forbidden{}
}

/*
GetRuntimeDetectionsCombinedV2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRuntimeDetectionsCombinedV2Forbidden struct {

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

// IsSuccess returns true when this get runtime detections combined v2 forbidden response has a 2xx status code
func (o *GetRuntimeDetectionsCombinedV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runtime detections combined v2 forbidden response has a 3xx status code
func (o *GetRuntimeDetectionsCombinedV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runtime detections combined v2 forbidden response has a 4xx status code
func (o *GetRuntimeDetectionsCombinedV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runtime detections combined v2 forbidden response has a 5xx status code
func (o *GetRuntimeDetectionsCombinedV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get runtime detections combined v2 forbidden response a status code equal to that given
func (o *GetRuntimeDetectionsCombinedV2Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get runtime detections combined v2 forbidden response
func (o *GetRuntimeDetectionsCombinedV2Forbidden) Code() int {
	return 403
}

func (o *GetRuntimeDetectionsCombinedV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/runtime-detections/v2][%d] getRuntimeDetectionsCombinedV2Forbidden  %+v", 403, o.Payload)
}

func (o *GetRuntimeDetectionsCombinedV2Forbidden) String() string {
	return fmt.Sprintf("[GET /container-security/combined/runtime-detections/v2][%d] getRuntimeDetectionsCombinedV2Forbidden  %+v", 403, o.Payload)
}

func (o *GetRuntimeDetectionsCombinedV2Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetRuntimeDetectionsCombinedV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRuntimeDetectionsCombinedV2TooManyRequests creates a GetRuntimeDetectionsCombinedV2TooManyRequests with default headers values
func NewGetRuntimeDetectionsCombinedV2TooManyRequests() *GetRuntimeDetectionsCombinedV2TooManyRequests {
	return &GetRuntimeDetectionsCombinedV2TooManyRequests{}
}

/*
GetRuntimeDetectionsCombinedV2TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetRuntimeDetectionsCombinedV2TooManyRequests struct {

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

// IsSuccess returns true when this get runtime detections combined v2 too many requests response has a 2xx status code
func (o *GetRuntimeDetectionsCombinedV2TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runtime detections combined v2 too many requests response has a 3xx status code
func (o *GetRuntimeDetectionsCombinedV2TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runtime detections combined v2 too many requests response has a 4xx status code
func (o *GetRuntimeDetectionsCombinedV2TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runtime detections combined v2 too many requests response has a 5xx status code
func (o *GetRuntimeDetectionsCombinedV2TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get runtime detections combined v2 too many requests response a status code equal to that given
func (o *GetRuntimeDetectionsCombinedV2TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get runtime detections combined v2 too many requests response
func (o *GetRuntimeDetectionsCombinedV2TooManyRequests) Code() int {
	return 429
}

func (o *GetRuntimeDetectionsCombinedV2TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/runtime-detections/v2][%d] getRuntimeDetectionsCombinedV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRuntimeDetectionsCombinedV2TooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/combined/runtime-detections/v2][%d] getRuntimeDetectionsCombinedV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRuntimeDetectionsCombinedV2TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetRuntimeDetectionsCombinedV2TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRuntimeDetectionsCombinedV2InternalServerError creates a GetRuntimeDetectionsCombinedV2InternalServerError with default headers values
func NewGetRuntimeDetectionsCombinedV2InternalServerError() *GetRuntimeDetectionsCombinedV2InternalServerError {
	return &GetRuntimeDetectionsCombinedV2InternalServerError{}
}

/*
GetRuntimeDetectionsCombinedV2InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetRuntimeDetectionsCombinedV2InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.CoreEntitiesResponse
}

// IsSuccess returns true when this get runtime detections combined v2 internal server error response has a 2xx status code
func (o *GetRuntimeDetectionsCombinedV2InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runtime detections combined v2 internal server error response has a 3xx status code
func (o *GetRuntimeDetectionsCombinedV2InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runtime detections combined v2 internal server error response has a 4xx status code
func (o *GetRuntimeDetectionsCombinedV2InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get runtime detections combined v2 internal server error response has a 5xx status code
func (o *GetRuntimeDetectionsCombinedV2InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get runtime detections combined v2 internal server error response a status code equal to that given
func (o *GetRuntimeDetectionsCombinedV2InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get runtime detections combined v2 internal server error response
func (o *GetRuntimeDetectionsCombinedV2InternalServerError) Code() int {
	return 500
}

func (o *GetRuntimeDetectionsCombinedV2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/runtime-detections/v2][%d] getRuntimeDetectionsCombinedV2InternalServerError  %+v", 500, o.Payload)
}

func (o *GetRuntimeDetectionsCombinedV2InternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/combined/runtime-detections/v2][%d] getRuntimeDetectionsCombinedV2InternalServerError  %+v", 500, o.Payload)
}

func (o *GetRuntimeDetectionsCombinedV2InternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *GetRuntimeDetectionsCombinedV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.CoreEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
