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

// RTRDeleteSessionReader is a Reader for the RTRDeleteSession structure.
type RTRDeleteSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRDeleteSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRTRDeleteSessionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRTRDeleteSessionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRTRDeleteSessionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRDeleteSessionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRDeleteSessionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRTRDeleteSessionNoContent creates a RTRDeleteSessionNoContent with default headers values
func NewRTRDeleteSessionNoContent() *RTRDeleteSessionNoContent {
	return &RTRDeleteSessionNoContent{}
}

/*
RTRDeleteSessionNoContent describes a response with status code 204, with default header values.

No Content
*/
type RTRDeleteSessionNoContent struct {

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

// IsSuccess returns true when this r t r delete session no content response has a 2xx status code
func (o *RTRDeleteSessionNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this r t r delete session no content response has a 3xx status code
func (o *RTRDeleteSessionNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete session no content response has a 4xx status code
func (o *RTRDeleteSessionNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r delete session no content response has a 5xx status code
func (o *RTRDeleteSessionNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete session no content response a status code equal to that given
func (o *RTRDeleteSessionNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the r t r delete session no content response
func (o *RTRDeleteSessionNoContent) Code() int {
	return 204
}

func (o *RTRDeleteSessionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/sessions/v1][%d] rTRDeleteSessionNoContent  %+v", 204, o.Payload)
}

func (o *RTRDeleteSessionNoContent) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/sessions/v1][%d] rTRDeleteSessionNoContent  %+v", 204, o.Payload)
}

func (o *RTRDeleteSessionNoContent) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeleteSessionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeleteSessionBadRequest creates a RTRDeleteSessionBadRequest with default headers values
func NewRTRDeleteSessionBadRequest() *RTRDeleteSessionBadRequest {
	return &RTRDeleteSessionBadRequest{}
}

/*
RTRDeleteSessionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RTRDeleteSessionBadRequest struct {

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

// IsSuccess returns true when this r t r delete session bad request response has a 2xx status code
func (o *RTRDeleteSessionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r delete session bad request response has a 3xx status code
func (o *RTRDeleteSessionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete session bad request response has a 4xx status code
func (o *RTRDeleteSessionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r delete session bad request response has a 5xx status code
func (o *RTRDeleteSessionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete session bad request response a status code equal to that given
func (o *RTRDeleteSessionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the r t r delete session bad request response
func (o *RTRDeleteSessionBadRequest) Code() int {
	return 400
}

func (o *RTRDeleteSessionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/sessions/v1][%d] rTRDeleteSessionBadRequest  %+v", 400, o.Payload)
}

func (o *RTRDeleteSessionBadRequest) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/sessions/v1][%d] rTRDeleteSessionBadRequest  %+v", 400, o.Payload)
}

func (o *RTRDeleteSessionBadRequest) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRDeleteSessionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeleteSessionUnauthorized creates a RTRDeleteSessionUnauthorized with default headers values
func NewRTRDeleteSessionUnauthorized() *RTRDeleteSessionUnauthorized {
	return &RTRDeleteSessionUnauthorized{}
}

/*
RTRDeleteSessionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RTRDeleteSessionUnauthorized struct {

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

// IsSuccess returns true when this r t r delete session unauthorized response has a 2xx status code
func (o *RTRDeleteSessionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r delete session unauthorized response has a 3xx status code
func (o *RTRDeleteSessionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete session unauthorized response has a 4xx status code
func (o *RTRDeleteSessionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r delete session unauthorized response has a 5xx status code
func (o *RTRDeleteSessionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete session unauthorized response a status code equal to that given
func (o *RTRDeleteSessionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the r t r delete session unauthorized response
func (o *RTRDeleteSessionUnauthorized) Code() int {
	return 401
}

func (o *RTRDeleteSessionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/sessions/v1][%d] rTRDeleteSessionUnauthorized  %+v", 401, o.Payload)
}

func (o *RTRDeleteSessionUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/sessions/v1][%d] rTRDeleteSessionUnauthorized  %+v", 401, o.Payload)
}

func (o *RTRDeleteSessionUnauthorized) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRDeleteSessionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeleteSessionForbidden creates a RTRDeleteSessionForbidden with default headers values
func NewRTRDeleteSessionForbidden() *RTRDeleteSessionForbidden {
	return &RTRDeleteSessionForbidden{}
}

/*
RTRDeleteSessionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RTRDeleteSessionForbidden struct {

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

// IsSuccess returns true when this r t r delete session forbidden response has a 2xx status code
func (o *RTRDeleteSessionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r delete session forbidden response has a 3xx status code
func (o *RTRDeleteSessionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete session forbidden response has a 4xx status code
func (o *RTRDeleteSessionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r delete session forbidden response has a 5xx status code
func (o *RTRDeleteSessionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete session forbidden response a status code equal to that given
func (o *RTRDeleteSessionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the r t r delete session forbidden response
func (o *RTRDeleteSessionForbidden) Code() int {
	return 403
}

func (o *RTRDeleteSessionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/sessions/v1][%d] rTRDeleteSessionForbidden  %+v", 403, o.Payload)
}

func (o *RTRDeleteSessionForbidden) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/sessions/v1][%d] rTRDeleteSessionForbidden  %+v", 403, o.Payload)
}

func (o *RTRDeleteSessionForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeleteSessionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeleteSessionTooManyRequests creates a RTRDeleteSessionTooManyRequests with default headers values
func NewRTRDeleteSessionTooManyRequests() *RTRDeleteSessionTooManyRequests {
	return &RTRDeleteSessionTooManyRequests{}
}

/*
RTRDeleteSessionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RTRDeleteSessionTooManyRequests struct {

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

// IsSuccess returns true when this r t r delete session too many requests response has a 2xx status code
func (o *RTRDeleteSessionTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r delete session too many requests response has a 3xx status code
func (o *RTRDeleteSessionTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete session too many requests response has a 4xx status code
func (o *RTRDeleteSessionTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r delete session too many requests response has a 5xx status code
func (o *RTRDeleteSessionTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete session too many requests response a status code equal to that given
func (o *RTRDeleteSessionTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the r t r delete session too many requests response
func (o *RTRDeleteSessionTooManyRequests) Code() int {
	return 429
}

func (o *RTRDeleteSessionTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/sessions/v1][%d] rTRDeleteSessionTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRDeleteSessionTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/sessions/v1][%d] rTRDeleteSessionTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRDeleteSessionTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeleteSessionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
