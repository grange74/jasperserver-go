package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateReorderInputControlsReader is a Reader for the UpdateReorderInputControls structure.
type UpdateReorderInputControlsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *UpdateReorderInputControlsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateReorderInputControlsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateReorderInputControlsOK creates a UpdateReorderInputControlsOK with default headers values
func NewUpdateReorderInputControlsOK() *UpdateReorderInputControlsOK {
	return &UpdateReorderInputControlsOK{}
}

/*UpdateReorderInputControlsOK handles this case with default header values.

UpdateReorderInputControlsOK update reorder input controls o k
*/
type UpdateReorderInputControlsOK struct {
}

func (o *UpdateReorderInputControlsOK) Error() string {
	return fmt.Sprintf("[PUT /reports/{reportUnitURI: .+}/inputControls][%d] updateReorderInputControlsOK ", 200)
}

func (o *UpdateReorderInputControlsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
