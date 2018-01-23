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

// NewProjectsUpdateParams creates a new ProjectsUpdateParams object
// with the default values initialized.
func NewProjectsUpdateParams() *ProjectsUpdateParams {
	var ()
	return &ProjectsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsUpdateParamsWithTimeout creates a new ProjectsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsUpdateParamsWithTimeout(timeout time.Duration) *ProjectsUpdateParams {
	var ()
	return &ProjectsUpdateParams{

		timeout: timeout,
	}
}

// NewProjectsUpdateParamsWithContext creates a new ProjectsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsUpdateParamsWithContext(ctx context.Context) *ProjectsUpdateParams {
	var ()
	return &ProjectsUpdateParams{

		Context: ctx,
	}
}

// NewProjectsUpdateParamsWithHTTPClient creates a new ProjectsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsUpdateParamsWithHTTPClient(client *http.Client) *ProjectsUpdateParams {
	var ()
	return &ProjectsUpdateParams{
		HTTPClient: client,
	}
}

/*ProjectsUpdateParams contains all the parameters to send to the API endpoint
for the projects update operation typically these are written to a http.Request
*/
type ProjectsUpdateParams struct {

	/*Namespace
	  User or team name.

	*/
	Namespace string
	/*Project
	  Project unique identifier expressed as UUID or name.

	*/
	Project string
	/*ProjectData*/
	ProjectData *models.ProjectData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects update params
func (o *ProjectsUpdateParams) WithTimeout(timeout time.Duration) *ProjectsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects update params
func (o *ProjectsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects update params
func (o *ProjectsUpdateParams) WithContext(ctx context.Context) *ProjectsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects update params
func (o *ProjectsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects update params
func (o *ProjectsUpdateParams) WithHTTPClient(client *http.Client) *ProjectsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects update params
func (o *ProjectsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the projects update params
func (o *ProjectsUpdateParams) WithNamespace(namespace string) *ProjectsUpdateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects update params
func (o *ProjectsUpdateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProject adds the project to the projects update params
func (o *ProjectsUpdateParams) WithProject(project string) *ProjectsUpdateParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the projects update params
func (o *ProjectsUpdateParams) SetProject(project string) {
	o.Project = project
}

// WithProjectData adds the projectData to the projects update params
func (o *ProjectsUpdateParams) WithProjectData(projectData *models.ProjectData) *ProjectsUpdateParams {
	o.SetProjectData(projectData)
	return o
}

// SetProjectData adds the projectData to the projects update params
func (o *ProjectsUpdateParams) SetProjectData(projectData *models.ProjectData) {
	o.ProjectData = projectData
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ProjectData != nil {
		if err := r.SetBodyParam(o.ProjectData); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
