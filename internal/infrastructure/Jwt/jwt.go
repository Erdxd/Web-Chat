package jwt

import (
	"Web-Chat/internal/domain/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey []byte

type JwtToken struct {
}

func NewJwtToken() *JwtToken {
	return &JwtToken{}
}
func (J *JwtToken) GenerateToken(UserId int) (string, error) {
	ActionTime := time.Now().Add(24 * time.Hour)
	Claims := model.Claims{
		User_id: UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(ActionTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	tokenStr, err := token.SignedString(token)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
func (J *JwtToken) ValidateToken(Token string) (*model.Claims, error) {
	Claims := &model.Claims{}
	token, err := jwt.ParseWithClaims(Token, Claims, func(t *jwt.Token) (interface{}, error) { return JwtKey, nil })
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}

	return Claims, nil
}
