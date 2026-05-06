package handlers

import (
	"anviz-mssql-api/api/httpx"
	"anviz-mssql-api/db"
	"anviz-mssql-api/models"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func GetDevicesHandler(conn *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pagination, ok := httpx.ParsePagination(w, r)
		if !ok {
			return
		}

		devices, err := db.GetDevices(conn, pagination)
		if err != nil {
			httpx.RespondError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, devices)
	}
}

func CreateDeviceHandler(conn *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var device models.Device
		if !httpx.BindJSON(w, r, &device) {
			return
		}
		if err := db.CreateDevice(conn, &device); err != nil {
			httpx.RespondError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusCreated, device)
	}
}

func UpdateDeviceHandler(conn *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			httpx.WriteError(w, http.StatusBadRequest, "Invalid ID")
			return
		}
		updates, ok := httpx.BindUpdates(w, r, deviceUpdateFields)
		if !ok {
			return
		}
		if err := db.UpdateDevice(conn, uint(id), updates); err != nil {
			httpx.RespondError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, map[string]string{"message": "Device updated"})
	}
}

func DeleteDeviceHandler(conn *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			httpx.WriteError(w, http.StatusBadRequest, "Invalid ID")
			return
		}
		if err := db.DeleteDevice(conn, uint(id)); err != nil {
			httpx.RespondError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, map[string]string{"message": "Device deleted"})
	}
}
