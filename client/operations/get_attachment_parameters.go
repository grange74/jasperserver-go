package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetAttachmentParams creates a new GetAttachmentParams object
// with the default values initialized.
func NewGetAttachmentParams() *GetAttachmentParams {
	var ()
	return &GetAttachmentParams{}
}

/*GetAttachmentParams contains all the parameters to send to the API endpoint
for the get attachment operation typically these are written to a http.Request
*/
type GetAttachmentParams struct {

	/*Attachment*/
	Attachment *string
	/*ExecutionID*/
	ExecutionID *string
	/*ExportID*/
	ExportID *string
}

// WithAttachment adds the attachment to the get attachment params
func (o *GetAttachmentParams) WithAttachment(attachment *string) *GetAttachmentParams {
	o.Attachment = attachment
	return o
}

// WithExecutionID adds the executionId to the get attachment params
func (o *GetAttachmentParams) WithExecutionID(executionId *string) *GetAttachmentParams {
	o.ExecutionID = executionId
	return o
}

// WithExportID adds the exportId to the get attachment params
func (o *GetAttachmentParams) WithExportID(exportId *string) *GetAttachmentParams {
	o.ExportID = exportId
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetAttachmentParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Attachment != nil {

		// path param attachment
		if err := r.SetPathParam("attachment", *o.Attachment); err != nil {
			return err
		}

	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}