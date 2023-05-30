// Code generated by go-swagger; DO NOT EDIT.

package falconx_sandbox

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

// QuerySubmissionsReader is a Reader for the QuerySubmissions structure.
type QuerySubmissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuerySubmissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQuerySubmissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQuerySubmissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQuerySubmissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQuerySubmissionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQuerySubmissionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQuerySubmissionsOK creates a QuerySubmissionsOK with default headers values
func NewQuerySubmissionsOK() *QuerySubmissionsOK {
	return &QuerySubmissionsOK{}
}

/*
QuerySubmissionsOK describes a response with status code 200, with default header values.

OK
*/
type QuerySubmissionsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecQueryResponse
}

// IsSuccess returns true when this query submissions o k response has a 2xx status code
func (o *QuerySubmissionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query submissions o k response has a 3xx status code
func (o *QuerySubmissionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query submissions o k response has a 4xx status code
func (o *QuerySubmissionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query submissions o k response has a 5xx status code
func (o *QuerySubmissionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query submissions o k response a status code equal to that given
func (o *QuerySubmissionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query submissions o k response
func (o *QuerySubmissionsOK) Code() int {
	return 200
}

func (o *QuerySubmissionsOK) Error() string {
	return fmt.Sprintf("[GET /falconx/queries/submissions/v1][%d] querySubmissionsOK  %+v", 200, o.Payload)
}

func (o *QuerySubmissionsOK) String() string {
	return fmt.Sprintf("[GET /falconx/queries/submissions/v1][%d] querySubmissionsOK  %+v", 200, o.Payload)
}

func (o *QuerySubmissionsOK) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *QuerySubmissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuerySubmissionsBadRequest creates a QuerySubmissionsBadRequest with default headers values
func NewQuerySubmissionsBadRequest() *QuerySubmissionsBadRequest {
	return &QuerySubmissionsBadRequest{}
}

/*
QuerySubmissionsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QuerySubmissionsBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecQueryResponse
}

// IsSuccess returns true when this query submissions bad request response has a 2xx status code
func (o *QuerySubmissionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query submissions bad request response has a 3xx status code
func (o *QuerySubmissionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query submissions bad request response has a 4xx status code
func (o *QuerySubmissionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query submissions bad request response has a 5xx status code
func (o *QuerySubmissionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query submissions bad request response a status code equal to that given
func (o *QuerySubmissionsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query submissions bad request response
func (o *QuerySubmissionsBadRequest) Code() int {
	return 400
}

func (o *QuerySubmissionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /falconx/queries/submissions/v1][%d] querySubmissionsBadRequest  %+v", 400, o.Payload)
}

func (o *QuerySubmissionsBadRequest) String() string {
	return fmt.Sprintf("[GET /falconx/queries/submissions/v1][%d] querySubmissionsBadRequest  %+v", 400, o.Payload)
}

func (o *QuerySubmissionsBadRequest) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *QuerySubmissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuerySubmissionsForbidden creates a QuerySubmissionsForbidden with default headers values
func NewQuerySubmissionsForbidden() *QuerySubmissionsForbidden {
	return &QuerySubmissionsForbidden{}
}

/*
QuerySubmissionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QuerySubmissionsForbidden struct {

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

// IsSuccess returns true when this query submissions forbidden response has a 2xx status code
func (o *QuerySubmissionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query submissions forbidden response has a 3xx status code
func (o *QuerySubmissionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query submissions forbidden response has a 4xx status code
func (o *QuerySubmissionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query submissions forbidden response has a 5xx status code
func (o *QuerySubmissionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query submissions forbidden response a status code equal to that given
func (o *QuerySubmissionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query submissions forbidden response
func (o *QuerySubmissionsForbidden) Code() int {
	return 403
}

func (o *QuerySubmissionsForbidden) Error() string {
	return fmt.Sprintf("[GET /falconx/queries/submissions/v1][%d] querySubmissionsForbidden  %+v", 403, o.Payload)
}

func (o *QuerySubmissionsForbidden) String() string {
	return fmt.Sprintf("[GET /falconx/queries/submissions/v1][%d] querySubmissionsForbidden  %+v", 403, o.Payload)
}

func (o *QuerySubmissionsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QuerySubmissionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQuerySubmissionsTooManyRequests creates a QuerySubmissionsTooManyRequests with default headers values
func NewQuerySubmissionsTooManyRequests() *QuerySubmissionsTooManyRequests {
	return &QuerySubmissionsTooManyRequests{}
}

/*
QuerySubmissionsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QuerySubmissionsTooManyRequests struct {

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

// IsSuccess returns true when this query submissions too many requests response has a 2xx status code
func (o *QuerySubmissionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query submissions too many requests response has a 3xx status code
func (o *QuerySubmissionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query submissions too many requests response has a 4xx status code
func (o *QuerySubmissionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query submissions too many requests response has a 5xx status code
func (o *QuerySubmissionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query submissions too many requests response a status code equal to that given
func (o *QuerySubmissionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query submissions too many requests response
func (o *QuerySubmissionsTooManyRequests) Code() int {
	return 429
}

func (o *QuerySubmissionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /falconx/queries/submissions/v1][%d] querySubmissionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *QuerySubmissionsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /falconx/queries/submissions/v1][%d] querySubmissionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *QuerySubmissionsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QuerySubmissionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQuerySubmissionsInternalServerError creates a QuerySubmissionsInternalServerError with default headers values
func NewQuerySubmissionsInternalServerError() *QuerySubmissionsInternalServerError {
	return &QuerySubmissionsInternalServerError{}
}

/*
QuerySubmissionsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QuerySubmissionsInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecQueryResponse
}

// IsSuccess returns true when this query submissions internal server error response has a 2xx status code
func (o *QuerySubmissionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query submissions internal server error response has a 3xx status code
func (o *QuerySubmissionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query submissions internal server error response has a 4xx status code
func (o *QuerySubmissionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query submissions internal server error response has a 5xx status code
func (o *QuerySubmissionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query submissions internal server error response a status code equal to that given
func (o *QuerySubmissionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query submissions internal server error response
func (o *QuerySubmissionsInternalServerError) Code() int {
	return 500
}

func (o *QuerySubmissionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /falconx/queries/submissions/v1][%d] querySubmissionsInternalServerError  %+v", 500, o.Payload)
}

func (o *QuerySubmissionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /falconx/queries/submissions/v1][%d] querySubmissionsInternalServerError  %+v", 500, o.Payload)
}

func (o *QuerySubmissionsInternalServerError) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *QuerySubmissionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
