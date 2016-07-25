package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// CreateGetThumbnailsFormEncodedReader is a Reader for the CreateGetThumbnailsFormEncoded structure.
type CreateGetThumbnailsFormEncodedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *CreateGetThumbnailsFormEncodedReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateGetThumbnailsFormEncodedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateGetThumbnailsFormEncodedOK creates a CreateGetThumbnailsFormEncodedOK with default headers values
func NewCreateGetThumbnailsFormEncodedOK() *CreateGetThumbnailsFormEncodedOK {
	return &CreateGetThumbnailsFormEncodedOK{}
}

/*CreateGetThumbnailsFormEncodedOK handles this case with default header values.

CreateGetThumbnailsFormEncodedOK create get thumbnails form encoded o k
*/
type CreateGetThumbnailsFormEncodedOK struct {
}

func (o *CreateGetThumbnailsFormEncodedOK) Error() string {
	return fmt.Sprintf("[POST /thumbnails][%d] createGetThumbnailsFormEncodedOK ", 200)
}

func (o *CreateGetThumbnailsFormEncodedOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
