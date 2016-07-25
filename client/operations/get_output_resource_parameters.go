package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetOutputResourceParams creates a new GetOutputResourceParams object
// with the default values initialized.
func NewGetOutputResourceParams() *GetOutputResourceParams {
	var (
		suppressContentDispositionDefault bool = bool(false)
	)
	return &GetOutputResourceParams{
		SuppressContentDisposition: &suppressContentDispositionDefault,
	}
}

/*GetOutputResourceParams contains all the parameters to send to the API endpoint
for the get output resource operation typically these are written to a http.Request
*/
type GetOutputResourceParams struct {

	/*ExecutionID*/
	ExecutionID *string
	/*ExportID*/
	ExportID *string
	/*SuppressContentDisposition*/
	SuppressContentDisposition *bool
}

// WithExecutionID adds the executionId to the get output resource params
func (o *GetOutputResourceParams) WithExecutionID(executionId *string) *GetOutputResourceParams {
	o.ExecutionID = executionId
	return o
}

// WithExportID adds the exportId to the get output resource params
func (o *GetOutputResourceParams) WithExportID(exportId *string) *GetOutputResourceParams {
	o.ExportID = exportId
	return o
}

// WithSuppressContentDisposition adds the suppressContentDisposition to the get output resource params
func (o *GetOutputResourceParams) WithSuppressContentDisposition(suppressContentDisposition *bool) *GetOutputResourceParams {
	o.SuppressContentDisposition = suppressContentDisposition
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutputResourceParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.ExecutionID != nil {

		// path param executionId
		if err := r.SetPathParam("executionId", *o.ExecutionID); err != nil {
			return err
		}

	}

	if o.ExportID != nil {

		// path param exportId
		if err := r.SetPathParam("exportId", *o.ExportID); err != nil {
			return err
		}

	}

	if o.SuppressContentDisposition != nil {

		// query param suppressContentDisposition
		var qrSuppressContentDisposition bool
		if o.SuppressContentDisposition != nil {
			qrSuppressContentDisposition = *o.SuppressContentDisposition
		}
		qSuppressContentDisposition := swag.FormatBool(qrSuppressContentDisposition)
		if qSuppressContentDisposition != "" {
			if err := r.SetQueryParam("suppressContentDisposition", qSuppressContentDisposition); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
