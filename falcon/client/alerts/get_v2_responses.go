// Code generated by go-swagger; DO NOT EDIT.

package alerts

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

// GetV2Reader is a Reader for the GetV2 structure.
type GetV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetV2TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /alerts/entities/alerts/v2] GetV2", response, response.Code())
	}
}

// NewGetV2OK creates a GetV2OK with default headers values
func NewGetV2OK() *GetV2OK {
	return &GetV2OK{}
}

/*
GetV2OK describes a response with status code 200, with default header values.

OK
*/
type GetV2OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DetectsapiPostEntitiesAlertsV2Response
}

// IsSuccess returns true when this get v2 o k response has a 2xx status code
func (o *GetV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v2 o k response has a 3xx status code
func (o *GetV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v2 o k response has a 4xx status code
func (o *GetV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v2 o k response has a 5xx status code
func (o *GetV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v2 o k response a status code equal to that given
func (o *GetV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v2 o k response
func (o *GetV2OK) Code() int {
	return 200
}

func (o *GetV2OK) Error() string {
	return fmt.Sprintf("[POST /alerts/entities/alerts/v2][%d] getV2OK  %+v", 200, o.Payload)
}

func (o *GetV2OK) String() string {
	return fmt.Sprintf("[POST /alerts/entities/alerts/v2][%d] getV2OK  %+v", 200, o.Payload)
}

func (o *GetV2OK) GetPayload() *models.DetectsapiPostEntitiesAlertsV2Response {
	return o.Payload
}

func (o *GetV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DetectsapiPostEntitiesAlertsV2Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV2BadRequest creates a GetV2BadRequest with default headers values
func NewGetV2BadRequest() *GetV2BadRequest {
	return &GetV2BadRequest{}
}

/*
GetV2BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetV2BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DetectsapiPostEntitiesAlertsV2Response
}

// IsSuccess returns true when this get v2 bad request response has a 2xx status code
func (o *GetV2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v2 bad request response has a 3xx status code
func (o *GetV2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v2 bad request response has a 4xx status code
func (o *GetV2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v2 bad request response has a 5xx status code
func (o *GetV2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get v2 bad request response a status code equal to that given
func (o *GetV2BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get v2 bad request response
func (o *GetV2BadRequest) Code() int {
	return 400
}

func (o *GetV2BadRequest) Error() string {
	return fmt.Sprintf("[POST /alerts/entities/alerts/v2][%d] getV2BadRequest  %+v", 400, o.Payload)
}

func (o *GetV2BadRequest) String() string {
	return fmt.Sprintf("[POST /alerts/entities/alerts/v2][%d] getV2BadRequest  %+v", 400, o.Payload)
}

func (o *GetV2BadRequest) GetPayload() *models.DetectsapiPostEntitiesAlertsV2Response {
	return o.Payload
}

func (o *GetV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DetectsapiPostEntitiesAlertsV2Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV2Forbidden creates a GetV2Forbidden with default headers values
func NewGetV2Forbidden() *GetV2Forbidden {
	return &GetV2Forbidden{}
}

/*
GetV2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetV2Forbidden struct {

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

// IsSuccess returns true when this get v2 forbidden response has a 2xx status code
func (o *GetV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v2 forbidden response has a 3xx status code
func (o *GetV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v2 forbidden response has a 4xx status code
func (o *GetV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v2 forbidden response has a 5xx status code
func (o *GetV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get v2 forbidden response a status code equal to that given
func (o *GetV2Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get v2 forbidden response
func (o *GetV2Forbidden) Code() int {
	return 403
}

func (o *GetV2Forbidden) Error() string {
	return fmt.Sprintf("[POST /alerts/entities/alerts/v2][%d] getV2Forbidden  %+v", 403, o.Payload)
}

func (o *GetV2Forbidden) String() string {
	return fmt.Sprintf("[POST /alerts/entities/alerts/v2][%d] getV2Forbidden  %+v", 403, o.Payload)
}

func (o *GetV2Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetV2TooManyRequests creates a GetV2TooManyRequests with default headers values
func NewGetV2TooManyRequests() *GetV2TooManyRequests {
	return &GetV2TooManyRequests{}
}

/*
GetV2TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetV2TooManyRequests struct {

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

// IsSuccess returns true when this get v2 too many requests response has a 2xx status code
func (o *GetV2TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v2 too many requests response has a 3xx status code
func (o *GetV2TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v2 too many requests response has a 4xx status code
func (o *GetV2TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v2 too many requests response has a 5xx status code
func (o *GetV2TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get v2 too many requests response a status code equal to that given
func (o *GetV2TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get v2 too many requests response
func (o *GetV2TooManyRequests) Code() int {
	return 429
}

func (o *GetV2TooManyRequests) Error() string {
	return fmt.Sprintf("[POST /alerts/entities/alerts/v2][%d] getV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetV2TooManyRequests) String() string {
	return fmt.Sprintf("[POST /alerts/entities/alerts/v2][%d] getV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetV2TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetV2TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetV2InternalServerError creates a GetV2InternalServerError with default headers values
func NewGetV2InternalServerError() *GetV2InternalServerError {
	return &GetV2InternalServerError{}
}

/*
GetV2InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetV2InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DetectsapiPostEntitiesAlertsV2Response
}

// IsSuccess returns true when this get v2 internal server error response has a 2xx status code
func (o *GetV2InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v2 internal server error response has a 3xx status code
func (o *GetV2InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v2 internal server error response has a 4xx status code
func (o *GetV2InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v2 internal server error response has a 5xx status code
func (o *GetV2InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get v2 internal server error response a status code equal to that given
func (o *GetV2InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get v2 internal server error response
func (o *GetV2InternalServerError) Code() int {
	return 500
}

func (o *GetV2InternalServerError) Error() string {
	return fmt.Sprintf("[POST /alerts/entities/alerts/v2][%d] getV2InternalServerError  %+v", 500, o.Payload)
}

func (o *GetV2InternalServerError) String() string {
	return fmt.Sprintf("[POST /alerts/entities/alerts/v2][%d] getV2InternalServerError  %+v", 500, o.Payload)
}

func (o *GetV2InternalServerError) GetPayload() *models.DetectsapiPostEntitiesAlertsV2Response {
	return o.Payload
}

func (o *GetV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DetectsapiPostEntitiesAlertsV2Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
