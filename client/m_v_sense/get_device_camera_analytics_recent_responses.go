// Code generated by go-swagger; DO NOT EDIT.

package m_v_sense

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetDeviceCameraAnalyticsRecentReader is a Reader for the GetDeviceCameraAnalyticsRecent structure.
type GetDeviceCameraAnalyticsRecentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceCameraAnalyticsRecentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceCameraAnalyticsRecentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDeviceCameraAnalyticsRecentOK creates a GetDeviceCameraAnalyticsRecentOK with default headers values
func NewGetDeviceCameraAnalyticsRecentOK() *GetDeviceCameraAnalyticsRecentOK {
	return &GetDeviceCameraAnalyticsRecentOK{}
}

/*GetDeviceCameraAnalyticsRecentOK handles this case with default header values.

Successful operation
*/
type GetDeviceCameraAnalyticsRecentOK struct {
	Payload interface{}
}

func (o *GetDeviceCameraAnalyticsRecentOK) Error() string {
	return fmt.Sprintf("[GET /devices/{serial}/camera/analytics/recent][%d] getDeviceCameraAnalyticsRecentOK  %+v", 200, o.Payload)
}

func (o *GetDeviceCameraAnalyticsRecentOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetDeviceCameraAnalyticsRecentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
