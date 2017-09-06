// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vanderbr/choicehealth_user-service/models"
)

// PutUserLoginOKCode is the HTTP code returned for type PutUserLoginOK
const PutUserLoginOKCode int = 200

/*PutUserLoginOK the user info associated with this token

swagger:response putUserLoginOK
*/
type PutUserLoginOK struct {

	/*
	  In: Body
	*/
	Payload *models.ProfileInfo `json:"body,omitempty"`
}

// NewPutUserLoginOK creates PutUserLoginOK with default headers values
func NewPutUserLoginOK() *PutUserLoginOK {
	return &PutUserLoginOK{}
}

// WithPayload adds the payload to the put user login o k response
func (o *PutUserLoginOK) WithPayload(payload *models.ProfileInfo) *PutUserLoginOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put user login o k response
func (o *PutUserLoginOK) SetPayload(payload *models.ProfileInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutUserLoginOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PutUserLoginDefault Generic Error used for most error responses - it returns a custom code and message depending on the reply context

swagger:response putUserLoginDefault
*/
type PutUserLoginDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewPutUserLoginDefault creates PutUserLoginDefault with default headers values
func NewPutUserLoginDefault(code int) *PutUserLoginDefault {
	if code <= 0 {
		code = 500
	}

	return &PutUserLoginDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put user login default response
func (o *PutUserLoginDefault) WithStatusCode(code int) *PutUserLoginDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put user login default response
func (o *PutUserLoginDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put user login default response
func (o *PutUserLoginDefault) WithPayload(payload *models.Response) *PutUserLoginDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put user login default response
func (o *PutUserLoginDefault) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutUserLoginDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
