// Code generated by go-swagger; DO NOT EDIT.

package firewall_management

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

// UpdateNetworkLocationsMetadataReader is a Reader for the UpdateNetworkLocationsMetadata structure.
type UpdateNetworkLocationsMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkLocationsMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkLocationsMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateNetworkLocationsMetadataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateNetworkLocationsMetadataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateNetworkLocationsMetadataTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /fwmgr/entities/network-locations-metadata/v1] update-network-locations-metadata", response, response.Code())
	}
}

// NewUpdateNetworkLocationsMetadataOK creates a UpdateNetworkLocationsMetadataOK with default headers values
func NewUpdateNetworkLocationsMetadataOK() *UpdateNetworkLocationsMetadataOK {
	return &UpdateNetworkLocationsMetadataOK{}
}

/*
UpdateNetworkLocationsMetadataOK describes a response with status code 200, with default header values.

OK
*/
type UpdateNetworkLocationsMetadataOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FwmgrMsaspecQueryResponse
}

// IsSuccess returns true when this update network locations metadata o k response has a 2xx status code
func (o *UpdateNetworkLocationsMetadataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network locations metadata o k response has a 3xx status code
func (o *UpdateNetworkLocationsMetadataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network locations metadata o k response has a 4xx status code
func (o *UpdateNetworkLocationsMetadataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network locations metadata o k response has a 5xx status code
func (o *UpdateNetworkLocationsMetadataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network locations metadata o k response a status code equal to that given
func (o *UpdateNetworkLocationsMetadataOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network locations metadata o k response
func (o *UpdateNetworkLocationsMetadataOK) Code() int {
	return 200
}

func (o *UpdateNetworkLocationsMetadataOK) Error() string {
	return fmt.Sprintf("[POST /fwmgr/entities/network-locations-metadata/v1][%d] updateNetworkLocationsMetadataOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkLocationsMetadataOK) String() string {
	return fmt.Sprintf("[POST /fwmgr/entities/network-locations-metadata/v1][%d] updateNetworkLocationsMetadataOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkLocationsMetadataOK) GetPayload() *models.FwmgrMsaspecQueryResponse {
	return o.Payload
}

func (o *UpdateNetworkLocationsMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FwmgrMsaspecQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNetworkLocationsMetadataBadRequest creates a UpdateNetworkLocationsMetadataBadRequest with default headers values
func NewUpdateNetworkLocationsMetadataBadRequest() *UpdateNetworkLocationsMetadataBadRequest {
	return &UpdateNetworkLocationsMetadataBadRequest{}
}

/*
UpdateNetworkLocationsMetadataBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateNetworkLocationsMetadataBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FwmgrMsaspecResponseFields
}

// IsSuccess returns true when this update network locations metadata bad request response has a 2xx status code
func (o *UpdateNetworkLocationsMetadataBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update network locations metadata bad request response has a 3xx status code
func (o *UpdateNetworkLocationsMetadataBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network locations metadata bad request response has a 4xx status code
func (o *UpdateNetworkLocationsMetadataBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update network locations metadata bad request response has a 5xx status code
func (o *UpdateNetworkLocationsMetadataBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update network locations metadata bad request response a status code equal to that given
func (o *UpdateNetworkLocationsMetadataBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update network locations metadata bad request response
func (o *UpdateNetworkLocationsMetadataBadRequest) Code() int {
	return 400
}

func (o *UpdateNetworkLocationsMetadataBadRequest) Error() string {
	return fmt.Sprintf("[POST /fwmgr/entities/network-locations-metadata/v1][%d] updateNetworkLocationsMetadataBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateNetworkLocationsMetadataBadRequest) String() string {
	return fmt.Sprintf("[POST /fwmgr/entities/network-locations-metadata/v1][%d] updateNetworkLocationsMetadataBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateNetworkLocationsMetadataBadRequest) GetPayload() *models.FwmgrMsaspecResponseFields {
	return o.Payload
}

func (o *UpdateNetworkLocationsMetadataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FwmgrMsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNetworkLocationsMetadataForbidden creates a UpdateNetworkLocationsMetadataForbidden with default headers values
func NewUpdateNetworkLocationsMetadataForbidden() *UpdateNetworkLocationsMetadataForbidden {
	return &UpdateNetworkLocationsMetadataForbidden{}
}

/*
UpdateNetworkLocationsMetadataForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateNetworkLocationsMetadataForbidden struct {

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

// IsSuccess returns true when this update network locations metadata forbidden response has a 2xx status code
func (o *UpdateNetworkLocationsMetadataForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update network locations metadata forbidden response has a 3xx status code
func (o *UpdateNetworkLocationsMetadataForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network locations metadata forbidden response has a 4xx status code
func (o *UpdateNetworkLocationsMetadataForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update network locations metadata forbidden response has a 5xx status code
func (o *UpdateNetworkLocationsMetadataForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update network locations metadata forbidden response a status code equal to that given
func (o *UpdateNetworkLocationsMetadataForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update network locations metadata forbidden response
func (o *UpdateNetworkLocationsMetadataForbidden) Code() int {
	return 403
}

func (o *UpdateNetworkLocationsMetadataForbidden) Error() string {
	return fmt.Sprintf("[POST /fwmgr/entities/network-locations-metadata/v1][%d] updateNetworkLocationsMetadataForbidden  %+v", 403, o.Payload)
}

func (o *UpdateNetworkLocationsMetadataForbidden) String() string {
	return fmt.Sprintf("[POST /fwmgr/entities/network-locations-metadata/v1][%d] updateNetworkLocationsMetadataForbidden  %+v", 403, o.Payload)
}

func (o *UpdateNetworkLocationsMetadataForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateNetworkLocationsMetadataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateNetworkLocationsMetadataTooManyRequests creates a UpdateNetworkLocationsMetadataTooManyRequests with default headers values
func NewUpdateNetworkLocationsMetadataTooManyRequests() *UpdateNetworkLocationsMetadataTooManyRequests {
	return &UpdateNetworkLocationsMetadataTooManyRequests{}
}

/*
UpdateNetworkLocationsMetadataTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateNetworkLocationsMetadataTooManyRequests struct {

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

// IsSuccess returns true when this update network locations metadata too many requests response has a 2xx status code
func (o *UpdateNetworkLocationsMetadataTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update network locations metadata too many requests response has a 3xx status code
func (o *UpdateNetworkLocationsMetadataTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network locations metadata too many requests response has a 4xx status code
func (o *UpdateNetworkLocationsMetadataTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update network locations metadata too many requests response has a 5xx status code
func (o *UpdateNetworkLocationsMetadataTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update network locations metadata too many requests response a status code equal to that given
func (o *UpdateNetworkLocationsMetadataTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the update network locations metadata too many requests response
func (o *UpdateNetworkLocationsMetadataTooManyRequests) Code() int {
	return 429
}

func (o *UpdateNetworkLocationsMetadataTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fwmgr/entities/network-locations-metadata/v1][%d] updateNetworkLocationsMetadataTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateNetworkLocationsMetadataTooManyRequests) String() string {
	return fmt.Sprintf("[POST /fwmgr/entities/network-locations-metadata/v1][%d] updateNetworkLocationsMetadataTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateNetworkLocationsMetadataTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateNetworkLocationsMetadataTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
