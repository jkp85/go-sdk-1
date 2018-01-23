// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// NewTeamsGroupsRemoveFromGroupParams creates a new TeamsGroupsRemoveFromGroupParams object
// with the default values initialized.
func NewTeamsGroupsRemoveFromGroupParams() *TeamsGroupsRemoveFromGroupParams {
	var ()
	return &TeamsGroupsRemoveFromGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTeamsGroupsRemoveFromGroupParamsWithTimeout creates a new TeamsGroupsRemoveFromGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTeamsGroupsRemoveFromGroupParamsWithTimeout(timeout time.Duration) *TeamsGroupsRemoveFromGroupParams {
	var ()
	return &TeamsGroupsRemoveFromGroupParams{

		timeout: timeout,
	}
}

// NewTeamsGroupsRemoveFromGroupParamsWithContext creates a new TeamsGroupsRemoveFromGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewTeamsGroupsRemoveFromGroupParamsWithContext(ctx context.Context) *TeamsGroupsRemoveFromGroupParams {
	var ()
	return &TeamsGroupsRemoveFromGroupParams{

		Context: ctx,
	}
}

// NewTeamsGroupsRemoveFromGroupParamsWithHTTPClient creates a new TeamsGroupsRemoveFromGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTeamsGroupsRemoveFromGroupParamsWithHTTPClient(client *http.Client) *TeamsGroupsRemoveFromGroupParams {
	var ()
	return &TeamsGroupsRemoveFromGroupParams{
		HTTPClient: client,
	}
}

/*TeamsGroupsRemoveFromGroupParams contains all the parameters to send to the API endpoint
for the teams groups remove from group operation typically these are written to a http.Request
*/
type TeamsGroupsRemoveFromGroupParams struct {

	/*Group
	  Group unique identifier expressed as UUID or name.

	*/
	Group string
	/*Team
	  Team unique identifier expressed as UUID or name.

	*/
	Team string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the teams groups remove from group params
func (o *TeamsGroupsRemoveFromGroupParams) WithTimeout(timeout time.Duration) *TeamsGroupsRemoveFromGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the teams groups remove from group params
func (o *TeamsGroupsRemoveFromGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the teams groups remove from group params
func (o *TeamsGroupsRemoveFromGroupParams) WithContext(ctx context.Context) *TeamsGroupsRemoveFromGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the teams groups remove from group params
func (o *TeamsGroupsRemoveFromGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the teams groups remove from group params
func (o *TeamsGroupsRemoveFromGroupParams) WithHTTPClient(client *http.Client) *TeamsGroupsRemoveFromGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the teams groups remove from group params
func (o *TeamsGroupsRemoveFromGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroup adds the group to the teams groups remove from group params
func (o *TeamsGroupsRemoveFromGroupParams) WithGroup(group string) *TeamsGroupsRemoveFromGroupParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the teams groups remove from group params
func (o *TeamsGroupsRemoveFromGroupParams) SetGroup(group string) {
	o.Group = group
}

// WithTeam adds the team to the teams groups remove from group params
func (o *TeamsGroupsRemoveFromGroupParams) WithTeam(team string) *TeamsGroupsRemoveFromGroupParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the teams groups remove from group params
func (o *TeamsGroupsRemoveFromGroupParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *TeamsGroupsRemoveFromGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param group
	if err := r.SetPathParam("group", o.Group); err != nil {
		return err
	}

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}