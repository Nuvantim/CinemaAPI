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
	fmt.Fprintln(w, "service berjalan...")
}

func GetData(w http.ResponseWriter, r *http.Request) {
	var data = Data{
		Data: "Nuvantim",
		Date: "17/10/2025",
	}
	helper.Success(w, data)
}
