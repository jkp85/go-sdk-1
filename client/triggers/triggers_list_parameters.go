// Code generated by go-swagger; DO NOT EDIT.

package triggers

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

// NewTriggersListParams creates a new TriggersListParams object
// with the default values initialized.
func NewTriggersListParams() *TriggersListParams {
	var ()
	return &TriggersListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTriggersListParamsWithTimeout creates a new TriggersListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTriggersListParamsWithTimeout(timeout time.Duration) *TriggersListParams {
	var ()
	return &TriggersListParams{

		timeout: timeout,
	}
}

// NewTriggersListParamsWithContext creates a new TriggersListParams object
// with the default values initialized, and the ability to set a context for a request
func NewTriggersListParamsWithContext(ctx context.Context) *TriggersListParams {
	var ()
	return &TriggersListParams{

		Context: ctx,
	}
}

// NewTriggersListParamsWithHTTPClient creates a new TriggersListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTriggersListParamsWithHTTPClient(client *http.Client) *TriggersListParams {
	var ()
	return &TriggersListParams{
		HTTPClient: client,
	}
}

/*TriggersListParams contains all the parameters to send to the API endpoint
for the triggers list operation typically these are written to a http.Request
*/
type TriggersListParams struct {

	/*Limit
	  Limit values returned.

	*/
	Limit *string
	/*Namespace
	  User or team name.

	*/
	Namespace string
	/*Offset
	  Offset list of values returned.

	*/
	Offset *string
	/*Ordering
	  Order list of values returned.

	*/
	Ordering *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the triggers list params
func (o *TriggersListParams) WithTimeout(timeout time.Duration) *TriggersListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the triggers list params
func (o *TriggersListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the triggers list params
func (o *TriggersListParams) WithContext(ctx context.Context) *TriggersListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the triggers list params
func (o *TriggersListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the triggers list params
func (o *TriggersListParams) WithHTTPClient(client *http.Client) *TriggersListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the triggers list params
func (o *TriggersListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the triggers list params
func (o *TriggersListParams) WithLimit(limit *string) *TriggersListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the triggers list params
func (o *TriggersListParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the triggers list params
func (o *TriggersListParams) WithNamespace(namespace string) *TriggersListParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the triggers list params
func (o *TriggersListParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the triggers list params
func (o *TriggersListParams) WithOffset(offset *string) *TriggersListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the triggers list params
func (o *TriggersListParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithOrdering adds the ordering to the triggers list params
func (o *TriggersListParams) WithOrdering(ordering *string) *TriggersListParams {
	o.SetOrdering(ordering)
	return o
}

// SetOrdering adds the ordering to the triggers list params
func (o *TriggersListParams) SetOrdering(ordering *string) {
	o.Ordering = ordering
}

// WriteToRequest writes these params to a swagger request
func (o *TriggersListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
