package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetPermissionsEntryPointRootReader is a Reader for the GetPermissionsEntryPointRoot structure.
type GetPermissionsEntryPointRootReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPermissionsEntryPointRootReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPermissionsEntryPointRootOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPermissionsEntryPointRootOK creates a GetPermissionsEntryPointRootOK with default headers values
func NewGetPermissionsEntryPointRootOK() *GetPermissionsEntryPointRootOK {
	return &GetPermissionsEntryPointRootOK{}
}

/*GetPermissionsEntryPointRootOK handles this case with default header values.

GetPermissionsEntryPointRootOK get permissions entry point root o k
*/
type GetPermissionsEntryPointRootOK struct {
}

func (o *GetPermissionsEntryPointRootOK) Error() string {
	return fmt.Sprintf("[GET /permissions][%d] getPermissionsEntryPointRootOK ", 200)
}

func (o *GetPermissionsEntryPointRootOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
