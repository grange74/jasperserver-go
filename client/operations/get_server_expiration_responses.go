package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// GetServerExpirationReader is a Reader for the GetServerExpiration structure.
type GetServerExpirationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetServerExpirationReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetServerExpirationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServerExpirationOK creates a GetServerExpirationOK with default headers values
func NewGetServerExpirationOK() *GetServerExpirationOK {
	return &GetServerExpirationOK{}
}

/*GetServerExpirationOK handles this case with default header values.

GetServerExpirationOK get server expiration o k
*/
type GetServerExpirationOK struct {
}

func (o *GetServerExpirationOK) Error() string {
	return fmt.Sprintf("[GET /expiration][%d] getServerExpirationOK ", 200)
}

func (o *GetServerExpirationOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}