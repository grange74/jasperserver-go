package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// UpdateCancelReportExecutionReader is a Reader for the UpdateCancelReportExecution structure.
type UpdateCancelReportExecutionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *UpdateCancelReportExecutionReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateCancelReportExecutionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateCancelReportExecutionOK creates a UpdateCancelReportExecutionOK with default headers values
func NewUpdateCancelReportExecutionOK() *UpdateCancelReportExecutionOK {
	return &UpdateCancelReportExecutionOK{}
}

/*UpdateCancelReportExecutionOK handles this case with default header values.

UpdateCancelReportExecutionOK update cancel report execution o k
*/
type UpdateCancelReportExecutionOK struct {
}

func (o *UpdateCancelReportExecutionOK) Error() string {
	return fmt.Sprintf("[PUT /{executionId}/status][%d] updateCancelReportExecutionOK ", 200)
}

func (o *UpdateCancelReportExecutionOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}