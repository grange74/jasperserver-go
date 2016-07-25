package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetResourcesParams creates a new GetResourcesParams object
// with the default values initialized.
func NewGetResourcesParams() *GetResourcesParams {
	var (
		forceFullPageDefault   bool = bool(false)
		forceTotalCountDefault bool = bool(false)
		recursiveDefault       bool = bool(true)
		showHiddenItemsDefault bool = bool(false)
	)
	return &GetResourcesParams{
		ForceFullPage:   &forceFullPageDefault,
		ForceTotalCount: &forceTotalCountDefault,
		Recursive:       &recursiveDefault,
		ShowHiddenItems: &showHiddenItemsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetResourcesParamsWithTimeout creates a new GetResourcesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetResourcesParamsWithTimeout(timeout time.Duration) *GetResourcesParams {
	var (
		forceFullPageDefault   bool = bool(false)
		forceTotalCountDefault bool = bool(false)
		recursiveDefault       bool = bool(true)
		showHiddenItemsDefault bool = bool(false)
	)
	return &GetResourcesParams{
		ForceFullPage:   &forceFullPageDefault,
		ForceTotalCount: &forceTotalCountDefault,
		Recursive:       &recursiveDefault,
		ShowHiddenItems: &showHiddenItemsDefault,

		timeout: timeout,
	}
}

/*GetResourcesParams contains all the parameters to send to the API endpoint
for the get resources operation typically these are written to a http.Request
*/
type GetResourcesParams struct {

	/*Accept*/
	Accept *string
	/*AccessType*/
	AccessType *string
	/*ExcludeFolder*/
	ExcludeFolder *string
	/*Expanded*/
	Expanded *bool
	/*FolderURI*/
	FolderURI *string
	/*ForceFullPage*/
	ForceFullPage *bool
	/*ForceTotalCount*/
	ForceTotalCount *bool
	/*Limit*/
	Limit *int32
	/*Offset*/
	Offset *int32
	/*Q*/
	Q *string
	/*Recursive*/
	Recursive *bool
	/*ShowHiddenItems*/
	ShowHiddenItems *bool
	/*SortBy*/
	SortBy *string
	/*Type*/
	Type *string

	timeout time.Duration
}

// WithAccept adds the accept to the get resources params
func (o *GetResourcesParams) WithAccept(Accept *string) *GetResourcesParams {
	o.Accept = Accept
	return o
}

// WithAccessType adds the accessType to the get resources params
func (o *GetResourcesParams) WithAccessType(AccessType *string) *GetResourcesParams {
	o.AccessType = AccessType
	return o
}

// WithExcludeFolder adds the excludeFolder to the get resources params
func (o *GetResourcesParams) WithExcludeFolder(ExcludeFolder *string) *GetResourcesParams {
	o.ExcludeFolder = ExcludeFolder
	return o
}

// WithExpanded adds the expanded to the get resources params
func (o *GetResourcesParams) WithExpanded(Expanded *bool) *GetResourcesParams {
	o.Expanded = Expanded
	return o
}

// WithFolderURI adds the folderUri to the get resources params
func (o *GetResourcesParams) WithFolderURI(FolderURI *string) *GetResourcesParams {
	o.FolderURI = FolderURI
	return o
}

// WithForceFullPage adds the forceFullPage to the get resources params
func (o *GetResourcesParams) WithForceFullPage(ForceFullPage *bool) *GetResourcesParams {
	o.ForceFullPage = ForceFullPage
	return o
}

// WithForceTotalCount adds the forceTotalCount to the get resources params
func (o *GetResourcesParams) WithForceTotalCount(ForceTotalCount *bool) *GetResourcesParams {
	o.ForceTotalCount = ForceTotalCount
	return o
}

// WithLimit adds the limit to the get resources params
func (o *GetResourcesParams) WithLimit(Limit *int32) *GetResourcesParams {
	o.Limit = Limit
	return o
}

// WithOffset adds the offset to the get resources params
func (o *GetResourcesParams) WithOffset(Offset *int32) *GetResourcesParams {
	o.Offset = Offset
	return o
}

// WithQ adds the q to the get resources params
func (o *GetResourcesParams) WithQ(Q *string) *GetResourcesParams {
	o.Q = Q
	return o
}

// WithRecursive adds the recursive to the get resources params
func (o *GetResourcesParams) WithRecursive(Recursive *bool) *GetResourcesParams {
	o.Recursive = Recursive
	return o
}

// WithShowHiddenItems adds the showHiddenItems to the get resources params
func (o *GetResourcesParams) WithShowHiddenItems(ShowHiddenItems *bool) *GetResourcesParams {
	o.ShowHiddenItems = ShowHiddenItems
	return o
}

// WithSortBy adds the sortBy to the get resources params
func (o *GetResourcesParams) WithSortBy(SortBy *string) *GetResourcesParams {
	o.SortBy = SortBy
	return o
}

// WithType adds the type to the get resources params
func (o *GetResourcesParams) WithType(Type *string) *GetResourcesParams {
	o.Type = Type
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.AccessType != nil {

		// query param accessType
		var qrAccessType string
		if o.AccessType != nil {
			qrAccessType = *o.AccessType
		}
		qAccessType := qrAccessType
		if qAccessType != "" {
			if err := r.SetQueryParam("accessType", qAccessType); err != nil {
				return err
			}
		}

	}

	if o.ExcludeFolder != nil {

		// query param excludeFolder
		var qrExcludeFolder string
		if o.ExcludeFolder != nil {
			qrExcludeFolder = *o.ExcludeFolder
		}
		qExcludeFolder := qrExcludeFolder
		if qExcludeFolder != "" {
			if err := r.SetQueryParam("excludeFolder", qExcludeFolder); err != nil {
				return err
			}
		}

	}

	if o.Expanded != nil {

		// query param expanded
		var qrExpanded bool
		if o.Expanded != nil {
			qrExpanded = *o.Expanded
		}
		qExpanded := swag.FormatBool(qrExpanded)
		if qExpanded != "" {
			if err := r.SetQueryParam("expanded", qExpanded); err != nil {
				return err
			}
		}

	}

	if o.FolderURI != nil {

		// query param folderUri
		var qrFolderURI string
		if o.FolderURI != nil {
			qrFolderURI = *o.FolderURI
		}
		qFolderURI := qrFolderURI
		if qFolderURI != "" {
			if err := r.SetQueryParam("folderUri", qFolderURI); err != nil {
				return err
			}
		}

	}

	if o.ForceFullPage != nil {

		// query param forceFullPage
		var qrForceFullPage bool
		if o.ForceFullPage != nil {
			qrForceFullPage = *o.ForceFullPage
		}
		qForceFullPage := swag.FormatBool(qrForceFullPage)
		if qForceFullPage != "" {
			if err := r.SetQueryParam("forceFullPage", qForceFullPage); err != nil {
				return err
			}
		}

	}

	if o.ForceTotalCount != nil {

		// query param forceTotalCount
		var qrForceTotalCount bool
		if o.ForceTotalCount != nil {
			qrForceTotalCount = *o.ForceTotalCount
		}
		qForceTotalCount := swag.FormatBool(qrForceTotalCount)
		if qForceTotalCount != "" {
			if err := r.SetQueryParam("forceTotalCount", qForceTotalCount); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Recursive != nil {

		// query param recursive
		var qrRecursive bool
		if o.Recursive != nil {
			qrRecursive = *o.Recursive
		}
		qRecursive := swag.FormatBool(qrRecursive)
		if qRecursive != "" {
			if err := r.SetQueryParam("recursive", qRecursive); err != nil {
				return err
			}
		}

	}

	if o.ShowHiddenItems != nil {

		// query param showHiddenItems
		var qrShowHiddenItems bool
		if o.ShowHiddenItems != nil {
			qrShowHiddenItems = *o.ShowHiddenItems
		}
		qShowHiddenItems := swag.FormatBool(qrShowHiddenItems)
		if qShowHiddenItems != "" {
			if err := r.SetQueryParam("showHiddenItems", qShowHiddenItems); err != nil {
				return err
			}
		}

	}

	if o.SortBy != nil {

		// query param sortBy
		var qrSortBy string
		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {
			if err := r.SetQueryParam("sortBy", qSortBy); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
