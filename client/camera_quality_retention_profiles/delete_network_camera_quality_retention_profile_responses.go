// Code generated by go-swagger; DO NOT EDIT.

package camera_quality_retention_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteNetworkCameraQualityRetentionProfileReader is a Reader for the DeleteNetworkCameraQualityRetentionProfile structure.
type DeleteNetworkCameraQualityRetentionProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkCameraQualityRetentionProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworkCameraQualityRetentionProfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteNetworkCameraQualityRetentionProfileNoContent creates a DeleteNetworkCameraQualityRetentionProfileNoContent with default headers values
func NewDeleteNetworkCameraQualityRetentionProfileNoContent() *DeleteNetworkCameraQualityRetentionProfileNoContent {
	return &DeleteNetworkCameraQualityRetentionProfileNoContent{}
}

/*DeleteNetworkCameraQualityRetentionProfileNoContent handles this case with default header values.

Successful operation
*/
type DeleteNetworkCameraQualityRetentionProfileNoContent struct {
}

func (o *DeleteNetworkCameraQualityRetentionProfileNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId}][%d] deleteNetworkCameraQualityRetentionProfileNoContent ", 204)
}

func (o *DeleteNetworkCameraQualityRetentionProfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
