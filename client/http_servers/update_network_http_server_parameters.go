// Code generated by go-swagger; DO NOT EDIT.

package http_servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/cisco-sso/meraki-cli/models"
)

// NewUpdateNetworkHTTPServerParams creates a new UpdateNetworkHTTPServerParams object
// with the default values initialized.
func NewUpdateNetworkHTTPServerParams() *UpdateNetworkHTTPServerParams {
	var ()
	return &UpdateNetworkHTTPServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkHTTPServerParamsWithTimeout creates a new UpdateNetworkHTTPServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNetworkHTTPServerParamsWithTimeout(timeout time.Duration) *UpdateNetworkHTTPServerParams {
	var ()
	return &UpdateNetworkHTTPServerParams{

		timeout: timeout,
	}
}

// NewUpdateNetworkHTTPServerParamsWithContext creates a new UpdateNetworkHTTPServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNetworkHTTPServerParamsWithContext(ctx context.Context) *UpdateNetworkHTTPServerParams {
	var ()
	return &UpdateNetworkHTTPServerParams{

		Context: ctx,
	}
}

// NewUpdateNetworkHTTPServerParamsWithHTTPClient creates a new UpdateNetworkHTTPServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNetworkHTTPServerParamsWithHTTPClient(client *http.Client) *UpdateNetworkHTTPServerParams {
	var ()
	return &UpdateNetworkHTTPServerParams{
		HTTPClient: client,
	}
}

/*UpdateNetworkHTTPServerParams contains all the parameters to send to the API endpoint
for the update network Http server operation typically these are written to a http.Request
*/
type UpdateNetworkHTTPServerParams struct {

	/*ID*/
	ID string
	/*NetworkID*/
	NetworkID string
	/*UpdateNetworkHTTPServer*/
	UpdateNetworkHTTPServer *models.UpdateNetworkHTTPServer

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update network Http server params
func (o *UpdateNetworkHTTPServerParams) WithTimeout(timeout time.Duration) *UpdateNetworkHTTPServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network Http server params
func (o *UpdateNetworkHTTPServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network Http server params
func (o *UpdateNetworkHTTPServerParams) WithContext(ctx context.Context) *UpdateNetworkHTTPServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network Http server params
func (o *UpdateNetworkHTTPServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network Http server params
func (o *UpdateNetworkHTTPServerParams) WithHTTPClient(client *http.Client) *UpdateNetworkHTTPServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network Http server params
func (o *UpdateNetworkHTTPServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update network Http server params
func (o *UpdateNetworkHTTPServerParams) WithID(id string) *UpdateNetworkHTTPServerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update network Http server params
func (o *UpdateNetworkHTTPServerParams) SetID(id string) {
	o.ID = id
}

// WithNetworkID adds the networkID to the update network Http server params
func (o *UpdateNetworkHTTPServerParams) WithNetworkID(networkID string) *UpdateNetworkHTTPServerParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network Http server params
func (o *UpdateNetworkHTTPServerParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkHTTPServer adds the updateNetworkHTTPServer to the update network Http server params
func (o *UpdateNetworkHTTPServerParams) WithUpdateNetworkHTTPServer(updateNetworkHTTPServer *models.UpdateNetworkHTTPServer) *UpdateNetworkHTTPServerParams {
	o.SetUpdateNetworkHTTPServer(updateNetworkHTTPServer)
	return o
}

// SetUpdateNetworkHTTPServer adds the updateNetworkHttpServer to the update network Http server params
func (o *UpdateNetworkHTTPServerParams) SetUpdateNetworkHTTPServer(updateNetworkHTTPServer *models.UpdateNetworkHTTPServer) {
	o.UpdateNetworkHTTPServer = updateNetworkHTTPServer
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkHTTPServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if o.UpdateNetworkHTTPServer != nil {
		if err := r.SetBodyParam(o.UpdateNetworkHTTPServer); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}