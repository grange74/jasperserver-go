package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetReportExecutionReader is a Reader for the GetReportExecution structure.
type GetReportExecutionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetReportExecutionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReportExecutionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReportExecutionOK creates a GetReportExecutionOK with default headers values
func NewGetReportExecutionOK() *GetReportExecutionOK {
	return &GetReportExecutionOK{}
}

/*GetReportExecutionOK handles this case with default header values.

GetReportExecutionOK get report execution o k
*/
type GetReportExecutionOK struct {
}

func (o *GetReportExecutionOK) Error() string {
	return fmt.Sprintf("[GET /{executionId}][%d] getReportExecutionOK ", 200)
}

func (o *GetReportExecutionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
