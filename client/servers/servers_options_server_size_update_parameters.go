// Code generated by go-swagger; DO NOT EDIT.

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

	models "github.com/IllumiDesk/go-sdk/models"
)

// NewServersOptionsServerSizeUpdateParams creates a new ServersOptionsServerSizeUpdateParams object
// with the default values initialized.
func NewServersOptionsServerSizeUpdateParams() *ServersOptionsServerSizeUpdateParams {
	var ()
	return &ServersOptionsServerSizeUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServersOptionsServerSizeUpdateParamsWithTimeout creates a new ServersOptionsServerSizeUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServersOptionsServerSizeUpdateParamsWithTimeout(timeout time.Duration) *ServersOptionsServerSizeUpdateParams {
	var ()
	return &ServersOptionsServerSizeUpdateParams{

		timeout: timeout,
	}
}

// NewServersOptionsServerSizeUpdateParamsWithContext creates a new ServersOptionsServerSizeUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewServersOptionsServerSizeUpdateParamsWithContext(ctx context.Context) *ServersOptionsServerSizeUpdateParams {
	var ()
	return &ServersOptionsServerSizeUpdateParams{

		Context: ctx,
	}
}

// NewServersOptionsServerSizeUpdateParamsWithHTTPClient creates a new ServersOptionsServerSizeUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServersOptionsServerSizeUpdateParamsWithHTTPClient(client *http.Client) *ServersOptionsServerSizeUpdateParams {
	var ()
	return &ServersOptionsServerSizeUpdateParams{
		HTTPClient: client,
	}
}

/*ServersOptionsServerSizeUpdateParams contains all the parameters to send to the API endpoint
for the servers options server size update operation typically these are written to a http.Request
*/
type ServersOptionsServerSizeUpdateParams struct {

	/*ServersizeData*/
	ServersizeData *models.ServerSizeData
	/*Size
	  Server size unique identifier expressed as UUID or name.

	*/
	Size string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the servers options server size update params
func (o *ServersOptionsServerSizeUpdateParams) WithTimeout(timeout time.Duration) *ServersOptionsServerSizeUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the servers options server size update params
func (o *ServersOptionsServerSizeUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the servers options server size update params
func (o *ServersOptionsServerSizeUpdateParams) WithContext(ctx context.Context) *ServersOptionsServerSizeUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the servers options server size update params
func (o *ServersOptionsServerSizeUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the servers options server size update params
func (o *ServersOptionsServerSizeUpdateParams) WithHTTPClient(client *http.Client) *ServersOptionsServerSizeUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the servers options server size update params
func (o *ServersOptionsServerSizeUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServersizeData adds the serversizeData to the servers options server size update params
func (o *ServersOptionsServerSizeUpdateParams) WithServersizeData(serversizeData *models.ServerSizeData) *ServersOptionsServerSizeUpdateParams {
	o.SetServersizeData(serversizeData)
	return o
}

// SetServersizeData adds the serversizeData to the servers options server size update params
func (o *ServersOptionsServerSizeUpdateParams) SetServersizeData(serversizeData *models.ServerSizeData) {
	o.ServersizeData = serversizeData
}

// WithSize adds the size to the servers options server size update params
func (o *ServersOptionsServerSizeUpdateParams) WithSize(size string) *ServersOptionsServerSizeUpdateParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the servers options server size update params
func (o *ServersOptionsServerSizeUpdateParams) SetSize(size string) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *ServersOptionsServerSizeUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ServersizeData != nil {
		if err := r.SetBodyParam(o.ServersizeData); err != nil {
			return err
		}
	}

	// path param size
	if err := r.SetPathParam("size", o.Size); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
