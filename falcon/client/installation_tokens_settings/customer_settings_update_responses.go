// Code generated by go-swagger; DO NOT EDIT.

package installation_tokens_settings

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

// CustomerSettingsUpdateReader is a Reader for the CustomerSettingsUpdate structure.
type CustomerSettingsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerSettingsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerSettingsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCustomerSettingsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCustomerSettingsUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCustomerSettingsUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCustomerSettingsUpdateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCustomerSettingsUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /installation-tokens/entities/customer-settings/v1] customer-settings-update", response, response.Code())
	}
}

// NewCustomerSettingsUpdateOK creates a CustomerSettingsUpdateOK with default headers values
func NewCustomerSettingsUpdateOK() *CustomerSettingsUpdateOK {
	return &CustomerSettingsUpdateOK{}
}

/*
CustomerSettingsUpdateOK describes a response with status code 200, with default header values.

OK
*/
type CustomerSettingsUpdateOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecQueryResponse
}

// IsSuccess returns true when this customer settings update o k response has a 2xx status code
func (o *CustomerSettingsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer settings update o k response has a 3xx status code
func (o *CustomerSettingsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer settings update o k response has a 4xx status code
func (o *CustomerSettingsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer settings update o k response has a 5xx status code
func (o *CustomerSettingsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer settings update o k response a status code equal to that given
func (o *CustomerSettingsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer settings update o k response
func (o *CustomerSettingsUpdateOK) Code() int {
	return 200
}

func (o *CustomerSettingsUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /installation-tokens/entities/customer-settings/v1][%d] customerSettingsUpdateOK  %+v", 200, o.Payload)
}

func (o *CustomerSettingsUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /installation-tokens/entities/customer-settings/v1][%d] customerSettingsUpdateOK  %+v", 200, o.Payload)
}

func (o *CustomerSettingsUpdateOK) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *CustomerSettingsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerSettingsUpdateBadRequest creates a CustomerSettingsUpdateBadRequest with default headers values
func NewCustomerSettingsUpdateBadRequest() *CustomerSettingsUpdateBadRequest {
	return &CustomerSettingsUpdateBadRequest{}
}

/*
CustomerSettingsUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CustomerSettingsUpdateBadRequest struct {

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

// IsSuccess returns true when this customer settings update bad request response has a 2xx status code
func (o *CustomerSettingsUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this customer settings update bad request response has a 3xx status code
func (o *CustomerSettingsUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer settings update bad request response has a 4xx status code
func (o *CustomerSettingsUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this customer settings update bad request response has a 5xx status code
func (o *CustomerSettingsUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this customer settings update bad request response a status code equal to that given
func (o *CustomerSettingsUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the customer settings update bad request response
func (o *CustomerSettingsUpdateBadRequest) Code() int {
	return 400
}

func (o *CustomerSettingsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /installation-tokens/entities/customer-settings/v1][%d] customerSettingsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *CustomerSettingsUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /installation-tokens/entities/customer-settings/v1][%d] customerSettingsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *CustomerSettingsUpdateBadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CustomerSettingsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCustomerSettingsUpdateForbidden creates a CustomerSettingsUpdateForbidden with default headers values
func NewCustomerSettingsUpdateForbidden() *CustomerSettingsUpdateForbidden {
	return &CustomerSettingsUpdateForbidden{}
}

/*
CustomerSettingsUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CustomerSettingsUpdateForbidden struct {

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

// IsSuccess returns true when this customer settings update forbidden response has a 2xx status code
func (o *CustomerSettingsUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this customer settings update forbidden response has a 3xx status code
func (o *CustomerSettingsUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer settings update forbidden response has a 4xx status code
func (o *CustomerSettingsUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this customer settings update forbidden response has a 5xx status code
func (o *CustomerSettingsUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this customer settings update forbidden response a status code equal to that given
func (o *CustomerSettingsUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the customer settings update forbidden response
func (o *CustomerSettingsUpdateForbidden) Code() int {
	return 403
}

func (o *CustomerSettingsUpdateForbidden) Error() string {
	return fmt.Sprintf("[PATCH /installation-tokens/entities/customer-settings/v1][%d] customerSettingsUpdateForbidden  %+v", 403, o.Payload)
}

func (o *CustomerSettingsUpdateForbidden) String() string {
	return fmt.Sprintf("[PATCH /installation-tokens/entities/customer-settings/v1][%d] customerSettingsUpdateForbidden  %+v", 403, o.Payload)
}

func (o *CustomerSettingsUpdateForbidden) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CustomerSettingsUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCustomerSettingsUpdateNotFound creates a CustomerSettingsUpdateNotFound with default headers values
func NewCustomerSettingsUpdateNotFound() *CustomerSettingsUpdateNotFound {
	return &CustomerSettingsUpdateNotFound{}
}

/*
CustomerSettingsUpdateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CustomerSettingsUpdateNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecQueryResponse
}

// IsSuccess returns true when this customer settings update not found response has a 2xx status code
func (o *CustomerSettingsUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this customer settings update not found response has a 3xx status code
func (o *CustomerSettingsUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer settings update not found response has a 4xx status code
func (o *CustomerSettingsUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this customer settings update not found response has a 5xx status code
func (o *CustomerSettingsUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this customer settings update not found response a status code equal to that given
func (o *CustomerSettingsUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the customer settings update not found response
func (o *CustomerSettingsUpdateNotFound) Code() int {
	return 404
}

func (o *CustomerSettingsUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /installation-tokens/entities/customer-settings/v1][%d] customerSettingsUpdateNotFound  %+v", 404, o.Payload)
}

func (o *CustomerSettingsUpdateNotFound) String() string {
	return fmt.Sprintf("[PATCH /installation-tokens/entities/customer-settings/v1][%d] customerSettingsUpdateNotFound  %+v", 404, o.Payload)
}

func (o *CustomerSettingsUpdateNotFound) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *CustomerSettingsUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerSettingsUpdateTooManyRequests creates a CustomerSettingsUpdateTooManyRequests with default headers values
func NewCustomerSettingsUpdateTooManyRequests() *CustomerSettingsUpdateTooManyRequests {
	return &CustomerSettingsUpdateTooManyRequests{}
}

/*
CustomerSettingsUpdateTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CustomerSettingsUpdateTooManyRequests struct {

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

// IsSuccess returns true when this customer settings update too many requests response has a 2xx status code
func (o *CustomerSettingsUpdateTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this customer settings update too many requests response has a 3xx status code
func (o *CustomerSettingsUpdateTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer settings update too many requests response has a 4xx status code
func (o *CustomerSettingsUpdateTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this customer settings update too many requests response has a 5xx status code
func (o *CustomerSettingsUpdateTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this customer settings update too many requests response a status code equal to that given
func (o *CustomerSettingsUpdateTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the customer settings update too many requests response
func (o *CustomerSettingsUpdateTooManyRequests) Code() int {
	return 429
}

func (o *CustomerSettingsUpdateTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /installation-tokens/entities/customer-settings/v1][%d] customerSettingsUpdateTooManyRequests  %+v", 429, o.Payload)
}

func (o *CustomerSettingsUpdateTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /installation-tokens/entities/customer-settings/v1][%d] customerSettingsUpdateTooManyRequests  %+v", 429, o.Payload)
}

func (o *CustomerSettingsUpdateTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CustomerSettingsUpdateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCustomerSettingsUpdateInternalServerError creates a CustomerSettingsUpdateInternalServerError with default headers values
func NewCustomerSettingsUpdateInternalServerError() *CustomerSettingsUpdateInternalServerError {
	return &CustomerSettingsUpdateInternalServerError{}
}

/*
CustomerSettingsUpdateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CustomerSettingsUpdateInternalServerError struct {

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

// IsSuccess returns true when this customer settings update internal server error response has a 2xx status code
func (o *CustomerSettingsUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this customer settings update internal server error response has a 3xx status code
func (o *CustomerSettingsUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer settings update internal server error response has a 4xx status code
func (o *CustomerSettingsUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer settings update internal server error response has a 5xx status code
func (o *CustomerSettingsUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this customer settings update internal server error response a status code equal to that given
func (o *CustomerSettingsUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the customer settings update internal server error response
func (o *CustomerSettingsUpdateInternalServerError) Code() int {
	return 500
}

func (o *CustomerSettingsUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /installation-tokens/entities/customer-settings/v1][%d] customerSettingsUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *CustomerSettingsUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /installation-tokens/entities/customer-settings/v1][%d] customerSettingsUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *CustomerSettingsUpdateInternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CustomerSettingsUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
