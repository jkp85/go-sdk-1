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

// NewProjectsDataSourcesStartParams creates a new ProjectsDataSourcesStartParams object
// with the default values initialized.
func NewProjectsDataSourcesStartParams() *ProjectsDataSourcesStartParams {
	var ()
	return &ProjectsDataSourcesStartParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsDataSourcesStartParamsWithTimeout creates a new ProjectsDataSourcesStartParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsDataSourcesStartParamsWithTimeout(timeout time.Duration) *ProjectsDataSourcesStartParams {
	var ()
	return &ProjectsDataSourcesStartParams{

		timeout: timeout,
	}
}

// NewProjectsDataSourcesStartParamsWithContext creates a new ProjectsDataSourcesStartParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsDataSourcesStartParamsWithContext(ctx context.Context) *ProjectsDataSourcesStartParams {
	var ()
	return &ProjectsDataSourcesStartParams{

		Context: ctx,
	}
}

// NewProjectsDataSourcesStartParamsWithHTTPClient creates a new ProjectsDataSourcesStartParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsDataSourcesStartParamsWithHTTPClient(client *http.Client) *ProjectsDataSourcesStartParams {
	var ()
	return &ProjectsDataSourcesStartParams{
		HTTPClient: client,
	}
}

/*ProjectsDataSourcesStartParams contains all the parameters to send to the API endpoint
for the projects data sources start operation typically these are written to a http.Request
*/
type ProjectsDataSourcesStartParams struct {

	/*Data*/
	Data ProjectsDataSourcesStartBody
	/*Namespace*/
	Namespace string
	/*ProjectPk*/
	ProjectPk string
	/*Server*/
	Server string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects data sources start params
func (o *ProjectsDataSourcesStartParams) WithTimeout(timeout time.Duration) *ProjectsDataSourcesStartParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects data sources start params
func (o *ProjectsDataSourcesStartParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects data sources start params
func (o *ProjectsDataSourcesStartParams) WithContext(ctx context.Context) *ProjectsDataSourcesStartParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects data sources start params
func (o *ProjectsDataSourcesStartParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects data sources start params
func (o *ProjectsDataSourcesStartParams) WithHTTPClient(client *http.Client) *ProjectsDataSourcesStartParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects data sources start params
func (o *ProjectsDataSourcesStartParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the projects data sources start params
func (o *ProjectsDataSourcesStartParams) WithData(data ProjectsDataSourcesStartBody) *ProjectsDataSourcesStartParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the projects data sources start params
func (o *ProjectsDataSourcesStartParams) SetData(data ProjectsDataSourcesStartBody) {
	o.Data = data
}

// WithNamespace adds the namespace to the projects data sources start params
func (o *ProjectsDataSourcesStartParams) WithNamespace(namespace string) *ProjectsDataSourcesStartParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects data sources start params
func (o *ProjectsDataSourcesStartParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectPk adds the projectPk to the projects data sources start params
func (o *ProjectsDataSourcesStartParams) WithProjectPk(projectPk string) *ProjectsDataSourcesStartParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects data sources start params
func (o *ProjectsDataSourcesStartParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WithServer adds the server to the projects data sources start params
func (o *ProjectsDataSourcesStartParams) WithServer(server string) *ProjectsDataSourcesStartParams {
	o.SetServer(server)
	return o
}

// SetServer adds the server to the projects data sources start params
func (o *ProjectsDataSourcesStartParams) SetServer(server string) {
	o.Server = server
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsDataSourcesStartParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Data); err != nil {
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

	// path param server
	if err := r.SetPathParam("server", o.Server); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}