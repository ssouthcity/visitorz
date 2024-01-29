package main

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed templates/*.html
var content embed.FS
var tmpl = template.Must(template.ParseFS(content, "**/*"))

type PageData struct {
	VisitorNumber int
	TotalVisitors int
}

func main() {
	visitorStore := visitorStoreFactory()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		visitorNumber, hasVisited := readVisitorCookie(r)
		if !hasVisited {
			visitorNumber = visitorStore.Increment()
			writeVisitorCookie(w, visitorNumber)
		}

		totalVisitors := visitorStore.Visitors()

		tmpl.ExecuteTemplate(w, "index.html", PageData{
			VisitorNumber: visitorNumber,
			TotalVisitors: totalVisitors,
		})
	})

	http.ListenAndServe(":5000", nil)
}
