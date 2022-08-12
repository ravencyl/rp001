package main

import (
	"net/http"
)

func Router(mux *http.ServeMux) *http.ServeMux {


	// mux.HandleFunc("/", viewer.HomePageViewer)
	// mux.HandleFunc("/upload", viewer.UploadPageViewer)
	// mux.HandleFunc("/install", viewer.InstallPageViewer)

	// fs := http.FileServer(http.Dir("asset"))
	// mux.Handle("/asset/", http.StripPrefix("/asset/", fs))
	return mux
}
