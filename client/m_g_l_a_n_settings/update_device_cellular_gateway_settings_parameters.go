// Code generated by go-swagger; DO NOT EDIT.

package m_g_l_a_n_settings

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

// NewUpdateDeviceCellularGatewaySettingsParams creates a new UpdateDeviceCellularGatewaySettingsParams object
// with the default values initialized.
func NewUpdateDeviceCellularGatewaySettingsParams() *UpdateDeviceCellularGatewaySettingsParams {
	var ()
	return &UpdateDeviceCellularGatewaySettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeviceCellularGatewaySettingsParamsWithTimeout creates a new UpdateDeviceCellularGatewaySettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateDeviceCellularGatewaySettingsParamsWithTimeout(timeout time.Duration) *UpdateDeviceCellularGatewaySettingsParams {
	var ()
	return &UpdateDeviceCellularGatewaySettingsParams{

		timeout: timeout,
	}
}

// NewUpdateDeviceCellularGatewaySettingsParamsWithContext creates a new UpdateDeviceCellularGatewaySettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateDeviceCellularGatewaySettingsParamsWithContext(ctx context.Context) *UpdateDeviceCellularGatewaySettingsParams {
	var ()
	return &UpdateDeviceCellularGatewaySettingsParams{

		Context: ctx,
	}
}

// NewUpdateDeviceCellularGatewaySettingsParamsWithHTTPClient creates a new UpdateDeviceCellularGatewaySettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateDeviceCellularGatewaySettingsParamsWithHTTPClient(client *http.Client) *UpdateDeviceCellularGatewaySettingsParams {
	var ()
	return &UpdateDeviceCellularGatewaySettingsParams{
		HTTPClient: client,
	}
}

/*UpdateDeviceCellularGatewaySettingsParams contains all the parameters to send to the API endpoint
for the update device cellular gateway settings operation typically these are written to a http.Request
*/
type UpdateDeviceCellularGatewaySettingsParams struct {

	/*Serial*/
	Serial string
	/*UpdateDeviceCellularGatewaySettings*/
	UpdateDeviceCellularGatewaySettings *models.UpdateDeviceCellularGatewaySettings

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update device cellular gateway settings params
func (o *UpdateDeviceCellularGatewaySettingsParams) WithTimeout(timeout time.Duration) *UpdateDeviceCellularGatewaySettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update device cellular gateway settings params
func (o *UpdateDeviceCellularGatewaySettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update device cellular gateway settings params
func (o *UpdateDeviceCellularGatewaySettingsParams) WithContext(ctx context.Context) *UpdateDeviceCellularGatewaySettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update device cellular gateway settings params
func (o *UpdateDeviceCellularGatewaySettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update device cellular gateway settings params
func (o *UpdateDeviceCellularGatewaySettingsParams) WithHTTPClient(client *http.Client) *UpdateDeviceCellularGatewaySettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update device cellular gateway settings params
func (o *UpdateDeviceCellularGatewaySettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the update device cellular gateway settings params
func (o *UpdateDeviceCellularGatewaySettingsParams) WithSerial(serial string) *UpdateDeviceCellularGatewaySettingsParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the update device cellular gateway settings params
func (o *UpdateDeviceCellularGatewaySettingsParams) SetSerial(serial string) {
	o.Serial = serial
}

// WithUpdateDeviceCellularGatewaySettings adds the updateDeviceCellularGatewaySettings to the update device cellular gateway settings params
func (o *UpdateDeviceCellularGatewaySettingsParams) WithUpdateDeviceCellularGatewaySettings(updateDeviceCellularGatewaySettings *models.UpdateDeviceCellularGatewaySettings) *UpdateDeviceCellularGatewaySettingsParams {
	o.SetUpdateDeviceCellularGatewaySettings(updateDeviceCellularGatewaySettings)
	return o
}

// SetUpdateDeviceCellularGatewaySettings adds the updateDeviceCellularGatewaySettings to the update device cellular gateway settings params
func (o *UpdateDeviceCellularGatewaySettingsParams) SetUpdateDeviceCellularGatewaySettings(updateDeviceCellularGatewaySettings *models.UpdateDeviceCellularGatewaySettings) {
	o.UpdateDeviceCellularGatewaySettings = updateDeviceCellularGatewaySettings
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeviceCellularGatewaySettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serial
	if err := r.SetPathParam("serial", o.Serial); err != nil {
		return err
	}

	if o.UpdateDeviceCellularGatewaySettings != nil {
		if err := r.SetBodyParam(o.UpdateDeviceCellularGatewaySettings); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
