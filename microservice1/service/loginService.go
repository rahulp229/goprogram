package service

type LoginService interface {
	Authenticate(user, password string) bool
}

type loginService struct {
	User     string
	Password string
}

func NewLoginService() LoginService{
	// here i hard code the credential we can fetch from db or some thing else
	return &loginService{
		User: "user",
		Password: "password",
	}
}

func (ls *loginService) Authenticate(user, password string) bool  {
	return ls.User == user && ls.Password == password
}