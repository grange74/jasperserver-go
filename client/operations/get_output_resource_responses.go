package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// GetOutputResourceReader is a Reader for the GetOutputResource structure.
type GetOutputResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetOutputResourceReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOutputResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOutputResourceOK creates a GetOutputResourceOK with default headers values
func NewGetOutputResourceOK() *GetOutputResourceOK {
	return &GetOutputResourceOK{}
}

/*GetOutputResourceOK handles this case with default header values.

GetOutputResourceOK get output resource o k
*/
type GetOutputResourceOK struct {
}

func (o *GetOutputResourceOK) Error() string {
	return fmt.Sprintf("[GET /{executionId}/exports/{exportId}/outputResource][%d] getOutputResourceOK ", 200)
}

func (o *GetOutputResourceOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
