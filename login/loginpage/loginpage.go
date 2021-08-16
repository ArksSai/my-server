package loginpage

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Account struct {
	id               int
	email            string
	phone            string
	username         string
	password         string
	create_at        int
	create_ip_at     string
	last_login_at    int
	last_login_ip_at string
	login_times      int
	status           int
}

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
			"error": errors.New("password!"),
		})
		return
	}

	golang_db.Connect()

}

func CheckUserIsExist(username string) bool {
	_, isExist := UserData[username]
	return isExist
}

func CheckPassword(p1 string, p2 string) error {
	if p1 == p2 {
		return nil
	} else {
		return errors.New("password is wrong.")
	}
}

func Auth(username string, password string) error {
	if isExist := CheckUserIsExist(username); isExist {
		return CheckPassword(UserData[username], password)
	} else {
		return errors.New("user is not exist")
	}
}
