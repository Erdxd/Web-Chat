package jwt

import (
	"Web-Chat/internal/domain/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtToken struct {
	Key []byte
}

func NewJwtToken(key []byte) *JwtToken {
	return &JwtToken{Key: key}
}
func (J *JwtToken) GenerateToken(UserId int, admin bool) (string, error) {
	ActionTime := time.Now().Add(24 * time.Hour)
	Claims := model.Claims{
		User_id: UserId,
		Admin:   admin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(ActionTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	tokenStr, err := token.SignedString(J.Key)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
func (J *JwtToken) ValidateToken(Token string) (*model.Claims, error) {
	Claims := &model.Claims{}
	token, err := jwt.ParseWithClaims(Token, Claims, func(t *jwt.Token) (interface{}, error) { return J.Key, nil })
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}

	return Claims, nil
}
