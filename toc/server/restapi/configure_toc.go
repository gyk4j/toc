// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/gyk4j/toc/toc/server/restapi/operations"
	"github.com/gyk4j/toc/toc/server/restapi/operations/archive"
	"github.com/gyk4j/toc/toc/server/restapi/operations/backup"
	logops "github.com/gyk4j/toc/toc/server/restapi/operations/log"
	"github.com/gyk4j/toc/toc/server/restapi/operations/quota"
	"github.com/gyk4j/toc/toc/server/restapi/operations/restoration"
	"github.com/gyk4j/toc/toc/server/restapi/operations/synchronization"

	"github.com/gyk4j/toc/toc/server/services"
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

	api.ArchiveArchiveDataHandler = archive.ArchiveDataHandlerFunc(func(params archive.ArchiveDataParams) middleware.Responder {
		services.ArchiveData()
		res := archive.NewArchiveDataOK()
		return res
	})

	api.LogExportLogHandler = logops.ExportLogHandlerFunc(func(params logops.ExportLogParams) middleware.Responder {
		var res middleware.Responder

		l := services.NewLog()
		if l != nil {
			res = logops.NewExportLogOK().WithPayload(l)
		} else {
			res = logops.NewExportLogInternalServerError()
		}

		return res
	})

	api.ArchiveGetArchiveHandler = archive.GetArchiveHandlerFunc(func(params archive.GetArchiveParams) middleware.Responder {
		var res middleware.Responder

		a := services.GetArchives()
		if a != nil {
			res = archive.NewGetArchiveOK().WithPayload(a)
		} else {
			res = archive.NewGetArchiveNotFound()
		}

		return res
	})

	api.ArchiveGetArchiveByIDHandler = archive.GetArchiveByIDHandlerFunc(func(params archive.GetArchiveByIDParams) middleware.Responder {
		var res middleware.Responder

		a := services.GetArchiveByID(params.ArchiveID)
		if a != nil {
			res = archive.NewGetArchiveByIDOK()
		} else {
			res = archive.NewGetArchiveByIDNotFound()
		}

		return res
	})

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
