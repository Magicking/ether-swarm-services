package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Magicking/ether-swarm-services/models"
)

/*GetInformationOK Information of the blockchain.

swagger:response getInformationOK
*/
type GetInformationOK struct {

	/*
	  In: Body
	*/
	Payload *models.Information `json:"body,omitempty"`
}

// NewGetInformationOK creates GetInformationOK with default headers values
func NewGetInformationOK() *GetInformationOK {
	return &GetInformationOK{}
}

// WithPayload adds the payload to the get information o k response
func (o *GetInformationOK) WithPayload(payload *models.Information) *GetInformationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get information o k response
func (o *GetInformationOK) SetPayload(payload *models.Information) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInformationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetInformationDefault Unexpected error

swagger:response getInformationDefault
*/
type GetInformationDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetInformationDefault creates GetInformationDefault with default headers values
func NewGetInformationDefault(code int) *GetInformationDefault {
	if code <= 0 {
		code = 500
	}

	return &GetInformationDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get information default response
func (o *GetInformationDefault) WithStatusCode(code int) *GetInformationDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get information default response
func (o *GetInformationDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get information default response
func (o *GetInformationDefault) WithPayload(payload *models.Error) *GetInformationDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get information default response
func (o *GetInformationDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInformationDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
