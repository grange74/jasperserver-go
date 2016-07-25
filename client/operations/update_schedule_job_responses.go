package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateScheduleJobReader is a Reader for the UpdateScheduleJob structure.
type UpdateScheduleJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *UpdateScheduleJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateScheduleJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateScheduleJobOK creates a UpdateScheduleJobOK with default headers values
func NewUpdateScheduleJobOK() *UpdateScheduleJobOK {
	return &UpdateScheduleJobOK{}
}

/*UpdateScheduleJobOK handles this case with default header values.

UpdateScheduleJobOK update schedule job o k
*/
type UpdateScheduleJobOK struct {
}

func (o *UpdateScheduleJobOK) Error() string {
	return fmt.Sprintf("[PUT /jobs][%d] updateScheduleJobOK ", 200)
}

func (o *UpdateScheduleJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
