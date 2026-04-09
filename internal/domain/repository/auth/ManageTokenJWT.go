package auth

import "Web-Chat/internal/domain/model"

type JwtToken interface {
	GenerateToken(User_id int, Admin bool) (string, error)
	ValidateToken(Token string) (*model.Claims, error)
}
