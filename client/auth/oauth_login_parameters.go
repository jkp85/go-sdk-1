// Code generated by go-swagger; DO NOT EDIT.

package auth

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

// NewOauthLoginParams creates a new OauthLoginParams object
// with the default values initialized.
func NewOauthLoginParams() *OauthLoginParams {
	var ()
	return &OauthLoginParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOauthLoginParamsWithTimeout creates a new OauthLoginParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOauthLoginParamsWithTimeout(timeout time.Duration) *OauthLoginParams {
	var ()
	return &OauthLoginParams{

		timeout: timeout,
	}
}

// NewOauthLoginParamsWithContext creates a new OauthLoginParams object
// with the default values initialized, and the ability to set a context for a request
func NewOauthLoginParamsWithContext(ctx context.Context) *OauthLoginParams {
	var ()
	return &OauthLoginParams{

		Context: ctx,
	}
}

// NewOauthLoginParamsWithHTTPClient creates a new OauthLoginParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOauthLoginParamsWithHTTPClient(client *http.Client) *OauthLoginParams {
	var ()
	return &OauthLoginParams{
		HTTPClient: client,
	}
}

/*OauthLoginParams contains all the parameters to send to the API endpoint
for the oauth login operation typically these are written to a http.Request
*/
type OauthLoginParams struct {

	/*Provider
	  OAuth2 provider

	*/
	Provider string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the oauth login params
func (o *OauthLoginParams) WithTimeout(timeout time.Duration) *OauthLoginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the oauth login params
func (o *OauthLoginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the oauth login params
func (o *OauthLoginParams) WithContext(ctx context.Context) *OauthLoginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the oauth login params
func (o *OauthLoginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the oauth login params
func (o *OauthLoginParams) WithHTTPClient(client *http.Client) *OauthLoginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the oauth login params
func (o *OauthLoginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProvider adds the provider to the oauth login params
func (o *OauthLoginParams) WithProvider(provider string) *OauthLoginParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the oauth login params
func (o *OauthLoginParams) SetProvider(provider string) {
	o.Provider = provider
}

// WriteToRequest writes these params to a swagger request
func (o *OauthLoginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param provider
	if err := r.SetPathParam("provider", o.Provider); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
