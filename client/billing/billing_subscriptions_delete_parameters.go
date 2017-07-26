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
)

// NewBillingSubscriptionsDeleteParams creates a new BillingSubscriptionsDeleteParams object
// with the default values initialized.
func NewBillingSubscriptionsDeleteParams() *BillingSubscriptionsDeleteParams {
	var ()
	return &BillingSubscriptionsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBillingSubscriptionsDeleteParamsWithTimeout creates a new BillingSubscriptionsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBillingSubscriptionsDeleteParamsWithTimeout(timeout time.Duration) *BillingSubscriptionsDeleteParams {
	var ()
	return &BillingSubscriptionsDeleteParams{

		timeout: timeout,
	}
}

// NewBillingSubscriptionsDeleteParamsWithContext creates a new BillingSubscriptionsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewBillingSubscriptionsDeleteParamsWithContext(ctx context.Context) *BillingSubscriptionsDeleteParams {
	var ()
	return &BillingSubscriptionsDeleteParams{

		Context: ctx,
	}
}

// NewBillingSubscriptionsDeleteParamsWithHTTPClient creates a new BillingSubscriptionsDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBillingSubscriptionsDeleteParamsWithHTTPClient(client *http.Client) *BillingSubscriptionsDeleteParams {
	var ()
	return &BillingSubscriptionsDeleteParams{
		HTTPClient: client,
	}
}

/*BillingSubscriptionsDeleteParams contains all the parameters to send to the API endpoint
for the billing subscriptions delete operation typically these are written to a http.Request
*/
type BillingSubscriptionsDeleteParams struct {

	/*ID
	  Subscription unique identifier expressed as UUID.

	*/
	ID string
	/*Namespace
	  User or team name.

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the billing subscriptions delete params
func (o *BillingSubscriptionsDeleteParams) WithTimeout(timeout time.Duration) *BillingSubscriptionsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the billing subscriptions delete params
func (o *BillingSubscriptionsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the billing subscriptions delete params
func (o *BillingSubscriptionsDeleteParams) WithContext(ctx context.Context) *BillingSubscriptionsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the billing subscriptions delete params
func (o *BillingSubscriptionsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the billing subscriptions delete params
func (o *BillingSubscriptionsDeleteParams) WithHTTPClient(client *http.Client) *BillingSubscriptionsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the billing subscriptions delete params
func (o *BillingSubscriptionsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the billing subscriptions delete params
func (o *BillingSubscriptionsDeleteParams) WithID(id string) *BillingSubscriptionsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the billing subscriptions delete params
func (o *BillingSubscriptionsDeleteParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the billing subscriptions delete params
func (o *BillingSubscriptionsDeleteParams) WithNamespace(namespace string) *BillingSubscriptionsDeleteParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the billing subscriptions delete params
func (o *BillingSubscriptionsDeleteParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *BillingSubscriptionsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
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
