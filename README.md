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



# gin相关集成

