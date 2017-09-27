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

// GetUserF2aHandlerFunc turns a function with the right signature into a get user f2a handler
type GetUserF2aHandlerFunc func(GetUserF2aParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserF2aHandlerFunc) Handle(params GetUserF2aParams) middleware.Responder {
	return fn(params)
}

// GetUserF2aHandler interface for that can handle valid get user f2a params
type GetUserF2aHandler interface {
	Handle(GetUserF2aParams) middleware.Responder
}

// NewGetUserF2a creates a new http.Handler for the get user f2a operation
func NewGetUserF2a(ctx *middleware.Context, handler GetUserF2aHandler) *GetUserF2a {
	return &GetUserF2a{Context: ctx, Handler: handler}
}

/*GetUserF2a swagger:route GET /user/f2a getUserF2a

generates a qr base64 encoded image and master code for the user to scan with the google authenticator and add it to the phone app

*/
type GetUserF2a struct {
	Context *middleware.Context
	Handler GetUserF2aHandler
}

func (o *GetUserF2a) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUserF2aParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetUserF2aOKBody get user f2a o k body
// swagger:model GetUserF2aOKBody

type GetUserF2aOKBody struct {

	// qr
	// Required: true
	Qr *string `json:"qr"`

	// secret
	// Required: true
	Secret *string `json:"secret"`
}

/* polymorph GetUserF2aOKBody qr false */

/* polymorph GetUserF2aOKBody secret false */

// Validate validates this get user f2a o k body
func (o *GetUserF2aOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateQr(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateSecret(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUserF2aOKBody) validateQr(formats strfmt.Registry) error {

	if err := validate.Required("getUserF2aOK"+"."+"qr", "body", o.Qr); err != nil {
		return err
	}

	return nil
}

func (o *GetUserF2aOKBody) validateSecret(formats strfmt.Registry) error {

	if err := validate.Required("getUserF2aOK"+"."+"secret", "body", o.Secret); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUserF2aOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUserF2aOKBody) UnmarshalBinary(b []byte) error {
	var res GetUserF2aOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}