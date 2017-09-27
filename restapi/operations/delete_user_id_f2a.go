// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteUserIDF2aHandlerFunc turns a function with the right signature into a delete user ID f2a handler
type DeleteUserIDF2aHandlerFunc func(DeleteUserIDF2aParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteUserIDF2aHandlerFunc) Handle(params DeleteUserIDF2aParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteUserIDF2aHandler interface for that can handle valid delete user ID f2a params
type DeleteUserIDF2aHandler interface {
	Handle(DeleteUserIDF2aParams, interface{}) middleware.Responder
}

// NewDeleteUserIDF2a creates a new http.Handler for the delete user ID f2a operation
func NewDeleteUserIDF2a(ctx *middleware.Context, handler DeleteUserIDF2aHandler) *DeleteUserIDF2a {
	return &DeleteUserIDF2a{Context: ctx, Handler: handler}
}

/*DeleteUserIDF2a swagger:route DELETE /user/{id}/f2a deleteUserIdF2a

disable 2 factor authenticaiton for an account.

*/
type DeleteUserIDF2a struct {
	Context *middleware.Context
	Handler DeleteUserIDF2aHandler
}

func (o *DeleteUserIDF2a) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteUserIDF2aParams()

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