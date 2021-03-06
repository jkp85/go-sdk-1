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

// NewNotificationReadParams creates a new NotificationReadParams object
// with the default values initialized.
func NewNotificationReadParams() *NotificationReadParams {
	var ()
	return &NotificationReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNotificationReadParamsWithTimeout creates a new NotificationReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNotificationReadParamsWithTimeout(timeout time.Duration) *NotificationReadParams {
	var ()
	return &NotificationReadParams{

		timeout: timeout,
	}
}

// NewNotificationReadParamsWithContext creates a new NotificationReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewNotificationReadParamsWithContext(ctx context.Context) *NotificationReadParams {
	var ()
	return &NotificationReadParams{

		Context: ctx,
	}
}

// NewNotificationReadParamsWithHTTPClient creates a new NotificationReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNotificationReadParamsWithHTTPClient(client *http.Client) *NotificationReadParams {
	var ()
	return &NotificationReadParams{
		HTTPClient: client,
	}
}

/*NotificationReadParams contains all the parameters to send to the API endpoint
for the notification read operation typically these are written to a http.Request
*/
type NotificationReadParams struct {

	/*Namespace
	  User or team data.

	*/
	Namespace string
	/*NotificationID
	  Notification UUID.

	*/
	NotificationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the notification read params
func (o *NotificationReadParams) WithTimeout(timeout time.Duration) *NotificationReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the notification read params
func (o *NotificationReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the notification read params
func (o *NotificationReadParams) WithContext(ctx context.Context) *NotificationReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the notification read params
func (o *NotificationReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the notification read params
func (o *NotificationReadParams) WithHTTPClient(client *http.Client) *NotificationReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the notification read params
func (o *NotificationReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the notification read params
func (o *NotificationReadParams) WithNamespace(namespace string) *NotificationReadParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the notification read params
func (o *NotificationReadParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithNotificationID adds the notificationID to the notification read params
func (o *NotificationReadParams) WithNotificationID(notificationID string) *NotificationReadParams {
	o.SetNotificationID(notificationID)
	return o
}

// SetNotificationID adds the notificationId to the notification read params
func (o *NotificationReadParams) SetNotificationID(notificationID string) {
	o.NotificationID = notificationID
}

// WriteToRequest writes these params to a swagger request
func (o *NotificationReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param notification_id
	if err := r.SetPathParam("notification_id", o.NotificationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
