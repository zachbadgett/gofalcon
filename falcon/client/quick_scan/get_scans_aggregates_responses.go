// Code generated by go-swagger; DO NOT EDIT.

package quick_scan

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

// GetScansAggregatesReader is a Reader for the GetScansAggregates structure.
type GetScansAggregatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScansAggregatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScansAggregatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetScansAggregatesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetScansAggregatesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScansAggregatesOK creates a GetScansAggregatesOK with default headers values
func NewGetScansAggregatesOK() *GetScansAggregatesOK {
	return &GetScansAggregatesOK{}
}

/*
GetScansAggregatesOK describes a response with status code 200, with default header values.

OK
*/
type GetScansAggregatesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64
}

// IsSuccess returns true when this get scans aggregates o k response has a 2xx status code
func (o *GetScansAggregatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get scans aggregates o k response has a 3xx status code
func (o *GetScansAggregatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scans aggregates o k response has a 4xx status code
func (o *GetScansAggregatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scans aggregates o k response has a 5xx status code
func (o *GetScansAggregatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get scans aggregates o k response a status code equal to that given
func (o *GetScansAggregatesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get scans aggregates o k response
func (o *GetScansAggregatesOK) Code() int {
	return 200
}

func (o *GetScansAggregatesOK) Error() string {
	return fmt.Sprintf("[POST /scanner/aggregates/scans/GET/v1][%d] getScansAggregatesOK ", 200)
}

func (o *GetScansAggregatesOK) String() string {
	return fmt.Sprintf("[POST /scanner/aggregates/scans/GET/v1][%d] getScansAggregatesOK ", 200)
}

func (o *GetScansAggregatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	return nil
}

// NewGetScansAggregatesForbidden creates a GetScansAggregatesForbidden with default headers values
func NewGetScansAggregatesForbidden() *GetScansAggregatesForbidden {
	return &GetScansAggregatesForbidden{}
}

/*
GetScansAggregatesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetScansAggregatesForbidden struct {

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

// IsSuccess returns true when this get scans aggregates forbidden response has a 2xx status code
func (o *GetScansAggregatesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scans aggregates forbidden response has a 3xx status code
func (o *GetScansAggregatesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scans aggregates forbidden response has a 4xx status code
func (o *GetScansAggregatesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scans aggregates forbidden response has a 5xx status code
func (o *GetScansAggregatesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get scans aggregates forbidden response a status code equal to that given
func (o *GetScansAggregatesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get scans aggregates forbidden response
func (o *GetScansAggregatesForbidden) Code() int {
	return 403
}

func (o *GetScansAggregatesForbidden) Error() string {
	return fmt.Sprintf("[POST /scanner/aggregates/scans/GET/v1][%d] getScansAggregatesForbidden  %+v", 403, o.Payload)
}

func (o *GetScansAggregatesForbidden) String() string {
	return fmt.Sprintf("[POST /scanner/aggregates/scans/GET/v1][%d] getScansAggregatesForbidden  %+v", 403, o.Payload)
}

func (o *GetScansAggregatesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetScansAggregatesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetScansAggregatesTooManyRequests creates a GetScansAggregatesTooManyRequests with default headers values
func NewGetScansAggregatesTooManyRequests() *GetScansAggregatesTooManyRequests {
	return &GetScansAggregatesTooManyRequests{}
}

/*
GetScansAggregatesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetScansAggregatesTooManyRequests struct {

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

// IsSuccess returns true when this get scans aggregates too many requests response has a 2xx status code
func (o *GetScansAggregatesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scans aggregates too many requests response has a 3xx status code
func (o *GetScansAggregatesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scans aggregates too many requests response has a 4xx status code
func (o *GetScansAggregatesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scans aggregates too many requests response has a 5xx status code
func (o *GetScansAggregatesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get scans aggregates too many requests response a status code equal to that given
func (o *GetScansAggregatesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get scans aggregates too many requests response
func (o *GetScansAggregatesTooManyRequests) Code() int {
	return 429
}

func (o *GetScansAggregatesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /scanner/aggregates/scans/GET/v1][%d] getScansAggregatesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScansAggregatesTooManyRequests) String() string {
	return fmt.Sprintf("[POST /scanner/aggregates/scans/GET/v1][%d] getScansAggregatesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScansAggregatesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetScansAggregatesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
