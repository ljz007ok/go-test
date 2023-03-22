# go-test
​		项目采用go1.19，使用go mod方式，使用gin和gorm框架集成环境。

# 项目搭建步骤

## 创建github项目

````shell
https://github.com/ljz007ok/go-test.git
````

## 修改项目.gitignore文件

​		增加如下内容，此文件添加不提交代码的文件。根据具体情况，添加相应的文件标识。

````
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with `go test -c`
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/

.idea
.vscode
*/.DS_Store
main.exe
main
go.sum
````

## github项目下载到本地

````shell
git clone https://github.com/ljz007ok/go-test.git
````

## 初始化项目

````shell
# 进入项目目录
cd go-test

# 初始化项目
go mod init github.com/ljz007ok/go-test

# 执行初始化命令后，会生成go.mod文件
````

## 引入gin-swag依赖

````shell
# 在任何目录下执行下面命令，但是不能在项目目录下，也就是不能在go.mod文件的目录下执行。
go install github.com/swaggo/swag/cmd/swag

# 会在gopath/bin目录下生成swag可执行文件
````

## 引入gin依赖

````shell
# 进入项目目录,如果不加版本，则默认用releases的latest版本，也就是releases的最新版本，目前最新版本就是v1.9.0
go get -u github.com/gin-gonic/gin@v1.9.0
````

## 引入gorm依赖

````shell
# 进入项目目录
# 这里是gorm2，如果gorm1则是go get -u github.com/jinzhu/gorm
# gorm在github上没有提供releases版本，那就下载最新的。
go get -u gorm.io/gorm
# 数据库驱动，不同数据库修改最后的名称
go get -u gorm.io/driver/mysql
````

## 引入gin参数校验依赖

````shell
go get -u github.com/go-playground/validator/v10
````

## 引入gin优雅关机和重启依赖

````shell
go get -u github.com/fvbock/endless
````

## 引入gin的日志zap

​		[Zap](https://github.com/uber-go/zap)是非常快的、结构化的，分日志级别的Go日志库。它同时提供了结构化日志记录和printf风格的日志记录。

​		golang提供的Logger没有日志级别，只有一个`Print`选项。不支持`INFO`/`DEBUG`等多个级别。对于错误日志，它有`Fatal`和`Panic`。Fatal日志通过调用`os.Exit(1)`来结束程序。Panic日志在写入日志消息之后抛出一个panic。但是它缺少一个ERROR日志级别，这个级别可以在不抛出panic或退出程序的情况下记录错误。缺乏日志格式化的能力——例如记录调用者的函数名和行号，格式化日期和时间格式。等等。不提供日志切割的能力。

​		

````shell
go get -u go.uber.org/zap
````

## 引入Viper配置文件管理依赖

````shell
go get -u github.com/spf13/viper
````



# gin相关集成

​		`gin.Default()`用到了gin框架内的两个默认中间件`Logger()`和`Recovery()`。

​		其中`Logger()`是把gin框架本身的日志输出到标准输出（我们本地开发调试时在终端输出的那些日志就是它的功劳），而`Recovery()`是在程序出现panic的时候恢复现场并写入500响应的。

