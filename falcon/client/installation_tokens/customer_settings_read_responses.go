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

// CustomerSettingsReadReader is a Reader for the CustomerSettingsRead structure.
type CustomerSettingsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerSettingsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerSettingsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCustomerSettingsReadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCustomerSettingsReadForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCustomerSettingsReadTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCustomerSettingsReadInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCustomerSettingsReadOK creates a CustomerSettingsReadOK with default headers values
func NewCustomerSettingsReadOK() *CustomerSettingsReadOK {
	return &CustomerSettingsReadOK{}
}

/*
CustomerSettingsReadOK describes a response with status code 200, with default header values.

OK
*/
type CustomerSettingsReadOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APICustomerSettingsResponseV1
}

// IsSuccess returns true when this customer settings read o k response has a 2xx status code
func (o *CustomerSettingsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer settings read o k response has a 3xx status code
func (o *CustomerSettingsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer settings read o k response has a 4xx status code
func (o *CustomerSettingsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer settings read o k response has a 5xx status code
func (o *CustomerSettingsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer settings read o k response a status code equal to that given
func (o *CustomerSettingsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer settings read o k response
func (o *CustomerSettingsReadOK) Code() int {
	return 200
}

func (o *CustomerSettingsReadOK) Error() string {
	return fmt.Sprintf("[GET /installation-tokens/entities/customer-settings/v1][%d] customerSettingsReadOK  %+v", 200, o.Payload)
}

func (o *CustomerSettingsReadOK) String() string {
	return fmt.Sprintf("[GET /installation-tokens/entities/customer-settings/v1][%d] customerSettingsReadOK  %+v", 200, o.Payload)
}

func (o *CustomerSettingsReadOK) GetPayload() *models.APICustomerSettingsResponseV1 {
	return o.Payload
}

func (o *CustomerSettingsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APICustomerSettingsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerSettingsReadBadRequest creates a CustomerSettingsReadBadRequest with default headers values
func NewCustomerSettingsReadBadRequest() *CustomerSettingsReadBadRequest {
	return &CustomerSettingsReadBadRequest{}
}

/*
CustomerSettingsReadBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CustomerSettingsReadBadRequest struct {

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

// IsSuccess returns true when this customer settings read bad request response has a 2xx status code
func (o *CustomerSettingsReadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this customer settings read bad request response has a 3xx status code
func (o *CustomerSettingsReadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer settings read bad request response has a 4xx status code
func (o *CustomerSettingsReadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this customer settings read bad request response has a 5xx status code
func (o *CustomerSettingsReadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this customer settings read bad request response a status code equal to that given
func (o *CustomerSettingsReadBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the customer settings read bad request response
func (o *CustomerSettingsReadBadRequest) Code() int {
	return 400
}

func (o *CustomerSettingsReadBadRequest) Error() string {
	return fmt.Sprintf("[GET /installation-tokens/entities/customer-settings/v1][%d] customerSettingsReadBadRequest  %+v", 400, o.Payload)
}

func (o *CustomerSettingsReadBadRequest) String() string {
	return fmt.Sprintf("[GET /installation-tokens/entities/customer-settings/v1][%d] customerSettingsReadBadRequest  %+v", 400, o.Payload)
}

func (o *CustomerSettingsReadBadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CustomerSettingsReadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCustomerSettingsReadForbidden creates a CustomerSettingsReadForbidden with default headers values
func NewCustomerSettingsReadForbidden() *CustomerSettingsReadForbidden {
	return &CustomerSettingsReadForbidden{}
}

/*
CustomerSettingsReadForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CustomerSettingsReadForbidden struct {

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

// IsSuccess returns true when this customer settings read forbidden response has a 2xx status code
func (o *CustomerSettingsReadForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this customer settings read forbidden response has a 3xx status code
func (o *CustomerSettingsReadForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer settings read forbidden response has a 4xx status code
func (o *CustomerSettingsReadForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this customer settings read forbidden response has a 5xx status code
func (o *CustomerSettingsReadForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this customer settings read forbidden response a status code equal to that given
func (o *CustomerSettingsReadForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the customer settings read forbidden response
func (o *CustomerSettingsReadForbidden) Code() int {
	return 403
}

func (o *CustomerSettingsReadForbidden) Error() string {
	return fmt.Sprintf("[GET /installation-tokens/entities/customer-settings/v1][%d] customerSettingsReadForbidden  %+v", 403, o.Payload)
}

func (o *CustomerSettingsReadForbidden) String() string {
	return fmt.Sprintf("[GET /installation-tokens/entities/customer-settings/v1][%d] customerSettingsReadForbidden  %+v", 403, o.Payload)
}

func (o *CustomerSettingsReadForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CustomerSettingsReadForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCustomerSettingsReadTooManyRequests creates a CustomerSettingsReadTooManyRequests with default headers values
func NewCustomerSettingsReadTooManyRequests() *CustomerSettingsReadTooManyRequests {
	return &CustomerSettingsReadTooManyRequests{}
}

/*
CustomerSettingsReadTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CustomerSettingsReadTooManyRequests struct {

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

// IsSuccess returns true when this customer settings read too many requests response has a 2xx status code
func (o *CustomerSettingsReadTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this customer settings read too many requests response has a 3xx status code
func (o *CustomerSettingsReadTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer settings read too many requests response has a 4xx status code
func (o *CustomerSettingsReadTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this customer settings read too many requests response has a 5xx status code
func (o *CustomerSettingsReadTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this customer settings read too many requests response a status code equal to that given
func (o *CustomerSettingsReadTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the customer settings read too many requests response
func (o *CustomerSettingsReadTooManyRequests) Code() int {
	return 429
}

func (o *CustomerSettingsReadTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /installation-tokens/entities/customer-settings/v1][%d] customerSettingsReadTooManyRequests  %+v", 429, o.Payload)
}

func (o *CustomerSettingsReadTooManyRequests) String() string {
	return fmt.Sprintf("[GET /installation-tokens/entities/customer-settings/v1][%d] customerSettingsReadTooManyRequests  %+v", 429, o.Payload)
}

func (o *CustomerSettingsReadTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CustomerSettingsReadTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCustomerSettingsReadInternalServerError creates a CustomerSettingsReadInternalServerError with default headers values
func NewCustomerSettingsReadInternalServerError() *CustomerSettingsReadInternalServerError {
	return &CustomerSettingsReadInternalServerError{}
}

/*
CustomerSettingsReadInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CustomerSettingsReadInternalServerError struct {

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

// IsSuccess returns true when this customer settings read internal server error response has a 2xx status code
func (o *CustomerSettingsReadInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this customer settings read internal server error response has a 3xx status code
func (o *CustomerSettingsReadInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer settings read internal server error response has a 4xx status code
func (o *CustomerSettingsReadInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer settings read internal server error response has a 5xx status code
func (o *CustomerSettingsReadInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this customer settings read internal server error response a status code equal to that given
func (o *CustomerSettingsReadInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the customer settings read internal server error response
func (o *CustomerSettingsReadInternalServerError) Code() int {
	return 500
}

func (o *CustomerSettingsReadInternalServerError) Error() string {
	return fmt.Sprintf("[GET /installation-tokens/entities/customer-settings/v1][%d] customerSettingsReadInternalServerError  %+v", 500, o.Payload)
}

func (o *CustomerSettingsReadInternalServerError) String() string {
	return fmt.Sprintf("[GET /installation-tokens/entities/customer-settings/v1][%d] customerSettingsReadInternalServerError  %+v", 500, o.Payload)
}

func (o *CustomerSettingsReadInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CustomerSettingsReadInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
