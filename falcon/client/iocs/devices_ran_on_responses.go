// Code generated by go-swagger; DO NOT EDIT.

package iocs

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

// DevicesRanOnReader is a Reader for the DevicesRanOn structure.
type DevicesRanOnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DevicesRanOnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDevicesRanOnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDevicesRanOnForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDevicesRanOnTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /indicators/queries/devices/v1] DevicesRanOn", response, response.Code())
	}
}

// NewDevicesRanOnOK creates a DevicesRanOnOK with default headers values
func NewDevicesRanOnOK() *DevicesRanOnOK {
	return &DevicesRanOnOK{}
}

/*
DevicesRanOnOK describes a response with status code 200, with default header values.

OK
*/
type DevicesRanOnOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.IocapiMsaReplyDevicesRanOn
}

// IsSuccess returns true when this devices ran on o k response has a 2xx status code
func (o *DevicesRanOnOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this devices ran on o k response has a 3xx status code
func (o *DevicesRanOnOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this devices ran on o k response has a 4xx status code
func (o *DevicesRanOnOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this devices ran on o k response has a 5xx status code
func (o *DevicesRanOnOK) IsServerError() bool {
	return false
}

// IsCode returns true when this devices ran on o k response a status code equal to that given
func (o *DevicesRanOnOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the devices ran on o k response
func (o *DevicesRanOnOK) Code() int {
	return 200
}

func (o *DevicesRanOnOK) Error() string {
	return fmt.Sprintf("[GET /indicators/queries/devices/v1][%d] devicesRanOnOK  %+v", 200, o.Payload)
}

func (o *DevicesRanOnOK) String() string {
	return fmt.Sprintf("[GET /indicators/queries/devices/v1][%d] devicesRanOnOK  %+v", 200, o.Payload)
}

func (o *DevicesRanOnOK) GetPayload() *models.IocapiMsaReplyDevicesRanOn {
	return o.Payload
}

func (o *DevicesRanOnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.IocapiMsaReplyDevicesRanOn)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDevicesRanOnForbidden creates a DevicesRanOnForbidden with default headers values
func NewDevicesRanOnForbidden() *DevicesRanOnForbidden {
	return &DevicesRanOnForbidden{}
}

/*
DevicesRanOnForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DevicesRanOnForbidden struct {

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

// IsSuccess returns true when this devices ran on forbidden response has a 2xx status code
func (o *DevicesRanOnForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this devices ran on forbidden response has a 3xx status code
func (o *DevicesRanOnForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this devices ran on forbidden response has a 4xx status code
func (o *DevicesRanOnForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this devices ran on forbidden response has a 5xx status code
func (o *DevicesRanOnForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this devices ran on forbidden response a status code equal to that given
func (o *DevicesRanOnForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the devices ran on forbidden response
func (o *DevicesRanOnForbidden) Code() int {
	return 403
}

func (o *DevicesRanOnForbidden) Error() string {
	return fmt.Sprintf("[GET /indicators/queries/devices/v1][%d] devicesRanOnForbidden  %+v", 403, o.Payload)
}

func (o *DevicesRanOnForbidden) String() string {
	return fmt.Sprintf("[GET /indicators/queries/devices/v1][%d] devicesRanOnForbidden  %+v", 403, o.Payload)
}

func (o *DevicesRanOnForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DevicesRanOnForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDevicesRanOnTooManyRequests creates a DevicesRanOnTooManyRequests with default headers values
func NewDevicesRanOnTooManyRequests() *DevicesRanOnTooManyRequests {
	return &DevicesRanOnTooManyRequests{}
}

/*
DevicesRanOnTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DevicesRanOnTooManyRequests struct {

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

// IsSuccess returns true when this devices ran on too many requests response has a 2xx status code
func (o *DevicesRanOnTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this devices ran on too many requests response has a 3xx status code
func (o *DevicesRanOnTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this devices ran on too many requests response has a 4xx status code
func (o *DevicesRanOnTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this devices ran on too many requests response has a 5xx status code
func (o *DevicesRanOnTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this devices ran on too many requests response a status code equal to that given
func (o *DevicesRanOnTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the devices ran on too many requests response
func (o *DevicesRanOnTooManyRequests) Code() int {
	return 429
}

func (o *DevicesRanOnTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /indicators/queries/devices/v1][%d] devicesRanOnTooManyRequests  %+v", 429, o.Payload)
}

func (o *DevicesRanOnTooManyRequests) String() string {
	return fmt.Sprintf("[GET /indicators/queries/devices/v1][%d] devicesRanOnTooManyRequests  %+v", 429, o.Payload)
}

func (o *DevicesRanOnTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DevicesRanOnTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
