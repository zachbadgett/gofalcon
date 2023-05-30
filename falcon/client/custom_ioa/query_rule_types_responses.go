// Code generated by go-swagger; DO NOT EDIT.

package custom_ioa

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

// QueryRuleTypesReader is a Reader for the QueryRuleTypes structure.
type QueryRuleTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryRuleTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryRuleTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQueryRuleTypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryRuleTypesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryRuleTypesOK creates a QueryRuleTypesOK with default headers values
func NewQueryRuleTypesOK() *QueryRuleTypesOK {
	return &QueryRuleTypesOK{}
}

/*
QueryRuleTypesOK describes a response with status code 200, with default header values.

OK
*/
type QueryRuleTypesOK struct {

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

// IsSuccess returns true when this query rule types o k response has a 2xx status code
func (o *QueryRuleTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query rule types o k response has a 3xx status code
func (o *QueryRuleTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query rule types o k response has a 4xx status code
func (o *QueryRuleTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query rule types o k response has a 5xx status code
func (o *QueryRuleTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query rule types o k response a status code equal to that given
func (o *QueryRuleTypesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query rule types o k response
func (o *QueryRuleTypesOK) Code() int {
	return 200
}

func (o *QueryRuleTypesOK) Error() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-types/v1][%d] queryRuleTypesOK  %+v", 200, o.Payload)
}

func (o *QueryRuleTypesOK) String() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-types/v1][%d] queryRuleTypesOK  %+v", 200, o.Payload)
}

func (o *QueryRuleTypesOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryRuleTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryRuleTypesForbidden creates a QueryRuleTypesForbidden with default headers values
func NewQueryRuleTypesForbidden() *QueryRuleTypesForbidden {
	return &QueryRuleTypesForbidden{}
}

/*
QueryRuleTypesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryRuleTypesForbidden struct {

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

// IsSuccess returns true when this query rule types forbidden response has a 2xx status code
func (o *QueryRuleTypesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query rule types forbidden response has a 3xx status code
func (o *QueryRuleTypesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query rule types forbidden response has a 4xx status code
func (o *QueryRuleTypesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query rule types forbidden response has a 5xx status code
func (o *QueryRuleTypesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query rule types forbidden response a status code equal to that given
func (o *QueryRuleTypesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query rule types forbidden response
func (o *QueryRuleTypesForbidden) Code() int {
	return 403
}

func (o *QueryRuleTypesForbidden) Error() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-types/v1][%d] queryRuleTypesForbidden  %+v", 403, o.Payload)
}

func (o *QueryRuleTypesForbidden) String() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-types/v1][%d] queryRuleTypesForbidden  %+v", 403, o.Payload)
}

func (o *QueryRuleTypesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryRuleTypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryRuleTypesTooManyRequests creates a QueryRuleTypesTooManyRequests with default headers values
func NewQueryRuleTypesTooManyRequests() *QueryRuleTypesTooManyRequests {
	return &QueryRuleTypesTooManyRequests{}
}

/*
QueryRuleTypesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryRuleTypesTooManyRequests struct {

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

// IsSuccess returns true when this query rule types too many requests response has a 2xx status code
func (o *QueryRuleTypesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query rule types too many requests response has a 3xx status code
func (o *QueryRuleTypesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query rule types too many requests response has a 4xx status code
func (o *QueryRuleTypesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query rule types too many requests response has a 5xx status code
func (o *QueryRuleTypesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query rule types too many requests response a status code equal to that given
func (o *QueryRuleTypesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query rule types too many requests response
func (o *QueryRuleTypesTooManyRequests) Code() int {
	return 429
}

func (o *QueryRuleTypesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-types/v1][%d] queryRuleTypesTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryRuleTypesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /ioarules/queries/rule-types/v1][%d] queryRuleTypesTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryRuleTypesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryRuleTypesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
