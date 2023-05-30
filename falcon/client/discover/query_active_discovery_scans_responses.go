// Code generated by go-swagger; DO NOT EDIT.

package discover

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

// QueryActiveDiscoveryScansReader is a Reader for the QueryActiveDiscoveryScans structure.
type QueryActiveDiscoveryScansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryActiveDiscoveryScansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryActiveDiscoveryScansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryActiveDiscoveryScansBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryActiveDiscoveryScansForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryActiveDiscoveryScansTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryActiveDiscoveryScansInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryActiveDiscoveryScansOK creates a QueryActiveDiscoveryScansOK with default headers values
func NewQueryActiveDiscoveryScansOK() *QueryActiveDiscoveryScansOK {
	return &QueryActiveDiscoveryScansOK{}
}

/*
QueryActiveDiscoveryScansOK describes a response with status code 200, with default header values.

OK
*/
type QueryActiveDiscoveryScansOK struct {

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

// IsSuccess returns true when this query active discovery scans o k response has a 2xx status code
func (o *QueryActiveDiscoveryScansOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query active discovery scans o k response has a 3xx status code
func (o *QueryActiveDiscoveryScansOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query active discovery scans o k response has a 4xx status code
func (o *QueryActiveDiscoveryScansOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query active discovery scans o k response has a 5xx status code
func (o *QueryActiveDiscoveryScansOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query active discovery scans o k response a status code equal to that given
func (o *QueryActiveDiscoveryScansOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query active discovery scans o k response
func (o *QueryActiveDiscoveryScansOK) Code() int {
	return 200
}

func (o *QueryActiveDiscoveryScansOK) Error() string {
	return fmt.Sprintf("[GET /discover/queries/active-discovery-scans/v1][%d] queryActiveDiscoveryScansOK  %+v", 200, o.Payload)
}

func (o *QueryActiveDiscoveryScansOK) String() string {
	return fmt.Sprintf("[GET /discover/queries/active-discovery-scans/v1][%d] queryActiveDiscoveryScansOK  %+v", 200, o.Payload)
}

func (o *QueryActiveDiscoveryScansOK) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *QueryActiveDiscoveryScansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryActiveDiscoveryScansBadRequest creates a QueryActiveDiscoveryScansBadRequest with default headers values
func NewQueryActiveDiscoveryScansBadRequest() *QueryActiveDiscoveryScansBadRequest {
	return &QueryActiveDiscoveryScansBadRequest{}
}

/*
QueryActiveDiscoveryScansBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryActiveDiscoveryScansBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this query active discovery scans bad request response has a 2xx status code
func (o *QueryActiveDiscoveryScansBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query active discovery scans bad request response has a 3xx status code
func (o *QueryActiveDiscoveryScansBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query active discovery scans bad request response has a 4xx status code
func (o *QueryActiveDiscoveryScansBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query active discovery scans bad request response has a 5xx status code
func (o *QueryActiveDiscoveryScansBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query active discovery scans bad request response a status code equal to that given
func (o *QueryActiveDiscoveryScansBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query active discovery scans bad request response
func (o *QueryActiveDiscoveryScansBadRequest) Code() int {
	return 400
}

func (o *QueryActiveDiscoveryScansBadRequest) Error() string {
	return fmt.Sprintf("[GET /discover/queries/active-discovery-scans/v1][%d] queryActiveDiscoveryScansBadRequest  %+v", 400, o.Payload)
}

func (o *QueryActiveDiscoveryScansBadRequest) String() string {
	return fmt.Sprintf("[GET /discover/queries/active-discovery-scans/v1][%d] queryActiveDiscoveryScansBadRequest  %+v", 400, o.Payload)
}

func (o *QueryActiveDiscoveryScansBadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *QueryActiveDiscoveryScansBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryActiveDiscoveryScansForbidden creates a QueryActiveDiscoveryScansForbidden with default headers values
func NewQueryActiveDiscoveryScansForbidden() *QueryActiveDiscoveryScansForbidden {
	return &QueryActiveDiscoveryScansForbidden{}
}

/*
QueryActiveDiscoveryScansForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryActiveDiscoveryScansForbidden struct {

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

// IsSuccess returns true when this query active discovery scans forbidden response has a 2xx status code
func (o *QueryActiveDiscoveryScansForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query active discovery scans forbidden response has a 3xx status code
func (o *QueryActiveDiscoveryScansForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query active discovery scans forbidden response has a 4xx status code
func (o *QueryActiveDiscoveryScansForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query active discovery scans forbidden response has a 5xx status code
func (o *QueryActiveDiscoveryScansForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query active discovery scans forbidden response a status code equal to that given
func (o *QueryActiveDiscoveryScansForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query active discovery scans forbidden response
func (o *QueryActiveDiscoveryScansForbidden) Code() int {
	return 403
}

func (o *QueryActiveDiscoveryScansForbidden) Error() string {
	return fmt.Sprintf("[GET /discover/queries/active-discovery-scans/v1][%d] queryActiveDiscoveryScansForbidden  %+v", 403, o.Payload)
}

func (o *QueryActiveDiscoveryScansForbidden) String() string {
	return fmt.Sprintf("[GET /discover/queries/active-discovery-scans/v1][%d] queryActiveDiscoveryScansForbidden  %+v", 403, o.Payload)
}

func (o *QueryActiveDiscoveryScansForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryActiveDiscoveryScansForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryActiveDiscoveryScansTooManyRequests creates a QueryActiveDiscoveryScansTooManyRequests with default headers values
func NewQueryActiveDiscoveryScansTooManyRequests() *QueryActiveDiscoveryScansTooManyRequests {
	return &QueryActiveDiscoveryScansTooManyRequests{}
}

/*
QueryActiveDiscoveryScansTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryActiveDiscoveryScansTooManyRequests struct {

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

// IsSuccess returns true when this query active discovery scans too many requests response has a 2xx status code
func (o *QueryActiveDiscoveryScansTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query active discovery scans too many requests response has a 3xx status code
func (o *QueryActiveDiscoveryScansTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query active discovery scans too many requests response has a 4xx status code
func (o *QueryActiveDiscoveryScansTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query active discovery scans too many requests response has a 5xx status code
func (o *QueryActiveDiscoveryScansTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query active discovery scans too many requests response a status code equal to that given
func (o *QueryActiveDiscoveryScansTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query active discovery scans too many requests response
func (o *QueryActiveDiscoveryScansTooManyRequests) Code() int {
	return 429
}

func (o *QueryActiveDiscoveryScansTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /discover/queries/active-discovery-scans/v1][%d] queryActiveDiscoveryScansTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryActiveDiscoveryScansTooManyRequests) String() string {
	return fmt.Sprintf("[GET /discover/queries/active-discovery-scans/v1][%d] queryActiveDiscoveryScansTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryActiveDiscoveryScansTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryActiveDiscoveryScansTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryActiveDiscoveryScansInternalServerError creates a QueryActiveDiscoveryScansInternalServerError with default headers values
func NewQueryActiveDiscoveryScansInternalServerError() *QueryActiveDiscoveryScansInternalServerError {
	return &QueryActiveDiscoveryScansInternalServerError{}
}

/*
QueryActiveDiscoveryScansInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryActiveDiscoveryScansInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this query active discovery scans internal server error response has a 2xx status code
func (o *QueryActiveDiscoveryScansInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query active discovery scans internal server error response has a 3xx status code
func (o *QueryActiveDiscoveryScansInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query active discovery scans internal server error response has a 4xx status code
func (o *QueryActiveDiscoveryScansInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query active discovery scans internal server error response has a 5xx status code
func (o *QueryActiveDiscoveryScansInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query active discovery scans internal server error response a status code equal to that given
func (o *QueryActiveDiscoveryScansInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query active discovery scans internal server error response
func (o *QueryActiveDiscoveryScansInternalServerError) Code() int {
	return 500
}

func (o *QueryActiveDiscoveryScansInternalServerError) Error() string {
	return fmt.Sprintf("[GET /discover/queries/active-discovery-scans/v1][%d] queryActiveDiscoveryScansInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryActiveDiscoveryScansInternalServerError) String() string {
	return fmt.Sprintf("[GET /discover/queries/active-discovery-scans/v1][%d] queryActiveDiscoveryScansInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryActiveDiscoveryScansInternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *QueryActiveDiscoveryScansInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
