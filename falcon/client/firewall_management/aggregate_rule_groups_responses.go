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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// AggregateRuleGroupsReader is a Reader for the AggregateRuleGroups structure.
type AggregateRuleGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregateRuleGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregateRuleGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAggregateRuleGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAggregateRuleGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAggregateRuleGroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAggregateRuleGroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAggregateRuleGroupsOK creates a AggregateRuleGroupsOK with default headers values
func NewAggregateRuleGroupsOK() *AggregateRuleGroupsOK {
	return &AggregateRuleGroupsOK{}
}

/*
AggregateRuleGroupsOK describes a response with status code 200, with default header values.

OK
*/
type AggregateRuleGroupsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FwmgrAPIAggregatesResponse
}

// IsSuccess returns true when this aggregate rule groups o k response has a 2xx status code
func (o *AggregateRuleGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aggregate rule groups o k response has a 3xx status code
func (o *AggregateRuleGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate rule groups o k response has a 4xx status code
func (o *AggregateRuleGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate rule groups o k response has a 5xx status code
func (o *AggregateRuleGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate rule groups o k response a status code equal to that given
func (o *AggregateRuleGroupsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the aggregate rule groups o k response
func (o *AggregateRuleGroupsOK) Code() int {
	return 200
}

func (o *AggregateRuleGroupsOK) Error() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/rule-groups/GET/v1][%d] aggregateRuleGroupsOK  %+v", 200, o.Payload)
}

func (o *AggregateRuleGroupsOK) String() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/rule-groups/GET/v1][%d] aggregateRuleGroupsOK  %+v", 200, o.Payload)
}

func (o *AggregateRuleGroupsOK) GetPayload() *models.FwmgrAPIAggregatesResponse {
	return o.Payload
}

func (o *AggregateRuleGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FwmgrAPIAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateRuleGroupsBadRequest creates a AggregateRuleGroupsBadRequest with default headers values
func NewAggregateRuleGroupsBadRequest() *AggregateRuleGroupsBadRequest {
	return &AggregateRuleGroupsBadRequest{}
}

/*
AggregateRuleGroupsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AggregateRuleGroupsBadRequest struct {

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

// IsSuccess returns true when this aggregate rule groups bad request response has a 2xx status code
func (o *AggregateRuleGroupsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate rule groups bad request response has a 3xx status code
func (o *AggregateRuleGroupsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate rule groups bad request response has a 4xx status code
func (o *AggregateRuleGroupsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate rule groups bad request response has a 5xx status code
func (o *AggregateRuleGroupsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate rule groups bad request response a status code equal to that given
func (o *AggregateRuleGroupsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the aggregate rule groups bad request response
func (o *AggregateRuleGroupsBadRequest) Code() int {
	return 400
}

func (o *AggregateRuleGroupsBadRequest) Error() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/rule-groups/GET/v1][%d] aggregateRuleGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *AggregateRuleGroupsBadRequest) String() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/rule-groups/GET/v1][%d] aggregateRuleGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *AggregateRuleGroupsBadRequest) GetPayload() *models.FwmgrMsaspecResponseFields {
	return o.Payload
}

func (o *AggregateRuleGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregateRuleGroupsForbidden creates a AggregateRuleGroupsForbidden with default headers values
func NewAggregateRuleGroupsForbidden() *AggregateRuleGroupsForbidden {
	return &AggregateRuleGroupsForbidden{}
}

/*
AggregateRuleGroupsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AggregateRuleGroupsForbidden struct {

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

// IsSuccess returns true when this aggregate rule groups forbidden response has a 2xx status code
func (o *AggregateRuleGroupsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate rule groups forbidden response has a 3xx status code
func (o *AggregateRuleGroupsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate rule groups forbidden response has a 4xx status code
func (o *AggregateRuleGroupsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate rule groups forbidden response has a 5xx status code
func (o *AggregateRuleGroupsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate rule groups forbidden response a status code equal to that given
func (o *AggregateRuleGroupsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the aggregate rule groups forbidden response
func (o *AggregateRuleGroupsForbidden) Code() int {
	return 403
}

func (o *AggregateRuleGroupsForbidden) Error() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/rule-groups/GET/v1][%d] aggregateRuleGroupsForbidden  %+v", 403, o.Payload)
}

func (o *AggregateRuleGroupsForbidden) String() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/rule-groups/GET/v1][%d] aggregateRuleGroupsForbidden  %+v", 403, o.Payload)
}

func (o *AggregateRuleGroupsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregateRuleGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregateRuleGroupsTooManyRequests creates a AggregateRuleGroupsTooManyRequests with default headers values
func NewAggregateRuleGroupsTooManyRequests() *AggregateRuleGroupsTooManyRequests {
	return &AggregateRuleGroupsTooManyRequests{}
}

/*
AggregateRuleGroupsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AggregateRuleGroupsTooManyRequests struct {

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

// IsSuccess returns true when this aggregate rule groups too many requests response has a 2xx status code
func (o *AggregateRuleGroupsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate rule groups too many requests response has a 3xx status code
func (o *AggregateRuleGroupsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate rule groups too many requests response has a 4xx status code
func (o *AggregateRuleGroupsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate rule groups too many requests response has a 5xx status code
func (o *AggregateRuleGroupsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate rule groups too many requests response a status code equal to that given
func (o *AggregateRuleGroupsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the aggregate rule groups too many requests response
func (o *AggregateRuleGroupsTooManyRequests) Code() int {
	return 429
}

func (o *AggregateRuleGroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/rule-groups/GET/v1][%d] aggregateRuleGroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregateRuleGroupsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/rule-groups/GET/v1][%d] aggregateRuleGroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregateRuleGroupsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregateRuleGroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregateRuleGroupsInternalServerError creates a AggregateRuleGroupsInternalServerError with default headers values
func NewAggregateRuleGroupsInternalServerError() *AggregateRuleGroupsInternalServerError {
	return &AggregateRuleGroupsInternalServerError{}
}

/*
AggregateRuleGroupsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type AggregateRuleGroupsInternalServerError struct {

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

// IsSuccess returns true when this aggregate rule groups internal server error response has a 2xx status code
func (o *AggregateRuleGroupsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate rule groups internal server error response has a 3xx status code
func (o *AggregateRuleGroupsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate rule groups internal server error response has a 4xx status code
func (o *AggregateRuleGroupsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate rule groups internal server error response has a 5xx status code
func (o *AggregateRuleGroupsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this aggregate rule groups internal server error response a status code equal to that given
func (o *AggregateRuleGroupsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the aggregate rule groups internal server error response
func (o *AggregateRuleGroupsInternalServerError) Code() int {
	return 500
}

func (o *AggregateRuleGroupsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/rule-groups/GET/v1][%d] aggregateRuleGroupsInternalServerError ", 500)
}

func (o *AggregateRuleGroupsInternalServerError) String() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/rule-groups/GET/v1][%d] aggregateRuleGroupsInternalServerError ", 500)
}

func (o *AggregateRuleGroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
