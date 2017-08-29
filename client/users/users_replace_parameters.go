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

	"github.com/3Blades/go-sdk/models"
)

// NewUsersReplaceParams creates a new UsersReplaceParams object
// with the default values initialized.
func NewUsersReplaceParams() *UsersReplaceParams {
	var ()
	return &UsersReplaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUsersReplaceParamsWithTimeout creates a new UsersReplaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUsersReplaceParamsWithTimeout(timeout time.Duration) *UsersReplaceParams {
	var ()
	return &UsersReplaceParams{

		timeout: timeout,
	}
}

// NewUsersReplaceParamsWithContext creates a new UsersReplaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewUsersReplaceParamsWithContext(ctx context.Context) *UsersReplaceParams {
	var ()
	return &UsersReplaceParams{

		Context: ctx,
	}
}

// NewUsersReplaceParamsWithHTTPClient creates a new UsersReplaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUsersReplaceParamsWithHTTPClient(client *http.Client) *UsersReplaceParams {
	var ()
	return &UsersReplaceParams{
		HTTPClient: client,
	}
}

/*UsersReplaceParams contains all the parameters to send to the API endpoint
for the users replace operation typically these are written to a http.Request
*/
type UsersReplaceParams struct {

	/*ID
	  User unique identifier expressed as UUID.

	*/
	ID string
	/*UserData*/
	UserData *models.UserData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the users replace params
func (o *UsersReplaceParams) WithTimeout(timeout time.Duration) *UsersReplaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users replace params
func (o *UsersReplaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users replace params
func (o *UsersReplaceParams) WithContext(ctx context.Context) *UsersReplaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users replace params
func (o *UsersReplaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users replace params
func (o *UsersReplaceParams) WithHTTPClient(client *http.Client) *UsersReplaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users replace params
func (o *UsersReplaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the users replace params
func (o *UsersReplaceParams) WithID(id string) *UsersReplaceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users replace params
func (o *UsersReplaceParams) SetID(id string) {
	o.ID = id
}

// WithUserData adds the userData to the users replace params
func (o *UsersReplaceParams) WithUserData(userData *models.UserData) *UsersReplaceParams {
	o.SetUserData(userData)
	return o
}

// SetUserData adds the userData to the users replace params
func (o *UsersReplaceParams) SetUserData(userData *models.UserData) {
	o.UserData = userData
}

// WriteToRequest writes these params to a swagger request
func (o *UsersReplaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.UserData == nil {
		o.UserData = new(models.UserData)
	}

	if err := r.SetBodyParam(o.UserData); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
