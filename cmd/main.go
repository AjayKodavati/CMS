package main

import (
	"log"
	"os"

	"github.com/AjayKodavati/CMS/config"
	"github.com/AjayKodavati/CMS/db"
	"github.com/AjayKodavati/CMS/router"
	"github.com/AjayKodavati/CMS/server"
)



func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	pool, err := db.InitDB()
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}
	defer pool.Close()
	
	server := server.NewServer(pool)
	router.SetupRouter(server)
	
	port := os.Getenv("PORT")
	server.Start(port)
}
