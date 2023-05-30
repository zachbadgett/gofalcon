// Code generated by go-swagger; DO NOT EDIT.

package device_control_policies

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

// QueryDeviceControlPoliciesReader is a Reader for the QueryDeviceControlPolicies structure.
type QueryDeviceControlPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryDeviceControlPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryDeviceControlPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryDeviceControlPoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryDeviceControlPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryDeviceControlPoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryDeviceControlPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryDeviceControlPoliciesOK creates a QueryDeviceControlPoliciesOK with default headers values
func NewQueryDeviceControlPoliciesOK() *QueryDeviceControlPoliciesOK {
	return &QueryDeviceControlPoliciesOK{}
}

/*
QueryDeviceControlPoliciesOK describes a response with status code 200, with default header values.

OK
*/
type QueryDeviceControlPoliciesOK struct {

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

// IsSuccess returns true when this query device control policies o k response has a 2xx status code
func (o *QueryDeviceControlPoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query device control policies o k response has a 3xx status code
func (o *QueryDeviceControlPoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query device control policies o k response has a 4xx status code
func (o *QueryDeviceControlPoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query device control policies o k response has a 5xx status code
func (o *QueryDeviceControlPoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query device control policies o k response a status code equal to that given
func (o *QueryDeviceControlPoliciesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query device control policies o k response
func (o *QueryDeviceControlPoliciesOK) Code() int {
	return 200
}

func (o *QueryDeviceControlPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /policy/queries/device-control/v1][%d] queryDeviceControlPoliciesOK  %+v", 200, o.Payload)
}

func (o *QueryDeviceControlPoliciesOK) String() string {
	return fmt.Sprintf("[GET /policy/queries/device-control/v1][%d] queryDeviceControlPoliciesOK  %+v", 200, o.Payload)
}

func (o *QueryDeviceControlPoliciesOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryDeviceControlPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryDeviceControlPoliciesBadRequest creates a QueryDeviceControlPoliciesBadRequest with default headers values
func NewQueryDeviceControlPoliciesBadRequest() *QueryDeviceControlPoliciesBadRequest {
	return &QueryDeviceControlPoliciesBadRequest{}
}

/*
QueryDeviceControlPoliciesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryDeviceControlPoliciesBadRequest struct {

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

// IsSuccess returns true when this query device control policies bad request response has a 2xx status code
func (o *QueryDeviceControlPoliciesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query device control policies bad request response has a 3xx status code
func (o *QueryDeviceControlPoliciesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query device control policies bad request response has a 4xx status code
func (o *QueryDeviceControlPoliciesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query device control policies bad request response has a 5xx status code
func (o *QueryDeviceControlPoliciesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query device control policies bad request response a status code equal to that given
func (o *QueryDeviceControlPoliciesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query device control policies bad request response
func (o *QueryDeviceControlPoliciesBadRequest) Code() int {
	return 400
}

func (o *QueryDeviceControlPoliciesBadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/queries/device-control/v1][%d] queryDeviceControlPoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *QueryDeviceControlPoliciesBadRequest) String() string {
	return fmt.Sprintf("[GET /policy/queries/device-control/v1][%d] queryDeviceControlPoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *QueryDeviceControlPoliciesBadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryDeviceControlPoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryDeviceControlPoliciesForbidden creates a QueryDeviceControlPoliciesForbidden with default headers values
func NewQueryDeviceControlPoliciesForbidden() *QueryDeviceControlPoliciesForbidden {
	return &QueryDeviceControlPoliciesForbidden{}
}

/*
QueryDeviceControlPoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryDeviceControlPoliciesForbidden struct {

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

// IsSuccess returns true when this query device control policies forbidden response has a 2xx status code
func (o *QueryDeviceControlPoliciesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query device control policies forbidden response has a 3xx status code
func (o *QueryDeviceControlPoliciesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query device control policies forbidden response has a 4xx status code
func (o *QueryDeviceControlPoliciesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query device control policies forbidden response has a 5xx status code
func (o *QueryDeviceControlPoliciesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query device control policies forbidden response a status code equal to that given
func (o *QueryDeviceControlPoliciesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query device control policies forbidden response
func (o *QueryDeviceControlPoliciesForbidden) Code() int {
	return 403
}

func (o *QueryDeviceControlPoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/queries/device-control/v1][%d] queryDeviceControlPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *QueryDeviceControlPoliciesForbidden) String() string {
	return fmt.Sprintf("[GET /policy/queries/device-control/v1][%d] queryDeviceControlPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *QueryDeviceControlPoliciesForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryDeviceControlPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryDeviceControlPoliciesTooManyRequests creates a QueryDeviceControlPoliciesTooManyRequests with default headers values
func NewQueryDeviceControlPoliciesTooManyRequests() *QueryDeviceControlPoliciesTooManyRequests {
	return &QueryDeviceControlPoliciesTooManyRequests{}
}

/*
QueryDeviceControlPoliciesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryDeviceControlPoliciesTooManyRequests struct {

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

// IsSuccess returns true when this query device control policies too many requests response has a 2xx status code
func (o *QueryDeviceControlPoliciesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query device control policies too many requests response has a 3xx status code
func (o *QueryDeviceControlPoliciesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query device control policies too many requests response has a 4xx status code
func (o *QueryDeviceControlPoliciesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query device control policies too many requests response has a 5xx status code
func (o *QueryDeviceControlPoliciesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query device control policies too many requests response a status code equal to that given
func (o *QueryDeviceControlPoliciesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query device control policies too many requests response
func (o *QueryDeviceControlPoliciesTooManyRequests) Code() int {
	return 429
}

func (o *QueryDeviceControlPoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/queries/device-control/v1][%d] queryDeviceControlPoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryDeviceControlPoliciesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /policy/queries/device-control/v1][%d] queryDeviceControlPoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryDeviceControlPoliciesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryDeviceControlPoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryDeviceControlPoliciesInternalServerError creates a QueryDeviceControlPoliciesInternalServerError with default headers values
func NewQueryDeviceControlPoliciesInternalServerError() *QueryDeviceControlPoliciesInternalServerError {
	return &QueryDeviceControlPoliciesInternalServerError{}
}

/*
QueryDeviceControlPoliciesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryDeviceControlPoliciesInternalServerError struct {

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

// IsSuccess returns true when this query device control policies internal server error response has a 2xx status code
func (o *QueryDeviceControlPoliciesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query device control policies internal server error response has a 3xx status code
func (o *QueryDeviceControlPoliciesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query device control policies internal server error response has a 4xx status code
func (o *QueryDeviceControlPoliciesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query device control policies internal server error response has a 5xx status code
func (o *QueryDeviceControlPoliciesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query device control policies internal server error response a status code equal to that given
func (o *QueryDeviceControlPoliciesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query device control policies internal server error response
func (o *QueryDeviceControlPoliciesInternalServerError) Code() int {
	return 500
}

func (o *QueryDeviceControlPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/queries/device-control/v1][%d] queryDeviceControlPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryDeviceControlPoliciesInternalServerError) String() string {
	return fmt.Sprintf("[GET /policy/queries/device-control/v1][%d] queryDeviceControlPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryDeviceControlPoliciesInternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryDeviceControlPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
