// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// NewPutNetworksNetworkIDParams creates a new PutNetworksNetworkIDParams object
// with the default values initialized.
func NewPutNetworksNetworkIDParams() *PutNetworksNetworkIDParams {
	var ()
	return &PutNetworksNetworkIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutNetworksNetworkIDParamsWithTimeout creates a new PutNetworksNetworkIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutNetworksNetworkIDParamsWithTimeout(timeout time.Duration) *PutNetworksNetworkIDParams {
	var ()
	return &PutNetworksNetworkIDParams{

		timeout: timeout,
	}
}

// NewPutNetworksNetworkIDParamsWithContext creates a new PutNetworksNetworkIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutNetworksNetworkIDParamsWithContext(ctx context.Context) *PutNetworksNetworkIDParams {
	var ()
	return &PutNetworksNetworkIDParams{

		Context: ctx,
	}
}

// NewPutNetworksNetworkIDParamsWithHTTPClient creates a new PutNetworksNetworkIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutNetworksNetworkIDParamsWithHTTPClient(client *http.Client) *PutNetworksNetworkIDParams {
	var ()
	return &PutNetworksNetworkIDParams{
		HTTPClient: client,
	}
}

/*PutNetworksNetworkIDParams contains all the parameters to send to the API endpoint
for the put networks network ID operation typically these are written to a http.Request
*/
type PutNetworksNetworkIDParams struct {

	/*Network
	  Full desired configuration of the network

	*/
	Network *models.Network
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put networks network ID params
func (o *PutNetworksNetworkIDParams) WithTimeout(timeout time.Duration) *PutNetworksNetworkIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put networks network ID params
func (o *PutNetworksNetworkIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put networks network ID params
func (o *PutNetworksNetworkIDParams) WithContext(ctx context.Context) *PutNetworksNetworkIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put networks network ID params
func (o *PutNetworksNetworkIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put networks network ID params
func (o *PutNetworksNetworkIDParams) WithHTTPClient(client *http.Client) *PutNetworksNetworkIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put networks network ID params
func (o *PutNetworksNetworkIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetwork adds the network to the put networks network ID params
func (o *PutNetworksNetworkIDParams) WithNetwork(network *models.Network) *PutNetworksNetworkIDParams {
	o.SetNetwork(network)
	return o
}

// SetNetwork adds the network to the put networks network ID params
func (o *PutNetworksNetworkIDParams) SetNetwork(network *models.Network) {
	o.Network = network
}

// WithNetworkID adds the networkID to the put networks network ID params
func (o *PutNetworksNetworkIDParams) WithNetworkID(networkID string) *PutNetworksNetworkIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put networks network ID params
func (o *PutNetworksNetworkIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutNetworksNetworkIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Network != nil {
		if err := r.SetBodyParam(o.Network); err != nil {
			return err
		}
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}