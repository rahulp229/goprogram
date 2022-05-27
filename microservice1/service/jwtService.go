package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtAuthService interface {
	GenerateToken(user, email string) string
}

type authClaims struct {
	User  string
	Email string
	jwt.StandardClaims
}

type jwtService struct {
	SecretKey string
	Issuer    string
}

func NewAuthJwtService() JwtAuthService {
	return &jwtService{
		SecretKey: "secret",
		Issuer:    "Rahul",
	}
}

func (js *jwtService) GenerateToken(user, email string) string {
	claims := authClaims{
		User:  user,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: time.Now().Add(time.Second * 30).Unix(), // token expiry timeout
			Id:        "",
			IssuedAt:  0,
			Issuer:    js.Issuer,
			NotBefore: 0,
			Subject:   "",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(js.SecretKey))
	// token is generated
	return t
}
