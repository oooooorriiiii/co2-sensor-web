package controllers

import (
	"co2-sensor-web/parser"
	"fmt"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseFiles("app/views/google.html"))

func viewChartHandler(w http.ResponseWriter, r *http.Request) {
	df := parser.SqlParser("co2ex.sqlite")

	err := templates.ExecuteTemplate(w, "google.html", df)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StartWebServer() error {
	http.HandleFunc("/chart/", viewChartHandler)
	return http.ListenAndServe(fmt.Sprintf(":%d", 8000), nil)
}
