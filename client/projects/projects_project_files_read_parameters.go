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

// NewProjectsProjectFilesReadParams creates a new ProjectsProjectFilesReadParams object
// with the default values initialized.
func NewProjectsProjectFilesReadParams() *ProjectsProjectFilesReadParams {
	var ()
	return &ProjectsProjectFilesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsProjectFilesReadParamsWithTimeout creates a new ProjectsProjectFilesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsProjectFilesReadParamsWithTimeout(timeout time.Duration) *ProjectsProjectFilesReadParams {
	var ()
	return &ProjectsProjectFilesReadParams{

		timeout: timeout,
	}
}

// NewProjectsProjectFilesReadParamsWithContext creates a new ProjectsProjectFilesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsProjectFilesReadParamsWithContext(ctx context.Context) *ProjectsProjectFilesReadParams {
	var ()
	return &ProjectsProjectFilesReadParams{

		Context: ctx,
	}
}

// NewProjectsProjectFilesReadParamsWithHTTPClient creates a new ProjectsProjectFilesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsProjectFilesReadParamsWithHTTPClient(client *http.Client) *ProjectsProjectFilesReadParams {
	var ()
	return &ProjectsProjectFilesReadParams{
		HTTPClient: client,
	}
}

/*ProjectsProjectFilesReadParams contains all the parameters to send to the API endpoint
for the projects project files read operation typically these are written to a http.Request
*/
type ProjectsProjectFilesReadParams struct {

	/*Content
	  Determines whether or not content is returned as base64. Defaults to false.

	*/
	Content *string
	/*ID
	  File unique identifier.

	*/
	ID string
	/*Namespace
	  User or team name.

	*/
	Namespace string
	/*Project
	  Project unique identifer.

	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects project files read params
func (o *ProjectsProjectFilesReadParams) WithTimeout(timeout time.Duration) *ProjectsProjectFilesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects project files read params
func (o *ProjectsProjectFilesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects project files read params
func (o *ProjectsProjectFilesReadParams) WithContext(ctx context.Context) *ProjectsProjectFilesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects project files read params
func (o *ProjectsProjectFilesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects project files read params
func (o *ProjectsProjectFilesReadParams) WithHTTPClient(client *http.Client) *ProjectsProjectFilesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects project files read params
func (o *ProjectsProjectFilesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContent adds the content to the projects project files read params
func (o *ProjectsProjectFilesReadParams) WithContent(content *string) *ProjectsProjectFilesReadParams {
	o.SetContent(content)
	return o
}

// SetContent adds the content to the projects project files read params
func (o *ProjectsProjectFilesReadParams) SetContent(content *string) {
	o.Content = content
}

// WithID adds the id to the projects project files read params
func (o *ProjectsProjectFilesReadParams) WithID(id string) *ProjectsProjectFilesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects project files read params
func (o *ProjectsProjectFilesReadParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the projects project files read params
func (o *ProjectsProjectFilesReadParams) WithNamespace(namespace string) *ProjectsProjectFilesReadParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects project files read params
func (o *ProjectsProjectFilesReadParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProject adds the project to the projects project files read params
func (o *ProjectsProjectFilesReadParams) WithProject(project string) *ProjectsProjectFilesReadParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the projects project files read params
func (o *ProjectsProjectFilesReadParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsProjectFilesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Content != nil {

		// query param content
		var qrContent string
		if o.Content != nil {
			qrContent = *o.Content
		}
		qContent := qrContent
		if qContent != "" {
			if err := r.SetQueryParam("content", qContent); err != nil {
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

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
