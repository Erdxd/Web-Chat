package main

import (
	"Web-Chat/internal/domain/service"
	"Web-Chat/internal/http/handlers"
	http1 "Web-Chat/internal/http/ws"

	hasher "Web-Chat/internal/infrastructure/Hasher"
	infrastructure "Web-Chat/internal/infrastructure/database"
	"Web-Chat/internal/repositories"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

func main() {
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	tmpl, err := template.ParseGlob("web/templates/*.html")
	if err != nil {
		panic(err)
	}

	err = godotenv.Load("database.env")
	if err != nil {
		log.Println(err)
	}
	UrlDb := os.Getenv("DATABASE_URL")
	db, err := infrastructure.InitDb(UrlDb)
	if err != nil {
		log.Println(err)
	}
	log.Println(UrlDb)
	log.Println(db)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/index.html")
	})
	hub := http1.NewHub()
	go hub.Run()
	MessageRepo := repositories.NewRepo(db)
	serviceM := service.NewServiceMessage(MessageRepo)
	handlerMain := http1.NewChatHandler(serviceM, hub, tmpl)
	UserRepo := repositories.NewUserRepo(db)
	Hasher := hasher.NewHasher()
	ServiceU := service.NewUserService(UserRepo, Hasher)
	HandlerUser := handlers.NewAuth(*ServiceU, tmpl)

	http.HandleFunc("/ws", handlerMain.OpenPipe)
	http.HandleFunc("/auth/login", HandlerUser.Login)
	http.HandleFunc("/auth/register", HandlerUser.CreateAcc)
	http.ListenAndServe(":8080", nil)
	log.Println("localhost:8080")
}
