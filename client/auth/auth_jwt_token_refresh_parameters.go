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

	models "github.com/IllumiDesk/go-sdk/models"
)

// NewAuthJwtTokenRefreshParams creates a new AuthJwtTokenRefreshParams object
// with the default values initialized.
func NewAuthJwtTokenRefreshParams() *AuthJwtTokenRefreshParams {
	var ()
	return &AuthJwtTokenRefreshParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAuthJwtTokenRefreshParamsWithTimeout creates a new AuthJwtTokenRefreshParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAuthJwtTokenRefreshParamsWithTimeout(timeout time.Duration) *AuthJwtTokenRefreshParams {
	var ()
	return &AuthJwtTokenRefreshParams{

		timeout: timeout,
	}
}

// NewAuthJwtTokenRefreshParamsWithContext creates a new AuthJwtTokenRefreshParams object
// with the default values initialized, and the ability to set a context for a request
func NewAuthJwtTokenRefreshParamsWithContext(ctx context.Context) *AuthJwtTokenRefreshParams {
	var ()
	return &AuthJwtTokenRefreshParams{

		Context: ctx,
	}
}

// NewAuthJwtTokenRefreshParamsWithHTTPClient creates a new AuthJwtTokenRefreshParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAuthJwtTokenRefreshParamsWithHTTPClient(client *http.Client) *AuthJwtTokenRefreshParams {
	var ()
	return &AuthJwtTokenRefreshParams{
		HTTPClient: client,
	}
}

/*AuthJwtTokenRefreshParams contains all the parameters to send to the API endpoint
for the auth jwt token refresh operation typically these are written to a http.Request
*/
type AuthJwtTokenRefreshParams struct {

	/*RefreshjwtData*/
	RefreshjwtData *models.RefreshJSONWebTokenData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the auth jwt token refresh params
func (o *AuthJwtTokenRefreshParams) WithTimeout(timeout time.Duration) *AuthJwtTokenRefreshParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auth jwt token refresh params
func (o *AuthJwtTokenRefreshParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auth jwt token refresh params
func (o *AuthJwtTokenRefreshParams) WithContext(ctx context.Context) *AuthJwtTokenRefreshParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auth jwt token refresh params
func (o *AuthJwtTokenRefreshParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auth jwt token refresh params
func (o *AuthJwtTokenRefreshParams) WithHTTPClient(client *http.Client) *AuthJwtTokenRefreshParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auth jwt token refresh params
func (o *AuthJwtTokenRefreshParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRefreshjwtData adds the refreshjwtData to the auth jwt token refresh params
func (o *AuthJwtTokenRefreshParams) WithRefreshjwtData(refreshjwtData *models.RefreshJSONWebTokenData) *AuthJwtTokenRefreshParams {
	o.SetRefreshjwtData(refreshjwtData)
	return o
}

// SetRefreshjwtData adds the refreshjwtData to the auth jwt token refresh params
func (o *AuthJwtTokenRefreshParams) SetRefreshjwtData(refreshjwtData *models.RefreshJSONWebTokenData) {
	o.RefreshjwtData = refreshjwtData
}

// WriteToRequest writes these params to a swagger request
func (o *AuthJwtTokenRefreshParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RefreshjwtData != nil {
		if err := r.SetBodyParam(o.RefreshjwtData); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
