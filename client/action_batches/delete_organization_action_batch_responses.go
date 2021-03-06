// Code generated by go-swagger; DO NOT EDIT.

package action_batches

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteOrganizationActionBatchReader is a Reader for the DeleteOrganizationActionBatch structure.
type DeleteOrganizationActionBatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationActionBatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOrganizationActionBatchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteOrganizationActionBatchNoContent creates a DeleteOrganizationActionBatchNoContent with default headers values
func NewDeleteOrganizationActionBatchNoContent() *DeleteOrganizationActionBatchNoContent {
	return &DeleteOrganizationActionBatchNoContent{}
}

/*DeleteOrganizationActionBatchNoContent handles this case with default header values.

Successful operation
*/
type DeleteOrganizationActionBatchNoContent struct {
}

func (o *DeleteOrganizationActionBatchNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organizationId}/actionBatches/{actionBatchId}][%d] deleteOrganizationActionBatchNoContent ", 204)
}

func (o *DeleteOrganizationActionBatchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
