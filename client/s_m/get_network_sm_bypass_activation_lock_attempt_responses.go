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

// GetNetworkSmBypassActivationLockAttemptReader is a Reader for the GetNetworkSmBypassActivationLockAttempt structure.
type GetNetworkSmBypassActivationLockAttemptReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSmBypassActivationLockAttemptReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSmBypassActivationLockAttemptOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkSmBypassActivationLockAttemptOK creates a GetNetworkSmBypassActivationLockAttemptOK with default headers values
func NewGetNetworkSmBypassActivationLockAttemptOK() *GetNetworkSmBypassActivationLockAttemptOK {
	return &GetNetworkSmBypassActivationLockAttemptOK{}
}

/*GetNetworkSmBypassActivationLockAttemptOK handles this case with default header values.

Successful operation
*/
type GetNetworkSmBypassActivationLockAttemptOK struct {
	Payload interface{}
}

func (o *GetNetworkSmBypassActivationLockAttemptOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sm/bypassActivationLockAttempts/{attemptId}][%d] getNetworkSmBypassActivationLockAttemptOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSmBypassActivationLockAttemptOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkSmBypassActivationLockAttemptOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
