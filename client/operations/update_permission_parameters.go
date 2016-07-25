package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewUpdatePermissionParams creates a new UpdatePermissionParams object
// with the default values initialized.
func NewUpdatePermissionParams() *UpdatePermissionParams {
	var ()
	return &UpdatePermissionParams{}
}

/*UpdatePermissionParams contains all the parameters to send to the API endpoint
for the update permission operation typically these are written to a http.Request
*/
type UpdatePermissionParams struct {

	/*URI*/
	URI *string
}

// WithURI adds the uri to the update permission params
func (o *UpdatePermissionParams) WithURI(uri *string) *UpdatePermissionParams {
	o.URI = uri
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePermissionParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.URI != nil {

		// query param uri
		var qrURI string
		if o.URI != nil {
			qrURI = *o.URI
		}
		qURI := qrURI
		if qURI != "" {
			if err := r.SetQueryParam("uri", qURI); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}