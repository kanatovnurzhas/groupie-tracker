package delivery

import (
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	fmt.Printf("Starting server at port 8080\nhttp://localhost:8080/\n")
	return s.httpServer.ListenAndServe()
}

// func Run(addr string, handler *http.ServeMux) error {
// 	if err := http.ListenAndServe(":8080", server.Route()); err != nil {
// 		return err
// 	}
// 	return nil
// }
