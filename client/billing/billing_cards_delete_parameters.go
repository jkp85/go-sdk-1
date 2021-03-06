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

// NewBillingCardsDeleteParams creates a new BillingCardsDeleteParams object
// with the default values initialized.
func NewBillingCardsDeleteParams() *BillingCardsDeleteParams {
	var ()
	return &BillingCardsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBillingCardsDeleteParamsWithTimeout creates a new BillingCardsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBillingCardsDeleteParamsWithTimeout(timeout time.Duration) *BillingCardsDeleteParams {
	var ()
	return &BillingCardsDeleteParams{

		timeout: timeout,
	}
}

// NewBillingCardsDeleteParamsWithContext creates a new BillingCardsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewBillingCardsDeleteParamsWithContext(ctx context.Context) *BillingCardsDeleteParams {
	var ()
	return &BillingCardsDeleteParams{

		Context: ctx,
	}
}

// NewBillingCardsDeleteParamsWithHTTPClient creates a new BillingCardsDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBillingCardsDeleteParamsWithHTTPClient(client *http.Client) *BillingCardsDeleteParams {
	var ()
	return &BillingCardsDeleteParams{
		HTTPClient: client,
	}
}

/*BillingCardsDeleteParams contains all the parameters to send to the API endpoint
for the billing cards delete operation typically these are written to a http.Request
*/
type BillingCardsDeleteParams struct {

	/*ID
	  Card unique identifier expressed as UUID.

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

// WithTimeout adds the timeout to the billing cards delete params
func (o *BillingCardsDeleteParams) WithTimeout(timeout time.Duration) *BillingCardsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the billing cards delete params
func (o *BillingCardsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the billing cards delete params
func (o *BillingCardsDeleteParams) WithContext(ctx context.Context) *BillingCardsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the billing cards delete params
func (o *BillingCardsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the billing cards delete params
func (o *BillingCardsDeleteParams) WithHTTPClient(client *http.Client) *BillingCardsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the billing cards delete params
func (o *BillingCardsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the billing cards delete params
func (o *BillingCardsDeleteParams) WithID(id string) *BillingCardsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the billing cards delete params
func (o *BillingCardsDeleteParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the billing cards delete params
func (o *BillingCardsDeleteParams) WithNamespace(namespace string) *BillingCardsDeleteParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the billing cards delete params
func (o *BillingCardsDeleteParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *BillingCardsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
