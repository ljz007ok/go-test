package action

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ljz007ok/go-test/global"
)

type userAction struct {
}

var action *userAction

// 实现单例
func GetUserAction() *userAction {
	global.Once.Do(func() {
		action = &userAction{}
	})
	return action
}

func (userAction) SelectUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello world!",
	})
}

func (userAction) FindUser(c *gin.Context) {

}

func (userAction) SaveUser(c *gin.Context) {

}

func (userAction) UpdateUser(c *gin.Context) {

}

func (userAction) DeleteUser(c *gin.Context) {

}

func init() {
	log.Println("user_action的init函数")
}
