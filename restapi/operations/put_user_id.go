// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutUserIDHandlerFunc turns a function with the right signature into a put user ID handler
type PutUserIDHandlerFunc func(PutUserIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PutUserIDHandlerFunc) Handle(params PutUserIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PutUserIDHandler interface for that can handle valid put user ID params
type PutUserIDHandler interface {
	Handle(PutUserIDParams, interface{}) middleware.Responder
}

// NewPutUserID creates a new http.Handler for the put user ID operation
func NewPutUserID(ctx *middleware.Context, handler PutUserIDHandler) *PutUserID {
	return &PutUserID{Context: ctx, Handler: handler}
}

/*PutUserID swagger:route PUT /user/{id} putUserId

updates an existing user, only submited fields will be updated so can ommit the ones that don't need updating

*/
type PutUserID struct {
	Context *middleware.Context
	Handler PutUserIDHandler
}

func (o *PutUserID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutUserIDParams()

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