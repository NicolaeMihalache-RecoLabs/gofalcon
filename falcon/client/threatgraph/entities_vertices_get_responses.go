// Code generated by go-swagger; DO NOT EDIT.

package threatgraph

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

// EntitiesVerticesGetReader is a Reader for the EntitiesVerticesGet structure.
type EntitiesVerticesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EntitiesVerticesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEntitiesVerticesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEntitiesVerticesGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEntitiesVerticesGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEntitiesVerticesGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEntitiesVerticesGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewEntitiesVerticesGetGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewEntitiesVerticesGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEntitiesVerticesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewEntitiesVerticesGetBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /threatgraph/entities/{vertex-type}/v1] entities_vertices_get", response, response.Code())
	}
}

// NewEntitiesVerticesGetOK creates a EntitiesVerticesGetOK with default headers values
func NewEntitiesVerticesGetOK() *EntitiesVerticesGetOK {
	return &EntitiesVerticesGetOK{}
}

/*
EntitiesVerticesGetOK describes a response with status code 200, with default header values.

OK
*/
type EntitiesVerticesGetOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ThreatgraphVertexDetailsResponse
}

// IsSuccess returns true when this entities vertices get o k response has a 2xx status code
func (o *EntitiesVerticesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this entities vertices get o k response has a 3xx status code
func (o *EntitiesVerticesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this entities vertices get o k response has a 4xx status code
func (o *EntitiesVerticesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this entities vertices get o k response has a 5xx status code
func (o *EntitiesVerticesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this entities vertices get o k response a status code equal to that given
func (o *EntitiesVerticesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the entities vertices get o k response
func (o *EntitiesVerticesGetOK) Code() int {
	return 200
}

func (o *EntitiesVerticesGetOK) Error() string {
	return fmt.Sprintf("[GET /threatgraph/entities/{vertex-type}/v1][%d] entitiesVerticesGetOK  %+v", 200, o.Payload)
}

func (o *EntitiesVerticesGetOK) String() string {
	return fmt.Sprintf("[GET /threatgraph/entities/{vertex-type}/v1][%d] entitiesVerticesGetOK  %+v", 200, o.Payload)
}

func (o *EntitiesVerticesGetOK) GetPayload() *models.ThreatgraphVertexDetailsResponse {
	return o.Payload
}

func (o *EntitiesVerticesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ThreatgraphVertexDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEntitiesVerticesGetBadRequest creates a EntitiesVerticesGetBadRequest with default headers values
func NewEntitiesVerticesGetBadRequest() *EntitiesVerticesGetBadRequest {
	return &EntitiesVerticesGetBadRequest{}
}

/*
EntitiesVerticesGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type EntitiesVerticesGetBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ThreatgraphVertexDetailsResponse
}

// IsSuccess returns true when this entities vertices get bad request response has a 2xx status code
func (o *EntitiesVerticesGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this entities vertices get bad request response has a 3xx status code
func (o *EntitiesVerticesGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this entities vertices get bad request response has a 4xx status code
func (o *EntitiesVerticesGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this entities vertices get bad request response has a 5xx status code
func (o *EntitiesVerticesGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this entities vertices get bad request response a status code equal to that given
func (o *EntitiesVerticesGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the entities vertices get bad request response
func (o *EntitiesVerticesGetBadRequest) Code() int {
	return 400
}

func (o *EntitiesVerticesGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /threatgraph/entities/{vertex-type}/v1][%d] entitiesVerticesGetBadRequest  %+v", 400, o.Payload)
}

func (o *EntitiesVerticesGetBadRequest) String() string {
	return fmt.Sprintf("[GET /threatgraph/entities/{vertex-type}/v1][%d] entitiesVerticesGetBadRequest  %+v", 400, o.Payload)
}

func (o *EntitiesVerticesGetBadRequest) GetPayload() *models.ThreatgraphVertexDetailsResponse {
	return o.Payload
}

func (o *EntitiesVerticesGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ThreatgraphVertexDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEntitiesVerticesGetUnauthorized creates a EntitiesVerticesGetUnauthorized with default headers values
func NewEntitiesVerticesGetUnauthorized() *EntitiesVerticesGetUnauthorized {
	return &EntitiesVerticesGetUnauthorized{}
}

/*
EntitiesVerticesGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type EntitiesVerticesGetUnauthorized struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ThreatgraphVertexDetailsResponse
}

// IsSuccess returns true when this entities vertices get unauthorized response has a 2xx status code
func (o *EntitiesVerticesGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this entities vertices get unauthorized response has a 3xx status code
func (o *EntitiesVerticesGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this entities vertices get unauthorized response has a 4xx status code
func (o *EntitiesVerticesGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this entities vertices get unauthorized response has a 5xx status code
func (o *EntitiesVerticesGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this entities vertices get unauthorized response a status code equal to that given
func (o *EntitiesVerticesGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the entities vertices get unauthorized response
func (o *EntitiesVerticesGetUnauthorized) Code() int {
	return 401
}

func (o *EntitiesVerticesGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /threatgraph/entities/{vertex-type}/v1][%d] entitiesVerticesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *EntitiesVerticesGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /threatgraph/entities/{vertex-type}/v1][%d] entitiesVerticesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *EntitiesVerticesGetUnauthorized) GetPayload() *models.ThreatgraphVertexDetailsResponse {
	return o.Payload
}

func (o *EntitiesVerticesGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ThreatgraphVertexDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEntitiesVerticesGetForbidden creates a EntitiesVerticesGetForbidden with default headers values
func NewEntitiesVerticesGetForbidden() *EntitiesVerticesGetForbidden {
	return &EntitiesVerticesGetForbidden{}
}

/*
EntitiesVerticesGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type EntitiesVerticesGetForbidden struct {

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

// IsSuccess returns true when this entities vertices get forbidden response has a 2xx status code
func (o *EntitiesVerticesGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this entities vertices get forbidden response has a 3xx status code
func (o *EntitiesVerticesGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this entities vertices get forbidden response has a 4xx status code
func (o *EntitiesVerticesGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this entities vertices get forbidden response has a 5xx status code
func (o *EntitiesVerticesGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this entities vertices get forbidden response a status code equal to that given
func (o *EntitiesVerticesGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the entities vertices get forbidden response
func (o *EntitiesVerticesGetForbidden) Code() int {
	return 403
}

func (o *EntitiesVerticesGetForbidden) Error() string {
	return fmt.Sprintf("[GET /threatgraph/entities/{vertex-type}/v1][%d] entitiesVerticesGetForbidden  %+v", 403, o.Payload)
}

func (o *EntitiesVerticesGetForbidden) String() string {
	return fmt.Sprintf("[GET /threatgraph/entities/{vertex-type}/v1][%d] entitiesVerticesGetForbidden  %+v", 403, o.Payload)
}

func (o *EntitiesVerticesGetForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *EntitiesVerticesGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewEntitiesVerticesGetNotFound creates a EntitiesVerticesGetNotFound with default headers values
func NewEntitiesVerticesGetNotFound() *EntitiesVerticesGetNotFound {
	return &EntitiesVerticesGetNotFound{}
}

/*
EntitiesVerticesGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type EntitiesVerticesGetNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ThreatgraphVertexDetailsResponse
}

// IsSuccess returns true when this entities vertices get not found response has a 2xx status code
func (o *EntitiesVerticesGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this entities vertices get not found response has a 3xx status code
func (o *EntitiesVerticesGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this entities vertices get not found response has a 4xx status code
func (o *EntitiesVerticesGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this entities vertices get not found response has a 5xx status code
func (o *EntitiesVerticesGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this entities vertices get not found response a status code equal to that given
func (o *EntitiesVerticesGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the entities vertices get not found response
func (o *EntitiesVerticesGetNotFound) Code() int {
	return 404
}

func (o *EntitiesVerticesGetNotFound) Error() string {
	return fmt.Sprintf("[GET /threatgraph/entities/{vertex-type}/v1][%d] entitiesVerticesGetNotFound  %+v", 404, o.Payload)
}

func (o *EntitiesVerticesGetNotFound) String() string {
	return fmt.Sprintf("[GET /threatgraph/entities/{vertex-type}/v1][%d] entitiesVerticesGetNotFound  %+v", 404, o.Payload)
}

func (o *EntitiesVerticesGetNotFound) GetPayload() *models.ThreatgraphVertexDetailsResponse {
	return o.Payload
}

func (o *EntitiesVerticesGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ThreatgraphVertexDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEntitiesVerticesGetGone creates a EntitiesVerticesGetGone with default headers values
func NewEntitiesVerticesGetGone() *EntitiesVerticesGetGone {
	return &EntitiesVerticesGetGone{}
}

/*
EntitiesVerticesGetGone describes a response with status code 410, with default header values.

Gone
*/
type EntitiesVerticesGetGone struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ThreatgraphVertexDetailsResponse
}

// IsSuccess returns true when this entities vertices get gone response has a 2xx status code
func (o *EntitiesVerticesGetGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this entities vertices get gone response has a 3xx status code
func (o *EntitiesVerticesGetGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this entities vertices get gone response has a 4xx status code
func (o *EntitiesVerticesGetGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this entities vertices get gone response has a 5xx status code
func (o *EntitiesVerticesGetGone) IsServerError() bool {
	return false
}

// IsCode returns true when this entities vertices get gone response a status code equal to that given
func (o *EntitiesVerticesGetGone) IsCode(code int) bool {
	return code == 410
}

// Code gets the status code for the entities vertices get gone response
func (o *EntitiesVerticesGetGone) Code() int {
	return 410
}

func (o *EntitiesVerticesGetGone) Error() string {
	return fmt.Sprintf("[GET /threatgraph/entities/{vertex-type}/v1][%d] entitiesVerticesGetGone  %+v", 410, o.Payload)
}

func (o *EntitiesVerticesGetGone) String() string {
	return fmt.Sprintf("[GET /threatgraph/entities/{vertex-type}/v1][%d] entitiesVerticesGetGone  %+v", 410, o.Payload)
}

func (o *EntitiesVerticesGetGone) GetPayload() *models.ThreatgraphVertexDetailsResponse {
	return o.Payload
}

func (o *EntitiesVerticesGetGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ThreatgraphVertexDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEntitiesVerticesGetTooManyRequests creates a EntitiesVerticesGetTooManyRequests with default headers values
func NewEntitiesVerticesGetTooManyRequests() *EntitiesVerticesGetTooManyRequests {
	return &EntitiesVerticesGetTooManyRequests{}
}

/*
EntitiesVerticesGetTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type EntitiesVerticesGetTooManyRequests struct {

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

// IsSuccess returns true when this entities vertices get too many requests response has a 2xx status code
func (o *EntitiesVerticesGetTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this entities vertices get too many requests response has a 3xx status code
func (o *EntitiesVerticesGetTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this entities vertices get too many requests response has a 4xx status code
func (o *EntitiesVerticesGetTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this entities vertices get too many requests response has a 5xx status code
func (o *EntitiesVerticesGetTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this entities vertices get too many requests response a status code equal to that given
func (o *EntitiesVerticesGetTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the entities vertices get too many requests response
func (o *EntitiesVerticesGetTooManyRequests) Code() int {
	return 429
}

func (o *EntitiesVerticesGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /threatgraph/entities/{vertex-type}/v1][%d] entitiesVerticesGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *EntitiesVerticesGetTooManyRequests) String() string {
	return fmt.Sprintf("[GET /threatgraph/entities/{vertex-type}/v1][%d] entitiesVerticesGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *EntitiesVerticesGetTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *EntitiesVerticesGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewEntitiesVerticesGetInternalServerError creates a EntitiesVerticesGetInternalServerError with default headers values
func NewEntitiesVerticesGetInternalServerError() *EntitiesVerticesGetInternalServerError {
	return &EntitiesVerticesGetInternalServerError{}
}

/*
EntitiesVerticesGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type EntitiesVerticesGetInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ThreatgraphVertexDetailsResponse
}

// IsSuccess returns true when this entities vertices get internal server error response has a 2xx status code
func (o *EntitiesVerticesGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this entities vertices get internal server error response has a 3xx status code
func (o *EntitiesVerticesGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this entities vertices get internal server error response has a 4xx status code
func (o *EntitiesVerticesGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this entities vertices get internal server error response has a 5xx status code
func (o *EntitiesVerticesGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this entities vertices get internal server error response a status code equal to that given
func (o *EntitiesVerticesGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the entities vertices get internal server error response
func (o *EntitiesVerticesGetInternalServerError) Code() int {
	return 500
}

func (o *EntitiesVerticesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /threatgraph/entities/{vertex-type}/v1][%d] entitiesVerticesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *EntitiesVerticesGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /threatgraph/entities/{vertex-type}/v1][%d] entitiesVerticesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *EntitiesVerticesGetInternalServerError) GetPayload() *models.ThreatgraphVertexDetailsResponse {
	return o.Payload
}

func (o *EntitiesVerticesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ThreatgraphVertexDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEntitiesVerticesGetBadGateway creates a EntitiesVerticesGetBadGateway with default headers values
func NewEntitiesVerticesGetBadGateway() *EntitiesVerticesGetBadGateway {
	return &EntitiesVerticesGetBadGateway{}
}

/*
EntitiesVerticesGetBadGateway describes a response with status code 502, with default header values.

Bad Gateway
*/
type EntitiesVerticesGetBadGateway struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ThreatgraphVertexDetailsResponse
}

// IsSuccess returns true when this entities vertices get bad gateway response has a 2xx status code
func (o *EntitiesVerticesGetBadGateway) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this entities vertices get bad gateway response has a 3xx status code
func (o *EntitiesVerticesGetBadGateway) IsRedirect() bool {
	return false
}

// IsClientError returns true when this entities vertices get bad gateway response has a 4xx status code
func (o *EntitiesVerticesGetBadGateway) IsClientError() bool {
	return false
}

// IsServerError returns true when this entities vertices get bad gateway response has a 5xx status code
func (o *EntitiesVerticesGetBadGateway) IsServerError() bool {
	return true
}

// IsCode returns true when this entities vertices get bad gateway response a status code equal to that given
func (o *EntitiesVerticesGetBadGateway) IsCode(code int) bool {
	return code == 502
}

// Code gets the status code for the entities vertices get bad gateway response
func (o *EntitiesVerticesGetBadGateway) Code() int {
	return 502
}

func (o *EntitiesVerticesGetBadGateway) Error() string {
	return fmt.Sprintf("[GET /threatgraph/entities/{vertex-type}/v1][%d] entitiesVerticesGetBadGateway  %+v", 502, o.Payload)
}

func (o *EntitiesVerticesGetBadGateway) String() string {
	return fmt.Sprintf("[GET /threatgraph/entities/{vertex-type}/v1][%d] entitiesVerticesGetBadGateway  %+v", 502, o.Payload)
}

func (o *EntitiesVerticesGetBadGateway) GetPayload() *models.ThreatgraphVertexDetailsResponse {
	return o.Payload
}

func (o *EntitiesVerticesGetBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ThreatgraphVertexDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
