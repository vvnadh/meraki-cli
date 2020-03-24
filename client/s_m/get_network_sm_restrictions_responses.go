// Code generated by go-swagger; DO NOT EDIT.

package s_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkSmRestrictionsReader is a Reader for the GetNetworkSmRestrictions structure.
type GetNetworkSmRestrictionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSmRestrictionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSmRestrictionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkSmRestrictionsOK creates a GetNetworkSmRestrictionsOK with default headers values
func NewGetNetworkSmRestrictionsOK() *GetNetworkSmRestrictionsOK {
	return &GetNetworkSmRestrictionsOK{}
}

/*GetNetworkSmRestrictionsOK handles this case with default header values.

Successful operation
*/
type GetNetworkSmRestrictionsOK struct {
	Payload interface{}
}

func (o *GetNetworkSmRestrictionsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sm/{deviceId}/restrictions][%d] getNetworkSmRestrictionsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSmRestrictionsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkSmRestrictionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
