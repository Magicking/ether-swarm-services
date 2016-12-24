package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	swag "github.com/go-openapi/swag"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/Magicking/ether-swarm-services/restapi/operations"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ../etherinfo.yml

var genesis_opts struct {
	Nonce       string `long:"nonce" env:"GG_NONCE" default:"0x0000000000000042" description:"Nonce of the genesis block"`
	Timestamp   string `long:"timestamp" env:"GG_TIMESTAMP" default:"0x0" description:"Timestamp of the genesis block"`
	ParentHash  string `long:"parent-hash" env:"GG_PARENT_HASH" default:"0x0" description:"Parent hash of the genesis block"`
	ExtraData   string `long:"extra-data" env:"GG_EXTRA_DATA" description:"Extra data of the genesis block"`
	GasLimit    string `long:"gas-limit" env:"GG_GAS_LIMIT" default:"0x8000000" description:"Gas limit of the genesis block"`
	Difficulty  string `long:"difficulty" env:"GG_DIFFICULTY" default:"0x400" description:"Initial difficulty"`
	Mixhash     string `long:"mixhash" env:"GG_MIX_HASH" default:"0x0" description:"Mixhash of the genesis block"`
	Coinbase    string `long:"coinbase" env:"GG_COINBASE" default:"0x0" description:"Coinbase of the genesis block"`
	Balance     string `long:"balance" env:"GG_BALANCE" default:"1000000000000000000" description:"Default balance in Wei for new account when new_allocator is used"`
}

func configureFlags(api *operations.EtherinfoAPI) {
	genesis_opts_grp := swag.CommandLineOptionsGroup{
		LongDescription:  "Default entries for genesis block",
		ShortDescription: "Genesis options",
		Options:          &genesis_opts,
	}
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{genesis_opts_grp}
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

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme string) {
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
