package service

type LoginService interface {
	LoginUser(user ,email string) bool
}

type loginService struct {
	User string
	Password string
}

func NewLoginService() LoginService{
	return &loginService{
		User:  "user",
		Password: "password",
	}
}

func (auth *loginService) LoginUser(user, password string) bool  {
	return auth.User == user && auth.Password == password
}


