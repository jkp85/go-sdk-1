// Code generated by go-swagger; DO NOT EDIT.

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

// NewProjectsServersDeleteParams creates a new ProjectsServersDeleteParams object
// with the default values initialized.
func NewProjectsServersDeleteParams() *ProjectsServersDeleteParams {
	var ()
	return &ProjectsServersDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsServersDeleteParamsWithTimeout creates a new ProjectsServersDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsServersDeleteParamsWithTimeout(timeout time.Duration) *ProjectsServersDeleteParams {
	var ()
	return &ProjectsServersDeleteParams{

		timeout: timeout,
	}
}

// NewProjectsServersDeleteParamsWithContext creates a new ProjectsServersDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsServersDeleteParamsWithContext(ctx context.Context) *ProjectsServersDeleteParams {
	var ()
	return &ProjectsServersDeleteParams{

		Context: ctx,
	}
}

// NewProjectsServersDeleteParamsWithHTTPClient creates a new ProjectsServersDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsServersDeleteParamsWithHTTPClient(client *http.Client) *ProjectsServersDeleteParams {
	var ()
	return &ProjectsServersDeleteParams{
		HTTPClient: client,
	}
}

/*ProjectsServersDeleteParams contains all the parameters to send to the API endpoint
for the projects servers delete operation typically these are written to a http.Request
*/
type ProjectsServersDeleteParams struct {

	/*Namespace
	  User or team name.

	*/
	Namespace string
	/*Project
	  Project unique identifier.

	*/
	Project string
	/*Server
	  User unique identifier.

	*/
	Server string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects servers delete params
func (o *ProjectsServersDeleteParams) WithTimeout(timeout time.Duration) *ProjectsServersDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects servers delete params
func (o *ProjectsServersDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects servers delete params
func (o *ProjectsServersDeleteParams) WithContext(ctx context.Context) *ProjectsServersDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects servers delete params
func (o *ProjectsServersDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects servers delete params
func (o *ProjectsServersDeleteParams) WithHTTPClient(client *http.Client) *ProjectsServersDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects servers delete params
func (o *ProjectsServersDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the projects servers delete params
func (o *ProjectsServersDeleteParams) WithNamespace(namespace string) *ProjectsServersDeleteParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects servers delete params
func (o *ProjectsServersDeleteParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProject adds the project to the projects servers delete params
func (o *ProjectsServersDeleteParams) WithProject(project string) *ProjectsServersDeleteParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the projects servers delete params
func (o *ProjectsServersDeleteParams) SetProject(project string) {
	o.Project = project
}

// WithServer adds the server to the projects servers delete params
func (o *ProjectsServersDeleteParams) WithServer(server string) *ProjectsServersDeleteParams {
	o.SetServer(server)
	return o
}

// SetServer adds the server to the projects servers delete params
func (o *ProjectsServersDeleteParams) SetServer(server string) {
	o.Server = server
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsServersDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
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
