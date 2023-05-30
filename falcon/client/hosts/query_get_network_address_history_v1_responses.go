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

// QueryGetNetworkAddressHistoryV1Reader is a Reader for the QueryGetNetworkAddressHistoryV1 structure.
type QueryGetNetworkAddressHistoryV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryGetNetworkAddressHistoryV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryGetNetworkAddressHistoryV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQueryGetNetworkAddressHistoryV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryGetNetworkAddressHistoryV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryGetNetworkAddressHistoryV1OK creates a QueryGetNetworkAddressHistoryV1OK with default headers values
func NewQueryGetNetworkAddressHistoryV1OK() *QueryGetNetworkAddressHistoryV1OK {
	return &QueryGetNetworkAddressHistoryV1OK{}
}

/*
QueryGetNetworkAddressHistoryV1OK describes a response with status code 200, with default header values.

OK
*/
type QueryGetNetworkAddressHistoryV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DeviceapiNetworkAddressHistoryResponseV1
}

// IsSuccess returns true when this query get network address history v1 o k response has a 2xx status code
func (o *QueryGetNetworkAddressHistoryV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query get network address history v1 o k response has a 3xx status code
func (o *QueryGetNetworkAddressHistoryV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query get network address history v1 o k response has a 4xx status code
func (o *QueryGetNetworkAddressHistoryV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query get network address history v1 o k response has a 5xx status code
func (o *QueryGetNetworkAddressHistoryV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this query get network address history v1 o k response a status code equal to that given
func (o *QueryGetNetworkAddressHistoryV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query get network address history v1 o k response
func (o *QueryGetNetworkAddressHistoryV1OK) Code() int {
	return 200
}

func (o *QueryGetNetworkAddressHistoryV1OK) Error() string {
	return fmt.Sprintf("[POST /devices/combined/devices/network-address-history/v1][%d] queryGetNetworkAddressHistoryV1OK  %+v", 200, o.Payload)
}

func (o *QueryGetNetworkAddressHistoryV1OK) String() string {
	return fmt.Sprintf("[POST /devices/combined/devices/network-address-history/v1][%d] queryGetNetworkAddressHistoryV1OK  %+v", 200, o.Payload)
}

func (o *QueryGetNetworkAddressHistoryV1OK) GetPayload() *models.DeviceapiNetworkAddressHistoryResponseV1 {
	return o.Payload
}

func (o *QueryGetNetworkAddressHistoryV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DeviceapiNetworkAddressHistoryResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryGetNetworkAddressHistoryV1Forbidden creates a QueryGetNetworkAddressHistoryV1Forbidden with default headers values
func NewQueryGetNetworkAddressHistoryV1Forbidden() *QueryGetNetworkAddressHistoryV1Forbidden {
	return &QueryGetNetworkAddressHistoryV1Forbidden{}
}

/*
QueryGetNetworkAddressHistoryV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryGetNetworkAddressHistoryV1Forbidden struct {

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

// IsSuccess returns true when this query get network address history v1 forbidden response has a 2xx status code
func (o *QueryGetNetworkAddressHistoryV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query get network address history v1 forbidden response has a 3xx status code
func (o *QueryGetNetworkAddressHistoryV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query get network address history v1 forbidden response has a 4xx status code
func (o *QueryGetNetworkAddressHistoryV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query get network address history v1 forbidden response has a 5xx status code
func (o *QueryGetNetworkAddressHistoryV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query get network address history v1 forbidden response a status code equal to that given
func (o *QueryGetNetworkAddressHistoryV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query get network address history v1 forbidden response
func (o *QueryGetNetworkAddressHistoryV1Forbidden) Code() int {
	return 403
}

func (o *QueryGetNetworkAddressHistoryV1Forbidden) Error() string {
	return fmt.Sprintf("[POST /devices/combined/devices/network-address-history/v1][%d] queryGetNetworkAddressHistoryV1Forbidden  %+v", 403, o.Payload)
}

func (o *QueryGetNetworkAddressHistoryV1Forbidden) String() string {
	return fmt.Sprintf("[POST /devices/combined/devices/network-address-history/v1][%d] queryGetNetworkAddressHistoryV1Forbidden  %+v", 403, o.Payload)
}

func (o *QueryGetNetworkAddressHistoryV1Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryGetNetworkAddressHistoryV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryGetNetworkAddressHistoryV1TooManyRequests creates a QueryGetNetworkAddressHistoryV1TooManyRequests with default headers values
func NewQueryGetNetworkAddressHistoryV1TooManyRequests() *QueryGetNetworkAddressHistoryV1TooManyRequests {
	return &QueryGetNetworkAddressHistoryV1TooManyRequests{}
}

/*
QueryGetNetworkAddressHistoryV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryGetNetworkAddressHistoryV1TooManyRequests struct {

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

// IsSuccess returns true when this query get network address history v1 too many requests response has a 2xx status code
func (o *QueryGetNetworkAddressHistoryV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query get network address history v1 too many requests response has a 3xx status code
func (o *QueryGetNetworkAddressHistoryV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query get network address history v1 too many requests response has a 4xx status code
func (o *QueryGetNetworkAddressHistoryV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query get network address history v1 too many requests response has a 5xx status code
func (o *QueryGetNetworkAddressHistoryV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query get network address history v1 too many requests response a status code equal to that given
func (o *QueryGetNetworkAddressHistoryV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query get network address history v1 too many requests response
func (o *QueryGetNetworkAddressHistoryV1TooManyRequests) Code() int {
	return 429
}

func (o *QueryGetNetworkAddressHistoryV1TooManyRequests) Error() string {
	return fmt.Sprintf("[POST /devices/combined/devices/network-address-history/v1][%d] queryGetNetworkAddressHistoryV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryGetNetworkAddressHistoryV1TooManyRequests) String() string {
	return fmt.Sprintf("[POST /devices/combined/devices/network-address-history/v1][%d] queryGetNetworkAddressHistoryV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryGetNetworkAddressHistoryV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryGetNetworkAddressHistoryV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
