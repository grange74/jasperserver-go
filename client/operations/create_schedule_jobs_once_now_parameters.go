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

// NewCreateScheduleJobsOnceNowParams creates a new CreateScheduleJobsOnceNowParams object
// with the default values initialized.
func NewCreateScheduleJobsOnceNowParams() *CreateScheduleJobsOnceNowParams {

	return &CreateScheduleJobsOnceNowParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateScheduleJobsOnceNowParamsWithTimeout creates a new CreateScheduleJobsOnceNowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateScheduleJobsOnceNowParamsWithTimeout(timeout time.Duration) *CreateScheduleJobsOnceNowParams {

	return &CreateScheduleJobsOnceNowParams{

		timeout: timeout,
	}
}

/*CreateScheduleJobsOnceNowParams contains all the parameters to send to the API endpoint
for the create schedule jobs once now operation typically these are written to a http.Request
*/
type CreateScheduleJobsOnceNowParams struct {
	timeout time.Duration
}

// WriteToRequest writes these params to a swagger request
func (o *CreateScheduleJobsOnceNowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
