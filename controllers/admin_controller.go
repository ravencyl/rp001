package controllers

import (
	"log"
	"rProject/core"
	"rProject/models"
)

func SystemAPIHandler(request_param []string) {
	log.Println("controllers.SystemAPIHandler")

	if request_param[1] == "migrate" {
		migrate_database()
	}
}

func migrate_database() {
	db := core.DB()
	db.AutoMigrate(&models.Task{})
}
