// Code generated by go-swagger; DO NOT EDIT.

package link_aggregations

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

// NewCreateNetworkSwitchLinkAggregationParams creates a new CreateNetworkSwitchLinkAggregationParams object
// with the default values initialized.
func NewCreateNetworkSwitchLinkAggregationParams() *CreateNetworkSwitchLinkAggregationParams {
	var ()
	return &CreateNetworkSwitchLinkAggregationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNetworkSwitchLinkAggregationParamsWithTimeout creates a new CreateNetworkSwitchLinkAggregationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateNetworkSwitchLinkAggregationParamsWithTimeout(timeout time.Duration) *CreateNetworkSwitchLinkAggregationParams {
	var ()
	return &CreateNetworkSwitchLinkAggregationParams{

		timeout: timeout,
	}
}

// NewCreateNetworkSwitchLinkAggregationParamsWithContext creates a new CreateNetworkSwitchLinkAggregationParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateNetworkSwitchLinkAggregationParamsWithContext(ctx context.Context) *CreateNetworkSwitchLinkAggregationParams {
	var ()
	return &CreateNetworkSwitchLinkAggregationParams{

		Context: ctx,
	}
}

// NewCreateNetworkSwitchLinkAggregationParamsWithHTTPClient creates a new CreateNetworkSwitchLinkAggregationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateNetworkSwitchLinkAggregationParamsWithHTTPClient(client *http.Client) *CreateNetworkSwitchLinkAggregationParams {
	var ()
	return &CreateNetworkSwitchLinkAggregationParams{
		HTTPClient: client,
	}
}

/*CreateNetworkSwitchLinkAggregationParams contains all the parameters to send to the API endpoint
for the create network switch link aggregation operation typically these are written to a http.Request
*/
type CreateNetworkSwitchLinkAggregationParams struct {

	/*CreateNetworkSwitchLinkAggregation*/
	CreateNetworkSwitchLinkAggregation *models.CreateNetworkSwitchLinkAggregation
	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create network switch link aggregation params
func (o *CreateNetworkSwitchLinkAggregationParams) WithTimeout(timeout time.Duration) *CreateNetworkSwitchLinkAggregationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create network switch link aggregation params
func (o *CreateNetworkSwitchLinkAggregationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create network switch link aggregation params
func (o *CreateNetworkSwitchLinkAggregationParams) WithContext(ctx context.Context) *CreateNetworkSwitchLinkAggregationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create network switch link aggregation params
func (o *CreateNetworkSwitchLinkAggregationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create network switch link aggregation params
func (o *CreateNetworkSwitchLinkAggregationParams) WithHTTPClient(client *http.Client) *CreateNetworkSwitchLinkAggregationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create network switch link aggregation params
func (o *CreateNetworkSwitchLinkAggregationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateNetworkSwitchLinkAggregation adds the createNetworkSwitchLinkAggregation to the create network switch link aggregation params
func (o *CreateNetworkSwitchLinkAggregationParams) WithCreateNetworkSwitchLinkAggregation(createNetworkSwitchLinkAggregation *models.CreateNetworkSwitchLinkAggregation) *CreateNetworkSwitchLinkAggregationParams {
	o.SetCreateNetworkSwitchLinkAggregation(createNetworkSwitchLinkAggregation)
	return o
}

// SetCreateNetworkSwitchLinkAggregation adds the createNetworkSwitchLinkAggregation to the create network switch link aggregation params
func (o *CreateNetworkSwitchLinkAggregationParams) SetCreateNetworkSwitchLinkAggregation(createNetworkSwitchLinkAggregation *models.CreateNetworkSwitchLinkAggregation) {
	o.CreateNetworkSwitchLinkAggregation = createNetworkSwitchLinkAggregation
}

// WithNetworkID adds the networkID to the create network switch link aggregation params
func (o *CreateNetworkSwitchLinkAggregationParams) WithNetworkID(networkID string) *CreateNetworkSwitchLinkAggregationParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the create network switch link aggregation params
func (o *CreateNetworkSwitchLinkAggregationParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNetworkSwitchLinkAggregationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateNetworkSwitchLinkAggregation != nil {
		if err := r.SetBodyParam(o.CreateNetworkSwitchLinkAggregation); err != nil {
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
