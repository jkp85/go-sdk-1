package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"os"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewProjectsProjectFilesPartialUpdateParams creates a new ProjectsProjectFilesPartialUpdateParams object
// with the default values initialized.
func NewProjectsProjectFilesPartialUpdateParams() *ProjectsProjectFilesPartialUpdateParams {
	var ()
	return &ProjectsProjectFilesPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsProjectFilesPartialUpdateParamsWithTimeout creates a new ProjectsProjectFilesPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsProjectFilesPartialUpdateParamsWithTimeout(timeout time.Duration) *ProjectsProjectFilesPartialUpdateParams {
	var ()
	return &ProjectsProjectFilesPartialUpdateParams{

		timeout: timeout,
	}
}

// NewProjectsProjectFilesPartialUpdateParamsWithContext creates a new ProjectsProjectFilesPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsProjectFilesPartialUpdateParamsWithContext(ctx context.Context) *ProjectsProjectFilesPartialUpdateParams {
	var ()
	return &ProjectsProjectFilesPartialUpdateParams{

		Context: ctx,
	}
}

// NewProjectsProjectFilesPartialUpdateParamsWithHTTPClient creates a new ProjectsProjectFilesPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsProjectFilesPartialUpdateParamsWithHTTPClient(client *http.Client) *ProjectsProjectFilesPartialUpdateParams {
	var ()
	return &ProjectsProjectFilesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*ProjectsProjectFilesPartialUpdateParams contains all the parameters to send to the API endpoint
for the projects project files partial update operation typically these are written to a http.Request
*/
type ProjectsProjectFilesPartialUpdateParams struct {

	/*File*/
	File *os.File
	/*ID*/
	ID string
	/*Namespace*/
	Namespace string
	/*Project*/
	Project *string
	/*ProjectPk*/
	ProjectPk string
	/*Public*/
	Public *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects project files partial update params
func (o *ProjectsProjectFilesPartialUpdateParams) WithTimeout(timeout time.Duration) *ProjectsProjectFilesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects project files partial update params
func (o *ProjectsProjectFilesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects project files partial update params
func (o *ProjectsProjectFilesPartialUpdateParams) WithContext(ctx context.Context) *ProjectsProjectFilesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects project files partial update params
func (o *ProjectsProjectFilesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects project files partial update params
func (o *ProjectsProjectFilesPartialUpdateParams) WithHTTPClient(client *http.Client) *ProjectsProjectFilesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects project files partial update params
func (o *ProjectsProjectFilesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFile adds the file to the projects project files partial update params
func (o *ProjectsProjectFilesPartialUpdateParams) WithFile(file *os.File) *ProjectsProjectFilesPartialUpdateParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the projects project files partial update params
func (o *ProjectsProjectFilesPartialUpdateParams) SetFile(file *os.File) {
	o.File = file
}

// WithID adds the id to the projects project files partial update params
func (o *ProjectsProjectFilesPartialUpdateParams) WithID(id string) *ProjectsProjectFilesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects project files partial update params
func (o *ProjectsProjectFilesPartialUpdateParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the projects project files partial update params
func (o *ProjectsProjectFilesPartialUpdateParams) WithNamespace(namespace string) *ProjectsProjectFilesPartialUpdateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects project files partial update params
func (o *ProjectsProjectFilesPartialUpdateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProject adds the project to the projects project files partial update params
func (o *ProjectsProjectFilesPartialUpdateParams) WithProject(project *string) *ProjectsProjectFilesPartialUpdateParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the projects project files partial update params
func (o *ProjectsProjectFilesPartialUpdateParams) SetProject(project *string) {
	o.Project = project
}

// WithProjectPk adds the projectPk to the projects project files partial update params
func (o *ProjectsProjectFilesPartialUpdateParams) WithProjectPk(projectPk string) *ProjectsProjectFilesPartialUpdateParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects project files partial update params
func (o *ProjectsProjectFilesPartialUpdateParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WithPublic adds the public to the projects project files partial update params
func (o *ProjectsProjectFilesPartialUpdateParams) WithPublic(public *bool) *ProjectsProjectFilesPartialUpdateParams {
	o.SetPublic(public)
	return o
}

// SetPublic adds the public to the projects project files partial update params
func (o *ProjectsProjectFilesPartialUpdateParams) SetPublic(public *bool) {
	o.Public = public
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsProjectFilesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.File != nil {

		if o.File != nil {

			// form file param file
			if err := r.SetFileParam("file", o.File); err != nil {
				return err
			}

		}

	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.Project != nil {

		// form param project
		var frProject string
		if o.Project != nil {
			frProject = *o.Project
		}
		fProject := frProject
		if fProject != "" {
			if err := r.SetFormParam("project", fProject); err != nil {
				return err
			}
		}

	}

	// path param project_pk
	if err := r.SetPathParam("project_pk", o.ProjectPk); err != nil {
		return err
	}

	if o.Public != nil {

		// form param public
		var frPublic bool
		if o.Public != nil {
			frPublic = *o.Public
		}
		fPublic := swag.FormatBool(frPublic)
		if fPublic != "" {
			if err := r.SetFormParam("public", fPublic); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}