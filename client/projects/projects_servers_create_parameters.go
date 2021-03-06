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

	models "github.com/IllumiDesk/go-sdk/models"
)

// NewProjectsServersCreateParams creates a new ProjectsServersCreateParams object
// with the default values initialized.
func NewProjectsServersCreateParams() *ProjectsServersCreateParams {
	var ()
	return &ProjectsServersCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsServersCreateParamsWithTimeout creates a new ProjectsServersCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsServersCreateParamsWithTimeout(timeout time.Duration) *ProjectsServersCreateParams {
	var ()
	return &ProjectsServersCreateParams{

		timeout: timeout,
	}
}

// NewProjectsServersCreateParamsWithContext creates a new ProjectsServersCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsServersCreateParamsWithContext(ctx context.Context) *ProjectsServersCreateParams {
	var ()
	return &ProjectsServersCreateParams{

		Context: ctx,
	}
}

// NewProjectsServersCreateParamsWithHTTPClient creates a new ProjectsServersCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsServersCreateParamsWithHTTPClient(client *http.Client) *ProjectsServersCreateParams {
	var ()
	return &ProjectsServersCreateParams{
		HTTPClient: client,
	}
}

/*ProjectsServersCreateParams contains all the parameters to send to the API endpoint
for the projects servers create operation typically these are written to a http.Request
*/
type ProjectsServersCreateParams struct {

	/*Namespace
	  User or team name.

	*/
	Namespace string
	/*Project
	  Project unique identifer expressed as UUID or name.

	*/
	Project string
	/*ServerData*/
	ServerData *models.ServerData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects servers create params
func (o *ProjectsServersCreateParams) WithTimeout(timeout time.Duration) *ProjectsServersCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects servers create params
func (o *ProjectsServersCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects servers create params
func (o *ProjectsServersCreateParams) WithContext(ctx context.Context) *ProjectsServersCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects servers create params
func (o *ProjectsServersCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects servers create params
func (o *ProjectsServersCreateParams) WithHTTPClient(client *http.Client) *ProjectsServersCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects servers create params
func (o *ProjectsServersCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the projects servers create params
func (o *ProjectsServersCreateParams) WithNamespace(namespace string) *ProjectsServersCreateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects servers create params
func (o *ProjectsServersCreateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProject adds the project to the projects servers create params
func (o *ProjectsServersCreateParams) WithProject(project string) *ProjectsServersCreateParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the projects servers create params
func (o *ProjectsServersCreateParams) SetProject(project string) {
	o.Project = project
}

// WithServerData adds the serverData to the projects servers create params
func (o *ProjectsServersCreateParams) WithServerData(serverData *models.ServerData) *ProjectsServersCreateParams {
	o.SetServerData(serverData)
	return o
}

// SetServerData adds the serverData to the projects servers create params
func (o *ProjectsServersCreateParams) SetServerData(serverData *models.ServerData) {
	o.ServerData = serverData
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsServersCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ServerData != nil {
		if err := r.SetBodyParam(o.ServerData); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
