package handlers

import (
	"Web-Chat/internal/domain/model"
	"Web-Chat/internal/domain/service"
	"net/http"
	"text/template"
	"time"
)

type Auth struct {
	AuthS     service.UserService
	templates *template.Template
}

func NewAuth(AuthS service.UserService, tmpl *template.Template) *Auth {
	return &Auth{AuthS: AuthS, templates: tmpl}
}
func (Au *Auth) CreateAcc(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("Name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	repeatPassword := r.FormValue("repeatpassword")
	Data := model.User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}
	err := Au.AuthS.CreateAcc(Data, repeatPassword)
	if err != nil {
		return
	}
	http.Redirect(w, r, "auth/login", http.StatusSeeOther)

}
func (Au *Auth) Login(w http.ResponseWriter, r *http.Request) {

}
