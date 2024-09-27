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

// GetAzureTenantConfigReader is a Reader for the GetAzureTenantConfig structure.
type GetAzureTenantConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAzureTenantConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAzureTenantConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewGetAzureTenantConfigMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAzureTenantConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAzureTenantConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAzureTenantConfigTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAzureTenantConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /kubernetes-protection/entities/config/azure/v1] GetAzureTenantConfig", response, response.Code())
	}
}

// NewGetAzureTenantConfigOK creates a GetAzureTenantConfigOK with default headers values
func NewGetAzureTenantConfigOK() *GetAzureTenantConfigOK {
	return &GetAzureTenantConfigOK{}
}

/*
GetAzureTenantConfigOK describes a response with status code 200, with default header values.

OK
*/
type GetAzureTenantConfigOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.K8sregGetAzureTenantConfigResp
}

// IsSuccess returns true when this get azure tenant config o k response has a 2xx status code
func (o *GetAzureTenantConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get azure tenant config o k response has a 3xx status code
func (o *GetAzureTenantConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get azure tenant config o k response has a 4xx status code
func (o *GetAzureTenantConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get azure tenant config o k response has a 5xx status code
func (o *GetAzureTenantConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get azure tenant config o k response a status code equal to that given
func (o *GetAzureTenantConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get azure tenant config o k response
func (o *GetAzureTenantConfigOK) Code() int {
	return 200
}

func (o *GetAzureTenantConfigOK) Error() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/config/azure/v1][%d] getAzureTenantConfigOK  %+v", 200, o.Payload)
}

func (o *GetAzureTenantConfigOK) String() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/config/azure/v1][%d] getAzureTenantConfigOK  %+v", 200, o.Payload)
}

func (o *GetAzureTenantConfigOK) GetPayload() *models.K8sregGetAzureTenantConfigResp {
	return o.Payload
}

func (o *GetAzureTenantConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.K8sregGetAzureTenantConfigResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAzureTenantConfigMultiStatus creates a GetAzureTenantConfigMultiStatus with default headers values
func NewGetAzureTenantConfigMultiStatus() *GetAzureTenantConfigMultiStatus {
	return &GetAzureTenantConfigMultiStatus{}
}

/*
GetAzureTenantConfigMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type GetAzureTenantConfigMultiStatus struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.K8sregGetAzureTenantConfigResp
}

// IsSuccess returns true when this get azure tenant config multi status response has a 2xx status code
func (o *GetAzureTenantConfigMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get azure tenant config multi status response has a 3xx status code
func (o *GetAzureTenantConfigMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get azure tenant config multi status response has a 4xx status code
func (o *GetAzureTenantConfigMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this get azure tenant config multi status response has a 5xx status code
func (o *GetAzureTenantConfigMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this get azure tenant config multi status response a status code equal to that given
func (o *GetAzureTenantConfigMultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the get azure tenant config multi status response
func (o *GetAzureTenantConfigMultiStatus) Code() int {
	return 207
}

func (o *GetAzureTenantConfigMultiStatus) Error() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/config/azure/v1][%d] getAzureTenantConfigMultiStatus  %+v", 207, o.Payload)
}

func (o *GetAzureTenantConfigMultiStatus) String() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/config/azure/v1][%d] getAzureTenantConfigMultiStatus  %+v", 207, o.Payload)
}

func (o *GetAzureTenantConfigMultiStatus) GetPayload() *models.K8sregGetAzureTenantConfigResp {
	return o.Payload
}

func (o *GetAzureTenantConfigMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.K8sregGetAzureTenantConfigResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAzureTenantConfigBadRequest creates a GetAzureTenantConfigBadRequest with default headers values
func NewGetAzureTenantConfigBadRequest() *GetAzureTenantConfigBadRequest {
	return &GetAzureTenantConfigBadRequest{}
}

/*
GetAzureTenantConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAzureTenantConfigBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.K8sregGetAzureTenantConfigResp
}

// IsSuccess returns true when this get azure tenant config bad request response has a 2xx status code
func (o *GetAzureTenantConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get azure tenant config bad request response has a 3xx status code
func (o *GetAzureTenantConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get azure tenant config bad request response has a 4xx status code
func (o *GetAzureTenantConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get azure tenant config bad request response has a 5xx status code
func (o *GetAzureTenantConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get azure tenant config bad request response a status code equal to that given
func (o *GetAzureTenantConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get azure tenant config bad request response
func (o *GetAzureTenantConfigBadRequest) Code() int {
	return 400
}

func (o *GetAzureTenantConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/config/azure/v1][%d] getAzureTenantConfigBadRequest  %+v", 400, o.Payload)
}

func (o *GetAzureTenantConfigBadRequest) String() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/config/azure/v1][%d] getAzureTenantConfigBadRequest  %+v", 400, o.Payload)
}

func (o *GetAzureTenantConfigBadRequest) GetPayload() *models.K8sregGetAzureTenantConfigResp {
	return o.Payload
}

func (o *GetAzureTenantConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.K8sregGetAzureTenantConfigResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAzureTenantConfigForbidden creates a GetAzureTenantConfigForbidden with default headers values
func NewGetAzureTenantConfigForbidden() *GetAzureTenantConfigForbidden {
	return &GetAzureTenantConfigForbidden{}
}

/*
GetAzureTenantConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAzureTenantConfigForbidden struct {

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

// IsSuccess returns true when this get azure tenant config forbidden response has a 2xx status code
func (o *GetAzureTenantConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get azure tenant config forbidden response has a 3xx status code
func (o *GetAzureTenantConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get azure tenant config forbidden response has a 4xx status code
func (o *GetAzureTenantConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get azure tenant config forbidden response has a 5xx status code
func (o *GetAzureTenantConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get azure tenant config forbidden response a status code equal to that given
func (o *GetAzureTenantConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get azure tenant config forbidden response
func (o *GetAzureTenantConfigForbidden) Code() int {
	return 403
}

func (o *GetAzureTenantConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/config/azure/v1][%d] getAzureTenantConfigForbidden  %+v", 403, o.Payload)
}

func (o *GetAzureTenantConfigForbidden) String() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/config/azure/v1][%d] getAzureTenantConfigForbidden  %+v", 403, o.Payload)
}

func (o *GetAzureTenantConfigForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetAzureTenantConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAzureTenantConfigTooManyRequests creates a GetAzureTenantConfigTooManyRequests with default headers values
func NewGetAzureTenantConfigTooManyRequests() *GetAzureTenantConfigTooManyRequests {
	return &GetAzureTenantConfigTooManyRequests{}
}

/*
GetAzureTenantConfigTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetAzureTenantConfigTooManyRequests struct {

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

// IsSuccess returns true when this get azure tenant config too many requests response has a 2xx status code
func (o *GetAzureTenantConfigTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get azure tenant config too many requests response has a 3xx status code
func (o *GetAzureTenantConfigTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get azure tenant config too many requests response has a 4xx status code
func (o *GetAzureTenantConfigTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get azure tenant config too many requests response has a 5xx status code
func (o *GetAzureTenantConfigTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get azure tenant config too many requests response a status code equal to that given
func (o *GetAzureTenantConfigTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get azure tenant config too many requests response
func (o *GetAzureTenantConfigTooManyRequests) Code() int {
	return 429
}

func (o *GetAzureTenantConfigTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/config/azure/v1][%d] getAzureTenantConfigTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAzureTenantConfigTooManyRequests) String() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/config/azure/v1][%d] getAzureTenantConfigTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAzureTenantConfigTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetAzureTenantConfigTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAzureTenantConfigInternalServerError creates a GetAzureTenantConfigInternalServerError with default headers values
func NewGetAzureTenantConfigInternalServerError() *GetAzureTenantConfigInternalServerError {
	return &GetAzureTenantConfigInternalServerError{}
}

/*
GetAzureTenantConfigInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAzureTenantConfigInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.K8sregGetAzureTenantConfigResp
}

// IsSuccess returns true when this get azure tenant config internal server error response has a 2xx status code
func (o *GetAzureTenantConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get azure tenant config internal server error response has a 3xx status code
func (o *GetAzureTenantConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get azure tenant config internal server error response has a 4xx status code
func (o *GetAzureTenantConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get azure tenant config internal server error response has a 5xx status code
func (o *GetAzureTenantConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get azure tenant config internal server error response a status code equal to that given
func (o *GetAzureTenantConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get azure tenant config internal server error response
func (o *GetAzureTenantConfigInternalServerError) Code() int {
	return 500
}

func (o *GetAzureTenantConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/config/azure/v1][%d] getAzureTenantConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAzureTenantConfigInternalServerError) String() string {
	return fmt.Sprintf("[GET /kubernetes-protection/entities/config/azure/v1][%d] getAzureTenantConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAzureTenantConfigInternalServerError) GetPayload() *models.K8sregGetAzureTenantConfigResp {
	return o.Payload
}

func (o *GetAzureTenantConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.K8sregGetAzureTenantConfigResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
