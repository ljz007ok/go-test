package user

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ljz007ok/go-test/action"
	"github.com/ljz007ok/go-test/global"
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

// 初始化函数，进行路由注册
func init() {
	log.Println("init")
	global.Options = append(global.Options, Routers)
}
