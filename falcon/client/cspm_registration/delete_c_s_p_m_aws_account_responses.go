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

// DeleteCSPMAwsAccountReader is a Reader for the DeleteCSPMAwsAccount structure.
type DeleteCSPMAwsAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCSPMAwsAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCSPMAwsAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewDeleteCSPMAwsAccountMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCSPMAwsAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteCSPMAwsAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteCSPMAwsAccountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteCSPMAwsAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCSPMAwsAccountOK creates a DeleteCSPMAwsAccountOK with default headers values
func NewDeleteCSPMAwsAccountOK() *DeleteCSPMAwsAccountOK {
	return &DeleteCSPMAwsAccountOK{}
}

/*
DeleteCSPMAwsAccountOK describes a response with status code 200, with default header values.

OK
*/
type DeleteCSPMAwsAccountOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaBaseEntitiesResponse
}

// IsSuccess returns true when this delete c s p m aws account o k response has a 2xx status code
func (o *DeleteCSPMAwsAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete c s p m aws account o k response has a 3xx status code
func (o *DeleteCSPMAwsAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete c s p m aws account o k response has a 4xx status code
func (o *DeleteCSPMAwsAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete c s p m aws account o k response has a 5xx status code
func (o *DeleteCSPMAwsAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete c s p m aws account o k response a status code equal to that given
func (o *DeleteCSPMAwsAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete c s p m aws account o k response
func (o *DeleteCSPMAwsAccountOK) Code() int {
	return 200
}

func (o *DeleteCSPMAwsAccountOK) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-aws/entities/account/v1][%d] deleteCSPMAwsAccountOK  %+v", 200, o.Payload)
}

func (o *DeleteCSPMAwsAccountOK) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-aws/entities/account/v1][%d] deleteCSPMAwsAccountOK  %+v", 200, o.Payload)
}

func (o *DeleteCSPMAwsAccountOK) GetPayload() *models.MsaBaseEntitiesResponse {
	return o.Payload
}

func (o *DeleteCSPMAwsAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaBaseEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCSPMAwsAccountMultiStatus creates a DeleteCSPMAwsAccountMultiStatus with default headers values
func NewDeleteCSPMAwsAccountMultiStatus() *DeleteCSPMAwsAccountMultiStatus {
	return &DeleteCSPMAwsAccountMultiStatus{}
}

/*
DeleteCSPMAwsAccountMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type DeleteCSPMAwsAccountMultiStatus struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaBaseEntitiesResponse
}

// IsSuccess returns true when this delete c s p m aws account multi status response has a 2xx status code
func (o *DeleteCSPMAwsAccountMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete c s p m aws account multi status response has a 3xx status code
func (o *DeleteCSPMAwsAccountMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete c s p m aws account multi status response has a 4xx status code
func (o *DeleteCSPMAwsAccountMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete c s p m aws account multi status response has a 5xx status code
func (o *DeleteCSPMAwsAccountMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this delete c s p m aws account multi status response a status code equal to that given
func (o *DeleteCSPMAwsAccountMultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the delete c s p m aws account multi status response
func (o *DeleteCSPMAwsAccountMultiStatus) Code() int {
	return 207
}

func (o *DeleteCSPMAwsAccountMultiStatus) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-aws/entities/account/v1][%d] deleteCSPMAwsAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *DeleteCSPMAwsAccountMultiStatus) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-aws/entities/account/v1][%d] deleteCSPMAwsAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *DeleteCSPMAwsAccountMultiStatus) GetPayload() *models.MsaBaseEntitiesResponse {
	return o.Payload
}

func (o *DeleteCSPMAwsAccountMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaBaseEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCSPMAwsAccountBadRequest creates a DeleteCSPMAwsAccountBadRequest with default headers values
func NewDeleteCSPMAwsAccountBadRequest() *DeleteCSPMAwsAccountBadRequest {
	return &DeleteCSPMAwsAccountBadRequest{}
}

/*
DeleteCSPMAwsAccountBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteCSPMAwsAccountBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaBaseEntitiesResponse
}

// IsSuccess returns true when this delete c s p m aws account bad request response has a 2xx status code
func (o *DeleteCSPMAwsAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete c s p m aws account bad request response has a 3xx status code
func (o *DeleteCSPMAwsAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete c s p m aws account bad request response has a 4xx status code
func (o *DeleteCSPMAwsAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete c s p m aws account bad request response has a 5xx status code
func (o *DeleteCSPMAwsAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete c s p m aws account bad request response a status code equal to that given
func (o *DeleteCSPMAwsAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete c s p m aws account bad request response
func (o *DeleteCSPMAwsAccountBadRequest) Code() int {
	return 400
}

func (o *DeleteCSPMAwsAccountBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-aws/entities/account/v1][%d] deleteCSPMAwsAccountBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteCSPMAwsAccountBadRequest) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-aws/entities/account/v1][%d] deleteCSPMAwsAccountBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteCSPMAwsAccountBadRequest) GetPayload() *models.MsaBaseEntitiesResponse {
	return o.Payload
}

func (o *DeleteCSPMAwsAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaBaseEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCSPMAwsAccountForbidden creates a DeleteCSPMAwsAccountForbidden with default headers values
func NewDeleteCSPMAwsAccountForbidden() *DeleteCSPMAwsAccountForbidden {
	return &DeleteCSPMAwsAccountForbidden{}
}

/*
DeleteCSPMAwsAccountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteCSPMAwsAccountForbidden struct {

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

// IsSuccess returns true when this delete c s p m aws account forbidden response has a 2xx status code
func (o *DeleteCSPMAwsAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete c s p m aws account forbidden response has a 3xx status code
func (o *DeleteCSPMAwsAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete c s p m aws account forbidden response has a 4xx status code
func (o *DeleteCSPMAwsAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete c s p m aws account forbidden response has a 5xx status code
func (o *DeleteCSPMAwsAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete c s p m aws account forbidden response a status code equal to that given
func (o *DeleteCSPMAwsAccountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete c s p m aws account forbidden response
func (o *DeleteCSPMAwsAccountForbidden) Code() int {
	return 403
}

func (o *DeleteCSPMAwsAccountForbidden) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-aws/entities/account/v1][%d] deleteCSPMAwsAccountForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCSPMAwsAccountForbidden) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-aws/entities/account/v1][%d] deleteCSPMAwsAccountForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCSPMAwsAccountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteCSPMAwsAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteCSPMAwsAccountTooManyRequests creates a DeleteCSPMAwsAccountTooManyRequests with default headers values
func NewDeleteCSPMAwsAccountTooManyRequests() *DeleteCSPMAwsAccountTooManyRequests {
	return &DeleteCSPMAwsAccountTooManyRequests{}
}

/*
DeleteCSPMAwsAccountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteCSPMAwsAccountTooManyRequests struct {

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

// IsSuccess returns true when this delete c s p m aws account too many requests response has a 2xx status code
func (o *DeleteCSPMAwsAccountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete c s p m aws account too many requests response has a 3xx status code
func (o *DeleteCSPMAwsAccountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete c s p m aws account too many requests response has a 4xx status code
func (o *DeleteCSPMAwsAccountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete c s p m aws account too many requests response has a 5xx status code
func (o *DeleteCSPMAwsAccountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete c s p m aws account too many requests response a status code equal to that given
func (o *DeleteCSPMAwsAccountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the delete c s p m aws account too many requests response
func (o *DeleteCSPMAwsAccountTooManyRequests) Code() int {
	return 429
}

func (o *DeleteCSPMAwsAccountTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-aws/entities/account/v1][%d] deleteCSPMAwsAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteCSPMAwsAccountTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-aws/entities/account/v1][%d] deleteCSPMAwsAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteCSPMAwsAccountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteCSPMAwsAccountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteCSPMAwsAccountInternalServerError creates a DeleteCSPMAwsAccountInternalServerError with default headers values
func NewDeleteCSPMAwsAccountInternalServerError() *DeleteCSPMAwsAccountInternalServerError {
	return &DeleteCSPMAwsAccountInternalServerError{}
}

/*
DeleteCSPMAwsAccountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteCSPMAwsAccountInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaBaseEntitiesResponse
}

// IsSuccess returns true when this delete c s p m aws account internal server error response has a 2xx status code
func (o *DeleteCSPMAwsAccountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete c s p m aws account internal server error response has a 3xx status code
func (o *DeleteCSPMAwsAccountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete c s p m aws account internal server error response has a 4xx status code
func (o *DeleteCSPMAwsAccountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete c s p m aws account internal server error response has a 5xx status code
func (o *DeleteCSPMAwsAccountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete c s p m aws account internal server error response a status code equal to that given
func (o *DeleteCSPMAwsAccountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete c s p m aws account internal server error response
func (o *DeleteCSPMAwsAccountInternalServerError) Code() int {
	return 500
}

func (o *DeleteCSPMAwsAccountInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-aws/entities/account/v1][%d] deleteCSPMAwsAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCSPMAwsAccountInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /cloud-connect-cspm-aws/entities/account/v1][%d] deleteCSPMAwsAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCSPMAwsAccountInternalServerError) GetPayload() *models.MsaBaseEntitiesResponse {
	return o.Payload
}

func (o *DeleteCSPMAwsAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaBaseEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
