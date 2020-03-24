// Code generated by go-swagger; DO NOT EDIT.

package net_flow_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkNetflowSettingsReader is a Reader for the GetNetworkNetflowSettings structure.
type GetNetworkNetflowSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkNetflowSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkNetflowSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkNetflowSettingsOK creates a GetNetworkNetflowSettingsOK with default headers values
func NewGetNetworkNetflowSettingsOK() *GetNetworkNetflowSettingsOK {
	return &GetNetworkNetflowSettingsOK{}
}

/*GetNetworkNetflowSettingsOK handles this case with default header values.

Successful operation
*/
type GetNetworkNetflowSettingsOK struct {
	Payload interface{}
}

func (o *GetNetworkNetflowSettingsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/netflowSettings][%d] getNetworkNetflowSettingsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkNetflowSettingsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkNetflowSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
