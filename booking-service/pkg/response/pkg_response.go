package response

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Error(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	if _, err := io.WriteString(w, err.Error()); err != nil {
		io.WriteString(w, fmt.Sprintf("failed get error response : %s", err.Error()))
	}
}

func Success[T any](w http.ResponseWriter, data T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
