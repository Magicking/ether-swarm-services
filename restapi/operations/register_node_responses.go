package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Magicking/ether-swarm-services/models"
)

/*RegisterNodeOK Indicate if the operation was successful.

swagger:response registerNodeOK
*/
type RegisterNodeOK struct {

	/*
	  In: Body
	*/
	Payload bool `json:"body,omitempty"`
}

// NewRegisterNodeOK creates RegisterNodeOK with default headers values
func NewRegisterNodeOK() *RegisterNodeOK {
	return &RegisterNodeOK{}
}

// WithPayload adds the payload to the register node o k response
func (o *RegisterNodeOK) WithPayload(payload bool) *RegisterNodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register node o k response
func (o *RegisterNodeOK) SetPayload(payload bool) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterNodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*RegisterNodeDefault Unexpected error

swagger:response registerNodeDefault
*/
type RegisterNodeDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRegisterNodeDefault creates RegisterNodeDefault with default headers values
func NewRegisterNodeDefault(code int) *RegisterNodeDefault {
	if code <= 0 {
		code = 500
	}

	return &RegisterNodeDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the register node default response
func (o *RegisterNodeDefault) WithStatusCode(code int) *RegisterNodeDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the register node default response
func (o *RegisterNodeDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the register node default response
func (o *RegisterNodeDefault) WithPayload(payload *models.Error) *RegisterNodeDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register node default response
func (o *RegisterNodeDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterNodeDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
