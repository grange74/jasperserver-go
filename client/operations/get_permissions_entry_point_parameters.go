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

// NewGetPermissionsEntryPointParams creates a new GetPermissionsEntryPointParams object
// with the default values initialized.
func NewGetPermissionsEntryPointParams() *GetPermissionsEntryPointParams {
	var ()
	return &GetPermissionsEntryPointParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPermissionsEntryPointParamsWithTimeout creates a new GetPermissionsEntryPointParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPermissionsEntryPointParamsWithTimeout(timeout time.Duration) *GetPermissionsEntryPointParams {
	var ()
	return &GetPermissionsEntryPointParams{

		timeout: timeout,
	}
}

/*GetPermissionsEntryPointParams contains all the parameters to send to the API endpoint
for the get permissions entry point operation typically these are written to a http.Request
*/
type GetPermissionsEntryPointParams struct {

	/*URI*/
	URI *string

	timeout time.Duration
}

// WithURI adds the uri to the get permissions entry point params
func (o *GetPermissionsEntryPointParams) WithURI(URI *string) *GetPermissionsEntryPointParams {
	o.URI = URI
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetPermissionsEntryPointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.URI != nil {

		// query param uri
		var qrURI string
		if o.URI != nil {
			qrURI = *o.URI
		}
		qURI := qrURI
		if qURI != "" {
			if err := r.SetQueryParam("uri", qURI); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
