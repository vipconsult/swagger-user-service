// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vanderbr/choicehealth_user-service/models"
)

// PostUserLoginOKCode is the HTTP code returned for type PostUserLoginOK
const PostUserLoginOKCode int = 200

/*PostUserLoginOK A jwt token to use for authentication.

swagger:response postUserLoginOK
*/
type PostUserLoginOK struct {

	/*
	  In: Body
	*/
	Payload *models.Jwt `json:"body,omitempty"`
}

// NewPostUserLoginOK creates PostUserLoginOK with default headers values
func NewPostUserLoginOK() *PostUserLoginOK {
	return &PostUserLoginOK{}
}

// WithPayload adds the payload to the post user login o k response
func (o *PostUserLoginOK) WithPayload(payload *models.Jwt) *PostUserLoginOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user login o k response
func (o *PostUserLoginOK) SetPayload(payload *models.Jwt) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserLoginOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUserLoginCreatedCode is the HTTP code returned for type PostUserLoginCreated
const PostUserLoginCreatedCode int = 201

/*PostUserLoginCreated Password change is required, hit the password reset endpoint with the generated jwt token

swagger:response postUserLoginCreated
*/
type PostUserLoginCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Jwt `json:"body,omitempty"`
}

// NewPostUserLoginCreated creates PostUserLoginCreated with default headers values
func NewPostUserLoginCreated() *PostUserLoginCreated {
	return &PostUserLoginCreated{}
}

// WithPayload adds the payload to the post user login created response
func (o *PostUserLoginCreated) WithPayload(payload *models.Jwt) *PostUserLoginCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user login created response
func (o *PostUserLoginCreated) SetPayload(payload *models.Jwt) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserLoginCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUserLoginPartialContentCode is the HTTP code returned for type PostUserLoginPartialContent
const PostUserLoginPartialContentCode int = 206

/*PostUserLoginPartialContent Account is with 2 factor authenticaiton so use the 2 factor endpoint to generate the final the jwt token.

swagger:response postUserLoginPartialContent
*/
type PostUserLoginPartialContent struct {

	/*
	  In: Body
	*/
	Payload *models.Jwt `json:"body,omitempty"`
}

// NewPostUserLoginPartialContent creates PostUserLoginPartialContent with default headers values
func NewPostUserLoginPartialContent() *PostUserLoginPartialContent {
	return &PostUserLoginPartialContent{}
}

// WithPayload adds the payload to the post user login partial content response
func (o *PostUserLoginPartialContent) WithPayload(payload *models.Jwt) *PostUserLoginPartialContent {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user login partial content response
func (o *PostUserLoginPartialContent) SetPayload(payload *models.Jwt) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserLoginPartialContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(206)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostUserLoginDefault Unexpected error

swagger:response postUserLoginDefault
*/
type PostUserLoginDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewPostUserLoginDefault creates PostUserLoginDefault with default headers values
func NewPostUserLoginDefault(code int) *PostUserLoginDefault {
	if code <= 0 {
		code = 500
	}

	return &PostUserLoginDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post user login default response
func (o *PostUserLoginDefault) WithStatusCode(code int) *PostUserLoginDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post user login default response
func (o *PostUserLoginDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post user login default response
func (o *PostUserLoginDefault) WithPayload(payload *models.Response) *PostUserLoginDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user login default response
func (o *PostUserLoginDefault) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserLoginDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
