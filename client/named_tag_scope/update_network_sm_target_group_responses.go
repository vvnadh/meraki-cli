// Code generated by go-swagger; DO NOT EDIT.

package named_tag_scope

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateNetworkSmTargetGroupReader is a Reader for the UpdateNetworkSmTargetGroup structure.
type UpdateNetworkSmTargetGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkSmTargetGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkSmTargetGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNetworkSmTargetGroupOK creates a UpdateNetworkSmTargetGroupOK with default headers values
func NewUpdateNetworkSmTargetGroupOK() *UpdateNetworkSmTargetGroupOK {
	return &UpdateNetworkSmTargetGroupOK{}
}

/*UpdateNetworkSmTargetGroupOK handles this case with default header values.

Successful operation
*/
type UpdateNetworkSmTargetGroupOK struct {
	Payload interface{}
}

func (o *UpdateNetworkSmTargetGroupOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/sm/targetGroups/{targetGroupId}][%d] updateNetworkSmTargetGroupOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkSmTargetGroupOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkSmTargetGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
