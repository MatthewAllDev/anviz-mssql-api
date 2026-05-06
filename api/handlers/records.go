package handlers

import (
	"anviz-mssql-api/api/httpx"
	"anviz-mssql-api/db"
	"anviz-mssql-api/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

func GetRecordsHandler(conn *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pagination, ok := httpx.ParsePagination(w, r)
		if !ok {
			return
		}

		recordIDStr := r.URL.Query().Get("record_id")
		var recordID *uint
		if recordIDStr != "" {
			id, err := strconv.ParseUint(recordIDStr, 10, 32)
			if err != nil {
				httpx.WriteError(w, http.StatusBadRequest, "Invalid record_id")
				return
			}
			rid := uint(id)
			recordID = &rid
		}

		userIDsStr := r.URL.Query().Get("user_ids")
		var userIDs []string
		if userIDsStr != "" {
			userIDs = strings.Split(userIDsStr, ",")
		}

		startTimeStr := r.URL.Query().Get("start_time")
		endTimeStr := r.URL.Query().Get("end_time")
		var startTime, endTime *time.Time
		if startTimeStr != "" {
			t, err := time.Parse(time.RFC3339, startTimeStr)
			if err != nil {
				httpx.WriteError(w, http.StatusBadRequest, "Invalid start_time")
				return
			}
			startTime = &t
		}
		if endTimeStr != "" {
			t, err := time.Parse(time.RFC3339, endTimeStr)
			if err != nil {
				httpx.WriteError(w, http.StatusBadRequest, "Invalid end_time")
				return
			}
			endTime = &t
		}

		records, err := db.GetRecords(conn, recordID, userIDs, startTime, endTime, pagination)
		if err != nil {
			httpx.RespondError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, records)
	}
}

func GetRecordsSinceHandler(conn *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pagination, ok := httpx.ParsePagination(w, r)
		if !ok {
			return
		}

		afterIDStr := r.URL.Query().Get("after_id")
		if afterIDStr == "" {
			httpx.WriteError(w, http.StatusBadRequest, "after_id is required")
			return
		}
		afterID, err := strconv.ParseUint(afterIDStr, 10, 32)
		if err != nil {
			httpx.WriteError(w, http.StatusBadRequest, "Invalid after_id")
			return
		}

		records, err := db.GetRecordsSince(conn, uint(afterID), pagination)
		if err != nil {
			httpx.RespondError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, records)
	}
}

func CreateRecordHandler(conn *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var record models.Record
		if !httpx.BindJSON(w, r, &record) {
			return
		}
		if err := db.CreateRecord(conn, &record); err != nil {
			httpx.RespondError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusCreated, record)
	}
}

func UpdateRecordHandler(conn *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			httpx.WriteError(w, http.StatusBadRequest, "Invalid ID")
			return
		}
		updates, ok := httpx.BindUpdates(w, r, recordUpdateFields)
		if !ok {
			return
		}
		if err := db.UpdateRecord(conn, uint(id), updates); err != nil {
			httpx.RespondError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, map[string]string{"message": "Record updated"})
	}
}

func DeleteRecordHandler(conn *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			httpx.WriteError(w, http.StatusBadRequest, "Invalid ID")
			return
		}
		if err := db.DeleteRecord(conn, uint(id)); err != nil {
			httpx.RespondError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, map[string]string{"message": "Record deleted"})
	}
}
