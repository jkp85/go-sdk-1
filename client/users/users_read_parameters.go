// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewUsersReadParams creates a new UsersReadParams object
// with the default values initialized.
func NewUsersReadParams() *UsersReadParams {
	var ()
	return &UsersReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUsersReadParamsWithTimeout creates a new UsersReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUsersReadParamsWithTimeout(timeout time.Duration) *UsersReadParams {
	var ()
	return &UsersReadParams{

		timeout: timeout,
	}
}

// NewUsersReadParamsWithContext creates a new UsersReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewUsersReadParamsWithContext(ctx context.Context) *UsersReadParams {
	var ()
	return &UsersReadParams{

		Context: ctx,
	}
}

// NewUsersReadParamsWithHTTPClient creates a new UsersReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUsersReadParamsWithHTTPClient(client *http.Client) *UsersReadParams {
	var ()
	return &UsersReadParams{
		HTTPClient: client,
	}
}

/*UsersReadParams contains all the parameters to send to the API endpoint
for the users read operation typically these are written to a http.Request
*/
type UsersReadParams struct {

	/*User
	  Unique identifier expressed as UUID or username.

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the users read params
func (o *UsersReadParams) WithTimeout(timeout time.Duration) *UsersReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users read params
func (o *UsersReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users read params
func (o *UsersReadParams) WithContext(ctx context.Context) *UsersReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users read params
func (o *UsersReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users read params
func (o *UsersReadParams) WithHTTPClient(client *http.Client) *UsersReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users read params
func (o *UsersReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUser adds the user to the users read params
func (o *UsersReadParams) WithUser(user string) *UsersReadParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the users read params
func (o *UsersReadParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *UsersReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user
	if err := r.SetPathParam("user", o.User); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
