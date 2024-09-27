// Code generated by go-swagger; DO NOT EDIT.

package cloud_snapshots

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

// ReadDeploymentsEntitiesReader is a Reader for the ReadDeploymentsEntities structure.
type ReadDeploymentsEntitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadDeploymentsEntitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadDeploymentsEntitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReadDeploymentsEntitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReadDeploymentsEntitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewReadDeploymentsEntitiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadDeploymentsEntitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /snapshots/entities/deployments/v1] ReadDeploymentsEntities", response, response.Code())
	}
}

// NewReadDeploymentsEntitiesOK creates a ReadDeploymentsEntitiesOK with default headers values
func NewReadDeploymentsEntitiesOK() *ReadDeploymentsEntitiesOK {
	return &ReadDeploymentsEntitiesOK{}
}

/*
ReadDeploymentsEntitiesOK describes a response with status code 200, with default header values.

OK
*/
type ReadDeploymentsEntitiesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DeploymentsEntityResponse
}

// IsSuccess returns true when this read deployments entities o k response has a 2xx status code
func (o *ReadDeploymentsEntitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read deployments entities o k response has a 3xx status code
func (o *ReadDeploymentsEntitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read deployments entities o k response has a 4xx status code
func (o *ReadDeploymentsEntitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read deployments entities o k response has a 5xx status code
func (o *ReadDeploymentsEntitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read deployments entities o k response a status code equal to that given
func (o *ReadDeploymentsEntitiesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read deployments entities o k response
func (o *ReadDeploymentsEntitiesOK) Code() int {
	return 200
}

func (o *ReadDeploymentsEntitiesOK) Error() string {
	return fmt.Sprintf("[GET /snapshots/entities/deployments/v1][%d] readDeploymentsEntitiesOK  %+v", 200, o.Payload)
}

func (o *ReadDeploymentsEntitiesOK) String() string {
	return fmt.Sprintf("[GET /snapshots/entities/deployments/v1][%d] readDeploymentsEntitiesOK  %+v", 200, o.Payload)
}

func (o *ReadDeploymentsEntitiesOK) GetPayload() *models.DeploymentsEntityResponse {
	return o.Payload
}

func (o *ReadDeploymentsEntitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DeploymentsEntityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadDeploymentsEntitiesBadRequest creates a ReadDeploymentsEntitiesBadRequest with default headers values
func NewReadDeploymentsEntitiesBadRequest() *ReadDeploymentsEntitiesBadRequest {
	return &ReadDeploymentsEntitiesBadRequest{}
}

/*
ReadDeploymentsEntitiesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ReadDeploymentsEntitiesBadRequest struct {

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

// IsSuccess returns true when this read deployments entities bad request response has a 2xx status code
func (o *ReadDeploymentsEntitiesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read deployments entities bad request response has a 3xx status code
func (o *ReadDeploymentsEntitiesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read deployments entities bad request response has a 4xx status code
func (o *ReadDeploymentsEntitiesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this read deployments entities bad request response has a 5xx status code
func (o *ReadDeploymentsEntitiesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this read deployments entities bad request response a status code equal to that given
func (o *ReadDeploymentsEntitiesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the read deployments entities bad request response
func (o *ReadDeploymentsEntitiesBadRequest) Code() int {
	return 400
}

func (o *ReadDeploymentsEntitiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /snapshots/entities/deployments/v1][%d] readDeploymentsEntitiesBadRequest  %+v", 400, o.Payload)
}

func (o *ReadDeploymentsEntitiesBadRequest) String() string {
	return fmt.Sprintf("[GET /snapshots/entities/deployments/v1][%d] readDeploymentsEntitiesBadRequest  %+v", 400, o.Payload)
}

func (o *ReadDeploymentsEntitiesBadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *ReadDeploymentsEntitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReadDeploymentsEntitiesForbidden creates a ReadDeploymentsEntitiesForbidden with default headers values
func NewReadDeploymentsEntitiesForbidden() *ReadDeploymentsEntitiesForbidden {
	return &ReadDeploymentsEntitiesForbidden{}
}

/*
ReadDeploymentsEntitiesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ReadDeploymentsEntitiesForbidden struct {

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

// IsSuccess returns true when this read deployments entities forbidden response has a 2xx status code
func (o *ReadDeploymentsEntitiesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read deployments entities forbidden response has a 3xx status code
func (o *ReadDeploymentsEntitiesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read deployments entities forbidden response has a 4xx status code
func (o *ReadDeploymentsEntitiesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this read deployments entities forbidden response has a 5xx status code
func (o *ReadDeploymentsEntitiesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this read deployments entities forbidden response a status code equal to that given
func (o *ReadDeploymentsEntitiesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the read deployments entities forbidden response
func (o *ReadDeploymentsEntitiesForbidden) Code() int {
	return 403
}

func (o *ReadDeploymentsEntitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /snapshots/entities/deployments/v1][%d] readDeploymentsEntitiesForbidden  %+v", 403, o.Payload)
}

func (o *ReadDeploymentsEntitiesForbidden) String() string {
	return fmt.Sprintf("[GET /snapshots/entities/deployments/v1][%d] readDeploymentsEntitiesForbidden  %+v", 403, o.Payload)
}

func (o *ReadDeploymentsEntitiesForbidden) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *ReadDeploymentsEntitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReadDeploymentsEntitiesTooManyRequests creates a ReadDeploymentsEntitiesTooManyRequests with default headers values
func NewReadDeploymentsEntitiesTooManyRequests() *ReadDeploymentsEntitiesTooManyRequests {
	return &ReadDeploymentsEntitiesTooManyRequests{}
}

/*
ReadDeploymentsEntitiesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ReadDeploymentsEntitiesTooManyRequests struct {

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

// IsSuccess returns true when this read deployments entities too many requests response has a 2xx status code
func (o *ReadDeploymentsEntitiesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read deployments entities too many requests response has a 3xx status code
func (o *ReadDeploymentsEntitiesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read deployments entities too many requests response has a 4xx status code
func (o *ReadDeploymentsEntitiesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this read deployments entities too many requests response has a 5xx status code
func (o *ReadDeploymentsEntitiesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this read deployments entities too many requests response a status code equal to that given
func (o *ReadDeploymentsEntitiesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the read deployments entities too many requests response
func (o *ReadDeploymentsEntitiesTooManyRequests) Code() int {
	return 429
}

func (o *ReadDeploymentsEntitiesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /snapshots/entities/deployments/v1][%d] readDeploymentsEntitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *ReadDeploymentsEntitiesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /snapshots/entities/deployments/v1][%d] readDeploymentsEntitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *ReadDeploymentsEntitiesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ReadDeploymentsEntitiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReadDeploymentsEntitiesInternalServerError creates a ReadDeploymentsEntitiesInternalServerError with default headers values
func NewReadDeploymentsEntitiesInternalServerError() *ReadDeploymentsEntitiesInternalServerError {
	return &ReadDeploymentsEntitiesInternalServerError{}
}

/*
ReadDeploymentsEntitiesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ReadDeploymentsEntitiesInternalServerError struct {

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

// IsSuccess returns true when this read deployments entities internal server error response has a 2xx status code
func (o *ReadDeploymentsEntitiesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read deployments entities internal server error response has a 3xx status code
func (o *ReadDeploymentsEntitiesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read deployments entities internal server error response has a 4xx status code
func (o *ReadDeploymentsEntitiesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read deployments entities internal server error response has a 5xx status code
func (o *ReadDeploymentsEntitiesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read deployments entities internal server error response a status code equal to that given
func (o *ReadDeploymentsEntitiesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read deployments entities internal server error response
func (o *ReadDeploymentsEntitiesInternalServerError) Code() int {
	return 500
}

func (o *ReadDeploymentsEntitiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /snapshots/entities/deployments/v1][%d] readDeploymentsEntitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadDeploymentsEntitiesInternalServerError) String() string {
	return fmt.Sprintf("[GET /snapshots/entities/deployments/v1][%d] readDeploymentsEntitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadDeploymentsEntitiesInternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *ReadDeploymentsEntitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
