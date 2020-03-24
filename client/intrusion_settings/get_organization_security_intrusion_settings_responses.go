// Code generated by go-swagger; DO NOT EDIT.

package intrusion_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetOrganizationSecurityIntrusionSettingsReader is a Reader for the GetOrganizationSecurityIntrusionSettings structure.
type GetOrganizationSecurityIntrusionSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationSecurityIntrusionSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationSecurityIntrusionSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganizationSecurityIntrusionSettingsOK creates a GetOrganizationSecurityIntrusionSettingsOK with default headers values
func NewGetOrganizationSecurityIntrusionSettingsOK() *GetOrganizationSecurityIntrusionSettingsOK {
	return &GetOrganizationSecurityIntrusionSettingsOK{}
}

/*GetOrganizationSecurityIntrusionSettingsOK handles this case with default header values.

Successful operation
*/
type GetOrganizationSecurityIntrusionSettingsOK struct {
	Payload interface{}
}

func (o *GetOrganizationSecurityIntrusionSettingsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/security/intrusionSettings][%d] getOrganizationSecurityIntrusionSettingsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationSecurityIntrusionSettingsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationSecurityIntrusionSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
