// Code generated by go-swagger; DO NOT EDIT.

package content_filtering_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateNetworkContentFilteringReader is a Reader for the UpdateNetworkContentFiltering structure.
type UpdateNetworkContentFilteringReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkContentFilteringReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkContentFilteringOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNetworkContentFilteringOK creates a UpdateNetworkContentFilteringOK with default headers values
func NewUpdateNetworkContentFilteringOK() *UpdateNetworkContentFilteringOK {
	return &UpdateNetworkContentFilteringOK{}
}

/*UpdateNetworkContentFilteringOK handles this case with default header values.

Successful operation
*/
type UpdateNetworkContentFilteringOK struct {
	Payload interface{}
}

func (o *UpdateNetworkContentFilteringOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/contentFiltering][%d] updateNetworkContentFilteringOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkContentFilteringOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkContentFilteringOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
