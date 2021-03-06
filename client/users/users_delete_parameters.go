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

// NewUsersDeleteParams creates a new UsersDeleteParams object
// with the default values initialized.
func NewUsersDeleteParams() *UsersDeleteParams {
	var ()
	return &UsersDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUsersDeleteParamsWithTimeout creates a new UsersDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUsersDeleteParamsWithTimeout(timeout time.Duration) *UsersDeleteParams {
	var ()
	return &UsersDeleteParams{

		timeout: timeout,
	}
}

// NewUsersDeleteParamsWithContext creates a new UsersDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewUsersDeleteParamsWithContext(ctx context.Context) *UsersDeleteParams {
	var ()
	return &UsersDeleteParams{

		Context: ctx,
	}
}

// NewUsersDeleteParamsWithHTTPClient creates a new UsersDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUsersDeleteParamsWithHTTPClient(client *http.Client) *UsersDeleteParams {
	var ()
	return &UsersDeleteParams{
		HTTPClient: client,
	}
}

/*UsersDeleteParams contains all the parameters to send to the API endpoint
for the users delete operation typically these are written to a http.Request
*/
type UsersDeleteParams struct {

	/*User
	  User identifier expressed as UUID or username.

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the users delete params
func (o *UsersDeleteParams) WithTimeout(timeout time.Duration) *UsersDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users delete params
func (o *UsersDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users delete params
func (o *UsersDeleteParams) WithContext(ctx context.Context) *UsersDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users delete params
func (o *UsersDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users delete params
func (o *UsersDeleteParams) WithHTTPClient(client *http.Client) *UsersDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users delete params
func (o *UsersDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUser adds the user to the users delete params
func (o *UsersDeleteParams) WithUser(user string) *UsersDeleteParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the users delete params
func (o *UsersDeleteParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *UsersDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
