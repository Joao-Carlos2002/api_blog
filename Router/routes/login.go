package routes

import (
	"log"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `form:"username" json:"username"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func LoginPost(c *gin.Context) {
	var user User
	err := c.ShouldBind(&user)
	if err != nil {
		panic(err.Error())
	}
	log.Println(user)
}
