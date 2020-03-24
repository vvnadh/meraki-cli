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

// GetDeviceCameraQualityAndRetentionSettingsReader is a Reader for the GetDeviceCameraQualityAndRetentionSettings structure.
type GetDeviceCameraQualityAndRetentionSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceCameraQualityAndRetentionSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceCameraQualityAndRetentionSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDeviceCameraQualityAndRetentionSettingsOK creates a GetDeviceCameraQualityAndRetentionSettingsOK with default headers values
func NewGetDeviceCameraQualityAndRetentionSettingsOK() *GetDeviceCameraQualityAndRetentionSettingsOK {
	return &GetDeviceCameraQualityAndRetentionSettingsOK{}
}

/*GetDeviceCameraQualityAndRetentionSettingsOK handles this case with default header values.

Successful operation
*/
type GetDeviceCameraQualityAndRetentionSettingsOK struct {
	Payload interface{}
}

func (o *GetDeviceCameraQualityAndRetentionSettingsOK) Error() string {
	return fmt.Sprintf("[GET /devices/{serial}/camera/qualityAndRetentionSettings][%d] getDeviceCameraQualityAndRetentionSettingsOK  %+v", 200, o.Payload)
}

func (o *GetDeviceCameraQualityAndRetentionSettingsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetDeviceCameraQualityAndRetentionSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
