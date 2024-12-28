package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService struct {
	secretKey []byte
}

func NewJWTService(secretKey []byte) *JWTService {
	return &JWTService{secretKey: secretKey}
}

func (s *JWTService) GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	return token.SignedString(s.secretKey)
}

func (s *JWTService) ValidateToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return s.secretKey, nil
	})
}
