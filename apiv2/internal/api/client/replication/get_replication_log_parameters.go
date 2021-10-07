// Code generated by go-swagger; DO NOT EDIT.

package replication

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

// NewGetReplicationLogParams creates a new GetReplicationLogParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetReplicationLogParams() *GetReplicationLogParams {
	return &GetReplicationLogParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetReplicationLogParamsWithTimeout creates a new GetReplicationLogParams object
// with the ability to set a timeout on a request.
func NewGetReplicationLogParamsWithTimeout(timeout time.Duration) *GetReplicationLogParams {
	return &GetReplicationLogParams{
		timeout: timeout,
	}
}

// NewGetReplicationLogParamsWithContext creates a new GetReplicationLogParams object
// with the ability to set a context for a request.
func NewGetReplicationLogParamsWithContext(ctx context.Context) *GetReplicationLogParams {
	return &GetReplicationLogParams{
		Context: ctx,
	}
}

// NewGetReplicationLogParamsWithHTTPClient creates a new GetReplicationLogParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetReplicationLogParamsWithHTTPClient(client *http.Client) *GetReplicationLogParams {
	return &GetReplicationLogParams{
		HTTPClient: client,
	}
}

/* GetReplicationLogParams contains all the parameters to send to the API endpoint
   for the get replication log operation.

   Typically these are written to a http.Request.
*/
type GetReplicationLogParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* ID.

	   The ID of the execution that the tasks belongs to.

	   Format: int64
	*/
	ID int64

	/* TaskID.

	   The ID of the task.

	   Format: int64
	*/
	TaskID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get replication log params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReplicationLogParams) WithDefaults() *GetReplicationLogParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get replication log params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReplicationLogParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get replication log params
func (o *GetReplicationLogParams) WithTimeout(timeout time.Duration) *GetReplicationLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get replication log params
func (o *GetReplicationLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get replication log params
func (o *GetReplicationLogParams) WithContext(ctx context.Context) *GetReplicationLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get replication log params
func (o *GetReplicationLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get replication log params
func (o *GetReplicationLogParams) WithHTTPClient(client *http.Client) *GetReplicationLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get replication log params
func (o *GetReplicationLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get replication log params
func (o *GetReplicationLogParams) WithXRequestID(xRequestID *string) *GetReplicationLogParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get replication log params
func (o *GetReplicationLogParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the get replication log params
func (o *GetReplicationLogParams) WithID(id int64) *GetReplicationLogParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get replication log params
func (o *GetReplicationLogParams) SetID(id int64) {
	o.ID = id
}

// WithTaskID adds the taskID to the get replication log params
func (o *GetReplicationLogParams) WithTaskID(taskID int64) *GetReplicationLogParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the get replication log params
func (o *GetReplicationLogParams) SetTaskID(taskID int64) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *GetReplicationLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param task_id
	if err := r.SetPathParam("task_id", swag.FormatInt64(o.TaskID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
