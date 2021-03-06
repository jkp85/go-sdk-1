// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/IllumiDesk/go-sdk/client/auth"
	"github.com/IllumiDesk/go-sdk/client/billing"
	"github.com/IllumiDesk/go-sdk/client/notifications"
	"github.com/IllumiDesk/go-sdk/client/operations"
	"github.com/IllumiDesk/go-sdk/client/projects"
	"github.com/IllumiDesk/go-sdk/client/search"
	"github.com/IllumiDesk/go-sdk/client/servers"
	"github.com/IllumiDesk/go-sdk/client/teams"
	"github.com/IllumiDesk/go-sdk/client/users"
)

// Default illumidesk HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.3blades.ai"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new illumidesk HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Illumidesk {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new illumidesk HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Illumidesk {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new illumidesk client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Illumidesk {
	cli := new(Illumidesk)
	cli.Transport = transport

	cli.Auth = auth.New(transport, formats)

	cli.Billing = billing.New(transport, formats)

	cli.Notifications = notifications.New(transport, formats)

	cli.Operations = operations.New(transport, formats)

	cli.Projects = projects.New(transport, formats)

	cli.Search = search.New(transport, formats)

	cli.Servers = servers.New(transport, formats)

	cli.Teams = teams.New(transport, formats)

	cli.Users = users.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Illumidesk is a client for illumidesk
type Illumidesk struct {
	Auth *auth.Client

	Billing *billing.Client

	Notifications *notifications.Client

	Operations *operations.Client

	Projects *projects.Client

	Search *search.Client

	Servers *servers.Client

	Teams *teams.Client

	Users *users.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Illumidesk) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Auth.SetTransport(transport)

	c.Billing.SetTransport(transport)

	c.Notifications.SetTransport(transport)

	c.Operations.SetTransport(transport)

	c.Projects.SetTransport(transport)

	c.Search.SetTransport(transport)

	c.Servers.SetTransport(transport)

	c.Teams.SetTransport(transport)

	c.Users.SetTransport(transport)

}
