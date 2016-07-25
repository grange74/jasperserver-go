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

// NewGetSpecificAttributeOfUserParams creates a new GetSpecificAttributeOfUserParams object
// with the default values initialized.
func NewGetSpecificAttributeOfUserParams() *GetSpecificAttributeOfUserParams {
	var ()
	return &GetSpecificAttributeOfUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSpecificAttributeOfUserParamsWithTimeout creates a new GetSpecificAttributeOfUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSpecificAttributeOfUserParamsWithTimeout(timeout time.Duration) *GetSpecificAttributeOfUserParams {
	var ()
	return &GetSpecificAttributeOfUserParams{

		timeout: timeout,
	}
}

/*GetSpecificAttributeOfUserParams contains all the parameters to send to the API endpoint
for the get specific attribute of user operation typically these are written to a http.Request
*/
type GetSpecificAttributeOfUserParams struct {

	/*Accept*/
	Accept *string
	/*Embedded*/
	Embedded *string
	/*AttrName*/
	AttrName *string
	/*Name*/
	Name *string

	timeout time.Duration
}

// WithAccept adds the accept to the get specific attribute of user params
func (o *GetSpecificAttributeOfUserParams) WithAccept(Accept *string) *GetSpecificAttributeOfUserParams {
	o.Accept = Accept
	return o
}

// WithEmbedded adds the embedded to the get specific attribute of user params
func (o *GetSpecificAttributeOfUserParams) WithEmbedded(Embedded *string) *GetSpecificAttributeOfUserParams {
	o.Embedded = Embedded
	return o
}

// WithAttrName adds the attrName to the get specific attribute of user params
func (o *GetSpecificAttributeOfUserParams) WithAttrName(AttrName *string) *GetSpecificAttributeOfUserParams {
	o.AttrName = AttrName
	return o
}

// WithName adds the name to the get specific attribute of user params
func (o *GetSpecificAttributeOfUserParams) WithName(Name *string) *GetSpecificAttributeOfUserParams {
	o.Name = Name
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetSpecificAttributeOfUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
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

	if o.Name != nil {

		// path param name
		if err := r.SetPathParam("name", *o.Name); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
