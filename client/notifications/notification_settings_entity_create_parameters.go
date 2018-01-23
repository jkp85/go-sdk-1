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

	models "github.com/IllumiDesk/go-sdk/models"
)

// NewNotificationSettingsEntityCreateParams creates a new NotificationSettingsEntityCreateParams object
// with the default values initialized.
func NewNotificationSettingsEntityCreateParams() *NotificationSettingsEntityCreateParams {
	var ()
	return &NotificationSettingsEntityCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNotificationSettingsEntityCreateParamsWithTimeout creates a new NotificationSettingsEntityCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNotificationSettingsEntityCreateParamsWithTimeout(timeout time.Duration) *NotificationSettingsEntityCreateParams {
	var ()
	return &NotificationSettingsEntityCreateParams{

		timeout: timeout,
	}
}

// NewNotificationSettingsEntityCreateParamsWithContext creates a new NotificationSettingsEntityCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewNotificationSettingsEntityCreateParamsWithContext(ctx context.Context) *NotificationSettingsEntityCreateParams {
	var ()
	return &NotificationSettingsEntityCreateParams{

		Context: ctx,
	}
}

// NewNotificationSettingsEntityCreateParamsWithHTTPClient creates a new NotificationSettingsEntityCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNotificationSettingsEntityCreateParamsWithHTTPClient(client *http.Client) *NotificationSettingsEntityCreateParams {
	var ()
	return &NotificationSettingsEntityCreateParams{
		HTTPClient: client,
	}
}

/*NotificationSettingsEntityCreateParams contains all the parameters to send to the API endpoint
for the notification settings entity create operation typically these are written to a http.Request
*/
type NotificationSettingsEntityCreateParams struct {

	/*Entity
	  Entity whose settings should be retrieved.

	*/
	Entity string
	/*Namespace
	  User or team name.

	*/
	Namespace string
	/*NotificationSettingsData*/
	NotificationSettingsData *models.NotificationSettingsData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the notification settings entity create params
func (o *NotificationSettingsEntityCreateParams) WithTimeout(timeout time.Duration) *NotificationSettingsEntityCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the notification settings entity create params
func (o *NotificationSettingsEntityCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the notification settings entity create params
func (o *NotificationSettingsEntityCreateParams) WithContext(ctx context.Context) *NotificationSettingsEntityCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the notification settings entity create params
func (o *NotificationSettingsEntityCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the notification settings entity create params
func (o *NotificationSettingsEntityCreateParams) WithHTTPClient(client *http.Client) *NotificationSettingsEntityCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the notification settings entity create params
func (o *NotificationSettingsEntityCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntity adds the entity to the notification settings entity create params
func (o *NotificationSettingsEntityCreateParams) WithEntity(entity string) *NotificationSettingsEntityCreateParams {
	o.SetEntity(entity)
	return o
}

// SetEntity adds the entity to the notification settings entity create params
func (o *NotificationSettingsEntityCreateParams) SetEntity(entity string) {
	o.Entity = entity
}

// WithNamespace adds the namespace to the notification settings entity create params
func (o *NotificationSettingsEntityCreateParams) WithNamespace(namespace string) *NotificationSettingsEntityCreateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the notification settings entity create params
func (o *NotificationSettingsEntityCreateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithNotificationSettingsData adds the notificationSettingsData to the notification settings entity create params
func (o *NotificationSettingsEntityCreateParams) WithNotificationSettingsData(notificationSettingsData *models.NotificationSettingsData) *NotificationSettingsEntityCreateParams {
	o.SetNotificationSettingsData(notificationSettingsData)
	return o
}

// SetNotificationSettingsData adds the notificationSettingsData to the notification settings entity create params
func (o *NotificationSettingsEntityCreateParams) SetNotificationSettingsData(notificationSettingsData *models.NotificationSettingsData) {
	o.NotificationSettingsData = notificationSettingsData
}

// WriteToRequest writes these params to a swagger request
func (o *NotificationSettingsEntityCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.NotificationSettingsData != nil {
		if err := r.SetBodyParam(o.NotificationSettingsData); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
