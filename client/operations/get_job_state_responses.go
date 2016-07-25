package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command


import (
  "io"
  "net/http"

  "github.com/go-swagger/go-swagger/httpkit"
  "github.com/go-swagger/go-swagger/swag"
  "github.com/go-swagger/go-swagger/errors"
  "github.com/go-swagger/go-swagger/httpkit/validate"
  "github.com/go-swagger/go-swagger/client"

  strfmt "github.com/go-swagger/go-swagger/strfmt"

  "github.com/retrievercommunications/jasperserver-go-client/models"
  
  
)

// GetJobStateReader is a Reader for the GetJobState structure.
type GetJobStateReader struct {
  formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetJobStateReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
  switch response.Code() {
  
    case 200:
      result := NewGetJobStateOK()
      if err := result.readResponse(response, consumer, o.formats); err != nil {
        return nil, err
      }
      return result, nil
  
    default:
      return nil, client.NewAPIError("unknown error", response, response.Code())
  }
}


// NewGetJobStateOK creates a GetJobStateOK with default headers values
func NewGetJobStateOK() *GetJobStateOK {
  return &GetJobStateOK{
    }
}

/*GetJobStateOK handles this case with default header values.

GetJobStateOK get job state o k
*/
type GetJobStateOK struct {
  
  
}


func (o *GetJobStateOK) Error() string {
	return fmt.Sprintf("[GET /{id: \d+}/state][%d] getJobStateOK ", 200)
}


func (o *GetJobStateOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {
  
  
  return nil
}




