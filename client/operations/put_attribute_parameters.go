package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewPutAttributeParams creates a new PutAttributeParams object
// with the default values initialized.
func NewPutAttributeParams() *PutAttributeParams {
	var ()
	return &PutAttributeParams{}
}

/*PutAttributeParams contains all the parameters to send to the API endpoint
for the put attribute operation typically these are written to a http.Request
*/
type PutAttributeParams struct {

	/*Accept*/
	Accept *string
	/*ContentType*/
	ContentType *string
	/*Embedded*/
	Embedded *string
	/*AttrName*/
	AttrName *string
}

// WithAccept adds the accept to the put attribute params
func (o *PutAttributeParams) WithAccept(accept *string) *PutAttributeParams {
	o.Accept = accept
	return o
}

// WithContentType adds the contentType to the put attribute params
func (o *PutAttributeParams) WithContentType(contentType *string) *PutAttributeParams {
	o.ContentType = contentType
	return o
}

// WithEmbedded adds the embedded to the put attribute params
func (o *PutAttributeParams) WithEmbedded(embedded *string) *PutAttributeParams {
	o.Embedded = embedded
	return o
}

// WithAttrName adds the attrName to the put attribute params
func (o *PutAttributeParams) WithAttrName(attrName *string) *PutAttributeParams {
	o.AttrName = attrName
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PutAttributeParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

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

	if o.ContentType != nil {

		// query param Content-Type
		var qrContentType string
		if o.ContentType != nil {
			qrContentType = *o.ContentType
		}
		qContentType := qrContentType
		if qContentType != "" {
			if err := r.SetQueryParam("Content-Type", qContentType); err != nil {
				return err
			}
		}

	}

	if o.Embedded != nil {

		// query param _embedded
		var qrEmbedded string
		if o.Embedded != nil {
			qrEmbedded = *o.Embedded
		}
		qEmbedded := qrEmbedded
		if qEmbedded != "" {
			if err := r.SetQueryParam("_embedded", qEmbedded); err != nil {
				return err
			}
		}

	}

	if o.AttrName != nil {

		// path param attrName
		if err := r.SetPathParam("attrName", *o.AttrName); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}