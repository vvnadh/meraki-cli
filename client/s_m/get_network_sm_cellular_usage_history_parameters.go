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

// NewGetNetworkSmCellularUsageHistoryParams creates a new GetNetworkSmCellularUsageHistoryParams object
// with the default values initialized.
func NewGetNetworkSmCellularUsageHistoryParams() *GetNetworkSmCellularUsageHistoryParams {
	var ()
	return &GetNetworkSmCellularUsageHistoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSmCellularUsageHistoryParamsWithTimeout creates a new GetNetworkSmCellularUsageHistoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkSmCellularUsageHistoryParamsWithTimeout(timeout time.Duration) *GetNetworkSmCellularUsageHistoryParams {
	var ()
	return &GetNetworkSmCellularUsageHistoryParams{

		timeout: timeout,
	}
}

// NewGetNetworkSmCellularUsageHistoryParamsWithContext creates a new GetNetworkSmCellularUsageHistoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkSmCellularUsageHistoryParamsWithContext(ctx context.Context) *GetNetworkSmCellularUsageHistoryParams {
	var ()
	return &GetNetworkSmCellularUsageHistoryParams{

		Context: ctx,
	}
}

// NewGetNetworkSmCellularUsageHistoryParamsWithHTTPClient creates a new GetNetworkSmCellularUsageHistoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkSmCellularUsageHistoryParamsWithHTTPClient(client *http.Client) *GetNetworkSmCellularUsageHistoryParams {
	var ()
	return &GetNetworkSmCellularUsageHistoryParams{
		HTTPClient: client,
	}
}

/*GetNetworkSmCellularUsageHistoryParams contains all the parameters to send to the API endpoint
for the get network sm cellular usage history operation typically these are written to a http.Request
*/
type GetNetworkSmCellularUsageHistoryParams struct {

	/*DeviceID*/
	DeviceID string
	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network sm cellular usage history params
func (o *GetNetworkSmCellularUsageHistoryParams) WithTimeout(timeout time.Duration) *GetNetworkSmCellularUsageHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network sm cellular usage history params
func (o *GetNetworkSmCellularUsageHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network sm cellular usage history params
func (o *GetNetworkSmCellularUsageHistoryParams) WithContext(ctx context.Context) *GetNetworkSmCellularUsageHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network sm cellular usage history params
func (o *GetNetworkSmCellularUsageHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network sm cellular usage history params
func (o *GetNetworkSmCellularUsageHistoryParams) WithHTTPClient(client *http.Client) *GetNetworkSmCellularUsageHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network sm cellular usage history params
func (o *GetNetworkSmCellularUsageHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get network sm cellular usage history params
func (o *GetNetworkSmCellularUsageHistoryParams) WithDeviceID(deviceID string) *GetNetworkSmCellularUsageHistoryParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get network sm cellular usage history params
func (o *GetNetworkSmCellularUsageHistoryParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithNetworkID adds the networkID to the get network sm cellular usage history params
func (o *GetNetworkSmCellularUsageHistoryParams) WithNetworkID(networkID string) *GetNetworkSmCellularUsageHistoryParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network sm cellular usage history params
func (o *GetNetworkSmCellularUsageHistoryParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSmCellularUsageHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
