package routers

import (
	"demo-api/Controller"
	"demo-api/service"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	City string `json:"city"`
}

type Users []User

func Routes() *gin.Engine {

	jwtService := service.NewJwtAuthService()
	authService := service.NewLoginService()
	loginCtrl := Controller.NewLoginController(jwtService, authService)

	r := gin.Default()
	r.POST("login", loginCtrl.Login)

	r.GET("users", ProfileMiddleware, func(ctx *gin.Context) {
		users := Users{
			{Name: "Rahul", City: "Mumbail"},
			{Name: "Adi", City: "Pune"},
		}

		ctx.JSON(http.StatusOK, users)
	})

	return r
}

func ProfileMiddleware(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth != "" {
		authkey := os.Getenv("AUTH_TOKEN")
		fmt.Println("authkey", authkey)
		if auth != authkey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "session has expired !")
			return
		}
		token, errr := jwt.Parse(authkey, func(token *jwt.Token) (interface{}, error) {
			somval, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("Invalid token..")
			}
			fmt.Println(">>>>>", somval)
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
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid token")
	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, "Authorization header is required")
}

// func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
// 		authorizationHeader := req.Header.Get("authorization")
// 		if authorizationHeader != "" {
// 			bearerToken := strings.Split(authorizationHeader, " ")
// 			if len(bearerToken) == 2 {
// 				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
// 					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 						return nil, fmt.Errorf("There was an error")
// 					}
// 					return []byte("secret"), nil
// 				})
// 				if error != nil {
// 					json.NewEncoder(w).Encode(error.Error())
// 					return
// 				}
// 				if token.Valid {
// 					context.Set(req, "decoded", token.Claims)
// 					next(w, req)
// 				} else {
// 					json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
// 				}
// 			}
// 		} else {
// 			json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
// 		}
// 	})
// }
