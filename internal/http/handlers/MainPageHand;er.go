package handlers

import (
	"Web-Chat/internal/domain/service"
	"Web-Chat/internal/http/middleware"
	"net/http"
	"text/template"
)

type MainPage struct {
	User      *service.UserService
	templates *template.Template
	AuthJwt   *middleware.JwtM
}

func NewMainPage(UserService *service.UserService, templates *template.Template, AuthJwt *middleware.JwtM) *MainPage {
	return &MainPage{User: UserService, templates: templates, AuthJwt: AuthJwt}
}
func Profile(w http.ResponseWriter, r *http.Request) {

}
