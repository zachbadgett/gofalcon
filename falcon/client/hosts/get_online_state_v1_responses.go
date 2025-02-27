// Code generated by go-swagger; DO NOT EDIT.

package hosts

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

// GetOnlineStateV1Reader is a Reader for the GetOnlineStateV1 structure.
type GetOnlineStateV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOnlineStateV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOnlineStateV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetOnlineStateV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOnlineStateV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetOnlineStateV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOnlineStateV1OK creates a GetOnlineStateV1OK with default headers values
func NewGetOnlineStateV1OK() *GetOnlineStateV1OK {
	return &GetOnlineStateV1OK{}
}

/*
GetOnlineStateV1OK describes a response with status code 200, with default header values.

OK
*/
type GetOnlineStateV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.StateOnlineStateRespV1
}

// IsSuccess returns true when this get online state v1 o k response has a 2xx status code
func (o *GetOnlineStateV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get online state v1 o k response has a 3xx status code
func (o *GetOnlineStateV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get online state v1 o k response has a 4xx status code
func (o *GetOnlineStateV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get online state v1 o k response has a 5xx status code
func (o *GetOnlineStateV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get online state v1 o k response a status code equal to that given
func (o *GetOnlineStateV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get online state v1 o k response
func (o *GetOnlineStateV1OK) Code() int {
	return 200
}

func (o *GetOnlineStateV1OK) Error() string {
	return fmt.Sprintf("[GET /devices/entities/online-state/v1][%d] getOnlineStateV1OK  %+v", 200, o.Payload)
}

func (o *GetOnlineStateV1OK) String() string {
	return fmt.Sprintf("[GET /devices/entities/online-state/v1][%d] getOnlineStateV1OK  %+v", 200, o.Payload)
}

func (o *GetOnlineStateV1OK) GetPayload() *models.StateOnlineStateRespV1 {
	return o.Payload
}

func (o *GetOnlineStateV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.StateOnlineStateRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOnlineStateV1Forbidden creates a GetOnlineStateV1Forbidden with default headers values
func NewGetOnlineStateV1Forbidden() *GetOnlineStateV1Forbidden {
	return &GetOnlineStateV1Forbidden{}
}

/*
GetOnlineStateV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetOnlineStateV1Forbidden struct {

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

// IsSuccess returns true when this get online state v1 forbidden response has a 2xx status code
func (o *GetOnlineStateV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get online state v1 forbidden response has a 3xx status code
func (o *GetOnlineStateV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get online state v1 forbidden response has a 4xx status code
func (o *GetOnlineStateV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get online state v1 forbidden response has a 5xx status code
func (o *GetOnlineStateV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get online state v1 forbidden response a status code equal to that given
func (o *GetOnlineStateV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get online state v1 forbidden response
func (o *GetOnlineStateV1Forbidden) Code() int {
	return 403
}

func (o *GetOnlineStateV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /devices/entities/online-state/v1][%d] getOnlineStateV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetOnlineStateV1Forbidden) String() string {
	return fmt.Sprintf("[GET /devices/entities/online-state/v1][%d] getOnlineStateV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetOnlineStateV1Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetOnlineStateV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetOnlineStateV1TooManyRequests creates a GetOnlineStateV1TooManyRequests with default headers values
func NewGetOnlineStateV1TooManyRequests() *GetOnlineStateV1TooManyRequests {
	return &GetOnlineStateV1TooManyRequests{}
}

/*
GetOnlineStateV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetOnlineStateV1TooManyRequests struct {

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

// IsSuccess returns true when this get online state v1 too many requests response has a 2xx status code
func (o *GetOnlineStateV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get online state v1 too many requests response has a 3xx status code
func (o *GetOnlineStateV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get online state v1 too many requests response has a 4xx status code
func (o *GetOnlineStateV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get online state v1 too many requests response has a 5xx status code
func (o *GetOnlineStateV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get online state v1 too many requests response a status code equal to that given
func (o *GetOnlineStateV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get online state v1 too many requests response
func (o *GetOnlineStateV1TooManyRequests) Code() int {
	return 429
}

func (o *GetOnlineStateV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /devices/entities/online-state/v1][%d] getOnlineStateV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOnlineStateV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /devices/entities/online-state/v1][%d] getOnlineStateV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOnlineStateV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetOnlineStateV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetOnlineStateV1Default creates a GetOnlineStateV1Default with default headers values
func NewGetOnlineStateV1Default(code int) *GetOnlineStateV1Default {
	return &GetOnlineStateV1Default{
		_statusCode: code,
	}
}

/*
GetOnlineStateV1Default describes a response with status code -1, with default header values.

OK
*/
type GetOnlineStateV1Default struct {
	_statusCode int

	Payload *models.StateOnlineStateRespV1
}

// IsSuccess returns true when this get online state v1 default response has a 2xx status code
func (o *GetOnlineStateV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get online state v1 default response has a 3xx status code
func (o *GetOnlineStateV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get online state v1 default response has a 4xx status code
func (o *GetOnlineStateV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get online state v1 default response has a 5xx status code
func (o *GetOnlineStateV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get online state v1 default response a status code equal to that given
func (o *GetOnlineStateV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get online state v1 default response
func (o *GetOnlineStateV1Default) Code() int {
	return o._statusCode
}

func (o *GetOnlineStateV1Default) Error() string {
	return fmt.Sprintf("[GET /devices/entities/online-state/v1][%d] GetOnlineState.V1 default  %+v", o._statusCode, o.Payload)
}

func (o *GetOnlineStateV1Default) String() string {
	return fmt.Sprintf("[GET /devices/entities/online-state/v1][%d] GetOnlineState.V1 default  %+v", o._statusCode, o.Payload)
}

func (o *GetOnlineStateV1Default) GetPayload() *models.StateOnlineStateRespV1 {
	return o.Payload
}

func (o *GetOnlineStateV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StateOnlineStateRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
