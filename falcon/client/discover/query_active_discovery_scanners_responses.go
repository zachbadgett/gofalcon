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

// QueryActiveDiscoveryScannersReader is a Reader for the QueryActiveDiscoveryScanners structure.
type QueryActiveDiscoveryScannersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryActiveDiscoveryScannersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryActiveDiscoveryScannersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryActiveDiscoveryScannersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryActiveDiscoveryScannersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryActiveDiscoveryScannersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryActiveDiscoveryScannersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryActiveDiscoveryScannersOK creates a QueryActiveDiscoveryScannersOK with default headers values
func NewQueryActiveDiscoveryScannersOK() *QueryActiveDiscoveryScannersOK {
	return &QueryActiveDiscoveryScannersOK{}
}

/*
QueryActiveDiscoveryScannersOK describes a response with status code 200, with default header values.

OK
*/
type QueryActiveDiscoveryScannersOK struct {

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

// IsSuccess returns true when this query active discovery scanners o k response has a 2xx status code
func (o *QueryActiveDiscoveryScannersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query active discovery scanners o k response has a 3xx status code
func (o *QueryActiveDiscoveryScannersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query active discovery scanners o k response has a 4xx status code
func (o *QueryActiveDiscoveryScannersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query active discovery scanners o k response has a 5xx status code
func (o *QueryActiveDiscoveryScannersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query active discovery scanners o k response a status code equal to that given
func (o *QueryActiveDiscoveryScannersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query active discovery scanners o k response
func (o *QueryActiveDiscoveryScannersOK) Code() int {
	return 200
}

func (o *QueryActiveDiscoveryScannersOK) Error() string {
	return fmt.Sprintf("[GET /discover/queries/active-discovery-scanners/v1][%d] queryActiveDiscoveryScannersOK  %+v", 200, o.Payload)
}

func (o *QueryActiveDiscoveryScannersOK) String() string {
	return fmt.Sprintf("[GET /discover/queries/active-discovery-scanners/v1][%d] queryActiveDiscoveryScannersOK  %+v", 200, o.Payload)
}

func (o *QueryActiveDiscoveryScannersOK) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *QueryActiveDiscoveryScannersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryActiveDiscoveryScannersBadRequest creates a QueryActiveDiscoveryScannersBadRequest with default headers values
func NewQueryActiveDiscoveryScannersBadRequest() *QueryActiveDiscoveryScannersBadRequest {
	return &QueryActiveDiscoveryScannersBadRequest{}
}

/*
QueryActiveDiscoveryScannersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryActiveDiscoveryScannersBadRequest struct {

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

// IsSuccess returns true when this query active discovery scanners bad request response has a 2xx status code
func (o *QueryActiveDiscoveryScannersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query active discovery scanners bad request response has a 3xx status code
func (o *QueryActiveDiscoveryScannersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query active discovery scanners bad request response has a 4xx status code
func (o *QueryActiveDiscoveryScannersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query active discovery scanners bad request response has a 5xx status code
func (o *QueryActiveDiscoveryScannersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query active discovery scanners bad request response a status code equal to that given
func (o *QueryActiveDiscoveryScannersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query active discovery scanners bad request response
func (o *QueryActiveDiscoveryScannersBadRequest) Code() int {
	return 400
}

func (o *QueryActiveDiscoveryScannersBadRequest) Error() string {
	return fmt.Sprintf("[GET /discover/queries/active-discovery-scanners/v1][%d] queryActiveDiscoveryScannersBadRequest  %+v", 400, o.Payload)
}

func (o *QueryActiveDiscoveryScannersBadRequest) String() string {
	return fmt.Sprintf("[GET /discover/queries/active-discovery-scanners/v1][%d] queryActiveDiscoveryScannersBadRequest  %+v", 400, o.Payload)
}

func (o *QueryActiveDiscoveryScannersBadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *QueryActiveDiscoveryScannersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryActiveDiscoveryScannersForbidden creates a QueryActiveDiscoveryScannersForbidden with default headers values
func NewQueryActiveDiscoveryScannersForbidden() *QueryActiveDiscoveryScannersForbidden {
	return &QueryActiveDiscoveryScannersForbidden{}
}

/*
QueryActiveDiscoveryScannersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryActiveDiscoveryScannersForbidden struct {

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

// IsSuccess returns true when this query active discovery scanners forbidden response has a 2xx status code
func (o *QueryActiveDiscoveryScannersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query active discovery scanners forbidden response has a 3xx status code
func (o *QueryActiveDiscoveryScannersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query active discovery scanners forbidden response has a 4xx status code
func (o *QueryActiveDiscoveryScannersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query active discovery scanners forbidden response has a 5xx status code
func (o *QueryActiveDiscoveryScannersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query active discovery scanners forbidden response a status code equal to that given
func (o *QueryActiveDiscoveryScannersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query active discovery scanners forbidden response
func (o *QueryActiveDiscoveryScannersForbidden) Code() int {
	return 403
}

func (o *QueryActiveDiscoveryScannersForbidden) Error() string {
	return fmt.Sprintf("[GET /discover/queries/active-discovery-scanners/v1][%d] queryActiveDiscoveryScannersForbidden  %+v", 403, o.Payload)
}

func (o *QueryActiveDiscoveryScannersForbidden) String() string {
	return fmt.Sprintf("[GET /discover/queries/active-discovery-scanners/v1][%d] queryActiveDiscoveryScannersForbidden  %+v", 403, o.Payload)
}

func (o *QueryActiveDiscoveryScannersForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryActiveDiscoveryScannersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryActiveDiscoveryScannersTooManyRequests creates a QueryActiveDiscoveryScannersTooManyRequests with default headers values
func NewQueryActiveDiscoveryScannersTooManyRequests() *QueryActiveDiscoveryScannersTooManyRequests {
	return &QueryActiveDiscoveryScannersTooManyRequests{}
}

/*
QueryActiveDiscoveryScannersTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryActiveDiscoveryScannersTooManyRequests struct {

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

// IsSuccess returns true when this query active discovery scanners too many requests response has a 2xx status code
func (o *QueryActiveDiscoveryScannersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query active discovery scanners too many requests response has a 3xx status code
func (o *QueryActiveDiscoveryScannersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query active discovery scanners too many requests response has a 4xx status code
func (o *QueryActiveDiscoveryScannersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query active discovery scanners too many requests response has a 5xx status code
func (o *QueryActiveDiscoveryScannersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query active discovery scanners too many requests response a status code equal to that given
func (o *QueryActiveDiscoveryScannersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query active discovery scanners too many requests response
func (o *QueryActiveDiscoveryScannersTooManyRequests) Code() int {
	return 429
}

func (o *QueryActiveDiscoveryScannersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /discover/queries/active-discovery-scanners/v1][%d] queryActiveDiscoveryScannersTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryActiveDiscoveryScannersTooManyRequests) String() string {
	return fmt.Sprintf("[GET /discover/queries/active-discovery-scanners/v1][%d] queryActiveDiscoveryScannersTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryActiveDiscoveryScannersTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryActiveDiscoveryScannersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryActiveDiscoveryScannersInternalServerError creates a QueryActiveDiscoveryScannersInternalServerError with default headers values
func NewQueryActiveDiscoveryScannersInternalServerError() *QueryActiveDiscoveryScannersInternalServerError {
	return &QueryActiveDiscoveryScannersInternalServerError{}
}

/*
QueryActiveDiscoveryScannersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryActiveDiscoveryScannersInternalServerError struct {

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

// IsSuccess returns true when this query active discovery scanners internal server error response has a 2xx status code
func (o *QueryActiveDiscoveryScannersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query active discovery scanners internal server error response has a 3xx status code
func (o *QueryActiveDiscoveryScannersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query active discovery scanners internal server error response has a 4xx status code
func (o *QueryActiveDiscoveryScannersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query active discovery scanners internal server error response has a 5xx status code
func (o *QueryActiveDiscoveryScannersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query active discovery scanners internal server error response a status code equal to that given
func (o *QueryActiveDiscoveryScannersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query active discovery scanners internal server error response
func (o *QueryActiveDiscoveryScannersInternalServerError) Code() int {
	return 500
}

func (o *QueryActiveDiscoveryScannersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /discover/queries/active-discovery-scanners/v1][%d] queryActiveDiscoveryScannersInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryActiveDiscoveryScannersInternalServerError) String() string {
	return fmt.Sprintf("[GET /discover/queries/active-discovery-scanners/v1][%d] queryActiveDiscoveryScannersInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryActiveDiscoveryScannersInternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *QueryActiveDiscoveryScannersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
