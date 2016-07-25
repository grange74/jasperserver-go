package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CreateConnectionReader is a Reader for the CreateConnection structure.
type CreateConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *CreateConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateConnectionOK creates a CreateConnectionOK with default headers values
func NewCreateConnectionOK() *CreateConnectionOK {
	return &CreateConnectionOK{}
}

/*CreateConnectionOK handles this case with default header values.

CreateConnectionOK create connection o k
*/
type CreateConnectionOK struct {
}

func (o *CreateConnectionOK) Error() string {
	return fmt.Sprintf("[POST /connections][%d] createConnectionOK ", 200)
}

func (o *CreateConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
