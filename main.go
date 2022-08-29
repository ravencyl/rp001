package main

import (
	"log"
	"net/http"
	"os"
	"rProject/core"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("main.main")
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	log.Println(
		"[main] start program",
	)

	core.DB()

	log.Fatal(http.ListenAndServe(":"+port, Router(mux)))
}
