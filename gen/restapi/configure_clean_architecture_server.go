// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/negiseijin/clean-architectur-swagger/gen/restapi/operations"
	"github.com/negiseijin/clean-architectur-swagger/gen/restapi/operations/todo"
	"github.com/negiseijin/clean-architectur-swagger/handler"
	"github.com/negiseijin/clean-architectur-swagger/infrastructure"
	"github.com/negiseijin/clean-architectur-swagger/infrastructure/persistence"
	"github.com/negiseijin/clean-architectur-swagger/migration"
)

//go:generate swagger generate server --target ../../gen --name CleanArchitectureServer --spec ../../swagger/swagger.yml --principal interface{} --exclude-main

func configureFlags(api *operations.CleanArchitectureServerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CleanArchitectureServerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// DB Settting
	db := infrastructure.NewDevelopmentDB()
	db.NewDB()

	m := migration.NewMigration(db.Connection)
	m.Migrate()

	th := handler.NewTodoHandler(
		persistence.DBRepository{
			DB: db.Connection,
		},
	)

	if api.TodoGetTodoListHandler != nil {
		api.TodoGetTodoListHandler = todo.GetTodoListHandlerFunc(func(params todo.GetTodoListParams) middleware.Responder {
			return th.GetTodoListHandler(params)
		})
	}

	if api.TodoGetTodoHandler != nil {
		api.TodoGetTodoHandler = todo.GetTodoHandlerFunc(func(params todo.GetTodoParams) middleware.Responder {
			return th.GetTodoHandler(params)
		})
	}

	if api.TodoCreateTodoHandler != nil {
		api.TodoCreateTodoHandler = todo.CreateTodoHandlerFunc(func(params todo.CreateTodoParams) middleware.Responder {
			return th.PostTodoHandler(params)
		})
	}

	api.PreServerShutdown = func() {}

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
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
