package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type regester struct {
	username string
	password string
}

var passage []string

func u(b regester) int {
	var usersname []string
	for j := 0; j <= 5; j++ {
		usersname = append(usersname, strconv.Itoa(j))
	}
	for i := 0; i <= 5; i++ {

		if usersname[i] == b.username {
			return 0
		}
	}
	return 1
}
func regerster(c *gin.Context) {
	var user regester
	user.username = c.Query("username")
	user.password = c.Query("password")
	a := u(user)
	if a == 0 {
		c.JSON(http.StatusOK, "注册失败")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"用户名": user.username,
			"密码":  user.password,
		})
	}
	cookie := &http.Cookie{
		Name:     user.username,
		Value:    user.password,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)
	c.Next()
}
func login(c *gin.Context) {
	cookie, err := c.Request.Cookie("user.username")
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"username": cookie.Name,
			"message":  "欢迎登录",
		})
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "您未完成注册",
		})
		c.Abort() // 不执行该请求的后续函数
	}
}
func send(c *gin.Context) {
	p := c.PostForm("passage")
	passage = append(passage, p)
	c.JSON(http.StatusOK, gin.H{
		"passage": passage,
	})
	c.Next()
}
func like(c *gin.Context) {
	c.String(http.StatusOK, "👍")

}
func delete(c *gin.Context) {
	c.String(http.StatusOK, "删除文章")
}

func main() {
	r := gin.Default()
	a := r.Group("/web")
	a.Use(regerster, login)
	{
		a.POST("/send", send, like)
		a.DELETE("/delete", delete)
	}
	err := r.Run()
	if err != nil {
		return
	}
}
