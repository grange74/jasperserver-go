package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CreateGetReportInputParametersViaPostReader is a Reader for the CreateGetReportInputParametersViaPost structure.
type CreateGetReportInputParametersViaPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *CreateGetReportInputParametersViaPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateGetReportInputParametersViaPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateGetReportInputParametersViaPostOK creates a CreateGetReportInputParametersViaPostOK with default headers values
func NewCreateGetReportInputParametersViaPostOK() *CreateGetReportInputParametersViaPostOK {
	return &CreateGetReportInputParametersViaPostOK{}
}

/*CreateGetReportInputParametersViaPostOK handles this case with default header values.

CreateGetReportInputParametersViaPostOK create get report input parameters via post o k
*/
type CreateGetReportInputParametersViaPostOK struct {
}

func (o *CreateGetReportInputParametersViaPostOK) Error() string {
	return fmt.Sprintf("[POST /{executionId}/parameters][%d] createGetReportInputParametersViaPostOK ", 200)
}

func (o *CreateGetReportInputParametersViaPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
