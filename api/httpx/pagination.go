package httpx

import (
	"anviz-mssql-api/db"
	"net/http"
	"strconv"
)

const (
	defaultLimit = 100
	maxLimit     = 1000
)

func ParsePagination(w http.ResponseWriter, r *http.Request) (db.Pagination, bool) {
	limit := defaultLimit
	if value := r.URL.Query().Get("limit"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil || parsed < 1 || parsed > maxLimit {
			WriteError(w, http.StatusBadRequest, "Invalid limit")
			return db.Pagination{}, false
		}
		limit = parsed
	}

	offset := 0
	if value := r.URL.Query().Get("offset"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil || parsed < 0 {
			WriteError(w, http.StatusBadRequest, "Invalid offset")
			return db.Pagination{}, false
		}
		offset = parsed
	}

	w.Header().Set("X-Limit", strconv.Itoa(limit))
	w.Header().Set("X-Offset", strconv.Itoa(offset))
	return db.Pagination{Limit: limit, Offset: offset}, true
}
