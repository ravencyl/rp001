package main

import (
	"net/http"
	"rProject/views"
)

func Router(mux *http.ServeMux) *http.ServeMux {

	mux.HandleFunc("/", views.DashboardViewer)
	// mux.HandleFunc("/upload", viewer.UploadPageViewer)
	// mux.HandleFunc("/install", viewer.InstallPageViewer)

	// fs := http.FileServer(http.Dir("asset"))
	// mux.Handle("/asset/", http.StripPrefix("/asset/", fs))
	mux.HandleFunc("/task/", views.TaskHandler)
	return mux
}
