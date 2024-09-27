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

// KubernetesIomEntitiesCombinedReader is a Reader for the KubernetesIomEntitiesCombined structure.
type KubernetesIomEntitiesCombinedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesIomEntitiesCombinedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesIomEntitiesCombinedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewKubernetesIomEntitiesCombinedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewKubernetesIomEntitiesCombinedTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesIomEntitiesCombinedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/combined/kubernetes-ioms/v1] KubernetesIomEntitiesCombined", response, response.Code())
	}
}

// NewKubernetesIomEntitiesCombinedOK creates a KubernetesIomEntitiesCombinedOK with default headers values
func NewKubernetesIomEntitiesCombinedOK() *KubernetesIomEntitiesCombinedOK {
	return &KubernetesIomEntitiesCombinedOK{}
}

/*
KubernetesIomEntitiesCombinedOK describes a response with status code 200, with default header values.

OK
*/
type KubernetesIomEntitiesCombinedOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.K8siomsKubernetesIOMEntityResponse
}

// IsSuccess returns true when this kubernetes iom entities combined o k response has a 2xx status code
func (o *KubernetesIomEntitiesCombinedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes iom entities combined o k response has a 3xx status code
func (o *KubernetesIomEntitiesCombinedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes iom entities combined o k response has a 4xx status code
func (o *KubernetesIomEntitiesCombinedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes iom entities combined o k response has a 5xx status code
func (o *KubernetesIomEntitiesCombinedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes iom entities combined o k response a status code equal to that given
func (o *KubernetesIomEntitiesCombinedOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kubernetes iom entities combined o k response
func (o *KubernetesIomEntitiesCombinedOK) Code() int {
	return 200
}

func (o *KubernetesIomEntitiesCombinedOK) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/kubernetes-ioms/v1][%d] kubernetesIomEntitiesCombinedOK  %+v", 200, o.Payload)
}

func (o *KubernetesIomEntitiesCombinedOK) String() string {
	return fmt.Sprintf("[GET /container-security/combined/kubernetes-ioms/v1][%d] kubernetesIomEntitiesCombinedOK  %+v", 200, o.Payload)
}

func (o *KubernetesIomEntitiesCombinedOK) GetPayload() *models.K8siomsKubernetesIOMEntityResponse {
	return o.Payload
}

func (o *KubernetesIomEntitiesCombinedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.K8siomsKubernetesIOMEntityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesIomEntitiesCombinedForbidden creates a KubernetesIomEntitiesCombinedForbidden with default headers values
func NewKubernetesIomEntitiesCombinedForbidden() *KubernetesIomEntitiesCombinedForbidden {
	return &KubernetesIomEntitiesCombinedForbidden{}
}

/*
KubernetesIomEntitiesCombinedForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesIomEntitiesCombinedForbidden struct {

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

// IsSuccess returns true when this kubernetes iom entities combined forbidden response has a 2xx status code
func (o *KubernetesIomEntitiesCombinedForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes iom entities combined forbidden response has a 3xx status code
func (o *KubernetesIomEntitiesCombinedForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes iom entities combined forbidden response has a 4xx status code
func (o *KubernetesIomEntitiesCombinedForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes iom entities combined forbidden response has a 5xx status code
func (o *KubernetesIomEntitiesCombinedForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes iom entities combined forbidden response a status code equal to that given
func (o *KubernetesIomEntitiesCombinedForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the kubernetes iom entities combined forbidden response
func (o *KubernetesIomEntitiesCombinedForbidden) Code() int {
	return 403
}

func (o *KubernetesIomEntitiesCombinedForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/kubernetes-ioms/v1][%d] kubernetesIomEntitiesCombinedForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesIomEntitiesCombinedForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/combined/kubernetes-ioms/v1][%d] kubernetesIomEntitiesCombinedForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesIomEntitiesCombinedForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *KubernetesIomEntitiesCombinedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewKubernetesIomEntitiesCombinedTooManyRequests creates a KubernetesIomEntitiesCombinedTooManyRequests with default headers values
func NewKubernetesIomEntitiesCombinedTooManyRequests() *KubernetesIomEntitiesCombinedTooManyRequests {
	return &KubernetesIomEntitiesCombinedTooManyRequests{}
}

/*
KubernetesIomEntitiesCombinedTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type KubernetesIomEntitiesCombinedTooManyRequests struct {

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

// IsSuccess returns true when this kubernetes iom entities combined too many requests response has a 2xx status code
func (o *KubernetesIomEntitiesCombinedTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes iom entities combined too many requests response has a 3xx status code
func (o *KubernetesIomEntitiesCombinedTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes iom entities combined too many requests response has a 4xx status code
func (o *KubernetesIomEntitiesCombinedTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes iom entities combined too many requests response has a 5xx status code
func (o *KubernetesIomEntitiesCombinedTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes iom entities combined too many requests response a status code equal to that given
func (o *KubernetesIomEntitiesCombinedTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the kubernetes iom entities combined too many requests response
func (o *KubernetesIomEntitiesCombinedTooManyRequests) Code() int {
	return 429
}

func (o *KubernetesIomEntitiesCombinedTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/kubernetes-ioms/v1][%d] kubernetesIomEntitiesCombinedTooManyRequests  %+v", 429, o.Payload)
}

func (o *KubernetesIomEntitiesCombinedTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/combined/kubernetes-ioms/v1][%d] kubernetesIomEntitiesCombinedTooManyRequests  %+v", 429, o.Payload)
}

func (o *KubernetesIomEntitiesCombinedTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *KubernetesIomEntitiesCombinedTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewKubernetesIomEntitiesCombinedInternalServerError creates a KubernetesIomEntitiesCombinedInternalServerError with default headers values
func NewKubernetesIomEntitiesCombinedInternalServerError() *KubernetesIomEntitiesCombinedInternalServerError {
	return &KubernetesIomEntitiesCombinedInternalServerError{}
}

/*
KubernetesIomEntitiesCombinedInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type KubernetesIomEntitiesCombinedInternalServerError struct {

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

// IsSuccess returns true when this kubernetes iom entities combined internal server error response has a 2xx status code
func (o *KubernetesIomEntitiesCombinedInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes iom entities combined internal server error response has a 3xx status code
func (o *KubernetesIomEntitiesCombinedInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes iom entities combined internal server error response has a 4xx status code
func (o *KubernetesIomEntitiesCombinedInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes iom entities combined internal server error response has a 5xx status code
func (o *KubernetesIomEntitiesCombinedInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes iom entities combined internal server error response a status code equal to that given
func (o *KubernetesIomEntitiesCombinedInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the kubernetes iom entities combined internal server error response
func (o *KubernetesIomEntitiesCombinedInternalServerError) Code() int {
	return 500
}

func (o *KubernetesIomEntitiesCombinedInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/combined/kubernetes-ioms/v1][%d] kubernetesIomEntitiesCombinedInternalServerError  %+v", 500, o.Payload)
}

func (o *KubernetesIomEntitiesCombinedInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/combined/kubernetes-ioms/v1][%d] kubernetesIomEntitiesCombinedInternalServerError  %+v", 500, o.Payload)
}

func (o *KubernetesIomEntitiesCombinedInternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *KubernetesIomEntitiesCombinedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
