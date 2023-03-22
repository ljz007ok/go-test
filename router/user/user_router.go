package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ljz007ok/go-test/action"
)

func Routers(group *gin.RouterGroup) {
	user := group.Group("/user")
	action := action.GetUserAction()
	{
		user.POST("/page", action.SelectUser)
		user.GET("/:id", action.FindUser)
		user.POST("/", action.SaveUser)
		user.PUT("/:id", action.UpdateUser)
		user.DELETE("/:id", action.DeleteUser)
	}
}
