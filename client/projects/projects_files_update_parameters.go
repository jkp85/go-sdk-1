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

// NewProjectsFilesUpdateParams creates a new ProjectsFilesUpdateParams object
// with the default values initialized.
func NewProjectsFilesUpdateParams() *ProjectsFilesUpdateParams {
	var ()
	return &ProjectsFilesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsFilesUpdateParamsWithTimeout creates a new ProjectsFilesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsFilesUpdateParamsWithTimeout(timeout time.Duration) *ProjectsFilesUpdateParams {
	var ()
	return &ProjectsFilesUpdateParams{

		timeout: timeout,
	}
}

// NewProjectsFilesUpdateParamsWithContext creates a new ProjectsFilesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsFilesUpdateParamsWithContext(ctx context.Context) *ProjectsFilesUpdateParams {
	var ()
	return &ProjectsFilesUpdateParams{

		Context: ctx,
	}
}

// NewProjectsFilesUpdateParamsWithHTTPClient creates a new ProjectsFilesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsFilesUpdateParamsWithHTTPClient(client *http.Client) *ProjectsFilesUpdateParams {
	var ()
	return &ProjectsFilesUpdateParams{
		HTTPClient: client,
	}
}

/*ProjectsFilesUpdateParams contains all the parameters to send to the API endpoint
for the projects files update operation typically these are written to a http.Request
*/
type ProjectsFilesUpdateParams struct {

	/*Data*/
	Data ProjectsFilesUpdateBody
	/*ID*/
	ID string
	/*Namespace*/
	Namespace string
	/*ProjectPk*/
	ProjectPk string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects files update params
func (o *ProjectsFilesUpdateParams) WithTimeout(timeout time.Duration) *ProjectsFilesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects files update params
func (o *ProjectsFilesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects files update params
func (o *ProjectsFilesUpdateParams) WithContext(ctx context.Context) *ProjectsFilesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects files update params
func (o *ProjectsFilesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects files update params
func (o *ProjectsFilesUpdateParams) WithHTTPClient(client *http.Client) *ProjectsFilesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects files update params
func (o *ProjectsFilesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the projects files update params
func (o *ProjectsFilesUpdateParams) WithData(data ProjectsFilesUpdateBody) *ProjectsFilesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the projects files update params
func (o *ProjectsFilesUpdateParams) SetData(data ProjectsFilesUpdateBody) {
	o.Data = data
}

// WithID adds the id to the projects files update params
func (o *ProjectsFilesUpdateParams) WithID(id string) *ProjectsFilesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects files update params
func (o *ProjectsFilesUpdateParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the projects files update params
func (o *ProjectsFilesUpdateParams) WithNamespace(namespace string) *ProjectsFilesUpdateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects files update params
func (o *ProjectsFilesUpdateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectPk adds the projectPk to the projects files update params
func (o *ProjectsFilesUpdateParams) WithProjectPk(projectPk string) *ProjectsFilesUpdateParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects files update params
func (o *ProjectsFilesUpdateParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsFilesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Data); err != nil {
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

	// path param project_pk
	if err := r.SetPathParam("project_pk", o.ProjectPk); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
