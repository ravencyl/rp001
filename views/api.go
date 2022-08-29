package views

import (
	"fmt"
	"log"
	"net/http"
	"rProject/controllers"
	"strings"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {
	param := strings.TrimPrefix(r.URL.Path, "/api/")
	if param == "" {
		log.Println(param)
		fmt.Fprintln(w, "wrong path ")
		return
	}

	request_param := strings.Split(param, "/")
	if request_param[0] == "system" {
		controllers.SystemAPIHandler(request_param)
	}

	fmt.Fprintln(w, "hello ")
}

// func API_system_migration(w http.ResponseWriter, r *http.Request) {
// 	param := strings.TrimPrefix(r.URL.Path, "/API/")
// }
