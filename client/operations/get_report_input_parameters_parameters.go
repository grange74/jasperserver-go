package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetReportInputParametersParams creates a new GetReportInputParametersParams object
// with the default values initialized.
func NewGetReportInputParametersParams() *GetReportInputParametersParams {
	var ()
	return &GetReportInputParametersParams{}
}

/*GetReportInputParametersParams contains all the parameters to send to the API endpoint
for the get report input parameters operation typically these are written to a http.Request
*/
type GetReportInputParametersParams struct {

	/*ReportUnitURI*/
	ReportUnitURI *string
}

// WithReportUnitURI adds the reportUnitUri to the get report input parameters params
func (o *GetReportInputParametersParams) WithReportUnitURI(reportUnitUri *string) *GetReportInputParametersParams {
	o.ReportUnitURI = reportUnitUri
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetReportInputParametersParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

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
