package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewCreateGetInputControlsInitialValuesViaPostParams creates a new CreateGetInputControlsInitialValuesViaPostParams object
// with the default values initialized.
func NewCreateGetInputControlsInitialValuesViaPostParams() *CreateGetInputControlsInitialValuesViaPostParams {
	var (
		freshDataDefault bool = bool(false)
	)
	return &CreateGetInputControlsInitialValuesViaPostParams{
		FreshData: &freshDataDefault,
	}
}

/*CreateGetInputControlsInitialValuesViaPostParams contains all the parameters to send to the API endpoint
for the create get input controls initial values via post operation typically these are written to a http.Request
*/
type CreateGetInputControlsInitialValuesViaPostParams struct {

	/*FreshData*/
	FreshData *bool
	/*ReportUnitURI*/
	ReportUnitURI *string
}

// WithFreshData adds the freshData to the create get input controls initial values via post params
func (o *CreateGetInputControlsInitialValuesViaPostParams) WithFreshData(freshData *bool) *CreateGetInputControlsInitialValuesViaPostParams {
	o.FreshData = freshData
	return o
}

// WithReportUnitURI adds the reportUnitUri to the create get input controls initial values via post params
func (o *CreateGetInputControlsInitialValuesViaPostParams) WithReportUnitURI(reportUnitUri *string) *CreateGetInputControlsInitialValuesViaPostParams {
	o.ReportUnitURI = reportUnitUri
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreateGetInputControlsInitialValuesViaPostParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.FreshData != nil {

		// query param freshData
		var qrFreshData bool
		if o.FreshData != nil {
			qrFreshData = *o.FreshData
		}
		qFreshData := swag.FormatBool(qrFreshData)
		if qFreshData != "" {
			if err := r.SetQueryParam("freshData", qFreshData); err != nil {
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