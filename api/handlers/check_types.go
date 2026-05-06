package handlers

import (
	"anviz-mssql-api/api/httpx"
	"anviz-mssql-api/db"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func GetCheckTypesHandler(conn *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pagination, ok := httpx.ParsePagination(w, r)
		if !ok {
			return
		}

		checkTypes, err := db.GetCheckTypes(conn, pagination)
		if err != nil {
			httpx.RespondError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, checkTypes)
	}
}

func UpdateCheckTypeHandler(conn *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			httpx.WriteError(w, http.StatusBadRequest, "Invalid ID")
			return
		}
		updates, ok := httpx.BindUpdates(w, r, checkTypeUpdateFields)
		if !ok {
			return
		}
		if err := db.UpdateCheckType(conn, uint(id), updates); err != nil {
			httpx.RespondError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, map[string]string{"message": "Check type updated"})
	}
}
