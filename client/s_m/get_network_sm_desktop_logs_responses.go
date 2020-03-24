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

// GetNetworkSmDesktopLogsReader is a Reader for the GetNetworkSmDesktopLogs structure.
type GetNetworkSmDesktopLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSmDesktopLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSmDesktopLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkSmDesktopLogsOK creates a GetNetworkSmDesktopLogsOK with default headers values
func NewGetNetworkSmDesktopLogsOK() *GetNetworkSmDesktopLogsOK {
	return &GetNetworkSmDesktopLogsOK{}
}

/*GetNetworkSmDesktopLogsOK handles this case with default header values.

Successful operation
*/
type GetNetworkSmDesktopLogsOK struct {
	Payload interface{}
}

func (o *GetNetworkSmDesktopLogsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/sm/{id}/desktopLogs][%d] getNetworkSmDesktopLogsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSmDesktopLogsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkSmDesktopLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}