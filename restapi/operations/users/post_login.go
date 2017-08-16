// Code generated by go-swagger; DO NOT EDIT.

package users

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

// PostLoginHandlerFunc turns a function with the right signature into a post login handler
type PostLoginHandlerFunc func(PostLoginParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostLoginHandlerFunc) Handle(params PostLoginParams) middleware.Responder {
	return fn(params)
}

// PostLoginHandler interface for that can handle valid post login params
type PostLoginHandler interface {
	Handle(PostLoginParams) middleware.Responder
}

// NewPostLogin creates a new http.Handler for the post login operation
func NewPostLogin(ctx *middleware.Context, handler PostLoginHandler) *PostLogin {
	return &PostLogin{Context: ctx, Handler: handler}
}

/*PostLogin swagger:route POST /login users postLogin

get an swt token to access protected endpoints

*/
type PostLogin struct {
	Context *middleware.Context
	Handler PostLoginHandler
}

func (o *PostLogin) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostLoginParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostLoginOKBody post login o k body
// swagger:model PostLoginOKBody
type PostLoginOKBody struct {

	// token
	// Required: true
	Token *string `json:"token"`
}

// Validate validates this post login o k body
func (o *PostLoginOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateToken(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLoginOKBody) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("postLoginOK"+"."+"token", "body", o.Token); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLoginOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLoginOKBody) UnmarshalBinary(b []byte) error {
	var res PostLoginOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}