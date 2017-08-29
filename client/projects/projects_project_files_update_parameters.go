// Code generated by go-swagger; DO NOT EDIT.

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

	strfmt "github.com/go-openapi/strfmt"
)

// NewProjectsProjectFilesUpdateParams creates a new ProjectsProjectFilesUpdateParams object
// with the default values initialized.
func NewProjectsProjectFilesUpdateParams() *ProjectsProjectFilesUpdateParams {
	var ()
	return &ProjectsProjectFilesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsProjectFilesUpdateParamsWithTimeout creates a new ProjectsProjectFilesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsProjectFilesUpdateParamsWithTimeout(timeout time.Duration) *ProjectsProjectFilesUpdateParams {
	var ()
	return &ProjectsProjectFilesUpdateParams{

		timeout: timeout,
	}
}

// NewProjectsProjectFilesUpdateParamsWithContext creates a new ProjectsProjectFilesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsProjectFilesUpdateParamsWithContext(ctx context.Context) *ProjectsProjectFilesUpdateParams {
	var ()
	return &ProjectsProjectFilesUpdateParams{

		Context: ctx,
	}
}

// NewProjectsProjectFilesUpdateParamsWithHTTPClient creates a new ProjectsProjectFilesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsProjectFilesUpdateParamsWithHTTPClient(client *http.Client) *ProjectsProjectFilesUpdateParams {
	var ()
	return &ProjectsProjectFilesUpdateParams{
		HTTPClient: client,
	}
}

/*ProjectsProjectFilesUpdateParams contains all the parameters to send to the API endpoint
for the projects project files update operation typically these are written to a http.Request
*/
type ProjectsProjectFilesUpdateParams struct {

	/*Base64Data
	  Fila data, represented as base64.

	*/
	Base64Data *string
	/*File
	  File to send, to create new file. This parameter is only used with form data and may include multiple files.

	*/
	File *os.File
	/*ID
	  File unique identifier.

	*/
	ID string
	/*Name
	  File name. May include path when creating file with base64 field.

	*/
	Name *string
	/*Namespace
	  User or team name.

	*/
	Namespace string
	/*Path
	  File path. Defaults to (/).

	*/
	Path *string
	/*ProjectID
	  Project unique identifer.

	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects project files update params
func (o *ProjectsProjectFilesUpdateParams) WithTimeout(timeout time.Duration) *ProjectsProjectFilesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects project files update params
func (o *ProjectsProjectFilesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects project files update params
func (o *ProjectsProjectFilesUpdateParams) WithContext(ctx context.Context) *ProjectsProjectFilesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects project files update params
func (o *ProjectsProjectFilesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects project files update params
func (o *ProjectsProjectFilesUpdateParams) WithHTTPClient(client *http.Client) *ProjectsProjectFilesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects project files update params
func (o *ProjectsProjectFilesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBase64Data adds the base64Data to the projects project files update params
func (o *ProjectsProjectFilesUpdateParams) WithBase64Data(base64Data *string) *ProjectsProjectFilesUpdateParams {
	o.SetBase64Data(base64Data)
	return o
}

// SetBase64Data adds the base64Data to the projects project files update params
func (o *ProjectsProjectFilesUpdateParams) SetBase64Data(base64Data *string) {
	o.Base64Data = base64Data
}

// WithFile adds the file to the projects project files update params
func (o *ProjectsProjectFilesUpdateParams) WithFile(file *os.File) *ProjectsProjectFilesUpdateParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the projects project files update params
func (o *ProjectsProjectFilesUpdateParams) SetFile(file *os.File) {
	o.File = file
}

// WithID adds the id to the projects project files update params
func (o *ProjectsProjectFilesUpdateParams) WithID(id string) *ProjectsProjectFilesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects project files update params
func (o *ProjectsProjectFilesUpdateParams) SetID(id string) {
	o.ID = id
}

// WithName adds the name to the projects project files update params
func (o *ProjectsProjectFilesUpdateParams) WithName(name *string) *ProjectsProjectFilesUpdateParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the projects project files update params
func (o *ProjectsProjectFilesUpdateParams) SetName(name *string) {
	o.Name = name
}

// WithNamespace adds the namespace to the projects project files update params
func (o *ProjectsProjectFilesUpdateParams) WithNamespace(namespace string) *ProjectsProjectFilesUpdateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects project files update params
func (o *ProjectsProjectFilesUpdateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPath adds the path to the projects project files update params
func (o *ProjectsProjectFilesUpdateParams) WithPath(path *string) *ProjectsProjectFilesUpdateParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the projects project files update params
func (o *ProjectsProjectFilesUpdateParams) SetPath(path *string) {
	o.Path = path
}

// WithProjectID adds the projectID to the projects project files update params
func (o *ProjectsProjectFilesUpdateParams) WithProjectID(projectID string) *ProjectsProjectFilesUpdateParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the projects project files update params
func (o *ProjectsProjectFilesUpdateParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsProjectFilesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Base64Data != nil {

		// form param base64_data
		var frBase64Data string
		if o.Base64Data != nil {
			frBase64Data = *o.Base64Data
		}
		fBase64Data := frBase64Data
		if fBase64Data != "" {
			if err := r.SetFormParam("base64_data", fBase64Data); err != nil {
				return err
			}
		}

	}

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

	if o.Name != nil {

		// form param name
		var frName string
		if o.Name != nil {
			frName = *o.Name
		}
		fName := frName
		if fName != "" {
			if err := r.SetFormParam("name", fName); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.Path != nil {

		// form param path
		var frPath string
		if o.Path != nil {
			frPath = *o.Path
		}
		fPath := frPath
		if fPath != "" {
			if err := r.SetFormParam("path", fPath); err != nil {
				return err
			}
		}

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
