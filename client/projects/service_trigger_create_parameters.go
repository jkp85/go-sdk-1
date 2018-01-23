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

	models "github.com/IllumiDesk/go-sdk/models"
)

// NewServiceTriggerCreateParams creates a new ServiceTriggerCreateParams object
// with the default values initialized.
func NewServiceTriggerCreateParams() *ServiceTriggerCreateParams {
	var ()
	return &ServiceTriggerCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServiceTriggerCreateParamsWithTimeout creates a new ServiceTriggerCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServiceTriggerCreateParamsWithTimeout(timeout time.Duration) *ServiceTriggerCreateParams {
	var ()
	return &ServiceTriggerCreateParams{

		timeout: timeout,
	}
}

// NewServiceTriggerCreateParamsWithContext creates a new ServiceTriggerCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewServiceTriggerCreateParamsWithContext(ctx context.Context) *ServiceTriggerCreateParams {
	var ()
	return &ServiceTriggerCreateParams{

		Context: ctx,
	}
}

// NewServiceTriggerCreateParamsWithHTTPClient creates a new ServiceTriggerCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServiceTriggerCreateParamsWithHTTPClient(client *http.Client) *ServiceTriggerCreateParams {
	var ()
	return &ServiceTriggerCreateParams{
		HTTPClient: client,
	}
}

/*ServiceTriggerCreateParams contains all the parameters to send to the API endpoint
for the service trigger create operation typically these are written to a http.Request
*/
type ServiceTriggerCreateParams struct {

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
	/*ServerAction
	  Server action.

	*/
	ServerAction *models.ServerActionData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the service trigger create params
func (o *ServiceTriggerCreateParams) WithTimeout(timeout time.Duration) *ServiceTriggerCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service trigger create params
func (o *ServiceTriggerCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service trigger create params
func (o *ServiceTriggerCreateParams) WithContext(ctx context.Context) *ServiceTriggerCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service trigger create params
func (o *ServiceTriggerCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service trigger create params
func (o *ServiceTriggerCreateParams) WithHTTPClient(client *http.Client) *ServiceTriggerCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service trigger create params
func (o *ServiceTriggerCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the service trigger create params
func (o *ServiceTriggerCreateParams) WithNamespace(namespace string) *ServiceTriggerCreateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the service trigger create params
func (o *ServiceTriggerCreateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProject adds the project to the service trigger create params
func (o *ServiceTriggerCreateParams) WithProject(project string) *ServiceTriggerCreateParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the service trigger create params
func (o *ServiceTriggerCreateParams) SetProject(project string) {
	o.Project = project
}

// WithServer adds the server to the service trigger create params
func (o *ServiceTriggerCreateParams) WithServer(server string) *ServiceTriggerCreateParams {
	o.SetServer(server)
	return o
}

// SetServer adds the server to the service trigger create params
func (o *ServiceTriggerCreateParams) SetServer(server string) {
	o.Server = server
}

// WithServerAction adds the serverAction to the service trigger create params
func (o *ServiceTriggerCreateParams) WithServerAction(serverAction *models.ServerActionData) *ServiceTriggerCreateParams {
	o.SetServerAction(serverAction)
	return o
}

// SetServerAction adds the serverAction to the service trigger create params
func (o *ServiceTriggerCreateParams) SetServerAction(serverAction *models.ServerActionData) {
	o.ServerAction = serverAction
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceTriggerCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ServerAction != nil {
		if err := r.SetBodyParam(o.ServerAction); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
