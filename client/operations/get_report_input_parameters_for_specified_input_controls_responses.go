package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// GetReportInputParametersForSpecifiedInputControlsReader is a Reader for the GetReportInputParametersForSpecifiedInputControls structure.
type GetReportInputParametersForSpecifiedInputControlsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetReportInputParametersForSpecifiedInputControlsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReportInputParametersForSpecifiedInputControlsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReportInputParametersForSpecifiedInputControlsOK creates a GetReportInputParametersForSpecifiedInputControlsOK with default headers values
func NewGetReportInputParametersForSpecifiedInputControlsOK() *GetReportInputParametersForSpecifiedInputControlsOK {
	return &GetReportInputParametersForSpecifiedInputControlsOK{}
}

/*GetReportInputParametersForSpecifiedInputControlsOK handles this case with default header values.

GetReportInputParametersForSpecifiedInputControlsOK get report input parameters for specified input controls o k
*/
type GetReportInputParametersForSpecifiedInputControlsOK struct {
}

func (o *GetReportInputParametersForSpecifiedInputControlsOK) Error() string {
	return fmt.Sprintf("[GET /{inputControlIds: [^;/]+(;[^;/]+)*}][%d] getReportInputParametersForSpecifiedInputControlsOK ", 200)
}

func (o *GetReportInputParametersForSpecifiedInputControlsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
