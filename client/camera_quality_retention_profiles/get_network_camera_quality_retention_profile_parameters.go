// Code generated by go-swagger; DO NOT EDIT.

package camera_quality_retention_profiles

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

// NewGetNetworkCameraQualityRetentionProfileParams creates a new GetNetworkCameraQualityRetentionProfileParams object
// with the default values initialized.
func NewGetNetworkCameraQualityRetentionProfileParams() *GetNetworkCameraQualityRetentionProfileParams {
	var ()
	return &GetNetworkCameraQualityRetentionProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkCameraQualityRetentionProfileParamsWithTimeout creates a new GetNetworkCameraQualityRetentionProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkCameraQualityRetentionProfileParamsWithTimeout(timeout time.Duration) *GetNetworkCameraQualityRetentionProfileParams {
	var ()
	return &GetNetworkCameraQualityRetentionProfileParams{

		timeout: timeout,
	}
}

// NewGetNetworkCameraQualityRetentionProfileParamsWithContext creates a new GetNetworkCameraQualityRetentionProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkCameraQualityRetentionProfileParamsWithContext(ctx context.Context) *GetNetworkCameraQualityRetentionProfileParams {
	var ()
	return &GetNetworkCameraQualityRetentionProfileParams{

		Context: ctx,
	}
}

// NewGetNetworkCameraQualityRetentionProfileParamsWithHTTPClient creates a new GetNetworkCameraQualityRetentionProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkCameraQualityRetentionProfileParamsWithHTTPClient(client *http.Client) *GetNetworkCameraQualityRetentionProfileParams {
	var ()
	return &GetNetworkCameraQualityRetentionProfileParams{
		HTTPClient: client,
	}
}

/*GetNetworkCameraQualityRetentionProfileParams contains all the parameters to send to the API endpoint
for the get network camera quality retention profile operation typically these are written to a http.Request
*/
type GetNetworkCameraQualityRetentionProfileParams struct {

	/*NetworkID*/
	NetworkID string
	/*QualityRetentionProfileID*/
	QualityRetentionProfileID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network camera quality retention profile params
func (o *GetNetworkCameraQualityRetentionProfileParams) WithTimeout(timeout time.Duration) *GetNetworkCameraQualityRetentionProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network camera quality retention profile params
func (o *GetNetworkCameraQualityRetentionProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network camera quality retention profile params
func (o *GetNetworkCameraQualityRetentionProfileParams) WithContext(ctx context.Context) *GetNetworkCameraQualityRetentionProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network camera quality retention profile params
func (o *GetNetworkCameraQualityRetentionProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network camera quality retention profile params
func (o *GetNetworkCameraQualityRetentionProfileParams) WithHTTPClient(client *http.Client) *GetNetworkCameraQualityRetentionProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network camera quality retention profile params
func (o *GetNetworkCameraQualityRetentionProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network camera quality retention profile params
func (o *GetNetworkCameraQualityRetentionProfileParams) WithNetworkID(networkID string) *GetNetworkCameraQualityRetentionProfileParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network camera quality retention profile params
func (o *GetNetworkCameraQualityRetentionProfileParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithQualityRetentionProfileID adds the qualityRetentionProfileID to the get network camera quality retention profile params
func (o *GetNetworkCameraQualityRetentionProfileParams) WithQualityRetentionProfileID(qualityRetentionProfileID string) *GetNetworkCameraQualityRetentionProfileParams {
	o.SetQualityRetentionProfileID(qualityRetentionProfileID)
	return o
}

// SetQualityRetentionProfileID adds the qualityRetentionProfileId to the get network camera quality retention profile params
func (o *GetNetworkCameraQualityRetentionProfileParams) SetQualityRetentionProfileID(qualityRetentionProfileID string) {
	o.QualityRetentionProfileID = qualityRetentionProfileID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkCameraQualityRetentionProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param qualityRetentionProfileId
	if err := r.SetPathParam("qualityRetentionProfileId", o.QualityRetentionProfileID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
