package views

import (
	"html/template"
	"log"
	"net/http"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {

	DatabaseMigration(w)
	// param := strings.TrimPrefix(r.URL.Path, "/admin/")
	// fmt.Println(param)

	// if param == "" {
	// fmt.Println("empty param, go default page")
	// }
}

func DatabaseMigration(w http.ResponseWriter) {
	tpl, err := template.ParseFiles(
		"./templates/admin/page_admin_base.html",
		"./templates/admin/block_slidebar.html",
		"./templates/admin/block_header.html",
		"./templates/admin/block_system.html",
	)

	if err != nil {
		log.Fatal(err)
	}

	tpl.Execute(w, Page{Title: "Admin-System"})
}
