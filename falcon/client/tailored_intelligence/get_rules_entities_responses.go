// Code generated by go-swagger; DO NOT EDIT.

package tailored_intelligence

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

// GetRulesEntitiesReader is a Reader for the GetRulesEntities structure.
type GetRulesEntitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRulesEntitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRulesEntitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRulesEntitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRulesEntitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRulesEntitiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRulesEntitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRulesEntitiesOK creates a GetRulesEntitiesOK with default headers values
func NewGetRulesEntitiesOK() *GetRulesEntitiesOK {
	return &GetRulesEntitiesOK{}
}

/*
GetRulesEntitiesOK describes a response with status code 200, with default header values.

OK
*/
type GetRulesEntitiesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainRuleEntitiesResponse
}

// IsSuccess returns true when this get rules entities o k response has a 2xx status code
func (o *GetRulesEntitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get rules entities o k response has a 3xx status code
func (o *GetRulesEntitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rules entities o k response has a 4xx status code
func (o *GetRulesEntitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get rules entities o k response has a 5xx status code
func (o *GetRulesEntitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get rules entities o k response a status code equal to that given
func (o *GetRulesEntitiesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get rules entities o k response
func (o *GetRulesEntitiesOK) Code() int {
	return 200
}

func (o *GetRulesEntitiesOK) Error() string {
	return fmt.Sprintf("[POST /ti/rules/entities/rules/GET/v2][%d] getRulesEntitiesOK  %+v", 200, o.Payload)
}

func (o *GetRulesEntitiesOK) String() string {
	return fmt.Sprintf("[POST /ti/rules/entities/rules/GET/v2][%d] getRulesEntitiesOK  %+v", 200, o.Payload)
}

func (o *GetRulesEntitiesOK) GetPayload() *models.DomainRuleEntitiesResponse {
	return o.Payload
}

func (o *GetRulesEntitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainRuleEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRulesEntitiesBadRequest creates a GetRulesEntitiesBadRequest with default headers values
func NewGetRulesEntitiesBadRequest() *GetRulesEntitiesBadRequest {
	return &GetRulesEntitiesBadRequest{}
}

/*
GetRulesEntitiesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetRulesEntitiesBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainRuleEntitiesResponse
}

// IsSuccess returns true when this get rules entities bad request response has a 2xx status code
func (o *GetRulesEntitiesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rules entities bad request response has a 3xx status code
func (o *GetRulesEntitiesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rules entities bad request response has a 4xx status code
func (o *GetRulesEntitiesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rules entities bad request response has a 5xx status code
func (o *GetRulesEntitiesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get rules entities bad request response a status code equal to that given
func (o *GetRulesEntitiesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get rules entities bad request response
func (o *GetRulesEntitiesBadRequest) Code() int {
	return 400
}

func (o *GetRulesEntitiesBadRequest) Error() string {
	return fmt.Sprintf("[POST /ti/rules/entities/rules/GET/v2][%d] getRulesEntitiesBadRequest  %+v", 400, o.Payload)
}

func (o *GetRulesEntitiesBadRequest) String() string {
	return fmt.Sprintf("[POST /ti/rules/entities/rules/GET/v2][%d] getRulesEntitiesBadRequest  %+v", 400, o.Payload)
}

func (o *GetRulesEntitiesBadRequest) GetPayload() *models.DomainRuleEntitiesResponse {
	return o.Payload
}

func (o *GetRulesEntitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainRuleEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRulesEntitiesForbidden creates a GetRulesEntitiesForbidden with default headers values
func NewGetRulesEntitiesForbidden() *GetRulesEntitiesForbidden {
	return &GetRulesEntitiesForbidden{}
}

/*
GetRulesEntitiesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRulesEntitiesForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64
}

// IsSuccess returns true when this get rules entities forbidden response has a 2xx status code
func (o *GetRulesEntitiesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rules entities forbidden response has a 3xx status code
func (o *GetRulesEntitiesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rules entities forbidden response has a 4xx status code
func (o *GetRulesEntitiesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rules entities forbidden response has a 5xx status code
func (o *GetRulesEntitiesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get rules entities forbidden response a status code equal to that given
func (o *GetRulesEntitiesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get rules entities forbidden response
func (o *GetRulesEntitiesForbidden) Code() int {
	return 403
}

func (o *GetRulesEntitiesForbidden) Error() string {
	return fmt.Sprintf("[POST /ti/rules/entities/rules/GET/v2][%d] getRulesEntitiesForbidden ", 403)
}

func (o *GetRulesEntitiesForbidden) String() string {
	return fmt.Sprintf("[POST /ti/rules/entities/rules/GET/v2][%d] getRulesEntitiesForbidden ", 403)
}

func (o *GetRulesEntitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	return nil
}

// NewGetRulesEntitiesTooManyRequests creates a GetRulesEntitiesTooManyRequests with default headers values
func NewGetRulesEntitiesTooManyRequests() *GetRulesEntitiesTooManyRequests {
	return &GetRulesEntitiesTooManyRequests{}
}

/*
GetRulesEntitiesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetRulesEntitiesTooManyRequests struct {

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

// IsSuccess returns true when this get rules entities too many requests response has a 2xx status code
func (o *GetRulesEntitiesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rules entities too many requests response has a 3xx status code
func (o *GetRulesEntitiesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rules entities too many requests response has a 4xx status code
func (o *GetRulesEntitiesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rules entities too many requests response has a 5xx status code
func (o *GetRulesEntitiesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get rules entities too many requests response a status code equal to that given
func (o *GetRulesEntitiesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get rules entities too many requests response
func (o *GetRulesEntitiesTooManyRequests) Code() int {
	return 429
}

func (o *GetRulesEntitiesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /ti/rules/entities/rules/GET/v2][%d] getRulesEntitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRulesEntitiesTooManyRequests) String() string {
	return fmt.Sprintf("[POST /ti/rules/entities/rules/GET/v2][%d] getRulesEntitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRulesEntitiesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetRulesEntitiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRulesEntitiesInternalServerError creates a GetRulesEntitiesInternalServerError with default headers values
func NewGetRulesEntitiesInternalServerError() *GetRulesEntitiesInternalServerError {
	return &GetRulesEntitiesInternalServerError{}
}

/*
GetRulesEntitiesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetRulesEntitiesInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainRuleEntitiesResponse
}

// IsSuccess returns true when this get rules entities internal server error response has a 2xx status code
func (o *GetRulesEntitiesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rules entities internal server error response has a 3xx status code
func (o *GetRulesEntitiesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rules entities internal server error response has a 4xx status code
func (o *GetRulesEntitiesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get rules entities internal server error response has a 5xx status code
func (o *GetRulesEntitiesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get rules entities internal server error response a status code equal to that given
func (o *GetRulesEntitiesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get rules entities internal server error response
func (o *GetRulesEntitiesInternalServerError) Code() int {
	return 500
}

func (o *GetRulesEntitiesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /ti/rules/entities/rules/GET/v2][%d] getRulesEntitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRulesEntitiesInternalServerError) String() string {
	return fmt.Sprintf("[POST /ti/rules/entities/rules/GET/v2][%d] getRulesEntitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRulesEntitiesInternalServerError) GetPayload() *models.DomainRuleEntitiesResponse {
	return o.Payload
}

func (o *GetRulesEntitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainRuleEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
