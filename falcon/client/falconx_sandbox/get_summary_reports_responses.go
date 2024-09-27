// Code generated by go-swagger; DO NOT EDIT.

package falconx_sandbox

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

// GetSummaryReportsReader is a Reader for the GetSummaryReports structure.
type GetSummaryReportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSummaryReportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSummaryReportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSummaryReportsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSummaryReportsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSummaryReportsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSummaryReportsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /falconx/entities/report-summaries/v1] GetSummaryReports", response, response.Code())
	}
}

// NewGetSummaryReportsOK creates a GetSummaryReportsOK with default headers values
func NewGetSummaryReportsOK() *GetSummaryReportsOK {
	return &GetSummaryReportsOK{}
}

/*
GetSummaryReportsOK describes a response with status code 200, with default header values.

OK
*/
type GetSummaryReportsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FalconxSummaryReportV1Response
}

// IsSuccess returns true when this get summary reports o k response has a 2xx status code
func (o *GetSummaryReportsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get summary reports o k response has a 3xx status code
func (o *GetSummaryReportsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get summary reports o k response has a 4xx status code
func (o *GetSummaryReportsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get summary reports o k response has a 5xx status code
func (o *GetSummaryReportsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get summary reports o k response a status code equal to that given
func (o *GetSummaryReportsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get summary reports o k response
func (o *GetSummaryReportsOK) Code() int {
	return 200
}

func (o *GetSummaryReportsOK) Error() string {
	return fmt.Sprintf("[GET /falconx/entities/report-summaries/v1][%d] getSummaryReportsOK  %+v", 200, o.Payload)
}

func (o *GetSummaryReportsOK) String() string {
	return fmt.Sprintf("[GET /falconx/entities/report-summaries/v1][%d] getSummaryReportsOK  %+v", 200, o.Payload)
}

func (o *GetSummaryReportsOK) GetPayload() *models.FalconxSummaryReportV1Response {
	return o.Payload
}

func (o *GetSummaryReportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FalconxSummaryReportV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSummaryReportsBadRequest creates a GetSummaryReportsBadRequest with default headers values
func NewGetSummaryReportsBadRequest() *GetSummaryReportsBadRequest {
	return &GetSummaryReportsBadRequest{}
}

/*
GetSummaryReportsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetSummaryReportsBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FalconxSummaryReportV1Response
}

// IsSuccess returns true when this get summary reports bad request response has a 2xx status code
func (o *GetSummaryReportsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get summary reports bad request response has a 3xx status code
func (o *GetSummaryReportsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get summary reports bad request response has a 4xx status code
func (o *GetSummaryReportsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get summary reports bad request response has a 5xx status code
func (o *GetSummaryReportsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get summary reports bad request response a status code equal to that given
func (o *GetSummaryReportsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get summary reports bad request response
func (o *GetSummaryReportsBadRequest) Code() int {
	return 400
}

func (o *GetSummaryReportsBadRequest) Error() string {
	return fmt.Sprintf("[GET /falconx/entities/report-summaries/v1][%d] getSummaryReportsBadRequest  %+v", 400, o.Payload)
}

func (o *GetSummaryReportsBadRequest) String() string {
	return fmt.Sprintf("[GET /falconx/entities/report-summaries/v1][%d] getSummaryReportsBadRequest  %+v", 400, o.Payload)
}

func (o *GetSummaryReportsBadRequest) GetPayload() *models.FalconxSummaryReportV1Response {
	return o.Payload
}

func (o *GetSummaryReportsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FalconxSummaryReportV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSummaryReportsForbidden creates a GetSummaryReportsForbidden with default headers values
func NewGetSummaryReportsForbidden() *GetSummaryReportsForbidden {
	return &GetSummaryReportsForbidden{}
}

/*
GetSummaryReportsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSummaryReportsForbidden struct {

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

// IsSuccess returns true when this get summary reports forbidden response has a 2xx status code
func (o *GetSummaryReportsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get summary reports forbidden response has a 3xx status code
func (o *GetSummaryReportsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get summary reports forbidden response has a 4xx status code
func (o *GetSummaryReportsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get summary reports forbidden response has a 5xx status code
func (o *GetSummaryReportsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get summary reports forbidden response a status code equal to that given
func (o *GetSummaryReportsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get summary reports forbidden response
func (o *GetSummaryReportsForbidden) Code() int {
	return 403
}

func (o *GetSummaryReportsForbidden) Error() string {
	return fmt.Sprintf("[GET /falconx/entities/report-summaries/v1][%d] getSummaryReportsForbidden  %+v", 403, o.Payload)
}

func (o *GetSummaryReportsForbidden) String() string {
	return fmt.Sprintf("[GET /falconx/entities/report-summaries/v1][%d] getSummaryReportsForbidden  %+v", 403, o.Payload)
}

func (o *GetSummaryReportsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetSummaryReportsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSummaryReportsTooManyRequests creates a GetSummaryReportsTooManyRequests with default headers values
func NewGetSummaryReportsTooManyRequests() *GetSummaryReportsTooManyRequests {
	return &GetSummaryReportsTooManyRequests{}
}

/*
GetSummaryReportsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetSummaryReportsTooManyRequests struct {

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

// IsSuccess returns true when this get summary reports too many requests response has a 2xx status code
func (o *GetSummaryReportsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get summary reports too many requests response has a 3xx status code
func (o *GetSummaryReportsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get summary reports too many requests response has a 4xx status code
func (o *GetSummaryReportsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get summary reports too many requests response has a 5xx status code
func (o *GetSummaryReportsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get summary reports too many requests response a status code equal to that given
func (o *GetSummaryReportsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get summary reports too many requests response
func (o *GetSummaryReportsTooManyRequests) Code() int {
	return 429
}

func (o *GetSummaryReportsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /falconx/entities/report-summaries/v1][%d] getSummaryReportsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSummaryReportsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /falconx/entities/report-summaries/v1][%d] getSummaryReportsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSummaryReportsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetSummaryReportsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSummaryReportsInternalServerError creates a GetSummaryReportsInternalServerError with default headers values
func NewGetSummaryReportsInternalServerError() *GetSummaryReportsInternalServerError {
	return &GetSummaryReportsInternalServerError{}
}

/*
GetSummaryReportsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetSummaryReportsInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FalconxSummaryReportV1Response
}

// IsSuccess returns true when this get summary reports internal server error response has a 2xx status code
func (o *GetSummaryReportsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get summary reports internal server error response has a 3xx status code
func (o *GetSummaryReportsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get summary reports internal server error response has a 4xx status code
func (o *GetSummaryReportsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get summary reports internal server error response has a 5xx status code
func (o *GetSummaryReportsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get summary reports internal server error response a status code equal to that given
func (o *GetSummaryReportsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get summary reports internal server error response
func (o *GetSummaryReportsInternalServerError) Code() int {
	return 500
}

func (o *GetSummaryReportsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /falconx/entities/report-summaries/v1][%d] getSummaryReportsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSummaryReportsInternalServerError) String() string {
	return fmt.Sprintf("[GET /falconx/entities/report-summaries/v1][%d] getSummaryReportsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSummaryReportsInternalServerError) GetPayload() *models.FalconxSummaryReportV1Response {
	return o.Payload
}

func (o *GetSummaryReportsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FalconxSummaryReportV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
