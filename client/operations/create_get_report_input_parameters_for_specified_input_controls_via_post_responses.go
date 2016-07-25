package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// CreateGetReportInputParametersForSpecifiedInputControlsViaPostReader is a Reader for the CreateGetReportInputParametersForSpecifiedInputControlsViaPost structure.
type CreateGetReportInputParametersForSpecifiedInputControlsViaPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *CreateGetReportInputParametersForSpecifiedInputControlsViaPostReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateGetReportInputParametersForSpecifiedInputControlsViaPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateGetReportInputParametersForSpecifiedInputControlsViaPostOK creates a CreateGetReportInputParametersForSpecifiedInputControlsViaPostOK with default headers values
func NewCreateGetReportInputParametersForSpecifiedInputControlsViaPostOK() *CreateGetReportInputParametersForSpecifiedInputControlsViaPostOK {
	return &CreateGetReportInputParametersForSpecifiedInputControlsViaPostOK{}
}

/*CreateGetReportInputParametersForSpecifiedInputControlsViaPostOK handles this case with default header values.

CreateGetReportInputParametersForSpecifiedInputControlsViaPostOK create get report input parameters for specified input controls via post o k
*/
type CreateGetReportInputParametersForSpecifiedInputControlsViaPostOK struct {
}

func (o *CreateGetReportInputParametersForSpecifiedInputControlsViaPostOK) Error() string {
	return fmt.Sprintf("[POST /{inputControlIds: [^;/]+(;[^;/]+)*}][%d] createGetReportInputParametersForSpecifiedInputControlsViaPostOK ", 200)
}

func (o *CreateGetReportInputParametersForSpecifiedInputControlsViaPostOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
