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

// NewAuthConvertTokenCreateParams creates a new AuthConvertTokenCreateParams object
// with the default values initialized.
func NewAuthConvertTokenCreateParams() *AuthConvertTokenCreateParams {

	return &AuthConvertTokenCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAuthConvertTokenCreateParamsWithTimeout creates a new AuthConvertTokenCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAuthConvertTokenCreateParamsWithTimeout(timeout time.Duration) *AuthConvertTokenCreateParams {

	return &AuthConvertTokenCreateParams{

		timeout: timeout,
	}
}

// NewAuthConvertTokenCreateParamsWithContext creates a new AuthConvertTokenCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewAuthConvertTokenCreateParamsWithContext(ctx context.Context) *AuthConvertTokenCreateParams {

	return &AuthConvertTokenCreateParams{

		Context: ctx,
	}
}

// NewAuthConvertTokenCreateParamsWithHTTPClient creates a new AuthConvertTokenCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAuthConvertTokenCreateParamsWithHTTPClient(client *http.Client) *AuthConvertTokenCreateParams {

	return &AuthConvertTokenCreateParams{
		HTTPClient: client,
	}
}

/*AuthConvertTokenCreateParams contains all the parameters to send to the API endpoint
for the auth convert token create operation typically these are written to a http.Request
*/
type AuthConvertTokenCreateParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the auth convert token create params
func (o *AuthConvertTokenCreateParams) WithTimeout(timeout time.Duration) *AuthConvertTokenCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auth convert token create params
func (o *AuthConvertTokenCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auth convert token create params
func (o *AuthConvertTokenCreateParams) WithContext(ctx context.Context) *AuthConvertTokenCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auth convert token create params
func (o *AuthConvertTokenCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auth convert token create params
func (o *AuthConvertTokenCreateParams) WithHTTPClient(client *http.Client) *AuthConvertTokenCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auth convert token create params
func (o *AuthConvertTokenCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AuthConvertTokenCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}