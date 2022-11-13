package handler

import (
	"net/http"
	"text/template"

	"01.alem.school/git/Nurzhas/groupie-tracker/internal/models"
)

func serveErrors(w http.ResponseWriter, Err models.InfoErr) {
	temp, err := template.ParseFiles("ui/html/error.html")
	w.WriteHeader(Err.StatusNumber)

	if err != nil {
		http.Error(w, Err.StatusNumberText, Err.StatusNumber)
		return
	}

	if err := temp.Execute(w, Err); err != nil {
		http.Error(w, Err.StatusNumberText, Err.StatusNumber)
		return
	}
}
