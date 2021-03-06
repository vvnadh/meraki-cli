// Code generated by go-swagger; DO NOT EDIT.

package switch_ports

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

// NewGetDeviceSwitchPortStatusesParams creates a new GetDeviceSwitchPortStatusesParams object
// with the default values initialized.
func NewGetDeviceSwitchPortStatusesParams() *GetDeviceSwitchPortStatusesParams {
	var ()
	return &GetDeviceSwitchPortStatusesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceSwitchPortStatusesParamsWithTimeout creates a new GetDeviceSwitchPortStatusesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeviceSwitchPortStatusesParamsWithTimeout(timeout time.Duration) *GetDeviceSwitchPortStatusesParams {
	var ()
	return &GetDeviceSwitchPortStatusesParams{

		timeout: timeout,
	}
}

// NewGetDeviceSwitchPortStatusesParamsWithContext creates a new GetDeviceSwitchPortStatusesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeviceSwitchPortStatusesParamsWithContext(ctx context.Context) *GetDeviceSwitchPortStatusesParams {
	var ()
	return &GetDeviceSwitchPortStatusesParams{

		Context: ctx,
	}
}

// NewGetDeviceSwitchPortStatusesParamsWithHTTPClient creates a new GetDeviceSwitchPortStatusesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeviceSwitchPortStatusesParamsWithHTTPClient(client *http.Client) *GetDeviceSwitchPortStatusesParams {
	var ()
	return &GetDeviceSwitchPortStatusesParams{
		HTTPClient: client,
	}
}

/*GetDeviceSwitchPortStatusesParams contains all the parameters to send to the API endpoint
for the get device switch port statuses operation typically these are written to a http.Request
*/
type GetDeviceSwitchPortStatusesParams struct {

	/*Serial*/
	Serial string
	/*T0
	  The beginning of the timespan for the data. The maximum lookback period is 31 days from today.

	*/
	T0 *string
	/*Timespan
	  The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.

	*/
	Timespan *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get device switch port statuses params
func (o *GetDeviceSwitchPortStatusesParams) WithTimeout(timeout time.Duration) *GetDeviceSwitchPortStatusesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device switch port statuses params
func (o *GetDeviceSwitchPortStatusesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device switch port statuses params
func (o *GetDeviceSwitchPortStatusesParams) WithContext(ctx context.Context) *GetDeviceSwitchPortStatusesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device switch port statuses params
func (o *GetDeviceSwitchPortStatusesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device switch port statuses params
func (o *GetDeviceSwitchPortStatusesParams) WithHTTPClient(client *http.Client) *GetDeviceSwitchPortStatusesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device switch port statuses params
func (o *GetDeviceSwitchPortStatusesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the get device switch port statuses params
func (o *GetDeviceSwitchPortStatusesParams) WithSerial(serial string) *GetDeviceSwitchPortStatusesParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the get device switch port statuses params
func (o *GetDeviceSwitchPortStatusesParams) SetSerial(serial string) {
	o.Serial = serial
}

// WithT0 adds the t0 to the get device switch port statuses params
func (o *GetDeviceSwitchPortStatusesParams) WithT0(t0 *string) *GetDeviceSwitchPortStatusesParams {
	o.SetT0(t0)
	return o
}

// SetT0 adds the t0 to the get device switch port statuses params
func (o *GetDeviceSwitchPortStatusesParams) SetT0(t0 *string) {
	o.T0 = t0
}

// WithTimespan adds the timespan to the get device switch port statuses params
func (o *GetDeviceSwitchPortStatusesParams) WithTimespan(timespan *float64) *GetDeviceSwitchPortStatusesParams {
	o.SetTimespan(timespan)
	return o
}

// SetTimespan adds the timespan to the get device switch port statuses params
func (o *GetDeviceSwitchPortStatusesParams) SetTimespan(timespan *float64) {
	o.Timespan = timespan
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceSwitchPortStatusesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serial
	if err := r.SetPathParam("serial", o.Serial); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
