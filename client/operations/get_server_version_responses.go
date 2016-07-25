package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// GetServerVersionReader is a Reader for the GetServerVersion structure.
type GetServerVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetServerVersionReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetServerVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServerVersionOK creates a GetServerVersionOK with default headers values
func NewGetServerVersionOK() *GetServerVersionOK {
	return &GetServerVersionOK{}
}

/*GetServerVersionOK handles this case with default header values.

GetServerVersionOK get server version o k
*/
type GetServerVersionOK struct {
}

func (o *GetServerVersionOK) Error() string {
	return fmt.Sprintf("[GET /version][%d] getServerVersionOK ", 200)
}

func (o *GetServerVersionOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
