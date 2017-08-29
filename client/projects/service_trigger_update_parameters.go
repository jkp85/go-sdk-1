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

// NewServiceTriggerUpdateParams creates a new ServiceTriggerUpdateParams object
// with the default values initialized.
func NewServiceTriggerUpdateParams() *ServiceTriggerUpdateParams {
	var ()
	return &ServiceTriggerUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServiceTriggerUpdateParamsWithTimeout creates a new ServiceTriggerUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServiceTriggerUpdateParamsWithTimeout(timeout time.Duration) *ServiceTriggerUpdateParams {
	var ()
	return &ServiceTriggerUpdateParams{

		timeout: timeout,
	}
}

// NewServiceTriggerUpdateParamsWithContext creates a new ServiceTriggerUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewServiceTriggerUpdateParamsWithContext(ctx context.Context) *ServiceTriggerUpdateParams {
	var ()
	return &ServiceTriggerUpdateParams{

		Context: ctx,
	}
}

// NewServiceTriggerUpdateParamsWithHTTPClient creates a new ServiceTriggerUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServiceTriggerUpdateParamsWithHTTPClient(client *http.Client) *ServiceTriggerUpdateParams {
	var ()
	return &ServiceTriggerUpdateParams{
		HTTPClient: client,
	}
}

/*ServiceTriggerUpdateParams contains all the parameters to send to the API endpoint
for the service trigger update operation typically these are written to a http.Request
*/
type ServiceTriggerUpdateParams struct {

	/*ID
	  Trigger identifier expressed as UUID.

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
	/*ServerAction*/
	ServerAction *models.ServerActionData
	/*ServerID
	  Server unique identifier expressed as UUID.

	*/
	ServerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the service trigger update params
func (o *ServiceTriggerUpdateParams) WithTimeout(timeout time.Duration) *ServiceTriggerUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service trigger update params
func (o *ServiceTriggerUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service trigger update params
func (o *ServiceTriggerUpdateParams) WithContext(ctx context.Context) *ServiceTriggerUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service trigger update params
func (o *ServiceTriggerUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service trigger update params
func (o *ServiceTriggerUpdateParams) WithHTTPClient(client *http.Client) *ServiceTriggerUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service trigger update params
func (o *ServiceTriggerUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the service trigger update params
func (o *ServiceTriggerUpdateParams) WithID(id string) *ServiceTriggerUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the service trigger update params
func (o *ServiceTriggerUpdateParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the service trigger update params
func (o *ServiceTriggerUpdateParams) WithNamespace(namespace string) *ServiceTriggerUpdateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the service trigger update params
func (o *ServiceTriggerUpdateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProjectID adds the projectID to the service trigger update params
func (o *ServiceTriggerUpdateParams) WithProjectID(projectID string) *ServiceTriggerUpdateParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the service trigger update params
func (o *ServiceTriggerUpdateParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithServerAction adds the serverAction to the service trigger update params
func (o *ServiceTriggerUpdateParams) WithServerAction(serverAction *models.ServerActionData) *ServiceTriggerUpdateParams {
	o.SetServerAction(serverAction)
	return o
}

// SetServerAction adds the serverAction to the service trigger update params
func (o *ServiceTriggerUpdateParams) SetServerAction(serverAction *models.ServerActionData) {
	o.ServerAction = serverAction
}

// WithServerID adds the serverID to the service trigger update params
func (o *ServiceTriggerUpdateParams) WithServerID(serverID string) *ServiceTriggerUpdateParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the service trigger update params
func (o *ServiceTriggerUpdateParams) SetServerID(serverID string) {
	o.ServerID = serverID
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceTriggerUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ServerAction == nil {
		o.ServerAction = new(models.ServerActionData)
	}

	if err := r.SetBodyParam(o.ServerAction); err != nil {
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
