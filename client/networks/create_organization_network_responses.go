// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateOrganizationNetworkReader is a Reader for the CreateOrganizationNetwork structure.
type CreateOrganizationNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrganizationNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateOrganizationNetworkCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateOrganizationNetworkCreated creates a CreateOrganizationNetworkCreated with default headers values
func NewCreateOrganizationNetworkCreated() *CreateOrganizationNetworkCreated {
	return &CreateOrganizationNetworkCreated{}
}

/*CreateOrganizationNetworkCreated handles this case with default header values.

Successful operation
*/
type CreateOrganizationNetworkCreated struct {
	Payload interface{}
}

func (o *CreateOrganizationNetworkCreated) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/networks][%d] createOrganizationNetworkCreated  %+v", 201, o.Payload)
}

func (o *CreateOrganizationNetworkCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateOrganizationNetworkCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
