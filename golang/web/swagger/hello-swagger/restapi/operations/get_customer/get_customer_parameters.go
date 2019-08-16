// Code generated by go-swagger; DO NOT EDIT.

package get_customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCustomerParams creates a new GetCustomerParams object
// no default values defined in spec.
func NewGetCustomerParams() GetCustomerParams {

	return GetCustomerParams{}
}

// GetCustomerParams contains all the bound params for the get customer operation
// typically these are obtained from a http.Request
//
// swagger:parameters getCustomer
type GetCustomerParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*email of the customer
	  Required: true
	  Pattern: [a-z0-9._%+-]+@[a-z0-9.-]+\.com$
	  In: path
	*/
	Email strfmt.Email
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetCustomerParams() beforehand.
func (o *GetCustomerParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rEmail, rhkEmail, _ := route.Params.GetOK("email")
	if err := o.bindEmail(rEmail, rhkEmail, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindEmail binds and validates parameter Email from path.
func (o *GetCustomerParams) bindEmail(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: email
	value, err := formats.Parse("email", raw)
	if err != nil {
		return errors.InvalidType("email", "path", "strfmt.Email", raw)
	}
	o.Email = *(value.(*strfmt.Email))

	if err := o.validateEmail(formats); err != nil {
		return err
	}

	return nil
}

// validateEmail carries on validations for parameter Email
func (o *GetCustomerParams) validateEmail(formats strfmt.Registry) error {

	if err := validate.Pattern("email", "path", o.Email.String(), `[a-z0-9._%+-]+@[a-z0-9.-]+\.com$`); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "path", "email", o.Email.String(), formats); err != nil {
		return err
	}
	return nil
}
