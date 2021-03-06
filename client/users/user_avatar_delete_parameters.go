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

// NewUserAvatarDeleteParams creates a new UserAvatarDeleteParams object
// with the default values initialized.
func NewUserAvatarDeleteParams() *UserAvatarDeleteParams {
	var ()
	return &UserAvatarDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserAvatarDeleteParamsWithTimeout creates a new UserAvatarDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserAvatarDeleteParamsWithTimeout(timeout time.Duration) *UserAvatarDeleteParams {
	var ()
	return &UserAvatarDeleteParams{

		timeout: timeout,
	}
}

// NewUserAvatarDeleteParamsWithContext creates a new UserAvatarDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserAvatarDeleteParamsWithContext(ctx context.Context) *UserAvatarDeleteParams {
	var ()
	return &UserAvatarDeleteParams{

		Context: ctx,
	}
}

// NewUserAvatarDeleteParamsWithHTTPClient creates a new UserAvatarDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserAvatarDeleteParamsWithHTTPClient(client *http.Client) *UserAvatarDeleteParams {
	var ()
	return &UserAvatarDeleteParams{
		HTTPClient: client,
	}
}

/*UserAvatarDeleteParams contains all the parameters to send to the API endpoint
for the user avatar delete operation typically these are written to a http.Request
*/
type UserAvatarDeleteParams struct {

	/*User
	  User unique identifier expressed as UUID or username.

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user avatar delete params
func (o *UserAvatarDeleteParams) WithTimeout(timeout time.Duration) *UserAvatarDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user avatar delete params
func (o *UserAvatarDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user avatar delete params
func (o *UserAvatarDeleteParams) WithContext(ctx context.Context) *UserAvatarDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user avatar delete params
func (o *UserAvatarDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user avatar delete params
func (o *UserAvatarDeleteParams) WithHTTPClient(client *http.Client) *UserAvatarDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user avatar delete params
func (o *UserAvatarDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUser adds the user to the user avatar delete params
func (o *UserAvatarDeleteParams) WithUser(user string) *UserAvatarDeleteParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the user avatar delete params
func (o *UserAvatarDeleteParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *UserAvatarDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
