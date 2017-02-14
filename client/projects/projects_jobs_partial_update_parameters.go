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

// NewProjectsJobsPartialUpdateParams creates a new ProjectsJobsPartialUpdateParams object
// with the default values initialized.
func NewProjectsJobsPartialUpdateParams() *ProjectsJobsPartialUpdateParams {
	var ()
	return &ProjectsJobsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsJobsPartialUpdateParamsWithTimeout creates a new ProjectsJobsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsJobsPartialUpdateParamsWithTimeout(timeout time.Duration) *ProjectsJobsPartialUpdateParams {
	var ()
	return &ProjectsJobsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewProjectsJobsPartialUpdateParamsWithContext creates a new ProjectsJobsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsJobsPartialUpdateParamsWithContext(ctx context.Context) *ProjectsJobsPartialUpdateParams {
	var ()
	return &ProjectsJobsPartialUpdateParams{

		Context: ctx,
	}
}

// NewProjectsJobsPartialUpdateParamsWithHTTPClient creates a new ProjectsJobsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsJobsPartialUpdateParamsWithHTTPClient(client *http.Client) *ProjectsJobsPartialUpdateParams {
	var ()
	return &ProjectsJobsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*ProjectsJobsPartialUpdateParams contains all the parameters to send to the API endpoint
for the projects jobs partial update operation typically these are written to a http.Request
*/
type ProjectsJobsPartialUpdateParams struct {

	/*Data*/
	Data ProjectsJobsPartialUpdateBody
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

// WithTimeout adds the timeout to the projects jobs partial update params
func (o *ProjectsJobsPartialUpdateParams) WithTimeout(timeout time.Duration) *ProjectsJobsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects jobs partial update params
func (o *ProjectsJobsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects jobs partial update params
func (o *ProjectsJobsPartialUpdateParams) WithContext(ctx context.Context) *ProjectsJobsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects jobs partial update params
func (o *ProjectsJobsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects jobs partial update params
func (o *ProjectsJobsPartialUpdateParams) WithHTTPClient(client *http.Client) *ProjectsJobsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects jobs partial update params
func (o *ProjectsJobsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the projects jobs partial update params
func (o *ProjectsJobsPartialUpdateParams) WithData(data ProjectsJobsPartialUpdateBody) *ProjectsJobsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the projects jobs partial update params
func (o *ProjectsJobsPartialUpdateParams) SetData(data ProjectsJobsPartialUpdateBody) {
	o.Data = data
}

// WithNamespace adds the namespace to the projects jobs partial update params
func (o *ProjectsJobsPartialUpdateParams) WithNamespace(namespace string) *ProjectsJobsPartialUpdateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects jobs partial update params
func (o *ProjectsJobsPartialUpdateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectPk adds the projectPk to the projects jobs partial update params
func (o *ProjectsJobsPartialUpdateParams) WithProjectPk(projectPk string) *ProjectsJobsPartialUpdateParams {
	o.SetProjectPk(projectPk)
	return o
}

// SetProjectPk adds the projectPk to the projects jobs partial update params
func (o *ProjectsJobsPartialUpdateParams) SetProjectPk(projectPk string) {
	o.ProjectPk = projectPk
}

// WithServer adds the server to the projects jobs partial update params
func (o *ProjectsJobsPartialUpdateParams) WithServer(server string) *ProjectsJobsPartialUpdateParams {
	o.SetServer(server)
	return o
}

// SetServer adds the server to the projects jobs partial update params
func (o *ProjectsJobsPartialUpdateParams) SetServer(server string) {
	o.Server = server
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsJobsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
