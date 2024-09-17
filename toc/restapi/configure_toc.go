// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
  "log"
  "time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
  "github.com/go-openapi/strfmt"

	"github.com/gyk4j/toc/toc/restapi/operations"
	"github.com/gyk4j/toc/toc/restapi/operations/archive"
	"github.com/gyk4j/toc/toc/restapi/operations/backup"
	logops "github.com/gyk4j/toc/toc/restapi/operations/log"
	"github.com/gyk4j/toc/toc/restapi/operations/quota"
	"github.com/gyk4j/toc/toc/restapi/operations/restoration"
	"github.com/gyk4j/toc/toc/restapi/operations/synchronization"
  
  "github.com/gyk4j/toc/toc/models"
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

	if api.ArchiveArchiveDataHandler == nil {
		api.ArchiveArchiveDataHandler = archive.ArchiveDataHandlerFunc(func(params archive.ArchiveDataParams) middleware.Responder {
			return middleware.NotImplemented("operation archive.ArchiveData has not yet been implemented")
		})
	}
	if api.LogExportLogHandler == nil {
		api.LogExportLogHandler = logops.ExportLogHandlerFunc(func(params logops.ExportLogParams) middleware.Responder {
			return middleware.NotImplemented("operation log.ExportLog has not yet been implemented")
		})
	}
	if api.ArchiveGetArchiveHandler == nil {
		api.ArchiveGetArchiveHandler = archive.GetArchiveHandlerFunc(func(params archive.GetArchiveParams) middleware.Responder {
			return middleware.NotImplemented("operation archive.GetArchive has not yet been implemented")
		})
	}
	//if api.ArchiveGetArchiveByIDHandler == nil {
		api.ArchiveGetArchiveByIDHandler = archive.GetArchiveByIDHandlerFunc(func(params archive.GetArchiveByIDParams) middleware.Responder {
			api.Logger("Hello %d!", params.ArchiveID)
      //return middleware.NotImplemented("xxx operation archive.GetArchiveByID has not yet been implemented")
      a := models.Archive{
        ID: params.ArchiveID,
        Status: "in-progress",
        Time: strfmt.DateTime(time.Now()),
        URL: "/v1/archives/",
      }
      res := archive.NewGetArchiveByIDOK()
      res.SetPayload(&a)
      return res
		})
	//}
  
	if api.BackupGetBackupHandler == nil {
		api.BackupGetBackupHandler = backup.GetBackupHandlerFunc(func(params backup.GetBackupParams) middleware.Responder {
			return middleware.NotImplemented("operation backup.GetBackup has not yet been implemented")
		})
	}
	if api.BackupGetBackupByIDHandler == nil {
		api.BackupGetBackupByIDHandler = backup.GetBackupByIDHandlerFunc(func(params backup.GetBackupByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation backup.GetBackupByID has not yet been implemented")
		})
	}
	if api.QuotaGetQuotaHandler == nil {
		api.QuotaGetQuotaHandler = quota.GetQuotaHandlerFunc(func(params quota.GetQuotaParams) middleware.Responder {
			return middleware.NotImplemented("operation quota.GetQuota has not yet been implemented")
		})
	}
	if api.RestorationGetRestorationHandler == nil {
		api.RestorationGetRestorationHandler = restoration.GetRestorationHandlerFunc(func(params restoration.GetRestorationParams) middleware.Responder {
			return middleware.NotImplemented("operation restoration.GetRestoration has not yet been implemented")
		})
	}
	if api.RestorationGetRestorationByIDHandler == nil {
		api.RestorationGetRestorationByIDHandler = restoration.GetRestorationByIDHandlerFunc(func(params restoration.GetRestorationByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation restoration.GetRestorationByID has not yet been implemented")
		})
	}
	if api.BackupNewBackupHandler == nil {
		api.BackupNewBackupHandler = backup.NewBackupHandlerFunc(func(params backup.NewBackupParams) middleware.Responder {
			return middleware.NotImplemented("operation backup.NewBackup has not yet been implemented")
		})
	}
	if api.RestorationNewRestorationHandler == nil {
		api.RestorationNewRestorationHandler = restoration.NewRestorationHandlerFunc(func(params restoration.NewRestorationParams) middleware.Responder {
			return middleware.NotImplemented("operation restoration.NewRestoration has not yet been implemented")
		})
	}
	if api.SynchronizationNewSynchronizationHandler == nil {
		api.SynchronizationNewSynchronizationHandler = synchronization.NewSynchronizationHandlerFunc(func(params synchronization.NewSynchronizationParams) middleware.Responder {
			return middleware.NotImplemented("operation synchronization.NewSynchronization has not yet been implemented")
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
