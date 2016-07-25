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

// NewUpdateCancelReportExecutionParams creates a new UpdateCancelReportExecutionParams object
// with the default values initialized.
func NewUpdateCancelReportExecutionParams() *UpdateCancelReportExecutionParams {
	var ()
	return &UpdateCancelReportExecutionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCancelReportExecutionParamsWithTimeout creates a new UpdateCancelReportExecutionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCancelReportExecutionParamsWithTimeout(timeout time.Duration) *UpdateCancelReportExecutionParams {
	var ()
	return &UpdateCancelReportExecutionParams{

		timeout: timeout,
	}
}

/*UpdateCancelReportExecutionParams contains all the parameters to send to the API endpoint
for the update cancel report execution operation typically these are written to a http.Request
*/
type UpdateCancelReportExecutionParams struct {

	/*ExecutionID*/
	ExecutionID *string

	timeout time.Duration
}

// WithExecutionID adds the executionId to the update cancel report execution params
func (o *UpdateCancelReportExecutionParams) WithExecutionID(ExecutionID *string) *UpdateCancelReportExecutionParams {
	o.ExecutionID = ExecutionID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCancelReportExecutionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.ExecutionID != nil {

		// path param executionId
		if err := r.SetPathParam("executionId", *o.ExecutionID); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
