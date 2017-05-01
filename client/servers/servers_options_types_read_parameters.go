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

// NewServersOptionsTypesReadParams creates a new ServersOptionsTypesReadParams object
// with the default values initialized.
func NewServersOptionsTypesReadParams() *ServersOptionsTypesReadParams {
	var ()
	return &ServersOptionsTypesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServersOptionsTypesReadParamsWithTimeout creates a new ServersOptionsTypesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServersOptionsTypesReadParamsWithTimeout(timeout time.Duration) *ServersOptionsTypesReadParams {
	var ()
	return &ServersOptionsTypesReadParams{

		timeout: timeout,
	}
}

// NewServersOptionsTypesReadParamsWithContext creates a new ServersOptionsTypesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewServersOptionsTypesReadParamsWithContext(ctx context.Context) *ServersOptionsTypesReadParams {
	var ()
	return &ServersOptionsTypesReadParams{

		Context: ctx,
	}
}

// NewServersOptionsTypesReadParamsWithHTTPClient creates a new ServersOptionsTypesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServersOptionsTypesReadParamsWithHTTPClient(client *http.Client) *ServersOptionsTypesReadParams {
	var ()
	return &ServersOptionsTypesReadParams{
		HTTPClient: client,
	}
}

/*ServersOptionsTypesReadParams contains all the parameters to send to the API endpoint
for the servers options types read operation typically these are written to a http.Request
*/
type ServersOptionsTypesReadParams struct {

	/*ID*/
	ID string
	/*Namespace*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the servers options types read params
func (o *ServersOptionsTypesReadParams) WithTimeout(timeout time.Duration) *ServersOptionsTypesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the servers options types read params
func (o *ServersOptionsTypesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the servers options types read params
func (o *ServersOptionsTypesReadParams) WithContext(ctx context.Context) *ServersOptionsTypesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the servers options types read params
func (o *ServersOptionsTypesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the servers options types read params
func (o *ServersOptionsTypesReadParams) WithHTTPClient(client *http.Client) *ServersOptionsTypesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the servers options types read params
func (o *ServersOptionsTypesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the servers options types read params
func (o *ServersOptionsTypesReadParams) WithID(id string) *ServersOptionsTypesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the servers options types read params
func (o *ServersOptionsTypesReadParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the servers options types read params
func (o *ServersOptionsTypesReadParams) WithNamespace(namespace string) *ServersOptionsTypesReadParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the servers options types read params
func (o *ServersOptionsTypesReadParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *ServersOptionsTypesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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