package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetReportInputParametersForSpecifiedInputControlsParams creates a new GetReportInputParametersForSpecifiedInputControlsParams object
// with the default values initialized.
func NewGetReportInputParametersForSpecifiedInputControlsParams() *GetReportInputParametersForSpecifiedInputControlsParams {
	var ()
	return &GetReportInputParametersForSpecifiedInputControlsParams{}
}

/*GetReportInputParametersForSpecifiedInputControlsParams contains all the parameters to send to the API endpoint
for the get report input parameters for specified input controls operation typically these are written to a http.Request
*/
type GetReportInputParametersForSpecifiedInputControlsParams struct {

	/*InputControlIds*/
	InputControlIds *string
	/*ReportUnitURI*/
	ReportUnitURI *string
}

// WithInputControlIds adds the inputControlIds to the get report input parameters for specified input controls params
func (o *GetReportInputParametersForSpecifiedInputControlsParams) WithInputControlIds(inputControlIds *string) *GetReportInputParametersForSpecifiedInputControlsParams {
	o.InputControlIds = inputControlIds
	return o
}

// WithReportUnitURI adds the reportUnitUri to the get report input parameters for specified input controls params
func (o *GetReportInputParametersForSpecifiedInputControlsParams) WithReportUnitURI(reportUnitUri *string) *GetReportInputParametersForSpecifiedInputControlsParams {
	o.ReportUnitURI = reportUnitUri
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetReportInputParametersForSpecifiedInputControlsParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.InputControlIds != nil {

		// query param inputControlIds
		var qrInputControlIds string
		if o.InputControlIds != nil {
			qrInputControlIds = *o.InputControlIds
		}
		qInputControlIds := qrInputControlIds
		if qInputControlIds != "" {
			if err := r.SetQueryParam("inputControlIds", qInputControlIds); err != nil {
				return err
			}
		}

	}

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
