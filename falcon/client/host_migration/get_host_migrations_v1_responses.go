// Code generated by go-swagger; DO NOT EDIT.

package host_migration

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

// GetHostMigrationsV1Reader is a Reader for the GetHostMigrationsV1 structure.
type GetHostMigrationsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHostMigrationsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHostMigrationsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetHostMigrationsV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetHostMigrationsV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetHostMigrationsV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetHostMigrationsV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetHostMigrationsV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /host-migration/entities/host-migrations/GET/v1] GetHostMigrationsV1", response, response.Code())
	}
}

// NewGetHostMigrationsV1OK creates a GetHostMigrationsV1OK with default headers values
func NewGetHostMigrationsV1OK() *GetHostMigrationsV1OK {
	return &GetHostMigrationsV1OK{}
}

/*
GetHostMigrationsV1OK describes a response with status code 200, with default header values.

OK
*/
type GetHostMigrationsV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIGetHostMigrationResponseV1
}

// IsSuccess returns true when this get host migrations v1 o k response has a 2xx status code
func (o *GetHostMigrationsV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get host migrations v1 o k response has a 3xx status code
func (o *GetHostMigrationsV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get host migrations v1 o k response has a 4xx status code
func (o *GetHostMigrationsV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get host migrations v1 o k response has a 5xx status code
func (o *GetHostMigrationsV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get host migrations v1 o k response a status code equal to that given
func (o *GetHostMigrationsV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get host migrations v1 o k response
func (o *GetHostMigrationsV1OK) Code() int {
	return 200
}

func (o *GetHostMigrationsV1OK) Error() string {
	return fmt.Sprintf("[POST /host-migration/entities/host-migrations/GET/v1][%d] getHostMigrationsV1OK  %+v", 200, o.Payload)
}

func (o *GetHostMigrationsV1OK) String() string {
	return fmt.Sprintf("[POST /host-migration/entities/host-migrations/GET/v1][%d] getHostMigrationsV1OK  %+v", 200, o.Payload)
}

func (o *GetHostMigrationsV1OK) GetPayload() *models.APIGetHostMigrationResponseV1 {
	return o.Payload
}

func (o *GetHostMigrationsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIGetHostMigrationResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostMigrationsV1BadRequest creates a GetHostMigrationsV1BadRequest with default headers values
func NewGetHostMigrationsV1BadRequest() *GetHostMigrationsV1BadRequest {
	return &GetHostMigrationsV1BadRequest{}
}

/*
GetHostMigrationsV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetHostMigrationsV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIGetHostMigrationResponseV1
}

// IsSuccess returns true when this get host migrations v1 bad request response has a 2xx status code
func (o *GetHostMigrationsV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get host migrations v1 bad request response has a 3xx status code
func (o *GetHostMigrationsV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get host migrations v1 bad request response has a 4xx status code
func (o *GetHostMigrationsV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get host migrations v1 bad request response has a 5xx status code
func (o *GetHostMigrationsV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get host migrations v1 bad request response a status code equal to that given
func (o *GetHostMigrationsV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get host migrations v1 bad request response
func (o *GetHostMigrationsV1BadRequest) Code() int {
	return 400
}

func (o *GetHostMigrationsV1BadRequest) Error() string {
	return fmt.Sprintf("[POST /host-migration/entities/host-migrations/GET/v1][%d] getHostMigrationsV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetHostMigrationsV1BadRequest) String() string {
	return fmt.Sprintf("[POST /host-migration/entities/host-migrations/GET/v1][%d] getHostMigrationsV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetHostMigrationsV1BadRequest) GetPayload() *models.APIGetHostMigrationResponseV1 {
	return o.Payload
}

func (o *GetHostMigrationsV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIGetHostMigrationResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostMigrationsV1Forbidden creates a GetHostMigrationsV1Forbidden with default headers values
func NewGetHostMigrationsV1Forbidden() *GetHostMigrationsV1Forbidden {
	return &GetHostMigrationsV1Forbidden{}
}

/*
GetHostMigrationsV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetHostMigrationsV1Forbidden struct {

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

// IsSuccess returns true when this get host migrations v1 forbidden response has a 2xx status code
func (o *GetHostMigrationsV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get host migrations v1 forbidden response has a 3xx status code
func (o *GetHostMigrationsV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get host migrations v1 forbidden response has a 4xx status code
func (o *GetHostMigrationsV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get host migrations v1 forbidden response has a 5xx status code
func (o *GetHostMigrationsV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get host migrations v1 forbidden response a status code equal to that given
func (o *GetHostMigrationsV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get host migrations v1 forbidden response
func (o *GetHostMigrationsV1Forbidden) Code() int {
	return 403
}

func (o *GetHostMigrationsV1Forbidden) Error() string {
	return fmt.Sprintf("[POST /host-migration/entities/host-migrations/GET/v1][%d] getHostMigrationsV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetHostMigrationsV1Forbidden) String() string {
	return fmt.Sprintf("[POST /host-migration/entities/host-migrations/GET/v1][%d] getHostMigrationsV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetHostMigrationsV1Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetHostMigrationsV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetHostMigrationsV1NotFound creates a GetHostMigrationsV1NotFound with default headers values
func NewGetHostMigrationsV1NotFound() *GetHostMigrationsV1NotFound {
	return &GetHostMigrationsV1NotFound{}
}

/*
GetHostMigrationsV1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetHostMigrationsV1NotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIGetHostMigrationResponseV1
}

// IsSuccess returns true when this get host migrations v1 not found response has a 2xx status code
func (o *GetHostMigrationsV1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get host migrations v1 not found response has a 3xx status code
func (o *GetHostMigrationsV1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get host migrations v1 not found response has a 4xx status code
func (o *GetHostMigrationsV1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get host migrations v1 not found response has a 5xx status code
func (o *GetHostMigrationsV1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get host migrations v1 not found response a status code equal to that given
func (o *GetHostMigrationsV1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get host migrations v1 not found response
func (o *GetHostMigrationsV1NotFound) Code() int {
	return 404
}

func (o *GetHostMigrationsV1NotFound) Error() string {
	return fmt.Sprintf("[POST /host-migration/entities/host-migrations/GET/v1][%d] getHostMigrationsV1NotFound  %+v", 404, o.Payload)
}

func (o *GetHostMigrationsV1NotFound) String() string {
	return fmt.Sprintf("[POST /host-migration/entities/host-migrations/GET/v1][%d] getHostMigrationsV1NotFound  %+v", 404, o.Payload)
}

func (o *GetHostMigrationsV1NotFound) GetPayload() *models.APIGetHostMigrationResponseV1 {
	return o.Payload
}

func (o *GetHostMigrationsV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIGetHostMigrationResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostMigrationsV1TooManyRequests creates a GetHostMigrationsV1TooManyRequests with default headers values
func NewGetHostMigrationsV1TooManyRequests() *GetHostMigrationsV1TooManyRequests {
	return &GetHostMigrationsV1TooManyRequests{}
}

/*
GetHostMigrationsV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetHostMigrationsV1TooManyRequests struct {

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

// IsSuccess returns true when this get host migrations v1 too many requests response has a 2xx status code
func (o *GetHostMigrationsV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get host migrations v1 too many requests response has a 3xx status code
func (o *GetHostMigrationsV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get host migrations v1 too many requests response has a 4xx status code
func (o *GetHostMigrationsV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get host migrations v1 too many requests response has a 5xx status code
func (o *GetHostMigrationsV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get host migrations v1 too many requests response a status code equal to that given
func (o *GetHostMigrationsV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get host migrations v1 too many requests response
func (o *GetHostMigrationsV1TooManyRequests) Code() int {
	return 429
}

func (o *GetHostMigrationsV1TooManyRequests) Error() string {
	return fmt.Sprintf("[POST /host-migration/entities/host-migrations/GET/v1][%d] getHostMigrationsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetHostMigrationsV1TooManyRequests) String() string {
	return fmt.Sprintf("[POST /host-migration/entities/host-migrations/GET/v1][%d] getHostMigrationsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetHostMigrationsV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetHostMigrationsV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetHostMigrationsV1InternalServerError creates a GetHostMigrationsV1InternalServerError with default headers values
func NewGetHostMigrationsV1InternalServerError() *GetHostMigrationsV1InternalServerError {
	return &GetHostMigrationsV1InternalServerError{}
}

/*
GetHostMigrationsV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetHostMigrationsV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIGetHostMigrationResponseV1
}

// IsSuccess returns true when this get host migrations v1 internal server error response has a 2xx status code
func (o *GetHostMigrationsV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get host migrations v1 internal server error response has a 3xx status code
func (o *GetHostMigrationsV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get host migrations v1 internal server error response has a 4xx status code
func (o *GetHostMigrationsV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get host migrations v1 internal server error response has a 5xx status code
func (o *GetHostMigrationsV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get host migrations v1 internal server error response a status code equal to that given
func (o *GetHostMigrationsV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get host migrations v1 internal server error response
func (o *GetHostMigrationsV1InternalServerError) Code() int {
	return 500
}

func (o *GetHostMigrationsV1InternalServerError) Error() string {
	return fmt.Sprintf("[POST /host-migration/entities/host-migrations/GET/v1][%d] getHostMigrationsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetHostMigrationsV1InternalServerError) String() string {
	return fmt.Sprintf("[POST /host-migration/entities/host-migrations/GET/v1][%d] getHostMigrationsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetHostMigrationsV1InternalServerError) GetPayload() *models.APIGetHostMigrationResponseV1 {
	return o.Payload
}

func (o *GetHostMigrationsV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIGetHostMigrationResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
