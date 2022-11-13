package handler

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"01.alem.school/git/Nurzhas/groupie-tracker/internal/models"
	"01.alem.school/git/Nurzhas/groupie-tracker/internal/service"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Err(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		Err(w, http.StatusMethodNotAllowed)
		return
	}

	artists := []models.Artists{}
	if service.ConvertJson(models.ArtistsURL, &artists) != nil {
		Err(w, http.StatusInternalServerError)
		return
	}
	temp, err := template.ParseFiles("ui/html/index.html")
	if err != nil {
		Err(w, http.StatusInternalServerError)
		return
	}

	if err = temp.Execute(w, artists); err != nil {
		Err(w, http.StatusInternalServerError)
		return
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		Err(w, http.StatusMethodNotAllowed)
		return
	}

	temp, err := template.ParseFiles("ui/html/artist.html")
	if err != nil {
		Err(w, http.StatusInternalServerError)
		return
	}

	artists := []models.Artists{}
	if service.ConvertJson(models.ArtistsURL, &artists) != nil {
		Err(w, http.StatusInternalServerError)
		return
	}

	idTemp := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(idTemp)
	if err != nil || id > len(artists) || id <= 0 {
		Err(w, http.StatusNotFound)
		return
	}

	if artists[id-1].InfoConcert() != nil {
		Err(w, http.StatusInternalServerError)
		return
	}

	if err = temp.Execute(w, artists[id-1]); err != nil {
		Err(w, http.StatusInternalServerError)
		return
	}
}

func Err(w http.ResponseWriter, code int) {
	serveErrors(w, models.InfoErr{http.StatusText(code), code})
}
