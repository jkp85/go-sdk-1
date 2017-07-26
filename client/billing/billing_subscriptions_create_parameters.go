// Code generated by go-swagger; DO NOT EDIT.

package billing

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

// NewBillingSubscriptionsCreateParams creates a new BillingSubscriptionsCreateParams object
// with the default values initialized.
func NewBillingSubscriptionsCreateParams() *BillingSubscriptionsCreateParams {
	var ()
	return &BillingSubscriptionsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBillingSubscriptionsCreateParamsWithTimeout creates a new BillingSubscriptionsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBillingSubscriptionsCreateParamsWithTimeout(timeout time.Duration) *BillingSubscriptionsCreateParams {
	var ()
	return &BillingSubscriptionsCreateParams{

		timeout: timeout,
	}
}

// NewBillingSubscriptionsCreateParamsWithContext creates a new BillingSubscriptionsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewBillingSubscriptionsCreateParamsWithContext(ctx context.Context) *BillingSubscriptionsCreateParams {
	var ()
	return &BillingSubscriptionsCreateParams{

		Context: ctx,
	}
}

// NewBillingSubscriptionsCreateParamsWithHTTPClient creates a new BillingSubscriptionsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBillingSubscriptionsCreateParamsWithHTTPClient(client *http.Client) *BillingSubscriptionsCreateParams {
	var ()
	return &BillingSubscriptionsCreateParams{
		HTTPClient: client,
	}
}

/*BillingSubscriptionsCreateParams contains all the parameters to send to the API endpoint
for the billing subscriptions create operation typically these are written to a http.Request
*/
type BillingSubscriptionsCreateParams struct {

	/*Data*/
	Data *models.SubscriptionData
	/*Namespace
	  User or team name.

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the billing subscriptions create params
func (o *BillingSubscriptionsCreateParams) WithTimeout(timeout time.Duration) *BillingSubscriptionsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the billing subscriptions create params
func (o *BillingSubscriptionsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the billing subscriptions create params
func (o *BillingSubscriptionsCreateParams) WithContext(ctx context.Context) *BillingSubscriptionsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the billing subscriptions create params
func (o *BillingSubscriptionsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the billing subscriptions create params
func (o *BillingSubscriptionsCreateParams) WithHTTPClient(client *http.Client) *BillingSubscriptionsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the billing subscriptions create params
func (o *BillingSubscriptionsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the billing subscriptions create params
func (o *BillingSubscriptionsCreateParams) WithData(data *models.SubscriptionData) *BillingSubscriptionsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the billing subscriptions create params
func (o *BillingSubscriptionsCreateParams) SetData(data *models.SubscriptionData) {
	o.Data = data
}

// WithNamespace adds the namespace to the billing subscriptions create params
func (o *BillingSubscriptionsCreateParams) WithNamespace(namespace string) *BillingSubscriptionsCreateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the billing subscriptions create params
func (o *BillingSubscriptionsCreateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *BillingSubscriptionsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data == nil {
		o.Data = new(models.SubscriptionData)
	}

	if err := r.SetBodyParam(o.Data); err != nil {
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
