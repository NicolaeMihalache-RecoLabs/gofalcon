// Code generated by go-swagger; DO NOT EDIT.

package firewall_policies

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

// QueryCombinedFirewallPoliciesReader is a Reader for the QueryCombinedFirewallPolicies structure.
type QueryCombinedFirewallPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryCombinedFirewallPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryCombinedFirewallPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryCombinedFirewallPoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryCombinedFirewallPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryCombinedFirewallPoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryCombinedFirewallPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /policy/combined/firewall/v1] queryCombinedFirewallPolicies", response, response.Code())
	}
}

// NewQueryCombinedFirewallPoliciesOK creates a QueryCombinedFirewallPoliciesOK with default headers values
func NewQueryCombinedFirewallPoliciesOK() *QueryCombinedFirewallPoliciesOK {
	return &QueryCombinedFirewallPoliciesOK{}
}

/*
QueryCombinedFirewallPoliciesOK describes a response with status code 200, with default header values.

OK
*/
type QueryCombinedFirewallPoliciesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FirewallRespV1
}

// IsSuccess returns true when this query combined firewall policies o k response has a 2xx status code
func (o *QueryCombinedFirewallPoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query combined firewall policies o k response has a 3xx status code
func (o *QueryCombinedFirewallPoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined firewall policies o k response has a 4xx status code
func (o *QueryCombinedFirewallPoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query combined firewall policies o k response has a 5xx status code
func (o *QueryCombinedFirewallPoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined firewall policies o k response a status code equal to that given
func (o *QueryCombinedFirewallPoliciesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query combined firewall policies o k response
func (o *QueryCombinedFirewallPoliciesOK) Code() int {
	return 200
}

func (o *QueryCombinedFirewallPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /policy/combined/firewall/v1][%d] queryCombinedFirewallPoliciesOK  %+v", 200, o.Payload)
}

func (o *QueryCombinedFirewallPoliciesOK) String() string {
	return fmt.Sprintf("[GET /policy/combined/firewall/v1][%d] queryCombinedFirewallPoliciesOK  %+v", 200, o.Payload)
}

func (o *QueryCombinedFirewallPoliciesOK) GetPayload() *models.FirewallRespV1 {
	return o.Payload
}

func (o *QueryCombinedFirewallPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FirewallRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedFirewallPoliciesBadRequest creates a QueryCombinedFirewallPoliciesBadRequest with default headers values
func NewQueryCombinedFirewallPoliciesBadRequest() *QueryCombinedFirewallPoliciesBadRequest {
	return &QueryCombinedFirewallPoliciesBadRequest{}
}

/*
QueryCombinedFirewallPoliciesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryCombinedFirewallPoliciesBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FirewallRespV1
}

// IsSuccess returns true when this query combined firewall policies bad request response has a 2xx status code
func (o *QueryCombinedFirewallPoliciesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined firewall policies bad request response has a 3xx status code
func (o *QueryCombinedFirewallPoliciesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined firewall policies bad request response has a 4xx status code
func (o *QueryCombinedFirewallPoliciesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined firewall policies bad request response has a 5xx status code
func (o *QueryCombinedFirewallPoliciesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined firewall policies bad request response a status code equal to that given
func (o *QueryCombinedFirewallPoliciesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query combined firewall policies bad request response
func (o *QueryCombinedFirewallPoliciesBadRequest) Code() int {
	return 400
}

func (o *QueryCombinedFirewallPoliciesBadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/combined/firewall/v1][%d] queryCombinedFirewallPoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *QueryCombinedFirewallPoliciesBadRequest) String() string {
	return fmt.Sprintf("[GET /policy/combined/firewall/v1][%d] queryCombinedFirewallPoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *QueryCombinedFirewallPoliciesBadRequest) GetPayload() *models.FirewallRespV1 {
	return o.Payload
}

func (o *QueryCombinedFirewallPoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FirewallRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedFirewallPoliciesForbidden creates a QueryCombinedFirewallPoliciesForbidden with default headers values
func NewQueryCombinedFirewallPoliciesForbidden() *QueryCombinedFirewallPoliciesForbidden {
	return &QueryCombinedFirewallPoliciesForbidden{}
}

/*
QueryCombinedFirewallPoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryCombinedFirewallPoliciesForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this query combined firewall policies forbidden response has a 2xx status code
func (o *QueryCombinedFirewallPoliciesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined firewall policies forbidden response has a 3xx status code
func (o *QueryCombinedFirewallPoliciesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined firewall policies forbidden response has a 4xx status code
func (o *QueryCombinedFirewallPoliciesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined firewall policies forbidden response has a 5xx status code
func (o *QueryCombinedFirewallPoliciesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined firewall policies forbidden response a status code equal to that given
func (o *QueryCombinedFirewallPoliciesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query combined firewall policies forbidden response
func (o *QueryCombinedFirewallPoliciesForbidden) Code() int {
	return 403
}

func (o *QueryCombinedFirewallPoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/combined/firewall/v1][%d] queryCombinedFirewallPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *QueryCombinedFirewallPoliciesForbidden) String() string {
	return fmt.Sprintf("[GET /policy/combined/firewall/v1][%d] queryCombinedFirewallPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *QueryCombinedFirewallPoliciesForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryCombinedFirewallPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedFirewallPoliciesTooManyRequests creates a QueryCombinedFirewallPoliciesTooManyRequests with default headers values
func NewQueryCombinedFirewallPoliciesTooManyRequests() *QueryCombinedFirewallPoliciesTooManyRequests {
	return &QueryCombinedFirewallPoliciesTooManyRequests{}
}

/*
QueryCombinedFirewallPoliciesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryCombinedFirewallPoliciesTooManyRequests struct {

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

// IsSuccess returns true when this query combined firewall policies too many requests response has a 2xx status code
func (o *QueryCombinedFirewallPoliciesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined firewall policies too many requests response has a 3xx status code
func (o *QueryCombinedFirewallPoliciesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined firewall policies too many requests response has a 4xx status code
func (o *QueryCombinedFirewallPoliciesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined firewall policies too many requests response has a 5xx status code
func (o *QueryCombinedFirewallPoliciesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined firewall policies too many requests response a status code equal to that given
func (o *QueryCombinedFirewallPoliciesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query combined firewall policies too many requests response
func (o *QueryCombinedFirewallPoliciesTooManyRequests) Code() int {
	return 429
}

func (o *QueryCombinedFirewallPoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/combined/firewall/v1][%d] queryCombinedFirewallPoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryCombinedFirewallPoliciesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /policy/combined/firewall/v1][%d] queryCombinedFirewallPoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryCombinedFirewallPoliciesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryCombinedFirewallPoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryCombinedFirewallPoliciesInternalServerError creates a QueryCombinedFirewallPoliciesInternalServerError with default headers values
func NewQueryCombinedFirewallPoliciesInternalServerError() *QueryCombinedFirewallPoliciesInternalServerError {
	return &QueryCombinedFirewallPoliciesInternalServerError{}
}

/*
QueryCombinedFirewallPoliciesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryCombinedFirewallPoliciesInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FirewallRespV1
}

// IsSuccess returns true when this query combined firewall policies internal server error response has a 2xx status code
func (o *QueryCombinedFirewallPoliciesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined firewall policies internal server error response has a 3xx status code
func (o *QueryCombinedFirewallPoliciesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined firewall policies internal server error response has a 4xx status code
func (o *QueryCombinedFirewallPoliciesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query combined firewall policies internal server error response has a 5xx status code
func (o *QueryCombinedFirewallPoliciesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query combined firewall policies internal server error response a status code equal to that given
func (o *QueryCombinedFirewallPoliciesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query combined firewall policies internal server error response
func (o *QueryCombinedFirewallPoliciesInternalServerError) Code() int {
	return 500
}

func (o *QueryCombinedFirewallPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/combined/firewall/v1][%d] queryCombinedFirewallPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryCombinedFirewallPoliciesInternalServerError) String() string {
	return fmt.Sprintf("[GET /policy/combined/firewall/v1][%d] queryCombinedFirewallPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryCombinedFirewallPoliciesInternalServerError) GetPayload() *models.FirewallRespV1 {
	return o.Payload
}

func (o *QueryCombinedFirewallPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FirewallRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
