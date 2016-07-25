package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetCalendarByNameParams creates a new GetCalendarByNameParams object
// with the default values initialized.
func NewGetCalendarByNameParams() *GetCalendarByNameParams {
	var ()
	return &GetCalendarByNameParams{}
}

/*GetCalendarByNameParams contains all the parameters to send to the API endpoint
for the get calendar by name operation typically these are written to a http.Request
*/
type GetCalendarByNameParams struct {

	/*CalendarName*/
	CalendarName *string
}

// WithCalendarName adds the calendarName to the get calendar by name params
func (o *GetCalendarByNameParams) WithCalendarName(calendarName *string) *GetCalendarByNameParams {
	o.CalendarName = calendarName
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetCalendarByNameParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.CalendarName != nil {

		// path param calendarName
		if err := r.SetPathParam("calendarName", *o.CalendarName); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
