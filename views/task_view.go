package views

import (
	"html/template"
	"log"
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

	tpl, err := template.ParseFiles(
		"./templates/page_base.html",
		"./templates/block_header.html",
		"./templates/page_task_create.html",
		"./templates/block_foot.html",
	)

	if err != nil {
		log.Fatal(err)
	}

	tpl.Execute(w, nil)
}
