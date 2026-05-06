package handlers

import (
	"anviz-mssql-api/api/httpx"
	"anviz-mssql-api/db"
	"anviz-mssql-api/models"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func GetUsersHandler(conn *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pagination, ok := httpx.ParsePagination(w, r)
		if !ok {
			return
		}

		userID := r.URL.Query().Get("user_id")
		var userIDFilter *string
		if userID != "" {
			userIDFilter = &userID
		}

		departmentIDStr := r.URL.Query().Get("department_id")
		var departmentID *uint
		if departmentIDStr != "" {
			id, err := strconv.ParseUint(departmentIDStr, 10, 32)
			if err != nil {
				httpx.WriteError(w, http.StatusBadRequest, "Invalid department_id")
				return
			}
			did := uint(id)
			departmentID = &did
		}

		users, err := db.GetUsers(conn, userIDFilter, departmentID, pagination)
		if err != nil {
			httpx.RespondError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, users)
	}
}

func CreateUserHandler(conn *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if !httpx.BindJSON(w, r, &user) {
			return
		}
		if err := db.CreateUser(conn, &user); err != nil {
			httpx.RespondError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusCreated, user)
	}
}

func UpdateUserHandler(conn *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		updates, ok := httpx.BindUpdates(w, r, userUpdateFields)
		if !ok {
			return
		}
		if err := db.UpdateUser(conn, id, updates); err != nil {
			httpx.RespondError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, map[string]string{"message": "User updated"})
	}
}

func DeleteUserHandler(conn *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if err := db.DeleteUser(conn, id); err != nil {
			httpx.RespondError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, map[string]string{"message": "User deleted"})
	}
}
