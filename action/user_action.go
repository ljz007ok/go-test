package action

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type userAction struct {
}

var instance *userAction
var once sync.Once

// 实现单例
func GetUserAction() *userAction {
	once.Do(func() {
		instance = &userAction{}
	})
	return instance
}

func SelectUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello world!",
	})
}

func FindUser(c *gin.Context) {

}

func SaveUser(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
