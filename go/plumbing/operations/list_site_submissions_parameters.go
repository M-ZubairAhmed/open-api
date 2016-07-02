package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListSiteSubmissionsParams creates a new ListSiteSubmissionsParams object
// with the default values initialized.
func NewListSiteSubmissionsParams() *ListSiteSubmissionsParams {
	var ()
	return &ListSiteSubmissionsParams{}
}

/*ListSiteSubmissionsParams contains all the parameters to send to the API endpoint
for the list site submissions operation typically these are written to a http.Request
*/
type ListSiteSubmissionsParams struct {

	/*SiteID*/
	SiteID string
}

// WithSiteID adds the siteId to the list site submissions params
func (o *ListSiteSubmissionsParams) WithSiteID(SiteID string) *ListSiteSubmissionsParams {
	o.SiteID = SiteID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ListSiteSubmissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}