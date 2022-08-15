package views

import (
	"html/template"
	"net/http"
	"strings"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	param := strings.TrimPrefix(r.URL.Path, "/task/")

	if param == "create" {
		CreateTask(w)
	}
	//  strings.TrimPrefix(req.URL.Path, "/provisions/"))
}

func CreateTask(w http.ResponseWriter) {
	var tpl = template.Must(template.ParseFiles("templates/page_task_create.html"))
	page := Page{
		Title: "Create new task",
	}

	tpl.Execute(w, page)
}
