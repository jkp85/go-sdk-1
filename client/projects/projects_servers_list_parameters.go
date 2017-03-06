package projects

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

// NewProjectsServersListParams creates a new ProjectsServersListParams object
// with the default values initialized.
func NewProjectsServersListParams() *ProjectsServersListParams {
	var ()
	return &ProjectsServersListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsServersListParamsWithTimeout creates a new ProjectsServersListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsServersListParamsWithTimeout(timeout time.Duration) *ProjectsServersListParams {
	var ()
	return &ProjectsServersListParams{

		timeout: timeout,
	}
}

// NewProjectsServersListParamsWithContext creates a new ProjectsServersListParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsServersListParamsWithContext(ctx context.Context) *ProjectsServersListParams {
	var ()
	return &ProjectsServersListParams{

		Context: ctx,
	}
}

// NewProjectsServersListParamsWithHTTPClient creates a new ProjectsServersListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsServersListParamsWithHTTPClient(client *http.Client) *ProjectsServersListParams {
	var ()
	return &ProjectsServersListParams{
		HTTPClient: client,
	}
}

/*ProjectsServersListParams contains all the parameters to send to the API endpoint
for the projects servers list operation typically these are written to a http.Request
*/
type ProjectsServersListParams struct {

	/*Limit*/
	Limit *string
	/*Namespace*/
	Namespace string
	/*Offset*/
	Offset *string
	/*ProjectPk*/
	ProjectPk string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects servers list params
func (o *ProjectsServersListParams) WithTimeout(timeout time.Duration) *ProjectsServersListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects servers list params
func (o *ProjectsServersListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects servers list params
func (o *ProjectsServersListParams) WithContext(ctx context.Context) *ProjectsServersListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects servers list params
func (o *ProjectsServersListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects servers list params
func (o *ProjectsServersListParams) WithHTTPClient(client *http.Client) *ProjectsServersListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects servers list params
func (o *ProjectsServersListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the projects servers list params
func (o *ProjectsServersListParams) WithLimit(limit *string) *ProjectsServersListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the projects servers list params
func (o *ProjectsServersListParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the projects servers list params
func (o *ProjectsServersListParams) WithNamespace(namespace string) *ProjectsServersListParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects servers list params
func (o *ProjectsServersListParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the projects servers list params
func (o *ProjectsServersListParams) WithOffset(offset *string) *ProjectsServersListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the projects servers list params
func (o *ProjectsServersListParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithProjectPk adds the projectPk to the projects servers list params
func (o *ProjectsServersListParams) WithProjectPk(projectPk string) *ProjectsServersListParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects servers list params
func (o *ProjectsServersListParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsServersListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param project_pk
	if err := r.SetPathParam("project_pk", o.ProjectPk); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
