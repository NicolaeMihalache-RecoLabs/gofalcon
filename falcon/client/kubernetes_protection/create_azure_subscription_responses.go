// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_protection

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

// CreateAzureSubscriptionReader is a Reader for the CreateAzureSubscription structure.
type CreateAzureSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAzureSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAzureSubscriptionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewCreateAzureSubscriptionMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAzureSubscriptionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateAzureSubscriptionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateAzureSubscriptionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateAzureSubscriptionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /kubernetes-protection/entities/accounts/azure/v1] CreateAzureSubscription", response, response.Code())
	}
}

// NewCreateAzureSubscriptionCreated creates a CreateAzureSubscriptionCreated with default headers values
func NewCreateAzureSubscriptionCreated() *CreateAzureSubscriptionCreated {
	return &CreateAzureSubscriptionCreated{}
}

/*
CreateAzureSubscriptionCreated describes a response with status code 201, with default header values.

Created
*/
type CreateAzureSubscriptionCreated struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaBaseEntitiesResponse
}

// IsSuccess returns true when this create azure subscription created response has a 2xx status code
func (o *CreateAzureSubscriptionCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create azure subscription created response has a 3xx status code
func (o *CreateAzureSubscriptionCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create azure subscription created response has a 4xx status code
func (o *CreateAzureSubscriptionCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create azure subscription created response has a 5xx status code
func (o *CreateAzureSubscriptionCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create azure subscription created response a status code equal to that given
func (o *CreateAzureSubscriptionCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create azure subscription created response
func (o *CreateAzureSubscriptionCreated) Code() int {
	return 201
}

func (o *CreateAzureSubscriptionCreated) Error() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/accounts/azure/v1][%d] createAzureSubscriptionCreated  %+v", 201, o.Payload)
}

func (o *CreateAzureSubscriptionCreated) String() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/accounts/azure/v1][%d] createAzureSubscriptionCreated  %+v", 201, o.Payload)
}

func (o *CreateAzureSubscriptionCreated) GetPayload() *models.MsaBaseEntitiesResponse {
	return o.Payload
}

func (o *CreateAzureSubscriptionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaBaseEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAzureSubscriptionMultiStatus creates a CreateAzureSubscriptionMultiStatus with default headers values
func NewCreateAzureSubscriptionMultiStatus() *CreateAzureSubscriptionMultiStatus {
	return &CreateAzureSubscriptionMultiStatus{}
}

/*
CreateAzureSubscriptionMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type CreateAzureSubscriptionMultiStatus struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaBaseEntitiesResponse
}

// IsSuccess returns true when this create azure subscription multi status response has a 2xx status code
func (o *CreateAzureSubscriptionMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create azure subscription multi status response has a 3xx status code
func (o *CreateAzureSubscriptionMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create azure subscription multi status response has a 4xx status code
func (o *CreateAzureSubscriptionMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this create azure subscription multi status response has a 5xx status code
func (o *CreateAzureSubscriptionMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this create azure subscription multi status response a status code equal to that given
func (o *CreateAzureSubscriptionMultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the create azure subscription multi status response
func (o *CreateAzureSubscriptionMultiStatus) Code() int {
	return 207
}

func (o *CreateAzureSubscriptionMultiStatus) Error() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/accounts/azure/v1][%d] createAzureSubscriptionMultiStatus  %+v", 207, o.Payload)
}

func (o *CreateAzureSubscriptionMultiStatus) String() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/accounts/azure/v1][%d] createAzureSubscriptionMultiStatus  %+v", 207, o.Payload)
}

func (o *CreateAzureSubscriptionMultiStatus) GetPayload() *models.MsaBaseEntitiesResponse {
	return o.Payload
}

func (o *CreateAzureSubscriptionMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaBaseEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAzureSubscriptionBadRequest creates a CreateAzureSubscriptionBadRequest with default headers values
func NewCreateAzureSubscriptionBadRequest() *CreateAzureSubscriptionBadRequest {
	return &CreateAzureSubscriptionBadRequest{}
}

/*
CreateAzureSubscriptionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateAzureSubscriptionBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaBaseEntitiesResponse
}

// IsSuccess returns true when this create azure subscription bad request response has a 2xx status code
func (o *CreateAzureSubscriptionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create azure subscription bad request response has a 3xx status code
func (o *CreateAzureSubscriptionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create azure subscription bad request response has a 4xx status code
func (o *CreateAzureSubscriptionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create azure subscription bad request response has a 5xx status code
func (o *CreateAzureSubscriptionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create azure subscription bad request response a status code equal to that given
func (o *CreateAzureSubscriptionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create azure subscription bad request response
func (o *CreateAzureSubscriptionBadRequest) Code() int {
	return 400
}

func (o *CreateAzureSubscriptionBadRequest) Error() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/accounts/azure/v1][%d] createAzureSubscriptionBadRequest  %+v", 400, o.Payload)
}

func (o *CreateAzureSubscriptionBadRequest) String() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/accounts/azure/v1][%d] createAzureSubscriptionBadRequest  %+v", 400, o.Payload)
}

func (o *CreateAzureSubscriptionBadRequest) GetPayload() *models.MsaBaseEntitiesResponse {
	return o.Payload
}

func (o *CreateAzureSubscriptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaBaseEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAzureSubscriptionForbidden creates a CreateAzureSubscriptionForbidden with default headers values
func NewCreateAzureSubscriptionForbidden() *CreateAzureSubscriptionForbidden {
	return &CreateAzureSubscriptionForbidden{}
}

/*
CreateAzureSubscriptionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateAzureSubscriptionForbidden struct {

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

// IsSuccess returns true when this create azure subscription forbidden response has a 2xx status code
func (o *CreateAzureSubscriptionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create azure subscription forbidden response has a 3xx status code
func (o *CreateAzureSubscriptionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create azure subscription forbidden response has a 4xx status code
func (o *CreateAzureSubscriptionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create azure subscription forbidden response has a 5xx status code
func (o *CreateAzureSubscriptionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create azure subscription forbidden response a status code equal to that given
func (o *CreateAzureSubscriptionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create azure subscription forbidden response
func (o *CreateAzureSubscriptionForbidden) Code() int {
	return 403
}

func (o *CreateAzureSubscriptionForbidden) Error() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/accounts/azure/v1][%d] createAzureSubscriptionForbidden  %+v", 403, o.Payload)
}

func (o *CreateAzureSubscriptionForbidden) String() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/accounts/azure/v1][%d] createAzureSubscriptionForbidden  %+v", 403, o.Payload)
}

func (o *CreateAzureSubscriptionForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateAzureSubscriptionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateAzureSubscriptionTooManyRequests creates a CreateAzureSubscriptionTooManyRequests with default headers values
func NewCreateAzureSubscriptionTooManyRequests() *CreateAzureSubscriptionTooManyRequests {
	return &CreateAzureSubscriptionTooManyRequests{}
}

/*
CreateAzureSubscriptionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateAzureSubscriptionTooManyRequests struct {

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

// IsSuccess returns true when this create azure subscription too many requests response has a 2xx status code
func (o *CreateAzureSubscriptionTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create azure subscription too many requests response has a 3xx status code
func (o *CreateAzureSubscriptionTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create azure subscription too many requests response has a 4xx status code
func (o *CreateAzureSubscriptionTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create azure subscription too many requests response has a 5xx status code
func (o *CreateAzureSubscriptionTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create azure subscription too many requests response a status code equal to that given
func (o *CreateAzureSubscriptionTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the create azure subscription too many requests response
func (o *CreateAzureSubscriptionTooManyRequests) Code() int {
	return 429
}

func (o *CreateAzureSubscriptionTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/accounts/azure/v1][%d] createAzureSubscriptionTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateAzureSubscriptionTooManyRequests) String() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/accounts/azure/v1][%d] createAzureSubscriptionTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateAzureSubscriptionTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateAzureSubscriptionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateAzureSubscriptionInternalServerError creates a CreateAzureSubscriptionInternalServerError with default headers values
func NewCreateAzureSubscriptionInternalServerError() *CreateAzureSubscriptionInternalServerError {
	return &CreateAzureSubscriptionInternalServerError{}
}

/*
CreateAzureSubscriptionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateAzureSubscriptionInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaBaseEntitiesResponse
}

// IsSuccess returns true when this create azure subscription internal server error response has a 2xx status code
func (o *CreateAzureSubscriptionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create azure subscription internal server error response has a 3xx status code
func (o *CreateAzureSubscriptionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create azure subscription internal server error response has a 4xx status code
func (o *CreateAzureSubscriptionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create azure subscription internal server error response has a 5xx status code
func (o *CreateAzureSubscriptionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create azure subscription internal server error response a status code equal to that given
func (o *CreateAzureSubscriptionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create azure subscription internal server error response
func (o *CreateAzureSubscriptionInternalServerError) Code() int {
	return 500
}

func (o *CreateAzureSubscriptionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/accounts/azure/v1][%d] createAzureSubscriptionInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateAzureSubscriptionInternalServerError) String() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/accounts/azure/v1][%d] createAzureSubscriptionInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateAzureSubscriptionInternalServerError) GetPayload() *models.MsaBaseEntitiesResponse {
	return o.Payload
}

func (o *CreateAzureSubscriptionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaBaseEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
