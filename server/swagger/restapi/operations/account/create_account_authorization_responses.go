// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"quocbang/go-swagger-api/server/swagger/models"
)

// CreateAccountAuthorizationOKCode is the HTTP code returned for type CreateAccountAuthorizationOK
const CreateAccountAuthorizationOKCode int = 200

/*CreateAccountAuthorizationOK OK

swagger:response createAccountAuthorizationOK
*/
type CreateAccountAuthorizationOK struct {
}

// NewCreateAccountAuthorizationOK creates CreateAccountAuthorizationOK with default headers values
func NewCreateAccountAuthorizationOK() *CreateAccountAuthorizationOK {

	return &CreateAccountAuthorizationOK{}
}

// WriteResponse to the client
func (o *CreateAccountAuthorizationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*CreateAccountAuthorizationDefault Unexpected error

swagger:response createAccountAuthorizationDefault
*/
type CreateAccountAuthorizationDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateAccountAuthorizationDefault creates CreateAccountAuthorizationDefault with default headers values
func NewCreateAccountAuthorizationDefault(code int) *CreateAccountAuthorizationDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateAccountAuthorizationDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create account authorization default response
func (o *CreateAccountAuthorizationDefault) WithStatusCode(code int) *CreateAccountAuthorizationDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create account authorization default response
func (o *CreateAccountAuthorizationDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create account authorization default response
func (o *CreateAccountAuthorizationDefault) WithPayload(payload *models.Error) *CreateAccountAuthorizationDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create account authorization default response
func (o *CreateAccountAuthorizationDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAccountAuthorizationDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}