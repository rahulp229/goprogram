package controllers

import (
	"microservice/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type User interface {
	GetUsers(c *gin.Context)
}

type user struct {
	db *database.DBConnection
}

func NewUsers(db *database.DBConnection) User {
	return &user{db: db}
}

//GetUsers
func (u *user) GetUsers(c *gin.Context) {
	collection := u.db.Client.Database(u.db.DBName).Collection("user")
	cur, err := collection.Find(c, bson.D{})
	if err != nil {
		c.JSON(http.StatusForbidden, err.Error())
		return
	}
	s := ""
	err = cur.All(c, s)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"sucess": s})
}
