// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vanderbr/choicehealth_user-service/models"
)

// GetUserRoleOKCode is the HTTP code returned for type GetUserRoleOK
const GetUserRoleOKCode int = 200

/*GetUserRoleOK full roles list

swagger:response getUserRoleOK
*/
type GetUserRoleOK struct {

	/*
	  In: Body
	*/
	Payload []*models.UserRole `json:"body,omitempty"`
}

// NewGetUserRoleOK creates GetUserRoleOK with default headers values
func NewGetUserRoleOK() *GetUserRoleOK {
	return &GetUserRoleOK{}
}

// WithPayload adds the payload to the get user role o k response
func (o *GetUserRoleOK) WithPayload(payload []*models.UserRole) *GetUserRoleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user role o k response
func (o *GetUserRoleOK) SetPayload(payload []*models.UserRole) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserRoleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.UserRole, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetUserRoleUnauthorizedCode is the HTTP code returned for type GetUserRoleUnauthorized
const GetUserRoleUnauthorizedCode int = 401

/*GetUserRoleUnauthorized Authentication is missing or invalid

swagger:response getUserRoleUnauthorized
*/
type GetUserRoleUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewGetUserRoleUnauthorized creates GetUserRoleUnauthorized with default headers values
func NewGetUserRoleUnauthorized() *GetUserRoleUnauthorized {
	return &GetUserRoleUnauthorized{}
}

// WithPayload adds the payload to the get user role unauthorized response
func (o *GetUserRoleUnauthorized) WithPayload(payload *models.Response) *GetUserRoleUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user role unauthorized response
func (o *GetUserRoleUnauthorized) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserRoleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetUserRoleDefault Unexpected error

swagger:response getUserRoleDefault
*/
type GetUserRoleDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewGetUserRoleDefault creates GetUserRoleDefault with default headers values
func NewGetUserRoleDefault(code int) *GetUserRoleDefault {
	if code <= 0 {
		code = 500
	}

	return &GetUserRoleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get user role default response
func (o *GetUserRoleDefault) WithStatusCode(code int) *GetUserRoleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get user role default response
func (o *GetUserRoleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get user role default response
func (o *GetUserRoleDefault) WithPayload(payload *models.Response) *GetUserRoleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user role default response
func (o *GetUserRoleDefault) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserRoleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
