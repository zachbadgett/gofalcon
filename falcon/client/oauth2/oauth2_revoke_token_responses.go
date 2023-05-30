// Code generated by go-swagger; DO NOT EDIT.

package oauth2

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

// Oauth2RevokeTokenReader is a Reader for the Oauth2RevokeToken structure.
type Oauth2RevokeTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Oauth2RevokeTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOauth2RevokeTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOauth2RevokeTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOauth2RevokeTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewOauth2RevokeTokenTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOauth2RevokeTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOauth2RevokeTokenOK creates a Oauth2RevokeTokenOK with default headers values
func NewOauth2RevokeTokenOK() *Oauth2RevokeTokenOK {
	return &Oauth2RevokeTokenOK{}
}

/*
Oauth2RevokeTokenOK describes a response with status code 200, with default header values.

Successfully revoked token
*/
type Oauth2RevokeTokenOK struct {

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

// IsSuccess returns true when this oauth2 revoke token o k response has a 2xx status code
func (o *Oauth2RevokeTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this oauth2 revoke token o k response has a 3xx status code
func (o *Oauth2RevokeTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this oauth2 revoke token o k response has a 4xx status code
func (o *Oauth2RevokeTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this oauth2 revoke token o k response has a 5xx status code
func (o *Oauth2RevokeTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this oauth2 revoke token o k response a status code equal to that given
func (o *Oauth2RevokeTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the oauth2 revoke token o k response
func (o *Oauth2RevokeTokenOK) Code() int {
	return 200
}

func (o *Oauth2RevokeTokenOK) Error() string {
	return fmt.Sprintf("[POST /oauth2/revoke][%d] oauth2RevokeTokenOK  %+v", 200, o.Payload)
}

func (o *Oauth2RevokeTokenOK) String() string {
	return fmt.Sprintf("[POST /oauth2/revoke][%d] oauth2RevokeTokenOK  %+v", 200, o.Payload)
}

func (o *Oauth2RevokeTokenOK) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *Oauth2RevokeTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewOauth2RevokeTokenBadRequest creates a Oauth2RevokeTokenBadRequest with default headers values
func NewOauth2RevokeTokenBadRequest() *Oauth2RevokeTokenBadRequest {
	return &Oauth2RevokeTokenBadRequest{}
}

/*
Oauth2RevokeTokenBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type Oauth2RevokeTokenBadRequest struct {

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

// IsSuccess returns true when this oauth2 revoke token bad request response has a 2xx status code
func (o *Oauth2RevokeTokenBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this oauth2 revoke token bad request response has a 3xx status code
func (o *Oauth2RevokeTokenBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this oauth2 revoke token bad request response has a 4xx status code
func (o *Oauth2RevokeTokenBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this oauth2 revoke token bad request response has a 5xx status code
func (o *Oauth2RevokeTokenBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this oauth2 revoke token bad request response a status code equal to that given
func (o *Oauth2RevokeTokenBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the oauth2 revoke token bad request response
func (o *Oauth2RevokeTokenBadRequest) Code() int {
	return 400
}

func (o *Oauth2RevokeTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /oauth2/revoke][%d] oauth2RevokeTokenBadRequest  %+v", 400, o.Payload)
}

func (o *Oauth2RevokeTokenBadRequest) String() string {
	return fmt.Sprintf("[POST /oauth2/revoke][%d] oauth2RevokeTokenBadRequest  %+v", 400, o.Payload)
}

func (o *Oauth2RevokeTokenBadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *Oauth2RevokeTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewOauth2RevokeTokenForbidden creates a Oauth2RevokeTokenForbidden with default headers values
func NewOauth2RevokeTokenForbidden() *Oauth2RevokeTokenForbidden {
	return &Oauth2RevokeTokenForbidden{}
}

/*
Oauth2RevokeTokenForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type Oauth2RevokeTokenForbidden struct {

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

// IsSuccess returns true when this oauth2 revoke token forbidden response has a 2xx status code
func (o *Oauth2RevokeTokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this oauth2 revoke token forbidden response has a 3xx status code
func (o *Oauth2RevokeTokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this oauth2 revoke token forbidden response has a 4xx status code
func (o *Oauth2RevokeTokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this oauth2 revoke token forbidden response has a 5xx status code
func (o *Oauth2RevokeTokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this oauth2 revoke token forbidden response a status code equal to that given
func (o *Oauth2RevokeTokenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the oauth2 revoke token forbidden response
func (o *Oauth2RevokeTokenForbidden) Code() int {
	return 403
}

func (o *Oauth2RevokeTokenForbidden) Error() string {
	return fmt.Sprintf("[POST /oauth2/revoke][%d] oauth2RevokeTokenForbidden  %+v", 403, o.Payload)
}

func (o *Oauth2RevokeTokenForbidden) String() string {
	return fmt.Sprintf("[POST /oauth2/revoke][%d] oauth2RevokeTokenForbidden  %+v", 403, o.Payload)
}

func (o *Oauth2RevokeTokenForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *Oauth2RevokeTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewOauth2RevokeTokenTooManyRequests creates a Oauth2RevokeTokenTooManyRequests with default headers values
func NewOauth2RevokeTokenTooManyRequests() *Oauth2RevokeTokenTooManyRequests {
	return &Oauth2RevokeTokenTooManyRequests{}
}

/*
Oauth2RevokeTokenTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type Oauth2RevokeTokenTooManyRequests struct {

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

// IsSuccess returns true when this oauth2 revoke token too many requests response has a 2xx status code
func (o *Oauth2RevokeTokenTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this oauth2 revoke token too many requests response has a 3xx status code
func (o *Oauth2RevokeTokenTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this oauth2 revoke token too many requests response has a 4xx status code
func (o *Oauth2RevokeTokenTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this oauth2 revoke token too many requests response has a 5xx status code
func (o *Oauth2RevokeTokenTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this oauth2 revoke token too many requests response a status code equal to that given
func (o *Oauth2RevokeTokenTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the oauth2 revoke token too many requests response
func (o *Oauth2RevokeTokenTooManyRequests) Code() int {
	return 429
}

func (o *Oauth2RevokeTokenTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /oauth2/revoke][%d] oauth2RevokeTokenTooManyRequests  %+v", 429, o.Payload)
}

func (o *Oauth2RevokeTokenTooManyRequests) String() string {
	return fmt.Sprintf("[POST /oauth2/revoke][%d] oauth2RevokeTokenTooManyRequests  %+v", 429, o.Payload)
}

func (o *Oauth2RevokeTokenTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *Oauth2RevokeTokenTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewOauth2RevokeTokenInternalServerError creates a Oauth2RevokeTokenInternalServerError with default headers values
func NewOauth2RevokeTokenInternalServerError() *Oauth2RevokeTokenInternalServerError {
	return &Oauth2RevokeTokenInternalServerError{}
}

/*
Oauth2RevokeTokenInternalServerError describes a response with status code 500, with default header values.

Failed to revoke token
*/
type Oauth2RevokeTokenInternalServerError struct {

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

// IsSuccess returns true when this oauth2 revoke token internal server error response has a 2xx status code
func (o *Oauth2RevokeTokenInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this oauth2 revoke token internal server error response has a 3xx status code
func (o *Oauth2RevokeTokenInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this oauth2 revoke token internal server error response has a 4xx status code
func (o *Oauth2RevokeTokenInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this oauth2 revoke token internal server error response has a 5xx status code
func (o *Oauth2RevokeTokenInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this oauth2 revoke token internal server error response a status code equal to that given
func (o *Oauth2RevokeTokenInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the oauth2 revoke token internal server error response
func (o *Oauth2RevokeTokenInternalServerError) Code() int {
	return 500
}

func (o *Oauth2RevokeTokenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /oauth2/revoke][%d] oauth2RevokeTokenInternalServerError  %+v", 500, o.Payload)
}

func (o *Oauth2RevokeTokenInternalServerError) String() string {
	return fmt.Sprintf("[POST /oauth2/revoke][%d] oauth2RevokeTokenInternalServerError  %+v", 500, o.Payload)
}

func (o *Oauth2RevokeTokenInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *Oauth2RevokeTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
