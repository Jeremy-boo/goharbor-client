// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v4/apiv2/model"
)

// GetHealthReader is a Reader for the GetHealth structure.
type GetHealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetHealthInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHealthOK creates a GetHealthOK with default headers values
func NewGetHealthOK() *GetHealthOK {
	return &GetHealthOK{}
}

/* GetHealthOK describes a response with status code 200, with default header values.

The health status of Harbor components
*/
type GetHealthOK struct {
	Payload *model.OverallHealthStatus
}

func (o *GetHealthOK) Error() string {
	return fmt.Sprintf("[GET /health][%d] getHealthOK  %+v", 200, o.Payload)
}
func (o *GetHealthOK) GetPayload() *model.OverallHealthStatus {
	return o.Payload
}

func (o *GetHealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.OverallHealthStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHealthInternalServerError creates a GetHealthInternalServerError with default headers values
func NewGetHealthInternalServerError() *GetHealthInternalServerError {
	return &GetHealthInternalServerError{}
}

/* GetHealthInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetHealthInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetHealthInternalServerError) Error() string {
	return fmt.Sprintf("[GET /health][%d] getHealthInternalServerError  %+v", 500, o.Payload)
}
func (o *GetHealthInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetHealthInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
