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

// NewServiceTriggerReadParams creates a new ServiceTriggerReadParams object
// with the default values initialized.
func NewServiceTriggerReadParams() *ServiceTriggerReadParams {
	var ()
	return &ServiceTriggerReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServiceTriggerReadParamsWithTimeout creates a new ServiceTriggerReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServiceTriggerReadParamsWithTimeout(timeout time.Duration) *ServiceTriggerReadParams {
	var ()
	return &ServiceTriggerReadParams{

		timeout: timeout,
	}
}

// NewServiceTriggerReadParamsWithContext creates a new ServiceTriggerReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewServiceTriggerReadParamsWithContext(ctx context.Context) *ServiceTriggerReadParams {
	var ()
	return &ServiceTriggerReadParams{

		Context: ctx,
	}
}

// NewServiceTriggerReadParamsWithHTTPClient creates a new ServiceTriggerReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServiceTriggerReadParamsWithHTTPClient(client *http.Client) *ServiceTriggerReadParams {
	var ()
	return &ServiceTriggerReadParams{
		HTTPClient: client,
	}
}

/*ServiceTriggerReadParams contains all the parameters to send to the API endpoint
for the service trigger read operation typically these are written to a http.Request
*/
type ServiceTriggerReadParams struct {

	/*ID
	  Trigger unique identifier.

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the service trigger read params
func (o *ServiceTriggerReadParams) WithTimeout(timeout time.Duration) *ServiceTriggerReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service trigger read params
func (o *ServiceTriggerReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service trigger read params
func (o *ServiceTriggerReadParams) WithContext(ctx context.Context) *ServiceTriggerReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service trigger read params
func (o *ServiceTriggerReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service trigger read params
func (o *ServiceTriggerReadParams) WithHTTPClient(client *http.Client) *ServiceTriggerReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service trigger read params
func (o *ServiceTriggerReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the service trigger read params
func (o *ServiceTriggerReadParams) WithID(id string) *ServiceTriggerReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the service trigger read params
func (o *ServiceTriggerReadParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the service trigger read params
func (o *ServiceTriggerReadParams) WithNamespace(namespace string) *ServiceTriggerReadParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the service trigger read params
func (o *ServiceTriggerReadParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectID adds the projectID to the service trigger read params
func (o *ServiceTriggerReadParams) WithProjectID(projectID string) *ServiceTriggerReadParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the service trigger read params
func (o *ServiceTriggerReadParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithServerID adds the serverID to the service trigger read params
func (o *ServiceTriggerReadParams) WithServerID(serverID string) *ServiceTriggerReadParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the service trigger read params
func (o *ServiceTriggerReadParams) SetServerID(serverID string) {
	o.ServerID = serverID
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceTriggerReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
