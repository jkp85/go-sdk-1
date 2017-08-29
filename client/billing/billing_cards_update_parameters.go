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

// NewBillingCardsUpdateParams creates a new BillingCardsUpdateParams object
// with the default values initialized.
func NewBillingCardsUpdateParams() *BillingCardsUpdateParams {
	var ()
	return &BillingCardsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBillingCardsUpdateParamsWithTimeout creates a new BillingCardsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBillingCardsUpdateParamsWithTimeout(timeout time.Duration) *BillingCardsUpdateParams {
	var ()
	return &BillingCardsUpdateParams{

		timeout: timeout,
	}
}

// NewBillingCardsUpdateParamsWithContext creates a new BillingCardsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewBillingCardsUpdateParamsWithContext(ctx context.Context) *BillingCardsUpdateParams {
	var ()
	return &BillingCardsUpdateParams{

		Context: ctx,
	}
}

// NewBillingCardsUpdateParamsWithHTTPClient creates a new BillingCardsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBillingCardsUpdateParamsWithHTTPClient(client *http.Client) *BillingCardsUpdateParams {
	var ()
	return &BillingCardsUpdateParams{
		HTTPClient: client,
	}
}

/*BillingCardsUpdateParams contains all the parameters to send to the API endpoint
for the billing cards update operation typically these are written to a http.Request
*/
type BillingCardsUpdateParams struct {

	/*CardData*/
	CardData *models.CardDataPutandPatch
	/*ID
	  Card unique identifier.

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

// WithTimeout adds the timeout to the billing cards update params
func (o *BillingCardsUpdateParams) WithTimeout(timeout time.Duration) *BillingCardsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the billing cards update params
func (o *BillingCardsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the billing cards update params
func (o *BillingCardsUpdateParams) WithContext(ctx context.Context) *BillingCardsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the billing cards update params
func (o *BillingCardsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the billing cards update params
func (o *BillingCardsUpdateParams) WithHTTPClient(client *http.Client) *BillingCardsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the billing cards update params
func (o *BillingCardsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCardData adds the cardData to the billing cards update params
func (o *BillingCardsUpdateParams) WithCardData(cardData *models.CardDataPutandPatch) *BillingCardsUpdateParams {
	o.SetCardData(cardData)
	return o
}

// SetCardData adds the cardData to the billing cards update params
func (o *BillingCardsUpdateParams) SetCardData(cardData *models.CardDataPutandPatch) {
	o.CardData = cardData
}

// WithID adds the id to the billing cards update params
func (o *BillingCardsUpdateParams) WithID(id string) *BillingCardsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the billing cards update params
func (o *BillingCardsUpdateParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the billing cards update params
func (o *BillingCardsUpdateParams) WithNamespace(namespace string) *BillingCardsUpdateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the billing cards update params
func (o *BillingCardsUpdateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *BillingCardsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CardData == nil {
		o.CardData = new(models.CardDataPutandPatch)
	}

	if err := r.SetBodyParam(o.CardData); err != nil {
		return err
	}

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
