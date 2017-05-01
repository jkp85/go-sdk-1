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

// NewProjectsDataSourcesPartialUpdateParams creates a new ProjectsDataSourcesPartialUpdateParams object
// with the default values initialized.
func NewProjectsDataSourcesPartialUpdateParams() *ProjectsDataSourcesPartialUpdateParams {
	var ()
	return &ProjectsDataSourcesPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsDataSourcesPartialUpdateParamsWithTimeout creates a new ProjectsDataSourcesPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsDataSourcesPartialUpdateParamsWithTimeout(timeout time.Duration) *ProjectsDataSourcesPartialUpdateParams {
	var ()
	return &ProjectsDataSourcesPartialUpdateParams{

		timeout: timeout,
	}
}

// NewProjectsDataSourcesPartialUpdateParamsWithContext creates a new ProjectsDataSourcesPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsDataSourcesPartialUpdateParamsWithContext(ctx context.Context) *ProjectsDataSourcesPartialUpdateParams {
	var ()
	return &ProjectsDataSourcesPartialUpdateParams{

		Context: ctx,
	}
}

// NewProjectsDataSourcesPartialUpdateParamsWithHTTPClient creates a new ProjectsDataSourcesPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsDataSourcesPartialUpdateParamsWithHTTPClient(client *http.Client) *ProjectsDataSourcesPartialUpdateParams {
	var ()
	return &ProjectsDataSourcesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*ProjectsDataSourcesPartialUpdateParams contains all the parameters to send to the API endpoint
for the projects data sources partial update operation typically these are written to a http.Request
*/
type ProjectsDataSourcesPartialUpdateParams struct {

	/*Data*/
	Data ProjectsDataSourcesPartialUpdateBody
	/*Namespace*/
	Namespace string
	/*ProjectPk*/
	ProjectPk string
	/*Server*/
	Server string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects data sources partial update params
func (o *ProjectsDataSourcesPartialUpdateParams) WithTimeout(timeout time.Duration) *ProjectsDataSourcesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects data sources partial update params
func (o *ProjectsDataSourcesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects data sources partial update params
func (o *ProjectsDataSourcesPartialUpdateParams) WithContext(ctx context.Context) *ProjectsDataSourcesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects data sources partial update params
func (o *ProjectsDataSourcesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects data sources partial update params
func (o *ProjectsDataSourcesPartialUpdateParams) WithHTTPClient(client *http.Client) *ProjectsDataSourcesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects data sources partial update params
func (o *ProjectsDataSourcesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the projects data sources partial update params
func (o *ProjectsDataSourcesPartialUpdateParams) WithData(data ProjectsDataSourcesPartialUpdateBody) *ProjectsDataSourcesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the projects data sources partial update params
func (o *ProjectsDataSourcesPartialUpdateParams) SetData(data ProjectsDataSourcesPartialUpdateBody) {
	o.Data = data
}

// WithNamespace adds the namespace to the projects data sources partial update params
func (o *ProjectsDataSourcesPartialUpdateParams) WithNamespace(namespace string) *ProjectsDataSourcesPartialUpdateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects data sources partial update params
func (o *ProjectsDataSourcesPartialUpdateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectPk adds the projectPk to the projects data sources partial update params
func (o *ProjectsDataSourcesPartialUpdateParams) WithProjectPk(projectPk string) *ProjectsDataSourcesPartialUpdateParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects data sources partial update params
func (o *ProjectsDataSourcesPartialUpdateParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WithServer adds the server to the projects data sources partial update params
func (o *ProjectsDataSourcesPartialUpdateParams) WithServer(server string) *ProjectsDataSourcesPartialUpdateParams {
	o.SetServer(server)
	return o
}

// SetServer adds the server to the projects data sources partial update params
func (o *ProjectsDataSourcesPartialUpdateParams) SetServer(server string) {
	o.Server = server
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsDataSourcesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Data); err != nil {
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

	// path param server
	if err := r.SetPathParam("server", o.Server); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}