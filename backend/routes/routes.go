package routes

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jatin-dua/clear-clutter/backend/app"
	"github.com/jatin-dua/clear-clutter/backend/store"
)

var data *store.Data

func RefreshData() {
	data = store.TemplateData()
	installedApps := app.InstalledApps(*data.Device)
	store.SetTemplateData(&installedApps, data.Device)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	RefreshData()
	tmpl := template.Must(template.ParseFiles("../frontend/public/index.html"))
	err := tmpl.Execute(w, store.TemplateData())
	if err != nil {
		log.Fatal(err)
	}
}

func processForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	fmt.Println("## Form-Data ##")
	for key, applications := range r.Form {
		fmt.Printf("%v = %v\n", key, applications)
		app.UninstallApps(applications)
	}

	http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
}

func SetupRoutes(port int) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Serve static files from the "static" directory
	staticDir := "../frontend/static"
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	r.Get("/", serveHome)
	r.Post("/", processForm)

	fmt.Printf("Application is running on http://localhost:%v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), r))
}
