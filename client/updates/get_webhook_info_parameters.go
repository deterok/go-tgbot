package updates

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

// NewGetWebhookInfoParams creates a new GetWebhookInfoParams object
// with the default values initialized.
func NewGetWebhookInfoParams() *GetWebhookInfoParams {

	return &GetWebhookInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWebhookInfoParamsWithTimeout creates a new GetWebhookInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWebhookInfoParamsWithTimeout(timeout time.Duration) *GetWebhookInfoParams {

	return &GetWebhookInfoParams{

		timeout: timeout,
	}
}

// NewGetWebhookInfoParamsWithContext creates a new GetWebhookInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWebhookInfoParamsWithContext(ctx context.Context) *GetWebhookInfoParams {

	return &GetWebhookInfoParams{

		Context: ctx,
	}
}

// NewGetWebhookInfoParamsWithHTTPClient creates a new GetWebhookInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWebhookInfoParamsWithHTTPClient(client *http.Client) *GetWebhookInfoParams {

	return &GetWebhookInfoParams{
		HTTPClient: client,
	}
}

/*GetWebhookInfoParams contains all the parameters to send to the API endpoint
for the get webhook info operation typically these are written to a http.Request
*/
type GetWebhookInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get webhook info params
func (o *GetWebhookInfoParams) WithTimeout(timeout time.Duration) *GetWebhookInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get webhook info params
func (o *GetWebhookInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get webhook info params
func (o *GetWebhookInfoParams) WithContext(ctx context.Context) *GetWebhookInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get webhook info params
func (o *GetWebhookInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get webhook info params
func (o *GetWebhookInfoParams) WithHTTPClient(client *http.Client) *GetWebhookInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get webhook info params
func (o *GetWebhookInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetWebhookInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}