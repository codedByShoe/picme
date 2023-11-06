package actions

import (
	"net/http"

	"github.com/codedByShoe/picme/html"
)

func StaticHandler(tpl html.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
