// Code generated by go-swagger; DO NOT EDIT.

package scanner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v4/apiv2/model"
)

// GetScannerReader is a Reader for the GetScanner structure.
type GetScannerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScannerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScannerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetScannerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScannerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScannerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScannerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScannerOK creates a GetScannerOK with default headers values
func NewGetScannerOK() *GetScannerOK {
	return &GetScannerOK{}
}

/* GetScannerOK describes a response with status code 200, with default header values.

The details of the scanner registration.
*/
type GetScannerOK struct {
	Payload *model.ScannerRegistration
}

func (o *GetScannerOK) Error() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}][%d] getScannerOK  %+v", 200, o.Payload)
}
func (o *GetScannerOK) GetPayload() *model.ScannerRegistration {
	return o.Payload
}

func (o *GetScannerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ScannerRegistration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScannerUnauthorized creates a GetScannerUnauthorized with default headers values
func NewGetScannerUnauthorized() *GetScannerUnauthorized {
	return &GetScannerUnauthorized{}
}

/* GetScannerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetScannerUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetScannerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}][%d] getScannerUnauthorized  %+v", 401, o.Payload)
}
func (o *GetScannerUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetScannerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScannerForbidden creates a GetScannerForbidden with default headers values
func NewGetScannerForbidden() *GetScannerForbidden {
	return &GetScannerForbidden{}
}

/* GetScannerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetScannerForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetScannerForbidden) Error() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}][%d] getScannerForbidden  %+v", 403, o.Payload)
}
func (o *GetScannerForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetScannerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScannerNotFound creates a GetScannerNotFound with default headers values
func NewGetScannerNotFound() *GetScannerNotFound {
	return &GetScannerNotFound{}
}

/* GetScannerNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetScannerNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetScannerNotFound) Error() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}][%d] getScannerNotFound  %+v", 404, o.Payload)
}
func (o *GetScannerNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetScannerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScannerInternalServerError creates a GetScannerInternalServerError with default headers values
func NewGetScannerInternalServerError() *GetScannerInternalServerError {
	return &GetScannerInternalServerError{}
}

/* GetScannerInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetScannerInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetScannerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}][%d] getScannerInternalServerError  %+v", 500, o.Payload)
}
func (o *GetScannerInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetScannerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
