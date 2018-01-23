// Code generated by go-swagger; DO NOT EDIT.

package notifications

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

// NewNotificationSettingsEntityReadParams creates a new NotificationSettingsEntityReadParams object
// with the default values initialized.
func NewNotificationSettingsEntityReadParams() *NotificationSettingsEntityReadParams {
	var ()
	return &NotificationSettingsEntityReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNotificationSettingsEntityReadParamsWithTimeout creates a new NotificationSettingsEntityReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNotificationSettingsEntityReadParamsWithTimeout(timeout time.Duration) *NotificationSettingsEntityReadParams {
	var ()
	return &NotificationSettingsEntityReadParams{

		timeout: timeout,
	}
}

// NewNotificationSettingsEntityReadParamsWithContext creates a new NotificationSettingsEntityReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewNotificationSettingsEntityReadParamsWithContext(ctx context.Context) *NotificationSettingsEntityReadParams {
	var ()
	return &NotificationSettingsEntityReadParams{

		Context: ctx,
	}
}

// NewNotificationSettingsEntityReadParamsWithHTTPClient creates a new NotificationSettingsEntityReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNotificationSettingsEntityReadParamsWithHTTPClient(client *http.Client) *NotificationSettingsEntityReadParams {
	var ()
	return &NotificationSettingsEntityReadParams{
		HTTPClient: client,
	}
}

/*NotificationSettingsEntityReadParams contains all the parameters to send to the API endpoint
for the notification settings entity read operation typically these are written to a http.Request
*/
type NotificationSettingsEntityReadParams struct {

	/*Entity
	  Entity whose settings should be retrieved.

	*/
	Entity string
	/*Namespace
	  User or team data.

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the notification settings entity read params
func (o *NotificationSettingsEntityReadParams) WithTimeout(timeout time.Duration) *NotificationSettingsEntityReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the notification settings entity read params
func (o *NotificationSettingsEntityReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the notification settings entity read params
func (o *NotificationSettingsEntityReadParams) WithContext(ctx context.Context) *NotificationSettingsEntityReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the notification settings entity read params
func (o *NotificationSettingsEntityReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the notification settings entity read params
func (o *NotificationSettingsEntityReadParams) WithHTTPClient(client *http.Client) *NotificationSettingsEntityReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the notification settings entity read params
func (o *NotificationSettingsEntityReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntity adds the entity to the notification settings entity read params
func (o *NotificationSettingsEntityReadParams) WithEntity(entity string) *NotificationSettingsEntityReadParams {
	o.SetEntity(entity)
	return o
}

// SetEntity adds the entity to the notification settings entity read params
func (o *NotificationSettingsEntityReadParams) SetEntity(entity string) {
	o.Entity = entity
}

// WithNamespace adds the namespace to the notification settings entity read params
func (o *NotificationSettingsEntityReadParams) WithNamespace(namespace string) *NotificationSettingsEntityReadParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the notification settings entity read params
func (o *NotificationSettingsEntityReadParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *NotificationSettingsEntityReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entity
	if err := r.SetPathParam("entity", o.Entity); err != nil {
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
