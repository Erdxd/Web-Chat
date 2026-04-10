package middleware

import (
	"Web-Chat/internal/domain/model"
	"Web-Chat/internal/domain/repository/auth"
	"net/http"
)

type JwtM struct {
	jwt auth.JwtToken
}

func NewJwtM(jwt auth.JwtToken) *JwtM {
	return &JwtM{jwt: jwt}
}
func (JM *JwtM) GetDataFromJwt(w http.ResponseWriter, r *http.Request) (*model.Claims, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Unauthorized", 401)
		return nil, err
	}
	claims, err := JM.jwt.ValidateToken(cookie.Value)
	if err != nil {
		http.Error(w, "Unauthorized", 401)
		return nil, err
	}
	return claims, err
}
