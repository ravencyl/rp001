package views

import (
	"html/template"
	"net/http"
)

func DashboardViewer(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("templates/dashboard.html"))
	page := Page{Title: "Create new Project"}
	tpl.Execute(w, page)
}
