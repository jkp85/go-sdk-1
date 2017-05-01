package hosts

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

// NewHostsListParams creates a new HostsListParams object
// with the default values initialized.
func NewHostsListParams() *HostsListParams {
	var ()
	return &HostsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHostsListParamsWithTimeout creates a new HostsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHostsListParamsWithTimeout(timeout time.Duration) *HostsListParams {
	var ()
	return &HostsListParams{

		timeout: timeout,
	}
}

// NewHostsListParamsWithContext creates a new HostsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewHostsListParamsWithContext(ctx context.Context) *HostsListParams {
	var ()
	return &HostsListParams{

		Context: ctx,
	}
}

// NewHostsListParamsWithHTTPClient creates a new HostsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHostsListParamsWithHTTPClient(client *http.Client) *HostsListParams {
	var ()
	return &HostsListParams{
		HTTPClient: client,
	}
}

/*HostsListParams contains all the parameters to send to the API endpoint
for the hosts list operation typically these are written to a http.Request
*/
type HostsListParams struct {

	/*Limit*/
	Limit *string
	/*Name*/
	Name *string
	/*Namespace*/
	Namespace string
	/*Offset*/
	Offset *string
	/*Ordering*/
	Ordering *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the hosts list params
func (o *HostsListParams) WithTimeout(timeout time.Duration) *HostsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hosts list params
func (o *HostsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hosts list params
func (o *HostsListParams) WithContext(ctx context.Context) *HostsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hosts list params
func (o *HostsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hosts list params
func (o *HostsListParams) WithHTTPClient(client *http.Client) *HostsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hosts list params
func (o *HostsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the hosts list params
func (o *HostsListParams) WithLimit(limit *string) *HostsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the hosts list params
func (o *HostsListParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithName adds the name to the hosts list params
func (o *HostsListParams) WithName(name *string) *HostsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the hosts list params
func (o *HostsListParams) SetName(name *string) {
	o.Name = name
}

// WithNamespace adds the namespace to the hosts list params
func (o *HostsListParams) WithNamespace(namespace string) *HostsListParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the hosts list params
func (o *HostsListParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the hosts list params
func (o *HostsListParams) WithOffset(offset *string) *HostsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the hosts list params
func (o *HostsListParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithOrdering adds the ordering to the hosts list params
func (o *HostsListParams) WithOrdering(ordering *string) *HostsListParams {
	o.SetOrdering(ordering)
	return o
}

// SetOrdering adds the ordering to the hosts list params
func (o *HostsListParams) SetOrdering(ordering *string) {
	o.Ordering = ordering
}

// WriteToRequest writes these params to a swagger request
func (o *HostsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
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

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
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