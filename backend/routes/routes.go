package routes

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jatin-dua/clear-clutter/backend/store"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../frontend/public/index.html"))
	err := tmpl.Execute(w, store.TemplateData())
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
