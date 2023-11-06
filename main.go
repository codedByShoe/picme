package main

import (
	"fmt"
	"net/http"

	"github.com/codedByShoe/picme/actions"
	"github.com/codedByShoe/picme/html"
	"github.com/codedByShoe/picme/templates"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", actions.StaticHandler(html.Must(html.ParseFS(templates.FS, "index.html"))))

	r.Get("/contact", actions.StaticHandler(html.Must(html.ParseFS(templates.FS, "contact.html"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on port 3000...")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
