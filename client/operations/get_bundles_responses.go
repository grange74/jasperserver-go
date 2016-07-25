package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// GetBundlesReader is a Reader for the GetBundles structure.
type GetBundlesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetBundlesReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBundlesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBundlesOK creates a GetBundlesOK with default headers values
func NewGetBundlesOK() *GetBundlesOK {
	return &GetBundlesOK{}
}

/*GetBundlesOK handles this case with default header values.

GetBundlesOK get bundles o k
*/
type GetBundlesOK struct {
}

func (o *GetBundlesOK) Error() string {
	return fmt.Sprintf("[GET /bundles][%d] getBundlesOK ", 200)
}

func (o *GetBundlesOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
