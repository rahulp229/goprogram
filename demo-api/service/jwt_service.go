package service

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JwtService interface {
	CreateToken(user, email string) string
}

type authClaims struct {
	Name  string
	Email string
	jwt.StandardClaims
}

type jwtService struct {
	SecretKey string
	Issure    string
}

func NewJwtAuthService() JwtService {
	return &jwtService{
		SecretKey: "secret",
		Issure:    "Rahul",
	}
}

func (auth *jwtService) CreateToken(user, email string) string {
	claim := authClaims{
		Name:  "rahul",
		Email: "rahul@gmail.com",
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: time.Now().Add(time.Second * 360).Unix(),
			Id:        user,
			IssuedAt:  0,
			Issuer:    "rahul",
			NotBefore: 0,
			Subject:   "",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	//encoded string
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		panic(err)
	}
	return t
}
