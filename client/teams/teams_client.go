// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new teams API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for teams API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
TeamsCreate creates a new team
*/
func (a *Client) TeamsCreate(params *TeamsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_create",
		Method:             "POST",
		PathPattern:        "/v1/teams/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamsCreateCreated), nil

}

/*
TeamsDelete deletes a team
*/
func (a *Client) TeamsDelete(params *TeamsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_delete",
		Method:             "DELETE",
		PathPattern:        "/v1/teams/{team}/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamsDeleteNoContent), nil

}

/*
TeamsGroupsAddToGroup adds user to group
*/
func (a *Client) TeamsGroupsAddToGroup(params *TeamsGroupsAddToGroupParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsGroupsAddToGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsGroupsAddToGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_groups_add_to_group",
		Method:             "POST",
		PathPattern:        "/v1/teams/{team}/groups/{group}/add/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsGroupsAddToGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamsGroupsAddToGroupOK), nil

}

/*
TeamsGroupsDelete deletes team group
*/
func (a *Client) TeamsGroupsDelete(params *TeamsGroupsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsGroupsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsGroupsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_groups_delete",
		Method:             "DELETE",
		PathPattern:        "/v1/teams/{team}/groups/{group}/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsGroupsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamsGroupsDeleteNoContent), nil

}

/*
TeamsGroupsList gets team groups
*/
func (a *Client) TeamsGroupsList(params *TeamsGroupsListParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsGroupsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsGroupsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_groups_list",
		Method:             "GET",
		PathPattern:        "/v1/teams/{team}/groups/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsGroupsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamsGroupsListOK), nil

}

/*
TeamsGroupsRead gets team group
*/
func (a *Client) TeamsGroupsRead(params *TeamsGroupsReadParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsGroupsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsGroupsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_groups_read",
		Method:             "GET",
		PathPattern:        "/v1/teams/{team}/groups/{group}/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsGroupsReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamsGroupsReadOK), nil

}

/*
TeamsGroupsRemoveFromGroup users removed from group
*/
func (a *Client) TeamsGroupsRemoveFromGroup(params *TeamsGroupsRemoveFromGroupParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsGroupsRemoveFromGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsGroupsRemoveFromGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_groups_remove_from_group",
		Method:             "POST",
		PathPattern:        "/v1/teams/{team}/groups/{group}/remove/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsGroupsRemoveFromGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamsGroupsRemoveFromGroupOK), nil

}

/*
TeamsGroupsReplace patches team group
*/
func (a *Client) TeamsGroupsReplace(params *TeamsGroupsReplaceParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsGroupsReplaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsGroupsReplaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_groups_replace",
		Method:             "PUT",
		PathPattern:        "/v1/teams/{team}/groups/{group}/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsGroupsReplaceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamsGroupsReplaceOK), nil

}

/*
TeamsGroupsUpdate patches team group
*/
func (a *Client) TeamsGroupsUpdate(params *TeamsGroupsUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsGroupsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsGroupsUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_groups_update",
		Method:             "PATCH",
		PathPattern:        "/v1/teams/{team}/groups/{group}/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsGroupsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamsGroupsUpdateOK), nil

}

/*
TeamsList gets teams
*/
func (a *Client) TeamsList(params *TeamsListParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_list",
		Method:             "GET",
		PathPattern:        "/v1/teams/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamsListOK), nil

}

/*
TeamsRead gets a team
*/
func (a *Client) TeamsRead(params *TeamsReadParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_read",
		Method:             "GET",
		PathPattern:        "/v1/teams/{team}/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamsReadOK), nil

}

/*
TeamsReplace replaces a team
*/
func (a *Client) TeamsReplace(params *TeamsReplaceParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsReplaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsReplaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_replace",
		Method:             "PUT",
		PathPattern:        "/v1/teams/{team}/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsReplaceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamsReplaceOK), nil

}

/*
TeamsUpdate updates a team
*/
func (a *Client) TeamsUpdate(params *TeamsUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_update",
		Method:             "PATCH",
		PathPattern:        "/v1/teams/{team}/",
		ProducesMediaTypes: []string{"application/json", "text/html"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TeamsUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
