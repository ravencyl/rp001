package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	param := strings.TrimPrefix(r.URL.Path, "/task/")

	if param == "create" {
		CreateTask(w)
	}

	if _, err := strconv.Atoi(param); err == nil {
		fmt.Printf("%q looks like a number.\n", param)
	}

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
