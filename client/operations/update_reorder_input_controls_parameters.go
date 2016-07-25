package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewUpdateReorderInputControlsParams creates a new UpdateReorderInputControlsParams object
// with the default values initialized.
func NewUpdateReorderInputControlsParams() *UpdateReorderInputControlsParams {
	var ()
	return &UpdateReorderInputControlsParams{}
}

/*UpdateReorderInputControlsParams contains all the parameters to send to the API endpoint
for the update reorder input controls operation typically these are written to a http.Request
*/
type UpdateReorderInputControlsParams struct {

	/*ReportUnitURI*/
	ReportUnitURI *string
}

// WithReportUnitURI adds the reportUnitUri to the update reorder input controls params
func (o *UpdateReorderInputControlsParams) WithReportUnitURI(reportUnitUri *string) *UpdateReorderInputControlsParams {
	o.ReportUnitURI = reportUnitUri
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateReorderInputControlsParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.ReportUnitURI != nil {

		// query param reportUnitURI
		var qrReportUnitURI string
		if o.ReportUnitURI != nil {
			qrReportUnitURI = *o.ReportUnitURI
		}
		qReportUnitURI := qrReportUnitURI
		if qReportUnitURI != "" {
			if err := r.SetQueryParam("reportUnitURI", qReportUnitURI); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}