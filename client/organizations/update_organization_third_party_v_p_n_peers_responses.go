// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateOrganizationThirdPartyVPNPeersReader is a Reader for the UpdateOrganizationThirdPartyVPNPeers structure.
type UpdateOrganizationThirdPartyVPNPeersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrganizationThirdPartyVPNPeersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrganizationThirdPartyVPNPeersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateOrganizationThirdPartyVPNPeersOK creates a UpdateOrganizationThirdPartyVPNPeersOK with default headers values
func NewUpdateOrganizationThirdPartyVPNPeersOK() *UpdateOrganizationThirdPartyVPNPeersOK {
	return &UpdateOrganizationThirdPartyVPNPeersOK{}
}

/*UpdateOrganizationThirdPartyVPNPeersOK handles this case with default header values.

Successful operation
*/
type UpdateOrganizationThirdPartyVPNPeersOK struct {
	Payload interface{}
}

func (o *UpdateOrganizationThirdPartyVPNPeersOK) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organizationId}/thirdPartyVPNPeers][%d] updateOrganizationThirdPartyVPNPeersOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationThirdPartyVPNPeersOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateOrganizationThirdPartyVPNPeersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
