// Code generated by go-swagger; DO NOT EDIT.

package switch_ports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetDeviceSwitchPortStatusesReader is a Reader for the GetDeviceSwitchPortStatuses structure.
type GetDeviceSwitchPortStatusesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceSwitchPortStatusesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceSwitchPortStatusesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDeviceSwitchPortStatusesOK creates a GetDeviceSwitchPortStatusesOK with default headers values
func NewGetDeviceSwitchPortStatusesOK() *GetDeviceSwitchPortStatusesOK {
	return &GetDeviceSwitchPortStatusesOK{}
}

/*GetDeviceSwitchPortStatusesOK handles this case with default header values.

Successful operation
*/
type GetDeviceSwitchPortStatusesOK struct {
	Payload interface{}
}

func (o *GetDeviceSwitchPortStatusesOK) Error() string {
	return fmt.Sprintf("[GET /devices/{serial}/switchPortStatuses][%d] getDeviceSwitchPortStatusesOK  %+v", 200, o.Payload)
}

func (o *GetDeviceSwitchPortStatusesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetDeviceSwitchPortStatusesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
