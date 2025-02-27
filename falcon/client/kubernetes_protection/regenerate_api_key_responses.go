// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_protection

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

// RegenerateAPIKeyReader is a Reader for the RegenerateAPIKey structure.
type RegenerateAPIKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegenerateAPIKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegenerateAPIKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewRegenerateAPIKeyMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRegenerateAPIKeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRegenerateAPIKeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRegenerateAPIKeyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRegenerateAPIKeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRegenerateAPIKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRegenerateAPIKeyOK creates a RegenerateAPIKeyOK with default headers values
func NewRegenerateAPIKeyOK() *RegenerateAPIKeyOK {
	return &RegenerateAPIKeyOK{}
}

/*
RegenerateAPIKeyOK describes a response with status code 200, with default header values.

OK
*/
type RegenerateAPIKeyOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.K8sregRegenAPIKeyResp
}

// IsSuccess returns true when this regenerate Api key o k response has a 2xx status code
func (o *RegenerateAPIKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this regenerate Api key o k response has a 3xx status code
func (o *RegenerateAPIKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this regenerate Api key o k response has a 4xx status code
func (o *RegenerateAPIKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this regenerate Api key o k response has a 5xx status code
func (o *RegenerateAPIKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this regenerate Api key o k response a status code equal to that given
func (o *RegenerateAPIKeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the regenerate Api key o k response
func (o *RegenerateAPIKeyOK) Code() int {
	return 200
}

func (o *RegenerateAPIKeyOK) Error() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/integration/api-key/v1][%d] regenerateApiKeyOK  %+v", 200, o.Payload)
}

func (o *RegenerateAPIKeyOK) String() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/integration/api-key/v1][%d] regenerateApiKeyOK  %+v", 200, o.Payload)
}

func (o *RegenerateAPIKeyOK) GetPayload() *models.K8sregRegenAPIKeyResp {
	return o.Payload
}

func (o *RegenerateAPIKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.K8sregRegenAPIKeyResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegenerateAPIKeyMultiStatus creates a RegenerateAPIKeyMultiStatus with default headers values
func NewRegenerateAPIKeyMultiStatus() *RegenerateAPIKeyMultiStatus {
	return &RegenerateAPIKeyMultiStatus{}
}

/*
RegenerateAPIKeyMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type RegenerateAPIKeyMultiStatus struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.K8sregRegenAPIKeyResp
}

// IsSuccess returns true when this regenerate Api key multi status response has a 2xx status code
func (o *RegenerateAPIKeyMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this regenerate Api key multi status response has a 3xx status code
func (o *RegenerateAPIKeyMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this regenerate Api key multi status response has a 4xx status code
func (o *RegenerateAPIKeyMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this regenerate Api key multi status response has a 5xx status code
func (o *RegenerateAPIKeyMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this regenerate Api key multi status response a status code equal to that given
func (o *RegenerateAPIKeyMultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the regenerate Api key multi status response
func (o *RegenerateAPIKeyMultiStatus) Code() int {
	return 207
}

func (o *RegenerateAPIKeyMultiStatus) Error() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/integration/api-key/v1][%d] regenerateApiKeyMultiStatus  %+v", 207, o.Payload)
}

func (o *RegenerateAPIKeyMultiStatus) String() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/integration/api-key/v1][%d] regenerateApiKeyMultiStatus  %+v", 207, o.Payload)
}

func (o *RegenerateAPIKeyMultiStatus) GetPayload() *models.K8sregRegenAPIKeyResp {
	return o.Payload
}

func (o *RegenerateAPIKeyMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.K8sregRegenAPIKeyResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegenerateAPIKeyBadRequest creates a RegenerateAPIKeyBadRequest with default headers values
func NewRegenerateAPIKeyBadRequest() *RegenerateAPIKeyBadRequest {
	return &RegenerateAPIKeyBadRequest{}
}

/*
RegenerateAPIKeyBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RegenerateAPIKeyBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.K8sregRegenAPIKeyResp
}

// IsSuccess returns true when this regenerate Api key bad request response has a 2xx status code
func (o *RegenerateAPIKeyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this regenerate Api key bad request response has a 3xx status code
func (o *RegenerateAPIKeyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this regenerate Api key bad request response has a 4xx status code
func (o *RegenerateAPIKeyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this regenerate Api key bad request response has a 5xx status code
func (o *RegenerateAPIKeyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this regenerate Api key bad request response a status code equal to that given
func (o *RegenerateAPIKeyBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the regenerate Api key bad request response
func (o *RegenerateAPIKeyBadRequest) Code() int {
	return 400
}

func (o *RegenerateAPIKeyBadRequest) Error() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/integration/api-key/v1][%d] regenerateApiKeyBadRequest  %+v", 400, o.Payload)
}

func (o *RegenerateAPIKeyBadRequest) String() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/integration/api-key/v1][%d] regenerateApiKeyBadRequest  %+v", 400, o.Payload)
}

func (o *RegenerateAPIKeyBadRequest) GetPayload() *models.K8sregRegenAPIKeyResp {
	return o.Payload
}

func (o *RegenerateAPIKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.K8sregRegenAPIKeyResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegenerateAPIKeyForbidden creates a RegenerateAPIKeyForbidden with default headers values
func NewRegenerateAPIKeyForbidden() *RegenerateAPIKeyForbidden {
	return &RegenerateAPIKeyForbidden{}
}

/*
RegenerateAPIKeyForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RegenerateAPIKeyForbidden struct {

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

// IsSuccess returns true when this regenerate Api key forbidden response has a 2xx status code
func (o *RegenerateAPIKeyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this regenerate Api key forbidden response has a 3xx status code
func (o *RegenerateAPIKeyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this regenerate Api key forbidden response has a 4xx status code
func (o *RegenerateAPIKeyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this regenerate Api key forbidden response has a 5xx status code
func (o *RegenerateAPIKeyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this regenerate Api key forbidden response a status code equal to that given
func (o *RegenerateAPIKeyForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the regenerate Api key forbidden response
func (o *RegenerateAPIKeyForbidden) Code() int {
	return 403
}

func (o *RegenerateAPIKeyForbidden) Error() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/integration/api-key/v1][%d] regenerateApiKeyForbidden  %+v", 403, o.Payload)
}

func (o *RegenerateAPIKeyForbidden) String() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/integration/api-key/v1][%d] regenerateApiKeyForbidden  %+v", 403, o.Payload)
}

func (o *RegenerateAPIKeyForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RegenerateAPIKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRegenerateAPIKeyTooManyRequests creates a RegenerateAPIKeyTooManyRequests with default headers values
func NewRegenerateAPIKeyTooManyRequests() *RegenerateAPIKeyTooManyRequests {
	return &RegenerateAPIKeyTooManyRequests{}
}

/*
RegenerateAPIKeyTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RegenerateAPIKeyTooManyRequests struct {

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

// IsSuccess returns true when this regenerate Api key too many requests response has a 2xx status code
func (o *RegenerateAPIKeyTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this regenerate Api key too many requests response has a 3xx status code
func (o *RegenerateAPIKeyTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this regenerate Api key too many requests response has a 4xx status code
func (o *RegenerateAPIKeyTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this regenerate Api key too many requests response has a 5xx status code
func (o *RegenerateAPIKeyTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this regenerate Api key too many requests response a status code equal to that given
func (o *RegenerateAPIKeyTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the regenerate Api key too many requests response
func (o *RegenerateAPIKeyTooManyRequests) Code() int {
	return 429
}

func (o *RegenerateAPIKeyTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/integration/api-key/v1][%d] regenerateApiKeyTooManyRequests  %+v", 429, o.Payload)
}

func (o *RegenerateAPIKeyTooManyRequests) String() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/integration/api-key/v1][%d] regenerateApiKeyTooManyRequests  %+v", 429, o.Payload)
}

func (o *RegenerateAPIKeyTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RegenerateAPIKeyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRegenerateAPIKeyInternalServerError creates a RegenerateAPIKeyInternalServerError with default headers values
func NewRegenerateAPIKeyInternalServerError() *RegenerateAPIKeyInternalServerError {
	return &RegenerateAPIKeyInternalServerError{}
}

/*
RegenerateAPIKeyInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type RegenerateAPIKeyInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.K8sregRegenAPIKeyResp
}

// IsSuccess returns true when this regenerate Api key internal server error response has a 2xx status code
func (o *RegenerateAPIKeyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this regenerate Api key internal server error response has a 3xx status code
func (o *RegenerateAPIKeyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this regenerate Api key internal server error response has a 4xx status code
func (o *RegenerateAPIKeyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this regenerate Api key internal server error response has a 5xx status code
func (o *RegenerateAPIKeyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this regenerate Api key internal server error response a status code equal to that given
func (o *RegenerateAPIKeyInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the regenerate Api key internal server error response
func (o *RegenerateAPIKeyInternalServerError) Code() int {
	return 500
}

func (o *RegenerateAPIKeyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/integration/api-key/v1][%d] regenerateApiKeyInternalServerError  %+v", 500, o.Payload)
}

func (o *RegenerateAPIKeyInternalServerError) String() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/integration/api-key/v1][%d] regenerateApiKeyInternalServerError  %+v", 500, o.Payload)
}

func (o *RegenerateAPIKeyInternalServerError) GetPayload() *models.K8sregRegenAPIKeyResp {
	return o.Payload
}

func (o *RegenerateAPIKeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.K8sregRegenAPIKeyResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegenerateAPIKeyDefault creates a RegenerateAPIKeyDefault with default headers values
func NewRegenerateAPIKeyDefault(code int) *RegenerateAPIKeyDefault {
	return &RegenerateAPIKeyDefault{
		_statusCode: code,
	}
}

/*
RegenerateAPIKeyDefault describes a response with status code -1, with default header values.

OK
*/
type RegenerateAPIKeyDefault struct {
	_statusCode int

	Payload *models.K8sregRegenAPIKeyResp
}

// IsSuccess returns true when this regenerate API key default response has a 2xx status code
func (o *RegenerateAPIKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this regenerate API key default response has a 3xx status code
func (o *RegenerateAPIKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this regenerate API key default response has a 4xx status code
func (o *RegenerateAPIKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this regenerate API key default response has a 5xx status code
func (o *RegenerateAPIKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this regenerate API key default response a status code equal to that given
func (o *RegenerateAPIKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the regenerate API key default response
func (o *RegenerateAPIKeyDefault) Code() int {
	return o._statusCode
}

func (o *RegenerateAPIKeyDefault) Error() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/integration/api-key/v1][%d] RegenerateAPIKey default  %+v", o._statusCode, o.Payload)
}

func (o *RegenerateAPIKeyDefault) String() string {
	return fmt.Sprintf("[POST /kubernetes-protection/entities/integration/api-key/v1][%d] RegenerateAPIKey default  %+v", o._statusCode, o.Payload)
}

func (o *RegenerateAPIKeyDefault) GetPayload() *models.K8sregRegenAPIKeyResp {
	return o.Payload
}

func (o *RegenerateAPIKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.K8sregRegenAPIKeyResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
