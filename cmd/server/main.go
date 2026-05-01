package main

import (
	"Web-Chat/internal/domain/service"
	"Web-Chat/internal/http/handlers"
	"Web-Chat/internal/http/middleware"
	http1 "Web-Chat/internal/http/ws"

	hasher "Web-Chat/internal/infrastructure/Hasher"
	jwt "Web-Chat/internal/infrastructure/Jwt"
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

	err = godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
	UrlDb := os.Getenv("DATABASE_URL")
	if UrlDb == "" {
		panic("App was started without Database_URL")
	}
	db, err := infrastructure.InitDb(UrlDb)
	if err != nil {
		log.Println(err)
	}

	hub := http1.NewHub()
	go hub.Run()

	jwt_token := os.Getenv("JWT_TOKEN")
	if jwt_token == "" {
		panic("App was started without JWT-token")
	}
	jwt_token2 := []byte(jwt_token)
	jwtservice := jwt.NewJwtToken(jwt_token2)
	jwtService2 := service.NewJwt(jwtservice)

	jwtMiddleware := middleware.NewJwtM(jwtservice)

	MessageRepo := repositories.NewRepo(db)
	serviceM := service.NewServiceMessage(MessageRepo)

	UserRepo := repositories.NewUserRepo(db)
	Hasher := hasher.NewHasher()
	ServiceU := service.NewUserService(UserRepo, Hasher)
	HandlerUser := handlers.NewAuth(*ServiceU, tmpl, jwtMiddleware, jwtService2)
	AdminRepo := repositories.NewAdminRepo(db)
	AdminService := service.NewAdminService(AdminRepo)
	RoomRepo := repositories.NewRoomRepo(db)
	RoomService := service.NewRoomService(RoomRepo, UserRepo, ServiceU)
	AdminHandler := handlers.NewAdminHandler(*AdminService, tmpl)
	handlerMain := http1.NewChatHandler(serviceM, hub, tmpl, jwtMiddleware, *ServiceU, *RoomService)
	ProfileHandler := handlers.NewProfileH(ServiceU, tmpl, jwtMiddleware)

	http.HandleFunc("/ws", handlerMain.OpenPipe)
	http.HandleFunc("/", jwtMiddleware.VerifUser(handlerMain.GetAllPrivateChats))
	http.HandleFunc("/auth/login", HandlerUser.Login)
	http.HandleFunc("/auth/register", HandlerUser.CreateAcc)
	http.HandleFunc("/ws/delete", handlerMain.DeleteMessage)
	http.HandleFunc("/admin/CheckUsers", jwtMiddleware.VerifAdmin(AdminHandler.CheckAllUsers))
	http.HandleFunc("/admin/FoundUser", jwtMiddleware.VerifAdmin(AdminHandler.FoundByUserId))
	http.HandleFunc("/admin/deleteUser", jwtMiddleware.VerifAdmin(AdminHandler.DeleteUser))
	http.HandleFunc("/profile", ProfileHandler.Profile)
	http.HandleFunc("/profile/redact/tag", ProfileHandler.RedactUserTag)
	http.HandleFunc("/profile/redact/name", ProfileHandler.RedactName)
	http.HandleFunc("/profile/redact/password", ProfileHandler.RedactPassword)
	http.HandleFunc("/users/search", handlerMain.FindUserByUserTag)
	http.HandleFunc("/chats/private", handlerMain.CreatePrivateChat)

	http.ListenAndServe(":8080", nil)
	log.Println("localhost:8080")
}
