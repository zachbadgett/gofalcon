// Code generated by go-swagger; DO NOT EDIT.

package mssp

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

// QueryCIDGroupsReader is a Reader for the QueryCIDGroups structure.
type QueryCIDGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryCIDGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryCIDGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQueryCIDGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryCIDGroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryCIDGroupsOK creates a QueryCIDGroupsOK with default headers values
func NewQueryCIDGroupsOK() *QueryCIDGroupsOK {
	return &QueryCIDGroupsOK{}
}

/*
QueryCIDGroupsOK describes a response with status code 200, with default header values.

OK
*/
type QueryCIDGroupsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query c Id groups o k response has a 2xx status code
func (o *QueryCIDGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query c Id groups o k response has a 3xx status code
func (o *QueryCIDGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query c Id groups o k response has a 4xx status code
func (o *QueryCIDGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query c Id groups o k response has a 5xx status code
func (o *QueryCIDGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query c Id groups o k response a status code equal to that given
func (o *QueryCIDGroupsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query c Id groups o k response
func (o *QueryCIDGroupsOK) Code() int {
	return 200
}

func (o *QueryCIDGroupsOK) Error() string {
	return fmt.Sprintf("[GET /mssp/queries/cid-groups/v1][%d] queryCIdGroupsOK  %+v", 200, o.Payload)
}

func (o *QueryCIDGroupsOK) String() string {
	return fmt.Sprintf("[GET /mssp/queries/cid-groups/v1][%d] queryCIdGroupsOK  %+v", 200, o.Payload)
}

func (o *QueryCIDGroupsOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryCIDGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCIDGroupsForbidden creates a QueryCIDGroupsForbidden with default headers values
func NewQueryCIDGroupsForbidden() *QueryCIDGroupsForbidden {
	return &QueryCIDGroupsForbidden{}
}

/*
QueryCIDGroupsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryCIDGroupsForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this query c Id groups forbidden response has a 2xx status code
func (o *QueryCIDGroupsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query c Id groups forbidden response has a 3xx status code
func (o *QueryCIDGroupsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query c Id groups forbidden response has a 4xx status code
func (o *QueryCIDGroupsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query c Id groups forbidden response has a 5xx status code
func (o *QueryCIDGroupsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query c Id groups forbidden response a status code equal to that given
func (o *QueryCIDGroupsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query c Id groups forbidden response
func (o *QueryCIDGroupsForbidden) Code() int {
	return 403
}

func (o *QueryCIDGroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /mssp/queries/cid-groups/v1][%d] queryCIdGroupsForbidden  %+v", 403, o.Payload)
}

func (o *QueryCIDGroupsForbidden) String() string {
	return fmt.Sprintf("[GET /mssp/queries/cid-groups/v1][%d] queryCIdGroupsForbidden  %+v", 403, o.Payload)
}

func (o *QueryCIDGroupsForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryCIDGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCIDGroupsTooManyRequests creates a QueryCIDGroupsTooManyRequests with default headers values
func NewQueryCIDGroupsTooManyRequests() *QueryCIDGroupsTooManyRequests {
	return &QueryCIDGroupsTooManyRequests{}
}

/*
QueryCIDGroupsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryCIDGroupsTooManyRequests struct {

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

// IsSuccess returns true when this query c Id groups too many requests response has a 2xx status code
func (o *QueryCIDGroupsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query c Id groups too many requests response has a 3xx status code
func (o *QueryCIDGroupsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query c Id groups too many requests response has a 4xx status code
func (o *QueryCIDGroupsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query c Id groups too many requests response has a 5xx status code
func (o *QueryCIDGroupsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query c Id groups too many requests response a status code equal to that given
func (o *QueryCIDGroupsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query c Id groups too many requests response
func (o *QueryCIDGroupsTooManyRequests) Code() int {
	return 429
}

func (o *QueryCIDGroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mssp/queries/cid-groups/v1][%d] queryCIdGroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryCIDGroupsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /mssp/queries/cid-groups/v1][%d] queryCIdGroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryCIDGroupsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryCIDGroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
