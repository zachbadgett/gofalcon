// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// QueryDeviceLoginHistoryReader is a Reader for the QueryDeviceLoginHistory structure.
type QueryDeviceLoginHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryDeviceLoginHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryDeviceLoginHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQueryDeviceLoginHistoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryDeviceLoginHistoryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryDeviceLoginHistoryOK creates a QueryDeviceLoginHistoryOK with default headers values
func NewQueryDeviceLoginHistoryOK() *QueryDeviceLoginHistoryOK {
	return &QueryDeviceLoginHistoryOK{}
}

/*
QueryDeviceLoginHistoryOK describes a response with status code 200, with default header values.

OK
*/
type QueryDeviceLoginHistoryOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DeviceapiLoginHistoryResponseV1
}

// IsSuccess returns true when this query device login history o k response has a 2xx status code
func (o *QueryDeviceLoginHistoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query device login history o k response has a 3xx status code
func (o *QueryDeviceLoginHistoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query device login history o k response has a 4xx status code
func (o *QueryDeviceLoginHistoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query device login history o k response has a 5xx status code
func (o *QueryDeviceLoginHistoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query device login history o k response a status code equal to that given
func (o *QueryDeviceLoginHistoryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query device login history o k response
func (o *QueryDeviceLoginHistoryOK) Code() int {
	return 200
}

func (o *QueryDeviceLoginHistoryOK) Error() string {
	return fmt.Sprintf("[POST /devices/combined/devices/login-history/v1][%d] queryDeviceLoginHistoryOK  %+v", 200, o.Payload)
}

func (o *QueryDeviceLoginHistoryOK) String() string {
	return fmt.Sprintf("[POST /devices/combined/devices/login-history/v1][%d] queryDeviceLoginHistoryOK  %+v", 200, o.Payload)
}

func (o *QueryDeviceLoginHistoryOK) GetPayload() *models.DeviceapiLoginHistoryResponseV1 {
	return o.Payload
}

func (o *QueryDeviceLoginHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DeviceapiLoginHistoryResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryDeviceLoginHistoryForbidden creates a QueryDeviceLoginHistoryForbidden with default headers values
func NewQueryDeviceLoginHistoryForbidden() *QueryDeviceLoginHistoryForbidden {
	return &QueryDeviceLoginHistoryForbidden{}
}

/*
QueryDeviceLoginHistoryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryDeviceLoginHistoryForbidden struct {

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

// IsSuccess returns true when this query device login history forbidden response has a 2xx status code
func (o *QueryDeviceLoginHistoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query device login history forbidden response has a 3xx status code
func (o *QueryDeviceLoginHistoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query device login history forbidden response has a 4xx status code
func (o *QueryDeviceLoginHistoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query device login history forbidden response has a 5xx status code
func (o *QueryDeviceLoginHistoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query device login history forbidden response a status code equal to that given
func (o *QueryDeviceLoginHistoryForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query device login history forbidden response
func (o *QueryDeviceLoginHistoryForbidden) Code() int {
	return 403
}

func (o *QueryDeviceLoginHistoryForbidden) Error() string {
	return fmt.Sprintf("[POST /devices/combined/devices/login-history/v1][%d] queryDeviceLoginHistoryForbidden  %+v", 403, o.Payload)
}

func (o *QueryDeviceLoginHistoryForbidden) String() string {
	return fmt.Sprintf("[POST /devices/combined/devices/login-history/v1][%d] queryDeviceLoginHistoryForbidden  %+v", 403, o.Payload)
}

func (o *QueryDeviceLoginHistoryForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryDeviceLoginHistoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryDeviceLoginHistoryTooManyRequests creates a QueryDeviceLoginHistoryTooManyRequests with default headers values
func NewQueryDeviceLoginHistoryTooManyRequests() *QueryDeviceLoginHistoryTooManyRequests {
	return &QueryDeviceLoginHistoryTooManyRequests{}
}

/*
QueryDeviceLoginHistoryTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryDeviceLoginHistoryTooManyRequests struct {

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

// IsSuccess returns true when this query device login history too many requests response has a 2xx status code
func (o *QueryDeviceLoginHistoryTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query device login history too many requests response has a 3xx status code
func (o *QueryDeviceLoginHistoryTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query device login history too many requests response has a 4xx status code
func (o *QueryDeviceLoginHistoryTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query device login history too many requests response has a 5xx status code
func (o *QueryDeviceLoginHistoryTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query device login history too many requests response a status code equal to that given
func (o *QueryDeviceLoginHistoryTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query device login history too many requests response
func (o *QueryDeviceLoginHistoryTooManyRequests) Code() int {
	return 429
}

func (o *QueryDeviceLoginHistoryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /devices/combined/devices/login-history/v1][%d] queryDeviceLoginHistoryTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryDeviceLoginHistoryTooManyRequests) String() string {
	return fmt.Sprintf("[POST /devices/combined/devices/login-history/v1][%d] queryDeviceLoginHistoryTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryDeviceLoginHistoryTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryDeviceLoginHistoryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
