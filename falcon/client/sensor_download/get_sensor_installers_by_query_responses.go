// Code generated by go-swagger; DO NOT EDIT.

package sensor_download

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

// GetSensorInstallersByQueryReader is a Reader for the GetSensorInstallersByQuery structure.
type GetSensorInstallersByQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSensorInstallersByQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSensorInstallersByQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSensorInstallersByQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSensorInstallersByQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSensorInstallersByQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSensorInstallersByQueryOK creates a GetSensorInstallersByQueryOK with default headers values
func NewGetSensorInstallersByQueryOK() *GetSensorInstallersByQueryOK {
	return &GetSensorInstallersByQueryOK{}
}

/*
GetSensorInstallersByQueryOK describes a response with status code 200, with default header values.

OK
*/
type GetSensorInstallersByQueryOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecQueryResponse
}

// IsSuccess returns true when this get sensor installers by query o k response has a 2xx status code
func (o *GetSensorInstallersByQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get sensor installers by query o k response has a 3xx status code
func (o *GetSensorInstallersByQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sensor installers by query o k response has a 4xx status code
func (o *GetSensorInstallersByQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sensor installers by query o k response has a 5xx status code
func (o *GetSensorInstallersByQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get sensor installers by query o k response a status code equal to that given
func (o *GetSensorInstallersByQueryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get sensor installers by query o k response
func (o *GetSensorInstallersByQueryOK) Code() int {
	return 200
}

func (o *GetSensorInstallersByQueryOK) Error() string {
	return fmt.Sprintf("[GET /sensors/queries/installers/v1][%d] getSensorInstallersByQueryOK  %+v", 200, o.Payload)
}

func (o *GetSensorInstallersByQueryOK) String() string {
	return fmt.Sprintf("[GET /sensors/queries/installers/v1][%d] getSensorInstallersByQueryOK  %+v", 200, o.Payload)
}

func (o *GetSensorInstallersByQueryOK) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *GetSensorInstallersByQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSensorInstallersByQueryBadRequest creates a GetSensorInstallersByQueryBadRequest with default headers values
func NewGetSensorInstallersByQueryBadRequest() *GetSensorInstallersByQueryBadRequest {
	return &GetSensorInstallersByQueryBadRequest{}
}

/*
GetSensorInstallersByQueryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetSensorInstallersByQueryBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecQueryResponse
}

// IsSuccess returns true when this get sensor installers by query bad request response has a 2xx status code
func (o *GetSensorInstallersByQueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sensor installers by query bad request response has a 3xx status code
func (o *GetSensorInstallersByQueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sensor installers by query bad request response has a 4xx status code
func (o *GetSensorInstallersByQueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sensor installers by query bad request response has a 5xx status code
func (o *GetSensorInstallersByQueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get sensor installers by query bad request response a status code equal to that given
func (o *GetSensorInstallersByQueryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get sensor installers by query bad request response
func (o *GetSensorInstallersByQueryBadRequest) Code() int {
	return 400
}

func (o *GetSensorInstallersByQueryBadRequest) Error() string {
	return fmt.Sprintf("[GET /sensors/queries/installers/v1][%d] getSensorInstallersByQueryBadRequest  %+v", 400, o.Payload)
}

func (o *GetSensorInstallersByQueryBadRequest) String() string {
	return fmt.Sprintf("[GET /sensors/queries/installers/v1][%d] getSensorInstallersByQueryBadRequest  %+v", 400, o.Payload)
}

func (o *GetSensorInstallersByQueryBadRequest) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *GetSensorInstallersByQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSensorInstallersByQueryForbidden creates a GetSensorInstallersByQueryForbidden with default headers values
func NewGetSensorInstallersByQueryForbidden() *GetSensorInstallersByQueryForbidden {
	return &GetSensorInstallersByQueryForbidden{}
}

/*
GetSensorInstallersByQueryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSensorInstallersByQueryForbidden struct {

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

// IsSuccess returns true when this get sensor installers by query forbidden response has a 2xx status code
func (o *GetSensorInstallersByQueryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sensor installers by query forbidden response has a 3xx status code
func (o *GetSensorInstallersByQueryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sensor installers by query forbidden response has a 4xx status code
func (o *GetSensorInstallersByQueryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sensor installers by query forbidden response has a 5xx status code
func (o *GetSensorInstallersByQueryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get sensor installers by query forbidden response a status code equal to that given
func (o *GetSensorInstallersByQueryForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get sensor installers by query forbidden response
func (o *GetSensorInstallersByQueryForbidden) Code() int {
	return 403
}

func (o *GetSensorInstallersByQueryForbidden) Error() string {
	return fmt.Sprintf("[GET /sensors/queries/installers/v1][%d] getSensorInstallersByQueryForbidden  %+v", 403, o.Payload)
}

func (o *GetSensorInstallersByQueryForbidden) String() string {
	return fmt.Sprintf("[GET /sensors/queries/installers/v1][%d] getSensorInstallersByQueryForbidden  %+v", 403, o.Payload)
}

func (o *GetSensorInstallersByQueryForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetSensorInstallersByQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSensorInstallersByQueryTooManyRequests creates a GetSensorInstallersByQueryTooManyRequests with default headers values
func NewGetSensorInstallersByQueryTooManyRequests() *GetSensorInstallersByQueryTooManyRequests {
	return &GetSensorInstallersByQueryTooManyRequests{}
}

/*
GetSensorInstallersByQueryTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetSensorInstallersByQueryTooManyRequests struct {

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

// IsSuccess returns true when this get sensor installers by query too many requests response has a 2xx status code
func (o *GetSensorInstallersByQueryTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sensor installers by query too many requests response has a 3xx status code
func (o *GetSensorInstallersByQueryTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sensor installers by query too many requests response has a 4xx status code
func (o *GetSensorInstallersByQueryTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sensor installers by query too many requests response has a 5xx status code
func (o *GetSensorInstallersByQueryTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get sensor installers by query too many requests response a status code equal to that given
func (o *GetSensorInstallersByQueryTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get sensor installers by query too many requests response
func (o *GetSensorInstallersByQueryTooManyRequests) Code() int {
	return 429
}

func (o *GetSensorInstallersByQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /sensors/queries/installers/v1][%d] getSensorInstallersByQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSensorInstallersByQueryTooManyRequests) String() string {
	return fmt.Sprintf("[GET /sensors/queries/installers/v1][%d] getSensorInstallersByQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSensorInstallersByQueryTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetSensorInstallersByQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
