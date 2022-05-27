package controllers

import (
	"microservice/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(c *gin.Context)
	FetchDataController(c *gin.Context)
}

type loginController struct {
	JwtAuthenticationService service.JwtAuthService
	LoginAuthService         service.LoginService
	TestSvc                  service.TestService
}

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func NewLoginAuthController(jServic service.JwtAuthService, lService service.LoginService, testsvc service.TestService) LoginController {
	return &loginController{
		JwtAuthenticationService: jServic,
		LoginAuthService:         lService,
		TestSvc:                  testsvc,
	}
}

func (lc *loginController) Login(c *gin.Context) {
	req := LoginRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	isAuthenticate := lc.LoginAuthService.Authenticate(req.UserName, req.Password)
	if isAuthenticate {
		accessToken := lc.JwtAuthenticationService.GenerateToken(req.UserName, req.Password)
		// setting up env variable
		os.Setenv("access_token", accessToken)
		c.JSON(http.StatusOK, gin.H{"token": accessToken})
		return
	}
	//lets see the demo
	c.JSON(http.StatusUnauthorized, "authentication failed")
}

type Requestbody struct {
	Fsyms string `form:"fsyms"`
	Tsyms string `form:"tsyms"`
}

func (lc *loginController) FetchDataController(c *gin.Context) {
	var req Requestbody
	err := c.ShouldBindQuery(&req)
	//fmt.Println("fsyms : ", err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	if req.Fsyms == "" || req.Tsyms == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "fsyms should not empty"})
		return
	}

	resp, err := lc.TestSvc.FetchData(req.Fsyms, req.Tsyms)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, resp)
}
