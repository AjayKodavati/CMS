package main

import (
	"log"

	"github.com/AjayKodavati/CMS/config"
	"github.com/AjayKodavati/CMS/db"
	"github.com/AjayKodavati/CMS/repository"
)



func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	err = db.InitDB()
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}
	defer db.Pool.Close()

	repositoryService := repository.SetUpDBRepositories(db.Pool)
	
}
