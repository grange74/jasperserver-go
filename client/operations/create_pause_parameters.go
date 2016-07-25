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

// NewCreatePauseParams creates a new CreatePauseParams object
// with the default values initialized.
func NewCreatePauseParams() *CreatePauseParams {

	return &CreatePauseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePauseParamsWithTimeout creates a new CreatePauseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePauseParamsWithTimeout(timeout time.Duration) *CreatePauseParams {

	return &CreatePauseParams{

		timeout: timeout,
	}
}

/*CreatePauseParams contains all the parameters to send to the API endpoint
for the create pause operation typically these are written to a http.Request
*/
type CreatePauseParams struct {
	timeout time.Duration
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePauseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
