package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetInformationHandlerFunc turns a function with the right signature into a get information handler
type GetInformationHandlerFunc func(GetInformationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetInformationHandlerFunc) Handle(params GetInformationParams) middleware.Responder {
	return fn(params)
}

// GetInformationHandler interface for that can handle valid get information params
type GetInformationHandler interface {
	Handle(GetInformationParams) middleware.Responder
}

// NewGetInformation creates a new http.Handler for the get information operation
func NewGetInformation(ctx *middleware.Context, handler GetInformationHandler) *GetInformation {
	return &GetInformation{Context: ctx, Handler: handler}
}

/*GetInformation swagger:route GET /blockchain/info getInformation

Return information necessary to participate to the blockchain

Return information necessary to bootstrap a node
e.g. genesis, bootnodes url, ...


*/
type GetInformation struct {
	Context *middleware.Context
	Handler GetInformationHandler
}

func (o *GetInformation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetInformationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
