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

	"github.com/3Blades/go-sdk/models"
)

// NewProjectsCollaboratorsUpdateParams creates a new ProjectsCollaboratorsUpdateParams object
// with the default values initialized.
func NewProjectsCollaboratorsUpdateParams() *ProjectsCollaboratorsUpdateParams {
	var ()
	return &ProjectsCollaboratorsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsCollaboratorsUpdateParamsWithTimeout creates a new ProjectsCollaboratorsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsCollaboratorsUpdateParamsWithTimeout(timeout time.Duration) *ProjectsCollaboratorsUpdateParams {
	var ()
	return &ProjectsCollaboratorsUpdateParams{

		timeout: timeout,
	}
}

// NewProjectsCollaboratorsUpdateParamsWithContext creates a new ProjectsCollaboratorsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsCollaboratorsUpdateParamsWithContext(ctx context.Context) *ProjectsCollaboratorsUpdateParams {
	var ()
	return &ProjectsCollaboratorsUpdateParams{

		Context: ctx,
	}
}

// NewProjectsCollaboratorsUpdateParamsWithHTTPClient creates a new ProjectsCollaboratorsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsCollaboratorsUpdateParamsWithHTTPClient(client *http.Client) *ProjectsCollaboratorsUpdateParams {
	var ()
	return &ProjectsCollaboratorsUpdateParams{
		HTTPClient: client,
	}
}

/*ProjectsCollaboratorsUpdateParams contains all the parameters to send to the API endpoint
for the projects collaborators update operation typically these are written to a http.Request
*/
type ProjectsCollaboratorsUpdateParams struct {

	/*CollaboratorData*/
	CollaboratorData *models.CollaboratorData
	/*ID*/
	ID string
	/*Namespace
	  User or team name.

	*/
	Namespace string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects collaborators update params
func (o *ProjectsCollaboratorsUpdateParams) WithTimeout(timeout time.Duration) *ProjectsCollaboratorsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects collaborators update params
func (o *ProjectsCollaboratorsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects collaborators update params
func (o *ProjectsCollaboratorsUpdateParams) WithContext(ctx context.Context) *ProjectsCollaboratorsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects collaborators update params
func (o *ProjectsCollaboratorsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects collaborators update params
func (o *ProjectsCollaboratorsUpdateParams) WithHTTPClient(client *http.Client) *ProjectsCollaboratorsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects collaborators update params
func (o *ProjectsCollaboratorsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCollaboratorData adds the collaboratorData to the projects collaborators update params
func (o *ProjectsCollaboratorsUpdateParams) WithCollaboratorData(collaboratorData *models.CollaboratorData) *ProjectsCollaboratorsUpdateParams {
	o.SetCollaboratorData(collaboratorData)
	return o
}

// SetCollaboratorData adds the collaboratorData to the projects collaborators update params
func (o *ProjectsCollaboratorsUpdateParams) SetCollaboratorData(collaboratorData *models.CollaboratorData) {
	o.CollaboratorData = collaboratorData
}

// WithID adds the id to the projects collaborators update params
func (o *ProjectsCollaboratorsUpdateParams) WithID(id string) *ProjectsCollaboratorsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects collaborators update params
func (o *ProjectsCollaboratorsUpdateParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the projects collaborators update params
func (o *ProjectsCollaboratorsUpdateParams) WithNamespace(namespace string) *ProjectsCollaboratorsUpdateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects collaborators update params
func (o *ProjectsCollaboratorsUpdateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectID adds the projectID to the projects collaborators update params
func (o *ProjectsCollaboratorsUpdateParams) WithProjectID(projectID string) *ProjectsCollaboratorsUpdateParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the projects collaborators update params
func (o *ProjectsCollaboratorsUpdateParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsCollaboratorsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CollaboratorData == nil {
		o.CollaboratorData = new(models.CollaboratorData)
	}

	if err := r.SetBodyParam(o.CollaboratorData); err != nil {
		return err
	}

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
