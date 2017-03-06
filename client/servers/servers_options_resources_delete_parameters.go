package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewServersOptionsResourcesDeleteParams creates a new ServersOptionsResourcesDeleteParams object
// with the default values initialized.
func NewServersOptionsResourcesDeleteParams() *ServersOptionsResourcesDeleteParams {
	var ()
	return &ServersOptionsResourcesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServersOptionsResourcesDeleteParamsWithTimeout creates a new ServersOptionsResourcesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServersOptionsResourcesDeleteParamsWithTimeout(timeout time.Duration) *ServersOptionsResourcesDeleteParams {
	var ()
	return &ServersOptionsResourcesDeleteParams{

		timeout: timeout,
	}
}

// NewServersOptionsResourcesDeleteParamsWithContext creates a new ServersOptionsResourcesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewServersOptionsResourcesDeleteParamsWithContext(ctx context.Context) *ServersOptionsResourcesDeleteParams {
	var ()
	return &ServersOptionsResourcesDeleteParams{

		Context: ctx,
	}
}

// NewServersOptionsResourcesDeleteParamsWithHTTPClient creates a new ServersOptionsResourcesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServersOptionsResourcesDeleteParamsWithHTTPClient(client *http.Client) *ServersOptionsResourcesDeleteParams {
	var ()
	return &ServersOptionsResourcesDeleteParams{
		HTTPClient: client,
	}
}

/*ServersOptionsResourcesDeleteParams contains all the parameters to send to the API endpoint
for the servers options resources delete operation typically these are written to a http.Request
*/
type ServersOptionsResourcesDeleteParams struct {

	/*ID*/
	ID string
	/*Namespace*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the servers options resources delete params
func (o *ServersOptionsResourcesDeleteParams) WithTimeout(timeout time.Duration) *ServersOptionsResourcesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the servers options resources delete params
func (o *ServersOptionsResourcesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the servers options resources delete params
func (o *ServersOptionsResourcesDeleteParams) WithContext(ctx context.Context) *ServersOptionsResourcesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the servers options resources delete params
func (o *ServersOptionsResourcesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the servers options resources delete params
func (o *ServersOptionsResourcesDeleteParams) WithHTTPClient(client *http.Client) *ServersOptionsResourcesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the servers options resources delete params
func (o *ServersOptionsResourcesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the servers options resources delete params
func (o *ServersOptionsResourcesDeleteParams) WithID(id string) *ServersOptionsResourcesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the servers options resources delete params
func (o *ServersOptionsResourcesDeleteParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the servers options resources delete params
func (o *ServersOptionsResourcesDeleteParams) WithNamespace(namespace string) *ServersOptionsResourcesDeleteParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the servers options resources delete params
func (o *ServersOptionsResourcesDeleteParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *ServersOptionsResourcesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
