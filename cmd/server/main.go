package main

import (
	"Web-Chat/internal/domain/service"
	http1 "Web-Chat/internal/http/ws"
	infrastructure "Web-Chat/internal/infrastructure/database"
	"Web-Chat/internal/repositories"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
	UrlDb := os.Getenv("DATABASE_URL")
	db, err := infrastructure.InitDb(UrlDb)
	if err != nil {
		log.Println(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/index.html")
	})
	hub := http1.NewHub()
	go hub.Run()
	MessageRepo := repositories.NewRepo(db)
	serviceM := service.NewServiceMessage(MessageRepo)
	handlerMain := http1.NewChatHandler(serviceM, hub)
	http.HandleFunc("/ws", handlerMain.OpenPipe)
	http.ListenAndServe(":8080", nil)
	log.Println("localhost:8080")
}
