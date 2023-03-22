package router

import (
	"log"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options = []Option{}

// 注册所有模块的路由配置
func include(opts ...Option) {
	options = append(options, opts...)
}

// 加载各个模块的初始化
func routerInit() *gin.Engine {
	r := gin.Default()
	for _, opt := range options {
		opt(r)
	}
	return r
}

// 启动gin服务
func Gininit() {
	// 加载多个APP的路由配置
	include(shop.Routers, blog.Routers)
	// 初始化路由
	r := routerInit()

	// 启动web服务器
	go func() {
		// 优雅实现关机和重启
		// 默认endless服务器会监听下列信号：
		// syscall.SIGHUP，syscall.SIGUSR1，syscall.SIGUSR2，syscall.SIGINT，syscall.SIGTERM和syscall.SIGTSTP
		// 接收到 SIGHUP 信号将触发`fork/restart` 实现优雅重启（kill -1 pid会发送SIGHUP信号）
		// 接收到 syscall.SIGINT或syscall.SIGTERM 信号将触发优雅关机
		// 接收到 SIGUSR2 信号将触发HammerTime
		// SIGUSR1 和 SIGTSTP 被用来触发一些用户自定义的hook函数
		err := endless.ListenAndServe(":8888", r)
		if err != nil {
			log.Fatalf("listen: %s\n", err)
		} else {
			log.Println("Server exiting")
		}
	}()
}
