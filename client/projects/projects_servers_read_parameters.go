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

// NewProjectsServersReadParams creates a new ProjectsServersReadParams object
// with the default values initialized.
func NewProjectsServersReadParams() *ProjectsServersReadParams {
	var ()
	return &ProjectsServersReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsServersReadParamsWithTimeout creates a new ProjectsServersReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsServersReadParamsWithTimeout(timeout time.Duration) *ProjectsServersReadParams {
	var ()
	return &ProjectsServersReadParams{

		timeout: timeout,
	}
}

// NewProjectsServersReadParamsWithContext creates a new ProjectsServersReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsServersReadParamsWithContext(ctx context.Context) *ProjectsServersReadParams {
	var ()
	return &ProjectsServersReadParams{

		Context: ctx,
	}
}

// NewProjectsServersReadParamsWithHTTPClient creates a new ProjectsServersReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsServersReadParamsWithHTTPClient(client *http.Client) *ProjectsServersReadParams {
	var ()
	return &ProjectsServersReadParams{
		HTTPClient: client,
	}
}

/*ProjectsServersReadParams contains all the parameters to send to the API endpoint
for the projects servers read operation typically these are written to a http.Request
*/
type ProjectsServersReadParams struct {

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

// WithTimeout adds the timeout to the projects servers read params
func (o *ProjectsServersReadParams) WithTimeout(timeout time.Duration) *ProjectsServersReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects servers read params
func (o *ProjectsServersReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects servers read params
func (o *ProjectsServersReadParams) WithContext(ctx context.Context) *ProjectsServersReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects servers read params
func (o *ProjectsServersReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects servers read params
func (o *ProjectsServersReadParams) WithHTTPClient(client *http.Client) *ProjectsServersReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects servers read params
func (o *ProjectsServersReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the projects servers read params
func (o *ProjectsServersReadParams) WithID(id string) *ProjectsServersReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects servers read params
func (o *ProjectsServersReadParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the projects servers read params
func (o *ProjectsServersReadParams) WithNamespace(namespace string) *ProjectsServersReadParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects servers read params
func (o *ProjectsServersReadParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectPk adds the projectPk to the projects servers read params
func (o *ProjectsServersReadParams) WithProjectPk(projectPk string) *ProjectsServersReadParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects servers read params
func (o *ProjectsServersReadParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsServersReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

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