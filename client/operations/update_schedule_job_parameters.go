package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewUpdateScheduleJobParams creates a new UpdateScheduleJobParams object
// with the default values initialized.
func NewUpdateScheduleJobParams() *UpdateScheduleJobParams {

	return &UpdateScheduleJobParams{}
}

/*UpdateScheduleJobParams contains all the parameters to send to the API endpoint
for the update schedule job operation typically these are written to a http.Request
*/
type UpdateScheduleJobParams struct {
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateScheduleJobParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}