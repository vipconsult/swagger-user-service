// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vanderbr/choicehealth_user-service/models"
)

// PostUserRoleOKCode is the HTTP code returned for type PostUserRoleOK
const PostUserRoleOKCode int = 200

/*PostUserRoleOK the id of the created role.

swagger:response postUserRoleOK
*/
type PostUserRoleOK struct {

	/*
	  In: Body
	*/
	Payload PostUserRoleOKBody `json:"body,omitempty"`
}

// NewPostUserRoleOK creates PostUserRoleOK with default headers values
func NewPostUserRoleOK() *PostUserRoleOK {
	return &PostUserRoleOK{}
}

// WithPayload adds the payload to the post user role o k response
func (o *PostUserRoleOK) WithPayload(payload PostUserRoleOKBody) *PostUserRoleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user role o k response
func (o *PostUserRoleOK) SetPayload(payload PostUserRoleOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserRoleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// PostUserRoleUnauthorizedCode is the HTTP code returned for type PostUserRoleUnauthorized
const PostUserRoleUnauthorizedCode int = 401

/*PostUserRoleUnauthorized Authentication is missing or invalid

swagger:response postUserRoleUnauthorized
*/
type PostUserRoleUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewPostUserRoleUnauthorized creates PostUserRoleUnauthorized with default headers values
func NewPostUserRoleUnauthorized() *PostUserRoleUnauthorized {
	return &PostUserRoleUnauthorized{}
}

// WithPayload adds the payload to the post user role unauthorized response
func (o *PostUserRoleUnauthorized) WithPayload(payload *models.Response) *PostUserRoleUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user role unauthorized response
func (o *PostUserRoleUnauthorized) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserRoleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostUserRoleDefault Unexpected error

swagger:response postUserRoleDefault
*/
type PostUserRoleDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewPostUserRoleDefault creates PostUserRoleDefault with default headers values
func NewPostUserRoleDefault(code int) *PostUserRoleDefault {
	if code <= 0 {
		code = 500
	}

	return &PostUserRoleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post user role default response
func (o *PostUserRoleDefault) WithStatusCode(code int) *PostUserRoleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post user role default response
func (o *PostUserRoleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post user role default response
func (o *PostUserRoleDefault) WithPayload(payload *models.Response) *PostUserRoleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user role default response
func (o *PostUserRoleDefault) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserRoleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}