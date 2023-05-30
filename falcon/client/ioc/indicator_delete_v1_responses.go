// Code generated by go-swagger; DO NOT EDIT.

package ioc

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

// IndicatorDeleteV1Reader is a Reader for the IndicatorDeleteV1 structure.
type IndicatorDeleteV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndicatorDeleteV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIndicatorDeleteV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewIndicatorDeleteV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewIndicatorDeleteV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIndicatorDeleteV1OK creates a IndicatorDeleteV1OK with default headers values
func NewIndicatorDeleteV1OK() *IndicatorDeleteV1OK {
	return &IndicatorDeleteV1OK{}
}

/*
IndicatorDeleteV1OK describes a response with status code 200, with default header values.

OK
*/
type IndicatorDeleteV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIIndicatorQueryRespV1
}

// IsSuccess returns true when this indicator delete v1 o k response has a 2xx status code
func (o *IndicatorDeleteV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this indicator delete v1 o k response has a 3xx status code
func (o *IndicatorDeleteV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this indicator delete v1 o k response has a 4xx status code
func (o *IndicatorDeleteV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this indicator delete v1 o k response has a 5xx status code
func (o *IndicatorDeleteV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this indicator delete v1 o k response a status code equal to that given
func (o *IndicatorDeleteV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the indicator delete v1 o k response
func (o *IndicatorDeleteV1OK) Code() int {
	return 200
}

func (o *IndicatorDeleteV1OK) Error() string {
	return fmt.Sprintf("[DELETE /iocs/entities/indicators/v1][%d] indicatorDeleteV1OK  %+v", 200, o.Payload)
}

func (o *IndicatorDeleteV1OK) String() string {
	return fmt.Sprintf("[DELETE /iocs/entities/indicators/v1][%d] indicatorDeleteV1OK  %+v", 200, o.Payload)
}

func (o *IndicatorDeleteV1OK) GetPayload() *models.APIIndicatorQueryRespV1 {
	return o.Payload
}

func (o *IndicatorDeleteV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIIndicatorQueryRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIndicatorDeleteV1Forbidden creates a IndicatorDeleteV1Forbidden with default headers values
func NewIndicatorDeleteV1Forbidden() *IndicatorDeleteV1Forbidden {
	return &IndicatorDeleteV1Forbidden{}
}

/*
IndicatorDeleteV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type IndicatorDeleteV1Forbidden struct {

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

// IsSuccess returns true when this indicator delete v1 forbidden response has a 2xx status code
func (o *IndicatorDeleteV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this indicator delete v1 forbidden response has a 3xx status code
func (o *IndicatorDeleteV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this indicator delete v1 forbidden response has a 4xx status code
func (o *IndicatorDeleteV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this indicator delete v1 forbidden response has a 5xx status code
func (o *IndicatorDeleteV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this indicator delete v1 forbidden response a status code equal to that given
func (o *IndicatorDeleteV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the indicator delete v1 forbidden response
func (o *IndicatorDeleteV1Forbidden) Code() int {
	return 403
}

func (o *IndicatorDeleteV1Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /iocs/entities/indicators/v1][%d] indicatorDeleteV1Forbidden  %+v", 403, o.Payload)
}

func (o *IndicatorDeleteV1Forbidden) String() string {
	return fmt.Sprintf("[DELETE /iocs/entities/indicators/v1][%d] indicatorDeleteV1Forbidden  %+v", 403, o.Payload)
}

func (o *IndicatorDeleteV1Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *IndicatorDeleteV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewIndicatorDeleteV1TooManyRequests creates a IndicatorDeleteV1TooManyRequests with default headers values
func NewIndicatorDeleteV1TooManyRequests() *IndicatorDeleteV1TooManyRequests {
	return &IndicatorDeleteV1TooManyRequests{}
}

/*
IndicatorDeleteV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type IndicatorDeleteV1TooManyRequests struct {

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

// IsSuccess returns true when this indicator delete v1 too many requests response has a 2xx status code
func (o *IndicatorDeleteV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this indicator delete v1 too many requests response has a 3xx status code
func (o *IndicatorDeleteV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this indicator delete v1 too many requests response has a 4xx status code
func (o *IndicatorDeleteV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this indicator delete v1 too many requests response has a 5xx status code
func (o *IndicatorDeleteV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this indicator delete v1 too many requests response a status code equal to that given
func (o *IndicatorDeleteV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the indicator delete v1 too many requests response
func (o *IndicatorDeleteV1TooManyRequests) Code() int {
	return 429
}

func (o *IndicatorDeleteV1TooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /iocs/entities/indicators/v1][%d] indicatorDeleteV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *IndicatorDeleteV1TooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /iocs/entities/indicators/v1][%d] indicatorDeleteV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *IndicatorDeleteV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *IndicatorDeleteV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
