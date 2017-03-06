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

// NewAuthGetTokenCreateParams creates a new AuthGetTokenCreateParams object
// with the default values initialized.
func NewAuthGetTokenCreateParams() *AuthGetTokenCreateParams {
	var ()
	return &AuthGetTokenCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAuthGetTokenCreateParamsWithTimeout creates a new AuthGetTokenCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAuthGetTokenCreateParamsWithTimeout(timeout time.Duration) *AuthGetTokenCreateParams {
	var ()
	return &AuthGetTokenCreateParams{

		timeout: timeout,
	}
}

// NewAuthGetTokenCreateParamsWithContext creates a new AuthGetTokenCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewAuthGetTokenCreateParamsWithContext(ctx context.Context) *AuthGetTokenCreateParams {
	var ()
	return &AuthGetTokenCreateParams{

		Context: ctx,
	}
}

// NewAuthGetTokenCreateParamsWithHTTPClient creates a new AuthGetTokenCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAuthGetTokenCreateParamsWithHTTPClient(client *http.Client) *AuthGetTokenCreateParams {
	var ()
	return &AuthGetTokenCreateParams{
		HTTPClient: client,
	}
}

/*AuthGetTokenCreateParams contains all the parameters to send to the API endpoint
for the auth get token create operation typically these are written to a http.Request
*/
type AuthGetTokenCreateParams struct {

	/*Data*/
	Data AuthGetTokenCreateBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the auth get token create params
func (o *AuthGetTokenCreateParams) WithTimeout(timeout time.Duration) *AuthGetTokenCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auth get token create params
func (o *AuthGetTokenCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auth get token create params
func (o *AuthGetTokenCreateParams) WithContext(ctx context.Context) *AuthGetTokenCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auth get token create params
func (o *AuthGetTokenCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auth get token create params
func (o *AuthGetTokenCreateParams) WithHTTPClient(client *http.Client) *AuthGetTokenCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auth get token create params
func (o *AuthGetTokenCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the auth get token create params
func (o *AuthGetTokenCreateParams) WithData(data AuthGetTokenCreateBody) *AuthGetTokenCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the auth get token create params
func (o *AuthGetTokenCreateParams) SetData(data AuthGetTokenCreateBody) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *AuthGetTokenCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Data); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
