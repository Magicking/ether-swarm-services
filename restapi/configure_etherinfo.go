package restapi

import (
	"fmt"
	"crypto/tls"
	"net/http"
	"log"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	swag "github.com/go-openapi/swag"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/Magicking/ether-swarm-services/models"
	"github.com/Magicking/ether-swarm-services/restapi/operations"

	"github.com/Magicking/ether-swarm-services/internal"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ../etherinfo.yml

var genesis_opts internal.GenesisConf

var svc_opts struct {
	DbDsn     string `long:"db-dsn" env:"DB_DSN" default:"host=localhost dbname=etherinfo sslmode=disable" description:"Data source name, currently only support postgres"`
	NetworkID int64 `long:"networkid" env:"NETWORK_ID" default:"6120" description:"Ethereum network id"`
}

func configureFlags(api *operations.EtherinfoAPI) {
	etherinfo_opts_grp := swag.CommandLineOptionsGroup{
		LongDescription:  "Configurations options for this service",
		ShortDescription: "Service options",
		Options:          &svc_opts,
	}
	genesis_opts_grp := swag.CommandLineOptionsGroup{
		LongDescription:  "Default entries for genesis block",
		ShortDescription: "Genesis options",
		Options:          &genesis_opts,
	}
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		genesis_opts_grp,
		etherinfo_opts_grp,
	}
}

func configureAPI(api *operations.EtherinfoAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	ctx, err := internal.NewContext(svc_opts.DbDsn)
	if err != nil {
		log.Fatalf("configureAPI: %v", err)
	}
	ctx.NetworkID = svc_opts.NetworkID

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.CreateGenesisHandler = operations.CreateGenesisHandlerFunc(func(params operations.CreateGenesisParams) middleware.Responder {
		genesis, err := internal.CreateGenesisHandler(ctx, &genesis_opts, params)
		if err != nil {
			msg := fmt.Sprintf("CreateGenesisHandlerFunc: %v", err)
			return operations.NewCreateGenesisDefault(500).WithPayload(&models.Error{
				Message: &msg,
			})
		}
		return operations.NewCreateGenesisOK().WithPayload(genesis)
	})
	api.GetInformationHandler = operations.GetInformationHandlerFunc(func(params operations.GetInformationParams) middleware.Responder {
		infos, err := internal.GetInformationHandler(ctx)
		if err != nil {
			msg := fmt.Sprintf("GetInformationHandlerFunc: %v", err)
			return operations.NewGetInformationDefault(500).WithPayload(&models.Error{
				Message: &msg,
			})
		}
		return operations.NewGetInformationOK().WithPayload(infos)
	})
	api.RegisterNodeHandler = operations.RegisterNodeHandlerFunc(func(params operations.RegisterNodeParams) middleware.Responder {
		fSuccess, err := internal.RegisterNodeHandler(ctx, params)
		if err != nil {
			msg := fmt.Sprintf("RegisterNodeHandlerFunc: %v", err)
			return operations.NewRegisterNodeDefault(500).WithPayload(&models.Error{
				Message: &msg,
			})
		}
		return operations.NewRegisterNodeOK().WithPayload(fSuccess)
	})
	api.StartNodeHandler = operations.StartNodeHandlerFunc(func(params operations.StartNodeParams) middleware.Responder {
		fSuccess, err := internal.StartNodeHandler(ctx, params)
		if err != nil {
			msg := fmt.Sprintf("StartNodeHandlerFunc: %v", err)
			return operations.NewStartNodeDefault(500).WithPayload(&models.Error{
				Message: &msg,
			})
		}
		return operations.NewStartNodeOK().WithPayload(fSuccess)
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
