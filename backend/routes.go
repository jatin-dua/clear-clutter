package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome"))
}

func SetupRoutes(port int) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", serveHome)

	fmt.Printf("Application is running on http://localhost:%v \n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), r))

}
