// Code generated by go-swagger; DO NOT EDIT.

package usergroup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v4/apiv2/model"
)

// ListUserGroupsReader is a Reader for the ListUserGroups structure.
type ListUserGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUserGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListUserGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListUserGroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListUserGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListUserGroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListUserGroupsOK creates a ListUserGroupsOK with default headers values
func NewListUserGroupsOK() *ListUserGroupsOK {
	return &ListUserGroupsOK{}
}

/* ListUserGroupsOK describes a response with status code 200, with default header values.

Get user group successfully.
*/
type ListUserGroupsOK struct {
	Payload []*model.UserGroup
}

func (o *ListUserGroupsOK) Error() string {
	return fmt.Sprintf("[GET /usergroups][%d] listUserGroupsOK  %+v", 200, o.Payload)
}
func (o *ListUserGroupsOK) GetPayload() []*model.UserGroup {
	return o.Payload
}

func (o *ListUserGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserGroupsUnauthorized creates a ListUserGroupsUnauthorized with default headers values
func NewListUserGroupsUnauthorized() *ListUserGroupsUnauthorized {
	return &ListUserGroupsUnauthorized{}
}

/* ListUserGroupsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListUserGroupsUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListUserGroupsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /usergroups][%d] listUserGroupsUnauthorized  %+v", 401, o.Payload)
}
func (o *ListUserGroupsUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListUserGroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListUserGroupsForbidden creates a ListUserGroupsForbidden with default headers values
func NewListUserGroupsForbidden() *ListUserGroupsForbidden {
	return &ListUserGroupsForbidden{}
}

/* ListUserGroupsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListUserGroupsForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListUserGroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /usergroups][%d] listUserGroupsForbidden  %+v", 403, o.Payload)
}
func (o *ListUserGroupsForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListUserGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListUserGroupsInternalServerError creates a ListUserGroupsInternalServerError with default headers values
func NewListUserGroupsInternalServerError() *ListUserGroupsInternalServerError {
	return &ListUserGroupsInternalServerError{}
}

/* ListUserGroupsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListUserGroupsInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListUserGroupsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /usergroups][%d] listUserGroupsInternalServerError  %+v", 500, o.Payload)
}
func (o *ListUserGroupsInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListUserGroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
