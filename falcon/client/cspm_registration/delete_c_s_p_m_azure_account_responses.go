// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

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

// DeleteCSPMAzureAccountReader is a Reader for the DeleteCSPMAzureAccount structure.
type DeleteCSPMAzureAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCSPMAzureAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCSPMAzureAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewDeleteCSPMAzureAccountMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCSPMAzureAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteCSPMAzureAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteCSPMAzureAccountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteCSPMAzureAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteCSPMAzureAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCSPMAzureAccountOK creates a DeleteCSPMAzureAccountOK with default headers values
func NewDeleteCSPMAzureAccountOK() *DeleteCSPMAzureAccountOK {
	return &DeleteCSPMAzureAccountOK{}
}

/*
DeleteCSPMAzureAccountOK describes a response with status code 200, with default header values.

OK
*/
type DeleteCSPMAzureAccountOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationBaseResponseV1
}

// IsSuccess returns true when this delete c s p m azure account o k response has a 2xx status code
func (o *DeleteCSPMAzureAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete c s p m azure account o k response has a 3xx status code
func (o *DeleteCSPMAzureAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete c s p m azure account o k response has a 4xx status code
func (o *DeleteCSPMAzureAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete c s p m azure account o k response has a 5xx status code
func (o *DeleteCSPMAzureAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete c s p m azure account o k response a status code equal to that given
func (o *DeleteCSPMAzureAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete c s p m azure account o k response
func (o *DeleteCSPMAzureAccountOK) Code() int {
	return 200
}

func (o *DeleteCSPMAzureAccountOK) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-azure/entities/account/v1][%d] deleteCSPMAzureAccountOK  %+v", 200, o.Payload)
}

func (o *DeleteCSPMAzureAccountOK) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-azure/entities/account/v1][%d] deleteCSPMAzureAccountOK  %+v", 200, o.Payload)
}

func (o *DeleteCSPMAzureAccountOK) GetPayload() *models.RegistrationBaseResponseV1 {
	return o.Payload
}

func (o *DeleteCSPMAzureAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationBaseResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCSPMAzureAccountMultiStatus creates a DeleteCSPMAzureAccountMultiStatus with default headers values
func NewDeleteCSPMAzureAccountMultiStatus() *DeleteCSPMAzureAccountMultiStatus {
	return &DeleteCSPMAzureAccountMultiStatus{}
}

/*
DeleteCSPMAzureAccountMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type DeleteCSPMAzureAccountMultiStatus struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationBaseResponseV1
}

// IsSuccess returns true when this delete c s p m azure account multi status response has a 2xx status code
func (o *DeleteCSPMAzureAccountMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete c s p m azure account multi status response has a 3xx status code
func (o *DeleteCSPMAzureAccountMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete c s p m azure account multi status response has a 4xx status code
func (o *DeleteCSPMAzureAccountMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete c s p m azure account multi status response has a 5xx status code
func (o *DeleteCSPMAzureAccountMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this delete c s p m azure account multi status response a status code equal to that given
func (o *DeleteCSPMAzureAccountMultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the delete c s p m azure account multi status response
func (o *DeleteCSPMAzureAccountMultiStatus) Code() int {
	return 207
}

func (o *DeleteCSPMAzureAccountMultiStatus) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-azure/entities/account/v1][%d] deleteCSPMAzureAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *DeleteCSPMAzureAccountMultiStatus) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-azure/entities/account/v1][%d] deleteCSPMAzureAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *DeleteCSPMAzureAccountMultiStatus) GetPayload() *models.RegistrationBaseResponseV1 {
	return o.Payload
}

func (o *DeleteCSPMAzureAccountMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationBaseResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCSPMAzureAccountBadRequest creates a DeleteCSPMAzureAccountBadRequest with default headers values
func NewDeleteCSPMAzureAccountBadRequest() *DeleteCSPMAzureAccountBadRequest {
	return &DeleteCSPMAzureAccountBadRequest{}
}

/*
DeleteCSPMAzureAccountBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteCSPMAzureAccountBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationBaseResponseV1
}

// IsSuccess returns true when this delete c s p m azure account bad request response has a 2xx status code
func (o *DeleteCSPMAzureAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete c s p m azure account bad request response has a 3xx status code
func (o *DeleteCSPMAzureAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete c s p m azure account bad request response has a 4xx status code
func (o *DeleteCSPMAzureAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete c s p m azure account bad request response has a 5xx status code
func (o *DeleteCSPMAzureAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete c s p m azure account bad request response a status code equal to that given
func (o *DeleteCSPMAzureAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete c s p m azure account bad request response
func (o *DeleteCSPMAzureAccountBadRequest) Code() int {
	return 400
}

func (o *DeleteCSPMAzureAccountBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-azure/entities/account/v1][%d] deleteCSPMAzureAccountBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteCSPMAzureAccountBadRequest) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-azure/entities/account/v1][%d] deleteCSPMAzureAccountBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteCSPMAzureAccountBadRequest) GetPayload() *models.RegistrationBaseResponseV1 {
	return o.Payload
}

func (o *DeleteCSPMAzureAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationBaseResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCSPMAzureAccountForbidden creates a DeleteCSPMAzureAccountForbidden with default headers values
func NewDeleteCSPMAzureAccountForbidden() *DeleteCSPMAzureAccountForbidden {
	return &DeleteCSPMAzureAccountForbidden{}
}

/*
DeleteCSPMAzureAccountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteCSPMAzureAccountForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this delete c s p m azure account forbidden response has a 2xx status code
func (o *DeleteCSPMAzureAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete c s p m azure account forbidden response has a 3xx status code
func (o *DeleteCSPMAzureAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete c s p m azure account forbidden response has a 4xx status code
func (o *DeleteCSPMAzureAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete c s p m azure account forbidden response has a 5xx status code
func (o *DeleteCSPMAzureAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete c s p m azure account forbidden response a status code equal to that given
func (o *DeleteCSPMAzureAccountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete c s p m azure account forbidden response
func (o *DeleteCSPMAzureAccountForbidden) Code() int {
	return 403
}

func (o *DeleteCSPMAzureAccountForbidden) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-azure/entities/account/v1][%d] deleteCSPMAzureAccountForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCSPMAzureAccountForbidden) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-azure/entities/account/v1][%d] deleteCSPMAzureAccountForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCSPMAzureAccountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteCSPMAzureAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteCSPMAzureAccountTooManyRequests creates a DeleteCSPMAzureAccountTooManyRequests with default headers values
func NewDeleteCSPMAzureAccountTooManyRequests() *DeleteCSPMAzureAccountTooManyRequests {
	return &DeleteCSPMAzureAccountTooManyRequests{}
}

/*
DeleteCSPMAzureAccountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteCSPMAzureAccountTooManyRequests struct {

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

// IsSuccess returns true when this delete c s p m azure account too many requests response has a 2xx status code
func (o *DeleteCSPMAzureAccountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete c s p m azure account too many requests response has a 3xx status code
func (o *DeleteCSPMAzureAccountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete c s p m azure account too many requests response has a 4xx status code
func (o *DeleteCSPMAzureAccountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete c s p m azure account too many requests response has a 5xx status code
func (o *DeleteCSPMAzureAccountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete c s p m azure account too many requests response a status code equal to that given
func (o *DeleteCSPMAzureAccountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the delete c s p m azure account too many requests response
func (o *DeleteCSPMAzureAccountTooManyRequests) Code() int {
	return 429
}

func (o *DeleteCSPMAzureAccountTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-azure/entities/account/v1][%d] deleteCSPMAzureAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteCSPMAzureAccountTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-azure/entities/account/v1][%d] deleteCSPMAzureAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteCSPMAzureAccountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteCSPMAzureAccountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteCSPMAzureAccountInternalServerError creates a DeleteCSPMAzureAccountInternalServerError with default headers values
func NewDeleteCSPMAzureAccountInternalServerError() *DeleteCSPMAzureAccountInternalServerError {
	return &DeleteCSPMAzureAccountInternalServerError{}
}

/*
DeleteCSPMAzureAccountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteCSPMAzureAccountInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAzureAccountResponseV1
}

// IsSuccess returns true when this delete c s p m azure account internal server error response has a 2xx status code
func (o *DeleteCSPMAzureAccountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete c s p m azure account internal server error response has a 3xx status code
func (o *DeleteCSPMAzureAccountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete c s p m azure account internal server error response has a 4xx status code
func (o *DeleteCSPMAzureAccountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete c s p m azure account internal server error response has a 5xx status code
func (o *DeleteCSPMAzureAccountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete c s p m azure account internal server error response a status code equal to that given
func (o *DeleteCSPMAzureAccountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete c s p m azure account internal server error response
func (o *DeleteCSPMAzureAccountInternalServerError) Code() int {
	return 500
}

func (o *DeleteCSPMAzureAccountInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-azure/entities/account/v1][%d] deleteCSPMAzureAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCSPMAzureAccountInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-azure/entities/account/v1][%d] deleteCSPMAzureAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCSPMAzureAccountInternalServerError) GetPayload() *models.RegistrationAzureAccountResponseV1 {
	return o.Payload
}

func (o *DeleteCSPMAzureAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAzureAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCSPMAzureAccountDefault creates a DeleteCSPMAzureAccountDefault with default headers values
func NewDeleteCSPMAzureAccountDefault(code int) *DeleteCSPMAzureAccountDefault {
	return &DeleteCSPMAzureAccountDefault{
		_statusCode: code,
	}
}

/*
DeleteCSPMAzureAccountDefault describes a response with status code -1, with default header values.

OK
*/
type DeleteCSPMAzureAccountDefault struct {
	_statusCode int

	Payload *models.RegistrationBaseResponseV1
}

// IsSuccess returns true when this delete c s p m azure account default response has a 2xx status code
func (o *DeleteCSPMAzureAccountDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete c s p m azure account default response has a 3xx status code
func (o *DeleteCSPMAzureAccountDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete c s p m azure account default response has a 4xx status code
func (o *DeleteCSPMAzureAccountDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete c s p m azure account default response has a 5xx status code
func (o *DeleteCSPMAzureAccountDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete c s p m azure account default response a status code equal to that given
func (o *DeleteCSPMAzureAccountDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete c s p m azure account default response
func (o *DeleteCSPMAzureAccountDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCSPMAzureAccountDefault) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-azure/entities/account/v1][%d] DeleteCSPMAzureAccount default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCSPMAzureAccountDefault) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-azure/entities/account/v1][%d] DeleteCSPMAzureAccount default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCSPMAzureAccountDefault) GetPayload() *models.RegistrationBaseResponseV1 {
	return o.Payload
}

func (o *DeleteCSPMAzureAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegistrationBaseResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
