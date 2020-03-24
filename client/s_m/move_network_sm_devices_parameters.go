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

	"github.com/cisco-sso/meraki-cli/models"
)

// NewMoveNetworkSmDevicesParams creates a new MoveNetworkSmDevicesParams object
// with the default values initialized.
func NewMoveNetworkSmDevicesParams() *MoveNetworkSmDevicesParams {
	var ()
	return &MoveNetworkSmDevicesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMoveNetworkSmDevicesParamsWithTimeout creates a new MoveNetworkSmDevicesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMoveNetworkSmDevicesParamsWithTimeout(timeout time.Duration) *MoveNetworkSmDevicesParams {
	var ()
	return &MoveNetworkSmDevicesParams{

		timeout: timeout,
	}
}

// NewMoveNetworkSmDevicesParamsWithContext creates a new MoveNetworkSmDevicesParams object
// with the default values initialized, and the ability to set a context for a request
func NewMoveNetworkSmDevicesParamsWithContext(ctx context.Context) *MoveNetworkSmDevicesParams {
	var ()
	return &MoveNetworkSmDevicesParams{

		Context: ctx,
	}
}

// NewMoveNetworkSmDevicesParamsWithHTTPClient creates a new MoveNetworkSmDevicesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMoveNetworkSmDevicesParamsWithHTTPClient(client *http.Client) *MoveNetworkSmDevicesParams {
	var ()
	return &MoveNetworkSmDevicesParams{
		HTTPClient: client,
	}
}

/*MoveNetworkSmDevicesParams contains all the parameters to send to the API endpoint
for the move network sm devices operation typically these are written to a http.Request
*/
type MoveNetworkSmDevicesParams struct {

	/*MoveNetworkSmDevices*/
	MoveNetworkSmDevices *models.MoveNetworkSmDevices
	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the move network sm devices params
func (o *MoveNetworkSmDevicesParams) WithTimeout(timeout time.Duration) *MoveNetworkSmDevicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the move network sm devices params
func (o *MoveNetworkSmDevicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the move network sm devices params
func (o *MoveNetworkSmDevicesParams) WithContext(ctx context.Context) *MoveNetworkSmDevicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the move network sm devices params
func (o *MoveNetworkSmDevicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the move network sm devices params
func (o *MoveNetworkSmDevicesParams) WithHTTPClient(client *http.Client) *MoveNetworkSmDevicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the move network sm devices params
func (o *MoveNetworkSmDevicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoveNetworkSmDevices adds the moveNetworkSmDevices to the move network sm devices params
func (o *MoveNetworkSmDevicesParams) WithMoveNetworkSmDevices(moveNetworkSmDevices *models.MoveNetworkSmDevices) *MoveNetworkSmDevicesParams {
	o.SetMoveNetworkSmDevices(moveNetworkSmDevices)
	return o
}

// SetMoveNetworkSmDevices adds the moveNetworkSmDevices to the move network sm devices params
func (o *MoveNetworkSmDevicesParams) SetMoveNetworkSmDevices(moveNetworkSmDevices *models.MoveNetworkSmDevices) {
	o.MoveNetworkSmDevices = moveNetworkSmDevices
}

// WithNetworkID adds the networkID to the move network sm devices params
func (o *MoveNetworkSmDevicesParams) WithNetworkID(networkID string) *MoveNetworkSmDevicesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the move network sm devices params
func (o *MoveNetworkSmDevicesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *MoveNetworkSmDevicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MoveNetworkSmDevices != nil {
		if err := r.SetBodyParam(o.MoveNetworkSmDevices); err != nil {
			return err
		}
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
