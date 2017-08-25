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

// NewProjectsServersStatsUpdateParams creates a new ProjectsServersStatsUpdateParams object
// with the default values initialized.
func NewProjectsServersStatsUpdateParams() *ProjectsServersStatsUpdateParams {
	var ()
	return &ProjectsServersStatsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsServersStatsUpdateParamsWithTimeout creates a new ProjectsServersStatsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsServersStatsUpdateParamsWithTimeout(timeout time.Duration) *ProjectsServersStatsUpdateParams {
	var ()
	return &ProjectsServersStatsUpdateParams{

		timeout: timeout,
	}
}

// NewProjectsServersStatsUpdateParamsWithContext creates a new ProjectsServersStatsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsServersStatsUpdateParamsWithContext(ctx context.Context) *ProjectsServersStatsUpdateParams {
	var ()
	return &ProjectsServersStatsUpdateParams{

		Context: ctx,
	}
}

// NewProjectsServersStatsUpdateParamsWithHTTPClient creates a new ProjectsServersStatsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsServersStatsUpdateParamsWithHTTPClient(client *http.Client) *ProjectsServersStatsUpdateParams {
	var ()
	return &ProjectsServersStatsUpdateParams{
		HTTPClient: client,
	}
}

/*ProjectsServersStatsUpdateParams contains all the parameters to send to the API endpoint
for the projects servers stats update operation typically these are written to a http.Request
*/
type ProjectsServersStatsUpdateParams struct {

	/*ID
	  Server statistics unique identifier expressed as UUID.

	*/
	ID string
	/*Namespace
	  User or team name.

	*/
	Namespace string
	/*ProjectID
	  Project unique identifier expressed as UUID.

	*/
	ProjectID string
	/*ServerID
	  Server unique identifier expressed as UUID.

	*/
	ServerID string
	/*ServerstatsData*/
	ServerstatsData *models.ServerStatisticsData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects servers stats update params
func (o *ProjectsServersStatsUpdateParams) WithTimeout(timeout time.Duration) *ProjectsServersStatsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects servers stats update params
func (o *ProjectsServersStatsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects servers stats update params
func (o *ProjectsServersStatsUpdateParams) WithContext(ctx context.Context) *ProjectsServersStatsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects servers stats update params
func (o *ProjectsServersStatsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects servers stats update params
func (o *ProjectsServersStatsUpdateParams) WithHTTPClient(client *http.Client) *ProjectsServersStatsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects servers stats update params
func (o *ProjectsServersStatsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the projects servers stats update params
func (o *ProjectsServersStatsUpdateParams) WithID(id string) *ProjectsServersStatsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects servers stats update params
func (o *ProjectsServersStatsUpdateParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the projects servers stats update params
func (o *ProjectsServersStatsUpdateParams) WithNamespace(namespace string) *ProjectsServersStatsUpdateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects servers stats update params
func (o *ProjectsServersStatsUpdateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectID adds the projectID to the projects servers stats update params
func (o *ProjectsServersStatsUpdateParams) WithProjectID(projectID string) *ProjectsServersStatsUpdateParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the projects servers stats update params
func (o *ProjectsServersStatsUpdateParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithServerID adds the serverID to the projects servers stats update params
func (o *ProjectsServersStatsUpdateParams) WithServerID(serverID string) *ProjectsServersStatsUpdateParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the projects servers stats update params
func (o *ProjectsServersStatsUpdateParams) SetServerID(serverID string) {
	o.ServerID = serverID
}

// WithServerstatsData adds the serverstatsData to the projects servers stats update params
func (o *ProjectsServersStatsUpdateParams) WithServerstatsData(serverstatsData *models.ServerStatisticsData) *ProjectsServersStatsUpdateParams {
	o.SetServerstatsData(serverstatsData)
	return o
}

// SetServerstatsData adds the serverstatsData to the projects servers stats update params
func (o *ProjectsServersStatsUpdateParams) SetServerstatsData(serverstatsData *models.ServerStatisticsData) {
	o.ServerstatsData = serverstatsData
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsServersStatsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param server_id
	if err := r.SetPathParam("server_id", o.ServerID); err != nil {
		return err
	}

	if o.ServerstatsData == nil {
		o.ServerstatsData = new(models.ServerStatisticsData)
	}

	if err := r.SetBodyParam(o.ServerstatsData); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
