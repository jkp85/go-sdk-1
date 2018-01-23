// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// NewTeamsListParams creates a new TeamsListParams object
// with the default values initialized.
func NewTeamsListParams() *TeamsListParams {
	var ()
	return &TeamsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTeamsListParamsWithTimeout creates a new TeamsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTeamsListParamsWithTimeout(timeout time.Duration) *TeamsListParams {
	var ()
	return &TeamsListParams{

		timeout: timeout,
	}
}

// NewTeamsListParamsWithContext creates a new TeamsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewTeamsListParamsWithContext(ctx context.Context) *TeamsListParams {
	var ()
	return &TeamsListParams{

		Context: ctx,
	}
}

// NewTeamsListParamsWithHTTPClient creates a new TeamsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTeamsListParamsWithHTTPClient(client *http.Client) *TeamsListParams {
	var ()
	return &TeamsListParams{
		HTTPClient: client,
	}
}

/*TeamsListParams contains all the parameters to send to the API endpoint
for the teams list operation typically these are written to a http.Request
*/
type TeamsListParams struct {

	/*Limit
	  Limit when getting data.

	*/
	Limit *string
	/*Offset
	  Offset when getting data.

	*/
	Offset *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the teams list params
func (o *TeamsListParams) WithTimeout(timeout time.Duration) *TeamsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the teams list params
func (o *TeamsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the teams list params
func (o *TeamsListParams) WithContext(ctx context.Context) *TeamsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the teams list params
func (o *TeamsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the teams list params
func (o *TeamsListParams) WithHTTPClient(client *http.Client) *TeamsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the teams list params
func (o *TeamsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the teams list params
func (o *TeamsListParams) WithLimit(limit *string) *TeamsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the teams list params
func (o *TeamsListParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithOffset adds the offset to the teams list params
func (o *TeamsListParams) WithOffset(offset *string) *TeamsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the teams list params
func (o *TeamsListParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *TeamsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
