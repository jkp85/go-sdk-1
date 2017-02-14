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

// NewProjectsDataSourcesStopParams creates a new ProjectsDataSourcesStopParams object
// with the default values initialized.
func NewProjectsDataSourcesStopParams() *ProjectsDataSourcesStopParams {
	var ()
	return &ProjectsDataSourcesStopParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsDataSourcesStopParamsWithTimeout creates a new ProjectsDataSourcesStopParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsDataSourcesStopParamsWithTimeout(timeout time.Duration) *ProjectsDataSourcesStopParams {
	var ()
	return &ProjectsDataSourcesStopParams{

		timeout: timeout,
	}
}

// NewProjectsDataSourcesStopParamsWithContext creates a new ProjectsDataSourcesStopParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsDataSourcesStopParamsWithContext(ctx context.Context) *ProjectsDataSourcesStopParams {
	var ()
	return &ProjectsDataSourcesStopParams{

		Context: ctx,
	}
}

// NewProjectsDataSourcesStopParamsWithHTTPClient creates a new ProjectsDataSourcesStopParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsDataSourcesStopParamsWithHTTPClient(client *http.Client) *ProjectsDataSourcesStopParams {
	var ()
	return &ProjectsDataSourcesStopParams{
		HTTPClient: client,
	}
}

/*ProjectsDataSourcesStopParams contains all the parameters to send to the API endpoint
for the projects data sources stop operation typically these are written to a http.Request
*/
type ProjectsDataSourcesStopParams struct {

	/*Data*/
	Data ProjectsDataSourcesStopBody
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

// WithTimeout adds the timeout to the projects data sources stop params
func (o *ProjectsDataSourcesStopParams) WithTimeout(timeout time.Duration) *ProjectsDataSourcesStopParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects data sources stop params
func (o *ProjectsDataSourcesStopParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects data sources stop params
func (o *ProjectsDataSourcesStopParams) WithContext(ctx context.Context) *ProjectsDataSourcesStopParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects data sources stop params
func (o *ProjectsDataSourcesStopParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects data sources stop params
func (o *ProjectsDataSourcesStopParams) WithHTTPClient(client *http.Client) *ProjectsDataSourcesStopParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects data sources stop params
func (o *ProjectsDataSourcesStopParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the projects data sources stop params
func (o *ProjectsDataSourcesStopParams) WithData(data ProjectsDataSourcesStopBody) *ProjectsDataSourcesStopParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the projects data sources stop params
func (o *ProjectsDataSourcesStopParams) SetData(data ProjectsDataSourcesStopBody) {
	o.Data = data
}

// WithNamespace adds the namespace to the projects data sources stop params
func (o *ProjectsDataSourcesStopParams) WithNamespace(namespace string) *ProjectsDataSourcesStopParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects data sources stop params
func (o *ProjectsDataSourcesStopParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectPk adds the projectPk to the projects data sources stop params
func (o *ProjectsDataSourcesStopParams) WithProjectPk(projectPk string) *ProjectsDataSourcesStopParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects data sources stop params
func (o *ProjectsDataSourcesStopParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WithServer adds the server to the projects data sources stop params
func (o *ProjectsDataSourcesStopParams) WithServer(server string) *ProjectsDataSourcesStopParams {
	o.SetServer(server)
	return o
}

// SetServer adds the server to the projects data sources stop params
func (o *ProjectsDataSourcesStopParams) SetServer(server string) {
	o.Server = server
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsDataSourcesStopParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
