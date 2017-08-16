// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"crypto/rsa"
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/rs/cors"
	graceful "github.com/tylerb/graceful"

	"github.com/choicehealth/user-service/restapi/operations"
	"github.com/choicehealth/user-service/restapi/operations/users"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ../swagger.yaml
var keysPaths = struct {
	Pub  string `long:"pub" description:"public key path for the swt token generation" default:"/tmp/app.rsa.pub"`
	Priv string `long:"priv" description:"private key path for the swt token verification" default:"/tmp/app.rsa"`
}{}

func configureFlags(api *operations.UserManagementAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		swag.CommandLineOptionsGroup{
			ShortDescription: "private pub key paths",
			Options:          &keysPaths,
		},
	}
}

func configureAPI(api *operations.UserManagementAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.UsersPostLoginHandler = users.PostLoginHandlerFunc(func(params users.PostLoginParams) middleware.Responder {
		token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
			"exp": time.Now().Add(time.Hour * 72).Unix(),
		})

		t, err := token.SignedString(signKey)
		if err != nil {
			return users.NewPostLoginDefault(0)
		}
		return users.NewPostLoginOK().WithPayload(users.PostLoginOKBody{Token: swag.String(t)})
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
func configureServer(s *graceful.Server, scheme, addr string) {

	if key, err := ioutil.ReadFile(keysPaths.Priv); err != nil {
		log.Fatalf("can't read public file:%v", err)
	} else {
		signKey, err = jwt.ParseRSAPrivateKeyFromPEM(key)
		if err != nil {
			log.Fatalf("error parsing the  private key :%v", err)
		}
	}
	if key, err := ioutil.ReadFile(keysPaths.Pub); err != nil {
		log.Fatalf("can't read private file:%v", err)

	} else {
		verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(key)
		if err != nil {
			log.Fatalf("error parsing the  public key :%v", err)
		}
	}
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler

	return handleCORS(handler)
}
