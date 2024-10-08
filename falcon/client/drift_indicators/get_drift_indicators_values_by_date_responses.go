// Code generated by go-swagger; DO NOT EDIT.

package drift_indicators

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

// GetDriftIndicatorsValuesByDateReader is a Reader for the GetDriftIndicatorsValuesByDate structure.
type GetDriftIndicatorsValuesByDateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDriftIndicatorsValuesByDateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDriftIndicatorsValuesByDateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetDriftIndicatorsValuesByDateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetDriftIndicatorsValuesByDateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDriftIndicatorsValuesByDateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/aggregates/drift-indicators/count-by-date/v1] GetDriftIndicatorsValuesByDate", response, response.Code())
	}
}

// NewGetDriftIndicatorsValuesByDateOK creates a GetDriftIndicatorsValuesByDateOK with default headers values
func NewGetDriftIndicatorsValuesByDateOK() *GetDriftIndicatorsValuesByDateOK {
	return &GetDriftIndicatorsValuesByDateOK{}
}

/*
GetDriftIndicatorsValuesByDateOK describes a response with status code 200, with default header values.

OK
*/
type GetDriftIndicatorsValuesByDateOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DriftindicatorsDriftIndicatorsFieldValue
}

// IsSuccess returns true when this get drift indicators values by date o k response has a 2xx status code
func (o *GetDriftIndicatorsValuesByDateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get drift indicators values by date o k response has a 3xx status code
func (o *GetDriftIndicatorsValuesByDateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get drift indicators values by date o k response has a 4xx status code
func (o *GetDriftIndicatorsValuesByDateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get drift indicators values by date o k response has a 5xx status code
func (o *GetDriftIndicatorsValuesByDateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get drift indicators values by date o k response a status code equal to that given
func (o *GetDriftIndicatorsValuesByDateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get drift indicators values by date o k response
func (o *GetDriftIndicatorsValuesByDateOK) Code() int {
	return 200
}

func (o *GetDriftIndicatorsValuesByDateOK) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/drift-indicators/count-by-date/v1][%d] getDriftIndicatorsValuesByDateOK  %+v", 200, o.Payload)
}

func (o *GetDriftIndicatorsValuesByDateOK) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/drift-indicators/count-by-date/v1][%d] getDriftIndicatorsValuesByDateOK  %+v", 200, o.Payload)
}

func (o *GetDriftIndicatorsValuesByDateOK) GetPayload() *models.DriftindicatorsDriftIndicatorsFieldValue {
	return o.Payload
}

func (o *GetDriftIndicatorsValuesByDateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DriftindicatorsDriftIndicatorsFieldValue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDriftIndicatorsValuesByDateForbidden creates a GetDriftIndicatorsValuesByDateForbidden with default headers values
func NewGetDriftIndicatorsValuesByDateForbidden() *GetDriftIndicatorsValuesByDateForbidden {
	return &GetDriftIndicatorsValuesByDateForbidden{}
}

/*
GetDriftIndicatorsValuesByDateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDriftIndicatorsValuesByDateForbidden struct {

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

// IsSuccess returns true when this get drift indicators values by date forbidden response has a 2xx status code
func (o *GetDriftIndicatorsValuesByDateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get drift indicators values by date forbidden response has a 3xx status code
func (o *GetDriftIndicatorsValuesByDateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get drift indicators values by date forbidden response has a 4xx status code
func (o *GetDriftIndicatorsValuesByDateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get drift indicators values by date forbidden response has a 5xx status code
func (o *GetDriftIndicatorsValuesByDateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get drift indicators values by date forbidden response a status code equal to that given
func (o *GetDriftIndicatorsValuesByDateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get drift indicators values by date forbidden response
func (o *GetDriftIndicatorsValuesByDateForbidden) Code() int {
	return 403
}

func (o *GetDriftIndicatorsValuesByDateForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/drift-indicators/count-by-date/v1][%d] getDriftIndicatorsValuesByDateForbidden  %+v", 403, o.Payload)
}

func (o *GetDriftIndicatorsValuesByDateForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/drift-indicators/count-by-date/v1][%d] getDriftIndicatorsValuesByDateForbidden  %+v", 403, o.Payload)
}

func (o *GetDriftIndicatorsValuesByDateForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetDriftIndicatorsValuesByDateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDriftIndicatorsValuesByDateTooManyRequests creates a GetDriftIndicatorsValuesByDateTooManyRequests with default headers values
func NewGetDriftIndicatorsValuesByDateTooManyRequests() *GetDriftIndicatorsValuesByDateTooManyRequests {
	return &GetDriftIndicatorsValuesByDateTooManyRequests{}
}

/*
GetDriftIndicatorsValuesByDateTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetDriftIndicatorsValuesByDateTooManyRequests struct {

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

// IsSuccess returns true when this get drift indicators values by date too many requests response has a 2xx status code
func (o *GetDriftIndicatorsValuesByDateTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get drift indicators values by date too many requests response has a 3xx status code
func (o *GetDriftIndicatorsValuesByDateTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get drift indicators values by date too many requests response has a 4xx status code
func (o *GetDriftIndicatorsValuesByDateTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get drift indicators values by date too many requests response has a 5xx status code
func (o *GetDriftIndicatorsValuesByDateTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get drift indicators values by date too many requests response a status code equal to that given
func (o *GetDriftIndicatorsValuesByDateTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get drift indicators values by date too many requests response
func (o *GetDriftIndicatorsValuesByDateTooManyRequests) Code() int {
	return 429
}

func (o *GetDriftIndicatorsValuesByDateTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/drift-indicators/count-by-date/v1][%d] getDriftIndicatorsValuesByDateTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDriftIndicatorsValuesByDateTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/drift-indicators/count-by-date/v1][%d] getDriftIndicatorsValuesByDateTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDriftIndicatorsValuesByDateTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetDriftIndicatorsValuesByDateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDriftIndicatorsValuesByDateInternalServerError creates a GetDriftIndicatorsValuesByDateInternalServerError with default headers values
func NewGetDriftIndicatorsValuesByDateInternalServerError() *GetDriftIndicatorsValuesByDateInternalServerError {
	return &GetDriftIndicatorsValuesByDateInternalServerError{}
}

/*
GetDriftIndicatorsValuesByDateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDriftIndicatorsValuesByDateInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.CoreEntitiesResponse
}

// IsSuccess returns true when this get drift indicators values by date internal server error response has a 2xx status code
func (o *GetDriftIndicatorsValuesByDateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get drift indicators values by date internal server error response has a 3xx status code
func (o *GetDriftIndicatorsValuesByDateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get drift indicators values by date internal server error response has a 4xx status code
func (o *GetDriftIndicatorsValuesByDateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get drift indicators values by date internal server error response has a 5xx status code
func (o *GetDriftIndicatorsValuesByDateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get drift indicators values by date internal server error response a status code equal to that given
func (o *GetDriftIndicatorsValuesByDateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get drift indicators values by date internal server error response
func (o *GetDriftIndicatorsValuesByDateInternalServerError) Code() int {
	return 500
}

func (o *GetDriftIndicatorsValuesByDateInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/drift-indicators/count-by-date/v1][%d] getDriftIndicatorsValuesByDateInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDriftIndicatorsValuesByDateInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/drift-indicators/count-by-date/v1][%d] getDriftIndicatorsValuesByDateInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDriftIndicatorsValuesByDateInternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *GetDriftIndicatorsValuesByDateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.CoreEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
