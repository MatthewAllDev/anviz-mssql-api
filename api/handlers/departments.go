package handlers

import (
	"anviz-mssql-api/api/httpx"
	"anviz-mssql-api/db"
	"anviz-mssql-api/models"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func GetDepartmentsHandler(conn *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pagination, ok := httpx.ParsePagination(w, r)
		if !ok {
			return
		}

		departments, err := db.GetDepartments(conn, pagination)
		if err != nil {
			httpx.RespondError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, departments)
	}
}

func CreateDepartmentHandler(conn *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dept models.Department
		if !httpx.BindJSON(w, r, &dept) {
			return
		}
		if err := db.CreateDepartment(conn, &dept); err != nil {
			httpx.RespondError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusCreated, dept)
	}
}

func UpdateDepartmentHandler(conn *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			httpx.WriteError(w, http.StatusBadRequest, "Invalid ID")
			return
		}
		updates, ok := httpx.BindUpdates(w, r, departmentUpdateFields)
		if !ok {
			return
		}
		if err := db.UpdateDepartment(conn, uint(id), updates); err != nil {
			httpx.RespondError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, map[string]string{"message": "Department updated"})
	}
}

func DeleteDepartmentHandler(conn *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			httpx.WriteError(w, http.StatusBadRequest, "Invalid ID")
			return
		}
		if err := db.DeleteDepartment(conn, uint(id)); err != nil {
			httpx.RespondError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, map[string]string{"message": "Department deleted"})
	}
}
