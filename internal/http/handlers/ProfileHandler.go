package handlers

import (
	"Web-Chat/internal/domain/service"
	"Web-Chat/internal/http/middleware"
	"net/http"
	"text/template"
)

type ProfileH struct {
	User      *service.UserService
	templates *template.Template
	AuthJwt   *middleware.JwtM
}

func NewMainPage(UserService *service.UserService, templates *template.Template, AuthJwt *middleware.JwtM) *ProfileH {
	return &ProfileH{User: UserService, templates: templates, AuthJwt: AuthJwt}
}
func (PH *ProfileH) Profile(w http.ResponseWriter, r *http.Request) {
	Claims, err := PH.AuthJwt.GetDataFromJwt(w, r)
	if err != nil {
		http.Error(w, "Unauthorized", 401)
		return
	}
	Data, err := PH.User.GetDataAboutUserForProfile(Claims.User_id)
	if err != nil {
		http.Error(w, "Unauthorized", 401)
		return
	}
	PH.templates.ExecuteTemplate(w, "Profile.html", Data)
}
func (PH *ProfileH) RedactUserTag(w http.ResponseWriter, r *http.Request) {
	CLaims, err := PH.AuthJwt.GetDataFromJwt(w, r)
	if err != nil {
		http.Error(w, "Unauthorized", 401)
		return
	}
	Usertag := r.FormValue("usertag")
	err = PH.User.RedactUserTag(Usertag, CLaims.User_id)
	if err != nil {
		http.Error(w, "Something is wrong", 500)
		return
	}
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}
func (PH *ProfileH) RedactPassword(w http.ResponseWriter, r *http.Request) {
	CLaims, err := PH.AuthJwt.GetDataFromJwt(w, r)
	if err != nil {
		http.Error(w, "Unauthorized", 401)
		return
	}
	NewPassword := r.FormValue("newpassword")
	RepeatPassword := r.FormValue("repeatpassword")
	err = PH.User.RedactPassword(NewPassword, RepeatPassword, CLaims.User_id)
	if err != nil {
		http.Error(w, "Something is wrong", 500)
		return
	}
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}
func (PH *ProfileH) RedactName(w http.ResponseWriter, r *http.Request) {
	CLaims, err := PH.AuthJwt.GetDataFromJwt(w, r)
	if err != nil {
		http.Error(w, "Unauthorized", 401)
		return
	}
	NewName := r.FormValue("newname")
	err = PH.User.RedactName(NewName, CLaims.User_id)
	if err != nil {
		http.Error(w, "Something is wrong", 500)
		return
	}
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}
