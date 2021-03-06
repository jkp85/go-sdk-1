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

// NewUsersSSHKeyListParams creates a new UsersSSHKeyListParams object
// with the default values initialized.
func NewUsersSSHKeyListParams() *UsersSSHKeyListParams {
	var ()
	return &UsersSSHKeyListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUsersSSHKeyListParamsWithTimeout creates a new UsersSSHKeyListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUsersSSHKeyListParamsWithTimeout(timeout time.Duration) *UsersSSHKeyListParams {
	var ()
	return &UsersSSHKeyListParams{

		timeout: timeout,
	}
}

// NewUsersSSHKeyListParamsWithContext creates a new UsersSSHKeyListParams object
// with the default values initialized, and the ability to set a context for a request
func NewUsersSSHKeyListParamsWithContext(ctx context.Context) *UsersSSHKeyListParams {
	var ()
	return &UsersSSHKeyListParams{

		Context: ctx,
	}
}

// NewUsersSSHKeyListParamsWithHTTPClient creates a new UsersSSHKeyListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUsersSSHKeyListParamsWithHTTPClient(client *http.Client) *UsersSSHKeyListParams {
	var ()
	return &UsersSSHKeyListParams{
		HTTPClient: client,
	}
}

/*UsersSSHKeyListParams contains all the parameters to send to the API endpoint
for the users ssh key list operation typically these are written to a http.Request
*/
type UsersSSHKeyListParams struct {

	/*User
	  User unique identifier expressed as UUID or username.

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the users ssh key list params
func (o *UsersSSHKeyListParams) WithTimeout(timeout time.Duration) *UsersSSHKeyListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users ssh key list params
func (o *UsersSSHKeyListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users ssh key list params
func (o *UsersSSHKeyListParams) WithContext(ctx context.Context) *UsersSSHKeyListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users ssh key list params
func (o *UsersSSHKeyListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users ssh key list params
func (o *UsersSSHKeyListParams) WithHTTPClient(client *http.Client) *UsersSSHKeyListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users ssh key list params
func (o *UsersSSHKeyListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUser adds the user to the users ssh key list params
func (o *UsersSSHKeyListParams) WithUser(user string) *UsersSSHKeyListParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the users ssh key list params
func (o *UsersSSHKeyListParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *UsersSSHKeyListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
