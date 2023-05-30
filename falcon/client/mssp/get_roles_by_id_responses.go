// Code generated by go-swagger; DO NOT EDIT.

package mssp

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

// GetRolesByIDReader is a Reader for the GetRolesByID structure.
type GetRolesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRolesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRolesByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewGetRolesByIDMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRolesByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRolesByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRolesByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRolesByIDOK creates a GetRolesByIDOK with default headers values
func NewGetRolesByIDOK() *GetRolesByIDOK {
	return &GetRolesByIDOK{}
}

/*
GetRolesByIDOK describes a response with status code 200, with default header values.

OK
*/
type GetRolesByIDOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainMSSPRoleResponseV1
}

// IsSuccess returns true when this get roles by Id o k response has a 2xx status code
func (o *GetRolesByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get roles by Id o k response has a 3xx status code
func (o *GetRolesByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles by Id o k response has a 4xx status code
func (o *GetRolesByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get roles by Id o k response has a 5xx status code
func (o *GetRolesByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles by Id o k response a status code equal to that given
func (o *GetRolesByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get roles by Id o k response
func (o *GetRolesByIDOK) Code() int {
	return 200
}

func (o *GetRolesByIDOK) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/mssp-roles/v1][%d] getRolesByIdOK  %+v", 200, o.Payload)
}

func (o *GetRolesByIDOK) String() string {
	return fmt.Sprintf("[GET /mssp/entities/mssp-roles/v1][%d] getRolesByIdOK  %+v", 200, o.Payload)
}

func (o *GetRolesByIDOK) GetPayload() *models.DomainMSSPRoleResponseV1 {
	return o.Payload
}

func (o *GetRolesByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainMSSPRoleResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesByIDMultiStatus creates a GetRolesByIDMultiStatus with default headers values
func NewGetRolesByIDMultiStatus() *GetRolesByIDMultiStatus {
	return &GetRolesByIDMultiStatus{}
}

/*
GetRolesByIDMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type GetRolesByIDMultiStatus struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainMSSPRoleResponseV1
}

// IsSuccess returns true when this get roles by Id multi status response has a 2xx status code
func (o *GetRolesByIDMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get roles by Id multi status response has a 3xx status code
func (o *GetRolesByIDMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles by Id multi status response has a 4xx status code
func (o *GetRolesByIDMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this get roles by Id multi status response has a 5xx status code
func (o *GetRolesByIDMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles by Id multi status response a status code equal to that given
func (o *GetRolesByIDMultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the get roles by Id multi status response
func (o *GetRolesByIDMultiStatus) Code() int {
	return 207
}

func (o *GetRolesByIDMultiStatus) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/mssp-roles/v1][%d] getRolesByIdMultiStatus  %+v", 207, o.Payload)
}

func (o *GetRolesByIDMultiStatus) String() string {
	return fmt.Sprintf("[GET /mssp/entities/mssp-roles/v1][%d] getRolesByIdMultiStatus  %+v", 207, o.Payload)
}

func (o *GetRolesByIDMultiStatus) GetPayload() *models.DomainMSSPRoleResponseV1 {
	return o.Payload
}

func (o *GetRolesByIDMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainMSSPRoleResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesByIDBadRequest creates a GetRolesByIDBadRequest with default headers values
func NewGetRolesByIDBadRequest() *GetRolesByIDBadRequest {
	return &GetRolesByIDBadRequest{}
}

/*
GetRolesByIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetRolesByIDBadRequest struct {

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

// IsSuccess returns true when this get roles by Id bad request response has a 2xx status code
func (o *GetRolesByIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles by Id bad request response has a 3xx status code
func (o *GetRolesByIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles by Id bad request response has a 4xx status code
func (o *GetRolesByIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get roles by Id bad request response has a 5xx status code
func (o *GetRolesByIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles by Id bad request response a status code equal to that given
func (o *GetRolesByIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get roles by Id bad request response
func (o *GetRolesByIDBadRequest) Code() int {
	return 400
}

func (o *GetRolesByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/mssp-roles/v1][%d] getRolesByIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetRolesByIDBadRequest) String() string {
	return fmt.Sprintf("[GET /mssp/entities/mssp-roles/v1][%d] getRolesByIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetRolesByIDBadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetRolesByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRolesByIDForbidden creates a GetRolesByIDForbidden with default headers values
func NewGetRolesByIDForbidden() *GetRolesByIDForbidden {
	return &GetRolesByIDForbidden{}
}

/*
GetRolesByIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRolesByIDForbidden struct {

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

// IsSuccess returns true when this get roles by Id forbidden response has a 2xx status code
func (o *GetRolesByIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles by Id forbidden response has a 3xx status code
func (o *GetRolesByIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles by Id forbidden response has a 4xx status code
func (o *GetRolesByIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get roles by Id forbidden response has a 5xx status code
func (o *GetRolesByIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles by Id forbidden response a status code equal to that given
func (o *GetRolesByIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get roles by Id forbidden response
func (o *GetRolesByIDForbidden) Code() int {
	return 403
}

func (o *GetRolesByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/mssp-roles/v1][%d] getRolesByIdForbidden  %+v", 403, o.Payload)
}

func (o *GetRolesByIDForbidden) String() string {
	return fmt.Sprintf("[GET /mssp/entities/mssp-roles/v1][%d] getRolesByIdForbidden  %+v", 403, o.Payload)
}

func (o *GetRolesByIDForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetRolesByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRolesByIDTooManyRequests creates a GetRolesByIDTooManyRequests with default headers values
func NewGetRolesByIDTooManyRequests() *GetRolesByIDTooManyRequests {
	return &GetRolesByIDTooManyRequests{}
}

/*
GetRolesByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetRolesByIDTooManyRequests struct {

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

// IsSuccess returns true when this get roles by Id too many requests response has a 2xx status code
func (o *GetRolesByIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles by Id too many requests response has a 3xx status code
func (o *GetRolesByIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles by Id too many requests response has a 4xx status code
func (o *GetRolesByIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get roles by Id too many requests response has a 5xx status code
func (o *GetRolesByIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles by Id too many requests response a status code equal to that given
func (o *GetRolesByIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get roles by Id too many requests response
func (o *GetRolesByIDTooManyRequests) Code() int {
	return 429
}

func (o *GetRolesByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/mssp-roles/v1][%d] getRolesByIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRolesByIDTooManyRequests) String() string {
	return fmt.Sprintf("[GET /mssp/entities/mssp-roles/v1][%d] getRolesByIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRolesByIDTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetRolesByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
