// Code generated by go-swagger; DO NOT EDIT.

package api_integrations

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

// GetCombinedPluginConfigsReader is a Reader for the GetCombinedPluginConfigs structure.
type GetCombinedPluginConfigsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCombinedPluginConfigsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCombinedPluginConfigsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCombinedPluginConfigsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCombinedPluginConfigsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCombinedPluginConfigsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCombinedPluginConfigsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCombinedPluginConfigsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /plugins/combined/configs/v1] GetCombinedPluginConfigs", response, response.Code())
	}
}

// NewGetCombinedPluginConfigsOK creates a GetCombinedPluginConfigsOK with default headers values
func NewGetCombinedPluginConfigsOK() *GetCombinedPluginConfigsOK {
	return &GetCombinedPluginConfigsOK{}
}

/*
GetCombinedPluginConfigsOK describes a response with status code 200, with default header values.

OK
*/
type GetCombinedPluginConfigsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainConfigsV1
}

// IsSuccess returns true when this get combined plugin configs o k response has a 2xx status code
func (o *GetCombinedPluginConfigsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get combined plugin configs o k response has a 3xx status code
func (o *GetCombinedPluginConfigsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get combined plugin configs o k response has a 4xx status code
func (o *GetCombinedPluginConfigsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get combined plugin configs o k response has a 5xx status code
func (o *GetCombinedPluginConfigsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get combined plugin configs o k response a status code equal to that given
func (o *GetCombinedPluginConfigsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get combined plugin configs o k response
func (o *GetCombinedPluginConfigsOK) Code() int {
	return 200
}

func (o *GetCombinedPluginConfigsOK) Error() string {
	return fmt.Sprintf("[GET /plugins/combined/configs/v1][%d] getCombinedPluginConfigsOK  %+v", 200, o.Payload)
}

func (o *GetCombinedPluginConfigsOK) String() string {
	return fmt.Sprintf("[GET /plugins/combined/configs/v1][%d] getCombinedPluginConfigsOK  %+v", 200, o.Payload)
}

func (o *GetCombinedPluginConfigsOK) GetPayload() *models.DomainConfigsV1 {
	return o.Payload
}

func (o *GetCombinedPluginConfigsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainConfigsV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCombinedPluginConfigsBadRequest creates a GetCombinedPluginConfigsBadRequest with default headers values
func NewGetCombinedPluginConfigsBadRequest() *GetCombinedPluginConfigsBadRequest {
	return &GetCombinedPluginConfigsBadRequest{}
}

/*
GetCombinedPluginConfigsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetCombinedPluginConfigsBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainConfigsV1
}

// IsSuccess returns true when this get combined plugin configs bad request response has a 2xx status code
func (o *GetCombinedPluginConfigsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get combined plugin configs bad request response has a 3xx status code
func (o *GetCombinedPluginConfigsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get combined plugin configs bad request response has a 4xx status code
func (o *GetCombinedPluginConfigsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get combined plugin configs bad request response has a 5xx status code
func (o *GetCombinedPluginConfigsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get combined plugin configs bad request response a status code equal to that given
func (o *GetCombinedPluginConfigsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get combined plugin configs bad request response
func (o *GetCombinedPluginConfigsBadRequest) Code() int {
	return 400
}

func (o *GetCombinedPluginConfigsBadRequest) Error() string {
	return fmt.Sprintf("[GET /plugins/combined/configs/v1][%d] getCombinedPluginConfigsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCombinedPluginConfigsBadRequest) String() string {
	return fmt.Sprintf("[GET /plugins/combined/configs/v1][%d] getCombinedPluginConfigsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCombinedPluginConfigsBadRequest) GetPayload() *models.DomainConfigsV1 {
	return o.Payload
}

func (o *GetCombinedPluginConfigsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainConfigsV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCombinedPluginConfigsForbidden creates a GetCombinedPluginConfigsForbidden with default headers values
func NewGetCombinedPluginConfigsForbidden() *GetCombinedPluginConfigsForbidden {
	return &GetCombinedPluginConfigsForbidden{}
}

/*
GetCombinedPluginConfigsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCombinedPluginConfigsForbidden struct {

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

// IsSuccess returns true when this get combined plugin configs forbidden response has a 2xx status code
func (o *GetCombinedPluginConfigsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get combined plugin configs forbidden response has a 3xx status code
func (o *GetCombinedPluginConfigsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get combined plugin configs forbidden response has a 4xx status code
func (o *GetCombinedPluginConfigsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get combined plugin configs forbidden response has a 5xx status code
func (o *GetCombinedPluginConfigsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get combined plugin configs forbidden response a status code equal to that given
func (o *GetCombinedPluginConfigsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get combined plugin configs forbidden response
func (o *GetCombinedPluginConfigsForbidden) Code() int {
	return 403
}

func (o *GetCombinedPluginConfigsForbidden) Error() string {
	return fmt.Sprintf("[GET /plugins/combined/configs/v1][%d] getCombinedPluginConfigsForbidden  %+v", 403, o.Payload)
}

func (o *GetCombinedPluginConfigsForbidden) String() string {
	return fmt.Sprintf("[GET /plugins/combined/configs/v1][%d] getCombinedPluginConfigsForbidden  %+v", 403, o.Payload)
}

func (o *GetCombinedPluginConfigsForbidden) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *GetCombinedPluginConfigsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCombinedPluginConfigsNotFound creates a GetCombinedPluginConfigsNotFound with default headers values
func NewGetCombinedPluginConfigsNotFound() *GetCombinedPluginConfigsNotFound {
	return &GetCombinedPluginConfigsNotFound{}
}

/*
GetCombinedPluginConfigsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetCombinedPluginConfigsNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainConfigsV1
}

// IsSuccess returns true when this get combined plugin configs not found response has a 2xx status code
func (o *GetCombinedPluginConfigsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get combined plugin configs not found response has a 3xx status code
func (o *GetCombinedPluginConfigsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get combined plugin configs not found response has a 4xx status code
func (o *GetCombinedPluginConfigsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get combined plugin configs not found response has a 5xx status code
func (o *GetCombinedPluginConfigsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get combined plugin configs not found response a status code equal to that given
func (o *GetCombinedPluginConfigsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get combined plugin configs not found response
func (o *GetCombinedPluginConfigsNotFound) Code() int {
	return 404
}

func (o *GetCombinedPluginConfigsNotFound) Error() string {
	return fmt.Sprintf("[GET /plugins/combined/configs/v1][%d] getCombinedPluginConfigsNotFound  %+v", 404, o.Payload)
}

func (o *GetCombinedPluginConfigsNotFound) String() string {
	return fmt.Sprintf("[GET /plugins/combined/configs/v1][%d] getCombinedPluginConfigsNotFound  %+v", 404, o.Payload)
}

func (o *GetCombinedPluginConfigsNotFound) GetPayload() *models.DomainConfigsV1 {
	return o.Payload
}

func (o *GetCombinedPluginConfigsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainConfigsV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCombinedPluginConfigsTooManyRequests creates a GetCombinedPluginConfigsTooManyRequests with default headers values
func NewGetCombinedPluginConfigsTooManyRequests() *GetCombinedPluginConfigsTooManyRequests {
	return &GetCombinedPluginConfigsTooManyRequests{}
}

/*
GetCombinedPluginConfigsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetCombinedPluginConfigsTooManyRequests struct {

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

// IsSuccess returns true when this get combined plugin configs too many requests response has a 2xx status code
func (o *GetCombinedPluginConfigsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get combined plugin configs too many requests response has a 3xx status code
func (o *GetCombinedPluginConfigsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get combined plugin configs too many requests response has a 4xx status code
func (o *GetCombinedPluginConfigsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get combined plugin configs too many requests response has a 5xx status code
func (o *GetCombinedPluginConfigsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get combined plugin configs too many requests response a status code equal to that given
func (o *GetCombinedPluginConfigsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get combined plugin configs too many requests response
func (o *GetCombinedPluginConfigsTooManyRequests) Code() int {
	return 429
}

func (o *GetCombinedPluginConfigsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /plugins/combined/configs/v1][%d] getCombinedPluginConfigsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCombinedPluginConfigsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /plugins/combined/configs/v1][%d] getCombinedPluginConfigsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCombinedPluginConfigsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCombinedPluginConfigsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCombinedPluginConfigsInternalServerError creates a GetCombinedPluginConfigsInternalServerError with default headers values
func NewGetCombinedPluginConfigsInternalServerError() *GetCombinedPluginConfigsInternalServerError {
	return &GetCombinedPluginConfigsInternalServerError{}
}

/*
GetCombinedPluginConfigsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetCombinedPluginConfigsInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainConfigsV1
}

// IsSuccess returns true when this get combined plugin configs internal server error response has a 2xx status code
func (o *GetCombinedPluginConfigsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get combined plugin configs internal server error response has a 3xx status code
func (o *GetCombinedPluginConfigsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get combined plugin configs internal server error response has a 4xx status code
func (o *GetCombinedPluginConfigsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get combined plugin configs internal server error response has a 5xx status code
func (o *GetCombinedPluginConfigsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get combined plugin configs internal server error response a status code equal to that given
func (o *GetCombinedPluginConfigsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get combined plugin configs internal server error response
func (o *GetCombinedPluginConfigsInternalServerError) Code() int {
	return 500
}

func (o *GetCombinedPluginConfigsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /plugins/combined/configs/v1][%d] getCombinedPluginConfigsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCombinedPluginConfigsInternalServerError) String() string {
	return fmt.Sprintf("[GET /plugins/combined/configs/v1][%d] getCombinedPluginConfigsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCombinedPluginConfigsInternalServerError) GetPayload() *models.DomainConfigsV1 {
	return o.Payload
}

func (o *GetCombinedPluginConfigsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainConfigsV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
