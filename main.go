package main

import (
	"fmt"
	"net/http"

	"github.com/codedByShoe/picme/controllers"
	"github.com/codedByShoe/picme/templates"
	"github.com/codedByShoe/picme/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "layout.html", "index.html", "partials.html"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "layout.html", "contact.html", "partials.html"))))

	usersC := controllers.Users{}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "layout.html", "signup.html", "partials.html"))
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on port 3000...")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
