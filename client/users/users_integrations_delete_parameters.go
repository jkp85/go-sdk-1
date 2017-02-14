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

// NewUsersIntegrationsDeleteParams creates a new UsersIntegrationsDeleteParams object
// with the default values initialized.
func NewUsersIntegrationsDeleteParams() *UsersIntegrationsDeleteParams {
	var ()
	return &UsersIntegrationsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUsersIntegrationsDeleteParamsWithTimeout creates a new UsersIntegrationsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUsersIntegrationsDeleteParamsWithTimeout(timeout time.Duration) *UsersIntegrationsDeleteParams {
	var ()
	return &UsersIntegrationsDeleteParams{

		timeout: timeout,
	}
}

// NewUsersIntegrationsDeleteParamsWithContext creates a new UsersIntegrationsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewUsersIntegrationsDeleteParamsWithContext(ctx context.Context) *UsersIntegrationsDeleteParams {
	var ()
	return &UsersIntegrationsDeleteParams{

		Context: ctx,
	}
}

// NewUsersIntegrationsDeleteParamsWithHTTPClient creates a new UsersIntegrationsDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUsersIntegrationsDeleteParamsWithHTTPClient(client *http.Client) *UsersIntegrationsDeleteParams {
	var ()
	return &UsersIntegrationsDeleteParams{
		HTTPClient: client,
	}
}

/*UsersIntegrationsDeleteParams contains all the parameters to send to the API endpoint
for the users integrations delete operation typically these are written to a http.Request
*/
type UsersIntegrationsDeleteParams struct {

	/*ID*/
	ID string
	/*Namespace*/
	Namespace string
	/*UserPk*/
	UserPk string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the users integrations delete params
func (o *UsersIntegrationsDeleteParams) WithTimeout(timeout time.Duration) *UsersIntegrationsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users integrations delete params
func (o *UsersIntegrationsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users integrations delete params
func (o *UsersIntegrationsDeleteParams) WithContext(ctx context.Context) *UsersIntegrationsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users integrations delete params
func (o *UsersIntegrationsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users integrations delete params
func (o *UsersIntegrationsDeleteParams) WithHTTPClient(client *http.Client) *UsersIntegrationsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users integrations delete params
func (o *UsersIntegrationsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the users integrations delete params
func (o *UsersIntegrationsDeleteParams) WithID(id string) *UsersIntegrationsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users integrations delete params
func (o *UsersIntegrationsDeleteParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the users integrations delete params
func (o *UsersIntegrationsDeleteParams) WithNamespace(namespace string) *UsersIntegrationsDeleteParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the users integrations delete params
func (o *UsersIntegrationsDeleteParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserPk adds the userPk to the users integrations delete params
func (o *UsersIntegrationsDeleteParams) WithUserPk(userPk string) *UsersIntegrationsDeleteParams {
	o.SetUserPk(userPk)
	return o
}

// SetUserPk adds the userPk to the users integrations delete params
func (o *UsersIntegrationsDeleteParams) SetUserPk(userPk string) {
	o.UserPk = userPk
}

// WriteToRequest writes these params to a swagger request
func (o *UsersIntegrationsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param user_pk
	if err := r.SetPathParam("user_pk", o.UserPk); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
