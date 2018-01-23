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

// NewTeamsBillingSubscriptionsListParams creates a new TeamsBillingSubscriptionsListParams object
// with the default values initialized.
func NewTeamsBillingSubscriptionsListParams() *TeamsBillingSubscriptionsListParams {
	var ()
	return &TeamsBillingSubscriptionsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTeamsBillingSubscriptionsListParamsWithTimeout creates a new TeamsBillingSubscriptionsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTeamsBillingSubscriptionsListParamsWithTimeout(timeout time.Duration) *TeamsBillingSubscriptionsListParams {
	var ()
	return &TeamsBillingSubscriptionsListParams{

		timeout: timeout,
	}
}

// NewTeamsBillingSubscriptionsListParamsWithContext creates a new TeamsBillingSubscriptionsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewTeamsBillingSubscriptionsListParamsWithContext(ctx context.Context) *TeamsBillingSubscriptionsListParams {
	var ()
	return &TeamsBillingSubscriptionsListParams{

		Context: ctx,
	}
}

// NewTeamsBillingSubscriptionsListParamsWithHTTPClient creates a new TeamsBillingSubscriptionsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTeamsBillingSubscriptionsListParamsWithHTTPClient(client *http.Client) *TeamsBillingSubscriptionsListParams {
	var ()
	return &TeamsBillingSubscriptionsListParams{
		HTTPClient: client,
	}
}

/*TeamsBillingSubscriptionsListParams contains all the parameters to send to the API endpoint
for the teams billing subscriptions list operation typically these are written to a http.Request
*/
type TeamsBillingSubscriptionsListParams struct {

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

// WithTimeout adds the timeout to the teams billing subscriptions list params
func (o *TeamsBillingSubscriptionsListParams) WithTimeout(timeout time.Duration) *TeamsBillingSubscriptionsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the teams billing subscriptions list params
func (o *TeamsBillingSubscriptionsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the teams billing subscriptions list params
func (o *TeamsBillingSubscriptionsListParams) WithContext(ctx context.Context) *TeamsBillingSubscriptionsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the teams billing subscriptions list params
func (o *TeamsBillingSubscriptionsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the teams billing subscriptions list params
func (o *TeamsBillingSubscriptionsListParams) WithHTTPClient(client *http.Client) *TeamsBillingSubscriptionsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the teams billing subscriptions list params
func (o *TeamsBillingSubscriptionsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the teams billing subscriptions list params
func (o *TeamsBillingSubscriptionsListParams) WithLimit(limit *string) *TeamsBillingSubscriptionsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the teams billing subscriptions list params
func (o *TeamsBillingSubscriptionsListParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithOffset adds the offset to the teams billing subscriptions list params
func (o *TeamsBillingSubscriptionsListParams) WithOffset(offset *string) *TeamsBillingSubscriptionsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the teams billing subscriptions list params
func (o *TeamsBillingSubscriptionsListParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithOrdering adds the ordering to the teams billing subscriptions list params
func (o *TeamsBillingSubscriptionsListParams) WithOrdering(ordering *string) *TeamsBillingSubscriptionsListParams {
	o.SetOrdering(ordering)
	return o
}

// SetOrdering adds the ordering to the teams billing subscriptions list params
func (o *TeamsBillingSubscriptionsListParams) SetOrdering(ordering *string) {
	o.Ordering = ordering
}

// WithTeam adds the team to the teams billing subscriptions list params
func (o *TeamsBillingSubscriptionsListParams) WithTeam(team string) *TeamsBillingSubscriptionsListParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the teams billing subscriptions list params
func (o *TeamsBillingSubscriptionsListParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *TeamsBillingSubscriptionsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
