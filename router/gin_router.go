package router

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ljz007ok/go-test/global"
)

// 注册所有模块的路由配置,注册路由，放在每个功能模块的路由里进行注册
//func include(opts ...global.Option) {
//	global.Options = append(global.Options, opts...)
//}

// 加载各个模块的初始化
func routerInit() *gin.Engine {
	// 强制日志颜色化
	gin.ForceConsoleColor()

	router := gin.Default()
	// 创建一个路由前缀，所有请求都以这个前缀开始，建议配置到配置文件里
	prefixRouter := router.Group("/test")

	// 对所有的模块已经注册的路由进行初始化
	for _, opt := range global.Options {
		opt(prefixRouter)
	}
	return router
}

// 启动gin服务
func Gininit() {
	router := routerInit()

	server := &http.Server{
		Addr:           ":8801",
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
	// 忽略os.Interrupt报错，用vscode是没有异常的，应该是跟编辑器有关
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

func init() {
	log.Println("gin_router执行init函数")
}
