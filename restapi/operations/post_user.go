// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"
)

// PostUserHandlerFunc turns a function with the right signature into a post user handler
type PostUserHandlerFunc func(PostUserParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostUserHandlerFunc) Handle(params PostUserParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostUserHandler interface for that can handle valid post user params
type PostUserHandler interface {
	Handle(PostUserParams, interface{}) middleware.Responder
}

// NewPostUser creates a new http.Handler for the post user operation
func NewPostUser(ctx *middleware.Context, handler PostUserHandler) *PostUser {
	return &PostUser{Context: ctx, Handler: handler}
}

/*PostUser swagger:route POST /user/ postUser

creates a new user

*/
type PostUser struct {
	Context *middleware.Context
	Handler PostUserHandler
}

func (o *PostUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostUserParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostUserOKBody post user o k body
// swagger:model PostUserOKBody

type PostUserOKBody struct {

	// id
	// Required: true
	ID *int64 `json:"id"`
}

/* polymorph PostUserOKBody id false */

// Validate validates this post user o k body
func (o *PostUserOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostUserOKBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("postUserOK"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostUserOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostUserOKBody) UnmarshalBinary(b []byte) error {
	var res PostUserOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}