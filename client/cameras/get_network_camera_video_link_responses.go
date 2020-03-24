// Code generated by go-swagger; DO NOT EDIT.

package cameras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkCameraVideoLinkReader is a Reader for the GetNetworkCameraVideoLink structure.
type GetNetworkCameraVideoLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkCameraVideoLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkCameraVideoLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkCameraVideoLinkOK creates a GetNetworkCameraVideoLinkOK with default headers values
func NewGetNetworkCameraVideoLinkOK() *GetNetworkCameraVideoLinkOK {
	return &GetNetworkCameraVideoLinkOK{}
}

/*GetNetworkCameraVideoLinkOK handles this case with default header values.

Successful operation
*/
type GetNetworkCameraVideoLinkOK struct {
	Payload interface{}
}

func (o *GetNetworkCameraVideoLinkOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/cameras/{serial}/videoLink][%d] getNetworkCameraVideoLinkOK  %+v", 200, o.Payload)
}

func (o *GetNetworkCameraVideoLinkOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkCameraVideoLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
