// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutUserRoleHandlerFunc turns a function with the right signature into a put user role handler
type PutUserRoleHandlerFunc func(PutUserRoleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PutUserRoleHandlerFunc) Handle(params PutUserRoleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PutUserRoleHandler interface for that can handle valid put user role params
type PutUserRoleHandler interface {
	Handle(PutUserRoleParams, interface{}) middleware.Responder
}

// NewPutUserRole creates a new http.Handler for the put user role operation
func NewPutUserRole(ctx *middleware.Context, handler PutUserRoleHandler) *PutUserRole {
	return &PutUserRole{Context: ctx, Handler: handler}
}

/*PutUserRole swagger:route PUT /user/role putUserRole

updates a role

*/
type PutUserRole struct {
	Context *middleware.Context
	Handler PutUserRoleHandler
}

func (o *PutUserRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutUserRoleParams()

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
