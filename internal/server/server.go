package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Log        *log.Logger
	HttpServer *http.Server
}

func CreateServer(l *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.ServeIndex)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	var s http.Server
	s.Addr = ":8080"
	s.Handler = mux
	s.ErrorLog = l
	s.ReadTimeout = time.Second * 5
	s.WriteTimeout = time.Second * 10
	s.IdleTimeout = time.Second * 15

	server := Server{
		Log:        l,
		HttpServer: &s,
	}
	return &server
}
