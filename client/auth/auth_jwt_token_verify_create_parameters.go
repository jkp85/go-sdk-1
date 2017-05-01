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

// NewAuthJwtTokenVerifyCreateParams creates a new AuthJwtTokenVerifyCreateParams object
// with the default values initialized.
func NewAuthJwtTokenVerifyCreateParams() *AuthJwtTokenVerifyCreateParams {
	var ()
	return &AuthJwtTokenVerifyCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAuthJwtTokenVerifyCreateParamsWithTimeout creates a new AuthJwtTokenVerifyCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAuthJwtTokenVerifyCreateParamsWithTimeout(timeout time.Duration) *AuthJwtTokenVerifyCreateParams {
	var ()
	return &AuthJwtTokenVerifyCreateParams{

		timeout: timeout,
	}
}

// NewAuthJwtTokenVerifyCreateParamsWithContext creates a new AuthJwtTokenVerifyCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewAuthJwtTokenVerifyCreateParamsWithContext(ctx context.Context) *AuthJwtTokenVerifyCreateParams {
	var ()
	return &AuthJwtTokenVerifyCreateParams{

		Context: ctx,
	}
}

// NewAuthJwtTokenVerifyCreateParamsWithHTTPClient creates a new AuthJwtTokenVerifyCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAuthJwtTokenVerifyCreateParamsWithHTTPClient(client *http.Client) *AuthJwtTokenVerifyCreateParams {
	var ()
	return &AuthJwtTokenVerifyCreateParams{
		HTTPClient: client,
	}
}

/*AuthJwtTokenVerifyCreateParams contains all the parameters to send to the API endpoint
for the auth jwt token verify create operation typically these are written to a http.Request
*/
type AuthJwtTokenVerifyCreateParams struct {

	/*Data*/
	Data AuthJwtTokenVerifyCreateBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the auth jwt token verify create params
func (o *AuthJwtTokenVerifyCreateParams) WithTimeout(timeout time.Duration) *AuthJwtTokenVerifyCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auth jwt token verify create params
func (o *AuthJwtTokenVerifyCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auth jwt token verify create params
func (o *AuthJwtTokenVerifyCreateParams) WithContext(ctx context.Context) *AuthJwtTokenVerifyCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auth jwt token verify create params
func (o *AuthJwtTokenVerifyCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auth jwt token verify create params
func (o *AuthJwtTokenVerifyCreateParams) WithHTTPClient(client *http.Client) *AuthJwtTokenVerifyCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auth jwt token verify create params
func (o *AuthJwtTokenVerifyCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the auth jwt token verify create params
func (o *AuthJwtTokenVerifyCreateParams) WithData(data AuthJwtTokenVerifyCreateBody) *AuthJwtTokenVerifyCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the auth jwt token verify create params
func (o *AuthJwtTokenVerifyCreateParams) SetData(data AuthJwtTokenVerifyCreateBody) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *AuthJwtTokenVerifyCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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