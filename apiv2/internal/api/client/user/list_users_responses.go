// Code generated by go-swagger; DO NOT EDIT.

package user

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

// ListUsersReader is a Reader for the ListUsers structure.
type ListUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListUsersOK creates a ListUsersOK with default headers values
func NewListUsersOK() *ListUsersOK {
	return &ListUsersOK{}
}

/* ListUsersOK describes a response with status code 200, with default header values.

return the list of users.
*/
type ListUsersOK struct {

	/* Link to previous page and next page
	 */
	Link string

	/* The total count of users
	 */
	XTotalCount int64

	Payload []*model.UserResp
}

func (o *ListUsersOK) Error() string {
	return fmt.Sprintf("[GET /users][%d] listUsersOK  %+v", 200, o.Payload)
}
func (o *ListUsersOK) GetPayload() []*model.UserResp {
	return o.Payload
}

func (o *ListUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListUsersUnauthorized creates a ListUsersUnauthorized with default headers values
func NewListUsersUnauthorized() *ListUsersUnauthorized {
	return &ListUsersUnauthorized{}
}

/* ListUsersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListUsersUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users][%d] listUsersUnauthorized  %+v", 401, o.Payload)
}
func (o *ListUsersUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListUsersForbidden creates a ListUsersForbidden with default headers values
func NewListUsersForbidden() *ListUsersForbidden {
	return &ListUsersForbidden{}
}

/* ListUsersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListUsersForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /users][%d] listUsersForbidden  %+v", 403, o.Payload)
}
func (o *ListUsersForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListUsersInternalServerError creates a ListUsersInternalServerError with default headers values
func NewListUsersInternalServerError() *ListUsersInternalServerError {
	return &ListUsersInternalServerError{}
}

/* ListUsersInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListUsersInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListUsersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users][%d] listUsersInternalServerError  %+v", 500, o.Payload)
}
func (o *ListUsersInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
