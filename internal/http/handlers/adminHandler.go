package handlers

import (
	"Web-Chat/internal/domain/service"
	"net/http"
	"strconv"
	"text/template"
)

type AdminHandler struct {
	AdminService service.AdminSerivce
	template     *template.Template
}

func NewAdminHandler(Admin service.AdminSerivce, template *template.Template) *AdminHandler {
	return &AdminHandler{AdminService: Admin, template: template}
}
func (AH *AdminHandler) CheckAllUsers(w http.ResponseWriter, r *http.Request) {
	AllUsers, err := AH.AdminService.CheckAllUsers()
	if err != nil {
		http.Error(w, "Cant select users", 500)
		return
	}
	AH.template.ExecuteTemplate(w, "admin.html", AllUsers)

}
func (AH *AdminHandler) FoundByUserId(w http.ResponseWriter, r *http.Request) {
	UserId, err := strconv.Atoi(r.FormValue("useridfound"))
	if err != nil {
		http.Error(w, "Something is worng", 500)
		return
	}
	user, err := AH.AdminService.FoundUserByUserId(UserId)
	if err != nil {
		http.Error(w, "Something is wrong", 500)
		return
	}
	AH.template.Execute(w, user)
}
func (AH *AdminHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	UserId, err := strconv.Atoi(r.FormValue("useridfound"))
	if err != nil {
		http.Error(w, "Cant found user with this ID", 204)
		return
	}
	err = AH.AdminService.DeleteUser(UserId)
	if err != nil {
		http.Error(w, "Something is wrong", 500)
		return
	}
	http.Redirect(w, r, "Admin/CheckUsers", http.StatusSeeOther)
}
