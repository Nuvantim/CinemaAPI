package handler

import (
	"fmt"
	"net/http"
)

func GetTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "service running...🔥")
}
