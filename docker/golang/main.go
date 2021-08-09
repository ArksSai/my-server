package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

//func test(context *gin.Context) {
//	context.HTML(http.StatusOK, "index.html", nil)
//}
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginAuth(c *gin.Context) {
	var (
		username string
		password string
	)

	if in, isExist := c.GetPostForm("username"); isExist && in != "" {
		username = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("user name!"),
		})
		return
	}
	if in, isExist := c.GetPostForm("password"); isExist && in != "" {
		password = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("password!"),
		})
		return
	}
	if err := Auth(username, password); err == nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"success": "Success!!",
		})
		return
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": err,
		})
		return
	}

}

func index_page(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func main() {
	server := gin.Default()
	//server.LoadHTMLGlob("template/html/*")
	server.LoadHTMLFiles("index.html", "template/html/login.html")
	server.Static("/assets", "template/assets")
	server.GET("/login", LoginPage)
	server.POST("/login", LoginAuth)
	//server.GET("/", index_page)
	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello Metafalica"})
	})
	server.Run(":8088")

}
