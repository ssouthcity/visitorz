package main

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed templates/*.html
var content embed.FS
var tmpl = template.Must(template.ParseFS(content, "**/*"))

func main() {
	visitorStore := visitorStoreFactory()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var visitors int

		if visitedCookieExists(r) {
			visitors = visitorStore.Visitors()
		} else {
			visitors = visitorStore.Increment()
			writeVisitedCookie(w)
		}

		tmpl.ExecuteTemplate(w, "index.html", visitors)
	})

	http.ListenAndServe(":5000", nil)
}
