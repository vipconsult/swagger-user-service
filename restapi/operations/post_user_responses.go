// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vanderbr/choicehealth_user-service/models"
)

// PostUserOKCode is the HTTP code returned for type PostUserOK
const PostUserOKCode int = 200

/*PostUserOK id of the created user.

swagger:response postUserOK
*/
type PostUserOK struct {

	/*
	  In: Body
	*/
	Payload PostUserOKBody `json:"body,omitempty"`
}

// NewPostUserOK creates PostUserOK with default headers values
func NewPostUserOK() *PostUserOK {
	return &PostUserOK{}
}

// WithPayload adds the payload to the post user o k response
func (o *PostUserOK) WithPayload(payload PostUserOKBody) *PostUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user o k response
func (o *PostUserOK) SetPayload(payload PostUserOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// PostUserConflictCode is the HTTP code returned for type PostUserConflict
const PostUserConflictCode int = 409

/*PostUserConflict Username already taken

swagger:response postUserConflict
*/
type PostUserConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewPostUserConflict creates PostUserConflict with default headers values
func NewPostUserConflict() *PostUserConflict {
	return &PostUserConflict{}
}

// WithPayload adds the payload to the post user conflict response
func (o *PostUserConflict) WithPayload(payload *models.Response) *PostUserConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user conflict response
func (o *PostUserConflict) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostUserDefault Generic Error used for most error responses - it returns a custom code and message depending on the reply context

swagger:response postUserDefault
*/
type PostUserDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewPostUserDefault creates PostUserDefault with default headers values
func NewPostUserDefault(code int) *PostUserDefault {
	if code <= 0 {
		code = 500
	}

	return &PostUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post user default response
func (o *PostUserDefault) WithStatusCode(code int) *PostUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post user default response
func (o *PostUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post user default response
func (o *PostUserDefault) WithPayload(payload *models.Response) *PostUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user default response
func (o *PostUserDefault) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}