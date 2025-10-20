package handler

import (
	"cinema/pkgs/helper"
	"fmt"
	"net/http"
)

type Data struct {
	Data string `json:"data"`
	Date string `json:"date"`
}

func GetTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "service running...")
}
