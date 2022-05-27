package Controller

import (
	"demo-api/dto"
	"demo-api/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(c *gin.Context)
}

type loginController struct {
	JwtService   service.JwtService
	LoginService service.LoginService
}

func NewLoginController(jwtService service.JwtService, loginService service.LoginService) LoginController {
	return loginController{
		JwtService:   jwtService,
		LoginService: loginService,
	}
}

func (l loginController) Login(c *gin.Context) {
	userRequest := dto.LoginRequest{}
	err := c.ShouldBind(&userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
	}
	isAuthenticate := l.LoginService.LoginUser(userRequest.Username, userRequest.Password)
	if isAuthenticate {
		token := l.JwtService.CreateToken(userRequest.Username, userRequest.Password)
		os.Setenv("AUTH_TOKEN", token)
		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}
	c.JSON(http.StatusUnauthorized, "authentication failed")
}
