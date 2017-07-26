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

// NewProjectsSyncedResourcesListParams creates a new ProjectsSyncedResourcesListParams object
// with the default values initialized.
func NewProjectsSyncedResourcesListParams() *ProjectsSyncedResourcesListParams {
	var ()
	return &ProjectsSyncedResourcesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsSyncedResourcesListParamsWithTimeout creates a new ProjectsSyncedResourcesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectsSyncedResourcesListParamsWithTimeout(timeout time.Duration) *ProjectsSyncedResourcesListParams {
	var ()
	return &ProjectsSyncedResourcesListParams{

		timeout: timeout,
	}
}

// NewProjectsSyncedResourcesListParamsWithContext creates a new ProjectsSyncedResourcesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectsSyncedResourcesListParamsWithContext(ctx context.Context) *ProjectsSyncedResourcesListParams {
	var ()
	return &ProjectsSyncedResourcesListParams{

		Context: ctx,
	}
}

// NewProjectsSyncedResourcesListParamsWithHTTPClient creates a new ProjectsSyncedResourcesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectsSyncedResourcesListParamsWithHTTPClient(client *http.Client) *ProjectsSyncedResourcesListParams {
	var ()
	return &ProjectsSyncedResourcesListParams{
		HTTPClient: client,
	}
}

/*ProjectsSyncedResourcesListParams contains all the parameters to send to the API endpoint
for the projects synced resources list operation typically these are written to a http.Request
*/
type ProjectsSyncedResourcesListParams struct {

	/*Limit
	  Limit when getting items.

	*/
	Limit *string
	/*Namespace
	  User or team name.

	*/
	Namespace string
	/*Offset
	  Offset when getting items.

	*/
	Offset *string
	/*Ordering
	  Ordering when getting items.

	*/
	Ordering *string
	/*ProjectID
	  Project unique identifier expressed as UUID.

	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the projects synced resources list params
func (o *ProjectsSyncedResourcesListParams) WithTimeout(timeout time.Duration) *ProjectsSyncedResourcesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects synced resources list params
func (o *ProjectsSyncedResourcesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects synced resources list params
func (o *ProjectsSyncedResourcesListParams) WithContext(ctx context.Context) *ProjectsSyncedResourcesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects synced resources list params
func (o *ProjectsSyncedResourcesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects synced resources list params
func (o *ProjectsSyncedResourcesListParams) WithHTTPClient(client *http.Client) *ProjectsSyncedResourcesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects synced resources list params
func (o *ProjectsSyncedResourcesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the projects synced resources list params
func (o *ProjectsSyncedResourcesListParams) WithLimit(limit *string) *ProjectsSyncedResourcesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the projects synced resources list params
func (o *ProjectsSyncedResourcesListParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the projects synced resources list params
func (o *ProjectsSyncedResourcesListParams) WithNamespace(namespace string) *ProjectsSyncedResourcesListParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the projects synced resources list params
func (o *ProjectsSyncedResourcesListParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the projects synced resources list params
func (o *ProjectsSyncedResourcesListParams) WithOffset(offset *string) *ProjectsSyncedResourcesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the projects synced resources list params
func (o *ProjectsSyncedResourcesListParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithOrdering adds the ordering to the projects synced resources list params
func (o *ProjectsSyncedResourcesListParams) WithOrdering(ordering *string) *ProjectsSyncedResourcesListParams {
	o.SetOrdering(ordering)
	return o
}

// SetOrdering adds the ordering to the projects synced resources list params
func (o *ProjectsSyncedResourcesListParams) SetOrdering(ordering *string) {
	o.Ordering = ordering
}

// WithProjectID adds the projectID to the projects synced resources list params
func (o *ProjectsSyncedResourcesListParams) WithProjectID(projectID string) *ProjectsSyncedResourcesListParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the projects synced resources list params
func (o *ProjectsSyncedResourcesListParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsSyncedResourcesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit string
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Ordering != nil {

		// query param ordering
		var qrOrdering string
		if o.Ordering != nil {
			qrOrdering = *o.Ordering
		}
		qOrdering := qrOrdering
		if qOrdering != "" {
			if err := r.SetQueryParam("ordering", qOrdering); err != nil {
				return err
			}
		}

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
