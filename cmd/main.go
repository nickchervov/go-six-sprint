package main

import (
	"log"
	"net/http"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	var logger log.Logger
	s := server.CreateServer(&logger)

	if err := http.ListenAndServe(s.HttpServer.Addr, s.HttpServer.Handler); err != nil {
		s.Log.Fatal(err)
	}
}
