package main

import (
	"fmt"
	"login/loginpage"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Server start")
	server := gin.Default()
	server.LoadHTMLFiles("template/html/login.html")
	server.Static("/assets", "template/assets")
	server.GET("/login", loginpage.LoginPage)
	//server.POST("/login", )

	server.Run(":8080")
}
