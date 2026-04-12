package service

import (
	"Web-Chat/internal/domain/repository/auth"
	"log"
)

type Jwt struct {
	Jwt auth.JwtToken
}

func NewJwt(jwt auth.JwtToken) *Jwt {
	return &Jwt{Jwt: jwt}
}
func (J *Jwt) CreateToken(userid int, admin bool) (string, error) {
	token, err := J.Jwt.GenerateToken(userid, admin)
	if err != nil {
		return "", err
	}
	log.Println(token)
	return token, err
}
