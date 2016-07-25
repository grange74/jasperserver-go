package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// GetCustomDataSourcesReader is a Reader for the GetCustomDataSources structure.
type GetCustomDataSourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetCustomDataSourcesReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCustomDataSourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCustomDataSourcesOK creates a GetCustomDataSourcesOK with default headers values
func NewGetCustomDataSourcesOK() *GetCustomDataSourcesOK {
	return &GetCustomDataSourcesOK{}
}

/*GetCustomDataSourcesOK handles this case with default header values.

GetCustomDataSourcesOK get custom data sources o k
*/
type GetCustomDataSourcesOK struct {
}

func (o *GetCustomDataSourcesOK) Error() string {
	return fmt.Sprintf("[GET /customDataSources][%d] getCustomDataSourcesOK ", 200)
}

func (o *GetCustomDataSourcesOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
