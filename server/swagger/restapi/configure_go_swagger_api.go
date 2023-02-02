// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	"quocbang/go-swagger-api/server"
	"quocbang/go-swagger-api/server/configs"
	"quocbang/go-swagger-api/server/impl/handlers"
	"quocbang/go-swagger-api/server/swagger/models"
	"quocbang/go-swagger-api/server/swagger/restapi/operations"
	"quocbang/go-swagger-api/server/swagger/restapi/operations/account"
)

var (
	options       = new(configs.Options)
	configuration = new(configs.Configs)
)

func configureFlags(api *operations.GoSwaggerAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
	api.CommandLineOptionsGroups = append(api.CommandLineOptionsGroups, swag.CommandLineOptionsGroup{
		ShortDescription: "Configuration Options",
		LongDescription:  "Configuration Options",
		Options:          options,
	})
}

// read config file and return to the pointer variables
func parseConfigurations(configFilePath string) (*configs.Configs, error) {
	configsFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var cfgs configs.Configs
	if err := yaml.Unmarshal(configsFile, &cfgs); err != nil {
		return nil, err
	}

	return &cfgs, nil

}

func configureAPI(api *operations.GoSwaggerAPIAPI) http.Handler {

	var err error
	// check file config
	configuration, err = parseConfigurations(options.ServerConfig)
	if err != nil {
		log.Fatalf("failed to parse configurations file. err= %s", err.Error())
	}

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

	dm, err := server.RegisterDataManager(*configuration)
	if err != nil {
		zap.L().Fatal("failed to register data manager", zap.Error(err))
	}
	// [NOTE] if you want try a test without real api, please switch import path from `/server/impl/mcom` to `/server/impl/mock`
	serviceConfig := handlers.ServiceConfig{
		TokenLifeTime: time.Duration(configuration.TokenExpiredSeconds) * time.Second,
	}

	if err := handlers.RegisterHandlers(dm, api, serviceConfig); err != nil {
		zap.L().Fatal("failed to register handlers", zap.Error(err))
	}

	// Applies when the "x-api-auth-key" header is set
	if api.APIKeyAuth == nil {
		api.APIKeyAuth = func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (api_key) x-api-auth-key from header param [x-api-auth-key] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.AccountChangePasswordHandler == nil {
		api.AccountChangePasswordHandler = account.ChangePasswordHandlerFunc(func(params account.ChangePasswordParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation account.ChangePassword has not yet been implemented")
		})
	}
	if api.CheckServerStatusHandler == nil {
		api.CheckServerStatusHandler = operations.CheckServerStatusHandlerFunc(func(params operations.CheckServerStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CheckServerStatus has not yet been implemented")
		})
	}
	if api.AccountCreateAccountAuthorizationHandler == nil {
		api.AccountCreateAccountAuthorizationHandler = account.CreateAccountAuthorizationHandlerFunc(func(params account.CreateAccountAuthorizationParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation account.CreateAccountAuthorization has not yet been implemented")
		})
	}
	if api.AccountDeleteAccountHandler == nil {
		api.AccountDeleteAccountHandler = account.DeleteAccountHandlerFunc(func(params account.DeleteAccountParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation account.DeleteAccount has not yet been implemented")
		})
	}
	if api.AccountGetRoleListHandler == nil {
		api.AccountGetRoleListHandler = account.GetRoleListHandlerFunc(func(params account.GetRoleListParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation account.GetRoleList has not yet been implemented")
		})
	}
	if api.AccountListAuthorizedAccountHandler == nil {
		api.AccountListAuthorizedAccountHandler = account.ListAuthorizedAccountHandlerFunc(func(params account.ListAuthorizedAccountParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation account.ListAuthorizedAccount has not yet been implemented")
		})
	}
	if api.AccountListUnauthorizedAccountHandler == nil {
		api.AccountListUnauthorizedAccountHandler = account.ListUnauthorizedAccountHandlerFunc(func(params account.ListUnauthorizedAccountParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation account.ListUnauthorizedAccount has not yet been implemented")
		})
	}
	if api.AccountLoginHandler == nil {
		api.AccountLoginHandler = account.LoginHandlerFunc(func(params account.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation account.Login has not yet been implemented")
		})
	}
	if api.AccountLogoutHandler == nil {
		api.AccountLogoutHandler = account.LogoutHandlerFunc(func(params account.LogoutParams) middleware.Responder {
			return middleware.NotImplemented("operation account.Logout has not yet been implemented")
		})
	}
	if api.AccountUpdateAccountAuthorizationHandler == nil {
		api.AccountUpdateAccountAuthorizationHandler = account.UpdateAccountAuthorizationHandlerFunc(func(params account.UpdateAccountAuthorizationParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation account.UpdateAccountAuthorization has not yet been implemented")
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
