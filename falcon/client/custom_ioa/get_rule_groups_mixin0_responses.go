// Code generated by go-swagger; DO NOT EDIT.

package custom_ioa

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

// GetRuleGroupsMixin0Reader is a Reader for the GetRuleGroupsMixin0 structure.
type GetRuleGroupsMixin0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRuleGroupsMixin0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRuleGroupsMixin0OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRuleGroupsMixin0Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRuleGroupsMixin0NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRuleGroupsMixin0TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /ioarules/entities/rule-groups/v1] get-rule-groupsMixin0", response, response.Code())
	}
}

// NewGetRuleGroupsMixin0OK creates a GetRuleGroupsMixin0OK with default headers values
func NewGetRuleGroupsMixin0OK() *GetRuleGroupsMixin0OK {
	return &GetRuleGroupsMixin0OK{}
}

/*
GetRuleGroupsMixin0OK describes a response with status code 200, with default header values.

OK
*/
type GetRuleGroupsMixin0OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIRuleGroupsResponse
}

// IsSuccess returns true when this get rule groups mixin0 o k response has a 2xx status code
func (o *GetRuleGroupsMixin0OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get rule groups mixin0 o k response has a 3xx status code
func (o *GetRuleGroupsMixin0OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rule groups mixin0 o k response has a 4xx status code
func (o *GetRuleGroupsMixin0OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get rule groups mixin0 o k response has a 5xx status code
func (o *GetRuleGroupsMixin0OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get rule groups mixin0 o k response a status code equal to that given
func (o *GetRuleGroupsMixin0OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get rule groups mixin0 o k response
func (o *GetRuleGroupsMixin0OK) Code() int {
	return 200
}

func (o *GetRuleGroupsMixin0OK) Error() string {
	return fmt.Sprintf("[GET /ioarules/entities/rule-groups/v1][%d] getRuleGroupsMixin0OK  %+v", 200, o.Payload)
}

func (o *GetRuleGroupsMixin0OK) String() string {
	return fmt.Sprintf("[GET /ioarules/entities/rule-groups/v1][%d] getRuleGroupsMixin0OK  %+v", 200, o.Payload)
}

func (o *GetRuleGroupsMixin0OK) GetPayload() *models.APIRuleGroupsResponse {
	return o.Payload
}

func (o *GetRuleGroupsMixin0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIRuleGroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuleGroupsMixin0Forbidden creates a GetRuleGroupsMixin0Forbidden with default headers values
func NewGetRuleGroupsMixin0Forbidden() *GetRuleGroupsMixin0Forbidden {
	return &GetRuleGroupsMixin0Forbidden{}
}

/*
GetRuleGroupsMixin0Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRuleGroupsMixin0Forbidden struct {

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

// IsSuccess returns true when this get rule groups mixin0 forbidden response has a 2xx status code
func (o *GetRuleGroupsMixin0Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rule groups mixin0 forbidden response has a 3xx status code
func (o *GetRuleGroupsMixin0Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rule groups mixin0 forbidden response has a 4xx status code
func (o *GetRuleGroupsMixin0Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rule groups mixin0 forbidden response has a 5xx status code
func (o *GetRuleGroupsMixin0Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get rule groups mixin0 forbidden response a status code equal to that given
func (o *GetRuleGroupsMixin0Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get rule groups mixin0 forbidden response
func (o *GetRuleGroupsMixin0Forbidden) Code() int {
	return 403
}

func (o *GetRuleGroupsMixin0Forbidden) Error() string {
	return fmt.Sprintf("[GET /ioarules/entities/rule-groups/v1][%d] getRuleGroupsMixin0Forbidden  %+v", 403, o.Payload)
}

func (o *GetRuleGroupsMixin0Forbidden) String() string {
	return fmt.Sprintf("[GET /ioarules/entities/rule-groups/v1][%d] getRuleGroupsMixin0Forbidden  %+v", 403, o.Payload)
}

func (o *GetRuleGroupsMixin0Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetRuleGroupsMixin0Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRuleGroupsMixin0NotFound creates a GetRuleGroupsMixin0NotFound with default headers values
func NewGetRuleGroupsMixin0NotFound() *GetRuleGroupsMixin0NotFound {
	return &GetRuleGroupsMixin0NotFound{}
}

/*
GetRuleGroupsMixin0NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetRuleGroupsMixin0NotFound struct {

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

// IsSuccess returns true when this get rule groups mixin0 not found response has a 2xx status code
func (o *GetRuleGroupsMixin0NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rule groups mixin0 not found response has a 3xx status code
func (o *GetRuleGroupsMixin0NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rule groups mixin0 not found response has a 4xx status code
func (o *GetRuleGroupsMixin0NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rule groups mixin0 not found response has a 5xx status code
func (o *GetRuleGroupsMixin0NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get rule groups mixin0 not found response a status code equal to that given
func (o *GetRuleGroupsMixin0NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get rule groups mixin0 not found response
func (o *GetRuleGroupsMixin0NotFound) Code() int {
	return 404
}

func (o *GetRuleGroupsMixin0NotFound) Error() string {
	return fmt.Sprintf("[GET /ioarules/entities/rule-groups/v1][%d] getRuleGroupsMixin0NotFound  %+v", 404, o.Payload)
}

func (o *GetRuleGroupsMixin0NotFound) String() string {
	return fmt.Sprintf("[GET /ioarules/entities/rule-groups/v1][%d] getRuleGroupsMixin0NotFound  %+v", 404, o.Payload)
}

func (o *GetRuleGroupsMixin0NotFound) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetRuleGroupsMixin0NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRuleGroupsMixin0TooManyRequests creates a GetRuleGroupsMixin0TooManyRequests with default headers values
func NewGetRuleGroupsMixin0TooManyRequests() *GetRuleGroupsMixin0TooManyRequests {
	return &GetRuleGroupsMixin0TooManyRequests{}
}

/*
GetRuleGroupsMixin0TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetRuleGroupsMixin0TooManyRequests struct {

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

// IsSuccess returns true when this get rule groups mixin0 too many requests response has a 2xx status code
func (o *GetRuleGroupsMixin0TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rule groups mixin0 too many requests response has a 3xx status code
func (o *GetRuleGroupsMixin0TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rule groups mixin0 too many requests response has a 4xx status code
func (o *GetRuleGroupsMixin0TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rule groups mixin0 too many requests response has a 5xx status code
func (o *GetRuleGroupsMixin0TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get rule groups mixin0 too many requests response a status code equal to that given
func (o *GetRuleGroupsMixin0TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get rule groups mixin0 too many requests response
func (o *GetRuleGroupsMixin0TooManyRequests) Code() int {
	return 429
}

func (o *GetRuleGroupsMixin0TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /ioarules/entities/rule-groups/v1][%d] getRuleGroupsMixin0TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRuleGroupsMixin0TooManyRequests) String() string {
	return fmt.Sprintf("[GET /ioarules/entities/rule-groups/v1][%d] getRuleGroupsMixin0TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRuleGroupsMixin0TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetRuleGroupsMixin0TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
