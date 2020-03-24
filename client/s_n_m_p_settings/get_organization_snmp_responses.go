// Code generated by go-swagger; DO NOT EDIT.

package s_n_m_p_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetOrganizationSnmpReader is a Reader for the GetOrganizationSnmp structure.
type GetOrganizationSnmpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationSnmpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationSnmpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganizationSnmpOK creates a GetOrganizationSnmpOK with default headers values
func NewGetOrganizationSnmpOK() *GetOrganizationSnmpOK {
	return &GetOrganizationSnmpOK{}
}

/*GetOrganizationSnmpOK handles this case with default header values.

Successful operation
*/
type GetOrganizationSnmpOK struct {
	Payload interface{}
}

func (o *GetOrganizationSnmpOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/snmp][%d] getOrganizationSnmpOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationSnmpOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationSnmpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
