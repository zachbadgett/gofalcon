// Code generated by go-swagger; DO NOT EDIT.

package sensor_update_policies

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

// QuerySensorUpdatePolicyMembersReader is a Reader for the QuerySensorUpdatePolicyMembers structure.
type QuerySensorUpdatePolicyMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuerySensorUpdatePolicyMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQuerySensorUpdatePolicyMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQuerySensorUpdatePolicyMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQuerySensorUpdatePolicyMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQuerySensorUpdatePolicyMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQuerySensorUpdatePolicyMembersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQuerySensorUpdatePolicyMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQuerySensorUpdatePolicyMembersOK creates a QuerySensorUpdatePolicyMembersOK with default headers values
func NewQuerySensorUpdatePolicyMembersOK() *QuerySensorUpdatePolicyMembersOK {
	return &QuerySensorUpdatePolicyMembersOK{}
}

/*
QuerySensorUpdatePolicyMembersOK describes a response with status code 200, with default header values.

OK
*/
type QuerySensorUpdatePolicyMembersOK struct {

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

// IsSuccess returns true when this query sensor update policy members o k response has a 2xx status code
func (o *QuerySensorUpdatePolicyMembersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query sensor update policy members o k response has a 3xx status code
func (o *QuerySensorUpdatePolicyMembersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query sensor update policy members o k response has a 4xx status code
func (o *QuerySensorUpdatePolicyMembersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query sensor update policy members o k response has a 5xx status code
func (o *QuerySensorUpdatePolicyMembersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query sensor update policy members o k response a status code equal to that given
func (o *QuerySensorUpdatePolicyMembersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query sensor update policy members o k response
func (o *QuerySensorUpdatePolicyMembersOK) Code() int {
	return 200
}

func (o *QuerySensorUpdatePolicyMembersOK) Error() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update-members/v1][%d] querySensorUpdatePolicyMembersOK  %+v", 200, o.Payload)
}

func (o *QuerySensorUpdatePolicyMembersOK) String() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update-members/v1][%d] querySensorUpdatePolicyMembersOK  %+v", 200, o.Payload)
}

func (o *QuerySensorUpdatePolicyMembersOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QuerySensorUpdatePolicyMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQuerySensorUpdatePolicyMembersBadRequest creates a QuerySensorUpdatePolicyMembersBadRequest with default headers values
func NewQuerySensorUpdatePolicyMembersBadRequest() *QuerySensorUpdatePolicyMembersBadRequest {
	return &QuerySensorUpdatePolicyMembersBadRequest{}
}

/*
QuerySensorUpdatePolicyMembersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QuerySensorUpdatePolicyMembersBadRequest struct {

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

// IsSuccess returns true when this query sensor update policy members bad request response has a 2xx status code
func (o *QuerySensorUpdatePolicyMembersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query sensor update policy members bad request response has a 3xx status code
func (o *QuerySensorUpdatePolicyMembersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query sensor update policy members bad request response has a 4xx status code
func (o *QuerySensorUpdatePolicyMembersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query sensor update policy members bad request response has a 5xx status code
func (o *QuerySensorUpdatePolicyMembersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query sensor update policy members bad request response a status code equal to that given
func (o *QuerySensorUpdatePolicyMembersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query sensor update policy members bad request response
func (o *QuerySensorUpdatePolicyMembersBadRequest) Code() int {
	return 400
}

func (o *QuerySensorUpdatePolicyMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update-members/v1][%d] querySensorUpdatePolicyMembersBadRequest  %+v", 400, o.Payload)
}

func (o *QuerySensorUpdatePolicyMembersBadRequest) String() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update-members/v1][%d] querySensorUpdatePolicyMembersBadRequest  %+v", 400, o.Payload)
}

func (o *QuerySensorUpdatePolicyMembersBadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QuerySensorUpdatePolicyMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQuerySensorUpdatePolicyMembersForbidden creates a QuerySensorUpdatePolicyMembersForbidden with default headers values
func NewQuerySensorUpdatePolicyMembersForbidden() *QuerySensorUpdatePolicyMembersForbidden {
	return &QuerySensorUpdatePolicyMembersForbidden{}
}

/*
QuerySensorUpdatePolicyMembersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QuerySensorUpdatePolicyMembersForbidden struct {

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

// IsSuccess returns true when this query sensor update policy members forbidden response has a 2xx status code
func (o *QuerySensorUpdatePolicyMembersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query sensor update policy members forbidden response has a 3xx status code
func (o *QuerySensorUpdatePolicyMembersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query sensor update policy members forbidden response has a 4xx status code
func (o *QuerySensorUpdatePolicyMembersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query sensor update policy members forbidden response has a 5xx status code
func (o *QuerySensorUpdatePolicyMembersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query sensor update policy members forbidden response a status code equal to that given
func (o *QuerySensorUpdatePolicyMembersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query sensor update policy members forbidden response
func (o *QuerySensorUpdatePolicyMembersForbidden) Code() int {
	return 403
}

func (o *QuerySensorUpdatePolicyMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update-members/v1][%d] querySensorUpdatePolicyMembersForbidden  %+v", 403, o.Payload)
}

func (o *QuerySensorUpdatePolicyMembersForbidden) String() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update-members/v1][%d] querySensorUpdatePolicyMembersForbidden  %+v", 403, o.Payload)
}

func (o *QuerySensorUpdatePolicyMembersForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QuerySensorUpdatePolicyMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQuerySensorUpdatePolicyMembersNotFound creates a QuerySensorUpdatePolicyMembersNotFound with default headers values
func NewQuerySensorUpdatePolicyMembersNotFound() *QuerySensorUpdatePolicyMembersNotFound {
	return &QuerySensorUpdatePolicyMembersNotFound{}
}

/*
QuerySensorUpdatePolicyMembersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type QuerySensorUpdatePolicyMembersNotFound struct {

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

// IsSuccess returns true when this query sensor update policy members not found response has a 2xx status code
func (o *QuerySensorUpdatePolicyMembersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query sensor update policy members not found response has a 3xx status code
func (o *QuerySensorUpdatePolicyMembersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query sensor update policy members not found response has a 4xx status code
func (o *QuerySensorUpdatePolicyMembersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this query sensor update policy members not found response has a 5xx status code
func (o *QuerySensorUpdatePolicyMembersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this query sensor update policy members not found response a status code equal to that given
func (o *QuerySensorUpdatePolicyMembersNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the query sensor update policy members not found response
func (o *QuerySensorUpdatePolicyMembersNotFound) Code() int {
	return 404
}

func (o *QuerySensorUpdatePolicyMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update-members/v1][%d] querySensorUpdatePolicyMembersNotFound  %+v", 404, o.Payload)
}

func (o *QuerySensorUpdatePolicyMembersNotFound) String() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update-members/v1][%d] querySensorUpdatePolicyMembersNotFound  %+v", 404, o.Payload)
}

func (o *QuerySensorUpdatePolicyMembersNotFound) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QuerySensorUpdatePolicyMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQuerySensorUpdatePolicyMembersTooManyRequests creates a QuerySensorUpdatePolicyMembersTooManyRequests with default headers values
func NewQuerySensorUpdatePolicyMembersTooManyRequests() *QuerySensorUpdatePolicyMembersTooManyRequests {
	return &QuerySensorUpdatePolicyMembersTooManyRequests{}
}

/*
QuerySensorUpdatePolicyMembersTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QuerySensorUpdatePolicyMembersTooManyRequests struct {

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

// IsSuccess returns true when this query sensor update policy members too many requests response has a 2xx status code
func (o *QuerySensorUpdatePolicyMembersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query sensor update policy members too many requests response has a 3xx status code
func (o *QuerySensorUpdatePolicyMembersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query sensor update policy members too many requests response has a 4xx status code
func (o *QuerySensorUpdatePolicyMembersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query sensor update policy members too many requests response has a 5xx status code
func (o *QuerySensorUpdatePolicyMembersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query sensor update policy members too many requests response a status code equal to that given
func (o *QuerySensorUpdatePolicyMembersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query sensor update policy members too many requests response
func (o *QuerySensorUpdatePolicyMembersTooManyRequests) Code() int {
	return 429
}

func (o *QuerySensorUpdatePolicyMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update-members/v1][%d] querySensorUpdatePolicyMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *QuerySensorUpdatePolicyMembersTooManyRequests) String() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update-members/v1][%d] querySensorUpdatePolicyMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *QuerySensorUpdatePolicyMembersTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QuerySensorUpdatePolicyMembersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQuerySensorUpdatePolicyMembersInternalServerError creates a QuerySensorUpdatePolicyMembersInternalServerError with default headers values
func NewQuerySensorUpdatePolicyMembersInternalServerError() *QuerySensorUpdatePolicyMembersInternalServerError {
	return &QuerySensorUpdatePolicyMembersInternalServerError{}
}

/*
QuerySensorUpdatePolicyMembersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QuerySensorUpdatePolicyMembersInternalServerError struct {

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

// IsSuccess returns true when this query sensor update policy members internal server error response has a 2xx status code
func (o *QuerySensorUpdatePolicyMembersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query sensor update policy members internal server error response has a 3xx status code
func (o *QuerySensorUpdatePolicyMembersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query sensor update policy members internal server error response has a 4xx status code
func (o *QuerySensorUpdatePolicyMembersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query sensor update policy members internal server error response has a 5xx status code
func (o *QuerySensorUpdatePolicyMembersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query sensor update policy members internal server error response a status code equal to that given
func (o *QuerySensorUpdatePolicyMembersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query sensor update policy members internal server error response
func (o *QuerySensorUpdatePolicyMembersInternalServerError) Code() int {
	return 500
}

func (o *QuerySensorUpdatePolicyMembersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update-members/v1][%d] querySensorUpdatePolicyMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *QuerySensorUpdatePolicyMembersInternalServerError) String() string {
	return fmt.Sprintf("[GET /policy/queries/sensor-update-members/v1][%d] querySensorUpdatePolicyMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *QuerySensorUpdatePolicyMembersInternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QuerySensorUpdatePolicyMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
