package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetServerVersionParams creates a new GetServerVersionParams object
// with the default values initialized.
func NewGetServerVersionParams() *GetServerVersionParams {

	return &GetServerVersionParams{}
}

/*GetServerVersionParams contains all the parameters to send to the API endpoint
for the get server version operation typically these are written to a http.Request
*/
type GetServerVersionParams struct {
}

// WriteToRequest writes these params to a swagger request
func (o *GetServerVersionParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
