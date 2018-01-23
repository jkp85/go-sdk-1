// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewTeamsBillingInvoiceItemsListParams creates a new TeamsBillingInvoiceItemsListParams object
// with the default values initialized.
func NewTeamsBillingInvoiceItemsListParams() *TeamsBillingInvoiceItemsListParams {
	var ()
	return &TeamsBillingInvoiceItemsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTeamsBillingInvoiceItemsListParamsWithTimeout creates a new TeamsBillingInvoiceItemsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTeamsBillingInvoiceItemsListParamsWithTimeout(timeout time.Duration) *TeamsBillingInvoiceItemsListParams {
	var ()
	return &TeamsBillingInvoiceItemsListParams{

		timeout: timeout,
	}
}

// NewTeamsBillingInvoiceItemsListParamsWithContext creates a new TeamsBillingInvoiceItemsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewTeamsBillingInvoiceItemsListParamsWithContext(ctx context.Context) *TeamsBillingInvoiceItemsListParams {
	var ()
	return &TeamsBillingInvoiceItemsListParams{

		Context: ctx,
	}
}

// NewTeamsBillingInvoiceItemsListParamsWithHTTPClient creates a new TeamsBillingInvoiceItemsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTeamsBillingInvoiceItemsListParamsWithHTTPClient(client *http.Client) *TeamsBillingInvoiceItemsListParams {
	var ()
	return &TeamsBillingInvoiceItemsListParams{
		HTTPClient: client,
	}
}

/*TeamsBillingInvoiceItemsListParams contains all the parameters to send to the API endpoint
for the teams billing invoice items list operation typically these are written to a http.Request
*/
type TeamsBillingInvoiceItemsListParams struct {

	/*InvoiceID
	  Invoice id, expressed as UUID.

	*/
	InvoiceID string
	/*Limit
	  Limit when getting items.

	*/
	Limit *string
	/*Offset
	  Offset when getting items.

	*/
	Offset *string
	/*Ordering
	  Ordering when getting items.

	*/
	Ordering *string
	/*Team
	  Team unique identifier expressed as UUID or name.

	*/
	Team string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the teams billing invoice items list params
func (o *TeamsBillingInvoiceItemsListParams) WithTimeout(timeout time.Duration) *TeamsBillingInvoiceItemsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the teams billing invoice items list params
func (o *TeamsBillingInvoiceItemsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the teams billing invoice items list params
func (o *TeamsBillingInvoiceItemsListParams) WithContext(ctx context.Context) *TeamsBillingInvoiceItemsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the teams billing invoice items list params
func (o *TeamsBillingInvoiceItemsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the teams billing invoice items list params
func (o *TeamsBillingInvoiceItemsListParams) WithHTTPClient(client *http.Client) *TeamsBillingInvoiceItemsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the teams billing invoice items list params
func (o *TeamsBillingInvoiceItemsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvoiceID adds the invoiceID to the teams billing invoice items list params
func (o *TeamsBillingInvoiceItemsListParams) WithInvoiceID(invoiceID string) *TeamsBillingInvoiceItemsListParams {
	o.SetInvoiceID(invoiceID)
	return o
}

// SetInvoiceID adds the invoiceId to the teams billing invoice items list params
func (o *TeamsBillingInvoiceItemsListParams) SetInvoiceID(invoiceID string) {
	o.InvoiceID = invoiceID
}

// WithLimit adds the limit to the teams billing invoice items list params
func (o *TeamsBillingInvoiceItemsListParams) WithLimit(limit *string) *TeamsBillingInvoiceItemsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the teams billing invoice items list params
func (o *TeamsBillingInvoiceItemsListParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithOffset adds the offset to the teams billing invoice items list params
func (o *TeamsBillingInvoiceItemsListParams) WithOffset(offset *string) *TeamsBillingInvoiceItemsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the teams billing invoice items list params
func (o *TeamsBillingInvoiceItemsListParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithOrdering adds the ordering to the teams billing invoice items list params
func (o *TeamsBillingInvoiceItemsListParams) WithOrdering(ordering *string) *TeamsBillingInvoiceItemsListParams {
	o.SetOrdering(ordering)
	return o
}

// SetOrdering adds the ordering to the teams billing invoice items list params
func (o *TeamsBillingInvoiceItemsListParams) SetOrdering(ordering *string) {
	o.Ordering = ordering
}

// WithTeam adds the team to the teams billing invoice items list params
func (o *TeamsBillingInvoiceItemsListParams) WithTeam(team string) *TeamsBillingInvoiceItemsListParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the teams billing invoice items list params
func (o *TeamsBillingInvoiceItemsListParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *TeamsBillingInvoiceItemsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param invoice_id
	if err := r.SetPathParam("invoice_id", o.InvoiceID); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit string
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Ordering != nil {

		// query param ordering
		var qrOrdering string
		if o.Ordering != nil {
			qrOrdering = *o.Ordering
		}
		qOrdering := qrOrdering
		if qOrdering != "" {
			if err := r.SetQueryParam("ordering", qOrdering); err != nil {
				return err
			}
		}

	}

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
