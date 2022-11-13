package delivery

import (
	"net/http"

	"01.alem.school/git/Nurzhas/groupie-tracker/internal/handler"
)

type Route struct {
	mux *http.ServeMux
}

func (s *Route) InitRoute() *http.ServeMux {
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/", handler.MainHandler)
	s.mux.HandleFunc("/artist/", handler.ArtistHandler)
	s.mux.Handle("/ui/css/", http.StripPrefix("/ui/css/", http.FileServer(http.Dir("ui/css"))))
	return s.mux
}
