// Code generated by go-swagger; DO NOT EDIT.

package configure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new configure API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for configure API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetConfigurations(params *GetConfigurationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetConfigurationsOK, error)

	GetInternalconfig(params *GetInternalconfigParams, authInfo runtime.ClientAuthInfoWriter) (*GetInternalconfigOK, error)

	UpdateConfigurations(params *UpdateConfigurationsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateConfigurationsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetConfigurations gets system configurations

  This endpoint is for retrieving system configurations that only provides for admin user.

*/
func (a *Client) GetConfigurations(params *GetConfigurationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetConfigurationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConfigurationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getConfigurations",
		Method:             "GET",
		PathPattern:        "/configurations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetConfigurationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetConfigurationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getConfigurations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetInternalconfig gets internal configurations

  This endpoint is for retrieving system configurations that only provides for internal api call.

*/
func (a *Client) GetInternalconfig(params *GetInternalconfigParams, authInfo runtime.ClientAuthInfoWriter) (*GetInternalconfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInternalconfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getInternalconfig",
		Method:             "GET",
		PathPattern:        "/internalconfig",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetInternalconfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInternalconfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInternalconfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateConfigurations modifies system configurations

  This endpoint is for modifying system configurations that only provides for admin user.

*/
func (a *Client) UpdateConfigurations(params *UpdateConfigurationsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateConfigurationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateConfigurationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateConfigurations",
		Method:             "PUT",
		PathPattern:        "/configurations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateConfigurationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateConfigurationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateConfigurations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
