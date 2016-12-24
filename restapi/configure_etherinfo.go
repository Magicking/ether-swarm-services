package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/Magicking/ether-swarm-services/restapi/operations"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ../etherinfo.yml

func configureFlags(api *operations.EtherinfoAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.EtherinfoAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.CreateGenesisHandler = operations.CreateGenesisHandlerFunc(func(params operations.CreateGenesisParams) middleware.Responder {
		return middleware.NotImplemented("operation .CreateGenesis has not yet been implemented")
	})
	api.GetInformationHandler = operations.GetInformationHandlerFunc(func(params operations.GetInformationParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetInformation has not yet been implemented")
	})
	api.RegisterNodeHandler = operations.RegisterNodeHandlerFunc(func(params operations.RegisterNodeParams) middleware.Responder {
		return middleware.NotImplemented("operation .RegisterNode has not yet been implemented")
	})
	api.StartNodeHandler = operations.StartNodeHandlerFunc(func(params operations.StartNodeParams) middleware.Responder {
		return middleware.NotImplemented("operation .StartNode has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
