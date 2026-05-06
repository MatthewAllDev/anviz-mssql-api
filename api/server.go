package api

import (
	"anviz-mssql-api/api/auth"
	"anviz-mssql-api/api/handlers"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
)

func SetupRouter(conn *gorm.DB, apiKeys *auth.APIKeyStore) http.Handler {
	mux := http.NewServeMux()
	handle := func(pattern string, scope auth.Scope, handler http.HandlerFunc) {
		mux.Handle(pattern, auth.Authorize(apiKeys, scope, handler))
	}

	handle("GET /api/v1/users", auth.ReadScope, handlers.GetUsersHandler(conn))
	handle("GET /api/v1/departments", auth.ReadScope, handlers.GetDepartmentsHandler(conn))
	handle("GET /api/v1/devices", auth.ReadScope, handlers.GetDevicesHandler(conn))
	handle("GET /api/v1/records", auth.ReadScope, handlers.GetRecordsHandler(conn))
	handle("GET /api/v1/records/since", auth.ReadScope, handlers.GetRecordsSinceHandler(conn))
	handle("GET /api/v1/check_types", auth.ReadScope, handlers.GetCheckTypesHandler(conn))

	handle("POST /api/v1/users", auth.CrudScope, handlers.CreateUserHandler(conn))
	handle("PUT /api/v1/users/{id}", auth.CrudScope, handlers.UpdateUserHandler(conn))
	handle("DELETE /api/v1/users/{id}", auth.CrudScope, handlers.DeleteUserHandler(conn))

	handle("POST /api/v1/departments", auth.CrudScope, handlers.CreateDepartmentHandler(conn))
	handle("PUT /api/v1/departments/{id}", auth.CrudScope, handlers.UpdateDepartmentHandler(conn))
	handle("DELETE /api/v1/departments/{id}", auth.CrudScope, handlers.DeleteDepartmentHandler(conn))

	handle("POST /api/v1/devices", auth.CrudScope, handlers.CreateDeviceHandler(conn))
	handle("PUT /api/v1/devices/{id}", auth.CrudScope, handlers.UpdateDeviceHandler(conn))
	handle("DELETE /api/v1/devices/{id}", auth.CrudScope, handlers.DeleteDeviceHandler(conn))

	handle("POST /api/v1/records", auth.CrudScope, handlers.CreateRecordHandler(conn))
	handle("PUT /api/v1/records/{id}", auth.CrudScope, handlers.UpdateRecordHandler(conn))
	handle("DELETE /api/v1/records/{id}", auth.CrudScope, handlers.DeleteRecordHandler(conn))

	handle("PUT /api/v1/check_types/{id}", auth.CrudScope, handlers.UpdateCheckTypeHandler(conn))

	return mux
}

func StartServer(conn *gorm.DB, apiKeys *auth.APIKeyStore, port string) {
	r := SetupRouter(conn, apiKeys)
	server := &http.Server{
		Addr:              ":" + port,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("Starting server on port %s", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}
