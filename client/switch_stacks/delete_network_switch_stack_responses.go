// Code generated by go-swagger; DO NOT EDIT.

package switch_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteNetworkSwitchStackReader is a Reader for the DeleteNetworkSwitchStack structure.
type DeleteNetworkSwitchStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkSwitchStackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworkSwitchStackNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteNetworkSwitchStackNoContent creates a DeleteNetworkSwitchStackNoContent with default headers values
func NewDeleteNetworkSwitchStackNoContent() *DeleteNetworkSwitchStackNoContent {
	return &DeleteNetworkSwitchStackNoContent{}
}

/*DeleteNetworkSwitchStackNoContent handles this case with default header values.

Successful operation
*/
type DeleteNetworkSwitchStackNoContent struct {
}

func (o *DeleteNetworkSwitchStackNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/switchStacks/{switchStackId}][%d] deleteNetworkSwitchStackNoContent ", 204)
}

func (o *DeleteNetworkSwitchStackNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
