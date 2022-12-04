package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Middleare(c *gin.Context) {

	cookie, err := c.Request.Cookie("username")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "游客你好",
		})
		c.Abort()
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": cookie,
		})
		c.Abort() // 不执行该请求的后续函数
	}

}
func logi(c *gin.Context) {
	username := c.Query("username")
	cookie := &http.Cookie{
		Name:     username,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": username,
	})
}

func main() {
	r := gin.Default()
	r.GET("/login", Middleare, login)
	r.Run()
}
