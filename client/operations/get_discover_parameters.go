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

// NewGetDiscoverParams creates a new GetDiscoverParams object
// with the default values initialized.
func NewGetDiscoverParams() *GetDiscoverParams {
	var ()
	return &GetDiscoverParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDiscoverParamsWithTimeout creates a new GetDiscoverParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDiscoverParamsWithTimeout(timeout time.Duration) *GetDiscoverParams {
	var ()
	return &GetDiscoverParams{

		timeout: timeout,
	}
}

/*GetDiscoverParams contains all the parameters to send to the API endpoint
for the get discover operation typically these are written to a http.Request
*/
type GetDiscoverParams struct {

	/*URI*/
	URI *string

	timeout time.Duration
}

// WithURI adds the uri to the get discover params
func (o *GetDiscoverParams) WithURI(URI *string) *GetDiscoverParams {
	o.URI = URI
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetDiscoverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
