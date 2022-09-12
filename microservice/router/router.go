package router

import (
	"log"
	"microservice/configuration"
	"microservice/controllers"
	"microservice/database"
	"microservice/executor"
	"microservice/service"
	"net/http"

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
	client := database.ConnectDb()
	//ctx := gin.Context{}
	testExecutor := executor.NewTestExecutor()
	testService := service.NewTestService(testExecutor)
	lController := controllers.NewLoginAuthController(testService)

	userController := controllers.NewUsers(client)
	//userRoutes := routes.Group("api/", userController.GetUsers)

	routes.GET("service/price", lController.FetchDataController)
	routes.GET("health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "success")
	})
	routes.GET("users", userController.GetUsers)
	return routes
}
