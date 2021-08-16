package login

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/golang_db"
)

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginAuth(c *gin.Context) {
	// receive account and password
	var (
		account  string
		password string
	)

	//get account
	if in, isExist := c.GetPostForm("username"); isExist && in != "" {
		account = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("Fill the user name!"),
		})
		return
	}
	//get password
	if in, isExist := c.GetPostForm("password"); isExist && in != "" {
		password = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("Fill the user password!"),
		})
		return
	}

	if golang_db.CheckPassword(account, password) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"success": "Login Success!!",
		})
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("Password incorrect."),
		})
	}

}
