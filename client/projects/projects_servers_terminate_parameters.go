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

// NewProjectsServersTerminateParams creates a new ProjectsServersTerminateParams object
// with the default values initialized.
func NewProjectsServersTerminateParams() *ProjectsServersTerminateParams {
	var ()
	return &ProjectsServersTerminateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsServersTerminateParamsWithTimeout creates a new ProjectsServersTerminateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsServersTerminateParamsWithTimeout(timeout time.Duration) *ProjectsServersTerminateParams {
	var ()
	return &ProjectsServersTerminateParams{

		timeout: timeout,
	}
}

// NewProjectsServersTerminateParamsWithContext creates a new ProjectsServersTerminateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsServersTerminateParamsWithContext(ctx context.Context) *ProjectsServersTerminateParams {
	var ()
	return &ProjectsServersTerminateParams{

		Context: ctx,
	}
}

// NewProjectsServersTerminateParamsWithHTTPClient creates a new ProjectsServersTerminateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsServersTerminateParamsWithHTTPClient(client *http.Client) *ProjectsServersTerminateParams {
	var ()
	return &ProjectsServersTerminateParams{
		HTTPClient: client,
	}
}

/*ProjectsServersTerminateParams contains all the parameters to send to the API endpoint
for the projects servers terminate operation typically these are written to a http.Request
*/
type ProjectsServersTerminateParams struct {

	/*ID
	  Server identifier expressed as UUID.

	*/
	ID string
	/*Namespace
	  User or team name.

	*/
	Namespace string
	/*ProjectID
	  Project identifier expressed as UUID.

	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects servers terminate params
func (o *ProjectsServersTerminateParams) WithTimeout(timeout time.Duration) *ProjectsServersTerminateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects servers terminate params
func (o *ProjectsServersTerminateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects servers terminate params
func (o *ProjectsServersTerminateParams) WithContext(ctx context.Context) *ProjectsServersTerminateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects servers terminate params
func (o *ProjectsServersTerminateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects servers terminate params
func (o *ProjectsServersTerminateParams) WithHTTPClient(client *http.Client) *ProjectsServersTerminateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects servers terminate params
func (o *ProjectsServersTerminateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the projects servers terminate params
func (o *ProjectsServersTerminateParams) WithID(id string) *ProjectsServersTerminateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects servers terminate params
func (o *ProjectsServersTerminateParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the projects servers terminate params
func (o *ProjectsServersTerminateParams) WithNamespace(namespace string) *ProjectsServersTerminateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects servers terminate params
func (o *ProjectsServersTerminateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectID adds the projectID to the projects servers terminate params
func (o *ProjectsServersTerminateParams) WithProjectID(projectID string) *ProjectsServersTerminateParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the projects servers terminate params
func (o *ProjectsServersTerminateParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsServersTerminateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
