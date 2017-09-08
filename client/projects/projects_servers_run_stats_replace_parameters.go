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

// NewProjectsServersRunStatsReplaceParams creates a new ProjectsServersRunStatsReplaceParams object
// with the default values initialized.
func NewProjectsServersRunStatsReplaceParams() *ProjectsServersRunStatsReplaceParams {
	var ()
	return &ProjectsServersRunStatsReplaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsServersRunStatsReplaceParamsWithTimeout creates a new ProjectsServersRunStatsReplaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsServersRunStatsReplaceParamsWithTimeout(timeout time.Duration) *ProjectsServersRunStatsReplaceParams {
	var ()
	return &ProjectsServersRunStatsReplaceParams{

		timeout: timeout,
	}
}

// NewProjectsServersRunStatsReplaceParamsWithContext creates a new ProjectsServersRunStatsReplaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsServersRunStatsReplaceParamsWithContext(ctx context.Context) *ProjectsServersRunStatsReplaceParams {
	var ()
	return &ProjectsServersRunStatsReplaceParams{

		Context: ctx,
	}
}

// NewProjectsServersRunStatsReplaceParamsWithHTTPClient creates a new ProjectsServersRunStatsReplaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsServersRunStatsReplaceParamsWithHTTPClient(client *http.Client) *ProjectsServersRunStatsReplaceParams {
	var ()
	return &ProjectsServersRunStatsReplaceParams{
		HTTPClient: client,
	}
}

/*ProjectsServersRunStatsReplaceParams contains all the parameters to send to the API endpoint
for the projects servers run stats replace operation typically these are written to a http.Request
*/
type ProjectsServersRunStatsReplaceParams struct {

	/*ID
	  Server run statistics expressed as UUID.

	*/
	ID string
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
	/*ServerrunstatsData*/
	ServerrunstatsData *models.ServerRunStatisticsData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects servers run stats replace params
func (o *ProjectsServersRunStatsReplaceParams) WithTimeout(timeout time.Duration) *ProjectsServersRunStatsReplaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects servers run stats replace params
func (o *ProjectsServersRunStatsReplaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects servers run stats replace params
func (o *ProjectsServersRunStatsReplaceParams) WithContext(ctx context.Context) *ProjectsServersRunStatsReplaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects servers run stats replace params
func (o *ProjectsServersRunStatsReplaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects servers run stats replace params
func (o *ProjectsServersRunStatsReplaceParams) WithHTTPClient(client *http.Client) *ProjectsServersRunStatsReplaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects servers run stats replace params
func (o *ProjectsServersRunStatsReplaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the projects servers run stats replace params
func (o *ProjectsServersRunStatsReplaceParams) WithID(id string) *ProjectsServersRunStatsReplaceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects servers run stats replace params
func (o *ProjectsServersRunStatsReplaceParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the projects servers run stats replace params
func (o *ProjectsServersRunStatsReplaceParams) WithNamespace(namespace string) *ProjectsServersRunStatsReplaceParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects servers run stats replace params
func (o *ProjectsServersRunStatsReplaceParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProject adds the project to the projects servers run stats replace params
func (o *ProjectsServersRunStatsReplaceParams) WithProject(project string) *ProjectsServersRunStatsReplaceParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the projects servers run stats replace params
func (o *ProjectsServersRunStatsReplaceParams) SetProject(project string) {
	o.Project = project
}

// WithServer adds the server to the projects servers run stats replace params
func (o *ProjectsServersRunStatsReplaceParams) WithServer(server string) *ProjectsServersRunStatsReplaceParams {
	o.SetServer(server)
	return o
}

// SetServer adds the server to the projects servers run stats replace params
func (o *ProjectsServersRunStatsReplaceParams) SetServer(server string) {
	o.Server = server
}

// WithServerrunstatsData adds the serverrunstatsData to the projects servers run stats replace params
func (o *ProjectsServersRunStatsReplaceParams) WithServerrunstatsData(serverrunstatsData *models.ServerRunStatisticsData) *ProjectsServersRunStatsReplaceParams {
	o.SetServerrunstatsData(serverrunstatsData)
	return o
}

// SetServerrunstatsData adds the serverrunstatsData to the projects servers run stats replace params
func (o *ProjectsServersRunStatsReplaceParams) SetServerrunstatsData(serverrunstatsData *models.ServerRunStatisticsData) {
	o.ServerrunstatsData = serverrunstatsData
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsServersRunStatsReplaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	// path param server
	if err := r.SetPathParam("server", o.Server); err != nil {
		return err
	}

	if o.ServerrunstatsData == nil {
		o.ServerrunstatsData = new(models.ServerRunStatisticsData)
	}

	if err := r.SetBodyParam(o.ServerrunstatsData); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
