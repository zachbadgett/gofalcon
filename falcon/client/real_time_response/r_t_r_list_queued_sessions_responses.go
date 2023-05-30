// Code generated by go-swagger; DO NOT EDIT.

package real_time_response

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

// RTRListQueuedSessionsReader is a Reader for the RTRListQueuedSessions structure.
type RTRListQueuedSessionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRListQueuedSessionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRTRListQueuedSessionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRTRListQueuedSessionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRTRListQueuedSessionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRListQueuedSessionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRTRListQueuedSessionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRListQueuedSessionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRTRListQueuedSessionsOK creates a RTRListQueuedSessionsOK with default headers values
func NewRTRListQueuedSessionsOK() *RTRListQueuedSessionsOK {
	return &RTRListQueuedSessionsOK{}
}

/*
RTRListQueuedSessionsOK describes a response with status code 200, with default header values.

success
*/
type RTRListQueuedSessionsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainQueuedSessionResponseWrapper
}

// IsSuccess returns true when this r t r list queued sessions o k response has a 2xx status code
func (o *RTRListQueuedSessionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this r t r list queued sessions o k response has a 3xx status code
func (o *RTRListQueuedSessionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list queued sessions o k response has a 4xx status code
func (o *RTRListQueuedSessionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r list queued sessions o k response has a 5xx status code
func (o *RTRListQueuedSessionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list queued sessions o k response a status code equal to that given
func (o *RTRListQueuedSessionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the r t r list queued sessions o k response
func (o *RTRListQueuedSessionsOK) Code() int {
	return 200
}

func (o *RTRListQueuedSessionsOK) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/queued-sessions/GET/v1][%d] rTRListQueuedSessionsOK  %+v", 200, o.Payload)
}

func (o *RTRListQueuedSessionsOK) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/queued-sessions/GET/v1][%d] rTRListQueuedSessionsOK  %+v", 200, o.Payload)
}

func (o *RTRListQueuedSessionsOK) GetPayload() *models.DomainQueuedSessionResponseWrapper {
	return o.Payload
}

func (o *RTRListQueuedSessionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainQueuedSessionResponseWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListQueuedSessionsBadRequest creates a RTRListQueuedSessionsBadRequest with default headers values
func NewRTRListQueuedSessionsBadRequest() *RTRListQueuedSessionsBadRequest {
	return &RTRListQueuedSessionsBadRequest{}
}

/*
RTRListQueuedSessionsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RTRListQueuedSessionsBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this r t r list queued sessions bad request response has a 2xx status code
func (o *RTRListQueuedSessionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r list queued sessions bad request response has a 3xx status code
func (o *RTRListQueuedSessionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list queued sessions bad request response has a 4xx status code
func (o *RTRListQueuedSessionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r list queued sessions bad request response has a 5xx status code
func (o *RTRListQueuedSessionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list queued sessions bad request response a status code equal to that given
func (o *RTRListQueuedSessionsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the r t r list queued sessions bad request response
func (o *RTRListQueuedSessionsBadRequest) Code() int {
	return 400
}

func (o *RTRListQueuedSessionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/queued-sessions/GET/v1][%d] rTRListQueuedSessionsBadRequest  %+v", 400, o.Payload)
}

func (o *RTRListQueuedSessionsBadRequest) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/queued-sessions/GET/v1][%d] rTRListQueuedSessionsBadRequest  %+v", 400, o.Payload)
}

func (o *RTRListQueuedSessionsBadRequest) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRListQueuedSessionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListQueuedSessionsUnauthorized creates a RTRListQueuedSessionsUnauthorized with default headers values
func NewRTRListQueuedSessionsUnauthorized() *RTRListQueuedSessionsUnauthorized {
	return &RTRListQueuedSessionsUnauthorized{}
}

/*
RTRListQueuedSessionsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RTRListQueuedSessionsUnauthorized struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this r t r list queued sessions unauthorized response has a 2xx status code
func (o *RTRListQueuedSessionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r list queued sessions unauthorized response has a 3xx status code
func (o *RTRListQueuedSessionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list queued sessions unauthorized response has a 4xx status code
func (o *RTRListQueuedSessionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r list queued sessions unauthorized response has a 5xx status code
func (o *RTRListQueuedSessionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list queued sessions unauthorized response a status code equal to that given
func (o *RTRListQueuedSessionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the r t r list queued sessions unauthorized response
func (o *RTRListQueuedSessionsUnauthorized) Code() int {
	return 401
}

func (o *RTRListQueuedSessionsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/queued-sessions/GET/v1][%d] rTRListQueuedSessionsUnauthorized  %+v", 401, o.Payload)
}

func (o *RTRListQueuedSessionsUnauthorized) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/queued-sessions/GET/v1][%d] rTRListQueuedSessionsUnauthorized  %+v", 401, o.Payload)
}

func (o *RTRListQueuedSessionsUnauthorized) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRListQueuedSessionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListQueuedSessionsForbidden creates a RTRListQueuedSessionsForbidden with default headers values
func NewRTRListQueuedSessionsForbidden() *RTRListQueuedSessionsForbidden {
	return &RTRListQueuedSessionsForbidden{}
}

/*
RTRListQueuedSessionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RTRListQueuedSessionsForbidden struct {

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

// IsSuccess returns true when this r t r list queued sessions forbidden response has a 2xx status code
func (o *RTRListQueuedSessionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r list queued sessions forbidden response has a 3xx status code
func (o *RTRListQueuedSessionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list queued sessions forbidden response has a 4xx status code
func (o *RTRListQueuedSessionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r list queued sessions forbidden response has a 5xx status code
func (o *RTRListQueuedSessionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list queued sessions forbidden response a status code equal to that given
func (o *RTRListQueuedSessionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the r t r list queued sessions forbidden response
func (o *RTRListQueuedSessionsForbidden) Code() int {
	return 403
}

func (o *RTRListQueuedSessionsForbidden) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/queued-sessions/GET/v1][%d] rTRListQueuedSessionsForbidden  %+v", 403, o.Payload)
}

func (o *RTRListQueuedSessionsForbidden) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/queued-sessions/GET/v1][%d] rTRListQueuedSessionsForbidden  %+v", 403, o.Payload)
}

func (o *RTRListQueuedSessionsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRListQueuedSessionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRListQueuedSessionsNotFound creates a RTRListQueuedSessionsNotFound with default headers values
func NewRTRListQueuedSessionsNotFound() *RTRListQueuedSessionsNotFound {
	return &RTRListQueuedSessionsNotFound{}
}

/*
RTRListQueuedSessionsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RTRListQueuedSessionsNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this r t r list queued sessions not found response has a 2xx status code
func (o *RTRListQueuedSessionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r list queued sessions not found response has a 3xx status code
func (o *RTRListQueuedSessionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list queued sessions not found response has a 4xx status code
func (o *RTRListQueuedSessionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r list queued sessions not found response has a 5xx status code
func (o *RTRListQueuedSessionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list queued sessions not found response a status code equal to that given
func (o *RTRListQueuedSessionsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the r t r list queued sessions not found response
func (o *RTRListQueuedSessionsNotFound) Code() int {
	return 404
}

func (o *RTRListQueuedSessionsNotFound) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/queued-sessions/GET/v1][%d] rTRListQueuedSessionsNotFound  %+v", 404, o.Payload)
}

func (o *RTRListQueuedSessionsNotFound) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/queued-sessions/GET/v1][%d] rTRListQueuedSessionsNotFound  %+v", 404, o.Payload)
}

func (o *RTRListQueuedSessionsNotFound) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRListQueuedSessionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListQueuedSessionsTooManyRequests creates a RTRListQueuedSessionsTooManyRequests with default headers values
func NewRTRListQueuedSessionsTooManyRequests() *RTRListQueuedSessionsTooManyRequests {
	return &RTRListQueuedSessionsTooManyRequests{}
}

/*
RTRListQueuedSessionsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RTRListQueuedSessionsTooManyRequests struct {

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

// IsSuccess returns true when this r t r list queued sessions too many requests response has a 2xx status code
func (o *RTRListQueuedSessionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r list queued sessions too many requests response has a 3xx status code
func (o *RTRListQueuedSessionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list queued sessions too many requests response has a 4xx status code
func (o *RTRListQueuedSessionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r list queued sessions too many requests response has a 5xx status code
func (o *RTRListQueuedSessionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list queued sessions too many requests response a status code equal to that given
func (o *RTRListQueuedSessionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the r t r list queued sessions too many requests response
func (o *RTRListQueuedSessionsTooManyRequests) Code() int {
	return 429
}

func (o *RTRListQueuedSessionsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/queued-sessions/GET/v1][%d] rTRListQueuedSessionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRListQueuedSessionsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/queued-sessions/GET/v1][%d] rTRListQueuedSessionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRListQueuedSessionsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRListQueuedSessionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
