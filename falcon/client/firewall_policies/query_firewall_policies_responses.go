// Code generated by go-swagger; DO NOT EDIT.

package firewall_policies

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

// QueryFirewallPoliciesReader is a Reader for the QueryFirewallPolicies structure.
type QueryFirewallPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryFirewallPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryFirewallPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryFirewallPoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryFirewallPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryFirewallPoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryFirewallPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryFirewallPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryFirewallPoliciesOK creates a QueryFirewallPoliciesOK with default headers values
func NewQueryFirewallPoliciesOK() *QueryFirewallPoliciesOK {
	return &QueryFirewallPoliciesOK{}
}

/*
QueryFirewallPoliciesOK describes a response with status code 200, with default header values.

OK
*/
type QueryFirewallPoliciesOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query firewall policies o k response has a 2xx status code
func (o *QueryFirewallPoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query firewall policies o k response has a 3xx status code
func (o *QueryFirewallPoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query firewall policies o k response has a 4xx status code
func (o *QueryFirewallPoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query firewall policies o k response has a 5xx status code
func (o *QueryFirewallPoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query firewall policies o k response a status code equal to that given
func (o *QueryFirewallPoliciesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query firewall policies o k response
func (o *QueryFirewallPoliciesOK) Code() int {
	return 200
}

func (o *QueryFirewallPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /policy/queries/firewall/v1][%d] queryFirewallPoliciesOK  %+v", 200, o.Payload)
}

func (o *QueryFirewallPoliciesOK) String() string {
	return fmt.Sprintf("[GET /policy/queries/firewall/v1][%d] queryFirewallPoliciesOK  %+v", 200, o.Payload)
}

func (o *QueryFirewallPoliciesOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryFirewallPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryFirewallPoliciesBadRequest creates a QueryFirewallPoliciesBadRequest with default headers values
func NewQueryFirewallPoliciesBadRequest() *QueryFirewallPoliciesBadRequest {
	return &QueryFirewallPoliciesBadRequest{}
}

/*
QueryFirewallPoliciesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryFirewallPoliciesBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query firewall policies bad request response has a 2xx status code
func (o *QueryFirewallPoliciesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query firewall policies bad request response has a 3xx status code
func (o *QueryFirewallPoliciesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query firewall policies bad request response has a 4xx status code
func (o *QueryFirewallPoliciesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query firewall policies bad request response has a 5xx status code
func (o *QueryFirewallPoliciesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query firewall policies bad request response a status code equal to that given
func (o *QueryFirewallPoliciesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query firewall policies bad request response
func (o *QueryFirewallPoliciesBadRequest) Code() int {
	return 400
}

func (o *QueryFirewallPoliciesBadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/queries/firewall/v1][%d] queryFirewallPoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *QueryFirewallPoliciesBadRequest) String() string {
	return fmt.Sprintf("[GET /policy/queries/firewall/v1][%d] queryFirewallPoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *QueryFirewallPoliciesBadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryFirewallPoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryFirewallPoliciesForbidden creates a QueryFirewallPoliciesForbidden with default headers values
func NewQueryFirewallPoliciesForbidden() *QueryFirewallPoliciesForbidden {
	return &QueryFirewallPoliciesForbidden{}
}

/*
QueryFirewallPoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryFirewallPoliciesForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this query firewall policies forbidden response has a 2xx status code
func (o *QueryFirewallPoliciesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query firewall policies forbidden response has a 3xx status code
func (o *QueryFirewallPoliciesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query firewall policies forbidden response has a 4xx status code
func (o *QueryFirewallPoliciesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query firewall policies forbidden response has a 5xx status code
func (o *QueryFirewallPoliciesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query firewall policies forbidden response a status code equal to that given
func (o *QueryFirewallPoliciesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query firewall policies forbidden response
func (o *QueryFirewallPoliciesForbidden) Code() int {
	return 403
}

func (o *QueryFirewallPoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/queries/firewall/v1][%d] queryFirewallPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *QueryFirewallPoliciesForbidden) String() string {
	return fmt.Sprintf("[GET /policy/queries/firewall/v1][%d] queryFirewallPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *QueryFirewallPoliciesForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryFirewallPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryFirewallPoliciesTooManyRequests creates a QueryFirewallPoliciesTooManyRequests with default headers values
func NewQueryFirewallPoliciesTooManyRequests() *QueryFirewallPoliciesTooManyRequests {
	return &QueryFirewallPoliciesTooManyRequests{}
}

/*
QueryFirewallPoliciesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryFirewallPoliciesTooManyRequests struct {

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

// IsSuccess returns true when this query firewall policies too many requests response has a 2xx status code
func (o *QueryFirewallPoliciesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query firewall policies too many requests response has a 3xx status code
func (o *QueryFirewallPoliciesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query firewall policies too many requests response has a 4xx status code
func (o *QueryFirewallPoliciesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query firewall policies too many requests response has a 5xx status code
func (o *QueryFirewallPoliciesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query firewall policies too many requests response a status code equal to that given
func (o *QueryFirewallPoliciesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query firewall policies too many requests response
func (o *QueryFirewallPoliciesTooManyRequests) Code() int {
	return 429
}

func (o *QueryFirewallPoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/queries/firewall/v1][%d] queryFirewallPoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryFirewallPoliciesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /policy/queries/firewall/v1][%d] queryFirewallPoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryFirewallPoliciesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryFirewallPoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryFirewallPoliciesInternalServerError creates a QueryFirewallPoliciesInternalServerError with default headers values
func NewQueryFirewallPoliciesInternalServerError() *QueryFirewallPoliciesInternalServerError {
	return &QueryFirewallPoliciesInternalServerError{}
}

/*
QueryFirewallPoliciesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryFirewallPoliciesInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query firewall policies internal server error response has a 2xx status code
func (o *QueryFirewallPoliciesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query firewall policies internal server error response has a 3xx status code
func (o *QueryFirewallPoliciesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query firewall policies internal server error response has a 4xx status code
func (o *QueryFirewallPoliciesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query firewall policies internal server error response has a 5xx status code
func (o *QueryFirewallPoliciesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query firewall policies internal server error response a status code equal to that given
func (o *QueryFirewallPoliciesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query firewall policies internal server error response
func (o *QueryFirewallPoliciesInternalServerError) Code() int {
	return 500
}

func (o *QueryFirewallPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/queries/firewall/v1][%d] queryFirewallPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryFirewallPoliciesInternalServerError) String() string {
	return fmt.Sprintf("[GET /policy/queries/firewall/v1][%d] queryFirewallPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryFirewallPoliciesInternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryFirewallPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryFirewallPoliciesDefault creates a QueryFirewallPoliciesDefault with default headers values
func NewQueryFirewallPoliciesDefault(code int) *QueryFirewallPoliciesDefault {
	return &QueryFirewallPoliciesDefault{
		_statusCode: code,
	}
}

/*
QueryFirewallPoliciesDefault describes a response with status code -1, with default header values.

OK
*/
type QueryFirewallPoliciesDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query firewall policies default response has a 2xx status code
func (o *QueryFirewallPoliciesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this query firewall policies default response has a 3xx status code
func (o *QueryFirewallPoliciesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this query firewall policies default response has a 4xx status code
func (o *QueryFirewallPoliciesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this query firewall policies default response has a 5xx status code
func (o *QueryFirewallPoliciesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this query firewall policies default response a status code equal to that given
func (o *QueryFirewallPoliciesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the query firewall policies default response
func (o *QueryFirewallPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *QueryFirewallPoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /policy/queries/firewall/v1][%d] queryFirewallPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *QueryFirewallPoliciesDefault) String() string {
	return fmt.Sprintf("[GET /policy/queries/firewall/v1][%d] queryFirewallPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *QueryFirewallPoliciesDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryFirewallPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
