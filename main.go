package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/codedByShoe/picme/actions"
	"github.com/codedByShoe/picme/html"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	// Parse templates
	tpl, err := html.Parse(filepath.Join("templates", "index.html"))
	if err != nil {
		panic(err)
	}
	r.Get("/", actions.StaticHandler(tpl))

	tpl, err = html.Parse(filepath.Join("templates", "contact.html"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", actions.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on port 3000...")
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
