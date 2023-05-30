// Code generated by go-swagger; DO NOT EDIT.

package falconx_sandbox

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

// GetSampleV2Reader is a Reader for the GetSampleV2 structure.
type GetSampleV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSampleV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSampleV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSampleV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSampleV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSampleV2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSampleV2TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSampleV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSampleV2OK creates a GetSampleV2OK with default headers values
func NewGetSampleV2OK() *GetSampleV2OK {
	return &GetSampleV2OK{}
}

/*
GetSampleV2OK describes a response with status code 200, with default header values.

OK
*/
type GetSampleV2OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload string
}

// IsSuccess returns true when this get sample v2 o k response has a 2xx status code
func (o *GetSampleV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get sample v2 o k response has a 3xx status code
func (o *GetSampleV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sample v2 o k response has a 4xx status code
func (o *GetSampleV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sample v2 o k response has a 5xx status code
func (o *GetSampleV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get sample v2 o k response a status code equal to that given
func (o *GetSampleV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get sample v2 o k response
func (o *GetSampleV2OK) Code() int {
	return 200
}

func (o *GetSampleV2OK) Error() string {
	return fmt.Sprintf("[GET /samples/entities/samples/v2][%d] getSampleV2OK  %+v", 200, o.Payload)
}

func (o *GetSampleV2OK) String() string {
	return fmt.Sprintf("[GET /samples/entities/samples/v2][%d] getSampleV2OK  %+v", 200, o.Payload)
}

func (o *GetSampleV2OK) GetPayload() string {
	return o.Payload
}

func (o *GetSampleV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSampleV2BadRequest creates a GetSampleV2BadRequest with default headers values
func NewGetSampleV2BadRequest() *GetSampleV2BadRequest {
	return &GetSampleV2BadRequest{}
}

/*
GetSampleV2BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetSampleV2BadRequest struct {

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

// IsSuccess returns true when this get sample v2 bad request response has a 2xx status code
func (o *GetSampleV2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sample v2 bad request response has a 3xx status code
func (o *GetSampleV2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sample v2 bad request response has a 4xx status code
func (o *GetSampleV2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sample v2 bad request response has a 5xx status code
func (o *GetSampleV2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get sample v2 bad request response a status code equal to that given
func (o *GetSampleV2BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get sample v2 bad request response
func (o *GetSampleV2BadRequest) Code() int {
	return 400
}

func (o *GetSampleV2BadRequest) Error() string {
	return fmt.Sprintf("[GET /samples/entities/samples/v2][%d] getSampleV2BadRequest  %+v", 400, o.Payload)
}

func (o *GetSampleV2BadRequest) String() string {
	return fmt.Sprintf("[GET /samples/entities/samples/v2][%d] getSampleV2BadRequest  %+v", 400, o.Payload)
}

func (o *GetSampleV2BadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetSampleV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSampleV2Forbidden creates a GetSampleV2Forbidden with default headers values
func NewGetSampleV2Forbidden() *GetSampleV2Forbidden {
	return &GetSampleV2Forbidden{}
}

/*
GetSampleV2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSampleV2Forbidden struct {

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

// IsSuccess returns true when this get sample v2 forbidden response has a 2xx status code
func (o *GetSampleV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sample v2 forbidden response has a 3xx status code
func (o *GetSampleV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sample v2 forbidden response has a 4xx status code
func (o *GetSampleV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sample v2 forbidden response has a 5xx status code
func (o *GetSampleV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get sample v2 forbidden response a status code equal to that given
func (o *GetSampleV2Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get sample v2 forbidden response
func (o *GetSampleV2Forbidden) Code() int {
	return 403
}

func (o *GetSampleV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /samples/entities/samples/v2][%d] getSampleV2Forbidden  %+v", 403, o.Payload)
}

func (o *GetSampleV2Forbidden) String() string {
	return fmt.Sprintf("[GET /samples/entities/samples/v2][%d] getSampleV2Forbidden  %+v", 403, o.Payload)
}

func (o *GetSampleV2Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetSampleV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSampleV2NotFound creates a GetSampleV2NotFound with default headers values
func NewGetSampleV2NotFound() *GetSampleV2NotFound {
	return &GetSampleV2NotFound{}
}

/*
GetSampleV2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetSampleV2NotFound struct {

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

// IsSuccess returns true when this get sample v2 not found response has a 2xx status code
func (o *GetSampleV2NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sample v2 not found response has a 3xx status code
func (o *GetSampleV2NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sample v2 not found response has a 4xx status code
func (o *GetSampleV2NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sample v2 not found response has a 5xx status code
func (o *GetSampleV2NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get sample v2 not found response a status code equal to that given
func (o *GetSampleV2NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get sample v2 not found response
func (o *GetSampleV2NotFound) Code() int {
	return 404
}

func (o *GetSampleV2NotFound) Error() string {
	return fmt.Sprintf("[GET /samples/entities/samples/v2][%d] getSampleV2NotFound  %+v", 404, o.Payload)
}

func (o *GetSampleV2NotFound) String() string {
	return fmt.Sprintf("[GET /samples/entities/samples/v2][%d] getSampleV2NotFound  %+v", 404, o.Payload)
}

func (o *GetSampleV2NotFound) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetSampleV2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSampleV2TooManyRequests creates a GetSampleV2TooManyRequests with default headers values
func NewGetSampleV2TooManyRequests() *GetSampleV2TooManyRequests {
	return &GetSampleV2TooManyRequests{}
}

/*
GetSampleV2TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetSampleV2TooManyRequests struct {

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

// IsSuccess returns true when this get sample v2 too many requests response has a 2xx status code
func (o *GetSampleV2TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sample v2 too many requests response has a 3xx status code
func (o *GetSampleV2TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sample v2 too many requests response has a 4xx status code
func (o *GetSampleV2TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sample v2 too many requests response has a 5xx status code
func (o *GetSampleV2TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get sample v2 too many requests response a status code equal to that given
func (o *GetSampleV2TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get sample v2 too many requests response
func (o *GetSampleV2TooManyRequests) Code() int {
	return 429
}

func (o *GetSampleV2TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /samples/entities/samples/v2][%d] getSampleV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSampleV2TooManyRequests) String() string {
	return fmt.Sprintf("[GET /samples/entities/samples/v2][%d] getSampleV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSampleV2TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetSampleV2TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSampleV2InternalServerError creates a GetSampleV2InternalServerError with default headers values
func NewGetSampleV2InternalServerError() *GetSampleV2InternalServerError {
	return &GetSampleV2InternalServerError{}
}

/*
GetSampleV2InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetSampleV2InternalServerError struct {

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

// IsSuccess returns true when this get sample v2 internal server error response has a 2xx status code
func (o *GetSampleV2InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sample v2 internal server error response has a 3xx status code
func (o *GetSampleV2InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sample v2 internal server error response has a 4xx status code
func (o *GetSampleV2InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sample v2 internal server error response has a 5xx status code
func (o *GetSampleV2InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get sample v2 internal server error response a status code equal to that given
func (o *GetSampleV2InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get sample v2 internal server error response
func (o *GetSampleV2InternalServerError) Code() int {
	return 500
}

func (o *GetSampleV2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /samples/entities/samples/v2][%d] getSampleV2InternalServerError  %+v", 500, o.Payload)
}

func (o *GetSampleV2InternalServerError) String() string {
	return fmt.Sprintf("[GET /samples/entities/samples/v2][%d] getSampleV2InternalServerError  %+v", 500, o.Payload)
}

func (o *GetSampleV2InternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetSampleV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
