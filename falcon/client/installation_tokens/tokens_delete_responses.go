// Code generated by go-swagger; DO NOT EDIT.

package installation_tokens

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

// TokensDeleteReader is a Reader for the TokensDelete structure.
type TokensDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TokensDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTokensDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTokensDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTokensDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTokensDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewTokensDeleteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTokensDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTokensDeleteOK creates a TokensDeleteOK with default headers values
func NewTokensDeleteOK() *TokensDeleteOK {
	return &TokensDeleteOK{}
}

/*
TokensDeleteOK describes a response with status code 200, with default header values.

OK
*/
type TokensDeleteOK struct {

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

// IsSuccess returns true when this tokens delete o k response has a 2xx status code
func (o *TokensDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tokens delete o k response has a 3xx status code
func (o *TokensDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tokens delete o k response has a 4xx status code
func (o *TokensDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tokens delete o k response has a 5xx status code
func (o *TokensDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tokens delete o k response a status code equal to that given
func (o *TokensDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tokens delete o k response
func (o *TokensDeleteOK) Code() int {
	return 200
}

func (o *TokensDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /installation-tokens/entities/tokens/v1][%d] tokensDeleteOK  %+v", 200, o.Payload)
}

func (o *TokensDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /installation-tokens/entities/tokens/v1][%d] tokensDeleteOK  %+v", 200, o.Payload)
}

func (o *TokensDeleteOK) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *TokensDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewTokensDeleteBadRequest creates a TokensDeleteBadRequest with default headers values
func NewTokensDeleteBadRequest() *TokensDeleteBadRequest {
	return &TokensDeleteBadRequest{}
}

/*
TokensDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TokensDeleteBadRequest struct {

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

// IsSuccess returns true when this tokens delete bad request response has a 2xx status code
func (o *TokensDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tokens delete bad request response has a 3xx status code
func (o *TokensDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tokens delete bad request response has a 4xx status code
func (o *TokensDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this tokens delete bad request response has a 5xx status code
func (o *TokensDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this tokens delete bad request response a status code equal to that given
func (o *TokensDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the tokens delete bad request response
func (o *TokensDeleteBadRequest) Code() int {
	return 400
}

func (o *TokensDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /installation-tokens/entities/tokens/v1][%d] tokensDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *TokensDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /installation-tokens/entities/tokens/v1][%d] tokensDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *TokensDeleteBadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *TokensDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewTokensDeleteForbidden creates a TokensDeleteForbidden with default headers values
func NewTokensDeleteForbidden() *TokensDeleteForbidden {
	return &TokensDeleteForbidden{}
}

/*
TokensDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TokensDeleteForbidden struct {

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

// IsSuccess returns true when this tokens delete forbidden response has a 2xx status code
func (o *TokensDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tokens delete forbidden response has a 3xx status code
func (o *TokensDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tokens delete forbidden response has a 4xx status code
func (o *TokensDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this tokens delete forbidden response has a 5xx status code
func (o *TokensDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this tokens delete forbidden response a status code equal to that given
func (o *TokensDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the tokens delete forbidden response
func (o *TokensDeleteForbidden) Code() int {
	return 403
}

func (o *TokensDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /installation-tokens/entities/tokens/v1][%d] tokensDeleteForbidden  %+v", 403, o.Payload)
}

func (o *TokensDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /installation-tokens/entities/tokens/v1][%d] tokensDeleteForbidden  %+v", 403, o.Payload)
}

func (o *TokensDeleteForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *TokensDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewTokensDeleteNotFound creates a TokensDeleteNotFound with default headers values
func NewTokensDeleteNotFound() *TokensDeleteNotFound {
	return &TokensDeleteNotFound{}
}

/*
TokensDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type TokensDeleteNotFound struct {

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

// IsSuccess returns true when this tokens delete not found response has a 2xx status code
func (o *TokensDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tokens delete not found response has a 3xx status code
func (o *TokensDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tokens delete not found response has a 4xx status code
func (o *TokensDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this tokens delete not found response has a 5xx status code
func (o *TokensDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this tokens delete not found response a status code equal to that given
func (o *TokensDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the tokens delete not found response
func (o *TokensDeleteNotFound) Code() int {
	return 404
}

func (o *TokensDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /installation-tokens/entities/tokens/v1][%d] tokensDeleteNotFound  %+v", 404, o.Payload)
}

func (o *TokensDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /installation-tokens/entities/tokens/v1][%d] tokensDeleteNotFound  %+v", 404, o.Payload)
}

func (o *TokensDeleteNotFound) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *TokensDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewTokensDeleteTooManyRequests creates a TokensDeleteTooManyRequests with default headers values
func NewTokensDeleteTooManyRequests() *TokensDeleteTooManyRequests {
	return &TokensDeleteTooManyRequests{}
}

/*
TokensDeleteTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type TokensDeleteTooManyRequests struct {

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

// IsSuccess returns true when this tokens delete too many requests response has a 2xx status code
func (o *TokensDeleteTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tokens delete too many requests response has a 3xx status code
func (o *TokensDeleteTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tokens delete too many requests response has a 4xx status code
func (o *TokensDeleteTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this tokens delete too many requests response has a 5xx status code
func (o *TokensDeleteTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this tokens delete too many requests response a status code equal to that given
func (o *TokensDeleteTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the tokens delete too many requests response
func (o *TokensDeleteTooManyRequests) Code() int {
	return 429
}

func (o *TokensDeleteTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /installation-tokens/entities/tokens/v1][%d] tokensDeleteTooManyRequests  %+v", 429, o.Payload)
}

func (o *TokensDeleteTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /installation-tokens/entities/tokens/v1][%d] tokensDeleteTooManyRequests  %+v", 429, o.Payload)
}

func (o *TokensDeleteTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *TokensDeleteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewTokensDeleteInternalServerError creates a TokensDeleteInternalServerError with default headers values
func NewTokensDeleteInternalServerError() *TokensDeleteInternalServerError {
	return &TokensDeleteInternalServerError{}
}

/*
TokensDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type TokensDeleteInternalServerError struct {

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

// IsSuccess returns true when this tokens delete internal server error response has a 2xx status code
func (o *TokensDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tokens delete internal server error response has a 3xx status code
func (o *TokensDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tokens delete internal server error response has a 4xx status code
func (o *TokensDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this tokens delete internal server error response has a 5xx status code
func (o *TokensDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this tokens delete internal server error response a status code equal to that given
func (o *TokensDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the tokens delete internal server error response
func (o *TokensDeleteInternalServerError) Code() int {
	return 500
}

func (o *TokensDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /installation-tokens/entities/tokens/v1][%d] tokensDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *TokensDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /installation-tokens/entities/tokens/v1][%d] tokensDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *TokensDeleteInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *TokensDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
