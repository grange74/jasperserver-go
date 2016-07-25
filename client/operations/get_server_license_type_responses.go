package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetServerLicenseTypeReader is a Reader for the GetServerLicenseType structure.
type GetServerLicenseTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetServerLicenseTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetServerLicenseTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServerLicenseTypeOK creates a GetServerLicenseTypeOK with default headers values
func NewGetServerLicenseTypeOK() *GetServerLicenseTypeOK {
	return &GetServerLicenseTypeOK{}
}

/*GetServerLicenseTypeOK handles this case with default header values.

GetServerLicenseTypeOK get server license type o k
*/
type GetServerLicenseTypeOK struct {
}

func (o *GetServerLicenseTypeOK) Error() string {
	return fmt.Sprintf("[GET /licenseType][%d] getServerLicenseTypeOK ", 200)
}

func (o *GetServerLicenseTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
