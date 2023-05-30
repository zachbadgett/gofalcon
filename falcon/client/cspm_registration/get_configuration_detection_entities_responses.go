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

// GetConfigurationDetectionEntitiesReader is a Reader for the GetConfigurationDetectionEntities structure.
type GetConfigurationDetectionEntitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigurationDetectionEntitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConfigurationDetectionEntitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConfigurationDetectionEntitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConfigurationDetectionEntitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConfigurationDetectionEntitiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConfigurationDetectionEntitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConfigurationDetectionEntitiesOK creates a GetConfigurationDetectionEntitiesOK with default headers values
func NewGetConfigurationDetectionEntitiesOK() *GetConfigurationDetectionEntitiesOK {
	return &GetConfigurationDetectionEntitiesOK{}
}

/*
GetConfigurationDetectionEntitiesOK describes a response with status code 200, with default header values.

OK
*/
type GetConfigurationDetectionEntitiesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationExternalIOMEventResponseV2
}

// IsSuccess returns true when this get configuration detection entities o k response has a 2xx status code
func (o *GetConfigurationDetectionEntitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get configuration detection entities o k response has a 3xx status code
func (o *GetConfigurationDetectionEntitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration detection entities o k response has a 4xx status code
func (o *GetConfigurationDetectionEntitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get configuration detection entities o k response has a 5xx status code
func (o *GetConfigurationDetectionEntitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration detection entities o k response a status code equal to that given
func (o *GetConfigurationDetectionEntitiesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get configuration detection entities o k response
func (o *GetConfigurationDetectionEntitiesOK) Code() int {
	return 200
}

func (o *GetConfigurationDetectionEntitiesOK) Error() string {
	return fmt.Sprintf("[GET /detects/entities/iom/v2][%d] getConfigurationDetectionEntitiesOK  %+v", 200, o.Payload)
}

func (o *GetConfigurationDetectionEntitiesOK) String() string {
	return fmt.Sprintf("[GET /detects/entities/iom/v2][%d] getConfigurationDetectionEntitiesOK  %+v", 200, o.Payload)
}

func (o *GetConfigurationDetectionEntitiesOK) GetPayload() *models.RegistrationExternalIOMEventResponseV2 {
	return o.Payload
}

func (o *GetConfigurationDetectionEntitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationExternalIOMEventResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationDetectionEntitiesBadRequest creates a GetConfigurationDetectionEntitiesBadRequest with default headers values
func NewGetConfigurationDetectionEntitiesBadRequest() *GetConfigurationDetectionEntitiesBadRequest {
	return &GetConfigurationDetectionEntitiesBadRequest{}
}

/*
GetConfigurationDetectionEntitiesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetConfigurationDetectionEntitiesBadRequest struct {

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

// IsSuccess returns true when this get configuration detection entities bad request response has a 2xx status code
func (o *GetConfigurationDetectionEntitiesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get configuration detection entities bad request response has a 3xx status code
func (o *GetConfigurationDetectionEntitiesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration detection entities bad request response has a 4xx status code
func (o *GetConfigurationDetectionEntitiesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get configuration detection entities bad request response has a 5xx status code
func (o *GetConfigurationDetectionEntitiesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration detection entities bad request response a status code equal to that given
func (o *GetConfigurationDetectionEntitiesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get configuration detection entities bad request response
func (o *GetConfigurationDetectionEntitiesBadRequest) Code() int {
	return 400
}

func (o *GetConfigurationDetectionEntitiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /detects/entities/iom/v2][%d] getConfigurationDetectionEntitiesBadRequest  %+v", 400, o.Payload)
}

func (o *GetConfigurationDetectionEntitiesBadRequest) String() string {
	return fmt.Sprintf("[GET /detects/entities/iom/v2][%d] getConfigurationDetectionEntitiesBadRequest  %+v", 400, o.Payload)
}

func (o *GetConfigurationDetectionEntitiesBadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetConfigurationDetectionEntitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetConfigurationDetectionEntitiesForbidden creates a GetConfigurationDetectionEntitiesForbidden with default headers values
func NewGetConfigurationDetectionEntitiesForbidden() *GetConfigurationDetectionEntitiesForbidden {
	return &GetConfigurationDetectionEntitiesForbidden{}
}

/*
GetConfigurationDetectionEntitiesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetConfigurationDetectionEntitiesForbidden struct {

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

// IsSuccess returns true when this get configuration detection entities forbidden response has a 2xx status code
func (o *GetConfigurationDetectionEntitiesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get configuration detection entities forbidden response has a 3xx status code
func (o *GetConfigurationDetectionEntitiesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration detection entities forbidden response has a 4xx status code
func (o *GetConfigurationDetectionEntitiesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get configuration detection entities forbidden response has a 5xx status code
func (o *GetConfigurationDetectionEntitiesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration detection entities forbidden response a status code equal to that given
func (o *GetConfigurationDetectionEntitiesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get configuration detection entities forbidden response
func (o *GetConfigurationDetectionEntitiesForbidden) Code() int {
	return 403
}

func (o *GetConfigurationDetectionEntitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /detects/entities/iom/v2][%d] getConfigurationDetectionEntitiesForbidden  %+v", 403, o.Payload)
}

func (o *GetConfigurationDetectionEntitiesForbidden) String() string {
	return fmt.Sprintf("[GET /detects/entities/iom/v2][%d] getConfigurationDetectionEntitiesForbidden  %+v", 403, o.Payload)
}

func (o *GetConfigurationDetectionEntitiesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetConfigurationDetectionEntitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetConfigurationDetectionEntitiesTooManyRequests creates a GetConfigurationDetectionEntitiesTooManyRequests with default headers values
func NewGetConfigurationDetectionEntitiesTooManyRequests() *GetConfigurationDetectionEntitiesTooManyRequests {
	return &GetConfigurationDetectionEntitiesTooManyRequests{}
}

/*
GetConfigurationDetectionEntitiesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetConfigurationDetectionEntitiesTooManyRequests struct {

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

// IsSuccess returns true when this get configuration detection entities too many requests response has a 2xx status code
func (o *GetConfigurationDetectionEntitiesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get configuration detection entities too many requests response has a 3xx status code
func (o *GetConfigurationDetectionEntitiesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration detection entities too many requests response has a 4xx status code
func (o *GetConfigurationDetectionEntitiesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get configuration detection entities too many requests response has a 5xx status code
func (o *GetConfigurationDetectionEntitiesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration detection entities too many requests response a status code equal to that given
func (o *GetConfigurationDetectionEntitiesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get configuration detection entities too many requests response
func (o *GetConfigurationDetectionEntitiesTooManyRequests) Code() int {
	return 429
}

func (o *GetConfigurationDetectionEntitiesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /detects/entities/iom/v2][%d] getConfigurationDetectionEntitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConfigurationDetectionEntitiesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /detects/entities/iom/v2][%d] getConfigurationDetectionEntitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConfigurationDetectionEntitiesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetConfigurationDetectionEntitiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetConfigurationDetectionEntitiesInternalServerError creates a GetConfigurationDetectionEntitiesInternalServerError with default headers values
func NewGetConfigurationDetectionEntitiesInternalServerError() *GetConfigurationDetectionEntitiesInternalServerError {
	return &GetConfigurationDetectionEntitiesInternalServerError{}
}

/*
GetConfigurationDetectionEntitiesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetConfigurationDetectionEntitiesInternalServerError struct {

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

// IsSuccess returns true when this get configuration detection entities internal server error response has a 2xx status code
func (o *GetConfigurationDetectionEntitiesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get configuration detection entities internal server error response has a 3xx status code
func (o *GetConfigurationDetectionEntitiesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration detection entities internal server error response has a 4xx status code
func (o *GetConfigurationDetectionEntitiesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get configuration detection entities internal server error response has a 5xx status code
func (o *GetConfigurationDetectionEntitiesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get configuration detection entities internal server error response a status code equal to that given
func (o *GetConfigurationDetectionEntitiesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get configuration detection entities internal server error response
func (o *GetConfigurationDetectionEntitiesInternalServerError) Code() int {
	return 500
}

func (o *GetConfigurationDetectionEntitiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /detects/entities/iom/v2][%d] getConfigurationDetectionEntitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConfigurationDetectionEntitiesInternalServerError) String() string {
	return fmt.Sprintf("[GET /detects/entities/iom/v2][%d] getConfigurationDetectionEntitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConfigurationDetectionEntitiesInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetConfigurationDetectionEntitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
