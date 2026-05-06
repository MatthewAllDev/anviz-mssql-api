package httpx

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const maxJSONBodySize = 1 << 20

func BindJSON(w http.ResponseWriter, r *http.Request, target any) bool {
	r.Body = http.MaxBytesReader(w, r.Body, maxJSONBodySize)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(target); err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return false
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		WriteError(w, http.StatusBadRequest, "Request body must contain a single JSON value")
		return false
	}
	return true
}

func BindUpdates(w http.ResponseWriter, r *http.Request, allowed map[string]string) (map[string]any, bool) {
	var input map[string]any
	if !BindJSON(w, r, &input) {
		return nil, false
	}
	if len(input) == 0 {
		WriteError(w, http.StatusBadRequest, "No fields to update")
		return nil, false
	}

	updates := make(map[string]any, len(input))
	for field, value := range input {
		column, ok := allowed[field]
		if !ok {
			WriteError(w, http.StatusBadRequest, "Unknown or read-only field: "+field)
			return nil, false
		}
		updates[column] = value
	}
	return updates, true
}
