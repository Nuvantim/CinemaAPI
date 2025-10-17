package handlers

import (
	"fmt"
	"net/http"
	"cinema/pkgs/helper"
)

type Data struct{
	Data string `json:"data"`
	Date string `json:"date"`
}
func GetTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "service berjalan...")
}

func GetTulit(w http.ResponseWriter, r *http.Request){
	var error = "terjadi error"
	helper.Error(w, error)
}

func GetData(w http.ResponseWriter, r *http.Request){
	var data = Data{
		Data : "Nuvantim",
		Date : "17/10/2025",
	}
	helper.Success(w,data)
}
