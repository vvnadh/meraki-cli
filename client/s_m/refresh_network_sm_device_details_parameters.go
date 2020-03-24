// Code generated by go-swagger; DO NOT EDIT.

package s_m

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
)

// NewRefreshNetworkSmDeviceDetailsParams creates a new RefreshNetworkSmDeviceDetailsParams object
// with the default values initialized.
func NewRefreshNetworkSmDeviceDetailsParams() *RefreshNetworkSmDeviceDetailsParams {
	var ()
	return &RefreshNetworkSmDeviceDetailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRefreshNetworkSmDeviceDetailsParamsWithTimeout creates a new RefreshNetworkSmDeviceDetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRefreshNetworkSmDeviceDetailsParamsWithTimeout(timeout time.Duration) *RefreshNetworkSmDeviceDetailsParams {
	var ()
	return &RefreshNetworkSmDeviceDetailsParams{

		timeout: timeout,
	}
}

// NewRefreshNetworkSmDeviceDetailsParamsWithContext creates a new RefreshNetworkSmDeviceDetailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewRefreshNetworkSmDeviceDetailsParamsWithContext(ctx context.Context) *RefreshNetworkSmDeviceDetailsParams {
	var ()
	return &RefreshNetworkSmDeviceDetailsParams{

		Context: ctx,
	}
}

// NewRefreshNetworkSmDeviceDetailsParamsWithHTTPClient creates a new RefreshNetworkSmDeviceDetailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRefreshNetworkSmDeviceDetailsParamsWithHTTPClient(client *http.Client) *RefreshNetworkSmDeviceDetailsParams {
	var ()
	return &RefreshNetworkSmDeviceDetailsParams{
		HTTPClient: client,
	}
}

/*RefreshNetworkSmDeviceDetailsParams contains all the parameters to send to the API endpoint
for the refresh network sm device details operation typically these are written to a http.Request
*/
type RefreshNetworkSmDeviceDetailsParams struct {

	/*DeviceID*/
	DeviceID string
	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the refresh network sm device details params
func (o *RefreshNetworkSmDeviceDetailsParams) WithTimeout(timeout time.Duration) *RefreshNetworkSmDeviceDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the refresh network sm device details params
func (o *RefreshNetworkSmDeviceDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the refresh network sm device details params
func (o *RefreshNetworkSmDeviceDetailsParams) WithContext(ctx context.Context) *RefreshNetworkSmDeviceDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the refresh network sm device details params
func (o *RefreshNetworkSmDeviceDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the refresh network sm device details params
func (o *RefreshNetworkSmDeviceDetailsParams) WithHTTPClient(client *http.Client) *RefreshNetworkSmDeviceDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the refresh network sm device details params
func (o *RefreshNetworkSmDeviceDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the refresh network sm device details params
func (o *RefreshNetworkSmDeviceDetailsParams) WithDeviceID(deviceID string) *RefreshNetworkSmDeviceDetailsParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the refresh network sm device details params
func (o *RefreshNetworkSmDeviceDetailsParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithNetworkID adds the networkID to the refresh network sm device details params
func (o *RefreshNetworkSmDeviceDetailsParams) WithNetworkID(networkID string) *RefreshNetworkSmDeviceDetailsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the refresh network sm device details params
func (o *RefreshNetworkSmDeviceDetailsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *RefreshNetworkSmDeviceDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
