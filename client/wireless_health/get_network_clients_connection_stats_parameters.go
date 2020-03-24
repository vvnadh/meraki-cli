// Code generated by go-swagger; DO NOT EDIT.

package wireless_health

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
	"github.com/go-openapi/swag"
)

// NewGetNetworkClientsConnectionStatsParams creates a new GetNetworkClientsConnectionStatsParams object
// with the default values initialized.
func NewGetNetworkClientsConnectionStatsParams() *GetNetworkClientsConnectionStatsParams {
	var ()
	return &GetNetworkClientsConnectionStatsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkClientsConnectionStatsParamsWithTimeout creates a new GetNetworkClientsConnectionStatsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkClientsConnectionStatsParamsWithTimeout(timeout time.Duration) *GetNetworkClientsConnectionStatsParams {
	var ()
	return &GetNetworkClientsConnectionStatsParams{

		timeout: timeout,
	}
}

// NewGetNetworkClientsConnectionStatsParamsWithContext creates a new GetNetworkClientsConnectionStatsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkClientsConnectionStatsParamsWithContext(ctx context.Context) *GetNetworkClientsConnectionStatsParams {
	var ()
	return &GetNetworkClientsConnectionStatsParams{

		Context: ctx,
	}
}

// NewGetNetworkClientsConnectionStatsParamsWithHTTPClient creates a new GetNetworkClientsConnectionStatsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkClientsConnectionStatsParamsWithHTTPClient(client *http.Client) *GetNetworkClientsConnectionStatsParams {
	var ()
	return &GetNetworkClientsConnectionStatsParams{
		HTTPClient: client,
	}
}

/*GetNetworkClientsConnectionStatsParams contains all the parameters to send to the API endpoint
for the get network clients connection stats operation typically these are written to a http.Request
*/
type GetNetworkClientsConnectionStatsParams struct {

	/*ApTag
	  Filter results by AP Tag

	*/
	ApTag *string
	/*NetworkID*/
	NetworkID string
	/*Ssid
	  Filter results by SSID

	*/
	Ssid *int32
	/*T0
	  The beginning of the timespan for the data. The maximum lookback period is 180 days from today.

	*/
	T0 *string
	/*T1
	  The end of the timespan for the data. t1 can be a maximum of 7 days after t0.

	*/
	T1 *string
	/*Timespan
	  The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.

	*/
	Timespan *float64
	/*Vlan
	  Filter results by VLAN

	*/
	Vlan *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network clients connection stats params
func (o *GetNetworkClientsConnectionStatsParams) WithTimeout(timeout time.Duration) *GetNetworkClientsConnectionStatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network clients connection stats params
func (o *GetNetworkClientsConnectionStatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network clients connection stats params
func (o *GetNetworkClientsConnectionStatsParams) WithContext(ctx context.Context) *GetNetworkClientsConnectionStatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network clients connection stats params
func (o *GetNetworkClientsConnectionStatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network clients connection stats params
func (o *GetNetworkClientsConnectionStatsParams) WithHTTPClient(client *http.Client) *GetNetworkClientsConnectionStatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network clients connection stats params
func (o *GetNetworkClientsConnectionStatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApTag adds the apTag to the get network clients connection stats params
func (o *GetNetworkClientsConnectionStatsParams) WithApTag(apTag *string) *GetNetworkClientsConnectionStatsParams {
	o.SetApTag(apTag)
	return o
}

// SetApTag adds the apTag to the get network clients connection stats params
func (o *GetNetworkClientsConnectionStatsParams) SetApTag(apTag *string) {
	o.ApTag = apTag
}

// WithNetworkID adds the networkID to the get network clients connection stats params
func (o *GetNetworkClientsConnectionStatsParams) WithNetworkID(networkID string) *GetNetworkClientsConnectionStatsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network clients connection stats params
func (o *GetNetworkClientsConnectionStatsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithSsid adds the ssid to the get network clients connection stats params
func (o *GetNetworkClientsConnectionStatsParams) WithSsid(ssid *int32) *GetNetworkClientsConnectionStatsParams {
	o.SetSsid(ssid)
	return o
}

// SetSsid adds the ssid to the get network clients connection stats params
func (o *GetNetworkClientsConnectionStatsParams) SetSsid(ssid *int32) {
	o.Ssid = ssid
}

// WithT0 adds the t0 to the get network clients connection stats params
func (o *GetNetworkClientsConnectionStatsParams) WithT0(t0 *string) *GetNetworkClientsConnectionStatsParams {
	o.SetT0(t0)
	return o
}

// SetT0 adds the t0 to the get network clients connection stats params
func (o *GetNetworkClientsConnectionStatsParams) SetT0(t0 *string) {
	o.T0 = t0
}

// WithT1 adds the t1 to the get network clients connection stats params
func (o *GetNetworkClientsConnectionStatsParams) WithT1(t1 *string) *GetNetworkClientsConnectionStatsParams {
	o.SetT1(t1)
	return o
}

// SetT1 adds the t1 to the get network clients connection stats params
func (o *GetNetworkClientsConnectionStatsParams) SetT1(t1 *string) {
	o.T1 = t1
}

// WithTimespan adds the timespan to the get network clients connection stats params
func (o *GetNetworkClientsConnectionStatsParams) WithTimespan(timespan *float64) *GetNetworkClientsConnectionStatsParams {
	o.SetTimespan(timespan)
	return o
}

// SetTimespan adds the timespan to the get network clients connection stats params
func (o *GetNetworkClientsConnectionStatsParams) SetTimespan(timespan *float64) {
	o.Timespan = timespan
}

// WithVlan adds the vlan to the get network clients connection stats params
func (o *GetNetworkClientsConnectionStatsParams) WithVlan(vlan *int32) *GetNetworkClientsConnectionStatsParams {
	o.SetVlan(vlan)
	return o
}

// SetVlan adds the vlan to the get network clients connection stats params
func (o *GetNetworkClientsConnectionStatsParams) SetVlan(vlan *int32) {
	o.Vlan = vlan
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkClientsConnectionStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApTag != nil {

		// query param apTag
		var qrApTag string
		if o.ApTag != nil {
			qrApTag = *o.ApTag
		}
		qApTag := qrApTag
		if qApTag != "" {
			if err := r.SetQueryParam("apTag", qApTag); err != nil {
				return err
			}
		}

	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if o.Ssid != nil {

		// query param ssid
		var qrSsid int32
		if o.Ssid != nil {
			qrSsid = *o.Ssid
		}
		qSsid := swag.FormatInt32(qrSsid)
		if qSsid != "" {
			if err := r.SetQueryParam("ssid", qSsid); err != nil {
				return err
			}
		}

	}

	if o.T0 != nil {

		// query param t0
		var qrT0 string
		if o.T0 != nil {
			qrT0 = *o.T0
		}
		qT0 := qrT0
		if qT0 != "" {
			if err := r.SetQueryParam("t0", qT0); err != nil {
				return err
			}
		}

	}

	if o.T1 != nil {

		// query param t1
		var qrT1 string
		if o.T1 != nil {
			qrT1 = *o.T1
		}
		qT1 := qrT1
		if qT1 != "" {
			if err := r.SetQueryParam("t1", qT1); err != nil {
				return err
			}
		}

	}

	if o.Timespan != nil {

		// query param timespan
		var qrTimespan float64
		if o.Timespan != nil {
			qrTimespan = *o.Timespan
		}
		qTimespan := swag.FormatFloat64(qrTimespan)
		if qTimespan != "" {
			if err := r.SetQueryParam("timespan", qTimespan); err != nil {
				return err
			}
		}

	}

	if o.Vlan != nil {

		// query param vlan
		var qrVlan int32
		if o.Vlan != nil {
			qrVlan = *o.Vlan
		}
		qVlan := swag.FormatInt32(qrVlan)
		if qVlan != "" {
			if err := r.SetQueryParam("vlan", qVlan); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
