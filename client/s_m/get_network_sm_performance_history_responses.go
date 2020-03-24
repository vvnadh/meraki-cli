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

// GetNetworkSmPerformanceHistoryReader is a Reader for the GetNetworkSmPerformanceHistory structure.
type GetNetworkSmPerformanceHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSmPerformanceHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSmPerformanceHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkSmPerformanceHistoryOK creates a GetNetworkSmPerformanceHistoryOK with default headers values
func NewGetNetworkSmPerformanceHistoryOK() *GetNetworkSmPerformanceHistoryOK {
	return &GetNetworkSmPerformanceHistoryOK{}
}

/*GetNetworkSmPerformanceHistoryOK handles this case with default header values.

Successful operation
*/
type GetNetworkSmPerformanceHistoryOK struct {
	Payload interface{}
}

func (o *GetNetworkSmPerformanceHistoryOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/sm/{id}/performanceHistory][%d] getNetworkSmPerformanceHistoryOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSmPerformanceHistoryOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkSmPerformanceHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
