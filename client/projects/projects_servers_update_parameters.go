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

// NewProjectsServersUpdateParams creates a new ProjectsServersUpdateParams object
// with the default values initialized.
func NewProjectsServersUpdateParams() *ProjectsServersUpdateParams {
	var ()
	return &ProjectsServersUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsServersUpdateParamsWithTimeout creates a new ProjectsServersUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsServersUpdateParamsWithTimeout(timeout time.Duration) *ProjectsServersUpdateParams {
	var ()
	return &ProjectsServersUpdateParams{

		timeout: timeout,
	}
}

// NewProjectsServersUpdateParamsWithContext creates a new ProjectsServersUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsServersUpdateParamsWithContext(ctx context.Context) *ProjectsServersUpdateParams {
	var ()
	return &ProjectsServersUpdateParams{

		Context: ctx,
	}
}

// NewProjectsServersUpdateParamsWithHTTPClient creates a new ProjectsServersUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsServersUpdateParamsWithHTTPClient(client *http.Client) *ProjectsServersUpdateParams {
	var ()
	return &ProjectsServersUpdateParams{
		HTTPClient: client,
	}
}

/*ProjectsServersUpdateParams contains all the parameters to send to the API endpoint
for the projects servers update operation typically these are written to a http.Request
*/
type ProjectsServersUpdateParams struct {

	/*Namespace
	  User or team name.

	*/
	Namespace string
	/*Project
	  Project unique identifier expressed as UUID or name.

	*/
	Project string
	/*Server
	  Server unique identifier expressed as UUID or name.

	*/
	Server string
	/*ServerData*/
	ServerData *models.ServerData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects servers update params
func (o *ProjectsServersUpdateParams) WithTimeout(timeout time.Duration) *ProjectsServersUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects servers update params
func (o *ProjectsServersUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects servers update params
func (o *ProjectsServersUpdateParams) WithContext(ctx context.Context) *ProjectsServersUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects servers update params
func (o *ProjectsServersUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects servers update params
func (o *ProjectsServersUpdateParams) WithHTTPClient(client *http.Client) *ProjectsServersUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects servers update params
func (o *ProjectsServersUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the projects servers update params
func (o *ProjectsServersUpdateParams) WithNamespace(namespace string) *ProjectsServersUpdateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects servers update params
func (o *ProjectsServersUpdateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProject adds the project to the projects servers update params
func (o *ProjectsServersUpdateParams) WithProject(project string) *ProjectsServersUpdateParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the projects servers update params
func (o *ProjectsServersUpdateParams) SetProject(project string) {
	o.Project = project
}

// WithServer adds the server to the projects servers update params
func (o *ProjectsServersUpdateParams) WithServer(server string) *ProjectsServersUpdateParams {
	o.SetServer(server)
	return o
}

// SetServer adds the server to the projects servers update params
func (o *ProjectsServersUpdateParams) SetServer(server string) {
	o.Server = server
}

// WithServerData adds the serverData to the projects servers update params
func (o *ProjectsServersUpdateParams) WithServerData(serverData *models.ServerData) *ProjectsServersUpdateParams {
	o.SetServerData(serverData)
	return o
}

// SetServerData adds the serverData to the projects servers update params
func (o *ProjectsServersUpdateParams) SetServerData(serverData *models.ServerData) {
	o.ServerData = serverData
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsServersUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ServerData == nil {
		o.ServerData = new(models.ServerData)
	}

	if err := r.SetBodyParam(o.ServerData); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
