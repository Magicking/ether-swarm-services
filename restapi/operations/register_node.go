package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// RegisterNodeHandlerFunc turns a function with the right signature into a register node handler
type RegisterNodeHandlerFunc func(RegisterNodeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RegisterNodeHandlerFunc) Handle(params RegisterNodeParams) middleware.Responder {
	return fn(params)
}

// RegisterNodeHandler interface for that can handle valid register node params
type RegisterNodeHandler interface {
	Handle(RegisterNodeParams) middleware.Responder
}

// NewRegisterNode creates a new http.Handler for the register node operation
func NewRegisterNode(ctx *middleware.Context, handler RegisterNodeHandler) *RegisterNode {
	return &RegisterNode{Context: ctx, Handler: handler}
}

/*RegisterNode swagger:route POST /blockchain/bootnode registerNode

Register a node

Register (boot)nodes that other nodes while connect to.


*/
type RegisterNode struct {
	Context *middleware.Context
	Handler RegisterNodeHandler
}

func (o *RegisterNode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewRegisterNodeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
