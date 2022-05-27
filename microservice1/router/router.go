package router

import (
	"fmt"
	"log"
	"microservice/configuration"
	"microservice/controllers"
	"microservice/executor"
	"microservice/service"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	City  string `json:"city"`
}

type Users []User

func SetRoutes() *gin.Engine {
	routes := gin.Default()
	_, err := configuration.LoadConfig()
	if err != nil {
		log.Fatalln("unable to load config")
	}

	jService := service.NewAuthJwtService()
	lService := service.NewLoginService()

	testExecutor := executor.NewTestExecutor()
	testService := service.NewTestService(testExecutor)
	lController := controllers.NewLoginAuthController(jService, lService, testService)

	routes.GET("users", ProfileMiddleware, func(c *gin.Context) {
		userData := Users{
			{Name: "Rahul", Email: "rahul@abc.com", City: "Mumbai"},
			{Name: "Joe", Email: "joe@abc.com", City: "NewYork"},
			{Name: "Jonny", Email: "jonny@abc.com", City: "Tokiyo"},
			{Name: "Sam", Email: "Sam@abc.com", City: "Paris"},
		}

		c.JSON(http.StatusOK, userData)
	})
	// below api will create new jwt token
	routes.POST("authenticate", lController.Login)

	routes.GET("service/price", lController.FetchDataController)

	return routes
}

func ProfileMiddleware(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth != "" {
		authKey := os.Getenv("access_token")
		if auth != authKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "session expired ...!")
			return
		}
		token, errr := jwt.Parse(authKey, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("wrong token")
			}
			return []byte("secret"), nil
		})
		if errr != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errr.Error()})
			return
		}
		if token.Valid {
			c.Next()
			return
		}

	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, "auth header is required")
}
