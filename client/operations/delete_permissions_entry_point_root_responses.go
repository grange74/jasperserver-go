package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeletePermissionsEntryPointRootReader is a Reader for the DeletePermissionsEntryPointRoot structure.
type DeletePermissionsEntryPointRootReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeletePermissionsEntryPointRootReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeletePermissionsEntryPointRootOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePermissionsEntryPointRootOK creates a DeletePermissionsEntryPointRootOK with default headers values
func NewDeletePermissionsEntryPointRootOK() *DeletePermissionsEntryPointRootOK {
	return &DeletePermissionsEntryPointRootOK{}
}

/*DeletePermissionsEntryPointRootOK handles this case with default header values.

DeletePermissionsEntryPointRootOK delete permissions entry point root o k
*/
type DeletePermissionsEntryPointRootOK struct {
}

func (o *DeletePermissionsEntryPointRootOK) Error() string {
	return fmt.Sprintf("[DELETE /permissions][%d] deletePermissionsEntryPointRootOK ", 200)
}

func (o *DeletePermissionsEntryPointRootOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
