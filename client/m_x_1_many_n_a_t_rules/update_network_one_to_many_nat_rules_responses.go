// Code generated by go-swagger; DO NOT EDIT.

package m_x_1_many_n_a_t_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateNetworkOneToManyNatRulesReader is a Reader for the UpdateNetworkOneToManyNatRules structure.
type UpdateNetworkOneToManyNatRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkOneToManyNatRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkOneToManyNatRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNetworkOneToManyNatRulesOK creates a UpdateNetworkOneToManyNatRulesOK with default headers values
func NewUpdateNetworkOneToManyNatRulesOK() *UpdateNetworkOneToManyNatRulesOK {
	return &UpdateNetworkOneToManyNatRulesOK{}
}

/*UpdateNetworkOneToManyNatRulesOK handles this case with default header values.

Successful operation
*/
type UpdateNetworkOneToManyNatRulesOK struct {
	Payload interface{}
}

func (o *UpdateNetworkOneToManyNatRulesOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/oneToManyNatRules][%d] updateNetworkOneToManyNatRulesOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkOneToManyNatRulesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkOneToManyNatRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
