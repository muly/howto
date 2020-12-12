// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetStudentsHandlerFunc turns a function with the right signature into a get students handler
type GetStudentsHandlerFunc func(GetStudentsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetStudentsHandlerFunc) Handle(params GetStudentsParams) middleware.Responder {
	return fn(params)
}

// GetStudentsHandler interface for that can handle valid get students params
type GetStudentsHandler interface {
	Handle(GetStudentsParams) middleware.Responder
}

// NewGetStudents creates a new http.Handler for the get students operation
func NewGetStudents(ctx *middleware.Context, handler GetStudentsHandler) *GetStudents {
	return &GetStudents{Context: ctx, Handler: handler}
}

/*GetStudents swagger:route GET /api/students getStudents

GetStudents get students API

*/
type GetStudents struct {
	Context *middleware.Context
	Handler GetStudentsHandler
}

func (o *GetStudents) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetStudentsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetStudentsOKBodyItems0 get students o k body items0
//
// swagger:model GetStudentsOKBodyItems0
type GetStudentsOKBodyItems0 struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this get students o k body items0
func (o *GetStudentsOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetStudentsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStudentsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetStudentsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
