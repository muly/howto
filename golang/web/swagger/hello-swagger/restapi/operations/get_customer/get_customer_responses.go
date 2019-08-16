// Code generated by go-swagger; DO NOT EDIT.

package get_customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetCustomerOKCode is the HTTP code returned for type GetCustomerOK
const GetCustomerOKCode int = 200

/*GetCustomerOK Customer found

swagger:response getCustomerOK
*/
type GetCustomerOK struct {
}

// NewGetCustomerOK creates GetCustomerOK with default headers values
func NewGetCustomerOK() *GetCustomerOK {

	return &GetCustomerOK{}
}

// WriteResponse to the client
func (o *GetCustomerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// GetCustomerNotFoundCode is the HTTP code returned for type GetCustomerNotFound
const GetCustomerNotFoundCode int = 404

/*GetCustomerNotFound Customer not found

swagger:response getCustomerNotFound
*/
type GetCustomerNotFound struct {
}

// NewGetCustomerNotFound creates GetCustomerNotFound with default headers values
func NewGetCustomerNotFound() *GetCustomerNotFound {

	return &GetCustomerNotFound{}
}

// WriteResponse to the client
func (o *GetCustomerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
