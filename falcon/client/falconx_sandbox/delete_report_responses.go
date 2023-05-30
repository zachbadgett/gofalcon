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

// DeleteReportReader is a Reader for the DeleteReport structure.
type DeleteReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteReportAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteReportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteReportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteReportTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteReportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteReportAccepted creates a DeleteReportAccepted with default headers values
func NewDeleteReportAccepted() *DeleteReportAccepted {
	return &DeleteReportAccepted{}
}

/*
DeleteReportAccepted describes a response with status code 202, with default header values.

Accepted
*/
type DeleteReportAccepted struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FalconxQueryResponse
}

// IsSuccess returns true when this delete report accepted response has a 2xx status code
func (o *DeleteReportAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete report accepted response has a 3xx status code
func (o *DeleteReportAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete report accepted response has a 4xx status code
func (o *DeleteReportAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete report accepted response has a 5xx status code
func (o *DeleteReportAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this delete report accepted response a status code equal to that given
func (o *DeleteReportAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the delete report accepted response
func (o *DeleteReportAccepted) Code() int {
	return 202
}

func (o *DeleteReportAccepted) Error() string {
	return fmt.Sprintf("[DELETE /falconx/entities/reports/v1][%d] deleteReportAccepted  %+v", 202, o.Payload)
}

func (o *DeleteReportAccepted) String() string {
	return fmt.Sprintf("[DELETE /falconx/entities/reports/v1][%d] deleteReportAccepted  %+v", 202, o.Payload)
}

func (o *DeleteReportAccepted) GetPayload() *models.FalconxQueryResponse {
	return o.Payload
}

func (o *DeleteReportAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FalconxQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteReportBadRequest creates a DeleteReportBadRequest with default headers values
func NewDeleteReportBadRequest() *DeleteReportBadRequest {
	return &DeleteReportBadRequest{}
}

/*
DeleteReportBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteReportBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FalconxErrorsOnly
}

// IsSuccess returns true when this delete report bad request response has a 2xx status code
func (o *DeleteReportBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete report bad request response has a 3xx status code
func (o *DeleteReportBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete report bad request response has a 4xx status code
func (o *DeleteReportBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete report bad request response has a 5xx status code
func (o *DeleteReportBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete report bad request response a status code equal to that given
func (o *DeleteReportBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete report bad request response
func (o *DeleteReportBadRequest) Code() int {
	return 400
}

func (o *DeleteReportBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /falconx/entities/reports/v1][%d] deleteReportBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteReportBadRequest) String() string {
	return fmt.Sprintf("[DELETE /falconx/entities/reports/v1][%d] deleteReportBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteReportBadRequest) GetPayload() *models.FalconxErrorsOnly {
	return o.Payload
}

func (o *DeleteReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FalconxErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteReportForbidden creates a DeleteReportForbidden with default headers values
func NewDeleteReportForbidden() *DeleteReportForbidden {
	return &DeleteReportForbidden{}
}

/*
DeleteReportForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteReportForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FalconxErrorsOnly
}

// IsSuccess returns true when this delete report forbidden response has a 2xx status code
func (o *DeleteReportForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete report forbidden response has a 3xx status code
func (o *DeleteReportForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete report forbidden response has a 4xx status code
func (o *DeleteReportForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete report forbidden response has a 5xx status code
func (o *DeleteReportForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete report forbidden response a status code equal to that given
func (o *DeleteReportForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete report forbidden response
func (o *DeleteReportForbidden) Code() int {
	return 403
}

func (o *DeleteReportForbidden) Error() string {
	return fmt.Sprintf("[DELETE /falconx/entities/reports/v1][%d] deleteReportForbidden  %+v", 403, o.Payload)
}

func (o *DeleteReportForbidden) String() string {
	return fmt.Sprintf("[DELETE /falconx/entities/reports/v1][%d] deleteReportForbidden  %+v", 403, o.Payload)
}

func (o *DeleteReportForbidden) GetPayload() *models.FalconxErrorsOnly {
	return o.Payload
}

func (o *DeleteReportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FalconxErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteReportNotFound creates a DeleteReportNotFound with default headers values
func NewDeleteReportNotFound() *DeleteReportNotFound {
	return &DeleteReportNotFound{}
}

/*
DeleteReportNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteReportNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FalconxErrorsOnly
}

// IsSuccess returns true when this delete report not found response has a 2xx status code
func (o *DeleteReportNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete report not found response has a 3xx status code
func (o *DeleteReportNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete report not found response has a 4xx status code
func (o *DeleteReportNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete report not found response has a 5xx status code
func (o *DeleteReportNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete report not found response a status code equal to that given
func (o *DeleteReportNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete report not found response
func (o *DeleteReportNotFound) Code() int {
	return 404
}

func (o *DeleteReportNotFound) Error() string {
	return fmt.Sprintf("[DELETE /falconx/entities/reports/v1][%d] deleteReportNotFound  %+v", 404, o.Payload)
}

func (o *DeleteReportNotFound) String() string {
	return fmt.Sprintf("[DELETE /falconx/entities/reports/v1][%d] deleteReportNotFound  %+v", 404, o.Payload)
}

func (o *DeleteReportNotFound) GetPayload() *models.FalconxErrorsOnly {
	return o.Payload
}

func (o *DeleteReportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FalconxErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteReportTooManyRequests creates a DeleteReportTooManyRequests with default headers values
func NewDeleteReportTooManyRequests() *DeleteReportTooManyRequests {
	return &DeleteReportTooManyRequests{}
}

/*
DeleteReportTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteReportTooManyRequests struct {

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

// IsSuccess returns true when this delete report too many requests response has a 2xx status code
func (o *DeleteReportTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete report too many requests response has a 3xx status code
func (o *DeleteReportTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete report too many requests response has a 4xx status code
func (o *DeleteReportTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete report too many requests response has a 5xx status code
func (o *DeleteReportTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete report too many requests response a status code equal to that given
func (o *DeleteReportTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the delete report too many requests response
func (o *DeleteReportTooManyRequests) Code() int {
	return 429
}

func (o *DeleteReportTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /falconx/entities/reports/v1][%d] deleteReportTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteReportTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /falconx/entities/reports/v1][%d] deleteReportTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteReportTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteReportTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteReportInternalServerError creates a DeleteReportInternalServerError with default headers values
func NewDeleteReportInternalServerError() *DeleteReportInternalServerError {
	return &DeleteReportInternalServerError{}
}

/*
DeleteReportInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteReportInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FalconxErrorsOnly
}

// IsSuccess returns true when this delete report internal server error response has a 2xx status code
func (o *DeleteReportInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete report internal server error response has a 3xx status code
func (o *DeleteReportInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete report internal server error response has a 4xx status code
func (o *DeleteReportInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete report internal server error response has a 5xx status code
func (o *DeleteReportInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete report internal server error response a status code equal to that given
func (o *DeleteReportInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete report internal server error response
func (o *DeleteReportInternalServerError) Code() int {
	return 500
}

func (o *DeleteReportInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /falconx/entities/reports/v1][%d] deleteReportInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteReportInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /falconx/entities/reports/v1][%d] deleteReportInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteReportInternalServerError) GetPayload() *models.FalconxErrorsOnly {
	return o.Payload
}

func (o *DeleteReportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FalconxErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
