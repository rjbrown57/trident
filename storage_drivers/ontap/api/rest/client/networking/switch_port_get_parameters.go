// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewSwitchPortGetParams creates a new SwitchPortGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSwitchPortGetParams() *SwitchPortGetParams {
	return &SwitchPortGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSwitchPortGetParamsWithTimeout creates a new SwitchPortGetParams object
// with the ability to set a timeout on a request.
func NewSwitchPortGetParamsWithTimeout(timeout time.Duration) *SwitchPortGetParams {
	return &SwitchPortGetParams{
		timeout: timeout,
	}
}

// NewSwitchPortGetParamsWithContext creates a new SwitchPortGetParams object
// with the ability to set a context for a request.
func NewSwitchPortGetParamsWithContext(ctx context.Context) *SwitchPortGetParams {
	return &SwitchPortGetParams{
		Context: ctx,
	}
}

// NewSwitchPortGetParamsWithHTTPClient creates a new SwitchPortGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSwitchPortGetParamsWithHTTPClient(client *http.Client) *SwitchPortGetParams {
	return &SwitchPortGetParams{
		HTTPClient: client,
	}
}

/* SwitchPortGetParams contains all the parameters to send to the API endpoint
   for the switch port get operation.

   Typically these are written to a http.Request.
*/
type SwitchPortGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* IdentityIndex.

	   Interface Index
	*/
	IdentityIndexPathParameter int64

	/* IdentityName.

	   Interface Name
	*/
	IdentityNamePathParameter string

	/* Switch.

	   Switch Name
	*/
	SwitchPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the switch port get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SwitchPortGetParams) WithDefaults() *SwitchPortGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the switch port get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SwitchPortGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the switch port get params
func (o *SwitchPortGetParams) WithTimeout(timeout time.Duration) *SwitchPortGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the switch port get params
func (o *SwitchPortGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the switch port get params
func (o *SwitchPortGetParams) WithContext(ctx context.Context) *SwitchPortGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the switch port get params
func (o *SwitchPortGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the switch port get params
func (o *SwitchPortGetParams) WithHTTPClient(client *http.Client) *SwitchPortGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the switch port get params
func (o *SwitchPortGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the switch port get params
func (o *SwitchPortGetParams) WithFieldsQueryParameter(fields []string) *SwitchPortGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the switch port get params
func (o *SwitchPortGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIdentityIndexPathParameter adds the identityIndex to the switch port get params
func (o *SwitchPortGetParams) WithIdentityIndexPathParameter(identityIndex int64) *SwitchPortGetParams {
	o.SetIdentityIndexPathParameter(identityIndex)
	return o
}

// SetIdentityIndexPathParameter adds the identityIndex to the switch port get params
func (o *SwitchPortGetParams) SetIdentityIndexPathParameter(identityIndex int64) {
	o.IdentityIndexPathParameter = identityIndex
}

// WithIdentityNamePathParameter adds the identityName to the switch port get params
func (o *SwitchPortGetParams) WithIdentityNamePathParameter(identityName string) *SwitchPortGetParams {
	o.SetIdentityNamePathParameter(identityName)
	return o
}

// SetIdentityNamePathParameter adds the identityName to the switch port get params
func (o *SwitchPortGetParams) SetIdentityNamePathParameter(identityName string) {
	o.IdentityNamePathParameter = identityName
}

// WithSwitchPathParameter adds the switchVar to the switch port get params
func (o *SwitchPortGetParams) WithSwitchPathParameter(switchVar string) *SwitchPortGetParams {
	o.SetSwitchPathParameter(switchVar)
	return o
}

// SetSwitchPathParameter adds the switch to the switch port get params
func (o *SwitchPortGetParams) SetSwitchPathParameter(switchVar string) {
	o.SwitchPathParameter = switchVar
}

// WriteToRequest writes these params to a swagger request
func (o *SwitchPortGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	// path param identity.index
	if err := r.SetPathParam("identity.index", swag.FormatInt64(o.IdentityIndexPathParameter)); err != nil {
		return err
	}

	// path param identity.name
	if err := r.SetPathParam("identity.name", o.IdentityNamePathParameter); err != nil {
		return err
	}

	// path param switch
	if err := r.SetPathParam("switch", o.SwitchPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSwitchPortGet binds the parameter fields
func (o *SwitchPortGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}
