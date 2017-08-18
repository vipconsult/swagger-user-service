// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/choicehealth/user-service/models"
)

// PostLoginOKCode is the HTTP code returned for type PostLoginOK
const PostLoginOKCode int = 200

/*PostLoginOK A token object.

swagger:response postLoginOK
*/
type PostLoginOK struct {

	/*
	  In: Body
	*/
	Payload PostLoginOKBody `json:"body,omitempty"`
}

// NewPostLoginOK creates PostLoginOK with default headers values
func NewPostLoginOK() *PostLoginOK {
	return &PostLoginOK{}
}

// WithPayload adds the payload to the post login o k response
func (o *PostLoginOK) WithPayload(payload PostLoginOKBody) *PostLoginOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post login o k response
func (o *PostLoginOK) SetPayload(payload PostLoginOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostLoginOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// PostLoginNotFoundCode is the HTTP code returned for type PostLoginNotFound
const PostLoginNotFoundCode int = 404

/*PostLoginNotFound Resource not found

swagger:response postLoginNotFound
*/
type PostLoginNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewPostLoginNotFound creates PostLoginNotFound with default headers values
func NewPostLoginNotFound() *PostLoginNotFound {
	return &PostLoginNotFound{}
}

// WithPayload adds the payload to the post login not found response
func (o *PostLoginNotFound) WithPayload(payload *models.Response) *PostLoginNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post login not found response
func (o *PostLoginNotFound) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostLoginNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostLoginDefault Unexpected error

swagger:response postLoginDefault
*/
type PostLoginDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewPostLoginDefault creates PostLoginDefault with default headers values
func NewPostLoginDefault(code int) *PostLoginDefault {
	if code <= 0 {
		code = 500
	}

	return &PostLoginDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post login default response
func (o *PostLoginDefault) WithStatusCode(code int) *PostLoginDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post login default response
func (o *PostLoginDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post login default response
func (o *PostLoginDefault) WithPayload(payload *models.Response) *PostLoginDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post login default response
func (o *PostLoginDefault) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostLoginDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
