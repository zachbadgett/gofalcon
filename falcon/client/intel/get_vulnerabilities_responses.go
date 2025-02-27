// Code generated by go-swagger; DO NOT EDIT.

package intel

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

// GetVulnerabilitiesReader is a Reader for the GetVulnerabilities structure.
type GetVulnerabilitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVulnerabilitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVulnerabilitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetVulnerabilitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVulnerabilitiesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetVulnerabilitiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVulnerabilitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetVulnerabilitiesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVulnerabilitiesOK creates a GetVulnerabilitiesOK with default headers values
func NewGetVulnerabilitiesOK() *GetVulnerabilitiesOK {
	return &GetVulnerabilitiesOK{}
}

/*
GetVulnerabilitiesOK describes a response with status code 200, with default header values.

OK
*/
type GetVulnerabilitiesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainVulnerabilityResponse
}

// IsSuccess returns true when this get vulnerabilities o k response has a 2xx status code
func (o *GetVulnerabilitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get vulnerabilities o k response has a 3xx status code
func (o *GetVulnerabilitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vulnerabilities o k response has a 4xx status code
func (o *GetVulnerabilitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vulnerabilities o k response has a 5xx status code
func (o *GetVulnerabilitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get vulnerabilities o k response a status code equal to that given
func (o *GetVulnerabilitiesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get vulnerabilities o k response
func (o *GetVulnerabilitiesOK) Code() int {
	return 200
}

func (o *GetVulnerabilitiesOK) Error() string {
	return fmt.Sprintf("[POST /intel/entities/vulnerabilities/GET/v1][%d] getVulnerabilitiesOK  %+v", 200, o.Payload)
}

func (o *GetVulnerabilitiesOK) String() string {
	return fmt.Sprintf("[POST /intel/entities/vulnerabilities/GET/v1][%d] getVulnerabilitiesOK  %+v", 200, o.Payload)
}

func (o *GetVulnerabilitiesOK) GetPayload() *models.DomainVulnerabilityResponse {
	return o.Payload
}

func (o *GetVulnerabilitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainVulnerabilityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVulnerabilitiesForbidden creates a GetVulnerabilitiesForbidden with default headers values
func NewGetVulnerabilitiesForbidden() *GetVulnerabilitiesForbidden {
	return &GetVulnerabilitiesForbidden{}
}

/*
GetVulnerabilitiesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetVulnerabilitiesForbidden struct {

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

// IsSuccess returns true when this get vulnerabilities forbidden response has a 2xx status code
func (o *GetVulnerabilitiesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get vulnerabilities forbidden response has a 3xx status code
func (o *GetVulnerabilitiesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vulnerabilities forbidden response has a 4xx status code
func (o *GetVulnerabilitiesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get vulnerabilities forbidden response has a 5xx status code
func (o *GetVulnerabilitiesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get vulnerabilities forbidden response a status code equal to that given
func (o *GetVulnerabilitiesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get vulnerabilities forbidden response
func (o *GetVulnerabilitiesForbidden) Code() int {
	return 403
}

func (o *GetVulnerabilitiesForbidden) Error() string {
	return fmt.Sprintf("[POST /intel/entities/vulnerabilities/GET/v1][%d] getVulnerabilitiesForbidden  %+v", 403, o.Payload)
}

func (o *GetVulnerabilitiesForbidden) String() string {
	return fmt.Sprintf("[POST /intel/entities/vulnerabilities/GET/v1][%d] getVulnerabilitiesForbidden  %+v", 403, o.Payload)
}

func (o *GetVulnerabilitiesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetVulnerabilitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVulnerabilitiesNotFound creates a GetVulnerabilitiesNotFound with default headers values
func NewGetVulnerabilitiesNotFound() *GetVulnerabilitiesNotFound {
	return &GetVulnerabilitiesNotFound{}
}

/*
GetVulnerabilitiesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetVulnerabilitiesNotFound struct {

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

// IsSuccess returns true when this get vulnerabilities not found response has a 2xx status code
func (o *GetVulnerabilitiesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get vulnerabilities not found response has a 3xx status code
func (o *GetVulnerabilitiesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vulnerabilities not found response has a 4xx status code
func (o *GetVulnerabilitiesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get vulnerabilities not found response has a 5xx status code
func (o *GetVulnerabilitiesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get vulnerabilities not found response a status code equal to that given
func (o *GetVulnerabilitiesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get vulnerabilities not found response
func (o *GetVulnerabilitiesNotFound) Code() int {
	return 404
}

func (o *GetVulnerabilitiesNotFound) Error() string {
	return fmt.Sprintf("[POST /intel/entities/vulnerabilities/GET/v1][%d] getVulnerabilitiesNotFound  %+v", 404, o.Payload)
}

func (o *GetVulnerabilitiesNotFound) String() string {
	return fmt.Sprintf("[POST /intel/entities/vulnerabilities/GET/v1][%d] getVulnerabilitiesNotFound  %+v", 404, o.Payload)
}

func (o *GetVulnerabilitiesNotFound) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetVulnerabilitiesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVulnerabilitiesTooManyRequests creates a GetVulnerabilitiesTooManyRequests with default headers values
func NewGetVulnerabilitiesTooManyRequests() *GetVulnerabilitiesTooManyRequests {
	return &GetVulnerabilitiesTooManyRequests{}
}

/*
GetVulnerabilitiesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetVulnerabilitiesTooManyRequests struct {

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

// IsSuccess returns true when this get vulnerabilities too many requests response has a 2xx status code
func (o *GetVulnerabilitiesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get vulnerabilities too many requests response has a 3xx status code
func (o *GetVulnerabilitiesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vulnerabilities too many requests response has a 4xx status code
func (o *GetVulnerabilitiesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get vulnerabilities too many requests response has a 5xx status code
func (o *GetVulnerabilitiesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get vulnerabilities too many requests response a status code equal to that given
func (o *GetVulnerabilitiesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get vulnerabilities too many requests response
func (o *GetVulnerabilitiesTooManyRequests) Code() int {
	return 429
}

func (o *GetVulnerabilitiesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /intel/entities/vulnerabilities/GET/v1][%d] getVulnerabilitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetVulnerabilitiesTooManyRequests) String() string {
	return fmt.Sprintf("[POST /intel/entities/vulnerabilities/GET/v1][%d] getVulnerabilitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetVulnerabilitiesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetVulnerabilitiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVulnerabilitiesInternalServerError creates a GetVulnerabilitiesInternalServerError with default headers values
func NewGetVulnerabilitiesInternalServerError() *GetVulnerabilitiesInternalServerError {
	return &GetVulnerabilitiesInternalServerError{}
}

/*
GetVulnerabilitiesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetVulnerabilitiesInternalServerError struct {

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

// IsSuccess returns true when this get vulnerabilities internal server error response has a 2xx status code
func (o *GetVulnerabilitiesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get vulnerabilities internal server error response has a 3xx status code
func (o *GetVulnerabilitiesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vulnerabilities internal server error response has a 4xx status code
func (o *GetVulnerabilitiesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vulnerabilities internal server error response has a 5xx status code
func (o *GetVulnerabilitiesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get vulnerabilities internal server error response a status code equal to that given
func (o *GetVulnerabilitiesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get vulnerabilities internal server error response
func (o *GetVulnerabilitiesInternalServerError) Code() int {
	return 500
}

func (o *GetVulnerabilitiesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /intel/entities/vulnerabilities/GET/v1][%d] getVulnerabilitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVulnerabilitiesInternalServerError) String() string {
	return fmt.Sprintf("[POST /intel/entities/vulnerabilities/GET/v1][%d] getVulnerabilitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVulnerabilitiesInternalServerError) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetVulnerabilitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVulnerabilitiesDefault creates a GetVulnerabilitiesDefault with default headers values
func NewGetVulnerabilitiesDefault(code int) *GetVulnerabilitiesDefault {
	return &GetVulnerabilitiesDefault{
		_statusCode: code,
	}
}

/*
GetVulnerabilitiesDefault describes a response with status code -1, with default header values.

OK
*/
type GetVulnerabilitiesDefault struct {
	_statusCode int

	Payload *models.DomainVulnerabilityResponse
}

// IsSuccess returns true when this get vulnerabilities default response has a 2xx status code
func (o *GetVulnerabilitiesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get vulnerabilities default response has a 3xx status code
func (o *GetVulnerabilitiesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get vulnerabilities default response has a 4xx status code
func (o *GetVulnerabilitiesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get vulnerabilities default response has a 5xx status code
func (o *GetVulnerabilitiesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get vulnerabilities default response a status code equal to that given
func (o *GetVulnerabilitiesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get vulnerabilities default response
func (o *GetVulnerabilitiesDefault) Code() int {
	return o._statusCode
}

func (o *GetVulnerabilitiesDefault) Error() string {
	return fmt.Sprintf("[POST /intel/entities/vulnerabilities/GET/v1][%d] GetVulnerabilities default  %+v", o._statusCode, o.Payload)
}

func (o *GetVulnerabilitiesDefault) String() string {
	return fmt.Sprintf("[POST /intel/entities/vulnerabilities/GET/v1][%d] GetVulnerabilities default  %+v", o._statusCode, o.Payload)
}

func (o *GetVulnerabilitiesDefault) GetPayload() *models.DomainVulnerabilityResponse {
	return o.Payload
}

func (o *GetVulnerabilitiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainVulnerabilityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
