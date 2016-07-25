package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewCreateFileViaFormParams creates a new CreateFileViaFormParams object
// with the default values initialized.
func NewCreateFileViaFormParams() *CreateFileViaFormParams {
	var (
		createFoldersDefault bool = bool(true)
	)
	return &CreateFileViaFormParams{
		CreateFolders: &createFoldersDefault,
	}
}

/*CreateFileViaFormParams contains all the parameters to send to the API endpoint
for the create file via form operation typically these are written to a http.Request
*/
type CreateFileViaFormParams struct {

	/*CreateFolders*/
	CreateFolders *bool
	/*URI*/
	URI *string
}

// WithCreateFolders adds the createFolders to the create file via form params
func (o *CreateFileViaFormParams) WithCreateFolders(createFolders *bool) *CreateFileViaFormParams {
	o.CreateFolders = createFolders
	return o
}

// WithURI adds the uri to the create file via form params
func (o *CreateFileViaFormParams) WithURI(uri *string) *CreateFileViaFormParams {
	o.URI = uri
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreateFileViaFormParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.CreateFolders != nil {

		// query param createFolders
		var qrCreateFolders bool
		if o.CreateFolders != nil {
			qrCreateFolders = *o.CreateFolders
		}
		qCreateFolders := swag.FormatBool(qrCreateFolders)
		if qCreateFolders != "" {
			if err := r.SetQueryParam("createFolders", qCreateFolders); err != nil {
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