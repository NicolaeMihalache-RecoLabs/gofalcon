// Code generated by go-swagger; DO NOT EDIT.

package user_management

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

// RetrieveUserUUIDsByCIDReader is a Reader for the RetrieveUserUUIDsByCID structure.
type RetrieveUserUUIDsByCIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveUserUUIDsByCIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrieveUserUUIDsByCIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRetrieveUserUUIDsByCIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRetrieveUserUUIDsByCIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRetrieveUserUUIDsByCIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /users/queries/user-uuids-by-cid/v1] RetrieveUserUUIDsByCID", response, response.Code())
	}
}

// NewRetrieveUserUUIDsByCIDOK creates a RetrieveUserUUIDsByCIDOK with default headers values
func NewRetrieveUserUUIDsByCIDOK() *RetrieveUserUUIDsByCIDOK {
	return &RetrieveUserUUIDsByCIDOK{}
}

/*
RetrieveUserUUIDsByCIDOK describes a response with status code 200, with default header values.

OK
*/
type RetrieveUserUUIDsByCIDOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this retrieve user u Ui ds by c Id o k response has a 2xx status code
func (o *RetrieveUserUUIDsByCIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this retrieve user u Ui ds by c Id o k response has a 3xx status code
func (o *RetrieveUserUUIDsByCIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retrieve user u Ui ds by c Id o k response has a 4xx status code
func (o *RetrieveUserUUIDsByCIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this retrieve user u Ui ds by c Id o k response has a 5xx status code
func (o *RetrieveUserUUIDsByCIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this retrieve user u Ui ds by c Id o k response a status code equal to that given
func (o *RetrieveUserUUIDsByCIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the retrieve user u Ui ds by c Id o k response
func (o *RetrieveUserUUIDsByCIDOK) Code() int {
	return 200
}

func (o *RetrieveUserUUIDsByCIDOK) Error() string {
	return fmt.Sprintf("[GET /users/queries/user-uuids-by-cid/v1][%d] retrieveUserUUiDsByCIdOK  %+v", 200, o.Payload)
}

func (o *RetrieveUserUUIDsByCIDOK) String() string {
	return fmt.Sprintf("[GET /users/queries/user-uuids-by-cid/v1][%d] retrieveUserUUiDsByCIdOK  %+v", 200, o.Payload)
}

func (o *RetrieveUserUUIDsByCIDOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *RetrieveUserUUIDsByCIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveUserUUIDsByCIDBadRequest creates a RetrieveUserUUIDsByCIDBadRequest with default headers values
func NewRetrieveUserUUIDsByCIDBadRequest() *RetrieveUserUUIDsByCIDBadRequest {
	return &RetrieveUserUUIDsByCIDBadRequest{}
}

/*
RetrieveUserUUIDsByCIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RetrieveUserUUIDsByCIDBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this retrieve user u Ui ds by c Id bad request response has a 2xx status code
func (o *RetrieveUserUUIDsByCIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retrieve user u Ui ds by c Id bad request response has a 3xx status code
func (o *RetrieveUserUUIDsByCIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retrieve user u Ui ds by c Id bad request response has a 4xx status code
func (o *RetrieveUserUUIDsByCIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this retrieve user u Ui ds by c Id bad request response has a 5xx status code
func (o *RetrieveUserUUIDsByCIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this retrieve user u Ui ds by c Id bad request response a status code equal to that given
func (o *RetrieveUserUUIDsByCIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the retrieve user u Ui ds by c Id bad request response
func (o *RetrieveUserUUIDsByCIDBadRequest) Code() int {
	return 400
}

func (o *RetrieveUserUUIDsByCIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/queries/user-uuids-by-cid/v1][%d] retrieveUserUUiDsByCIdBadRequest  %+v", 400, o.Payload)
}

func (o *RetrieveUserUUIDsByCIDBadRequest) String() string {
	return fmt.Sprintf("[GET /users/queries/user-uuids-by-cid/v1][%d] retrieveUserUUiDsByCIdBadRequest  %+v", 400, o.Payload)
}

func (o *RetrieveUserUUIDsByCIDBadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *RetrieveUserUUIDsByCIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveUserUUIDsByCIDForbidden creates a RetrieveUserUUIDsByCIDForbidden with default headers values
func NewRetrieveUserUUIDsByCIDForbidden() *RetrieveUserUUIDsByCIDForbidden {
	return &RetrieveUserUUIDsByCIDForbidden{}
}

/*
RetrieveUserUUIDsByCIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RetrieveUserUUIDsByCIDForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this retrieve user u Ui ds by c Id forbidden response has a 2xx status code
func (o *RetrieveUserUUIDsByCIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retrieve user u Ui ds by c Id forbidden response has a 3xx status code
func (o *RetrieveUserUUIDsByCIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retrieve user u Ui ds by c Id forbidden response has a 4xx status code
func (o *RetrieveUserUUIDsByCIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this retrieve user u Ui ds by c Id forbidden response has a 5xx status code
func (o *RetrieveUserUUIDsByCIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this retrieve user u Ui ds by c Id forbidden response a status code equal to that given
func (o *RetrieveUserUUIDsByCIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the retrieve user u Ui ds by c Id forbidden response
func (o *RetrieveUserUUIDsByCIDForbidden) Code() int {
	return 403
}

func (o *RetrieveUserUUIDsByCIDForbidden) Error() string {
	return fmt.Sprintf("[GET /users/queries/user-uuids-by-cid/v1][%d] retrieveUserUUiDsByCIdForbidden  %+v", 403, o.Payload)
}

func (o *RetrieveUserUUIDsByCIDForbidden) String() string {
	return fmt.Sprintf("[GET /users/queries/user-uuids-by-cid/v1][%d] retrieveUserUUiDsByCIdForbidden  %+v", 403, o.Payload)
}

func (o *RetrieveUserUUIDsByCIDForbidden) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *RetrieveUserUUIDsByCIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveUserUUIDsByCIDTooManyRequests creates a RetrieveUserUUIDsByCIDTooManyRequests with default headers values
func NewRetrieveUserUUIDsByCIDTooManyRequests() *RetrieveUserUUIDsByCIDTooManyRequests {
	return &RetrieveUserUUIDsByCIDTooManyRequests{}
}

/*
RetrieveUserUUIDsByCIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RetrieveUserUUIDsByCIDTooManyRequests struct {

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

// IsSuccess returns true when this retrieve user u Ui ds by c Id too many requests response has a 2xx status code
func (o *RetrieveUserUUIDsByCIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retrieve user u Ui ds by c Id too many requests response has a 3xx status code
func (o *RetrieveUserUUIDsByCIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retrieve user u Ui ds by c Id too many requests response has a 4xx status code
func (o *RetrieveUserUUIDsByCIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this retrieve user u Ui ds by c Id too many requests response has a 5xx status code
func (o *RetrieveUserUUIDsByCIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this retrieve user u Ui ds by c Id too many requests response a status code equal to that given
func (o *RetrieveUserUUIDsByCIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the retrieve user u Ui ds by c Id too many requests response
func (o *RetrieveUserUUIDsByCIDTooManyRequests) Code() int {
	return 429
}

func (o *RetrieveUserUUIDsByCIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /users/queries/user-uuids-by-cid/v1][%d] retrieveUserUUiDsByCIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *RetrieveUserUUIDsByCIDTooManyRequests) String() string {
	return fmt.Sprintf("[GET /users/queries/user-uuids-by-cid/v1][%d] retrieveUserUUiDsByCIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *RetrieveUserUUIDsByCIDTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RetrieveUserUUIDsByCIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
