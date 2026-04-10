package service

import (
	"Web-Chat/internal/domain/repository/auth"
	jwt "Web-Chat/internal/infrastructure/Jwt"
)

type Jwt struct {
	Jwt auth.JwtToken
}

func NewJwt(jwt auth.JwtToken) *Jwt {
	return &Jwt{Jwt: jwt}
}
func (J *Jwt) CreateToken(userid int, admin bool) (string, error) {
	token, err := jwt.NewJwtToken().GenerateToken(userid, admin)
	if err != nil {
		return "", nil
	}
	return token, nil
}
