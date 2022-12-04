package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type register struct {
	username string
	password string
}

func U(b register) int {
	var usersname []string
	for j := 0; j <= 5; j++ {
		usersname = append(usersname, string(j))
	}
	for i := 0; i <= 5; i++ {

		if usersname[i] == b.username {
			return 0
		}
	}
	return 1
}
func main() {

	r := gin.Default()
	r.GET("/register", func(c *gin.Context) {
		var user register
		user.username = c.Query("username")
		user.password = c.Query("password")
		a := U(user)
		if a == 0 {
			c.JSON(http.StatusOK, "注册失败")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"用户名": user.username,
				"密码":  user.password,
			})
		}
	})
	r.Run()
}
