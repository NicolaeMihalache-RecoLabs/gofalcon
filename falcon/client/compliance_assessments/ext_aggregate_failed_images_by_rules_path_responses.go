// Code generated by go-swagger; DO NOT EDIT.

package compliance_assessments

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

// ExtAggregateFailedImagesByRulesPathReader is a Reader for the ExtAggregateFailedImagesByRulesPath structure.
type ExtAggregateFailedImagesByRulesPathReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtAggregateFailedImagesByRulesPathReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtAggregateFailedImagesByRulesPathOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExtAggregateFailedImagesByRulesPathBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewExtAggregateFailedImagesByRulesPathUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExtAggregateFailedImagesByRulesPathForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewExtAggregateFailedImagesByRulesPathTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewExtAggregateFailedImagesByRulesPathInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-compliance/aggregates/failed-images-by-rules/v2] extAggregateFailedImagesByRulesPath", response, response.Code())
	}
}

// NewExtAggregateFailedImagesByRulesPathOK creates a ExtAggregateFailedImagesByRulesPathOK with default headers values
func NewExtAggregateFailedImagesByRulesPathOK() *ExtAggregateFailedImagesByRulesPathOK {
	return &ExtAggregateFailedImagesByRulesPathOK{}
}

/*
ExtAggregateFailedImagesByRulesPathOK describes a response with status code 200, with default header values.

OK
*/
type ExtAggregateFailedImagesByRulesPathOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAggregateFailedAssetsByRulesResponse
}

// IsSuccess returns true when this ext aggregate failed images by rules path o k response has a 2xx status code
func (o *ExtAggregateFailedImagesByRulesPathOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ext aggregate failed images by rules path o k response has a 3xx status code
func (o *ExtAggregateFailedImagesByRulesPathOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ext aggregate failed images by rules path o k response has a 4xx status code
func (o *ExtAggregateFailedImagesByRulesPathOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ext aggregate failed images by rules path o k response has a 5xx status code
func (o *ExtAggregateFailedImagesByRulesPathOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ext aggregate failed images by rules path o k response a status code equal to that given
func (o *ExtAggregateFailedImagesByRulesPathOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ext aggregate failed images by rules path o k response
func (o *ExtAggregateFailedImagesByRulesPathOK) Code() int {
	return 200
}

func (o *ExtAggregateFailedImagesByRulesPathOK) Error() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-images-by-rules/v2][%d] extAggregateFailedImagesByRulesPathOK  %+v", 200, o.Payload)
}

func (o *ExtAggregateFailedImagesByRulesPathOK) String() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-images-by-rules/v2][%d] extAggregateFailedImagesByRulesPathOK  %+v", 200, o.Payload)
}

func (o *ExtAggregateFailedImagesByRulesPathOK) GetPayload() *models.DomainAggregateFailedAssetsByRulesResponse {
	return o.Payload
}

func (o *ExtAggregateFailedImagesByRulesPathOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAggregateFailedAssetsByRulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtAggregateFailedImagesByRulesPathBadRequest creates a ExtAggregateFailedImagesByRulesPathBadRequest with default headers values
func NewExtAggregateFailedImagesByRulesPathBadRequest() *ExtAggregateFailedImagesByRulesPathBadRequest {
	return &ExtAggregateFailedImagesByRulesPathBadRequest{}
}

/*
ExtAggregateFailedImagesByRulesPathBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExtAggregateFailedImagesByRulesPathBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAggregateFailedAssetsByRulesResponse
}

// IsSuccess returns true when this ext aggregate failed images by rules path bad request response has a 2xx status code
func (o *ExtAggregateFailedImagesByRulesPathBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ext aggregate failed images by rules path bad request response has a 3xx status code
func (o *ExtAggregateFailedImagesByRulesPathBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ext aggregate failed images by rules path bad request response has a 4xx status code
func (o *ExtAggregateFailedImagesByRulesPathBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ext aggregate failed images by rules path bad request response has a 5xx status code
func (o *ExtAggregateFailedImagesByRulesPathBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ext aggregate failed images by rules path bad request response a status code equal to that given
func (o *ExtAggregateFailedImagesByRulesPathBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the ext aggregate failed images by rules path bad request response
func (o *ExtAggregateFailedImagesByRulesPathBadRequest) Code() int {
	return 400
}

func (o *ExtAggregateFailedImagesByRulesPathBadRequest) Error() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-images-by-rules/v2][%d] extAggregateFailedImagesByRulesPathBadRequest  %+v", 400, o.Payload)
}

func (o *ExtAggregateFailedImagesByRulesPathBadRequest) String() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-images-by-rules/v2][%d] extAggregateFailedImagesByRulesPathBadRequest  %+v", 400, o.Payload)
}

func (o *ExtAggregateFailedImagesByRulesPathBadRequest) GetPayload() *models.DomainAggregateFailedAssetsByRulesResponse {
	return o.Payload
}

func (o *ExtAggregateFailedImagesByRulesPathBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAggregateFailedAssetsByRulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtAggregateFailedImagesByRulesPathUnauthorized creates a ExtAggregateFailedImagesByRulesPathUnauthorized with default headers values
func NewExtAggregateFailedImagesByRulesPathUnauthorized() *ExtAggregateFailedImagesByRulesPathUnauthorized {
	return &ExtAggregateFailedImagesByRulesPathUnauthorized{}
}

/*
ExtAggregateFailedImagesByRulesPathUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ExtAggregateFailedImagesByRulesPathUnauthorized struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAggregateFailedAssetsByRulesResponse
}

// IsSuccess returns true when this ext aggregate failed images by rules path unauthorized response has a 2xx status code
func (o *ExtAggregateFailedImagesByRulesPathUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ext aggregate failed images by rules path unauthorized response has a 3xx status code
func (o *ExtAggregateFailedImagesByRulesPathUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ext aggregate failed images by rules path unauthorized response has a 4xx status code
func (o *ExtAggregateFailedImagesByRulesPathUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ext aggregate failed images by rules path unauthorized response has a 5xx status code
func (o *ExtAggregateFailedImagesByRulesPathUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ext aggregate failed images by rules path unauthorized response a status code equal to that given
func (o *ExtAggregateFailedImagesByRulesPathUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the ext aggregate failed images by rules path unauthorized response
func (o *ExtAggregateFailedImagesByRulesPathUnauthorized) Code() int {
	return 401
}

func (o *ExtAggregateFailedImagesByRulesPathUnauthorized) Error() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-images-by-rules/v2][%d] extAggregateFailedImagesByRulesPathUnauthorized  %+v", 401, o.Payload)
}

func (o *ExtAggregateFailedImagesByRulesPathUnauthorized) String() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-images-by-rules/v2][%d] extAggregateFailedImagesByRulesPathUnauthorized  %+v", 401, o.Payload)
}

func (o *ExtAggregateFailedImagesByRulesPathUnauthorized) GetPayload() *models.DomainAggregateFailedAssetsByRulesResponse {
	return o.Payload
}

func (o *ExtAggregateFailedImagesByRulesPathUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAggregateFailedAssetsByRulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtAggregateFailedImagesByRulesPathForbidden creates a ExtAggregateFailedImagesByRulesPathForbidden with default headers values
func NewExtAggregateFailedImagesByRulesPathForbidden() *ExtAggregateFailedImagesByRulesPathForbidden {
	return &ExtAggregateFailedImagesByRulesPathForbidden{}
}

/*
ExtAggregateFailedImagesByRulesPathForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ExtAggregateFailedImagesByRulesPathForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAggregateFailedAssetsByRulesResponse
}

// IsSuccess returns true when this ext aggregate failed images by rules path forbidden response has a 2xx status code
func (o *ExtAggregateFailedImagesByRulesPathForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ext aggregate failed images by rules path forbidden response has a 3xx status code
func (o *ExtAggregateFailedImagesByRulesPathForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ext aggregate failed images by rules path forbidden response has a 4xx status code
func (o *ExtAggregateFailedImagesByRulesPathForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ext aggregate failed images by rules path forbidden response has a 5xx status code
func (o *ExtAggregateFailedImagesByRulesPathForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ext aggregate failed images by rules path forbidden response a status code equal to that given
func (o *ExtAggregateFailedImagesByRulesPathForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the ext aggregate failed images by rules path forbidden response
func (o *ExtAggregateFailedImagesByRulesPathForbidden) Code() int {
	return 403
}

func (o *ExtAggregateFailedImagesByRulesPathForbidden) Error() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-images-by-rules/v2][%d] extAggregateFailedImagesByRulesPathForbidden  %+v", 403, o.Payload)
}

func (o *ExtAggregateFailedImagesByRulesPathForbidden) String() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-images-by-rules/v2][%d] extAggregateFailedImagesByRulesPathForbidden  %+v", 403, o.Payload)
}

func (o *ExtAggregateFailedImagesByRulesPathForbidden) GetPayload() *models.DomainAggregateFailedAssetsByRulesResponse {
	return o.Payload
}

func (o *ExtAggregateFailedImagesByRulesPathForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAggregateFailedAssetsByRulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtAggregateFailedImagesByRulesPathTooManyRequests creates a ExtAggregateFailedImagesByRulesPathTooManyRequests with default headers values
func NewExtAggregateFailedImagesByRulesPathTooManyRequests() *ExtAggregateFailedImagesByRulesPathTooManyRequests {
	return &ExtAggregateFailedImagesByRulesPathTooManyRequests{}
}

/*
ExtAggregateFailedImagesByRulesPathTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ExtAggregateFailedImagesByRulesPathTooManyRequests struct {

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

// IsSuccess returns true when this ext aggregate failed images by rules path too many requests response has a 2xx status code
func (o *ExtAggregateFailedImagesByRulesPathTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ext aggregate failed images by rules path too many requests response has a 3xx status code
func (o *ExtAggregateFailedImagesByRulesPathTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ext aggregate failed images by rules path too many requests response has a 4xx status code
func (o *ExtAggregateFailedImagesByRulesPathTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this ext aggregate failed images by rules path too many requests response has a 5xx status code
func (o *ExtAggregateFailedImagesByRulesPathTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this ext aggregate failed images by rules path too many requests response a status code equal to that given
func (o *ExtAggregateFailedImagesByRulesPathTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the ext aggregate failed images by rules path too many requests response
func (o *ExtAggregateFailedImagesByRulesPathTooManyRequests) Code() int {
	return 429
}

func (o *ExtAggregateFailedImagesByRulesPathTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-images-by-rules/v2][%d] extAggregateFailedImagesByRulesPathTooManyRequests  %+v", 429, o.Payload)
}

func (o *ExtAggregateFailedImagesByRulesPathTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-images-by-rules/v2][%d] extAggregateFailedImagesByRulesPathTooManyRequests  %+v", 429, o.Payload)
}

func (o *ExtAggregateFailedImagesByRulesPathTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ExtAggregateFailedImagesByRulesPathTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewExtAggregateFailedImagesByRulesPathInternalServerError creates a ExtAggregateFailedImagesByRulesPathInternalServerError with default headers values
func NewExtAggregateFailedImagesByRulesPathInternalServerError() *ExtAggregateFailedImagesByRulesPathInternalServerError {
	return &ExtAggregateFailedImagesByRulesPathInternalServerError{}
}

/*
ExtAggregateFailedImagesByRulesPathInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ExtAggregateFailedImagesByRulesPathInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAggregateFailedAssetsByRulesResponse
}

// IsSuccess returns true when this ext aggregate failed images by rules path internal server error response has a 2xx status code
func (o *ExtAggregateFailedImagesByRulesPathInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ext aggregate failed images by rules path internal server error response has a 3xx status code
func (o *ExtAggregateFailedImagesByRulesPathInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ext aggregate failed images by rules path internal server error response has a 4xx status code
func (o *ExtAggregateFailedImagesByRulesPathInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ext aggregate failed images by rules path internal server error response has a 5xx status code
func (o *ExtAggregateFailedImagesByRulesPathInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ext aggregate failed images by rules path internal server error response a status code equal to that given
func (o *ExtAggregateFailedImagesByRulesPathInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the ext aggregate failed images by rules path internal server error response
func (o *ExtAggregateFailedImagesByRulesPathInternalServerError) Code() int {
	return 500
}

func (o *ExtAggregateFailedImagesByRulesPathInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-images-by-rules/v2][%d] extAggregateFailedImagesByRulesPathInternalServerError  %+v", 500, o.Payload)
}

func (o *ExtAggregateFailedImagesByRulesPathInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-images-by-rules/v2][%d] extAggregateFailedImagesByRulesPathInternalServerError  %+v", 500, o.Payload)
}

func (o *ExtAggregateFailedImagesByRulesPathInternalServerError) GetPayload() *models.DomainAggregateFailedAssetsByRulesResponse {
	return o.Payload
}

func (o *ExtAggregateFailedImagesByRulesPathInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAggregateFailedAssetsByRulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
