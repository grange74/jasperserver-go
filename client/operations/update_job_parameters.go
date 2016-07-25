package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewUpdateJobParams creates a new UpdateJobParams object
// with the default values initialized.
func NewUpdateJobParams() *UpdateJobParams {
	var ()
	return &UpdateJobParams{}
}

/*UpdateJobParams contains all the parameters to send to the API endpoint
for the update job operation typically these are written to a http.Request
*/
type UpdateJobParams struct {

	/*ID*/
	ID *int64
}

// WithID adds the id to the update job params
func (o *UpdateJobParams) WithID(id *int64) *UpdateJobParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateJobParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.ID != nil {

		// query param id
		var qrID int64
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := swag.FormatInt64(qrID)
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}