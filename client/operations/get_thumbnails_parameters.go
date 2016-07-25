package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetThumbnailsParams creates a new GetThumbnailsParams object
// with the default values initialized.
func NewGetThumbnailsParams() *GetThumbnailsParams {
	var ()
	return &GetThumbnailsParams{}
}

/*GetThumbnailsParams contains all the parameters to send to the API endpoint
for the get thumbnails operation typically these are written to a http.Request
*/
type GetThumbnailsParams struct {

	/*Accept*/
	Accept *string
	/*DefaultAllowed*/
	DefaultAllowed *bool
	/*URI*/
	URI *string
}

// WithAccept adds the accept to the get thumbnails params
func (o *GetThumbnailsParams) WithAccept(accept *string) *GetThumbnailsParams {
	o.Accept = accept
	return o
}

// WithDefaultAllowed adds the defaultAllowed to the get thumbnails params
func (o *GetThumbnailsParams) WithDefaultAllowed(defaultAllowed *bool) *GetThumbnailsParams {
	o.DefaultAllowed = defaultAllowed
	return o
}

// WithURI adds the uri to the get thumbnails params
func (o *GetThumbnailsParams) WithURI(uri *string) *GetThumbnailsParams {
	o.URI = uri
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetThumbnailsParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Accept != nil {

		// query param Accept
		var qrAccept string
		if o.Accept != nil {
			qrAccept = *o.Accept
		}
		qAccept := qrAccept
		if qAccept != "" {
			if err := r.SetQueryParam("Accept", qAccept); err != nil {
				return err
			}
		}

	}

	if o.DefaultAllowed != nil {

		// query param defaultAllowed
		var qrDefaultAllowed bool
		if o.DefaultAllowed != nil {
			qrDefaultAllowed = *o.DefaultAllowed
		}
		qDefaultAllowed := swag.FormatBool(qrDefaultAllowed)
		if qDefaultAllowed != "" {
			if err := r.SetQueryParam("defaultAllowed", qDefaultAllowed); err != nil {
				return err
			}
		}

	}

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