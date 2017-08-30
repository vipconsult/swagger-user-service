// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	swag "github.com/go-openapi/swag"
)

// DeleteUserRoleHandlerFunc turns a function with the right signature into a delete user role handler
type DeleteUserRoleHandlerFunc func(DeleteUserRoleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteUserRoleHandlerFunc) Handle(params DeleteUserRoleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteUserRoleHandler interface for that can handle valid delete user role params
type DeleteUserRoleHandler interface {
	Handle(DeleteUserRoleParams, interface{}) middleware.Responder
}

// NewDeleteUserRole creates a new http.Handler for the delete user role operation
func NewDeleteUserRole(ctx *middleware.Context, handler DeleteUserRoleHandler) *DeleteUserRole {
	return &DeleteUserRole{Context: ctx, Handler: handler}
}

/*DeleteUserRole swagger:route DELETE /user/role deleteUserRole

deletes a role

*/
type DeleteUserRole struct {
	Context *middleware.Context
	Handler DeleteUserRoleHandler
}

func (o *DeleteUserRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteUserRoleParams()

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

// DeleteUserRoleBody delete user role body
// swagger:model DeleteUserRoleBody
type DeleteUserRoleBody struct {

	// id
	// Required: true
	ID *int64 `json:"id"`
}

// MarshalBinary interface implementation
func (o *DeleteUserRoleBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteUserRoleBody) UnmarshalBinary(b []byte) error {
	var res DeleteUserRoleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
