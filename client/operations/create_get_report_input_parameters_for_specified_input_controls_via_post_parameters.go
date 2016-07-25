package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCreateGetReportInputParametersForSpecifiedInputControlsViaPostParams creates a new CreateGetReportInputParametersForSpecifiedInputControlsViaPostParams object
// with the default values initialized.
func NewCreateGetReportInputParametersForSpecifiedInputControlsViaPostParams() *CreateGetReportInputParametersForSpecifiedInputControlsViaPostParams {
	var ()
	return &CreateGetReportInputParametersForSpecifiedInputControlsViaPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateGetReportInputParametersForSpecifiedInputControlsViaPostParamsWithTimeout creates a new CreateGetReportInputParametersForSpecifiedInputControlsViaPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateGetReportInputParametersForSpecifiedInputControlsViaPostParamsWithTimeout(timeout time.Duration) *CreateGetReportInputParametersForSpecifiedInputControlsViaPostParams {
	var ()
	return &CreateGetReportInputParametersForSpecifiedInputControlsViaPostParams{

		timeout: timeout,
	}
}

/*CreateGetReportInputParametersForSpecifiedInputControlsViaPostParams contains all the parameters to send to the API endpoint
for the create get report input parameters for specified input controls via post operation typically these are written to a http.Request
*/
type CreateGetReportInputParametersForSpecifiedInputControlsViaPostParams struct {

	/*InputControlIds*/
	InputControlIds *string
	/*ReportUnitURI*/
	ReportUnitURI *string

	timeout time.Duration
}

// WithInputControlIds adds the inputControlIds to the create get report input parameters for specified input controls via post params
func (o *CreateGetReportInputParametersForSpecifiedInputControlsViaPostParams) WithInputControlIds(InputControlIds *string) *CreateGetReportInputParametersForSpecifiedInputControlsViaPostParams {
	o.InputControlIds = InputControlIds
	return o
}

// WithReportUnitURI adds the reportUnitUri to the create get report input parameters for specified input controls via post params
func (o *CreateGetReportInputParametersForSpecifiedInputControlsViaPostParams) WithReportUnitURI(ReportUnitURI *string) *CreateGetReportInputParametersForSpecifiedInputControlsViaPostParams {
	o.ReportUnitURI = ReportUnitURI
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreateGetReportInputParametersForSpecifiedInputControlsViaPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
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
