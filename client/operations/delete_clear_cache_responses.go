package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteClearCacheReader is a Reader for the DeleteClearCache structure.
type DeleteClearCacheReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteClearCacheReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteClearCacheOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteClearCacheOK creates a DeleteClearCacheOK with default headers values
func NewDeleteClearCacheOK() *DeleteClearCacheOK {
	return &DeleteClearCacheOK{}
}

/*DeleteClearCacheOK handles this case with default header values.

DeleteClearCacheOK delete clear cache o k
*/
type DeleteClearCacheOK struct {
}

func (o *DeleteClearCacheOK) Error() string {
	return fmt.Sprintf("[DELETE /{cacheId}][%d] deleteClearCacheOK ", 200)
}

func (o *DeleteClearCacheOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
