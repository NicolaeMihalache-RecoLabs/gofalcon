// Code generated by go-swagger; DO NOT EDIT.

package filevantage

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

// UpdateRuleGroupPrecedenceReader is a Reader for the UpdateRuleGroupPrecedence structure.
type UpdateRuleGroupPrecedenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRuleGroupPrecedenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRuleGroupPrecedenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRuleGroupPrecedenceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRuleGroupPrecedenceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateRuleGroupPrecedenceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateRuleGroupPrecedenceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateRuleGroupPrecedenceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /filevantage/entities/rule-groups-rule-precedence/v1] updateRuleGroupPrecedence", response, response.Code())
	}
}

// NewUpdateRuleGroupPrecedenceOK creates a UpdateRuleGroupPrecedenceOK with default headers values
func NewUpdateRuleGroupPrecedenceOK() *UpdateRuleGroupPrecedenceOK {
	return &UpdateRuleGroupPrecedenceOK{}
}

/*
UpdateRuleGroupPrecedenceOK describes a response with status code 200, with default header values.

Rule precedence order has been set.
*/
type UpdateRuleGroupPrecedenceOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RulegroupsResponse
}

// IsSuccess returns true when this update rule group precedence o k response has a 2xx status code
func (o *UpdateRuleGroupPrecedenceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update rule group precedence o k response has a 3xx status code
func (o *UpdateRuleGroupPrecedenceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rule group precedence o k response has a 4xx status code
func (o *UpdateRuleGroupPrecedenceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update rule group precedence o k response has a 5xx status code
func (o *UpdateRuleGroupPrecedenceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update rule group precedence o k response a status code equal to that given
func (o *UpdateRuleGroupPrecedenceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update rule group precedence o k response
func (o *UpdateRuleGroupPrecedenceOK) Code() int {
	return 200
}

func (o *UpdateRuleGroupPrecedenceOK) Error() string {
	return fmt.Sprintf("[PATCH /filevantage/entities/rule-groups-rule-precedence/v1][%d] updateRuleGroupPrecedenceOK  %+v", 200, o.Payload)
}

func (o *UpdateRuleGroupPrecedenceOK) String() string {
	return fmt.Sprintf("[PATCH /filevantage/entities/rule-groups-rule-precedence/v1][%d] updateRuleGroupPrecedenceOK  %+v", 200, o.Payload)
}

func (o *UpdateRuleGroupPrecedenceOK) GetPayload() *models.RulegroupsResponse {
	return o.Payload
}

func (o *UpdateRuleGroupPrecedenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RulegroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRuleGroupPrecedenceBadRequest creates a UpdateRuleGroupPrecedenceBadRequest with default headers values
func NewUpdateRuleGroupPrecedenceBadRequest() *UpdateRuleGroupPrecedenceBadRequest {
	return &UpdateRuleGroupPrecedenceBadRequest{}
}

/*
UpdateRuleGroupPrecedenceBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateRuleGroupPrecedenceBadRequest struct {

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

// IsSuccess returns true when this update rule group precedence bad request response has a 2xx status code
func (o *UpdateRuleGroupPrecedenceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update rule group precedence bad request response has a 3xx status code
func (o *UpdateRuleGroupPrecedenceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rule group precedence bad request response has a 4xx status code
func (o *UpdateRuleGroupPrecedenceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update rule group precedence bad request response has a 5xx status code
func (o *UpdateRuleGroupPrecedenceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update rule group precedence bad request response a status code equal to that given
func (o *UpdateRuleGroupPrecedenceBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update rule group precedence bad request response
func (o *UpdateRuleGroupPrecedenceBadRequest) Code() int {
	return 400
}

func (o *UpdateRuleGroupPrecedenceBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /filevantage/entities/rule-groups-rule-precedence/v1][%d] updateRuleGroupPrecedenceBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateRuleGroupPrecedenceBadRequest) String() string {
	return fmt.Sprintf("[PATCH /filevantage/entities/rule-groups-rule-precedence/v1][%d] updateRuleGroupPrecedenceBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateRuleGroupPrecedenceBadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *UpdateRuleGroupPrecedenceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRuleGroupPrecedenceForbidden creates a UpdateRuleGroupPrecedenceForbidden with default headers values
func NewUpdateRuleGroupPrecedenceForbidden() *UpdateRuleGroupPrecedenceForbidden {
	return &UpdateRuleGroupPrecedenceForbidden{}
}

/*
UpdateRuleGroupPrecedenceForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateRuleGroupPrecedenceForbidden struct {

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

// IsSuccess returns true when this update rule group precedence forbidden response has a 2xx status code
func (o *UpdateRuleGroupPrecedenceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update rule group precedence forbidden response has a 3xx status code
func (o *UpdateRuleGroupPrecedenceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rule group precedence forbidden response has a 4xx status code
func (o *UpdateRuleGroupPrecedenceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update rule group precedence forbidden response has a 5xx status code
func (o *UpdateRuleGroupPrecedenceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update rule group precedence forbidden response a status code equal to that given
func (o *UpdateRuleGroupPrecedenceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update rule group precedence forbidden response
func (o *UpdateRuleGroupPrecedenceForbidden) Code() int {
	return 403
}

func (o *UpdateRuleGroupPrecedenceForbidden) Error() string {
	return fmt.Sprintf("[PATCH /filevantage/entities/rule-groups-rule-precedence/v1][%d] updateRuleGroupPrecedenceForbidden  %+v", 403, o.Payload)
}

func (o *UpdateRuleGroupPrecedenceForbidden) String() string {
	return fmt.Sprintf("[PATCH /filevantage/entities/rule-groups-rule-precedence/v1][%d] updateRuleGroupPrecedenceForbidden  %+v", 403, o.Payload)
}

func (o *UpdateRuleGroupPrecedenceForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateRuleGroupPrecedenceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRuleGroupPrecedenceConflict creates a UpdateRuleGroupPrecedenceConflict with default headers values
func NewUpdateRuleGroupPrecedenceConflict() *UpdateRuleGroupPrecedenceConflict {
	return &UpdateRuleGroupPrecedenceConflict{}
}

/*
UpdateRuleGroupPrecedenceConflict describes a response with status code 409, with default header values.

Conflict
*/
type UpdateRuleGroupPrecedenceConflict struct {

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

// IsSuccess returns true when this update rule group precedence conflict response has a 2xx status code
func (o *UpdateRuleGroupPrecedenceConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update rule group precedence conflict response has a 3xx status code
func (o *UpdateRuleGroupPrecedenceConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rule group precedence conflict response has a 4xx status code
func (o *UpdateRuleGroupPrecedenceConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update rule group precedence conflict response has a 5xx status code
func (o *UpdateRuleGroupPrecedenceConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update rule group precedence conflict response a status code equal to that given
func (o *UpdateRuleGroupPrecedenceConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update rule group precedence conflict response
func (o *UpdateRuleGroupPrecedenceConflict) Code() int {
	return 409
}

func (o *UpdateRuleGroupPrecedenceConflict) Error() string {
	return fmt.Sprintf("[PATCH /filevantage/entities/rule-groups-rule-precedence/v1][%d] updateRuleGroupPrecedenceConflict  %+v", 409, o.Payload)
}

func (o *UpdateRuleGroupPrecedenceConflict) String() string {
	return fmt.Sprintf("[PATCH /filevantage/entities/rule-groups-rule-precedence/v1][%d] updateRuleGroupPrecedenceConflict  %+v", 409, o.Payload)
}

func (o *UpdateRuleGroupPrecedenceConflict) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *UpdateRuleGroupPrecedenceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRuleGroupPrecedenceTooManyRequests creates a UpdateRuleGroupPrecedenceTooManyRequests with default headers values
func NewUpdateRuleGroupPrecedenceTooManyRequests() *UpdateRuleGroupPrecedenceTooManyRequests {
	return &UpdateRuleGroupPrecedenceTooManyRequests{}
}

/*
UpdateRuleGroupPrecedenceTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateRuleGroupPrecedenceTooManyRequests struct {

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

// IsSuccess returns true when this update rule group precedence too many requests response has a 2xx status code
func (o *UpdateRuleGroupPrecedenceTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update rule group precedence too many requests response has a 3xx status code
func (o *UpdateRuleGroupPrecedenceTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rule group precedence too many requests response has a 4xx status code
func (o *UpdateRuleGroupPrecedenceTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update rule group precedence too many requests response has a 5xx status code
func (o *UpdateRuleGroupPrecedenceTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update rule group precedence too many requests response a status code equal to that given
func (o *UpdateRuleGroupPrecedenceTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the update rule group precedence too many requests response
func (o *UpdateRuleGroupPrecedenceTooManyRequests) Code() int {
	return 429
}

func (o *UpdateRuleGroupPrecedenceTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /filevantage/entities/rule-groups-rule-precedence/v1][%d] updateRuleGroupPrecedenceTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateRuleGroupPrecedenceTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /filevantage/entities/rule-groups-rule-precedence/v1][%d] updateRuleGroupPrecedenceTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateRuleGroupPrecedenceTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateRuleGroupPrecedenceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRuleGroupPrecedenceInternalServerError creates a UpdateRuleGroupPrecedenceInternalServerError with default headers values
func NewUpdateRuleGroupPrecedenceInternalServerError() *UpdateRuleGroupPrecedenceInternalServerError {
	return &UpdateRuleGroupPrecedenceInternalServerError{}
}

/*
UpdateRuleGroupPrecedenceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateRuleGroupPrecedenceInternalServerError struct {

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

// IsSuccess returns true when this update rule group precedence internal server error response has a 2xx status code
func (o *UpdateRuleGroupPrecedenceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update rule group precedence internal server error response has a 3xx status code
func (o *UpdateRuleGroupPrecedenceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rule group precedence internal server error response has a 4xx status code
func (o *UpdateRuleGroupPrecedenceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update rule group precedence internal server error response has a 5xx status code
func (o *UpdateRuleGroupPrecedenceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update rule group precedence internal server error response a status code equal to that given
func (o *UpdateRuleGroupPrecedenceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update rule group precedence internal server error response
func (o *UpdateRuleGroupPrecedenceInternalServerError) Code() int {
	return 500
}

func (o *UpdateRuleGroupPrecedenceInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /filevantage/entities/rule-groups-rule-precedence/v1][%d] updateRuleGroupPrecedenceInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateRuleGroupPrecedenceInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /filevantage/entities/rule-groups-rule-precedence/v1][%d] updateRuleGroupPrecedenceInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateRuleGroupPrecedenceInternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *UpdateRuleGroupPrecedenceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
