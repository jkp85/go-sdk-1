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

	/*Data*/
	Data ProjectsServersStatsUpdateBody
	/*ID*/
	ID string
	/*Namespace*/
	Namespace string
	/*ProjectPk*/
	ProjectPk string
	/*ServerPk*/
	ServerPk string

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

// WithData adds the data to the projects servers stats update params
func (o *ProjectsServersStatsUpdateParams) WithData(data ProjectsServersStatsUpdateBody) *ProjectsServersStatsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the projects servers stats update params
func (o *ProjectsServersStatsUpdateParams) SetData(data ProjectsServersStatsUpdateBody) {
	o.Data = data
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

// WithProjectPk adds the projectPk to the projects servers stats update params
func (o *ProjectsServersStatsUpdateParams) WithProjectPk(projectPk string) *ProjectsServersStatsUpdateParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects servers stats update params
func (o *ProjectsServersStatsUpdateParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WithServerPk adds the serverPk to the projects servers stats update params
func (o *ProjectsServersStatsUpdateParams) WithServerPk(serverPk string) *ProjectsServersStatsUpdateParams {
	o.SetServerPk(serverPk)
	return o
}

// SetServerPk adds the serverPk to the projects servers stats update params
func (o *ProjectsServersStatsUpdateParams) SetServerPk(serverPk string) {
	o.ServerPk = serverPk
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsServersStatsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Data); err != nil {
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

	// path param project_pk
	if err := r.SetPathParam("project_pk", o.ProjectPk); err != nil {
		return err
	}

	// path param server_pk
	if err := r.SetPathParam("server_pk", o.ServerPk); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}