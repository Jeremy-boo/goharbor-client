// Code generated by go-swagger; DO NOT EDIT.

package quota

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetQuotaParams creates a new GetQuotaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetQuotaParams() *GetQuotaParams {
	return &GetQuotaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetQuotaParamsWithTimeout creates a new GetQuotaParams object
// with the ability to set a timeout on a request.
func NewGetQuotaParamsWithTimeout(timeout time.Duration) *GetQuotaParams {
	return &GetQuotaParams{
		timeout: timeout,
	}
}

// NewGetQuotaParamsWithContext creates a new GetQuotaParams object
// with the ability to set a context for a request.
func NewGetQuotaParamsWithContext(ctx context.Context) *GetQuotaParams {
	return &GetQuotaParams{
		Context: ctx,
	}
}

// NewGetQuotaParamsWithHTTPClient creates a new GetQuotaParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetQuotaParamsWithHTTPClient(client *http.Client) *GetQuotaParams {
	return &GetQuotaParams{
		HTTPClient: client,
	}
}

/* GetQuotaParams contains all the parameters to send to the API endpoint
   for the get quota operation.

   Typically these are written to a http.Request.
*/
type GetQuotaParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* ID.

	   Quota ID
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get quota params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQuotaParams) WithDefaults() *GetQuotaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get quota params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQuotaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get quota params
func (o *GetQuotaParams) WithTimeout(timeout time.Duration) *GetQuotaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get quota params
func (o *GetQuotaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get quota params
func (o *GetQuotaParams) WithContext(ctx context.Context) *GetQuotaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get quota params
func (o *GetQuotaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get quota params
func (o *GetQuotaParams) WithHTTPClient(client *http.Client) *GetQuotaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get quota params
func (o *GetQuotaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get quota params
func (o *GetQuotaParams) WithXRequestID(xRequestID *string) *GetQuotaParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get quota params
func (o *GetQuotaParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the get quota params
func (o *GetQuotaParams) WithID(id int64) *GetQuotaParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get quota params
func (o *GetQuotaParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetQuotaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
