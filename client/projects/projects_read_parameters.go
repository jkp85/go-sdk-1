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

// NewProjectsReadParams creates a new ProjectsReadParams object
// with the default values initialized.
func NewProjectsReadParams() *ProjectsReadParams {
	var ()
	return &ProjectsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsReadParamsWithTimeout creates a new ProjectsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsReadParamsWithTimeout(timeout time.Duration) *ProjectsReadParams {
	var ()
	return &ProjectsReadParams{

		timeout: timeout,
	}
}

// NewProjectsReadParamsWithContext creates a new ProjectsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsReadParamsWithContext(ctx context.Context) *ProjectsReadParams {
	var ()
	return &ProjectsReadParams{

		Context: ctx,
	}
}

// NewProjectsReadParamsWithHTTPClient creates a new ProjectsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsReadParamsWithHTTPClient(client *http.Client) *ProjectsReadParams {
	var ()
	return &ProjectsReadParams{
		HTTPClient: client,
	}
}

/*ProjectsReadParams contains all the parameters to send to the API endpoint
for the projects read operation typically these are written to a http.Request
*/
type ProjectsReadParams struct {

	/*Namespace
	  User or team name.

	*/
	Namespace string
	/*Project
	  Project unique identifier expressed as UUID or name.

	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects read params
func (o *ProjectsReadParams) WithTimeout(timeout time.Duration) *ProjectsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects read params
func (o *ProjectsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects read params
func (o *ProjectsReadParams) WithContext(ctx context.Context) *ProjectsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects read params
func (o *ProjectsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects read params
func (o *ProjectsReadParams) WithHTTPClient(client *http.Client) *ProjectsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects read params
func (o *ProjectsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the projects read params
func (o *ProjectsReadParams) WithNamespace(namespace string) *ProjectsReadParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects read params
func (o *ProjectsReadParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProject adds the project to the projects read params
func (o *ProjectsReadParams) WithProject(project string) *ProjectsReadParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the projects read params
func (o *ProjectsReadParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
