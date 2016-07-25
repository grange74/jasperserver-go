package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// DeleteResourcesReader is a Reader for the DeleteResources structure.
type DeleteResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteResourcesReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteResourcesOK creates a DeleteResourcesOK with default headers values
func NewDeleteResourcesOK() *DeleteResourcesOK {
	return &DeleteResourcesOK{}
}

/*DeleteResourcesOK handles this case with default header values.

DeleteResourcesOK delete resources o k
*/
type DeleteResourcesOK struct {
}

func (o *DeleteResourcesOK) Error() string {
	return fmt.Sprintf("[DELETE /resources][%d] deleteResourcesOK ", 200)
}

func (o *DeleteResourcesOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}