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

// NewProjectsJobsStartParams creates a new ProjectsJobsStartParams object
// with the default values initialized.
func NewProjectsJobsStartParams() *ProjectsJobsStartParams {
	var ()
	return &ProjectsJobsStartParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsJobsStartParamsWithTimeout creates a new ProjectsJobsStartParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsJobsStartParamsWithTimeout(timeout time.Duration) *ProjectsJobsStartParams {
	var ()
	return &ProjectsJobsStartParams{

		timeout: timeout,
	}
}

// NewProjectsJobsStartParamsWithContext creates a new ProjectsJobsStartParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsJobsStartParamsWithContext(ctx context.Context) *ProjectsJobsStartParams {
	var ()
	return &ProjectsJobsStartParams{

		Context: ctx,
	}
}

// NewProjectsJobsStartParamsWithHTTPClient creates a new ProjectsJobsStartParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsJobsStartParamsWithHTTPClient(client *http.Client) *ProjectsJobsStartParams {
	var ()
	return &ProjectsJobsStartParams{
		HTTPClient: client,
	}
}

/*ProjectsJobsStartParams contains all the parameters to send to the API endpoint
for the projects jobs start operation typically these are written to a http.Request
*/
type ProjectsJobsStartParams struct {

	/*Data*/
	Data ProjectsJobsStartBody
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

// WithTimeout adds the timeout to the projects jobs start params
func (o *ProjectsJobsStartParams) WithTimeout(timeout time.Duration) *ProjectsJobsStartParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects jobs start params
func (o *ProjectsJobsStartParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects jobs start params
func (o *ProjectsJobsStartParams) WithContext(ctx context.Context) *ProjectsJobsStartParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects jobs start params
func (o *ProjectsJobsStartParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects jobs start params
func (o *ProjectsJobsStartParams) WithHTTPClient(client *http.Client) *ProjectsJobsStartParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects jobs start params
func (o *ProjectsJobsStartParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the projects jobs start params
func (o *ProjectsJobsStartParams) WithData(data ProjectsJobsStartBody) *ProjectsJobsStartParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the projects jobs start params
func (o *ProjectsJobsStartParams) SetData(data ProjectsJobsStartBody) {
	o.Data = data
}

// WithNamespace adds the namespace to the projects jobs start params
func (o *ProjectsJobsStartParams) WithNamespace(namespace string) *ProjectsJobsStartParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects jobs start params
func (o *ProjectsJobsStartParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectPk adds the projectPk to the projects jobs start params
func (o *ProjectsJobsStartParams) WithProjectPk(projectPk string) *ProjectsJobsStartParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects jobs start params
func (o *ProjectsJobsStartParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WithServer adds the server to the projects jobs start params
func (o *ProjectsJobsStartParams) WithServer(server string) *ProjectsJobsStartParams {
	o.SetServer(server)
	return o
}

// SetServer adds the server to the projects jobs start params
func (o *ProjectsJobsStartParams) SetServer(server string) {
	o.Server = server
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsJobsStartParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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