package router

import (
	"context"
	"github.com/ljz007ok/go-test/router/user"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

type option func(*gin.RouterGroup)

var options = []option{}

// 注册所有模块的路由配置
func include(opts ...option) {
	options = append(options, opts...)
}

// 加载各个模块的初始化
func routerInit() *gin.Engine {
	router := gin.Default()
	// 创建一个路由前缀，所有请求都以这个前缀开始，建议配置到配置文件里
	prefixRouter := router.Group("/test")

	// 所有模块都需要在这里注册
	include(user.Routers)

	// 对所有的模块的路由进行初始化
	for _, opt := range options {
		opt(prefixRouter)
	}
	return router
}

// 启动gin服务
func Gininit() {
	router := routerInit()

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    180 * time.Second,
		WriteTimeout:   180 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// 启动web服务器
	go func() {
		// 优雅地重启或停止
		// 服务连接
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
