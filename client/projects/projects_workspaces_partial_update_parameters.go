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

// NewProjectsWorkspacesPartialUpdateParams creates a new ProjectsWorkspacesPartialUpdateParams object
// with the default values initialized.
func NewProjectsWorkspacesPartialUpdateParams() *ProjectsWorkspacesPartialUpdateParams {
	var ()
	return &ProjectsWorkspacesPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsWorkspacesPartialUpdateParamsWithTimeout creates a new ProjectsWorkspacesPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsWorkspacesPartialUpdateParamsWithTimeout(timeout time.Duration) *ProjectsWorkspacesPartialUpdateParams {
	var ()
	return &ProjectsWorkspacesPartialUpdateParams{

		timeout: timeout,
	}
}

// NewProjectsWorkspacesPartialUpdateParamsWithContext creates a new ProjectsWorkspacesPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsWorkspacesPartialUpdateParamsWithContext(ctx context.Context) *ProjectsWorkspacesPartialUpdateParams {
	var ()
	return &ProjectsWorkspacesPartialUpdateParams{

		Context: ctx,
	}
}

// NewProjectsWorkspacesPartialUpdateParamsWithHTTPClient creates a new ProjectsWorkspacesPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsWorkspacesPartialUpdateParamsWithHTTPClient(client *http.Client) *ProjectsWorkspacesPartialUpdateParams {
	var ()
	return &ProjectsWorkspacesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*ProjectsWorkspacesPartialUpdateParams contains all the parameters to send to the API endpoint
for the projects workspaces partial update operation typically these are written to a http.Request
*/
type ProjectsWorkspacesPartialUpdateParams struct {

	/*Data*/
	Data ProjectsWorkspacesPartialUpdateBody
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

// WithTimeout adds the timeout to the projects workspaces partial update params
func (o *ProjectsWorkspacesPartialUpdateParams) WithTimeout(timeout time.Duration) *ProjectsWorkspacesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects workspaces partial update params
func (o *ProjectsWorkspacesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects workspaces partial update params
func (o *ProjectsWorkspacesPartialUpdateParams) WithContext(ctx context.Context) *ProjectsWorkspacesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects workspaces partial update params
func (o *ProjectsWorkspacesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects workspaces partial update params
func (o *ProjectsWorkspacesPartialUpdateParams) WithHTTPClient(client *http.Client) *ProjectsWorkspacesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects workspaces partial update params
func (o *ProjectsWorkspacesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the projects workspaces partial update params
func (o *ProjectsWorkspacesPartialUpdateParams) WithData(data ProjectsWorkspacesPartialUpdateBody) *ProjectsWorkspacesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the projects workspaces partial update params
func (o *ProjectsWorkspacesPartialUpdateParams) SetData(data ProjectsWorkspacesPartialUpdateBody) {
	o.Data = data
}

// WithNamespace adds the namespace to the projects workspaces partial update params
func (o *ProjectsWorkspacesPartialUpdateParams) WithNamespace(namespace string) *ProjectsWorkspacesPartialUpdateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects workspaces partial update params
func (o *ProjectsWorkspacesPartialUpdateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectPk adds the projectPk to the projects workspaces partial update params
func (o *ProjectsWorkspacesPartialUpdateParams) WithProjectPk(projectPk string) *ProjectsWorkspacesPartialUpdateParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects workspaces partial update params
func (o *ProjectsWorkspacesPartialUpdateParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WithServer adds the server to the projects workspaces partial update params
func (o *ProjectsWorkspacesPartialUpdateParams) WithServer(server string) *ProjectsWorkspacesPartialUpdateParams {
	o.SetServer(server)
	return o
}

// SetServer adds the server to the projects workspaces partial update params
func (o *ProjectsWorkspacesPartialUpdateParams) SetServer(server string) {
	o.Server = server
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsWorkspacesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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