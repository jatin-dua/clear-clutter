package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Data struct {
	Packages *[]string
}

var data *Data

func SetData(d *[]string) {
	data = &Data{
		Packages: d,
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../frontend/public/index.html"))
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

func SetupRoutes(port int) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", serveHome)

	fmt.Printf("Application is running on http://localhost:%v \n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), r))
}
