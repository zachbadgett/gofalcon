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

// GetCIDGroupMembersByReader is a Reader for the GetCIDGroupMembersBy structure.
type GetCIDGroupMembersByReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCIDGroupMembersByReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCIDGroupMembersByOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewGetCIDGroupMembersByMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCIDGroupMembersByBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCIDGroupMembersByForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCIDGroupMembersByTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCIDGroupMembersByOK creates a GetCIDGroupMembersByOK with default headers values
func NewGetCIDGroupMembersByOK() *GetCIDGroupMembersByOK {
	return &GetCIDGroupMembersByOK{}
}

/*
GetCIDGroupMembersByOK describes a response with status code 200, with default header values.

OK
*/
type GetCIDGroupMembersByOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainCIDGroupMembersResponseV1
}

// IsSuccess returns true when this get c Id group members by o k response has a 2xx status code
func (o *GetCIDGroupMembersByOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get c Id group members by o k response has a 3xx status code
func (o *GetCIDGroupMembersByOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c Id group members by o k response has a 4xx status code
func (o *GetCIDGroupMembersByOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get c Id group members by o k response has a 5xx status code
func (o *GetCIDGroupMembersByOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get c Id group members by o k response a status code equal to that given
func (o *GetCIDGroupMembersByOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get c Id group members by o k response
func (o *GetCIDGroupMembersByOK) Code() int {
	return 200
}

func (o *GetCIDGroupMembersByOK) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v1][%d] getCIdGroupMembersByOK  %+v", 200, o.Payload)
}

func (o *GetCIDGroupMembersByOK) String() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v1][%d] getCIdGroupMembersByOK  %+v", 200, o.Payload)
}

func (o *GetCIDGroupMembersByOK) GetPayload() *models.DomainCIDGroupMembersResponseV1 {
	return o.Payload
}

func (o *GetCIDGroupMembersByOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainCIDGroupMembersResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCIDGroupMembersByMultiStatus creates a GetCIDGroupMembersByMultiStatus with default headers values
func NewGetCIDGroupMembersByMultiStatus() *GetCIDGroupMembersByMultiStatus {
	return &GetCIDGroupMembersByMultiStatus{}
}

/*
GetCIDGroupMembersByMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type GetCIDGroupMembersByMultiStatus struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainCIDGroupMembersResponseV1
}

// IsSuccess returns true when this get c Id group members by multi status response has a 2xx status code
func (o *GetCIDGroupMembersByMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get c Id group members by multi status response has a 3xx status code
func (o *GetCIDGroupMembersByMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c Id group members by multi status response has a 4xx status code
func (o *GetCIDGroupMembersByMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this get c Id group members by multi status response has a 5xx status code
func (o *GetCIDGroupMembersByMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this get c Id group members by multi status response a status code equal to that given
func (o *GetCIDGroupMembersByMultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the get c Id group members by multi status response
func (o *GetCIDGroupMembersByMultiStatus) Code() int {
	return 207
}

func (o *GetCIDGroupMembersByMultiStatus) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v1][%d] getCIdGroupMembersByMultiStatus  %+v", 207, o.Payload)
}

func (o *GetCIDGroupMembersByMultiStatus) String() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v1][%d] getCIdGroupMembersByMultiStatus  %+v", 207, o.Payload)
}

func (o *GetCIDGroupMembersByMultiStatus) GetPayload() *models.DomainCIDGroupMembersResponseV1 {
	return o.Payload
}

func (o *GetCIDGroupMembersByMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainCIDGroupMembersResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCIDGroupMembersByBadRequest creates a GetCIDGroupMembersByBadRequest with default headers values
func NewGetCIDGroupMembersByBadRequest() *GetCIDGroupMembersByBadRequest {
	return &GetCIDGroupMembersByBadRequest{}
}

/*
GetCIDGroupMembersByBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetCIDGroupMembersByBadRequest struct {

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

// IsSuccess returns true when this get c Id group members by bad request response has a 2xx status code
func (o *GetCIDGroupMembersByBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c Id group members by bad request response has a 3xx status code
func (o *GetCIDGroupMembersByBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c Id group members by bad request response has a 4xx status code
func (o *GetCIDGroupMembersByBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c Id group members by bad request response has a 5xx status code
func (o *GetCIDGroupMembersByBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get c Id group members by bad request response a status code equal to that given
func (o *GetCIDGroupMembersByBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get c Id group members by bad request response
func (o *GetCIDGroupMembersByBadRequest) Code() int {
	return 400
}

func (o *GetCIDGroupMembersByBadRequest) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v1][%d] getCIdGroupMembersByBadRequest  %+v", 400, o.Payload)
}

func (o *GetCIDGroupMembersByBadRequest) String() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v1][%d] getCIdGroupMembersByBadRequest  %+v", 400, o.Payload)
}

func (o *GetCIDGroupMembersByBadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetCIDGroupMembersByBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCIDGroupMembersByForbidden creates a GetCIDGroupMembersByForbidden with default headers values
func NewGetCIDGroupMembersByForbidden() *GetCIDGroupMembersByForbidden {
	return &GetCIDGroupMembersByForbidden{}
}

/*
GetCIDGroupMembersByForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCIDGroupMembersByForbidden struct {

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

// IsSuccess returns true when this get c Id group members by forbidden response has a 2xx status code
func (o *GetCIDGroupMembersByForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c Id group members by forbidden response has a 3xx status code
func (o *GetCIDGroupMembersByForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c Id group members by forbidden response has a 4xx status code
func (o *GetCIDGroupMembersByForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c Id group members by forbidden response has a 5xx status code
func (o *GetCIDGroupMembersByForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get c Id group members by forbidden response a status code equal to that given
func (o *GetCIDGroupMembersByForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get c Id group members by forbidden response
func (o *GetCIDGroupMembersByForbidden) Code() int {
	return 403
}

func (o *GetCIDGroupMembersByForbidden) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v1][%d] getCIdGroupMembersByForbidden  %+v", 403, o.Payload)
}

func (o *GetCIDGroupMembersByForbidden) String() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v1][%d] getCIdGroupMembersByForbidden  %+v", 403, o.Payload)
}

func (o *GetCIDGroupMembersByForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetCIDGroupMembersByForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCIDGroupMembersByTooManyRequests creates a GetCIDGroupMembersByTooManyRequests with default headers values
func NewGetCIDGroupMembersByTooManyRequests() *GetCIDGroupMembersByTooManyRequests {
	return &GetCIDGroupMembersByTooManyRequests{}
}

/*
GetCIDGroupMembersByTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetCIDGroupMembersByTooManyRequests struct {

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

// IsSuccess returns true when this get c Id group members by too many requests response has a 2xx status code
func (o *GetCIDGroupMembersByTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c Id group members by too many requests response has a 3xx status code
func (o *GetCIDGroupMembersByTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c Id group members by too many requests response has a 4xx status code
func (o *GetCIDGroupMembersByTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c Id group members by too many requests response has a 5xx status code
func (o *GetCIDGroupMembersByTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get c Id group members by too many requests response a status code equal to that given
func (o *GetCIDGroupMembersByTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get c Id group members by too many requests response
func (o *GetCIDGroupMembersByTooManyRequests) Code() int {
	return 429
}

func (o *GetCIDGroupMembersByTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v1][%d] getCIdGroupMembersByTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCIDGroupMembersByTooManyRequests) String() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v1][%d] getCIdGroupMembersByTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCIDGroupMembersByTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCIDGroupMembersByTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
