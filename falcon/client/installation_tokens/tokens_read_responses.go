// Code generated by go-swagger; DO NOT EDIT.

package installation_tokens

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

// TokensReadReader is a Reader for the TokensRead structure.
type TokensReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TokensReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTokensReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTokensReadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTokensReadForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewTokensReadTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTokensReadInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /installation-tokens/entities/tokens/v1] tokens-read", response, response.Code())
	}
}

// NewTokensReadOK creates a TokensReadOK with default headers values
func NewTokensReadOK() *TokensReadOK {
	return &TokensReadOK{}
}

/*
TokensReadOK describes a response with status code 200, with default header values.

OK
*/
type TokensReadOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APITokenDetailsResponseV1
}

// IsSuccess returns true when this tokens read o k response has a 2xx status code
func (o *TokensReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tokens read o k response has a 3xx status code
func (o *TokensReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tokens read o k response has a 4xx status code
func (o *TokensReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tokens read o k response has a 5xx status code
func (o *TokensReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tokens read o k response a status code equal to that given
func (o *TokensReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tokens read o k response
func (o *TokensReadOK) Code() int {
	return 200
}

func (o *TokensReadOK) Error() string {
	return fmt.Sprintf("[GET /installation-tokens/entities/tokens/v1][%d] tokensReadOK  %+v", 200, o.Payload)
}

func (o *TokensReadOK) String() string {
	return fmt.Sprintf("[GET /installation-tokens/entities/tokens/v1][%d] tokensReadOK  %+v", 200, o.Payload)
}

func (o *TokensReadOK) GetPayload() *models.APITokenDetailsResponseV1 {
	return o.Payload
}

func (o *TokensReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APITokenDetailsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokensReadBadRequest creates a TokensReadBadRequest with default headers values
func NewTokensReadBadRequest() *TokensReadBadRequest {
	return &TokensReadBadRequest{}
}

/*
TokensReadBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TokensReadBadRequest struct {

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

// IsSuccess returns true when this tokens read bad request response has a 2xx status code
func (o *TokensReadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tokens read bad request response has a 3xx status code
func (o *TokensReadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tokens read bad request response has a 4xx status code
func (o *TokensReadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this tokens read bad request response has a 5xx status code
func (o *TokensReadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this tokens read bad request response a status code equal to that given
func (o *TokensReadBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the tokens read bad request response
func (o *TokensReadBadRequest) Code() int {
	return 400
}

func (o *TokensReadBadRequest) Error() string {
	return fmt.Sprintf("[GET /installation-tokens/entities/tokens/v1][%d] tokensReadBadRequest  %+v", 400, o.Payload)
}

func (o *TokensReadBadRequest) String() string {
	return fmt.Sprintf("[GET /installation-tokens/entities/tokens/v1][%d] tokensReadBadRequest  %+v", 400, o.Payload)
}

func (o *TokensReadBadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *TokensReadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewTokensReadForbidden creates a TokensReadForbidden with default headers values
func NewTokensReadForbidden() *TokensReadForbidden {
	return &TokensReadForbidden{}
}

/*
TokensReadForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TokensReadForbidden struct {

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

// IsSuccess returns true when this tokens read forbidden response has a 2xx status code
func (o *TokensReadForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tokens read forbidden response has a 3xx status code
func (o *TokensReadForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tokens read forbidden response has a 4xx status code
func (o *TokensReadForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this tokens read forbidden response has a 5xx status code
func (o *TokensReadForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this tokens read forbidden response a status code equal to that given
func (o *TokensReadForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the tokens read forbidden response
func (o *TokensReadForbidden) Code() int {
	return 403
}

func (o *TokensReadForbidden) Error() string {
	return fmt.Sprintf("[GET /installation-tokens/entities/tokens/v1][%d] tokensReadForbidden  %+v", 403, o.Payload)
}

func (o *TokensReadForbidden) String() string {
	return fmt.Sprintf("[GET /installation-tokens/entities/tokens/v1][%d] tokensReadForbidden  %+v", 403, o.Payload)
}

func (o *TokensReadForbidden) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *TokensReadForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewTokensReadTooManyRequests creates a TokensReadTooManyRequests with default headers values
func NewTokensReadTooManyRequests() *TokensReadTooManyRequests {
	return &TokensReadTooManyRequests{}
}

/*
TokensReadTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type TokensReadTooManyRequests struct {

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

// IsSuccess returns true when this tokens read too many requests response has a 2xx status code
func (o *TokensReadTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tokens read too many requests response has a 3xx status code
func (o *TokensReadTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tokens read too many requests response has a 4xx status code
func (o *TokensReadTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this tokens read too many requests response has a 5xx status code
func (o *TokensReadTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this tokens read too many requests response a status code equal to that given
func (o *TokensReadTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the tokens read too many requests response
func (o *TokensReadTooManyRequests) Code() int {
	return 429
}

func (o *TokensReadTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /installation-tokens/entities/tokens/v1][%d] tokensReadTooManyRequests  %+v", 429, o.Payload)
}

func (o *TokensReadTooManyRequests) String() string {
	return fmt.Sprintf("[GET /installation-tokens/entities/tokens/v1][%d] tokensReadTooManyRequests  %+v", 429, o.Payload)
}

func (o *TokensReadTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *TokensReadTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewTokensReadInternalServerError creates a TokensReadInternalServerError with default headers values
func NewTokensReadInternalServerError() *TokensReadInternalServerError {
	return &TokensReadInternalServerError{}
}

/*
TokensReadInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type TokensReadInternalServerError struct {

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

// IsSuccess returns true when this tokens read internal server error response has a 2xx status code
func (o *TokensReadInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tokens read internal server error response has a 3xx status code
func (o *TokensReadInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tokens read internal server error response has a 4xx status code
func (o *TokensReadInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this tokens read internal server error response has a 5xx status code
func (o *TokensReadInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this tokens read internal server error response a status code equal to that given
func (o *TokensReadInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the tokens read internal server error response
func (o *TokensReadInternalServerError) Code() int {
	return 500
}

func (o *TokensReadInternalServerError) Error() string {
	return fmt.Sprintf("[GET /installation-tokens/entities/tokens/v1][%d] tokensReadInternalServerError  %+v", 500, o.Payload)
}

func (o *TokensReadInternalServerError) String() string {
	return fmt.Sprintf("[GET /installation-tokens/entities/tokens/v1][%d] tokensReadInternalServerError  %+v", 500, o.Payload)
}

func (o *TokensReadInternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *TokensReadInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
