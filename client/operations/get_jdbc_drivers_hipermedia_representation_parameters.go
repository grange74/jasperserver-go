package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetJdbcDriversHipermediaRepresentationParams creates a new GetJdbcDriversHipermediaRepresentationParams object
// with the default values initialized.
func NewGetJdbcDriversHipermediaRepresentationParams() *GetJdbcDriversHipermediaRepresentationParams {

	return &GetJdbcDriversHipermediaRepresentationParams{}
}

/*GetJdbcDriversHipermediaRepresentationParams contains all the parameters to send to the API endpoint
for the get jdbc drivers hipermedia representation operation typically these are written to a http.Request
*/
type GetJdbcDriversHipermediaRepresentationParams struct {
}

// WriteToRequest writes these params to a swagger request
func (o *GetJdbcDriversHipermediaRepresentationParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
