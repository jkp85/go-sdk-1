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

// NewProjectsWorkspacesReadParams creates a new ProjectsWorkspacesReadParams object
// with the default values initialized.
func NewProjectsWorkspacesReadParams() *ProjectsWorkspacesReadParams {
	var ()
	return &ProjectsWorkspacesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsWorkspacesReadParamsWithTimeout creates a new ProjectsWorkspacesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsWorkspacesReadParamsWithTimeout(timeout time.Duration) *ProjectsWorkspacesReadParams {
	var ()
	return &ProjectsWorkspacesReadParams{

		timeout: timeout,
	}
}

// NewProjectsWorkspacesReadParamsWithContext creates a new ProjectsWorkspacesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsWorkspacesReadParamsWithContext(ctx context.Context) *ProjectsWorkspacesReadParams {
	var ()
	return &ProjectsWorkspacesReadParams{

		Context: ctx,
	}
}

// NewProjectsWorkspacesReadParamsWithHTTPClient creates a new ProjectsWorkspacesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsWorkspacesReadParamsWithHTTPClient(client *http.Client) *ProjectsWorkspacesReadParams {
	var ()
	return &ProjectsWorkspacesReadParams{
		HTTPClient: client,
	}
}

/*ProjectsWorkspacesReadParams contains all the parameters to send to the API endpoint
for the projects workspaces read operation typically these are written to a http.Request
*/
type ProjectsWorkspacesReadParams struct {

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

// WithTimeout adds the timeout to the projects workspaces read params
func (o *ProjectsWorkspacesReadParams) WithTimeout(timeout time.Duration) *ProjectsWorkspacesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects workspaces read params
func (o *ProjectsWorkspacesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects workspaces read params
func (o *ProjectsWorkspacesReadParams) WithContext(ctx context.Context) *ProjectsWorkspacesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects workspaces read params
func (o *ProjectsWorkspacesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects workspaces read params
func (o *ProjectsWorkspacesReadParams) WithHTTPClient(client *http.Client) *ProjectsWorkspacesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects workspaces read params
func (o *ProjectsWorkspacesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the projects workspaces read params
func (o *ProjectsWorkspacesReadParams) WithNamespace(namespace string) *ProjectsWorkspacesReadParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects workspaces read params
func (o *ProjectsWorkspacesReadParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectPk adds the projectPk to the projects workspaces read params
func (o *ProjectsWorkspacesReadParams) WithProjectPk(projectPk string) *ProjectsWorkspacesReadParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects workspaces read params
func (o *ProjectsWorkspacesReadParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WithServer adds the server to the projects workspaces read params
func (o *ProjectsWorkspacesReadParams) WithServer(server string) *ProjectsWorkspacesReadParams {
	o.SetServer(server)
	return o
}

// SetServer adds the server to the projects workspaces read params
func (o *ProjectsWorkspacesReadParams) SetServer(server string) {
	o.Server = server
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsWorkspacesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

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