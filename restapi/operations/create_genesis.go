package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateGenesisHandlerFunc turns a function with the right signature into a create genesis handler
type CreateGenesisHandlerFunc func(CreateGenesisParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateGenesisHandlerFunc) Handle(params CreateGenesisParams) middleware.Responder {
	return fn(params)
}

// CreateGenesisHandler interface for that can handle valid create genesis params
type CreateGenesisHandler interface {
	Handle(CreateGenesisParams) middleware.Responder
}

// NewCreateGenesis creates a new http.Handler for the create genesis operation
func NewCreateGenesis(ctx *middleware.Context, handler CreateGenesisHandler) *CreateGenesis {
	return &CreateGenesis{Context: ctx, Handler: handler}
}

/*CreateGenesis swagger:route POST /blockchain/create createGenesis

Create genesis

Create a genesis and lock it in database, if the optional genesis
contains allocators new_allocator is ignored.


*/
type CreateGenesis struct {
	Context *middleware.Context
	Handler CreateGenesisHandler
}

func (o *CreateGenesis) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewCreateGenesisParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
