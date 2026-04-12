package handlers

import (
	"Web-Chat/internal/domain/model"
	"Web-Chat/internal/domain/service"
	"Web-Chat/internal/http/middleware"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Auth struct {
	AuthS     service.UserService
	templates *template.Template
	Jwt       *service.Jwt
	Authjwt   *middleware.JwtM
}

func NewAuth(AuthS service.UserService, tmpl *template.Template, AuthJwt *middleware.JwtM, jwtService *service.Jwt) *Auth {
	return &Auth{AuthS: AuthS, templates: tmpl, Authjwt: AuthJwt, Jwt: jwtService}
}
func (Au *Auth) CreateAcc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		repeatPassword := r.FormValue("password_repeat")
		Data := model.User{
			Name:      name,
			Email:     email,
			Password:  password,
			CreatedAt: time.Now(),
		}
		log.Println(Data)
		err := Au.AuthS.CreateAcc(Data, repeatPassword)
		if err != nil {
			log.Println(err)
			http.Error(w, "Something is wrong", 400)
			return
		}
		http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
	}

	Au.templates.ExecuteTemplate(w, "register.html", nil)

}
func (Au *Auth) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		Email := r.FormValue("email")
		PasswordFromUser := r.FormValue("password")
		err := Au.AuthS.Login(Email, PasswordFromUser)

		if err != nil {
			log.Println(err)
			http.Error(w, "Something is wrong", 400)
			return
		}
		Userid, admin, err := Au.AuthS.User.GetUserDataForJWT(Email)
		if err != nil {
			log.Println(err)
			http.Error(w, "Cant Auth you", 401)
			return
		}
		token, err := Au.Jwt.CreateToken(Userid, admin)

		if err != nil {
			log.Println(err)
			return
		}
		log.Println(token)
		if token == "" {
			http.Error(w, "Unauthorized", 401)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    token,
			Path:     "/",
			MaxAge:   86400,
			HttpOnly: true,
			Secure:   false,
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	Au.templates.ExecuteTemplate(w, "login.html", nil)

}
