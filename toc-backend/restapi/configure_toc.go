// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"
	"strings"

	"github.com/rs/cors"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"github.com/gyk4j/toc/toc/server/controllers"
	"github.com/gyk4j/toc/toc/server/restapi/operations"
	"github.com/gyk4j/toc/toc/server/restapi/operations/archive"
	"github.com/gyk4j/toc/toc/server/restapi/operations/backup"
	logops "github.com/gyk4j/toc/toc/server/restapi/operations/log"
	"github.com/gyk4j/toc/toc/server/restapi/operations/quota"
	"github.com/gyk4j/toc/toc/server/restapi/operations/restoration"
	"github.com/gyk4j/toc/toc/server/restapi/operations/synchronization"
)

//go:generate swagger generate server --target ..\..\sw --name Toc --spec ..\swagger.yaml --principal interface{}

func configureFlags(api *operations.TocAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TocAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.ArchiveArchiveDataHandler = archive.ArchiveDataHandlerFunc(controllers.NewArchive)
	api.LogExportLogHandler = logops.ExportLogHandlerFunc(controllers.ExportLog)
	api.ArchiveGetArchiveHandler = archive.GetArchiveHandlerFunc(controllers.GetArchive)
	api.ArchiveGetArchiveByIDHandler = archive.GetArchiveByIDHandlerFunc(controllers.GetArchiveByID)
	api.BackupGetBackupHandler = backup.GetBackupHandlerFunc(controllers.GetBackups)
	api.BackupGetBackupByIDHandler = backup.GetBackupByIDHandlerFunc(controllers.GetBackupByID)
	api.QuotaGetQuotaHandler = quota.GetQuotaHandlerFunc(controllers.GetQuotas)
	api.RestorationGetRestorationHandler = restoration.GetRestorationHandlerFunc(controllers.GetRestorations)
	api.RestorationGetRestorationByIDHandler = restoration.GetRestorationByIDHandlerFunc(controllers.GetRestorationByID)
	api.BackupNewBackupHandler = backup.NewBackupHandlerFunc(controllers.NewBackup)
	api.RestorationNewRestorationHandler = restoration.NewRestorationHandlerFunc(controllers.NewRestoration)
	api.SynchronizationNewSynchronizationHandler = synchronization.NewSynchronizationHandlerFunc(controllers.NewSynchronization)

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
	// NOTE: Using http.FileServer does not allow chaining of the REST API HandlerFunc "handler" parameter.
	//       If used, the UI static web files can be served, but the API endpoints would not work.
	//       Thus, this method is commented out, but left for reference purpose.
	// staticHandler := http.StripPrefix("/web/", http.FileServer(http.Dir("web")))
	// http.Handle("/web/", staticHandler)
	// return staticHandler

	handlerStatic := FileServerMiddleware(handler)
	handlerCORS := cors.Default().Handler
	return handlerCORS(handlerStatic)
}

// References:
// - https://notes.elmiko.dev/2016/07/11/go-swagger-logging.html
// - https://baltig.infn.it/minio/console/-/blob/7fa3279a934f215983a660c9898cb38b66a4f025/restapi/configure_mcs.go
// FileServerMiddleware serves files from the static folder
func FileServerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/v1/") || strings.HasPrefix(r.URL.Path, "/api/") {
			next.ServeHTTP(w, r)
		} else {
			http.ServeFile(w, r, "web/"+r.URL.Path[1:])
		}
	})
}