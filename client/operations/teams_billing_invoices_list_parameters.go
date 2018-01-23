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

// NewTeamsBillingInvoicesListParams creates a new TeamsBillingInvoicesListParams object
// with the default values initialized.
func NewTeamsBillingInvoicesListParams() *TeamsBillingInvoicesListParams {
	var ()
	return &TeamsBillingInvoicesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTeamsBillingInvoicesListParamsWithTimeout creates a new TeamsBillingInvoicesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTeamsBillingInvoicesListParamsWithTimeout(timeout time.Duration) *TeamsBillingInvoicesListParams {
	var ()
	return &TeamsBillingInvoicesListParams{

		timeout: timeout,
	}
}

// NewTeamsBillingInvoicesListParamsWithContext creates a new TeamsBillingInvoicesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewTeamsBillingInvoicesListParamsWithContext(ctx context.Context) *TeamsBillingInvoicesListParams {
	var ()
	return &TeamsBillingInvoicesListParams{

		Context: ctx,
	}
}

// NewTeamsBillingInvoicesListParamsWithHTTPClient creates a new TeamsBillingInvoicesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTeamsBillingInvoicesListParamsWithHTTPClient(client *http.Client) *TeamsBillingInvoicesListParams {
	var ()
	return &TeamsBillingInvoicesListParams{
		HTTPClient: client,
	}
}

/*TeamsBillingInvoicesListParams contains all the parameters to send to the API endpoint
for the teams billing invoices list operation typically these are written to a http.Request
*/
type TeamsBillingInvoicesListParams struct {

	/*Limit
	  Limit when getting items.

	*/
	Limit *string
	/*Offset
	  Offset when getting items.

	*/
	Offset *string
	/*Team
	  Team unique identifier expressed as UUID or name.

	*/
	Team string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the teams billing invoices list params
func (o *TeamsBillingInvoicesListParams) WithTimeout(timeout time.Duration) *TeamsBillingInvoicesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the teams billing invoices list params
func (o *TeamsBillingInvoicesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the teams billing invoices list params
func (o *TeamsBillingInvoicesListParams) WithContext(ctx context.Context) *TeamsBillingInvoicesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the teams billing invoices list params
func (o *TeamsBillingInvoicesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the teams billing invoices list params
func (o *TeamsBillingInvoicesListParams) WithHTTPClient(client *http.Client) *TeamsBillingInvoicesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the teams billing invoices list params
func (o *TeamsBillingInvoicesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the teams billing invoices list params
func (o *TeamsBillingInvoicesListParams) WithLimit(limit *string) *TeamsBillingInvoicesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the teams billing invoices list params
func (o *TeamsBillingInvoicesListParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithOffset adds the offset to the teams billing invoices list params
func (o *TeamsBillingInvoicesListParams) WithOffset(offset *string) *TeamsBillingInvoicesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the teams billing invoices list params
func (o *TeamsBillingInvoicesListParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithTeam adds the team to the teams billing invoices list params
func (o *TeamsBillingInvoicesListParams) WithTeam(team string) *TeamsBillingInvoicesListParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the teams billing invoices list params
func (o *TeamsBillingInvoicesListParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *TeamsBillingInvoicesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
