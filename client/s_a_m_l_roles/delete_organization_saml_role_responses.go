// Code generated by go-swagger; DO NOT EDIT.

package s_a_m_l_roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteOrganizationSamlRoleReader is a Reader for the DeleteOrganizationSamlRole structure.
type DeleteOrganizationSamlRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationSamlRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOrganizationSamlRoleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteOrganizationSamlRoleNoContent creates a DeleteOrganizationSamlRoleNoContent with default headers values
func NewDeleteOrganizationSamlRoleNoContent() *DeleteOrganizationSamlRoleNoContent {
	return &DeleteOrganizationSamlRoleNoContent{}
}

/*DeleteOrganizationSamlRoleNoContent handles this case with default header values.

Successful operation
*/
type DeleteOrganizationSamlRoleNoContent struct {
}

func (o *DeleteOrganizationSamlRoleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organizationId}/samlRoles/{samlRoleId}][%d] deleteOrganizationSamlRoleNoContent ", 204)
}

func (o *DeleteOrganizationSamlRoleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
