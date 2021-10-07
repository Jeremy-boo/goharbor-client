// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/mittwald/goharbor-client/v4/apiv2/model"
)

// ListAllRepositoriesReader is a Reader for the ListAllRepositories structure.
type ListAllRepositoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllRepositoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAllRepositoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListAllRepositoriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListAllRepositoriesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAllRepositoriesOK creates a ListAllRepositoriesOK with default headers values
func NewListAllRepositoriesOK() *ListAllRepositoriesOK {
	return &ListAllRepositoriesOK{}
}

/* ListAllRepositoriesOK describes a response with status code 200, with default header values.

Success
*/
type ListAllRepositoriesOK struct {

	/* Link refers to the previous page and next page
	 */
	Link string

	/* The total count of repositories
	 */
	XTotalCount int64

	Payload []*model.Repository
}

func (o *ListAllRepositoriesOK) Error() string {
	return fmt.Sprintf("[GET /repositories][%d] listAllRepositoriesOK  %+v", 200, o.Payload)
}
func (o *ListAllRepositoriesOK) GetPayload() []*model.Repository {
	return o.Payload
}

func (o *ListAllRepositoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllRepositoriesBadRequest creates a ListAllRepositoriesBadRequest with default headers values
func NewListAllRepositoriesBadRequest() *ListAllRepositoriesBadRequest {
	return &ListAllRepositoriesBadRequest{}
}

/* ListAllRepositoriesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ListAllRepositoriesBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListAllRepositoriesBadRequest) Error() string {
	return fmt.Sprintf("[GET /repositories][%d] listAllRepositoriesBadRequest  %+v", 400, o.Payload)
}
func (o *ListAllRepositoriesBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListAllRepositoriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListAllRepositoriesInternalServerError creates a ListAllRepositoriesInternalServerError with default headers values
func NewListAllRepositoriesInternalServerError() *ListAllRepositoriesInternalServerError {
	return &ListAllRepositoriesInternalServerError{}
}

/* ListAllRepositoriesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListAllRepositoriesInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListAllRepositoriesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /repositories][%d] listAllRepositoriesInternalServerError  %+v", 500, o.Payload)
}
func (o *ListAllRepositoriesInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListAllRepositoriesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
